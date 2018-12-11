package main

import (
    "io/ioutil"
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
    router.GET("/car/:id", getCar)
}

func getFoo(c *gin.Context) {
    if bytes, e := ioutil.ReadFile(`./static/json/get-foo.json`); e == nil {
        data, _ := simplejson.NewJson(bytes)
        c.JSON(http.StatusOK, data)
    } else {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Server Error"})
    }
}

func getBar(c *gin.Context) {
    if bytes, e := ioutil.ReadFile(`./static/json/get-bar.json`); e == nil {
        data, _ := simplejson.NewJson(bytes)
        c.JSON(http.StatusOK, data)
    } else {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Server Error"})
    }
}

func getCar(c *gin.Context) {
    id := c.Param("id")
    if bytes, e := ioutil.ReadFile(`./static/json/car/get-car` + id + `.json`); e == nil {
        data, _ := simplejson.NewJson(bytes)
        c.JSON(http.StatusOK, data)
    } else {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Server Error"})
    }
}
