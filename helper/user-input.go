package helper

import "fmt"

func GetUserInput() (string, string, string, uint8) {
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

	return firstName, lastName, email, userTickets
}
