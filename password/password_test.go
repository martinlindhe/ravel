package password

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPassword(t *testing.T) {

	clear := "password"
	hash1 := Make(clear)
	hash2 := Make(clear)
	hash3 := Make(clear)

	// always 60 characters
	assert.Equal(t, 60, len(hash1))
	assert.Equal(t, 60, len(hash2))
	assert.Equal(t, 60, len(hash3))

	// never same hash from same input
	assert.Equal(t, false, hash1 == hash2)
	assert.Equal(t, false, hash2 == hash3)
	assert.Equal(t, false, hash3 == hash1)
}
