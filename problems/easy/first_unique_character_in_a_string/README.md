# First Unique Character in a String

## Условие

Найти **индекс** первого символа в строке, который встречается **ровно один раз**. Если такого символа нет, вернуть `-1`.

## Примеры

Input: `s = "leetcode"`  
Output: `0`  
Объяснение: 'l' встречается один раз на позиции 0

Input: `s = "loveleetcode"`  
Output: `2`  
Объяснение: 'v' первый уникальный на позиции 2

Input: `s = "aabb"`  
Output: `-1`  
Объяснение: Нет уникальных символов

---

## Решение: Подсчёт частоты

```go
func FirstUniqueCharacterInAString(s string) int {
	counter := make(map[rune]int)

	// Считаем частоту каждого символа
	for _, v := range s {
		counter[v]++
	}

	// Ищем первый символ с частотой 1
	for i, v := range s {
		if counter[v] == 1 {
			return i
		}
	}
	return -1
}
```

**Идея:** Два прохода — первый подсчитываем частоты, второй ищем первый с частотой 1.

## Сложность

| Метрика | Значение |
|---------|----------|
| **Time** | O(n) |
| **Space** | O(1) |

Два прохода по строке, размер хеш-таблицы ограничен размером алфавита.

---

## Альтернативы

### Сортировка (O(n log n))
Отсортировать и проверить соседних соседей, но медленнее.

### С сохранением индекса (без второго прохода)
```go
minIndex := len(s)
for char, count := range counter {
    if count == 1 {
        for i, v := range s {
            if v == char {
                minIndex = min(minIndex, i)
                break
            }
        }
    }
}
```
Более сложно и медленнее.

---

## Оригинал условия из LeetCode

[First Unique Character in a String](https://leetcode.com/problems/first-unique-character-in-a-string/)
