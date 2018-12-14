//+build ignore

package main

import (
	"fmt"
	"runtime"
)

func NewtonSqrt(x int) float64 {
	z := float64(x)
	var y float64
	for i := 0; i < 10; i++ {
		y = z*z - float64(x)
		z = z - y/(2*z)
	}
	return z
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

func main() {
	mySwitch()

	sliceTest()

	rangeTest()

	mapTest()

	adderTest()
}
