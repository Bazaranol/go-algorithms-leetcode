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
- [Happy Number ](./problems/easy/happy_number)
- [Missing Number ](./problems/easy/missing_number)
- [Move Zeroes ](./problems/easy/move_zeroes)
- [Number Of One Bits ](./problems/easy/number_of_one_bits)
- [Palindrome Number ](./problems/easy/palindrome_number)
- [Power Of Four ](./problems/easy/power_of_four)
- [Power Of Three ](./problems/easy/power_of_three)
- [Power Of Two ](./problems/easy/power_of_two)
- [Remove Duplicates From Sorted Array ](./problems/easy/remove_duplicates_from_sorted_array)
- [Roman To Integer ](./problems/easy/roman_to_integer)
- [Single Number ](./problems/easy/single_number)
- [To Buy And Sell Stock ](./problems/easy/to_buy_and_sell_stock)
- [Two Sum ](./problems/easy/two_sum)

### Medium

- [Count Primes ](./problems/medium/count_primes)
- [Maximum Subarray ](./problems/medium/maximum_subarray)
- [Reverse Integer ](./problems/medium/reverse_integer)
- [Three Sum ](./problems/medium/three_sum)
- [Two Sum 2 ](./problems/medium/two_sum_2)

### Hard


<!-- END_PROBLEMS -->
