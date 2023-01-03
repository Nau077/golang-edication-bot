# Бот по теории языка go, систем-дизайну и генерации задачек

### Функции бота

- Бот доступен по url https://t.me/GolangEdicationBot
- Команда /start активирует главное меню
- В меню "задачки" генерируется рандомная задача, кнопка с ответом прилагается
- Предложить задачи и контент для тем можете в issues здесь https://github.com/Nau077/golang-edication-bot

### Сделано с помощью:

- [![golang][golang]][https://go.dev/]
- [![tg-bot-api][tg-bot-api]][https://go-telegram-bot-api.dev/]
- [![postgresql][postgresql]][https://www.postgresql.org/]

### Как запустить локально

- установить docker и docker-compose
- создать config.json в ./static по образцу и подобию ./static/config.dist
- убедиться, что установлен go 1.19
- вызвать make run/db
- вызвать make run/bot
