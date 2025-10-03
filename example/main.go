package main

import (
	"fmt"

	"github.com/exorich-lab/test-module-go"
)

func main() {
	// Пример текста для анализа
	sampleText := `Привет мир! Это пример текста для демонстрации возможностей нашего модуля textutils.
	
	Наш модуль предоставляет различные утилиты для работы с текстом: анализ статистики, проверку палиндромов,
	извлечение email адресов и URL, а также множество других полезных функций.
	
	Для связи используйте: admin@example.com или support@domain.org.
	Также посетите наш сайт: https://example.com и документацию: https://docs.example.com/api.`

	fmt.Println("=== Демонстрация модуля textutils ===\n")

	// 1. Анализ текста
	fmt.Println("1. Анализ текста:")
	stats := textutils.AnalyzeText(sampleText)
	fmt.Printf("   Слов: %d\n", stats.Words)
	fmt.Printf("   Символов (без пробелов): %d\n", stats.Characters)
	fmt.Printf("   Предложений: %d\n", stats.Sentences)
	fmt.Printf("   Абзацев: %d\n", stats.Paragraphs)
	fmt.Println()

	// 2. Проверка палиндромов
	fmt.Println("2. Проверка палиндромов:")
	palindromes := []string{
		"топот",
		"A man, a plan, a canal: Panama",
		"hello",
		"казак",
	}
	for _, word := range palindromes {
		result := textutils.IsPalindrome(word)
		fmt.Printf("   '%s' - палиндром: %t\n", word, result)
	}
	fmt.Println()

	// 3. Переворот строки
	fmt.Println("3. Переворот строки:")
	original := "Привет мир"
	reversed := textutils.ReverseString(original)
	fmt.Printf("   Оригинал: '%s'\n", original)
	fmt.Printf("   Перевернуто: '%s'\n", reversed)
	fmt.Println()

	// 4. Капитализация слов
	fmt.Println("4. Капитализация слов:")
	sentence := "hello world from go module"
	capitalized := textutils.CapitalizeWords(sentence)
	fmt.Printf("   Оригинал: '%s'\n", sentence)
	fmt.Printf("   С капитализацией: '%s'\n", capitalized)
	fmt.Println()

	// 5. Удаление лишних пробелов
	fmt.Println("5. Удаление лишних пробелов:")
	messyText := "  Это    текст	с   лишними\n\nпробелами  "
	cleaned := textutils.RemoveExtraSpaces(messyText)
	fmt.Printf("   Оригинал: '%s'\n", messyText)
	fmt.Printf("   Очищено: '%s'\n", cleaned)
	fmt.Println()

	// 6. Извлечение email адресов
	fmt.Println("6. Извлечение email адресов:")
	emails := textutils.ExtractEmails(sampleText)
	fmt.Printf("   Найдено email адресов: %d\n", len(emails))
	for i, email := range emails {
		fmt.Printf("   %d. %s\n", i+1, email)
	}
	fmt.Println()

	// 7. Извлечение URL
	fmt.Println("7. Извлечение URL:")
	urls := textutils.ExtractURLs(sampleText)
	fmt.Printf("   Найдено URL: %d\n", len(urls))
	for i, url := range urls {
		fmt.Printf("   %d. %s\n", i+1, url)
	}
	fmt.Println()

	// 8. Интерактивный режим
	fmt.Println("8. Интерактивный режим:")
	fmt.Println("Введите текст для анализа (или 'exit' для выхода):")
	
	var input string
	for {
		fmt.Print("> ")
		_, err := fmt.Scanln(&input)
		if err != nil {
			// Обработка пустой строки
			continue
		}
		
		if input == "exit" {
			break
		}
		
		// Анализ введенного текста
		inputStats := textutils.AnalyzeText(input)
		fmt.Printf("   Слов: %d, Символов: %d, Предложений: %d\n", 
			inputStats.Words, inputStats.Characters, inputStats.Sentences)
		
		// Проверка на палиндром
		if textutils.IsPalindrome(input) {
			fmt.Printf("   '%s' - это палиндром!\n", input)
		}
	}
	
	fmt.Println("\nСпасибо за использование textutils!")
}