# Algorithms in Go

Репозиторий с решениями алгоритмических задач (на LeetCode) на Go.

## Структура

* `problems/easy` — простые задачи
* `problems/medium` — средние
* `problems/hard` — сложные

Каждая задача содержит:

* решение (`solution.go`)
* тесты (`solution_test.go`)
* описание (`README.md`)

---

##  Как запустить

### 1. Клонировать репозиторий

```bash
git clone https://github.com/bazaranol/go-algorithms-leetcode.git
cd go-algorithms-leetcode
```

### 2. Запустить тесты

```bash
go test ./...
```

---

## Пример запуска конкретной задачи

```bash
go test ./problems/easy/roman_to_integer
```

---



## Стек

* Go (Golang)
* Стандартная библиотека (`testing`)

## Команда для создания нового решения

```./scripts/create_problem.sh easy roman_to_integer```

## Команда для запуска всех бенчмарков 

```go test ./... -bench=.```

## Установка прекоммит хука
### Создаем ссылку на гит файл
```ln -s ../../scripts/pre-commit.sh .git/hooks/pre-commit``` 
### Делаем исполнительным
```chmod +x scripts/pre-commit.sh``` 

## Проверка покрытия тестами
### Команда в консоли
```go test ./... -cover```

### Как визуализировать

```go test ./... -cover```
```go tool cover -html=coverage.out```
