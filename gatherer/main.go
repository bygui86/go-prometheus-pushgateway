package main

import (
	"fmt"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

var (
	completionTime = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "db_backup_last_completion_timestamp_seconds",
		Help: "The timestamp of the last completion of a DB backup, successful or not.",
	})
	duration = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "db_backup_duration_seconds",
		Help: "The duration of the last DB backup in seconds.",
	})
	records = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "db_backup_records_processed",
		Help: "The number of records processed in the last DB backup.",
	})
	/*
		successTime is added to pusher only in case of successful backup.
		We could as well register it with all others
		This example demonstrates that you can mix Gatherers and Collectors when handling a Pusher.
	*/
	successTime = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "db_backup_last_success_timestamp_seconds",
		Help: "The timestamp of the last successful completion of a DB backup.",
	})
)

func main() {
	fmt.Println("Start prom-pushgateway db-backup")

	// We use Prometheus register function to benefit from the consistency checks that happen during registration.
	prometheus.MustRegister(completionTime, duration, records)

	pusher :=
		push.New("http://localhost:9091", "db-backup").
			Gatherer(prometheus.DefaultGatherer).
			Grouping("env", "local")

	start := time.Now()
	n, backErr := performBackup()

	completionTime.SetToCurrentTime()
	duration.Set(time.Since(start).Seconds()) // Note that time.Since only uses a monotonic clock in Go1.9+.
	records.Set(float64(n))

	if backErr != nil {
		fmt.Println("DB backup failed:", backErr)
	} else {
		successTime.SetToCurrentTime()
		pusher.Collector(successTime)
	}

	// Add is used here rather than Push to not delete a previously pushed
	// success timestamp in case of a failure of this backup.
	err := pusher.Add()

	var result string
	if err != nil {
		result = fmt.Sprintf("ERROR - %s", err.Error())
	} else {
		result = "SUCCESS"
	}

	fmt.Printf("prom-pushgateway db-backup completed: %s \n", result)
}

// Perform the backup and return the number of backed up records and any applicable error.
func performBackup() (int, error) {
	time.Sleep(1 * time.Second)
	return 42, nil
}
