package utils

import "testing"

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "빈 문자열",
			input:    "",
			expected: true,
		},
		{
			name:     "공백만 있는 문자열",
			input:    "   ",
			expected: true,
		},
		{
			name:     "일반 문자열",
			input:    "hello",
			expected: false,
		},
		{
			name:     "공백이 있는 문자열",
			input:    "  hello  ",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsEmpty(tt.input)
			if result != tt.expected {
				t.Errorf("IsEmpty(%q) = %v, expected %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestIsNotEmpty(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "빈 문자열",
			input:    "",
			expected: false,
		},
		{
			name:     "공백만 있는 문자열",
			input:    "   ",
			expected: false,
		},
		{
			name:     "일반 문자열",
			input:    "hello",
			expected: true,
		},
		{
			name:     "공백이 있는 문자열",
			input:    "  hello  ",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsNotEmpty(tt.input)
			if result != tt.expected {
				t.Errorf("IsNotEmpty(%q) = %v, expected %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestDefaultIfEmpty(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		defaultValue string
		expected     string
	}{
		{
			name:         "빈 문자열",
			input:        "",
			defaultValue: "default",
			expected:     "default",
		},
		{
			name:         "공백만 있는 문자열",
			input:        "   ",
			defaultValue: "default",
			expected:     "default",
		},
		{
			name:         "일반 문자열",
			input:        "hello",
			defaultValue: "default",
			expected:     "hello",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DefaultIfEmpty(tt.input, tt.defaultValue)
			if result != tt.expected {
				t.Errorf("DefaultIfEmpty(%q, %q) = %q, expected %q", tt.input, tt.defaultValue, result, tt.expected)
			}
		})
	}
}