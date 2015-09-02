package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9 // integer division
	y = sum - x
	return
}

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

var x = 42

func main() {
	rand.Seed(42)
	// Shouldn't this throw an error at compile time?
	//fmt.Println(rand.Intn(0))
	fmt.Println(rand.Intn(1000))

	fmt.Printf("Now you have %g problems.\n", math.Nextafter(2, 3))
	fmt.Println(math.Pi)
	fmt.Println(add(123, 456))

	a, b := swap("Hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(20))

	i, j := 1, x+2
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)

	var (
		ToBe bool = false
		//MaxInt uint64     = -1 // errors for overflowing??
		MaxInt uint64     = 1<<64 - 1 // but 1<<64 - 1 doesn't...
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)
	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf("%T(0x%x)\n", MaxInt, MaxInt)
	fmt.Printf(f, z, z)

	// This is kind of odd
	fmt.Println(needInt(Small))
	//fmt.Println(needInt(Big))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

}
