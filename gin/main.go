package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/wiput1999/go102/hello/fizzbuzz"
)

type response struct {
	Number  string `json:"number"`
	Message string `json:"message"`
}

func main() {
	r := gin.Default()

	r.GET("/fizzbuzzr", func(c *gin.Context) {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		c.JSON(200, fizzBuzzByRandom(r1))
	})

	r.GET("/fizzbuzz/:number", func(c *gin.Context) {
		jwtToken := c.GetHeader("Authorization")[7:]

		token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte("AllYourBase"), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"message": err.Error(),
			})
			return
		}

		if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		} else {
			c.JSON(http.StatusUnauthorized, nil)
			return
		}

		number, err := strconv.Atoi(c.Param("number"))

		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]string{
				"message": err.Error(),
			})
			return
		}

		c.JSON(200, response{
			Message: fizzbuzz.FizzBuzz(number),
			Number:  c.Param("number"),
		})
	})

	r.POST("/users", userHandler)

	r.GET("/token", tokenHandler)

	r.Run()
}

func userHandler(c *gin.Context) {
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

func tokenHandler(c *gin.Context) {
	signingKey := []byte("AllYourBase")
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
		Issuer:    "Wiput",
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := jwtToken.SignedString(signingKey)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"token": ss,
	})
}

type randomer interface {
	Intn(int) int
}

func fizzBuzzByRandom(r randomer) response {
	number := r.Intn(100)
	return response{
		Message: fizzbuzz.FizzBuzz(number),
		Number:  strconv.Itoa(number),
	}
}
