package main

import (
	"log"
	"os"

	"github.com/GraphZC/lmwn-jr-2024/internal/config"
	"github.com/GraphZC/lmwn-jr-2024/internal/router"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
}

func main() {
	r := gin.Default()
	router.SetupRouter(r)

	PORT := os.Getenv("PORT")
	log.Println("🚀 Server started on port " + PORT)

	err := r.Run(":" + PORT)
	if err != nil {
		log.Fatalln("Error to start server: ", err.Error())
	}

}
