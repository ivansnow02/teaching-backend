package encrypt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncEmail(t *testing.T) {
	enc, err := EncEmail("test@example.com")
	assert.Nil(t, err)
	dec, err := DecEmail(enc)
	assert.Nil(t, err)
	assert.Equal(t, "test@example.com", dec)
}
