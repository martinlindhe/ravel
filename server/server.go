package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/martinlindhe/ubique.se/tpl"
)

func InitDB(host string, port int, user string, password string, database string) (gorm.DB, error) {

	protHost := ""

	if host == "" || host == "127.0.0.1" || host == "localhost" {
		protHost = "unix"
	} else {
		sPort := ""
		if port != 3306 {
			sPort = ":" + fmt.Sprintf("%d", port)
		}

		protHost = "tcp(" + host + sPort + ")"
	}

	connectionString := user + ":" + password + "@" + protHost + "/" + database +
		"?charset=utf8&parseTime=True" //&loc=Local"

	log.Printf("Connecting to %s", connectionString)

	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		fmt.Printf("Connection failed: %s\n", connectionString)
		return db, err
	}

	// Open doesn't open a connection. Validate DSN data:
	err = db.DB().Ping()
	if err != nil {
		return db, err
	}

	//db.DB().SetMaxIdleConns(10)
	//db.DB().SetMaxOpenConns(100)

	return db, nil
}

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
