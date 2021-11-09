# О проекте

Телеграм бот, который парсит youtube каналы и находит последнее видео и отправляет его в группу.


## Запуск приложения

1. Создать в `config/` файл config.yaml и добавить данные:
```yaml
TELEGRAM_API: <telegram_api>
TELEGRAM_TOKEN: <telegram_token>
TELEGRAM_CHAT_ID: <chat_id>
YOUTUBE_SEARCH_URL: https://www.googleapis.com/youtube/v3/search
YOUTUBE_VIDEO_URL: https://www.youtube.com/watch?v=
YOUTUBE_API_TOKEN: <youtube_api_token>
DATABASE_URL: <connect_url>
```

2. Запустить файл `cmd/bot/main.go` с помощью команды `go run ./bot/main.go`


## Структура приложения
```
    cmd/bot/main.go         # Файл запуска 
    config/                 # Конфиги
    internal
        bot/bot.go          # Функции для работы с ботом
        bot/youtube.go      # Функции для работы youtube
        database        
            database.go     # Подключение к БД
            service.go      # Функции для работы с БД
    models/models.go        # Модели
    .gitignore
    README.md
```