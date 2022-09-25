package main

import (
	"fmt"
	"time"
)

func main() {

	// var nemo = []string{"nemo"}

	// var everyone = []string{"dory", "bruce", "marlin", "nemo", "gill", "bloat",
	// "nigel", "squirt", "darla", "hank"}

	// var large = new Array(99).fill("nemo")

	var large = [10000000]string{"nemo"}

	findNemo(large)
}

func findNemo(array [10000000]string) {
	t0 := time.Now()
	for _, s := range array {
		if s == "nemo" {
			fmt.Println("Found NEMO!")
		}
	}
	t1 := time.Now()
	fmt.Printf("Call to find Nemo took %v (milliseconds)", t1.Sub(t0))
}
