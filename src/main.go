package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.IndentedJSON(200, gin.H{
			"hello": "world",
		})
	})
	router.Run("localhost:4000")
}
