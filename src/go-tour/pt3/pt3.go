package main

import (
	"fmt"
	//	"golang.org/x/tour/pic"
)

func pointers() {
	i, j := 42, 2701
	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(*p)
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p1 = &Vertex{1, 2} // has type *Vertex
)

func structs() {
	fmt.Println(Vertex{1, 2})

	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v)

	p := &v
	p.X = 1e9
	fmt.Println(v)

	fmt.Println(v1, v2, v3, p1)
}

func arrays() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
}

func slices() {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(s)

	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d] == %d\n", i, s[i])
	}

	fmt.Println("s[1:4] ==", s[1:4])
	fmt.Println("s[:3] ==", s[:3])
	fmt.Println("s[4:] ==", s[4:])

	a := make([]int, 5) // len=5 cap=5
	printSlice("a", a)
	b := make([]int, 0, 5) // len=0 cap=5
	printSlice("b", b)
	c := b[:2] // len=2 cap=5
	printSlice("c", c)
	d := c[2:5] // len=3 cap=3
	printSlice("d", d)

	var z []int
	printSlice("z", z)
	if z == nil {
		fmt.Println("nil!")
	}

	// append works on nil slices
	z = append(z, 0)
	printSlice("z", z)
	z = append(z, 1)
	printSlice("z", z)
	z = append(z, 2, 3, 4) // has a len=5 cap=6 from resizing I think
	printSlice("z", z)
}
func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func ranges() {
	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	pow = make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i)
	}

	for _, v := range pow {
		fmt.Printf("%d\n", v)
	}

	for i := range pow {
		fmt.Printf("%d\n", pow[i])
	}
}

func Pic(dx, dy int) [][]uint8 {
	ret := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		ret[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			ret[y][x] = uint8(x ^ y)
		}
	}
	return ret
}

type Coordinate struct {
	Lat, Long float64
}

func maps() {
	var m = make(map[string]Coordinate)
	m["Bell Labs"] = Coordinate{
		40.68433, -74.39967, // why is this last comma necessary??
	}
	fmt.Println(m["Bell Labs"])
}

func main() {
	pointers()
	structs()
	arrays()
	slices()
	ranges()
	//pic.Show(Pic) // Note: This will print a base64 encoded image.
	maps()
}
