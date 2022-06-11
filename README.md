# goweatherbot

Погодный Telegram бот на golang. Умеет показывать погоду по введенному городу, работает с openweahermap

### Установка
Скопировать .env.example в .env и задать там ключи

    OPENWEATHER_API=  
    BOT_TOKEN=

Для запуска бота локально

    go run ./cmd/bot .
    
Для запуска бота через Docker

    docker-compose up --build
