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

## Команды (Makefile)

| Команда | Что делает |
|---|---|
| `make test` | Запустить все тесты |
| `make bench` | Запустить все бенчмарки |
| `make lint` | Запустить линтер (golangci-lint) |
| `make run name=two_sum` | Тесты одной задачи по имени |
| `make run-bench name=two_sum` | Бенчмарк одной задачи |
| `make new level=easy name=climbing_stairs` | Создать новую задачу |
| `make readme` | Обновить список задач в README |
| `make cover` | Покрытие тестами с визуализацией |

## Установка прекоммит хука
### Создаем ссылку на гит файл
```ln -s ../../scripts/pre-commit.sh .git/hooks/pre-commit``` 
### Делаем исполнительным
```chmod +x scripts/pre-commit.sh``` 

### Либо для всей папки 
```chmod +x scripts/*.sh```

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
