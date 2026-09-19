package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSubcommand(t *testing.T) {
	assert.True(t, isSubcommand("compile"))
	assert.True(t, isSubcommand("format"))
	assert.False(t, isSubcommand("compilex"))
	assert.False(t, isSubcommand("--help"))
	assert.False(t, isSubcommand(""))
}
