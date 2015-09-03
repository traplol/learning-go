package main

import (
	"fmt"
	"math"
)

func loops() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum = 1
	for sum < 1000 { // wut? this is a while loop.
		sum += sum
	}
	fmt.Println(sum)

	//for {
	//	// This is an infinite loop.
	//}
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	// This is kind of strange, the first statement is local to
	// the scope of the if/else.
	if v := math.Pow(x, n); v < lim {
		return v
		// odd that 'else' MUST be on the same line as the '}'
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func Abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func Sqrt(x float64) float64 {
	switch {
	case math.IsInf(x, 1) || math.IsNaN(x) || x == 0:
		return x
	case x < 0:
		return math.NaN()
	}

	z := float64(1)
	var oldZ float64
	const epsilon = 0.00000001

	for delta := 1.0; delta > epsilon; {
		oldZ = z
		z = z - (z*z-x)/(2*z)
		delta = Abs(z - oldZ)
	}
	return z
}

//func ifTest() int {
//	if i := 42; i != 42 {
//		return i
//	}
//	return i // compile error: undefined: i
//}

func testSqrt(x float64) {
	var m, o = Sqrt(x), math.Sqrt(x)
	var d = m - o
	fmt.Println("Sqrt(", x, ") =", m, "| math.Sqrt(", x, ") =", o, "-- delta:", d)
}

func main() {
	loops()
	fmt.Println(sqrt(2), sqrt(-4), sqrt(-1))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
	testSqrt(1)
	testSqrt(2)
	testSqrt(3)
	testSqrt(10)
	testSqrt(-1)
	testSqrt(0)
	testSqrt(math.Inf(1))
	testSqrt(math.NaN())
}
