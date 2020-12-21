package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/assert.v1"
)

func TestCreateUser(t *testing.T) {

	gin.SetMode(gin.TestMode)

	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}
	samples := []struct {
		inputJSON  string
		statusCode int
		username   string
		email      string
	}{
		{
			inputJSON:  `{"username":"Pet", "email": "pet@example.com", "password": "password"}`,
			statusCode: 201,
			username:   "Pet",
			email:      "pet@example.com",
		},
		{
			inputJSON:  `{"username":"Frank", "email": "pet@example.com", "password": "password"}`,
			statusCode: 500,
		},
		{
			inputJSON:  `{"username":"Pet", "email": "grand@example.com", "password": "password"}`,
			statusCode: 500,
		},
		{
			inputJSON:  `{"username":"Kan", "email": "kanexample.com", "password": "password"}`,
			statusCode: 422,
		},
		{
			inputJSON:  `{"username": "", "email": "kan@example.com", "password": "password"}`,
			statusCode: 422,
		},
		{
			inputJSON:  `{"username": "Kan", "email": "", "password": "password"}`,
			statusCode: 422,
		},
		{
			inputJSON:  `{"username": "Kan", "email": "kan@example.com", "password": ""}`,
			statusCode: 422,
		},
	}

	for _, v := range samples {

		r := gin.Default()
		r.POST("/users", server.CreateUser)
		req, err := http.NewRequest(http.MethodPost, "/users", bytes.NewBufferString(v.inputJSON))
		if err != nil {
			t.Errorf("this is the error: %v\n", err)
		}
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		responseInterface := make(map[string]interface{})
		err = json.Unmarshal([]byte(rr.Body.String()), &responseInterface)
		if err != nil {
			t.Errorf("Cannot convert to json: %v", err)
		}

		assert.Equal(t, rr.Code, v.statusCode)
		if v.statusCode == 201 {
			//casting the interface to map:
			responseMap := responseInterface["response"].(map[string]interface{})
			assert.Equal(t, responseMap["username"], v.username)
			assert.Equal(t, responseMap["email"], v.email)
		}
		if v.statusCode == 422 || v.statusCode == 500 {
			responseMap := responseInterface["error"].(map[string]interface{})

			if responseMap["Taken_email"] != nil {
				assert.Equal(t, responseMap["Taken_email"], "Email Already Taken")
			}
			if responseMap["Taken_username"] != nil {
				assert.Equal(t, responseMap["Taken_username"], "Username Already Taken")
			}
			if responseMap["Invalid_email"] != nil {
				assert.Equal(t, responseMap["Invalid_email"], "Invalid Email")
			}
			if responseMap["Required_username"] != nil {
				assert.Equal(t, responseMap["Required_username"], "Required Username")
			}
			if responseMap["Required_email"] != nil {
				assert.Equal(t, responseMap["Required_email"], "Required Email")
			}
			if responseMap["Required_password"] != nil {
				assert.Equal(t, responseMap["Required_password"], "Required Password")
			}
		}
	}
}
