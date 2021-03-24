package api

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"loso/models"
	"net/http"
	"strconv"
)

// UserDatabase interface for encapsulating database access.
type UserDatabase interface {
	InsertUser(user *models.User) (*models.User, error)
	GetUserByIDs(ids []primitive.ObjectID) []*models.User
	GetUsers(paging *models.Filter) []*models.User
	CountUser() string
	GetUserByName(name string) *models.User
}

// UserApi provides handlers for managing user ln01
type UserAPI struct {
	DB UserDatabase
}

var validate = validator.New()

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

// GetUserByIDs returns user by id
func (a *UserAPI) GetUserByIDs(ctx *gin.Context) {
	withIDs(ctx, "id", func(ids []primitive.ObjectID) {
		ctx.JSON(http.StatusOK, a.DB.GetUserByIDs(ids))
	})
}

//GetUserByName
func (a *UserAPI) GetUserByUserbame(c *gin.Context) {
	username := c.Query("username")

	c.JSON(http.StatusOK, a.DB.GetUserByName(username))
}

// GetUsers returns all users
// _end=5&_order=DESC&_sort=id&_start=0 adapt react-admin
func (a *UserAPI) GetUsers(ctx *gin.Context) {
	var (
		start int64
		end   int64
		sort  string
		order int
	)
	id := ctx.DefaultQuery("id", "")
	if id != "" {
		a.GetUserByIDs(ctx)
		return
	}
	start, _ = strconv.ParseInt(ctx.DefaultQuery("_start", "0"), 10, 64)
	end, _ = strconv.ParseInt(ctx.DefaultQuery("_end", "10"), 10, 64)
	sort = ctx.DefaultQuery("_sort", "_id")
	order = 1

	if sort == "id" {
		sort = "_id"
	}

	if ctx.DefaultQuery("_order", "DESC") == "DESC" {
		order = -1
	}

	limit := end - start
	users := a.DB.GetUsers(
		&models.Filter{
			Skip:      &start,
			Limit:     &limit,
			SortKey:   sort,
			SortVal:   order,
			Condition: nil,
		})

	ctx.Header("Doc-Count ", a.DB.CountUser())
	ctx.JSON(200, users)
}
