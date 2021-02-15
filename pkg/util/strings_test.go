package util_test

import (
	"testing"

	"github.com/garethjevans/inspect/pkg/util"
	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	strings := []string{"one", "two", "three"}

	assert.True(t, util.Contains(strings, "one"))
	assert.True(t, util.Contains(strings, "two"))
	assert.True(t, util.Contains(strings, "three"))
	assert.False(t, util.Contains(strings, "four"))
}
