package main

import (
	"os"

	"github.com/62teknologi-test-alfatah/db"
	"github.com/62teknologi-test-alfatah/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	routes.Router(r)
	port := ":8080"
	db.Database.Migrate()
	getPort := os.Getenv("PORT")
	if getPort != "" {
		port = ":" + getPort
	}
	r.Run(port)
}
