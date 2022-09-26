package main

import (
	"fmt"
)

func main() {

	// var nemo = []string{"nemo"}

	// var everyone = []string{"dory", "bruce", "marlin", "nemo", "gill", "bloat",
	// "nigel", "squirt", "darla", "hank"}

	// var large = new Array(99).fill("nemo")

	var large = [10000000]string{"nemo"}

	findNemo(large) // 0(n) --> Linear Time
}

func findNemo(array [10000000]string) {
	for _, s := range array {
		if s == "nemo" {
			fmt.Println("Found NEMO!")
		}
	}
}
