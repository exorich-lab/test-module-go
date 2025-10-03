// Package textutils предоставляет утилиты для обработки и анализа текста
package textutils

import (
	"regexp"
	"strings"
	"unicode"
)

// WordCounter структура для подсчета статистики слов в тексте
type WordCounter struct {
	Words      int
	Characters int
	Sentences  int
	Paragraphs int
}

// AnalyzeText анализирует текст и возвращает статистику
func AnalyzeText(text string) WordCounter {
	counter := WordCounter{}
	
	// Подсчет символов (без пробелов)
	counter.Characters = len(strings.ReplaceAll(text, " ", ""))
	
	// Подсчет слов
	wordRegex := regexp.MustCompile(`\b\w+\b`)
	words := wordRegex.FindAllString(text, -1)
	counter.Words = len(words)
	
	// Подсчет предложений
	sentenceRegex := regexp.MustCompile(`[.!?]+`)
	sentences := sentenceRegex.FindAllString(text, -1)
	if len(sentences) > 0 {
		counter.Sentences = len(sentences)
	} else if len(strings.TrimSpace(text)) > 0 {
		counter.Sentences = 1
	}
	
	// Подсчет абзацев
	if text == "" {
		counter.Paragraphs = 1
	} else {
		paragraphs := strings.Split(text, "\n\n")
		counter.Paragraphs = len(paragraphs)
	}
	
	return counter
}

// ReverseString переворачивает строку
func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// IsPalindrome проверяет, является ли строка палиндромом
func IsPalindrome(s string) bool {
	// Удаляем все не-буквенные символы и приводим к нижнему регистру
	cleaned := regexp.MustCompile(`[^a-zA-Zа-яА-Я0-9]`).ReplaceAllString(s, "")
	cleaned = strings.ToLower(cleaned)
	
	return cleaned == ReverseString(cleaned)
}

// CapitalizeWords делает первую букву каждого слова заглавной
func CapitalizeWords(s string) string {
	if s == "" {
		return ""
	}
	
	words := strings.Fields(s)
	for i, word := range words {
		if len(word) > 0 {
			runes := []rune(word)
			runes[0] = unicode.ToUpper(runes[0])
			words[i] = string(runes)
		}
	}
	return strings.Join(words, " ")
}

// RemoveExtraSpaces удаляет лишние пробелы в тексте
func RemoveExtraSpaces(s string) string {
	spaceRegex := regexp.MustCompile(`\s+`)
	return spaceRegex.ReplaceAllString(strings.TrimSpace(s), " ")
}

// ExtractEmails извлекает все email адреса из текста
func ExtractEmails(s string) []string {
	emailRegex := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
	emails := emailRegex.FindAllString(s, -1)
	if emails == nil {
		return []string{}
	}
	return emails
}

// ExtractURLs извлекает все URL из текста
func ExtractURLs(s string) []string {
	urlRegex := regexp.MustCompile(`https?://[^\s]+`)
	urls := urlRegex.FindAllString(s, -1)
	if urls == nil {
		return []string{}
	}
	return urls
}