package functions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateToken(t *testing.T) {
	assert := assert.New(t)

	assert.Nil(ValidateToken("12345678901234567890123456789012"))
	assert.NotNil(ValidateToken(""))
	assert.NotNil(ValidateToken("123"))
	assert.NotNil(ValidateToken("1234567890123456789012345678901212"))
}
