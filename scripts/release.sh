#!/bin/bash

# Скрипт для публикации Go модуля на GitHub
# Требует установленного gh CLI и авторизации

set -e

# Проверяем наличие gh CLI
if ! command -v gh &> /dev/null; then
    echo "Ошибка: GitHub CLI (gh) не установлен. Установите его с https://cli.github.com/"
    exit 1
fi

# Проверяем авторизацию
if ! gh auth status &> /dev/null; then
    echo "Ошибка: Не авторизованы в GitHub CLI. Выполните 'gh auth login'"
    exit 1
fi

# Получаем информацию о модуле
MODULE_NAME=$(go list -m)
VERSION=${1:-"v1.0.0"}

echo "Подготовка к публикации модуля $MODULE_NAME версии $VERSION"

# Добавляем удаленный репозиторий
echo "Добавление удаленного репозитория..."
git remote add origin https://github.com/$MODULE_NAME.git
git push -u origin main

# Создаем релиз
echo "Создание релиза $VERSION..."
gh release create $VERSION --title "Release $VERSION" --notes "Первый релиз textutils модуля

## Возможности:
- Анализ текста (подсчет слов, символов, предложений, абзацев)
- Проверка палиндромов
- Манипуляции со строками (переворот, капитализация)
- Извлечение email адресов и URL
- Очистка текста от лишних пробелов

## Установка:
\`\`\`bash
go get $MODULE_NAME@$VERSION
\`\`\`

## Использование:
\`\`\`go
import \"$MODULE_NAME\"
\`\`\`"

echo "✅ Модуль успешно опубликован на GitHub!"
echo "📦 Репозиторий: https://github.com/$MODULE_NAME"
echo "🏷️ Релиз: https://github.com/$MODULE_NAME/releases/tag/$VERSION"
echo "📚 Документация: https://pkg.go.dev/$MODULE_NAME@$VERSION"