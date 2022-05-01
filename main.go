package main

import "fmt"

func main() {
	const name = "Go Conference"
	const tickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to %v booking application\n", name)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", tickets, remainingTickets)
	fmt.Println("Get your your tickets here to attend")
}
