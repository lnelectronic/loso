package api

import (
	"context"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"loso/models"
	"net/http"
	"strconv"
)

// UserApi provides handlers for managing user ln01
type UserAPI struct {
	DB UserDatabase
}

// UserDatabase interface for encapsulating database access.
type UserDatabase interface {
	InsertUser(user *models.User) (*models.User, error)
	GetUserByIDs(ids []primitive.ObjectID) []*models.User
	GetUsers(paging *models.Filter) []*models.User
	CountUser() string
	GetUserByName(name string) *models.User
	Signin(c *gin.Context)
	//FindByUser(username string) (*models.User,error)
	CheckSignin(ctx context.Context, u *models.User) error
}

// UserRepository defines methods the service layer expects
// any repository it interacts with to implement
type UserRepository interface {
	FindByUser(username string) (*models.User, error)
}

var validate = validator.New()

// InserUser creates a User.
func (a *UserAPI) InsertUser(ctx *gin.Context) {

	user := &models.User{}
	var pd, _ = hashPassword(user.Passwd)

	log.Println(pd)
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
	user.Passwd = pd

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
	start, _ = strconv.ParseInt(ctx.DefaultQuery("_skip", "0"), 10, 64)
	end, _ = strconv.ParseInt(ctx.DefaultQuery("_limit", "10"), 10, 64)
	sort = ctx.DefaultQuery("_sort", "_id")
	order = 1

	if sort == "id" {
		sort = "_id"
	}

	if ctx.DefaultQuery("_order", "DESC") == "DESC" {
		order = -1
	}
	if ctx.Query("_order") == "-1" {
		order = -1
	}

	limit := end - start

	// now instance param
	users := a.DB.GetUsers(
		&models.Filter{
			Skip:      &start,
			Limit:     &limit,
			SortKey:   sort,
			SortVal:   order,
			Condition: nil,
		})

	ctx.Header("Doc-Count ", a.DB.CountUser())
	ctx.JSON(http.StatusOK, users)
}

//// Signing used to authenticate extant user
//func (a *UserAPI) Signing(c *gin.Context) {
//	var req signinReq
//	log.Println(req.Passwd,req.Username)
//	if err := c.ShouldBind(req); err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{
//			"Error1": err.Error(),
//		})
//		return
//	}else {
//		log.Println("bine ok")
//	}
//	//
//	//if ok := bindData(c, &req); !ok {
//	//	return
//	//}
//
//	if err := validate.Struct(req); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{
//			"Error": err.Error(),
//		})
//		return
//	}
//
//	u := &models.User{
//		Username: req.Username,
//		Passwd:   req.Passwd,
//	}
//	log.Println(u.Username)
//	log.Println(u.Passwd)
//	ctx := c.Request.Context()
//	//err := h.UserService.Signing(ctx, u)
//	err := a.DB.CheckSignin(ctx, u)
//	if err != nil {
//		log.Printf("Failed to sign in user: %v\n", err.Error())
//		c.JSON(apperrors.Status(err), gin.H{
//			"error": err,
//		})
//		return
//	}
//
//	//tokens, err := h.TokenService.NewPairFromUser(ctx, u, "")
//	//
//	//if err != nil {
//	//	log.Printf("Failed to create tokens for user: %v\n", err.Error())
//	//
//	//	c.JSON(apperrors.Status(err), gin.H{
//	//		"error": err,
//	//	})
//	//	return
//	//}
//	//
//	c.JSON(http.StatusOK, gin.H{
//		"tokens": "tokens",
//	})
//}
//
//

// Signing reaches our to a UserRepository check if the user exists
// and then compares the supplied password with the provided password
// if a valid email/password combo is provided, u will hold all
// available user fields
//func (a *UserAPI) CheckSignin(ctx context.Context, u *models.User) error {
//	//uFetched, err := s.UserRepository.FindByEmail(ctx, u.Email)
//
//	var check = u.Username
//	uFerched,err :=a.DB.FindByUser(check)
//
//	// Will return NotAuthorized to client to omit details of why
//	if err != nil {
//		return apperrors.NewAuthorization("Invalid email and password combination")
//	}
//
//	// verify password - we previously created this method
//	match, err := comparePasswords(uFerched.Passwd, u.Passwd)
//
//	if err != nil {
//		return apperrors.NewInternal()
//	}
//
//	if !match {
//		return apperrors.NewAuthorization("Invalid email and password combination")
//	}
//
//	*u = *uFetched
//	return nil
//}
//
//
