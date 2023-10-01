package routes

import (
	"net/http"

	"github.com/avdushin/rgoauth/pkg/handlers"
	"github.com/gorilla/mux"
)

// @ Роуты авторизации
func SetupAuthRoutes(router *mux.Router) {
	router.HandleFunc("/register", handlers.RegisterHandler).Methods(http.MethodPost)
	router.HandleFunc("/login", handlers.LoginHandler).Methods(http.MethodPost)
}
