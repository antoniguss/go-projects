package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"message":"pong"}`, w.Body.String())
}

func TestAddRoute(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()

	// Test valid addition
	input := map[string]int{
		"number1": 5,
		"number2": 3,
	}
	jsonValue, _ := json.Marshal(input)
	req, _ := http.NewRequest("POST", "/add", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var response map[string]int
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, 8, response["result"])

	// Test invalid input
	w = httptest.NewRecorder()
	invalidInput := `{"number1": "invalid"}`
	req, _ = http.NewRequest("POST", "/add", bytes.NewBuffer([]byte(invalidInput)))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
}

func TestSubtractRoute(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()

	input := map[string]int{
		"number1": 10,
		"number2": 3,
	}
	jsonValue, _ := json.Marshal(input)
	req, _ := http.NewRequest("POST", "/subtract", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var response map[string]int
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, 7, response["result"])
}

func TestMultiplyRoute(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()

	input := map[string]int{
		"number1": 4,
		"number2": 5,
	}
	jsonValue, _ := json.Marshal(input)
	req, _ := http.NewRequest("POST", "/multiply", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var response map[string]int
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, 20, response["result"])
}

func TestDivideRoute(t *testing.T) {
	router := setupRouter()

	t.Run("Valid Division", func(t *testing.T) {
		w := httptest.NewRecorder()
		input := map[string]int{
			"number1": 10,
			"number2": 2,
		}
		jsonValue, _ := json.Marshal(input)
		req, _ := http.NewRequest("POST", "/divide", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		var response map[string]int
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, 5, response["result"])
	})

	t.Run("Division by Zero", func(t *testing.T) {
		w := httptest.NewRecorder()
		input := map[string]int{
			"number1": 10,
			"number2": 0,
		}
		jsonValue, _ := json.Marshal(input)
		req, _ := http.NewRequest("POST", "/divide", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Code)
		var response map[string]string
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, "Division by zero", response["error"])
	})
}
