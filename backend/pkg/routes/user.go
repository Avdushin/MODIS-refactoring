package routes

import (
	"net/http"

	"github.com/avdushin/rgoauth/pkg/handlers"
	"github.com/gorilla/mux"
)

// @ Пользовательские роуты
func SetupUserRoutes(router *mux.Router) {
	router.HandleFunc("/api/users/{id}", handlers.GetUser).Methods(http.MethodGet)
	router.HandleFunc("/api/users/{id}", handlers.UpdateUser).Methods(http.MethodPut)
}
