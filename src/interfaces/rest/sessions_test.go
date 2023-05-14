package rest

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/slices"
)

func TestCreateSessionIntegration(t *testing.T) {
	assert := assert.New(t)

	client := createDBCollection()
	userId := insertUserInDB(client)
	defer deleteUserInDB(client, userId)

	pass := "somepassword"
	hashed := encryptPassword(pass)
	updateUserInDB(client, userId, map[string]interface{}{"password": hashed})
	user := findUserInDB(client, userId)

	res := httptest.NewRecorder()
	payload := fmt.Sprintf(`{
		"email": "%s",
		"password": "%s"
	}`, user["email"], pass)
	req, _ := http.NewRequest("POST", "/sessions/", strings.NewReader(payload))
	router.ServeHTTP(res, req)

	bodyStr := res.Body.String()
	body := parseBodyString(bodyStr)

	assert.Equal(200, res.Code)
	assert.NotNil(body["token"])
}

func TestCreateSessionInvalidPayloadIntegration(t *testing.T) {
	assert := assert.New(t)

	client := createDBCollection()
	userId := insertUserInDB(client)
	defer deleteUserInDB(client, userId)

	res := httptest.NewRecorder()
	payload := `{"email": "testexample.com", "password": "123"}`
	req, _ := http.NewRequest("POST", "/sessions/", strings.NewReader(payload))
	router.ServeHTTP(res, req)

	bodyStr := res.Body.String()
	body := parseBodyString(bodyStr)
	fields := []string{"email", "password"}

	assert.Equal(400, res.Code)
	errors := body["errors"].([]interface{})
	for i := 0; i < len(errors); i++ {
		err := errors[i].(map[string]interface{})
		hasField := slices.Contains(fields, strings.ToLower(err["field"].(string)))
		assert.True(hasField)
	}
}

func TestCreateSessionInvalidCredentialsIntegration(t *testing.T) {
	assert := assert.New(t)

	client := createDBCollection()
	userId := insertUserInDB(client)
	defer deleteUserInDB(client, userId)

	pass := "validpassword"
	hashed := encryptPassword(pass)
	updateUserInDB(client, userId, map[string]interface{}{"password": hashed})
	user := findUserInDB(client, userId)

	res := httptest.NewRecorder()
	payload := fmt.Sprintf(`{
		"email": "%s",
		"password": "invalidpass"
	}`, user["email"])
	req, _ := http.NewRequest("POST", "/sessions/", strings.NewReader(payload))
	router.ServeHTTP(res, req)

	bodyStr := res.Body.String()
	body := parseBodyString(bodyStr)

	err := body["error"].(map[string]interface{})
	assert.Equal(401, res.Code)
	assert.Equal("invalid-credentials", err["code"])
}

func TestCreateSessionUserNotFoundIntegration(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	payload := `{
		"email": "email@notexists.com",
		"password": "somepass"
	}`
	req, _ := http.NewRequest("POST", "/sessions/", strings.NewReader(payload))
	router.ServeHTTP(res, req)

	bodyStr := res.Body.String()
	body := parseBodyString(bodyStr)

	err := body["error"].(map[string]interface{})
	assert.Equal(401, res.Code)
	assert.Equal("invalid-credentials", err["code"])
}

func TestSetSessionPasswordIntegration(t *testing.T) {
	assert := assert.New(t)

	client := createDBCollection()
	userId := insertUserInDB(client)
	defer deleteUserInDB(client, userId)

	res := httptest.NewRecorder()
	body := `{"password": "Test@123"}`
	url := fmt.Sprintf("/sessions/users/%s/passwords", userId)
	req, _ := http.NewRequest("PUT", url, strings.NewReader(body))
	router.ServeHTTP(res, req)

	user := findUserInDB(client, userId)

	assert.Equal(204, res.Code)
	assert.Equal("", res.Body.String())
	assert.NotNil(user["password"])
}

func TestSetSessionPasswordInvalidPayloadIntegration(t *testing.T) {
	assert := assert.New(t)

	client := createDBCollection()
	userId := insertUserInDB(client)
	defer deleteUserInDB(client, userId)

	res := httptest.NewRecorder()
	payload := `{"password": "123"}`
	url := fmt.Sprintf("/sessions/users/%s/passwords", userId)
	req, _ := http.NewRequest("PUT", url, strings.NewReader(payload))
	router.ServeHTTP(res, req)

	bodyStr := res.Body.String()
	body := parseBodyString(bodyStr)

	assert.Equal(400, res.Code)
	assert.GreaterOrEqual(1, len(body["errors"].([]interface{})))
}

func TestSetSessionPasswordUserNotFoundIntegration(t *testing.T) {
	assert := assert.New(t)
	router := setupRouter()

	res := httptest.NewRecorder()
	payload := `{"password": "123456"}`
	url := "/sessions/users/notexists/passwords"
	req, _ := http.NewRequest("PUT", url, strings.NewReader(payload))
	router.ServeHTTP(res, req)

	bodyStr := res.Body.String()
	body := parseBodyString(bodyStr)

	err := body["error"].(map[string]interface{})
	assert.Equal(404, res.Code)
	assert.Equal("user-not-found", err["code"])
}

func encryptPassword(pass string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	return string(hash)
}
