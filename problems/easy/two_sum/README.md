# Two Sum

## Условие

Дан массив целых чисел `nums` и число `target`.
Нужно вернуть индексы двух чисел, которые в сумме дают `target`.

Гарантируется, что существует ровно одно решение.

---

## Примеры

**Input:**
nums = [2,7,11,15], target = 9
**Output:**
[0,1]

---

## Подход

Используется хеш-таблица (map):

1. Проходим по массиву
2. Для каждого числа считаем:
   `complement = target - value`
3. Проверяем:

   * если complement уже есть → нашли ответ
   * иначе сохраняем текущее число в map

---

## Сложность

* Time: O(n)
* Space: O(n)


## Идея

Вместо двойного цикла (O(n²))
```go
func twoSum(nums []int, target int) []int {
    for i := 0; i <= len(nums)-1; i++ {
        for j := i + 1; j < len(nums); j++ {
            if target == nums[i] + nums[j] {
                return []int{i, j}
            }
        }
    }
    return nil
}
```
мы запоминаем уже пройденные числа и ищем пару за O(1).

## Оригинал условия из LeetCode

[Ссылка на условие](https://leetcode.com/problems/two-sum/)
