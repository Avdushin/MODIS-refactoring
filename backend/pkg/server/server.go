package server

import (
	"net/http"

	"github.com/avdushin/rgoauth/pkg/logger"
	"github.com/avdushin/rgoauth/vars"
)

// @ Функци StartServer
// ? Запускает Dev сервер по адресу http://localhost:4000
func StartServer(port string) error {
	logger.INFO("Сервер запущен по адресу", vars.DBHost, port)
	err := http.ListenAndServe(port, vars.CorsHandler)
	if err != nil {
		return err
	}
	return nil
}

// @ Функци StartProdServer
// ? Запускает Production сервер по адресу https://localhost:4000
func StartProdServer(port, cert, key string) error {
	logger.INFO("Сервер запущен по адрессу", vars.DBHost, port)
	//@ Запуск сервера с поддержкой SSL (HTTPS) соединения (Production mode)
	/* Атрибуты:
	*  1. Номер порта, на котором будет запущен сервер
	*  2. SSL сертификат, выданный хостингом
	*  3. SSL ключ сертификата, выданный хостингом
	 */
	err := http.ListenAndServeTLS(vars.PORT, vars.Cert, vars.Key, vars.CorsHandler)
	if err != nil {
		return err
	}
	return nil
}
