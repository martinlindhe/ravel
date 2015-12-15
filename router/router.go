package router

import (
	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/martinlindhe/ubique.se/tpl"
	_ "github.com/martinlindhe/ubique.se/tpl/layout"
)

// https://gin-gonic.github.io/gin/
func InitRoutes() *gin.Engine {

	// Creates a router without any middleware by default
	r := gin.New()

	// Global middleware
	r.Use(gin.Logger())

	r.Use(gzip.Gzip(gzip.DefaultCompression))

	r.GET("/", tpl.Index)

	r.GET("/ping", pingController)

	r.Static("/js", "./public/js")
	r.Static("/css", "./public/css")
	r.Static("/img", "./public/img")
	r.LoadHTMLFiles("./public/index.html")

	return r
}

// curl -v "http://localhost:8080/ping"
func pingController(c *gin.Context) {
	c.String(200, "pong")
}
