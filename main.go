package main

import (
	"log"
	"net/http"
	"os"
	"sesi-11/database"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db, err := database.ConnectDB(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
	)

	if err != nil {
		panic(err)
	}

	if db != nil {
		log.Println("db connected")
	}

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})

	router.Run(":" + os.Getenv("PORT"))
}
