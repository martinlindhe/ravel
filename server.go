package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/martinlindhe/ubique.se/dotenv"
	"github.com/martinlindhe/ubique.se/migration"
	"github.com/martinlindhe/ubique.se/seed"
	"github.com/martinlindhe/ubique.se/server"
)

func main() {

	err := dotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	routes := server.InitRoutes()

	dbUser := dotenv.Get("DB_USER", "user")
	dbPass := dotenv.Get("DB_PASS", "pass")
	dbName := dotenv.Get("DB_NAME", "ubique")
	dbHost := dotenv.Get("DB_HOST", "localhost")
	dbPort := dotenv.GetInt("DB_PORT", 3306)

	db, err := server.InitDB(dbHost, dbPort, dbUser, dbPass, dbName)
	if err != nil {
		fmt.Printf("Error connecting to database: %s\n", err)
		os.Exit(1)
	}
	defer db.Close()

	// Enable Logger
	db.LogMode(true)

	migration.Migrate(&db)

	seed.Seed(&db)

	fmt.Println(db)

	routes.Run(":8080") // listen and serve on 0.0.0.0:8080
}
