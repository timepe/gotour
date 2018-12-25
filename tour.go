package main

import (
	"fmt"
	"io"
	"math"
	"runtime"
	"strings"
	"time"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func NewtonSqrt(x int) (float64, error) {
	z := float64(x)
	if z < 0 {
		return -1, ErrNegativeSqrt(z)
	}
	var y float64
	for i := 0; i < 10; i++ {
		y = z*z - float64(x)
		z = z - y/(2*z)
	}
	return z, nil
}

func testNewtonSqrt() {
	fmt.Println(NewtonSqrt(-2))
}

func mySwitch() {
	fmt.Println("Go runs on")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Printf("%s.", os)
	}
}

func sliceTest() {
	var a []int
	printSlice("a", a)

	a = append(a, 0)
	printSlice("a", a)

	a = append(a, 2, 3, 4)
	printSlice("a", a)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func rangeTest() {
	ranges := []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range ranges {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	for i := range ranges {
		fmt.Println("index: ", i)
	}

	for _, v := range ranges {
		fmt.Println("values: ", v)
	}
}

func mapTest() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value: ", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value: ", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value: ", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value: ", v, "Present? ", ok)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func adderTest() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func vertexTest() {
	v := &Vertex{3, 4}
	fmt.Println(v.Abs())
}

type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func testIpAddr() {
	addrs := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for n, a := range addrs {
		fmt.Printf("%v: %v\n", n, a)
	}
}

//error handler

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v,%s", e.What, e.When)
}

func run() error {
	return &MyError{time.Now(), "it didn't work"}
}

func testRun() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

func testReader() {
	r := strings.NewReader("Hello, Reader")

	b := make([]byte, 8)

	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err= %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

type MyReader struct{}

func (r *MyReader) Read(b []byte) (int, error) {
	b[0] = 'A'
	return 1, nil
}

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum
}

func testChan() {
	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}

func testChannelFull() {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func testFibChan() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func testFib2() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 9; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fibonacci2(c, quit)
}

func main() {
	mySwitch()

	sliceTest()

	rangeTest()

	mapTest()

	adderTest()

	vertexTest()

	fmt.Println(math.Sqrt2)

	testIpAddr()

	testRun()

	testNewtonSqrt()

	testReader()

	testChan()

	testChannelFull()

	testFibChan()

	testFib2()
}
