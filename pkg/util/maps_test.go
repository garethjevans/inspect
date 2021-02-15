package util_test

import (
	"github.com/garethjevans/inspect/pkg/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAllKeys(t *testing.T) {
	map1 := map[string]string{}
	map1["key1"] = "value1"
	map1["key2"] = "value2"

	map2 := map[string]string{}
	map2["key1"] = "value1"
	map2["key3"] = "value3"

	allkeys := util.AllKeys(map1, map2)
	assert.Equal(t, allkeys, []string{"key1", "key2", "key3"})
}