.PHONY: help test benchmark lint build example clean release install

# Переменные
MODULE_NAME := $(shell go list -m)
VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo "v1.0.0")
GO_FILES := $(shell find . -name "*.go" -type f)

# По умолчанию
help: ## Показать доступные команды
	@echo "Доступные команды:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

# Тестирование
test: ## Запустить все тесты
	go test -v -race -coverprofile=coverage.out -covermode=atomic ./...

test-short: ## Запустить быстрые тесты
	go test -short -v ./...

benchmark: ## Запустить бенчмарки
	go test -bench=. -benchmem ./...

coverage: test ## Показать покрытие тестами
	go tool cover -html=coverage.out -o coverage.html
	@echo "Отчет о покрытии сохранен в coverage.html"

# Проверка кода
lint: ## Запустить линтер
	golangci-lint run

vet: ## Запустить go vet
	go vet ./...

fmt: ## Форматировать код
	go fmt ./...
	@if [ "$(gofmt -s -l . | wc -l)" -gt 0 ]; then \
		echo "Следующие файлы не отформатированы:"; \
		gofmt -s -l .; \
		exit 1; \
	fi

# Сборка
build: ## Собрать модуль
	go build -v ./...

example: ## Собрать и запустить пример
	cd example && go run main.go

install: ## Установить модуль локально
	go install ./...

# Очистка
clean: ## Очистить временные файлы
	rm -f coverage.out coverage.html
	rm -f *.test
	go clean -cache -testcache

# Публикация
release: ## Опубликовать релиз (требует gh CLI)
	@echo "Публикация версии $(VERSION)"
	./scripts/release.sh $(VERSION)

tag: ## Создать git тег
	git tag -a $(VERSION) -m "Release $(VERSION)"
	git push origin $(VERSION)

# Разработка
deps: ## Скачать зависимости
	go mod download
	go mod tidy

verify: ## Проверить зависимости
	go mod verify

update-deps: ## Обновить зависимости
	go get -u ./...
	go mod tidy

# Полная проверка перед коммитом
pre-commit: fmt vet test lint ## Полная проверка кода
	@echo "✅ Все проверки пройдены!"

# Инициализация проекта для новой версии
setup-dev: ## Настроить окружение для разработки
	@echo "Настройка окружения для разработки..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@echo "✅ Среда разработки настроена!"

# Показать информацию о модуле
info: ## Показать информацию о модуле
	@echo "Модуль: $(MODULE_NAME)"
	@echo "Версия: $(VERSION)"
	@echo "Go файлы: $(words $(GO_FILES))"
	@go version