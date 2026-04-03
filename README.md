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

### Либо для всей папки 
```chmod +x scripts/*.sh```

## Команда для автогенерации задач в README 
```./scripts/generate_readme.sh```

## Проверка покрытия тестами
### Команда в консоли
```go test ./... -cover```

### Как визуализировать

```go test ./... -coverprofile=coverage.out```
```go tool cover -html=coverage.out```

---

## Список задач (автогенерируется)

<!-- START_PROBLEMS -->

### Easy

- [Fizz Buzz ](./problems/easy/fizz_buzz)
- [Palindrome Number ](./problems/easy/palindrome_number)
- [Roman To Integer ](./problems/easy/roman_to_integer)
- [To Buy And Sell Stock ](./problems/easy/to_buy_and_sell_stock)
- [Two Sum ](./problems/easy/two_sum)

### Medium

- [Count Primes ](./problems/medium/count_primes)
- [Reverse Integer ](./problems/medium/reverse_integer)

### Hard


<!-- END_PROBLEMS -->
