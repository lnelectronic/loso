package api

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"loso/models"
	"net/http"
)

// UserDatabase interface for encapsulating database access.
type UserDatabase interface {
	InsertUser(user *models.User) (*models.User, error)
}

// UserApi provides handlers for managing user ln01.
type UserAPI struct {
	DB UserDatabase
}

var validate = validator.New()

// InserUser creates a User.
//func (a *UserAPI) InsertUser(ctx *gin.Context) {
//	var user = models.User{}
//	err := ctx.ShouldBindJSON(&user)
//	if err != nil {
//		ctx.AbortWithError(http.StatusInternalServerError, errors.New("Error: Check Data insert."))
//		return
//	}
//
//	err = validate.Struct(&user)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	result, err := a.DB.InsertUser(user.New())
//	if errors.Is(err, sql.ErrNoRows) {
//		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Server not found."})
//		return
//	}
//	ctx.JSON(http.StatusOK, result)
//	return
//}
//
//

// InserUser creates a User.
func (a *UserAPI) InsertUser(ctx *gin.Context) {

	user := &models.User{}
	if err := ctx.ShouldBind(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	if err := validate.Struct(user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	result, err := a.DB.InsertUser(user.New())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}
	// Error form database
	if errors.Is(err, sql.ErrNoRows) {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Server not found."})
		return
	}

	ctx.JSON(http.StatusOK, result)
	return
}
