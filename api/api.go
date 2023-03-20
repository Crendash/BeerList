package api

import (
	"Beer-BackendV1/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// List all users
func Users(context *gin.Context) {
	users, err := model.ReturnAllUsers()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"users": users})
}

// Create user and add to DB
func CreateUser(context *gin.Context) {
	var input model.User

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := model.User{
		Name:    input.Name,
		Surname: input.Surname,
		Balance: input.Balance,
	}

	savedUser, err := user.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"user": savedUser})
}

// Update user balance
func UpdateUserBalance(context *gin.Context) {
	var input model.BalanceInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := model.FindUserByName(input.Name, input.Surname)

	if err != nil {
		// UserNotFound
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	model.UpdateBalance(user, input.Balance)
}

// Delete user
func DeleteUser(context *gin.Context) {
	var input model.UserName

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := model.FindUserByName(input.Name, input.Surname)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	model.Delete(user)
}
