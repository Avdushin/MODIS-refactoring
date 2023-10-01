package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// @ Отправляем запрос на сервер
func serverRequest(w http.ResponseWriter, r *http.Request) {
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

func TestServerRequest(t *testing.T) {
	// Создаем фейковый HTTP запрос (GET запрос на корневой путь)
	req, err := http.NewRequest("GET", "http://localhost:4000", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Создаем фейковый ResponseWriter
	rr := httptest.NewRecorder()

	// Вызываем функцию ServerRequest с фейковыми запросом и ResponseWriter
	serverRequest(rr, req)

	// Проверяем статус ответа (должен быть 200 OK)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Ожидался статус 200 OK, но получено: %v", status)
	}

	// Проверяем содержимое ответа (должен быть "OK")
	expectedResponse := "OK"
	actualResponse := rr.Body.String()
	if actualResponse != expectedResponse {
		t.Errorf("Ожидался ответ \"%s\", но получено: \"%s\"", expectedResponse, actualResponse)
	}
}
