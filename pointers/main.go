package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         //point to i
	fmt.Println(*p) //read i thour the pointer
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}
