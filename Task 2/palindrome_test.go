package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Simple palindrome",
			input:    "radar",
			expected: true,
		},
		{
			name:     "Phrase with spaces and punctuation",
			input:    "A man, a plan, a canal: Panama",
			expected: true,
		},
		{
			name:     "Mixed case palindrome",
			input:    "RaCeCaR",
			expected: true,
		},
		{
			name:     "Not a palindrome",
			input:    "hello",
			expected: false,
		},
		{
			name:     "Empty string",
			input:    "",
			expected: true,
		},
		{
			name:     "Single character",
			input:    "a",
			expected: true,
		},
		{
			name:     "Palindrome with numbers",
			input:    "12321",
			expected: true,
		},
		{
			name:     "Palindrome with special characters",
			input:    "Madam, I'm Adam",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isPalindrome(tt.input)
			if result != tt.expected {
				t.Errorf("isPalindrome() = %v, want %v", result, tt.expected)
			}
		})
	}
}
