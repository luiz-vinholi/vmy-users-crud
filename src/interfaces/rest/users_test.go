package rest

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/luiz-vinholi/vmy-users-crud/src/infra/services"
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slices"
)

func TestGetUserIntegration(t *testing.T) {
	assert := assert.New(t)

	client := createDBCollection()
	userId := insertUserInDB(client)
	defer deleteUserInDB(client, userId)

	auth := services.NewAuth()
	token, _ := auth.GenerateToken(map[string]interface{}{"id": userId})

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users/"+userId, nil)
	req.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(res, req)

	bodyStr := res.Body.String()
	body := parseBodyString(bodyStr)

	assert.Equal(200, res.Code)
	assert.NotNil(body["id"])
	assert.NotNil(body["name"])
	assert.NotNil(body["email"])
	assert.Nil(body["password"])
	assert.NotNil(body["birthDate"])
	assert.NotNil(body["age"])
	assert.NotNil(body["address"])
	assert.NotNil(body["address"].(map[string]interface{})["street"])
	assert.NotNil(body["address"].(map[string]interface{})["city"])
	assert.NotNil(body["address"].(map[string]interface{})["state"])
}

func TestGetUserInvalidTokenIntegration(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users/obiwan", nil)
	req.Header.Set("Authorization", "Bearer INVALID_TOKEN")
	router.ServeHTTP(res, req)

	bodyStr := res.Body.String()
	body := parseBodyString(bodyStr)

	assert.Equal(401, res.Code)
	assert.Equal("invalid-authorization-token", body["error"].(map[string]interface{})["code"])
}

func TestGetUserNotFoundIntegration(t *testing.T) {
	assert := assert.New(t)

	client := createDBCollection()
	userId := insertUserInDB(client)
	defer deleteUserInDB(client, userId)

	auth := services.NewAuth()
	token, _ := auth.GenerateToken(map[string]interface{}{"id": userId})

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users/invalidid", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(res, req)

	bodyStr := res.Body.String()
	body := parseBodyString(bodyStr)

	assert.Equal(404, res.Code)
	assert.Equal("user-not-found", body["error"].(map[string]interface{})["code"])
}

func TestGetUsersIntegration(t *testing.T) {
	assert := assert.New(t)

	client := createDBCollection()
	userId := insertUserInDB(client)
	userId2 := insertUserInDB(client)
	defer deleteUserInDB(client, userId)
	defer deleteUserInDB(client, userId2)

	auth := services.NewAuth()
	token, _ := auth.GenerateToken(map[string]interface{}{"id": userId})

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users/?", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(res, req)

	bodyStr := res.Body.String()
	body := parseBodyString(bodyStr)

	assert.Equal(200, res.Code)
	assert.Equal(int64(2), int64(body["total"].(float64)))
	assert.Equal(2, int(len(body["users"].([]interface{}))))
	users := body["users"].([]interface{})
	for i := 0; i < len(users); i++ {
		user := users[i].(map[string]interface{})
		assert.NotNil(user["id"])
		assert.NotNil(user["name"])
		assert.NotNil(user["email"])
		assert.Nil(user["password"])
		assert.NotNil(user["birthDate"])
		assert.NotNil(user["age"])
		assert.NotNil(user["address"])
		assert.NotNil(user["address"].(map[string]interface{})["street"])
		assert.NotNil(user["address"].(map[string]interface{})["city"])
		assert.NotNil(user["address"].(map[string]interface{})["state"])
	}
}

func TestGetUsersInvalidTokenIntegration(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users/", nil)
	req.Header.Set("Authorization", "Bearer INVALID_TOKEN")
	router.ServeHTTP(res, req)

	bodyStr := res.Body.String()
	body := parseBodyString(bodyStr)

	assert.Equal(401, res.Code)
	assert.Equal("invalid-authorization-token", body["error"].(map[string]interface{})["code"])
}

func TestGetUsersWithPaginationIntegration(t *testing.T) {
	assert := assert.New(t)

	client := createDBCollection()
	userId := insertUserInDB(client)
	userId2 := insertUserInDB(client)
	userId3 := insertUserInDB(client)
	userId4 := insertUserInDB(client)
	userIds := []string{userId, userId2, userId3, userId4}
	defer deleteUsersInDB(client, userIds)

	auth := services.NewAuth()
	token, _ := auth.GenerateToken(map[string]interface{}{"id": userId})

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users/?offset=2&limit=2", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(res, req)

	bodyStr := res.Body.String()
	body := parseBodyString(bodyStr)

	assert.Equal(200, res.Code)
	assert.Equal(4, int(body["total"].(float64)))
	assert.Equal(2, int(len(body["users"].([]interface{}))))
	users := body["users"].([]interface{})
	for i := 0; i < len(users); i++ {
		user := users[i].(map[string]interface{})
		assert.NotNil(user["id"])
		assert.NotNil(user["name"])
		assert.NotNil(user["email"])
		assert.Nil(user["password"])
		assert.NotNil(user["birthDate"])
		assert.NotNil(user["age"])
		assert.NotNil(user["address"])
		assert.NotNil(user["address"].(map[string]interface{})["street"])
		assert.NotNil(user["address"].(map[string]interface{})["city"])
		assert.NotNil(user["address"].(map[string]interface{})["state"])
	}
}

func TestGetUsersHighLimitIntegration(t *testing.T) {
	assert := assert.New(t)

	client := createDBCollection()
	userId := insertUserInDB(client)
	defer deleteUserInDB(client, userId)

	auth := services.NewAuth()
	token, _ := auth.GenerateToken(map[string]interface{}{"id": userId})

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users/?limit=200", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(res, req)

	bodyStr := res.Body.String()
	body := parseBodyString(bodyStr)
	fields := []string{"limit", "offset"}

	assert.Equal(400, res.Code)
	errors := body["errors"].([]interface{})
	for i := 0; i < len(errors); i++ {
		err := errors[i].(map[string]interface{})
		hasField := slices.Contains(fields, strings.ToLower(err["field"].(string)))
		assert.True(hasField)
	}
}

func TestGetUsersInvalidQueryIntegration(t *testing.T) {
	assert := assert.New(t)

	client := createDBCollection()
	userId := insertUserInDB(client)
	defer deleteUserInDB(client, userId)

	auth := services.NewAuth()
	token, _ := auth.GenerateToken(map[string]interface{}{"id": userId})

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users/?limit=invalid&offset=invalid", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(res, req)

	bodyStr := res.Body.String()
	body := parseBodyString(bodyStr)

	assert.Equal(400, res.Code)
	assert.NotNil(body["error"])
}

func TestCreateUserIntegration(t *testing.T) {
	assert := assert.New(t)

	client := createDBCollection()
	userId := insertUserInDB(client)
	defer deleteUserInDB(client, userId)

	auth := services.NewAuth()
	token, _ := auth.GenerateToken(map[string]interface{}{"id": userId})

	res := httptest.NewRecorder()
	payload := `{
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
	req, _ := http.NewRequest("POST", "/users/", strings.NewReader(payload))
	req.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(res, req)

	bodyStr := res.Body.String()
	body := parseBodyString(bodyStr)

	createdUser := findUserInDB(client, body["id"].(string))

	assert.Equal(200, res.Code)
	assert.NotNil(body["id"])
	assert.Equal("test.create.user@example.com", createdUser["email"])
}

func TestCreateUserInvalidTokenIntegration(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/users/", nil)
	req.Header.Set("Authorization", "Bearer INVALID_TOKEN")
	router.ServeHTTP(res, req)

	bodyStr := res.Body.String()
	body := parseBodyString(bodyStr)

	assert.Equal(401, res.Code)
	assert.Equal("invalid-authorization-token", body["error"].(map[string]interface{})["code"])
}

func TestCreateUserInvalidPayloadIntegration(t *testing.T) {
	assert := assert.New(t)

	client := createDBCollection()
	userId := insertUserInDB(client)
	defer deleteUserInDB(client, userId)

	auth := services.NewAuth()
	token, _ := auth.GenerateToken(map[string]interface{}{"id": userId})

	res := httptest.NewRecorder()
	payload := `{
		"name": "Test User",
		"email": "testexample.com",
		"password": "calkestis",
    "address": {
			"state": "Paraná",
			"country": "Brasil" 
    }
	}`
	req, _ := http.NewRequest("POST", "/users/", strings.NewReader(payload))
	req.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(res, req)

	bodyStr := res.Body.String()
	body := parseBodyString(bodyStr)
	fields := []string{"email", "password", "street", "city", "state", "country", "birthdate"}

	assert.Equal(400, res.Code)
	errors := body["errors"].([]interface{})
	for i := 0; i < len(errors); i++ {
		err := errors[i].(map[string]interface{})
		field := err["field"].(string)
		hasField := slices.Contains(fields, strings.ToLower(field))
		assert.True(hasField, field)
	}
}

func TestCreateUserWithEmailInUseIntegration(t *testing.T) {
	assert := assert.New(t)

	client := createDBCollection()
	userId := insertUserInDB(client)
	defer deleteUserInDB(client, userId)
	user := findUserInDB(client, userId)
	email := user["email"].(string)

	auth := services.NewAuth()
	token, _ := auth.GenerateToken(map[string]interface{}{"id": userId})

	res := httptest.NewRecorder()
	payload := fmt.Sprintf(`{
		"name": "Test User",
		"email": "%s",
		"birthDate": "1980-06-21",
    "address": {
			"street": "Rua do Teste, 550 - Integração",
			"city": "Mandaguari",
			"state": "Paraná",
			"country": "Brasil" 
    }
	}`, email)
	req, _ := http.NewRequest("POST", "/users/", strings.NewReader(payload))
	req.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(res, req)

	bodyStr := res.Body.String()
	body := parseBodyString(bodyStr)

	assert.Equal(400, res.Code)
	assert.Equal("email-in-use", body["error"].(map[string]interface{})["code"])
}

func TestUpdateUserIntregration(t *testing.T) {
	assert := assert.New(t)

	client := createDBCollection()
	userId := insertUserInDB(client)
	defer deleteUserInDB(client, userId)
	userBefore := findUserInDB(client, userId)

	auth := services.NewAuth()
	token, _ := auth.GenerateToken(map[string]interface{}{"id": userId})

	res := httptest.NewRecorder()
	newName := "Paz Vizsla"
	newStreet := "Como deve ser"
	payload := fmt.Sprintf(`{
		"name": "%s",
    "address": {
			"street": "%s"
    }
	}`, newName, newStreet)
	req, _ := http.NewRequest("PATCH", "/users/"+userId, strings.NewReader(payload))
	req.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(res, req)

	userAfter := findUserInDB(client, userId)

	assert.Equal(204, res.Code)
	assert.Equal("", res.Body.String())
	assert.NotEqual(userBefore["name"], userAfter["name"])
	assert.NotEqual(
		userBefore["address"].(map[string]interface{})["street"],
		userAfter["address"].(map[string]interface{})["street"])
	assert.Equal(userBefore["email"], userAfter["email"])
	assert.Equal(
		userBefore["address"].(map[string]interface{})["city"],
		userBefore["address"].(map[string]interface{})["city"])
}

func TestUpdateUserInvalidTokenIntegration(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("PATCH", "/users/ahsoka", nil)
	req.Header.Set("Authorization", "Bearer INVALID_TOKEN")
	router.ServeHTTP(res, req)

	bodyStr := res.Body.String()
	body := parseBodyString(bodyStr)

	assert.Equal(401, res.Code)
	assert.Equal("invalid-authorization-token", body["error"].(map[string]interface{})["code"])
}

func TestUpdateUserInvalidPayloadIntregration(t *testing.T) {
	assert := assert.New(t)

	client := createDBCollection()
	userId := insertUserInDB(client)
	defer deleteUserInDB(client, userId)
	userBefore := findUserInDB(client, userId)

	auth := services.NewAuth()
	token, _ := auth.GenerateToken(map[string]interface{}{"id": userId})

	res := httptest.NewRecorder()
	payload := `{
		"birthDate": "26-08-2000"
	}`
	req, _ := http.NewRequest("PATCH", "/users/"+userId, strings.NewReader(payload))
	req.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(res, req)

	userAfter := findUserInDB(client, userId)

	bodyStr := res.Body.String()
	body := parseBodyString(bodyStr)
	err := body["errors"].([]interface{})[0]

	assert.Equal(400, res.Code)
	assert.Equal(
		"birthdate",
		strings.ToLower(err.(map[string]interface{})["field"].(string)))
	assert.Equal(userBefore["birthDate"], userAfter["birthDate"])
}

func TestUpdateUserNotFoundIntregration(t *testing.T) {
	assert := assert.New(t)

	client := createDBCollection()
	userId := insertUserInDB(client)
	defer deleteUserInDB(client, userId)

	auth := services.NewAuth()
	token, _ := auth.GenerateToken(map[string]interface{}{"id": userId})

	res := httptest.NewRecorder()
	payload := `{
		"birthDate": "2000-08-26"
	}`
	req, _ := http.NewRequest("PATCH", "/users/notexists", strings.NewReader(payload))
	req.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(res, req)

	bodyStr := res.Body.String()
	body := parseBodyString(bodyStr)
	err := body["error"].(map[string]interface{})

	assert.Equal(404, res.Code)
	assert.Equal("user-not-found", err["code"])
}

func TestDeleteUserIntegration(t *testing.T) {
	assert := assert.New(t)

	client := createDBCollection()
	userId := insertUserInDB(client)
	defer deleteUserInDB(client, userId)

	auth := services.NewAuth()
	token, _ := auth.GenerateToken(map[string]interface{}{"id": userId})

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/users/"+userId, nil)
	req.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(res, req)

	deletedUser := findUserInDB(client, userId)

	assert.Equal(204, res.Code)
	assert.Equal("", res.Body.String())
	assert.Nil(deletedUser)
}

func TestDeleteUserInvalidTokenIntegration(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/users/ahsoka", nil)
	req.Header.Set("Authorization", "Bearer INVALID_TOKEN")
	router.ServeHTTP(res, req)

	bodyStr := res.Body.String()
	body := parseBodyString(bodyStr)

	assert.Equal(401, res.Code)
	assert.Equal("invalid-authorization-token", body["error"].(map[string]interface{})["code"])
}

func TestDeleteUserNotFoundIntegration(t *testing.T) {
	assert := assert.New(t)

	client := createDBCollection()
	userId := insertUserInDB(client)
	defer deleteUserInDB(client, userId)

	auth := services.NewAuth()
	token, _ := auth.GenerateToken(map[string]interface{}{"id": userId})

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/users/notexist", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(res, req)

	bodyStr := res.Body.String()
	body := parseBodyString(bodyStr)
	err := body["error"].(map[string]interface{})

	assert.Equal(404, res.Code)
	assert.Equal("user-not-found", err["code"])
}
