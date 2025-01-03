package repository

import (
	"github.com/mahdibouaziz/go-testing-playground/webapp/pkg/data"
)

type DatabaseRepo interface {
	// return the connection to the DB
	// Connection() *sql.DB

	// AllUsers returns all users as a slice of *data.User
	AllUsers() ([]*data.User, error)

	// GetUser returns one user by id
	GetUser(id int) (*data.User, error)

	// GetUserByEmail returns one user by email address
	GetUserByEmail(email string) (*data.User, error)

	// UpdateUser updates one user in the database
	UpdateUser(u data.User) error

	// DeleteUser deletes one user from the database, by id
	DeleteUser(id int) error

	// InsertUser inserts a new user into the database, and returns the ID of the newly inserted row
	InsertUser(user data.User) (int, error)

	// ResetPassword is the method we will use to change a user's password.
	ResetPassword(id int, password string) error

	// InsertUserImage inserts a user profile image into the database.
	InsertUserImage(i data.UserImage) (int, error)
}
