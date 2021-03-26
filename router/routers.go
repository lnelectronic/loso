// ---------------------------------------------------------------------------
// LN-ELECTRINIC PROJECT LN-16C10R
// wwww.ln-electronic.com  ProjectManager : @JJOY, @Kimera
// FileData: 25/3/2564 14:27
// ---------------------------------------------------------------------------
package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"loso/database"
	"loso/router/api"
	"net/http"
	"os"
	"strings"
)

// InitGin Creates router
func InitGin(db *database.LnDatabase) *gin.Engine {

	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()
	// Logging to file original gin method.  os.Stdout
	f, _ := os.Create("server.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	g := gin.New()
	g.Use(gin.Recovery())
	g.Use(gin.LoggerWithFormatter(newLog))
	g.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path

		if strings.HasPrefix(path, "/show") {
			fmt.Println("No idea  wait  int the future.")
		}
		c.JSON(http.StatusNotFound, gin.H{
			"Error": "page not found",
		})
	})

	userHandler := api.UserAPI{DB: db}

	u := g.Group("/user")
	{
		u.POST("", userHandler.GetUsers)
		//u.GET("/ping", test)
		u.POST("/getbyid", userHandler.GetUserByIDs)
		u.POST("/insert", userHandler.InsertUser)
		//u.GET("getuser", userHandler.GetUsers)
		u.POST("/getbyname", userHandler.GetUserByUserbame)
	}

	log.Println("Gin Engin: Active...")

	return g
}
