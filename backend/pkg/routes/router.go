package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// @ Роутер
func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	//@ Обработчик OPTIONS запросов
	//? Обрабатывает запросы на регистрацию и авторизацию
	router.Methods(http.MethodOptions).HandlerFunc(handleOptions)

	//? Вызов функций из пакета routes для настройки роутов
	// @ Роуты авторизации
	SetupAuthRoutes(router)
	//@ Пользовательские роуты
	SetupUserRoutes(router)

	return router
}

// @ Разрешенные функции
func handleOptions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}
