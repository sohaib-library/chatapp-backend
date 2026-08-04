package main

import (
	"github.com/sohaib-library/chatapp-backend/database"
	"github.com/sohaib-library/chatapp-backend/route"

	"github.com/gin-gonic/gin"
)

func main() {
	DB := database.Database(".env")

	database.Migrations(DB)

	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	route.RegisterRoute(router, DB)

	router.Run(":8000")
}
