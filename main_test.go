package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Тест: запрос корректен, возвращается 200 и тело не пустое
func TestMainHandlerSuccess(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)
	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(MainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// Проверяем, что статус ответа 200 OK
	require.Equal(t, http.StatusOK, responseRecorder.Code)

	// Проверяем, что тело не пустое
	assert.NotEmpty(t, responseRecorder.Body.String())
}

// Тест: город не поддерживается, возвращается 400 и ошибка
func TestMainHandlerWrongCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=spb", nil)
	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(MainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// Проверяем, что статус ответа 400
	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)

	// Проверяем, что сообщение об ошибке корректное
	assert.Equal(t, "wrong city value", responseRecorder.Body.String())
}

// Тест: count больше чем количество кафе, возвращаются все кафе
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)
	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(MainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// Проверяем, что статус ответа 200 OK
	require.Equal(t, http.StatusOK, responseRecorder.Code)

	// Проверяем, что возвращено ровно 4 кафе (все доступные)
	assert.Equal(t, "Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент", responseRecorder.Body.String())
	assert.Len(t, strings.Split(responseRecorder.Body.String(), ","), 4)
}
