package repository

import (
	"fmt"
	"strings"
	"userProfileManagment/model"

	"gorm.io/gorm"
)

// UserRepository defines methods for interacting with the users table
type UserRepository interface {
	CreateUser(user model.User) error
	GetUserByID(id uint) (*model.User, error)
	UpdateUser(user *model.User, id uint) error
	DeleteUser(id uint) error
	UserExists(id uint) bool
}

// userRepository is the implementation of UserRepository
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new instance of userRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

// CreateUser inserts a new user into the database
func (r *userRepository) CreateUser(user model.User) error {
	// Correct table name ("users" instead of "user")
	query := "INSERT INTO users (name, email, contact, address) VALUES (?, ?, ?, ?)"

	// Execute the query with the user data
	err := r.db.Exec(query, user.Name, user.Email, user.Contact, user.Address).Error
	if err != nil {
		return fmt.Errorf("error inserting user: %v", err)
	}
	return nil
}

// GetUserByID retrieves a user by its ID
func (r *userRepository) GetUserByID(id uint) (*model.User, error) {
	// Correct query with parameterized ID
	query := "SELECT * FROM users WHERE id = ?"
	var user model.User

	// Execute the query and map the result to the user object
	err := r.db.Raw(query, id).Scan(&user).Error
	if err != nil {
		return nil, fmt.Errorf("error retrieving user: %v", err)
	}

	// Return the user object
	return &user, nil
}

// UpdateUser updates an existing user in the database
func (r *userRepository) UpdateUser(user *model.User, id uint) error {
	// Start building the update query dynamically
	query := "UPDATE users SET "
	args := []interface{}{}
	setClauses := []string{}

	// Conditionally add fields to the update query based on the non-empty values in the user struct
	if user.Name != "" {
		setClauses = append(setClauses, "name = ?")
		args = append(args, user.Name)
	}
	if user.Email != "" {
		setClauses = append(setClauses, "email = ?")
		args = append(args, user.Email)
	}
	if user.Contact != "" {
		setClauses = append(setClauses, "contact = ?")
		args = append(args, user.Contact)
	}
	if user.Address != "" {
		setClauses = append(setClauses, "address = ?")
		args = append(args, user.Address)
	}

	// If no fields are provided to update, return an error
	if len(setClauses) == 0 {
		return fmt.Errorf("no valid fields to update")
	}

	// Join the set clauses with commas
	query += strings.Join(setClauses, ", ") // Corrected to use strings.Join

	// Add the WHERE clause to the query
	query += " WHERE id = ?"
	args = append(args, id)

	// Execute the query
	err := r.db.Exec(query, args...).Error
	if err != nil {
		return fmt.Errorf("error updating user: %v", err)
	}

	return nil
} // DeleteUser deletes a user by its ID
func (r *userRepository) DeleteUser(id uint) error {
	// Correct query for deleting user by ID
	query := "DELETE FROM users WHERE id = ?"
	err := r.db.Exec(query, id).Error
	if err != nil {
		return fmt.Errorf("error deleting user: %v", err)
	}
	return nil
}
func (r *userRepository) UserExists(id uint) bool {
	var count int64
	err := r.db.Model(&model.User{}).Where("id = ?", id).Count(&count).Error
	if err != nil || count == 0 {
		return false
	}
	return true
}
