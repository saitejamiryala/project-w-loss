package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/saitejamiryala/project-w-loss/model"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("hello saiteja")

	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	port := os.Getenv("APPLICATION_PORT")
	log.Println("Starting server on port", port)

	connStr := "host=" + os.Getenv("DB_HOST") + " port=5432 user=" + os.Getenv("DB_USER_NAME") + " password= dbname=" + os.Getenv("DB_NAME") + " sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	} else {
		log.Println("Successfully connected to db")
	}

	defer db.Close()

	r := gin.Default()

	r.POST("/track", func(c *gin.Context) {

		var calories model.CaloriesConsumed

		if err := c.ShouldBindJSON(&calories); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, calories)
		fmt.Println("api is working")
	})

	connErr := r.Run(":" + port)
	if connErr != nil {
		return
	}

}
