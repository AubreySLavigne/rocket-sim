package main

import (
	"fmt"
	"math"
)

func main() {
	// Earth is an oblate spheroid, but we will aproximate it as a sphere
	// for now
	earth := sphere{
		// Mass of Earth is 5.9722 * 10^24 kg
		mass: 5.9722 * math.Pow10(27),
	}
	fmt.Printf("Earth: %v", earth)
}
