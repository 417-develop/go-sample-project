package main

import (
    "io/ioutil"
    "net/http"

    simplejson "github.com/bitly/go-simplejson"
    "github.com/gin-gonic/gin"
)

var r *gin.Engine

func main() {
    initializeRoutes().Run(":8088")
}

func initializeRoutes() *gin.Engine {
    r = gin.Default()
    r.GET("/foo", getFoo)
    r.GET("/bar", getBar)
    r.GET("/car/:id", getCar)
    return r
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
        c.JSON(http.StatusNotFound, gin.H{"message": "Resource is not found"})
    }
}
