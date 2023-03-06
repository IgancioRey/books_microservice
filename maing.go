package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/IgancioRey/books_microservice/router"
	"github.com/IgancioRey/books_microservice/utils/db"
	"github.com/gin-gonic/gin"
)

const defaultPort = "8080"

var (
	ginRouter *gin.Engine
)

func main() {
	ginRouter = gin.Default()
	router.MapUrls(ginRouter)
	err := db.Init_db()
	defer db.Disconect_db()

	if err != nil {
		fmt.Println("Cannot init db")
		fmt.Println(err)
		return
	}
	fmt.Println("Starting server")
	serverPort := os.Getenv("PORT")
	if serverPort == "" {
		serverPort = defaultPort
	}
	http.ListenAndServe(":"+serverPort, ginRouter)
	//ginRouter.Run(":" + serverPort)
}
