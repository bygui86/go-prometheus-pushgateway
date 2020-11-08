package main

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

func main() {
	fmt.Println("Start prom-pushgateway local-test")

	completionTime := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "local_test_last_completion",
		Help: "The timestamp of the last successful completion of a local test.",
	})
	completionTime.SetToCurrentTime()

	err :=
		push.New("http://localhost:9091", "local-test").
			Collector(completionTime).
			Grouping("env", "local").
			Grouping("phase", "test").
			Push()
	// INFO: alternative to .Push(), we could use .Add()
	// Add works like push, but only previously pushed metrics with the same name
	// (and the same job and other grouping labels) will be replaced.

	var result string
	if err != nil {
		result = fmt.Sprintf("ERROR - %s", err.Error())
	} else {
		result = "SUCCESS"
	}

	fmt.Printf("prom-pushgateway local-test completed: %s \n", result)
}
