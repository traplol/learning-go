package main

import (
	"fmt"
	"golang.org/x/tour/reader"
	"image"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"strings"
	"time"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func typeMethods() {
	v := &Vertex{3, 4}
	fmt.Println(v.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f)
	fmt.Println(f.Abs())

}

type Abser interface {
	Abs() float64
}

type Reader interface {
	Read(b []byte) (n int, err error)
}

type Writer interface {
	Write(b []byte) (n int, err error)
}

type ReaderWriter interface {
	Reader
	Writer
}

type Person struct {
	Name string
	Age  uint
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func interfaces() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f
	fmt.Println(a.Abs())
	a = &v
	fmt.Println(a.Abs())

	var w Writer
	w = os.Stdout
	fmt.Fprintf(w, "hello, writer\n")

	p := Person{"Max", 24}
	fmt.Println(p)

	ip := IPAddr{255, 255, 255, 255}
	fmt.Println(ip)
}

type MyError struct {
	When time.Time
	What string
}

func (err *MyError) Error() string {
	return fmt.Sprintf("at %v: %s", err.When, err.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work...",
	}
}

type ErrSqrtOfInf float64
type ErrSqrtOfNaN float64
type ErrSqrtOfNegativeNumber float64

func (err ErrSqrtOfInf) Error() string {
	return "Square root of infinity."
}
func (err ErrSqrtOfNaN) Error() string {
	return "Square root of NaN."
}
func (err ErrSqrtOfNegativeNumber) Error() string {
	return fmt.Sprintf("Square root of a negative number(%f).", err)
}

func Abs(f float64) float64 {
	if f < 0 {
		return -f
	}
	return f
}

func Sqrt(x float64) (float64, error) {
	switch {
	case x == 0:
		return x, nil
	case math.IsInf(x, 1):
		return x, ErrSqrtOfInf(x)
	case math.IsNaN(x):
		return x, ErrSqrtOfNaN(x)
	case x < 0:
		return math.NaN(), ErrSqrtOfNegativeNumber(x)
	}

	z := float64(1)
	var oldZ float64
	const epsilon = 0.00000001

	for delta := 1.0; delta > epsilon; {
		oldZ = z
		z = z - (z*z-x)/(2*z)
		delta = Abs(z - oldZ)
	}
	return z, nil
}

func testSqrt(n float64) {
	v, err := Sqrt(n)
	fmt.Printf("Sqrt(%v) = %v, error: %v\n", n, v, err)
}

func errors() {
	if err := run(); err != nil {
		fmt.Println(err)
	}

	testSqrt(1)
	testSqrt(2)
	testSqrt(3)
	testSqrt(10)
	testSqrt(-1)
	testSqrt(0)
	testSqrt(math.Inf(1))
	testSqrt(math.NaN())
}

type MyReader struct{}

func (MyReader) Read(b []byte) (int, error) {
	i := 0
	for ; i < len(b); i++ {
		b[i] = 'A'
	}
	return i - 1, nil
}

type rot13Reader struct {
	r io.Reader
}

func rot(n int, c byte) byte {
	n %= 26
	if 'a' <= c && c <= 'z' {
		c -= 'a'
		c = byte(c + byte(n))
		c %= 26
		return c + 'a'
	}
	if 'A' <= c && c <= 'Z' {
		c -= 'A'
		c = byte(c + byte(n))
		c %= 26
		return c + 'A'
	}
	return c
}

func (r rot13Reader) Read(b []byte) (int, error) {
	s, err := r.r.Read(b)
	if err != nil {
		return s, err
	}
	if len(b) == 0 {
		return 0, nil
	}

	for i, v := range b {
		b[i] = rot(13, byte(v))
	}
	return s, err
}

func readers() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}

	reader.Validate(MyReader{})
	rot13test()

	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r13 := rot13Reader{s}
	io.Copy(os.Stdout, &r13)
	fmt.Println()
}

func rot13test() {
	for i := 'a'; i <= 'z'; i++ {
		c := rot(13, rot(13, byte(i)))
		if byte(i) != c {
			fmt.Println("rot13 tests failed!")
			return
		}
	}
	fmt.Println("rot13 tests passed!")
}

type String string

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (s *Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s%s %s", s.Greeting, s.Punct, s.Who)
}

func httpServer() {
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}

func images() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}

func main() {
	typeMethods()
	interfaces()
	errors()
	readers()
	//httpServer()
	images()
}
