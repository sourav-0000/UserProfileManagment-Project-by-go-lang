package controller

import (
	"net/http"
	"userProfileManagment/model"
	"userProfileManagment/service"

	"github.com/gin-gonic/gin"
)

// UserController defines HTTP handler methods for users
type UserController struct {
	service service.UserService
}

// NewUserController creates a new instance of UserController
func NewUserController(service service.UserService) *UserController {
	return &UserController{service: service}
}

// CreateUser handles the creation of a new user
func (c *UserController) CreateUser(ct *gin.Context) {
	// Define the request struct to capture the incoming JSON data
	var request struct {
		Name    string `json:"name" binding:"required"`
		Email   string `json:"email" binding:"required,email"`
		Contact string `json:"contact" binding:"required"`
		Address string `json:"address" binding:"required"`
	}

	// Bind the incoming JSON request body to the 'request' struct
	if err := ct.ShouldBindJSON(&request); err != nil {
		// If binding fails, return a 400 error with a message
		ct.JSON(http.StatusBadRequest, gin.H{"error": "Request is not valid"})
		return
	}

	// Now map the request data to the user model
	user := model.User{
		Name:    request.Name,
		Email:   request.Email,
		Contact: request.Contact,
		Address: request.Address,
	}

	// Call the service to create the user
	if err := c.service.CreateUser(user); err != nil {
		// If the creation fails, return a 500 error with a message
		ct.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save the user details"})
		return
	}

	// Return success message if the user is successfully stored
	ct.JSON(http.StatusCreated, gin.H{
		"message": "Successfully stored user details",
	})
}

// GetUser handles fetching a user by ID
func (c *UserController) GetUser(ct *gin.Context) {
	var request struct {
		Id uint `json:"id" binding:"required"`
	}
	if err := ct.ShouldBindJSON(&request); err != nil {
		ct.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	user, err := c.service.GetUserByID(uint(request.Id))
	if err != nil {
		ct.JSON(http.StatusInternalServerError, gin.H{"error": "Id not found"})
		return
	}

	ct.JSON(http.StatusOK, user)
}

// UpdateUser handles updating an existing user
func (c *UserController) UpdateUser(ct *gin.Context) {
	// Request struct to bind JSON payload
	var request struct {
		Id      uint   `json:"id" binding:"required"`
		Name    string `json:"name"`
		Email   string `json:"email"`
		Contact string `json:"contact"`
		Address string `json:"address"`
	}

	// Bind the incoming JSON request to the request struct
	if err := ct.ShouldBindJSON(&request); err != nil {
		ct.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Fetch the user from the database (assuming this function exists in the service layer)
	user, err := c.service.GetUserByID(request.Id)
	if err != nil {
		// Return 500 if there is an error fetching the user
		ct.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
		return
	}

	if user == nil {
		// Return 404 if no user was found
		ct.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Update the user's fields with the request data
	user.Name = request.Name
	user.Email = request.Email
	user.Contact = request.Contact
	user.Address = request.Address

	// Call the service method to update the user in the database
	if err := c.service.UpdateUser(user, request.Id); err != nil {
		// If an error occurs in the service layer, return a 500 status
		ct.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user details"})
		return
	}

	// Return success response
	ct.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

// DeleteUser handles deleting a user by ID
func (c *UserController) DeleteUser(ctx *gin.Context) {
	// Define a struct to bind the incoming JSON body
	var request struct {
		Id uint `json:"id" binding:"required"`
	}

	// Bind the request body to the struct
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request, ID is required"})
		return
	}

	// Call the service layer to delete the user
	err := c.service.DeleteUser(request.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// If successful, respond with a success message
	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
