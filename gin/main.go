package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type response struct {
	Number  string `json:"number"`
	Message string `json:"message"`
}

func main() {
	r := gin.Default()

	r.GET("/fizzbuzz/:number", func(c *gin.Context) {
		_, err := strconv.Atoi(c.Param("number"))

		if err == nil {
			c.JSON(200, response{
				Message: c.Param("number"),
				Number:  c.Param("number"),
			})
		}
	})

	r.POST("/users", userHandlers)

	r.Run()
}

func userHandlers(c *gin.Context) {
	var users Users

	err := c.Bind(&users)

	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.XML(200, users)
}
