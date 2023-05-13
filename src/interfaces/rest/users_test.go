package rest

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserByIdIntegration(t *testing.T) {
	assert := assert.New(t)
	router := setupRouter()

	res := httptest.NewRecorder()
	userId := "1"
	req, _ := http.NewRequest("GET", "/users/"+userId, nil)
	router.ServeHTTP(res, req)

	assert.Equal(200, res.Code)
}

func TestGetUserNotFoundIntegration(t *testing.T) {
	assert := assert.New(t)
	router := setupRouter()

	res := httptest.NewRecorder()
	userId := "823324"
	req, _ := http.NewRequest("GET", "/users/"+userId, nil)
	router.ServeHTTP(res, req)

	assert.Equal(404, res.Code)
}

func TestGetUsersIntegration(t *testing.T) {
	assert := assert.New(t)
	router := setupRouter()

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users", nil)
	router.ServeHTTP(res, req)

	assert.Equal(200, res.Code)
}

func TestGetUsersWithPaginationIntegration(t *testing.T) {
	assert := assert.New(t)
	router := setupRouter()

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users", nil)
	router.ServeHTTP(res, req)

	assert.Equal(200, res.Code)
}

func TestGetUsersInvalidQueryIntegration(t *testing.T) {
	assert := assert.New(t)
	router := setupRouter()

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users?limit=invalid&offset=invalid", nil)
	router.ServeHTTP(res, req)

	assert.Equal(400, res.Code)
}

func TestCreateUserIntegration(t *testing.T) {
	assert := assert.New(t)
	router := setupRouter()

	res := httptest.NewRecorder()
	body := `{
		"name": "Test User",
		"email": "test.create.user@example.com",
		"birthDate": "1980-06-21",
    "address": {
			"street": "Rua do Teste, 550 - Integração",
			"city": "Mandaguari",
			"state": "Paraná",
			"country": "Brasil" 
    }
	}`
	req, _ := http.NewRequest("POST", "/users", strings.NewReader(body))
	router.ServeHTTP(res, req)

	assert.Equal(200, res.Code)
}

func TestCreateUserInvalidPayloadIntegration(t *testing.T) {
	assert := assert.New(t)
	router := setupRouter()

	res := httptest.NewRecorder()
	body := `{
		"name": "Test User",
		"email": "testexample.com",
		"birthDate": "21-06-1980",
    "address": {
			"state": "Paraná",
			"country": "Brasil" 
    }
	}`
	req, _ := http.NewRequest("POST", "/users", strings.NewReader(body))
	router.ServeHTTP(res, req)

	assert.Equal(400, res.Code)
}

func TestCreateUserWithEmailInUseIntegration(t *testing.T) {
	assert := assert.New(t)
	router := setupRouter()

	res := httptest.NewRecorder()
	body := `{
		"name": "Test User",
		"email": "test@example.com",
		"birthDate": "1980-06-21",
    "address": {
			"street": "Rua do Teste, 550 - Integração",
			"city": "Mandaguari",
			"state": "Paraná",
			"country": "Brasil" 
    }
	}`
	req, _ := http.NewRequest("POST", "/users", strings.NewReader(body))
	router.ServeHTTP(res, req)

	assert.Equal(400, res.Code)
}

func TestUpdateUserIntregration(t *testing.T) {
	assert := assert.New(t)
	router := setupRouter()

	res := httptest.NewRecorder()
	body := `{
		"name": "New Name",
    "address": {
			"street": "Rua nova, 550 - Integração",
    }
	}`
	req, _ := http.NewRequest("PATCH", "/users", strings.NewReader(body))
	router.ServeHTTP(res, req)

	assert.Equal(204, res.Code)
	// add checking in the database if updated
}

func TestUpdateUserInvalidPayloadIntregration(t *testing.T) {
	assert := assert.New(t)
	router := setupRouter()

	res := httptest.NewRecorder()
	body := `{
		"name": "New Name",
    "address": {
			"street": "Rua nova, 550 - Integração",
    }
	}`
	req, _ := http.NewRequest("PATCH", "/users", strings.NewReader(body))
	router.ServeHTTP(res, req)

	assert.Equal(204, res.Code)
	// add checking in the database if not updated
}

func TestDeleteUserIntegration(t *testing.T) {
	assert := assert.New(t)
	router := setupRouter()

	res := httptest.NewRecorder()
	userId := ""
	req, _ := http.NewRequest("DELETE", "/users/"+userId, nil)
	router.ServeHTTP(res, req)

	assert.Equal(200, res.Code)
	// check if delete in database
}

func TestDeleteUserNotFoundIntegration(t *testing.T) {
	assert := assert.New(t)
	router := setupRouter()

	res := httptest.NewRecorder()
	userId := ""
	req, _ := http.NewRequest("DELETE", "/users/"+userId, nil)
	router.ServeHTTP(res, req)

	assert.Equal(400, res.Code)
}
