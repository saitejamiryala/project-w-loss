package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/saitejamiryala/project-w-loss/model"
)

func main() {
	fmt.Println("hello saiteja")

	r := gin.Default()

	r.POST("/track", func(c *gin.Context) {

		var calories model.CaloriesConsumed

		if err := c.ShouldBindJSON(&calories); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("api is working")
	})

	err := r.Run(":9669")
	if err != nil {
		return
	}

}
