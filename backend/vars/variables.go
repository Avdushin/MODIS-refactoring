package vars

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var (
	//@ app
	VERSION = "1.3"
	//@ Получение переменных окружения из файла .env
	_      = godotenv.Load()
	PORT   = os.Getenv("PORT")
	DBUser = os.Getenv("DB_USER")
	DBPass = os.Getenv("DB_PASS")
	DBName = os.Getenv("DB_NAME")
	DBHost = os.Getenv("DB_HOST")
	DBPort = os.Getenv("DB_PORT")
	//@ DB connection...
	DBConn = fmt.Sprintf("%s:%s@tcp(%s:%s)/", DBUser, DBPass, DBHost, DBPort)
	//@ Пути до SSL/TLS сертификатов
	Cert = "/var/www/certs/domain-example.ru.pub"
	Key  = "/var/www/private/domain-example.ru.key"
	//@ Routing
	Router      *mux.Router
	CorsHandler http.Handler
)
