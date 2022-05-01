package main

import "fmt"

func main() {
	name := "Go Conference"
	const tickets uint8 = 50
	var remainingTickets uint8 = 50

	fmt.Printf("name type is %T, remainingTickets type is %T\n", name, remainingTickets)

	fmt.Printf("Welcome to %v booking application\n", name)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", tickets, remainingTickets)
	fmt.Println("Get your your tickets here to attend")

	var userName string
	var userTickets int

	userName = "Volodymyr"
	userTickets = 7

	fmt.Printf("User %v want to buy %v tickets", userName, userTickets)

}
