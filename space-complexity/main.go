package main

import "fmt"

func main() {

	booooo([]int{1, 2, 3, 4, 5, 6})

}

func booooo(n []int) {
	for i := 0; i < len(n); i++ {
		fmt.Println("boooooo!")
	}
}
