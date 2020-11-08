
# VARIABLES
# -


# CONFIG
.PHONY: help print-variables
.DEFAULT_GOAL := help


# ACTIONS

## infra

run-pushgateway :		## Run Prometheus Pushgateway
	docker run -d --name prom-pushgateway -p 9091:9091 prom/pushgateway:v1.3.0

open-pushgateway-ui :		## Open Prometheus Pushgateway
	open http://localhost:9091

show-pushgateway-metrics :		## Show Prometheus Pushgateway exposed metrics
	open http://localhost:9091/metrics

## app

build-collector :		## Build application with collector
	cd collector/ && make build

build-gatherer :		## Build application with gatherer
	cd gatherer/ && make build

build-all : build-collector build-gatherer		## Build all applications

run-collector : build-collector		## Run application with collector
	cd collector/ && make run

run-gatherer : build-gatherer		## Run application with gatherer
	cd gatherer/ && make run

run-all : run-collector run-gatherer		## Run all applications

## helpers

help :		## Help
	@echo ""
	@echo "*** \033[33mMakefile help\033[0m ***"
	@echo ""
	@echo "Targets list:"
	@grep -E '^[a-zA-Z_-]+ :.*?## .*$$' $(MAKEFILE_LIST) | sort -k 1,1 | awk 'BEGIN {FS = ":.*?## "}; {printf "\t\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	@echo ""

print-variables :		## Print variables values
	@echo ""
	@echo "*** \033[33mMakefile variables\033[0m ***"
	@echo ""
	@echo "- - - makefile - - -"
	@echo "MAKE: $(MAKE)"
	@echo "MAKEFILES: $(MAKEFILES)"
	@echo "MAKEFILE_LIST: $(MAKEFILE_LIST)"
	@echo "- - -"
	@echo ""
