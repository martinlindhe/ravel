package ravel

import (
	"log"

	"github.com/nicksnyder/go-i18n/i18n"
)

// T translates given string
func T(s string) string {

	tf, err := i18n.Tfunc("sv-SE", "en-US")
	if err != nil {
		log.Fatal("[translation] Tfunc error", err)
	}

	return tf(s)
}
