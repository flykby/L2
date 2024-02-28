package main


import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnwrap(t *testing.T) {
	testCases := []struct {
		name   string
		value  string
		expect string
	}{
		{
			name:   "a4bc2d5e",
			value:  "a4bc2d5e",
			expect: "aaaabccddddde",
		}, {
			name:   "abcd",
			value:  "abcd",
			expect: "abcd",
		}, {
			name:   "qwe\\4\\5",
			value:  "qwe\\4\\5",
			expect: "qwe45",
		}, {
			name:   "qwe\\45",
			value:  "qwe\\45",
			expect: "qwe44444",
		}, {
			name:   "qwe\\\\5",
			value:  "qwe\\\\5",
			expect: "qwe\\\\\\\\\\",
		}, {
			name:   "a4bc2d5e\\",
			value:  "a4bc2d5e\\",
			expect: "",
		}, 
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			str, _ := Unpacking(testCase.value)
			
			assert.Equal(t, testCase.expect, str)
		})
	}
}