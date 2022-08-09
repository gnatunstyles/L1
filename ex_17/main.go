package main

import "fmt"

func search(nums []int, target int) int {
	var middle, left int
	right := len(nums) - 1
	if nums[left] == target {
		return left
	}
	if nums[right] == target {
		return right
	}

	for left < right-1 {
		middle = (left + right) / 2
		if target > nums[middle] {
			left = middle
		} else {
			right = middle
		}
	}

	if nums[right] == target {
		return right
	} else {
		return -1
	}
}

func main() {
	fmt.Println(search([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 4))
}
