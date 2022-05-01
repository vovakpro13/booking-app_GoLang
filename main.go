package main

import (
	"booking-app/helper"
	"fmt"
)

const Conference = "Go Conference"

var remainingTickets uint8 = 50
var bookings = make([]helper.UserData, 0)

func main() {
	helper.GreetUsers(Conference, remainingTickets)

	for remainingTickets > 0 {
		firstName, lastName, email, userTickets := helper.GetUserInput()

		errors := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if len(errors) > 0 {
			helper.PrintErrors(errors)
			continue
		}

		bookTicket(firstName, lastName, email, userTickets)

		if remainingTickets == 0 {
			fmt.Println("Our conference is booked out. Come back next year.")
			break
		}
	}

}
