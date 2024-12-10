package database

import (
	"errors"
	"signup-login/model"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	// Initialize a mock database
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Override the global DB variable with the mock DB
	DB = db

	// Mock data
	mockUser := model.User{
		UserID:   1,
		Name:     "John Doe",
		Username: "johndoe",
		Password: "hashedpassword",
	}

	// Mock the query and result
	rows := sqlmock.NewRows([]string{"id", "name", "username", "password"}).
		AddRow(mockUser.UserID, mockUser.Name, mockUser.Username, mockUser.Password)

	mock.ExpectQuery("Select \\* from user where username == \\(\\?\\)").
		WithArgs("johndoe").
		WillReturnRows(rows)

	// Call the GetUser function
	user, err := GetUser("johndoe")

	// Assert the results
	assert.NoError(t, err)
	assert.Equal(t, mockUser, user)

	// Ensure all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUser_NotFound(t *testing.T) {
	// Initialize a mock database
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Override the global DB variable with the mock DB
	DB = db

	// Mock no rows returned
	mock.ExpectQuery("Select \\* from user where username == \\(\\?\\)").
		WithArgs("unknownuser").
		WillReturnRows(sqlmock.NewRows(nil))

	// Call the GetUser function
	user, err := GetUser("unknownuser")

	// Assert the results
	assert.Error(t, err)
	assert.Equal(t, "user does not exist", err.Error())
	assert.Equal(t, model.User{}, user)

	// Ensure all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestAddUser(t *testing.T) {
	// Initialize a mock database
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Override the global DB variable with the mock DB
	DB = db

	// Mock the Exec for insertion
	mock.ExpectExec("INSERT INTO user\\(name, username, password\\) VALUES\\(\\?,\\?,\\?\\)").
		WithArgs("John Doe", "johndoe", sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Call the AddUser function
	err = AddUser("John Doe", "johndoe", "plaintextpassword")

	// Assert the results
	assert.NoError(t, err)

	// Ensure all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestAddUser_Error(t *testing.T) {
	// Initialize a mock database
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Override the global DB variable with the mock DB
	DB = db

	// Mock the Exec for insertion to return an error
	mock.ExpectExec("INSERT INTO user\\(name, username, password\\) VALUES\\(\\?,\\?,\\?\\)").
		WithArgs("John Doe", "johndoe", sqlmock.AnyArg()).
		WillReturnError(errors.New("duplicate entry"))

	// Call the AddUser function
	err = AddUser("John Doe", "johndoe", "plaintextpassword")

	// Assert the results
	assert.Error(t, err)
	assert.Equal(t, "duplicate entry", err.Error())

	// Ensure all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}
