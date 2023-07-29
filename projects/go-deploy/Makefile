up:
	docker-compose up -d

down:
	docker-compose down

run:
	go run main.go

exec:
	docker-compose exec godeploy bash

build:
	GOOS=linux go build -o server .

.PHONY: up, down, run, exec, build