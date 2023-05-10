package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func Test_incrementKey(t *testing.T) {

	hostport = "localhost:6379"

	mcPostBody := map[string]interface{}{
		"key":   "age",
		"value": 19,
	}
	body, _ := json.Marshal(mcPostBody)
	r := gin.Default()
	r.POST("/redis/incr", incrementKey)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/redis/incr", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Код ответа не совпадает с ожидаемым")
	assert.JSONEq(t, `{"value": 20}`, w.Body.String(), "Тело ответа не совпадает с ожидаемым")
}

func Test_getSign(t *testing.T) {
	mcPostBody := map[string]interface{}{
		"text": "test",
		"key":  "test123",
	}
	body, _ := json.Marshal(mcPostBody)
	r := gin.Default()
	r.POST("/sign/hmacsha512", getSign)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/sign/hmacsha512", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Код ответа не совпадает с ожидаемым")
	assert.Equal(t, "\"b936cee86c9f87aa5d3c6f2e84cb5a4239a5fe50480a6ec66b70ab5b1f4ac6730c6c515421b327ec1d69402e53dfb49ad7381eb067b338fd7b0cb22247225d47\"",
		w.Body.String(), "Тело ответа не совпадает с ожидаемым")
}

func Test_addUser(t *testing.T) {
	db, err = sql.Open("postgres", "host=localhost port=5433 user=postgres password=postgres dbname=zen sslmode=disable")
	if err != nil {
		return
	}
	defer db.Close()

	mcPostBody := map[string]interface{}{
		"name": "Alex",
		"age":  21,
	}
	body, _ := json.Marshal(mcPostBody)
	r := gin.Default()
	r.POST("/postgres/users", addUser)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/postgres/users", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code, "Код ответа не совпадает с ожидаемым")
	assert.JSONEq(t, `{"id": 1}`, w.Body.String(), "Тело ответа не совпадает с ожидаемым")
}
