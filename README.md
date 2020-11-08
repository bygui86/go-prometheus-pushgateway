
# go-prometheus-pushgateway

Sample projec to understand how to use Prometheus Pushgateway in Golang

## Run

1. Start Prometheus Pushgateway
```shell script
make run-pushgateway
```

2. Run application
```shell script
make run-all
```

3. Open Prometheus Pushgateway UI
```shell script
make open-pushgateway-ui
```

4. Show Prometheus Pushgateway exposed metrics
```shell script
make show-pushgateway-metrics
```

## Links

- https://prometheus.io/docs/practices/pushing/
- https://prometheus.io/docs/instrumenting/pushing/
- https://github.com/prometheus/pushgateway
- https://pkg.go.dev/github.com/prometheus/client_golang/prometheus/push
