package main

import (
	"fmt"
)

func main() {

	var nemo = []string{"nemo"}

	findNemo(nemo)
}

func findNemo(array []string) {
	for _, s := range array {
		if s == "nemo" {
			fmt.Println("Found NEMO!")
		}
	}
}
