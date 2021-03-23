package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"loso/models"
)

// UserDatabase interface for encapsulating database access.
type UserDatabase interface {
	InsertUser(user *models.User) (*models.User, error)
}

// UserApi provides handlers for managing user ln01.
type UserAPI struct {
	DB UserDatabase
}

// InserUser creates a User.
func (a *UserAPI) InsertUser(ctx *gin.Context) {
	var user = models.User{}
	if err := ctx.ShouldBindJSON(&user); err == nil {
		result, err := a.DB.InsertUser(user.New())

		if err != nil {
			ctx.JSON(203, err)
		}
		ctx.JSON(200, result)
	} else {
		ctx.AbortWithError(500, errors.New("LN : Sorry  error"))
	}
}
