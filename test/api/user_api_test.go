package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUserEndpointBadRequst(t *testing.T) {
	router := SetUpRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/users", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)

}

func TestCrateUserEndpointSuccess(t *testing.T) {
	router := SetUpRouter()

	bodyRequest := `{
						"firstname": "test_firstname", 
						"lastname": "test_lastname", 
						"email": "test_email@email.com",
						"password": "password123", 
						"gender": "Male"
					}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/users", strings.NewReader(bodyRequest))
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
	assert.Equal(t, true, strings.Contains(w.Body.String(), "id"))

}
