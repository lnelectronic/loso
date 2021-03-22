package main

import (
	"github.com/gin-gonic/gin"
	"loso/config"
	"loso/database"
	"loso/router/api"
	"net/http"
)

func test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "00001",
	})
}
func main() {

	db, err := database.NewCon("ln-smt")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	Handler := api.UserAPI{DB: db}

	r := gin.Default()
	r.GET("/ping", test)
	r.POST("/add", Handler.InsertUser)

	r.Run(config.ServerHost)

}
