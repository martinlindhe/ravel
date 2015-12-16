package ravel

import (
	"log"
	"os"

	"github.com/martinlindhe/ravel/env"
	"github.com/nicksnyder/go-i18n/i18n"
)

// Bootstrap initializes the framework
func Bootstrap() {
	err := env.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
		os.Exit(1)
	}

	i18n.MustLoadTranslationFile("resources/lang/en-US.yaml")

	routerInstance = initRouter()
}
