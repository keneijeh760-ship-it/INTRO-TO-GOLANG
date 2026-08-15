package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50
	fmt.Printf("Welcome to the %v Booking application\n", conferenceName)
	fmt.Println("Get your tickets here to attend!")
	fmt.Printf("There are %v tickets in total and %v tickets still available for purchase.\n", conferenceTickets, remainingTickets)

	var username string
	var userTickets int
	var bookings = []string{}

	for {

		fmt.Println("What is your name? ")

		fmt.Scan(&username)
		fmt.Println("How many tickets do you wanna get? ")
		fmt.Scan(&userTickets)

		remainingTickets -= userTickets

		bookings = append(bookings, username)

		fmt.Printf("Hello %v, you have %v tickets!\n", username, userTickets)
		fmt.Printf("There are now %v tickets left!\n", remainingTickets)

		var newTable = []string{}

		for _, booking := range bookings {

			var firstname = strings.Fields(booking)
			var name = firstname[0]
			newTable = append(newTable, name)

		}
	}

}
