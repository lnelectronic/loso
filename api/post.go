package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"loso/model"
)

// The PostDatabase interface for encapsulating database access.
type PostDatabase interface {
	GetPosts(paging *model.Paging) []*model.Post
	GetPostByID(id primitive.ObjectID) *model.Post
	CreatePost(post *model.Post) *model.Post
	UpdatePost(post *model.Post) *model.Post
	DeletePostByID(id primitive.ObjectID) error
	CountPost(condition interface{}) string
}

// The PostAPI provides handlers for managing posts.
type PostAPI struct {
	DB PostDatabase
}

// CreatePost creates a post.
func (a *PostAPI) CreatePost(ctx *gin.Context) {
	var post = model.Post{}
	if err := ctx.ShouldBind(&post); err == nil {
		if result := a.DB.CreatePost(post.New()); result != nil {
			ctx.JSON(201, result)
		} else {
			ctx.AbortWithError(500, errors.New("CreatePost error"))
		}
	} else {
		ctx.AbortWithError(500, errors.New("ShouldBind error"))
	}
}
