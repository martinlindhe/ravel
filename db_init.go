package ravel

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // for mysql support
	"github.com/jinzhu/gorm"
	"github.com/martinlindhe/ravel/env"
	_ "github.com/mattn/go-sqlite3" // for sqlite3 support
)

// InitDB sets up the database connection from your .env settings
func InitDB() (gorm.DB, error) {

	driver := env.Get("DB_DRIVER", "sqlite3")
	connectionString := generateConnectionString(driver)

	log.Printf("Connecting to %s", connectionString)

	db, err := gorm.Open(driver, connectionString)
	if err != nil {
		fmt.Printf("Connection failed: %s\n", connectionString)
		return db, err
	}

	db.Set("gorm:table_options", "ENGINE=InnoDB")

	// Open doesn't open a connection. Validate DSN data:
	err = db.DB().Ping()
	if err != nil {
		return db, err
	}

	return db, nil
}

func generateConnectionString(driver string) string {

	user := env.Get("DB_USER", "root")
	password := env.Get("DB_PASS", "")
	database := env.Get("DB_NAME", "ravel")
	host := env.Get("DB_HOST", "localhost")
	port := env.GetInt("DB_PORT", 3306)

	connectionString := ""

	switch {
	case driver == "sqlite3":
		// XXX need helper func to get folder root
		connectionString = env.Get("DB_FILE", "./app.db")

	case driver == "mysql":
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

		connectionString = user + ":" + password + "@" + protHost + "/" + database +
			"?charset=utf8&parseTime=True" //&loc=Local"
	}
	return connectionString
}
