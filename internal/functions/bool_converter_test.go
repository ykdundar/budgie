package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBoolConverter(t *testing.T) {
	assert := assert.New(t)

	trueCase := BoolConverter(true)
	falseCase := BoolConverter(false)

	assert.Equal(1, trueCase, "they should be equal")
	assert.Equal(0, falseCase, "they should be equal")

	/*
	if trueCase != 1 {
		t.Errorf("BoolConverter(true) is expected to return 1, but returned %d", trueCase)
	}

	if falseCase != 0 {
		t.Errorf("BoolConverter(false) is expected to return 0, but returned %d", falseCase)
	}
	 */
}
