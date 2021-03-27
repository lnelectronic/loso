// ---------------------------------------------------------------------------
// LN-ELECTRINIC PROJECT LN-16C10R
// wwww.ln-electronic.com  ProjectManager : @JJOY, @Kimera
// FileData: 28/3/2564 1:56
// ---------------------------------------------------------------------------
package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"loso/models"
	"loso/models/apperrors"
	"net/http"
)

// signinReq is not exported
type signinReq struct {
	Username string `bson:"username" json:"username" validate:"required"`
	Passwd   string `bson:"passwd" json:"passwd" validate:"required,gte=6,lte=30"`
}

// Signing used to authenticate extant user
func (a *UserAPI) Signins(c *gin.Context) {
	var req signinReq
	log.Println(req.Passwd, req.Username)
	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error1": err.Error(),
		})
		return
	} else {
		log.Println("bine ok")
	}
	//
	//if ok := bindData(c, &req); !ok {
	//	return
	//}

	if err := validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	u := &models.User{
		Username: req.Username,
		Passwd:   req.Passwd,
	}
	log.Println(u.Username)
	log.Println(u.Passwd)
	ctx := c.Request.Context()
	//err := h.UserService.Signing(ctx, u)
	err := a.DB.CheckSignin(ctx, u)
	if err != nil {
		log.Printf("Failed to sign in user: %v\n", err.Error())
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	//tokens, err := h.TokenService.NewPairFromUser(ctx, u, "")
	//
	//if err != nil {
	//	log.Printf("Failed to create tokens for user: %v\n", err.Error())
	//
	//	c.JSON(apperrors.Status(err), gin.H{
	//		"error": err,
	//	})
	//	return
	//}
	//
	c.JSON(http.StatusOK, gin.H{
		"tokens": "tokens",
	})
}
