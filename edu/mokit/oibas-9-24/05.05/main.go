package main

import (
	"log"
	"sort"
)

func BinarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left < right {
		log.Println(left, right)
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
func main() {
	data := []int{53, 22, 12, 90}

	stdData := make([]int, len(data))
	copy(stdData, data)
	sort.Ints(stdData)
	log.Println(stdData)
	log.Println(BinarySearch(data, 22))
}
