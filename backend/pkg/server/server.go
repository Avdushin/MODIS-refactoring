package server

import (
	"net/http"

	"github.com/avdushin/rgoauth/pkg/logger"
	"github.com/avdushin/rgoauth/vars"
)

// @ Запуск Dev сервера по адресу http://localhost:4000
func StartServer(port string) {
	logger.INFO("Сервер запущен по адрессу", vars.DBHost, port)
	logger.FATAL(http.ListenAndServe(port, vars.CorsHandler))
}

func StartProdServer(port, cert, key string) {
	logger.INFO("Сервер запущен по адрессу", vars.DBHost, port)
	//@ Запуск сервера с поддержкой SSL (HTTPS) соединения (Production mode)
	/* Атрибуты:
	*  1. Номер порта, на котором будет запущен сервер
	*  2. SSL сертификат, выданный хостингом
	*  3. SSL ключ сертификата, выданный хостингом
	 */
	logger.FATAL(http.ListenAndServeTLS(vars.PORT, vars.Cert, vars.Key, vars.CorsHandler))
}
