.PHONY: build

build:
	go build -v ./cmd/time-booking
.DEFAULT_GOAL := build