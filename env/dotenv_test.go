package env

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {

	// create dotenv
	envFile := "/tmp/dotenv-test"

	d := []byte("TEST=wop\n" +
		"TEST0=wop\n" +
		"STR=hö hej\n" +
		"NUM=12356\n" +
		"NOTANUM=ababb\n" +
		"BOOL1=YES\n" +
		"BOOL2=ON\n" +
		"BOOL3=TRUE\n" +
		"BOOL4=false\n")

	err := ioutil.WriteFile(envFile, d, 0644)
	if err != nil {
		panic(err)
	}

	Load(envFile)

	assert.Equal(t, "hö hej", Get("STR", ""))
	assert.Equal(t, "default", Get("NOTFOUNDSTR", "default"))

	assert.Equal(t, 12356, GetInt("NUM", 0))
	assert.Equal(t, 0, GetInt("NOTANUM", 0))
	assert.Equal(t, 10, GetInt("NOTFOUNDINT", 10))

	assert.Equal(t, true, GetBool("BOOL1", false))
	assert.Equal(t, true, GetBool("BOOL2", false))
	assert.Equal(t, true, GetBool("BOOL3", false))
	assert.Equal(t, false, GetBool("BOOL4", true))

	assert.Equal(t, true, GetBool("NOTFOUNDBOOL", true))
	assert.Equal(t, false, GetBool("NOTFOUNDBOOL", false))

	os.Remove(envFile)
}
