package main

import (
	"booking-app/helper"
	"fmt"
	"time"
)

func bookTicket(
	firstName string, lastName string, email string,
	userTickets uint8,
) {
	remainingTickets = remainingTickets - userTickets

	userData := helper.UserData{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Tickets:   userTickets,
	}

	bookings = append(bookings, userData)

	wg.Add(1)
	go sendTicket(userData)

	fmt.Printf("%v \n", bookings)
	fmt.Printf(
		"Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n",
		firstName, lastName, userTickets, email,
	)

	fmt.Printf("%v tickets remaining for %v \n", remainingTickets, Conference)
	fmt.Printf("These are all first names of our visitors: %v \n", helper.GetFirstNames(bookings))
}

func sendTicket(user helper.UserData) {
	time.Sleep(10 * time.Second)

	ticket := fmt.Sprintf("%v tickets for %v %v", user.Tickets, user.FirstName, user.LastName)

	fmt.Println("###########")
	fmt.Printf("Send ticket:\n %v \nto email address %v\n", ticket, user.Email)
	fmt.Println("###########")
	wg.Done()
}
