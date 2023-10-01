package main

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/avdushin/rgoauth/pkg/logger"
	"github.com/avdushin/rgoauth/pkg/models/database"
	"github.com/avdushin/rgoauth/pkg/routes"
	"github.com/avdushin/rgoauth/pkg/server"
	"github.com/avdushin/rgoauth/pkg/utils"
	v "github.com/avdushin/rgoauth/vars"
)

// ? Функции, необходимые при запуске программы - запускаются однократно
func init() {
	//@ Инициализация логгера
	logger.InitLogger()
	//@ Инициализация таблиц БД
	database.InitTables()
	//@ Бэкап БД, который запускается каждые 4 часа
	go database.RunDatabaseBackup(4)
	//@ Инициализация роутера и настройка роутов
	v.Router = routes.SetupRouter()
	//@ corsHandler
	v.CorsHandler = utils.CreateCorsHandler(v.Router)

	// ? Если переменная `PORT` не была определена в файле `.env`, то по-умолчанию сервер будет запущен на порту :4000
	if v.PORT == "" {
		v.PORT = ":4000"
	}
}

func main() {
	// Запуск сервера без SSL (DEV)
	server.StartServer(v.PORT)
	// Запуск сервера с поддержкой SSL (PROD)
	// server.StartProdServer(v.PORT, v.Cert, v.Key)
}
