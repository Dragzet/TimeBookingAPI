.PHONY: build up down

build:
	go build -v ./cmd/time-booking

up:
	docker-compose up -d --build

down:
	docker-compose down

.DEFAULT_GOAL := up