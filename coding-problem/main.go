package main

import "fmt"

func main() {
	// Given 2 arrays, create a function that let's a user know (true/false) whether these two arrays contain any commin items
	// For example:
	// const array1 = ["a", "b", "c", "x"]
	// const array2 = ["z", "y", "i"]
	// should return false
	//------------------
	// const array1 = ["a", "b", "c", "x"]
	// const array2 = ["z", "y", "x"]
	// should return true

	// 2 parameters - arrays - no size limit
	// return true or false

	var array1 = []string{"a", "b", "c", "x"}
	var array2 = []string{"z", "y", "x"}

	fmt.Println(containsCommonItem(array1, array2))
	fmt.Println(containsCommonItem2(array1, array2))
}

func containsCommonItem(arr1 []string, arr2 []string) bool {
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			if arr1[i] == arr2[j] {
				return true
			}
		}
	}
	return false
} // O(a*b)

func containsCommonItem2(arr1 []string, arr2 []string) bool {
	// loop through first array and create object where properties === items in the array
	m := make(map[string]bool)
	for i := 0; i < len(arr1); i++ {
		if !m[arr1[i]] {
			item := arr1[i]
			m[item] = true
		}
	}

	// loopo through second array and check if item in second array exists on created object
	for j := 0; j < len(arr2); j++ {
		if m[arr2[j]] {
			return true
		}
	}

	return false
} // O(a+b)
