package helper

import (
	"fmt"
)

type UserData struct {
	FirstName string
	LastName  string
	Email     string
	Tickets   uint8
}

const tickets uint8 = 50

func GreetUsers(conference string, remainingTickets uint8) {
	fmt.Printf("Welcome to %v booking application\n", conference)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", tickets, remainingTickets)
	fmt.Println("Get your your tickets here to attend")
}

func GetFirstNames(bookings []UserData) []string {
	var firstNames []string

	for _, booking := range bookings {
		firstNames = append(firstNames, booking.FirstName)
	}

	return firstNames
}
