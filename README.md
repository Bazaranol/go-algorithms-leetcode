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