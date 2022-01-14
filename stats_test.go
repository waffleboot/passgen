package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordStats(t *testing.T) {
	stats := makeWordStats("asshole")

	assert.Equal(t, 6, len(stats.replace))
	assert.Equal(t, 3, len(stats.numbers))
	assert.Equal(t, 3, len(stats.symbols))

	assert.Equal(t, '@', stats.replace[0].to)
	assert.Equal(t, '$', stats.replace[1].to)
	assert.Equal(t, '$', stats.replace[2].to)
	assert.Equal(t, '0', stats.replace[4].to)
	assert.Equal(t, '1', stats.replace[5].to)
	assert.Equal(t, '3', stats.replace[6].to)
}
