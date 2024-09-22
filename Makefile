.PHONY: build up down logs

build:
	go build -v ./cmd/time-booking

up:
	docker-compose up -d --build

down:
	docker-compose down

logs:
	docker-compose logs -f

.DEFAULT_GOAL := up