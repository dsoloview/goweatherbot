# goweatherbot

Погодный Telegram бот на golang. 

### Установка
Скопировать .env.example в .env и задать там

    OPENWEATHER_API=  
    BOT_TOKEN=

Для запуска бота локально

    go run ./cmd/bot .
    
Для запуска бота через Docker

    docker-compose up --build
