package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Printf("x is %T\n", x)
	fmt.Printf("y is %T\n", y)
	fmt.Printf("f is %T\n", f)
	fmt.Printf("z is %T\n", z)
	fmt.Println(f)
	fmt.Println(z)
}
