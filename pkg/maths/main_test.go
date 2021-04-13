package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
	assert := assert.New(t)
	result := Add(2, 3)
	assert.Equal(5, result, "2 + 3 should be 5")
}
