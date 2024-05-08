package utils_test

import (
	"testing"

	"github.com/sorrawichYooboon/gogeneric/utils"
	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	a := 1
	b := 100

	result := utils.Min(a, b)

	assert.Equal(t, a, result)
}

func TestMax(t *testing.T) {
	a := 1
	b := 100

	result := utils.Max(a, b)

	assert.Equal(t, b, result)
}
