package main

import (
  "github.com/gin-gonic/gin"
  "net/http"
)
 func main() {
   rooter := gin.Default()

   rooter.GET("/ping", func(c *gin.Context) {
   		c.JSON(http.StatusOK, gin.H{
   			"message": "pong",
   		})
   	})
   	rooter.Run(":8088") // listen and serve on 0.0.0.0:8080
  }
