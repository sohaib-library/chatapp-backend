package main

import (
	
	"chatapp-backend/database"
	"chatapp-backend/route"
	"github.com/gin-gonic/gin"

)

func main() {
	DB := database.Database(".env")
	defer DB.Close()

	database.Migertions(DB)

	router := gin.Default()

	route.RegisterRoute(router, DB)

	router.Run(":8000")
}
