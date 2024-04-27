# Car catalog
## Библиотеки
- github.com/gin-gonic/gin - Gin Web Framework
- github.com/go-pg/pg/v10 - PostgreSQL client and ORM for Golang
- github.com/joho/godotenv - GoDotEnv
- github.com/sirupsen/logrus - Logrus
- github.com/swaggo/swag/cmd/swag

## config.env
| cvar               |     value      |           description           |
|:-------------------|:--------------:|:-------------------------------:|
| DB_ADDRESS         | localhost:5432 |              host               |
| DB_USER            | postgres |              user               |
| DB_PASSWORD        | root |              pass               |
| DB_NAME            | cartest |            database             |
| HOST_URL           | localhost:8080 |              Адрес              |
| HOST_RELATIVE_PATH | /cars |              Путь               |
| EXTERNAL_API_URL   | http://localhost:8083/info |       Адрес внешнего API        |
| LOG_LEVEL          | localhost:5432 | Уровень лога `info` или `debug` |
