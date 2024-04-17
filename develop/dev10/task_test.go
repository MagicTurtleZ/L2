package main

import (
	"errors"
	"testing"
)

// Тест для customTelnet с корректным подключением
func TestCustomTelnet_SuccessfulConnection(t *testing.T) {
	host = "0.0.0.0:8081"
	makeTimeout("10s")
	err := customTelnet()

	// Проверка на отсутствие ошибок
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

// Тест для customTelnet с ошибкой при подключении
func TestCustomTelnet_ConnectionError(t *testing.T) {
	expectedErr := errors.New("conection server error")
	host = "0.0.0.0:8082"
	makeTimeout("10s")
	// Вызов функции customTelnet с мокированным подключением
	err := customTelnet()

	// Проверка на ожидаемую ошибку
	if err == nil || err.Error() != expectedErr.Error() {
		t.Errorf("Expected error: %v, but got: %v", expectedErr, err)
	}
}
