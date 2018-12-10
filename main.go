package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	router = gin.Default()
	initializeRoutes()
	router.Run(":8088")
}

func initializeRoutes() {
	router.GET("/foo", getFoo)
	router.GET("/bar", getBar)
}

func getFoo(c *gin.Context) {
	bytes, e := ioutil.ReadFile(`./static/json/get-foo.json`)
	if e != nil {
		log.Fatal(e)
	}
	data, e := simplejson.NewJson(bytes)
	fmt.Println(data)
	c.JSON(http.StatusOK, data)
}

func getBar(c *gin.Context) {
	bytes, e := ioutil.ReadFile(`./static/json/get-bar.json`)
	if e != nil {
		log.Fatal(e)
	}
	data, e := simplejson.NewJson(bytes)
	fmt.Println(data)
	c.JSON(http.StatusOK, data)
}
