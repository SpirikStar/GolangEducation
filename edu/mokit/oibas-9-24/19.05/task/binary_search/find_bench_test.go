package binarysearch

import (
	"slices"
	"testing"
)

func makeSorted(n int) []int {
	nums := make([]int, n)
	for i := range nums {
		nums[i] = i * 2
	}
	return nums
}

func BenchmarkFind(b *testing.B) {
	nums := makeSorted(1_000_000)
	target := 999_998 // существует
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Find(nums, target)
	}
}

func BenchmarkSlicesBinarySearch(b *testing.B) {
	nums := makeSorted(1_000_000)
	target := 999_998
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slices.BinarySearch(nums, target)
	}
}

func BenchmarkFind_miss(b *testing.B) {
	nums := makeSorted(1_000_000)
	target := 999_999 // между элементами
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Find(nums, target)
	}
}
