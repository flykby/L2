package main


import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestAnagram(t *testing.T) {
	data := []struct {
		input    []string
		expect map[string][]string
	}{
		{
			input:	[]string{"тяпка", "ПЯТАК", "Пятка", "бетон", "СЛИТОК", "столик", "листок"},
			expect: map[string][]string{"слиток": {"листок", "слиток", "столик"}, "тяпка": {"пятак", "пятка", "тяпка"}},
		},
		{
			input:	[]string{},
			expect: map[string][]string{},
		},
	}

	for _, d := range data {
		t.Run("find anagram", func(t *testing.T) {
			result := isAnagram(d.input)

			assert.Equal(t, d.expect, result)
		})
	}
}