package textutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnalyzeText(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		expected WordCounter
	}{
		{
			name: "Simple text",
			text: "Hello world.",
			expected: WordCounter{
				Words:      2,
				Characters: 11,
				Sentences:  1,
				Paragraphs: 1,
			},
		},
		{
			name: "Multiple sentences",
			text: "Hello world. How are you? I'm fine!",
			expected: WordCounter{
				Words:      8,
				Characters: 29,
				Sentences:  3,
				Paragraphs: 1,
			},
		},
		{
			name: "Multiple paragraphs",
			text: "First paragraph.\n\nSecond paragraph with more words.",
			expected: WordCounter{
				Words:      7,
				Characters: 46,
				Sentences:  2,
				Paragraphs: 2,
			},
		},
		{
			name: "Empty text",
			text: "",
			expected: WordCounter{
				Words:      0,
				Characters: 0,
				Sentences:  0,
				Paragraphs: 1,
			},
		},
		{
			name: "Text without sentence markers",
			text: "Just some words without punctuation",
			expected: WordCounter{
				Words:      5,
				Characters: 31,
				Sentences:  1,
				Paragraphs: 1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AnalyzeText(tt.text)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestReverseString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Simple word",
			input:    "hello",
			expected: "olleh",
		},
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "Unicode characters",
			input:    "привет",
			expected: "тевирп",
		},
		{
			name:     "Mixed characters",
			input:    "Go123",
			expected: "321oG",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ReverseString(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Simple palindrome",
			input:    "racecar",
			expected: true,
		},
		{
			name:     "Not a palindrome",
			input:    "hello",
			expected: false,
		},
		{
			name:     "Palindrome with spaces and punctuation",
			input:    "A man, a plan, a canal: Panama",
			expected: true,
		},
		{
			name:     "Cyrillic palindrome",
			input:    "топот",
			expected: true,
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsPalindrome(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestCapitalizeWords(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Simple sentence",
			input:    "hello world",
			expected: "Hello World",
		},
		{
			name:     "Already capitalized",
			input:    "Hello World",
			expected: "Hello World",
		},
		{
			name:     "Mixed case",
			input:    "hELLo wORLd",
			expected: "HELLo WORLd",
		},
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "Multiple spaces",
			input:    "hello   world",
			expected: "Hello World",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CapitalizeWords(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestRemoveExtraSpaces(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Multiple spaces",
			input:    "hello   world",
			expected: "hello world",
		},
		{
			name:     "Tabs and newlines",
			input:    "hello\t\tworld\n\n",
			expected: "hello world",
		},
		{
			name:     "Leading and trailing spaces",
			input:    "  hello world  ",
			expected: "hello world",
		},
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "Only spaces",
			input:    "   ",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RemoveExtraSpaces(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestExtractEmails(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "Single email",
			input:    "Contact me at user@example.com",
			expected: []string{"user@example.com"},
		},
		{
			name:     "Multiple emails",
			input:    "Contact user1@example.com or user2@example.org",
			expected: []string{"user1@example.com", "user2@example.org"},
		},
		{
			name:     "No emails",
			input:    "No emails here",
			expected: []string{},
		},
		{
			name:     "Complex emails",
			input:    "Contact test.email+tag@domain.co.uk",
			expected: []string{"test.email+tag@domain.co.uk"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ExtractEmails(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestExtractURLs(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "Single HTTP URL",
			input:    "Visit http://example.com",
			expected: []string{"http://example.com"},
		},
		{
			name:     "Single HTTPS URL",
			input:    "Visit https://example.com",
			expected: []string{"https://example.com"},
		},
		{
			name:     "Multiple URLs",
			input:    "Visit http://example.com and https://test.org",
			expected: []string{"http://example.com", "https://test.org"},
		},
		{
			name:     "Complex URLs",
			input:    "Visit https://example.com/path?query=value&other=test",
			expected: []string{"https://example.com/path?query=value&other=test"},
		},
		{
			name:     "No URLs",
			input:    "No URLs here",
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ExtractURLs(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// Бенчмарк для производительности
func BenchmarkAnalyzeText(b *testing.B) {
	text := `This is a sample text for benchmarking. It contains multiple sentences and paragraphs.
	
	We want to test how efficiently our AnalyzeText function processes larger texts.
	This helps us understand the performance characteristics of our implementation.
	
	The benchmark should give us insights into potential optimizations.`
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		AnalyzeText(text)
	}
}

func BenchmarkReverseString(b *testing.B) {
	text := "This is a sample text for benchmarking the reverse string function"
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ReverseString(text)
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	text := "A man, a plan, a canal: Panama"
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsPalindrome(text)
	}
}