package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/fizzbuzz/:number", func(c *gin.Context) {
		number, err := strconv.Atoi(c.Param("number"))

		if err == nil {
			c.JSON(200, gin.H{
				"message": number,
			})
		}
	})

	r.Run()
}
