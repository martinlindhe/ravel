package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/martinlindhe/ubique.se/db"
	"github.com/martinlindhe/ubique.se/dotenv"
	"github.com/martinlindhe/ubique.se/migration"
	"github.com/martinlindhe/ubique.se/router"
	"github.com/martinlindhe/ubique.se/seed"
	"github.com/martinlindhe/ubique.se/views"
	"github.com/nicksnyder/go-i18n/i18n"
)

func bootstrap() {
	err := dotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
		os.Exit(1)
	}

	i18n.MustLoadTranslationFile("resources/lang/en-US.yaml")
}

func main() {

	bootstrap()

	dbUser := dotenv.Get("DB_USER", "user")
	dbPass := dotenv.Get("DB_PASS", "pass")
	dbName := dotenv.Get("DB_NAME", "ubique")
	dbHost := dotenv.Get("DB_HOST", "localhost")
	dbPort := dotenv.GetInt("DB_PORT", 3306)

	db, err := db.Init("mysql", dbHost, dbPort, dbUser, dbPass, dbName)
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

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
