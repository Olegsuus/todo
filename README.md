Это REST API для управления задачами (TODO-лист).
Проект реализован на Go с использованием Fiber, PostgreSQL и миграций через goose. Также предусмотрена документация API с использованием Swagger.

Быстрый старт

1. Клонирование репозитория
        

      git clone https://github.com/your_username/todo.git
      cd todo

2. Настройка конфигурации

* После клонирования переименуйте файл .env.example в .env и отредактируйте его:
1) CONFIG_PATH – путь до файла конфигурации (например, ./configs/local.yaml).
2) POSTGRES_HOST, POSTGRES_PORT, POSTGRES_DB, POSTGRES_USER, POSTGRES_PASSWORD – параметры подключения к базе данных.

* Переименуйте файл configs/local.yaml.example в configs/local.yaml и отредактируйте:
1) Укажите нужный порт для приложения (например, 5555).
2) Задайте путь до файла логов (например, logs/ingo.log).
3) Другие необходимые настройки для окружения.

3. Запуск проекта через Docker Compose

Соберите и запустите все контейнеры (PostgreSQL, приложение):

      make compose

4. Применение миграций

После того, как контейнеры подняты, примените миграции к базе данных:

      make compose-migrate

5. Документация API

Swagger документация доступна по адресу:
      
      http://localhost:{port}tasks/swagger/index.html#/


Где {port} — это порт, указанный в вашем configs/local.yaml (например, 5555).
Если документация не открывается сразу, проверьте, что сервер запущен и маршрутизация для Swagger настроена.



- Команды Makefile
1)	make build – сборка приложения (локально).
2)	make run – запуск приложения (локально).
3)	make docker-build – сборка Docker-образа.
4)	make docker-run – запуск Docker-контейнера.
5)	make migrate – применение миграций локально.
6)	make compose – запуск окружения через docker-compose.
7)	make compose-migrate – применение миграций в контейнерной базе.
8)	make compose-reset – сброс окружения и повторный запуск.

- Структура проекта
1)	cmd/app – точка входа в приложение.
2)	configs/ – файлы конфигураций.
3)  local.yaml – конфигурация для локального запуска.
4) migrations/ – SQL-миграции (goose).