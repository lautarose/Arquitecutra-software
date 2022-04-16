package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDivision(t *testing.T) {
	assert := assert.New(t)

	res, _ := Division(4, 2)

	assert.Equal(t, res, 2)
}
