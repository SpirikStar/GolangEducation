package binarysearch

import (
	"slices"
	"testing"
)

func TestFind(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{name: "empty", nums: nil, target: 1, want: -1},
		{name: "empty slice", nums: []int{}, target: 1, want: -1},
		{name: "single hit", nums: []int{42}, target: 42, want: 0},
		{name: "single miss", nums: []int{42}, target: 7, want: -1},
		{name: "first element", nums: []int{1, 3, 5, 7, 9}, target: 1, want: 0},
		{name: "last element", nums: []int{1, 3, 5, 7, 9}, target: 9, want: 4},
		{name: "middle", nums: []int{1, 3, 5, 7, 9}, target: 5, want: 2},
		{name: "miss between", nums: []int{1, 3, 5, 7, 9}, target: 4, want: -1},
		{name: "miss below", nums: []int{1, 3, 5, 7, 9}, target: 0, want: -1},
		{name: "miss above", nums: []int{1, 3, 5, 7, 9}, target: 10, want: -1},
		{name: "duplicates", nums: []int{1, 2, 2, 2, 3}, target: 2, want: -2}, // want=-2: любой валидный индекс
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Find(tt.nums, tt.target)
			if tt.want == -2 {
				if got < 0 || tt.nums[got] != tt.target {
					t.Fatalf("Find(%v, %d) = %d, want index with value %d", tt.nums, tt.target, got, tt.target)
				}
				return
			}
			if got != tt.want {
				t.Fatalf("Find(%v, %d) = %d, want %d", tt.nums, tt.target, got, tt.want)
			}
			std, ok := slices.BinarySearch(tt.nums, tt.target)
			if ok && got != std {
				t.Fatalf("mismatch with slices.BinarySearch: got %d, std %d", got, std)
			}
			if !ok && got != -1 {
				t.Fatalf("expected -1 when slices.BinarySearch misses, got %d", got)
			}
		})
	}
}

func TestFindOK(t *testing.T) {
	i, ok := FindOK([]int{1, 2, 3}, 2)
	if !ok || i != 1 {
		t.Fatalf("FindOK() = (%d, %v), want (1, true)", i, ok)
	}
	_, ok = FindOK([]int{1, 2, 3}, 9)
	if ok {
		t.Fatal("FindOK() should miss")
	}
}

func TestFindOrdered_strings(t *testing.T) {
	words := []string{"apple", "banana", "cherry"}
	if got := FindOrdered(words, "banana"); got != 1 {
		t.Fatalf("FindOrdered strings = %d, want 1", got)
	}
}
