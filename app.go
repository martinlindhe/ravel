package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/martinlindhe/ravel/db"
	"github.com/martinlindhe/ravel/env"
	"github.com/martinlindhe/ravel/migration"
	"github.com/martinlindhe/ravel/router"
	"github.com/martinlindhe/ravel/seed"
	"github.com/martinlindhe/ravel/views"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nicksnyder/go-i18n/i18n"
)

func bootstrap() {
	err := env.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
		os.Exit(1)
	}

	i18n.MustLoadTranslationFile("resources/lang/en-US.yaml")
}

func main() {

	bootstrap()

	appPort := env.GetInt("APP_PORT", 8080)

	db, err := db.Init()
	if err != nil {
		log.Fatalf("Error connecting to database: %s\n", err)
		os.Exit(1)
	}
	defer db.Close()

	// Enable Logger
	db.LogMode(true)

	migration.Migrate(&db)

	seed.Seed(&db)

	fmt.Println(db)

	r := router.Init()

	// XXX  cannot use tpl.Index (type func() string) as type gin.HandlerFunc in argument to r.RouterGroup.GET
	//r.GET("/", tpl.Index())

	r.GET("/", func(c *gin.Context) {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(200, views.Index())
	})

	// listen and serve on 0.0.0.0:8080
	listenAt := fmt.Sprintf(":%d", appPort)

	log.Printf("Starting http server on %s\n", listenAt)

	r.Run(listenAt)
}
