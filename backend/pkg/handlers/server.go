package handlers

import "net/http"

// ServerRequest обрабатывает запросы и отправляет ответ "OK"
func ServerRequest(w http.ResponseWriter, r *http.Request) {
	// Устанавливаем заголовок ответа
	w.Header().Set("Content-Type", "text/plain")

	// Устанавливаем статус ответа на 200 OK
	w.WriteHeader(http.StatusOK)

	// Отправляем ответ "OK"
	_, err := w.Write([]byte("OK"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
