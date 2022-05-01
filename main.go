package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "Go Conference"
	const tickets uint8 = 50
	var remainingTickets uint8 = 50
	var bookings []string

	fmt.Printf("name type is %T, remainingTickets type is %T\n", name, remainingTickets)

	fmt.Printf("Welcome to %v booking application\n", name)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", tickets, remainingTickets)
	fmt.Println("Get your your tickets here to attend")

	for remainingTickets > 0 {
		var firstName string
		var lastName string
		var email string
		var userTickets uint8

		fmt.Println("What is your first name? ")
		fmt.Scan(&firstName)

		fmt.Println("What is your last name? ")
		fmt.Scan(&lastName)

		fmt.Println("What is your email? ")
		fmt.Scan(&email)

		fmt.Println("How many tickets do you want to book? ")
		fmt.Scan(&userTickets)

		if remainingTickets < userTickets {
			fmt.Printf("We have only %v tickets, so you can`t book %v tickets\n", remainingTickets, userTickets)
			continue
		}

		remainingTickets = remainingTickets - userTickets

		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf(
			"Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n",
			firstName, lastName, userTickets, email,
		)

		fmt.Printf("%v tickets remaining for %v \n", remainingTickets, name)

		var firstNames []string

		for _, booking := range bookings {
			var firstName = strings.Fields(booking)[0]

			firstNames = append(firstNames, firstName)
		}

		fmt.Printf("These are all first names of our visitors: %v \n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Our conference is booked out. Come back next year.")
			break
		}
	}

}
