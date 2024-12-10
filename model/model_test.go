package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	mockUser := User{
		UserID:   1,
		Name:     "John Doe",
		Username: "johndoe",
		Password: "hashedpassword",
	}
	err := mockUser.HashPassword(mockUser.Password)
	assert.NoError(t, err)
}

func TestCheckCorrectPassword(t *testing.T) {
	mockUser := User{
		UserID:   1,
		Name:     "John Doe",
		Username: "johndoe",
		Password: "hashedpassword",
	}
	err := mockUser.HashPassword(mockUser.Password)
	assert.NoError(t, err)
	err = mockUser.CheckPassword("hashedpassword")
	assert.NoError(t, err)
}

func TestCheckWrongPassword(t *testing.T) {
	mockUser := User{
		UserID:   1,
		Name:     "John Doe",
		Username: "johndoe",
		Password: "hashedpassword",
	}
	err := mockUser.HashPassword(mockUser.Password)
	assert.NoError(t, err)
	err = mockUser.CheckPassword("nonpassword")
	assert.Error(t, err)
}
