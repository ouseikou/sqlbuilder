package facade

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNormalizeArrValue(t *testing.T) {
	orgArr1 := []interface{}{"a", "b", "c"}
	finalStr1 := Normalize1Arr(orgArr1)
	assert.EqualValues(t, "'a','b','c'", finalStr1)

	orgArr2 := []interface{}{1, 2, 3}
	finalStr2 := Normalize1Arr(orgArr2)
	assert.EqualValues(t, "1,2,3", finalStr2)

	orgArr3 := []interface{}{true, false, true}
	finalStr3 := Normalize1Arr(orgArr3)
	assert.EqualValues(t, "true,false,true", finalStr3)

	finalStr4 := Normalize1Arr(nil)
	assert.EqualValues(t, "", finalStr4)
}
