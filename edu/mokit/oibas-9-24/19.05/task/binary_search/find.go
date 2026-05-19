// Package binarysearch реализует бинарный поиск в отсортированном слайсе.
package binarysearch

// Ordered — типы, поддерживающие операторы < и ==.
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
	~float32 | ~float64 |
	~string
}

// Find ищет target в nums, отсортированном по неубыванию.
// Возвращает индекс элемента или -1, если target отсутствует.
//
// Временная сложность: O(log n), память: O(1).
//
// Инвариант цикла: если target присутствует в nums,
// его индекс лежит в закрытом интервале [left, right].
func Find(nums []int, target int) int {
	return FindOrdered(nums, target)
}

// FindOrdered — обобщённый бинарный поиск для упорядочиваемых типов.
func FindOrdered[T Ordered](nums []T, target T) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		switch {
		case nums[mid] == target:
			return mid
		case nums[mid] < target:
			left = mid + 1
		default:
			right = mid - 1
		}
	}
	return -1
}

// FindOK — вариант API с явным признаком успеха.
func FindOK(nums []int, target int) (index int, ok bool) {
	if i := Find(nums, target); i >= 0 {
		return i, true
	}
	return 0, false
}
