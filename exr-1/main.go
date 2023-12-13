package main

import "fmt"

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Printf("%T \t %T \t%T \t%T \t%T \t", i, j, c, python, java)
}

