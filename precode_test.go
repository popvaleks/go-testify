package main

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	uri := fmt.Sprintf("/?city=moscow&count=%d", totalCount+1)
	req := httptest.NewRequest(http.MethodGet, uri, nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code, "expected 200 OK")

	itemList := strings.Split(responseRecorder.Body.String(), ",")
	assert.Len(t, itemList, totalCount, "expected 4 items")
}

func TestMainHandlerWithCorrectParams(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/?city=moscow&count=1", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code, "expected 200 OK")
	assert.NotEmpty(t, responseRecorder.Body, "expected not empty response")
}

func TestMainHandlerWithWrongCity(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/?city=biysk&count=1", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code, "expected 400 ERR")
	assert.Equal(t, responseRecorder.Body.String(), "wrong city value", "expected other error message")
}

// ниже доп тесты не по заданию
func TestMainHandlerWithoutCount(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/?city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code, "expected 400 ERR")
	assert.Equal(t, responseRecorder.Body.String(), "count missing", "expected other error message")
}

func TestMainHandlerWithoutIncorrectCount(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/?city=moscow&count=asd", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code, "expected 400 ERR")
	assert.Equal(t, responseRecorder.Body.String(), "wrong count value", "expected other error message")
}

func TestMainHandlerWithoutCity(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/?count=2", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code, "expected 400 ERR")
	assert.Equal(t, responseRecorder.Body.String(), "wrong city value", "expected other error message")
}
