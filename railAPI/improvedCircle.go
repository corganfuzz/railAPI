package main

import (
	"fmt"
	"math"
)

// Circle bro
type nCircle struct {
	x, y, r float64
}

var c nCircle

func circleArea(c *nCircle) float64 {
	return math.Pi * c.r * c.r
}

func main() {

	c := nCircle{0, 0, 5}

	fmt.Println("Circle Area:", circleArea(&c))

}
