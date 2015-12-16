package router

import "github.com/gin-gonic/gin"

// Init boots up Gin for routing, https://gin-gonic.github.io/gin/
func Init() *gin.Engine {

	// Creates a router without any middleware by default
	r := gin.New()

	// Global middleware
	r.Use(gin.Logger())

	//	r.Use(gzip.Gzip(gzip.DefaultCompression))

	r.GET("/ping", pingController)

	r.Static("/js", "./public/js")
	r.Static("/css", "./public/css")
	r.Static("/img", "./public/img")
	r.Static("/fonts", "./public/fonts")
	r.Static("/flags", "./public/flags")
	//r.LoadHTMLFiles("./public/index.html")

	return r
}

// curl -v "http://localhost:8080/ping"
func pingController(c *gin.Context) {
	c.JSON(200, gin.H{"pong": "now"})
}
