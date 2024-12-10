package service

import (
	"signup-login/database"
)

// service function for user login
func UserLogin(username string, password string) error {
	user, err := database.GetUser(username)
	if err != nil {
		return err
	}

	passwordErr := user.CheckPassword(password)
	if passwordErr != nil {
		return passwordErr
	}

	return nil
}

// service function for user sign up
func UserSignUp(name string, username string, password string) error {
	return database.AddUser(name, username, password)
}
