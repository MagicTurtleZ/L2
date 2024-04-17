package main

import (
	"os"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Тест для случая успешного получения данных
func TestCustomWget_Success(t *testing.T) {
	// Мокирование сервера
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("test data"))
	}))
	defer ts.Close()

	// Вызов функции customWget
	err := customWget(ts.URL)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Проверка наличия файла
	if _, err := os.Stat("website.txt"); os.IsNotExist(err) {
		t.Error("File not found")
	}

	// Проверка содержимого файла
	data, err := os.ReadFile("website.txt")
	if err != nil {
		t.Errorf("Error reading file: %v", err)
	}
	if string(data) != "test data" {
		t.Errorf("Unexpected file content: %s", data)
	}

	// Удаление созданного файла
	os.Remove("website.txt")
}

// Тест для случая ошибки при получении данных
func TestCustomWget_Error(t *testing.T) {
	// Вызов функции customWget с неверным URL
	err := customWget("invalid-url")
	if err == nil {
		t.Error("Expected error, got nil")
	}
}
