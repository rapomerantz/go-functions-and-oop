package main

import (
	"fmt"
)

type Cube struct {
	depth  float64
	width  float64
	height float64
}

func (c *Cube) volumn() float64 {
	return c.depth * c.height * c.width
}

func main() {
	fmt.Println("in main")

	c := Cube{depth: 4, width: 4, height: 4}

	fmt.Print(c)

	fmt.Println("volume of cube is ", c.volumn())
}
