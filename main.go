package main

import "fmt"

func main() {
	b := 0.0
	a := 0.0
	fmt.Scan(&b)
	fmt.Scan(&a)
	fmt.Printf("%f\n", (b*a)/2)
}
