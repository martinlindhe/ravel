package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/martinlindhe/ravel"
	"github.com/martinlindhe/ravel/env"
	"github.com/martinlindhe/ravel/router"
	"github.com/martinlindhe/ravel/views"
)

func main() {

	ravel.Bootstrap()

	appPort := env.GetInt("APP_PORT", 8080)

	db, err := ravel.InitDB()
	if err != nil {
		log.Fatalf("Error connecting to database: %s\n", err)
		os.Exit(1)
	}
	defer db.Close()

	// Enable Logger
	db.LogMode(true)

	ravel.Migrate(&db)
	ravel.Seed(&db)

	fmt.Println(db)

	r := router.Init()

	// r.GET("/", views.Index()) // XXX: cant get this form to work with gorazor views
	r.GET("/", func(c *gin.Context) {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(200, views.Index())
	})

	// listen and serve on 0.0.0.0:8080
	listenAt := fmt.Sprintf(":%d", appPort)

	log.Printf("Starting http server on %s\n", listenAt)

	r.Run(listenAt)
}
