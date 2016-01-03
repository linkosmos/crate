package crate

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var filterTestFunc = func(input string) bool {
	return strings.HasPrefix(input, "some")
}

var filterTests = []struct {
	input   string
	present bool
}{
	{"some.1", false},
	{"some.2", false},
	{"no", true},
}

func TestFilterFunc(t *testing.T) {
	c := New(1)
	c.AttachFilter(filterTestFunc)

	for _, test := range filterTests {
		c.Add(test.input)
		assert.Equal(t, test.present, c.Exist(test.input),
			fmt.Sprintf("Expected %t for %s", test.present, test.input))
	}

	// Only one item shouldn't be filtered
	assert.Len(t, c.Map.Data, 1)
}

var addTests = []struct {
	input    string
	expected bool
}{
	{"one", true},
	{"one", true},
	{"two", true},
}

func TestAdd(t *testing.T) {
	c := New(100)
	for _, test := range addTests {
		c.Add(test.input)

		got := c.Exist(test.input)

		assert.Equal(t, got, test.expected)
	}
}
