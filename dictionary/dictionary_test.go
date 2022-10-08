package dictionary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDictionariesContainExpectedAmountOfEntries(t *testing.T) {
	combined, first, latter := GetOldSwedishDictionary()

	assert.Equal(t, 22572, len(first), "Incorrect amount of entries for first volumes!")
	assert.Equal(t, 19172, len(latter), "Incorrect amount of entries for latter volumes!")
	assert.Equal(t, 41744, len(combined), "Incorrect amount of entries for combined volumes!")
	assert.Equal(t, len(first)+len(latter), len(combined), "Combined should consist of first & latter volumes!")
}
