package main

import "fmt"

func main() {
	adams := 42
	fmt.Printf(`42 is a bunary, %b\n`, adams)
	fmt.Printf(`42 is a bunary, %v\n`, adams)

	a, b, c, d, e, f := 0, 1, 2, 3, 4, 5

	fmt.Printf(`%b\t %#x\n`, a, a)
	fmt.Printf(`%b\t %#x\n`, b, b)
	fmt.Printf(`%b\t %#x\n`, c, c)
	fmt.Printf(`%b\t %#x\n`, d, d)
	fmt.Printf(`%b\t %#x\n`, e, e)
	fmt.Printf(`%b\t %#x\n`, f, f)
}
