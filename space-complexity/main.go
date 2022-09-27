package main

import "fmt"

func main() {

	booooo([]int{1, 2, 3, 4, 5, 6})
	fmt.Println(arrayOfHiNTimes(6))

}

func booooo(n []int) {
	for i := 0; i < len(n); i++ {
		fmt.Println("boooooo!")
	}
}

func arrayOfHiNTimes(n int) []string {
	var hiArray = []string{}
	for i := 0; i < n; i++ {
		hiArray = append(hiArray, "hi")
	}
	return hiArray
}
