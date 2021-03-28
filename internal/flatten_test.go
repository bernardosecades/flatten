// +build unit

package flatten

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenZeroLevel(t *testing.T) {

	js := "{\"input\":[\"♣\", \"♦\", \"♥\"]}"
	a := make(map[string]interface{})
	err := json.Unmarshal([]byte(js), &a)

	if err != nil {
		assert.Fail(t, err.Error())
	}

	result, depth := ApplyFlatten(a["input"].([]interface{}))
	resultBytes, err := json.Marshal(result)

	if err != nil {
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, 0, depth)
	assert.Equal(t, BytesToString(resultBytes), "[\"♣\",\"♦\",\"♥\"]")

}

func TestFlattenMoreThanZeroLevel(t *testing.T) {

	js := "{\"input\":[1, [[2, 3, [[4], 5], 6, [7, 8]], 9, [10, [[[11]]]], 12, 13], 14]}"
	a := make(map[string]interface{})
	err := json.Unmarshal([]byte(js), &a)

	if err != nil {
		assert.Fail(t, err.Error())
	}

	result, depth := ApplyFlatten(a["input"].([]interface{}))
	resultBytes, err := json.Marshal(result)

	if err != nil {
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, 5, depth)
	assert.Equal(t, BytesToString(resultBytes), "[1,2,3,4,5,6,7,8,9,10,11,12,13,14]")
}

func TestFlattenOneElementButAlotOfLevels(t *testing.T) {

	js := "{\"input\":[[[[[[[3]]]]]]]}"
	a := make(map[string]interface{})
	err := json.Unmarshal([]byte(js), &a)

	if err != nil {
		assert.Fail(t, err.Error())
	}

	result, depth := ApplyFlatten(a["input"].([]interface{}))
	resultBytes, err := json.Marshal(result)

	if err != nil {
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, 6, depth)
	assert.Equal(t, BytesToString(resultBytes), "[3]")
}

func TestFlattenWithMixTypeValues(t *testing.T) {

	js := "{\"input\":[3, [\"berni\", [3], 3.5], \"3\"]}"
	a := make(map[string]interface{})
	err := json.Unmarshal([]byte(js), &a)

	if err != nil {
		assert.Fail(t, err.Error())
	}

	result, depth := ApplyFlatten(a["input"].([]interface{}))
	resultBytes, err := json.Marshal(result)

	if err != nil {
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, 2, depth)
	assert.Equal(t, BytesToString(resultBytes), "[3,\"berni\",3,3.5,\"3\"]")
}

func TestFlattenPepe(t *testing.T) {

	js := "{\"input\":[]}"
	a := make(map[string]interface{})
	err := json.Unmarshal([]byte(js), &a)

	if err != nil {
		assert.Fail(t, err.Error())
	}

	result, depth := ApplyFlatten(a["input"].([]interface{}))

	assert.Equal(t, 0, depth)
	assert.Nil(t, result)
}

func BytesToString(data []byte) string {
	return string(data[:])
}