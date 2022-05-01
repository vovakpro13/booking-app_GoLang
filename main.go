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

	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Println("What is your first name? ")
	fmt.Scan(&firstName)

	fmt.Println("What is your last name? ")
	fmt.Scan(&lastName)

	fmt.Println("What is your email? ")
	fmt.Scan(&email)

	fmt.Println("How many tickets do you want to book? ")
	fmt.Scan(&userTickets)

	fmt.Printf(
		"Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n",
		firstName, lastName, userTickets, email,
	)

}
