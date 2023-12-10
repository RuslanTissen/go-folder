package main

import "fmt"

func main() {
	var conferenceName = "Go conference"
	const conferenceTieckets = 50
	const remainingTieckets = 50

	fmt.Println("Welcome to", conferenceName, "booking aplication")
	fmt.Println("We have total", conferenceTieckets, "tieckets and", remainingTieckets, "are still available.")
	fmt.Println("Hello world!")
}
