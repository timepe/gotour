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

func main() {
	mySwitch()
}
