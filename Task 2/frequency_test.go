package main

import (
	"reflect"
	"testing"
)

func TestFrequency(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string]int
	}{
		{
			name:  "Simple sentence",
			input: "Hello world hello",
			expected: map[string]int{
				"hello": 2,
				"world": 1,
			},
		},
		{
			name:  "With punctuation",
			input: "Hello, world! Hello.",
			expected: map[string]int{
				"hello": 2,
				"world": 1,
			},
		},
		{
			name:  "Mixed case",
			input: "Hello WORLD hello World",
			expected: map[string]int{
				"hello": 2,
				"world": 2,
			},
		},
		{
			name:     "Empty string",
			input:    "",
			expected: map[string]int{},
		},
		{
			name:  "Multiple spaces",
			input: "hello   world  hello",
			expected: map[string]int{
				"hello": 2,
				"world": 1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := frequency(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("frequency() = %v, want %v", result, tt.expected)
			}
		})
	}
}
