package controller

import (
	"Beer-BackendV1/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AddEntry will add a new entry with the given JSON context data.
// Can only add if user is in database | first create User than add Beverages
func AddEntry(context *gin.Context) {
	var input model.Entry

	// checks if body of request is JSON
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// gets user from the input variables
	user, err := model.FindUserByName(input.Name, input.Surname)

	if err != nil {
		// UserNotFound
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.Name, input.Surname = user.Name, user.Surname

	// assigns a new Entry to the User Entries
	savedEntry, err := input.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// returns the new Entry
	context.JSON(http.StatusCreated, gin.H{"data": savedEntry})
}

func GetAllEntries(context *gin.Context) {
	entries, err := model.ReturnAllEntries()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"beverages": entries})
}

func GetAllEntriesOfUser(context *gin.Context) {
	var input model.UserName

	// checks if body of request is JSON
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	entries, err := model.ReturnAllEntriesOfUser(input)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"beverages": entries})
}
