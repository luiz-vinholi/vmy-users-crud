package rest

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateSessionIntegration(t *testing.T) {
	assert := assert.New(t)
	router := setupRouter()

	res := httptest.NewRecorder()
	body := `{"email": "test@example.com", "password": "test@123"}`
	req, _ := http.NewRequest("POST", "/sessions", strings.NewReader(body))
	router.ServeHTTP(res, req)

	log.Printf("%s", res.Body.String())
	assert.Equal(200, res.Code)
}

func TestCreateSessionInvalidPayloadIntegration(t *testing.T) {
	assert := assert.New(t)
	router := setupRouter()

	res := httptest.NewRecorder()
	body := `{"email": "testexample.com", "password": "123"}`
	req, _ := http.NewRequest("POST", "/sessions", strings.NewReader(body))
	router.ServeHTTP(res, req)

	assert.Equal(400, res.Code)
}

func TestCreateSessionInvalidCredentialsIntegration(t *testing.T) {
	assert := assert.New(t)
	router := setupRouter()

	res := httptest.NewRecorder()
	body := `{"email": "test@example.com", "password": "invalidpass"}`
	req, _ := http.NewRequest("POST", "/sessions", strings.NewReader(body))
	router.ServeHTTP(res, req)

	assert.Equal(401, res.Code)
}

func TestSetSessionPasswordIntegration(t *testing.T) {
	assert := assert.New(t)
	router := setupRouter()

	res := httptest.NewRecorder()
	body := `{"password": "Test@123"}`
	userId := ""
	url := fmt.Sprintf("/sessions/users/%s/passwords", userId)
	req, _ := http.NewRequest("PUT", url, strings.NewReader(body))
	router.ServeHTTP(res, req)

	assert.Equal(204, res.Code)
}

func TestSetSessionPasswordInvalidPayloadIntegration(t *testing.T) {
	assert := assert.New(t)
	router := setupRouter()

	res := httptest.NewRecorder()
	body := `{"password": "123"}`
	userId := ""
	url := fmt.Sprintf("/sessions/users/%s/passwords", userId)
	req, _ := http.NewRequest("PUT", url, strings.NewReader(body))
	router.ServeHTTP(res, req)

	assert.Equal(400, res.Code)
}

func TestSetSessionPasswordUserNotFoundIntegration(t *testing.T) {
	assert := assert.New(t)
	router := setupRouter()

	res := httptest.NewRecorder()
	body := `{"password": "123"}`
	userId := ""
	url := fmt.Sprintf("/sessions/users/%s/passwords", userId)
	req, _ := http.NewRequest("PUT", url, strings.NewReader(body))
	router.ServeHTTP(res, req)

	assert.Equal(404, res.Code)
}
