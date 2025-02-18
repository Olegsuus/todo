.PHONY: build run docker-build docker-run migrate compose compose-migrate compose-reset

POSTGRES_HOST ?= localhost
POSTGRES_PORT ?= 5432
POSTGRES_DB ?= todo_db
POSTGRES_USER ?= postgres
POSTGRES_PASSWORD ?= postgres

build:
	@echo "Сборка приложения todo..."
	go build -o app ./cmd/app

run:
	@echo "Запуск приложения todo с конфигом: ./configs/local.yaml"
	./app

docker-build:
	@echo "Сборка Docker-образа todo..."
	docker build -t todo .

docker-run:
	@echo "Запуск Docker-контейнера todo..."
	docker run -p 5555:5555 todo

migrate:
	@echo "Применение миграций (локально)"
	goose -dir migrations postgres "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable" up

compose-migrate:
	@echo "Применяем миграции в контейнере"
	docker run --rm --network todo_default todo \
	  goose -dir migrations postgres "postgres://postgres:postgres@postgres:5432/todo_db?sslmode=disable" up

compose:
	@echo "Поднимаем окружение todo через docker-compose..."
	docker-compose up --build

compose-reset:
	docker-compose down -v && docker-compose up --build