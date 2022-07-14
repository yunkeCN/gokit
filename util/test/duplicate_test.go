package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yunkeCN/gokit/util"
)

func TestDeDuplicateUintSlice(t *testing.T) {
	var input = []uint{1, 2, 3, 4, 5, 3, 2, 1, 3, 4}
	assert.Equal(t, util.DeDuplicateUintSlice(input), []uint{1, 2, 3, 4, 5})
}
