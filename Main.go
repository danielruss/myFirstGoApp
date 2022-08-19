package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func add(x int, y int) int {
	return x + y
}

// function can return multiple items
// and if multiple string are passed in your
// can leave out the type from all but the last.
func swap(x, y string) (string, string) {
	return y, x
}

// you can have named return values..
// https://go.dev/tour/basics/7
// use with care..
// x and y are returned. bare return.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	// from the go tour...
	fmt.Println("Hello, world.")
	fmt.Println("The current time is ", time.Now())
	// let play with random number... (not very random??)
	fmt.Println("Random integer ", rand.Intn(10), " pi=", math.Pi)
	fmt.Println("1+2=", add(1, 2))
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(23))
}
