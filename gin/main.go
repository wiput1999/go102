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

	r.Run()
}
