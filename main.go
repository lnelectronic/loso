package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"loso/config"
	"loso/database"
	"loso/router"
	"net/http"
)

func test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "00001",
	})
}
func main() {

	db, err := database.NewCon(config.Hostmgo, "ln-smt")
	if err != nil {
		panic(err)
	} else {
		log.Println("LN-ELECTRONIC Project SmartFarm")
		log.Println("Server MongoDb: Active...")
	}
	defer db.Close()

	r := router.InitGin(db)
	r.Run(config.ServerHost)

}
