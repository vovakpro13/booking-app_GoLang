package main

import (
	"booking-app/helper"
	"fmt"
)

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

	fmt.Printf("%v tickets remaining for %v \n", remainingTickets, Conference)
	fmt.Printf("These are all first names of our visitors: %v \n", helper.GetFirstNames(bookings))
}
