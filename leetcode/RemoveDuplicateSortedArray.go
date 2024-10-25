package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Initialize the slow pointer
	k := 0

	// Iterate over the array with the fast pointer
	for i := 1; i < len(nums); i++ {
		// If the current element is different from the last unique element
		if nums[i] != nums[k] {
			// Move the slow pointer and update the element
			k++
			nums[k] = nums[i]
		}
	}

	// Return the number of unique elements
	return k + 1
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	k := removeDuplicates(nums)
	fmt.Println(k, nums[:k]) // Output: 5, [0 1 2 3 4]
}
