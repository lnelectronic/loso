package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"log"
	"loso/config"
	"loso/database"
	"loso/router"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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

	engine := router.InitGin(db)
	srv := &http.Server{
		Addr:    config.ServerHost,
		Handler: engine,
	}

	// Initializing server
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	// a timeout of 5 seconds. Wait for interrupt signal to gracefully shutdown the server with
	quit := make(chan os.Signal)
	// kill  default syscall.SIGTERM kill -2  kill -9  but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")

}
