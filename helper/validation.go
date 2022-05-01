package helper

import (
	"fmt"
	"strings"
)

func ValidateUserInput(
	firstName string, lastName string, email string, userTickets uint8, remainingTickets uint8,
) []string {
	var errorMessages []string

	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isEmailValid := strings.Contains(email, "@")
	isTicketsCountValid := userTickets > 0 && userTickets <= remainingTickets

	if !isTicketsCountValid {
		var message = fmt.Sprintf(
			"We have only %v tickets, so you can`t book %v tickets",
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

func PrintErrors(errors []string) {
	for _, message := range errors {
		fmt.Printf("Validation error: %v. \n", message)
	}
}
