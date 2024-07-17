.PHONY: build
build:
	@go build -o ./bin/weather-cli

.PHONY: run
run: build
	@./bin/weather-cli 
