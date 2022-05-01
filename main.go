package main

import (
	"fmt"
	"strings"
)

var conference = "Go Conference"

const tickets uint8 = 50

var remainingTickets uint8 = 50
var bookings []string

func main() {
	greetUsers()

	for remainingTickets > 0 {
		firstName, lastName, email, userTickets := getUserInput()

		errors := validateUserInput(firstName, lastName, email, userTickets)

		if len(errors) > 0 {
			printErrors(errors)
			continue
		}

		bookTicket(firstName, lastName, email, userTickets)

		if remainingTickets == 0 {
			fmt.Println("Our conference is booked out. Come back next year.")
			break
		}
	}

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conference)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", tickets, remainingTickets)
	fmt.Println("Get your your tickets here to attend")
}

func getUserInput() (string, string, string, uint8) {
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

func getFirstNames() []string {
	var firstNames []string

	for _, booking := range bookings {
		var firstName = strings.Fields(booking)[0]
		firstNames = append(firstNames, firstName)
	}

	return firstNames
}

func validateUserInput(
	firstName string, lastName string, email string, userTickets uint8,
) []string {
	var errorMessages []string

	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isEmailValid := strings.Contains(email, "@")
	isTicketsCountValid := userTickets > 0 && userTickets <= remainingTickets

	if !isTicketsCountValid {
		var message = fmt.Sprintf(
			"We have only %v tickets, so you can`t book %v tickets\n",
			remainingTickets, userTickets,
		)

		errorMessages = append(errorMessages, message)
	}

	if !isValidName {
		errorMessages = append(errorMessages, "Your first name or last name are too short.")
	}

	if !isEmailValid {
		errorMessages = append(errorMessages, "Your email is not contain @ symbol")
	}

	return errorMessages
}

func printErrors(errors []string) {
	for _, message := range errors {
		fmt.Printf("Validation error: %v. \n", message)
	}
}

func bookTicket(
	firstName string, lastName string, email string,
	userTickets uint8,
) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf(
		"Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n",
		firstName, lastName, userTickets, email,
	)

	fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conference)
	fmt.Printf("These are all first names of our visitors: %v \n", getFirstNames())
}
