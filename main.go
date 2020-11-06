package main

import (
	"fmt"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

func main() {
	fmt.Println("Start prom-pushgateway local-test")

	pusher := push.New("http://localhost:9091", "local-test")

	time.Sleep(3 * time.Second)

	completionTime := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "local_test_last_completion",
		Help: "The timestamp of the last successful completion of a local test.",
	})
	completionTime.SetToCurrentTime()

	pusher.
		Collector(completionTime).
		Grouping("env", "local").
		Grouping("phase", "test")

	err := pusher.Push()

	var result string
	if err != nil {
		result = fmt.Sprintf("ERROR - %s", err.Error())
	} else {
		result = "SUCCESS"
	}

	fmt.Printf("prom-pushgateway local-test completed: %s \n", result)
}
