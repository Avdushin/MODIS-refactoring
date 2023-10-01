package routes

import (
	"net/http"

	"github.com/avdushin/rgoauth/pkg/handlers"
	"github.com/gorilla/mux"
)

// @ Роуты для работы с сервером
func SetupServerRoutes(router *mux.Router) {
	router.HandleFunc("/", handlers.RegisterHandler).Methods(http.MethodPost)
}
