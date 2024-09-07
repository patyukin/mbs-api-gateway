.PHONY: start stop rebuild gen up down restart

up:
	docker compose up -d

down:
	docker compose down

start:
	docker compose start

stop:
	docker compose stop

restart:
	docker compose restart

rebuild:
	docker compose down -v --remove-orphans
	docker compose up -d --build

gen:
	go run github.com/swaggo/swag/cmd/swag@latest init -g cmd/api_gateway/main.go -o ./docs/
