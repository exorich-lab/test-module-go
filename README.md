# TextUtils

[![Go Reference](https://pkg.go.dev/badge/github.com/chugunok/textutils.svg)](https://pkg.go.dev/github.com/chugunok/textutils)
[![Go Report Card](https://goreportcard.com/badge/github.com/chugunok/textutils)](https://goreportcard.com/report/github.com/chugunok/textutils)
[![CI](https://github.com/chugunok/textutils/workflows/CI/badge.svg)](https://github.com/chugunok/textutils/actions)

TextUtils - это Go модуль, предоставляющий набор утилит для обработки и анализа текста. Модуль включает функции для статистического анализа текста, проверки палиндромов, извлечения email адресов и URL, а также другие полезные операции со строками.

## Установка

```bash
go get github.com/chugunok/textutils
```

## Возможности

- 📊 **Анализ текста**: подсчет слов, символов, предложений и абзацев
- 🔄 **Манипуляции со строками**: переворот строк, капитализация слов
- 🔍 **Проверка палиндромов**: поддержка кириллицы и латиницы
- 📧 **Извлечение данных**: email адреса и URL из текста
- 🧹 **Очистка текста**: удаление лишних пробелов

## Использование

### Анализ текста

```go
package main

import (
    "fmt"
    "github.com/chugunok/textutils"
)

func main() {
    text := "Пример текста для анализа. Он содержит несколько предложений!"
    
    stats := textutils.AnalyzeText(text)
    fmt.Printf("Слов: %d\n", stats.Words)
    fmt.Printf("Символов: %d\n", stats.Characters)
    fmt.Printf("Предложений: %d\n", stats.Sentences)
    fmt.Printf("Абзацев: %d\n", stats.Paragraphs)
}
```

### Проверка палиндромов

```go
fmt.Println(textutils.IsPalindrome("топот"))        // true
fmt.Println(textutils.IsPalindrome("hello"))        // false
fmt.Println(textutils.IsPalindrome("A man, a plan, a canal: Panama")) // true
```

### Манипуляции со строками

```go
// Переворот строки
fmt.Println(textutils.ReverseString("hello"))  // "olleh"

// Капитализация слов
fmt.Println(textutils.CapitalizeWords("hello world"))  // "Hello World"

// Удаление лишних пробелов
fmt.Println(textutils.RemoveExtraSpaces("  hello   world  "))  // "hello world"
```

### Извлечение данных

```go
text := "Свяжитесь с нами по email: admin@example.com или посетите https://example.com"

// Извлечение email адресов
emails := textutils.ExtractEmails(text)
fmt.Println(emails)  // ["admin@example.com"]

// Извлечение URL
urls := textutils.ExtractURLs(text)
fmt.Println(urls)    // ["https://example.com"]
```

## API Documentation

### `AnalyzeText(text string) WordCounter`

Анализирует текст и возвращает статистику:

- `Words`: количество слов
- `Characters`: количество символов (без пробелов)
- `Sentences`: количество предложений
- `Paragraphs`: количество абзацев

### `ReverseString(s string) string`

Переворачивает строку, поддерживает Unicode символы.

### `IsPalindrome(s string) bool`

Проверяет, является ли строка палиндромом. Игнорирует пробелы, знаки препинания и регистр букв.

### `CapitalizeWords(s string) string`

Делает первую букву каждого слова заглавной.

### `RemoveExtraSpaces(s string) string`

Удаляет лишние пробелы, табы и переносы строк в тексте.

### `ExtractEmails(s string) []string`

Извлекает все email адреса из текста.

### `ExtractURLs(s string) []string`

Извлекает все URL (http/https) из текста.

## Примеры

Запустите пример использования:

```bash
cd example
go run main.go
```

## Тестирование

Запустите тесты:

```bash
go test -v
```

Запустите бенчмарки:

```bash
go test -bench=.
```

## Тестовое покрытие

```bash
go test -cover
```

## Вклад

Мы приветствуем вклад в развитие проекта! Пожалуйста, создайте issue для обсуждения изменений или отправьте pull request.

## Лицензия

Этот проект распространяется под лицензией MIT. Подробности в файле [LICENSE](LICENSE).

## Разработка

Для разработки модуля:

1. Клонируйте репозиторий
2. Установите зависимости: `go mod tidy`
3. Запустите тесты: `go test -v`
4. Вносите изменения и не забывайте про тесты!

## TODO

- [ ] Добавить поддержку更多 языков
- [ ] Реализовать функцию определения языка текста
- [ ] Добавить морфологический анализ
- [ ] Улучшить производительность для больших текстов