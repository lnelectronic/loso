// ---------------------------------------------------------------------------
// LN-ELECTRINIC PROJECT LN-16C10R
// wwww.ln-electronic.com  ProjectManager : @JJOY, @Kimera
// FileData: 24/3/2564 23:07
// ---------------------------------------------------------------------------
package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func withIDs(ctx *gin.Context, name string, f func(id []primitive.ObjectID)) {
	ids, b := ctx.GetQueryArray(name)
	objectIds := []primitive.ObjectID{}
	abort := errors.New("invalid id")
	if b {
		for _, id := range ids {
			if objID, err := primitive.ObjectIDFromHex(id); err == nil {
				objectIds = append(objectIds, objID)
			} else {
				ctx.AbortWithError(http.StatusBadRequest, abort)
			}
		}
		f(objectIds)
	} else {
		ctx.AbortWithError(http.StatusBadRequest, abort)
	}
}

func withID(ctx *gin.Context, name string, f func(id primitive.ObjectID)) {
	if id, err := primitive.ObjectIDFromHex(ctx.Param(name)); err == nil {
		f(id)
	} else {
		ctx.AbortWithError(400, errors.New("invalid id"))
	}
}
