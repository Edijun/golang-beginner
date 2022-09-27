package main

import "fmt"

func main() {
	fmt.Println(hasPairWithSum([]int{6, 4, 3, 2, 1, 7}, 9))
	fmt.Println(hasPairWithSum2([]int{6, 4, 3, 2, 1, 7}, 9))
}

// naive (brute force) O(n^2)
func hasPairWithSum(arr []int, sum int) bool {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == sum {
				return true
			}
		}
	}
	return false
}

// better
func hasPairWithSum2(arr []int, sum int) bool {
	mySet := make(map[int]bool)
	for i := 0; i < len(arr); i++ {
		if mySet[arr[i]] {
			return true
		}
		item := sum - arr[i]
		mySet[item] = true
	}
	return false

}
