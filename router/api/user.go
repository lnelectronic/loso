package api

import (
	"errors"
	"github.com/gin-gonic/gin"
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

// InserUser creates a User.
func (a *UserAPI) InsertUser(ctx *gin.Context) {
	var user = models.User{}
	if err := ctx.ShouldBindJSON(&user); err == nil {
		result, err := a.DB.InsertUser(user.New())

		if err != nil {
			ctx.JSON(http.StatusOK, err)
		}
		ctx.JSON(200, result)
	} else {
		//ctx.AbortWithError(500, errors.New("LN : Sorry  error"))
		ctx.AbortWithError(http.StatusBadRequest, errors.New("LN : Sorry  error"))
	}
}
