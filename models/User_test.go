package models

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestCreateUser(t *testing.T) {
	err := refreshUserTable()
	if err != nil {
		t.Fatal(err)
	}
	newUser := User{
		ID:       1,
		Email:    "test@example.com",
		Username: "test",
		Password: "password",
	}
	savedUser, err := newUser.CreateUser(server.DB)
	if err != nil {
		t.Errorf("error creating user: %v", err)
		return
	}
	assert.Equal(t, newUser.ID, savedUser.ID)
	assert.Equal(t, newUser.Email, savedUser.Email)
	assert.Equal(t, newUser.Username, savedUser.Username)

}
