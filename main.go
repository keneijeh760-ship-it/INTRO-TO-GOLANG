package main

import (
	"fmt"
	"intro-to-golang/helper"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50
	greetingStatement(conferenceName, conferenceTickets, remainingTickets)

	var username string
	var userTickets int
	var bookings = []string{}
	var Email string
	var city string

	for remainingTickets > 0 && len(bookings) < 50 {

		userInput(username, Email, userTickets, city, remainingTickets, bookings)

	}

}

func greetingStatement(confName string, ticketNumber int, remainingTcikets int) {

	fmt.Printf("Welcome to %v, we have about  %v available with %v tickets remaining, just giv us your email and where you book atand you would be ready to go", confName, ticketNumber, remainingTcikets)

}

func userInput(username string, Email string, userTickets int, city string, remainingTickets int, bookings []string) {
	fmt.Println("What is your name? ")

	fmt.Scan(&username)

	isValidUserName, isValidUserTickets, isValidEmail := helper.ValidInput(username, Email, userTickets)

	if !isValidUserName {
		fmt.Printf("%v is not a valid name broo", username)

	}
	fmt.Println("Please enter your email: ")
	fmt.Scan(&Email)

	if !isValidEmail {
		fmt.Printf("%v is not a valid email address", Email)
	}
	fmt.Println("How many tickets do you wanna get? ")
	fmt.Scan(&userTickets)

	if !isValidUserTickets {
		fmt.Printf("%v is not a valid number of tickets", userTickets)
	}

	if userTickets > remainingTickets {
		fmt.Printf("You can't do that nigga, there only %v tickets left! \n Basically you cannot buy %v tickets\n", remainingTickets, userTickets)

	} else {

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

	fmt.Println("What city are you booking it for? ")
	fmt.Scan(&city)

	switch city {
	case "New York":
		fmt.Println("You are booking for New York")
	case "Los Angeles":
		fmt.Println("You are booking for Los Angeles")
	case "Chicago":
		fmt.Println("You are booking for Chicago")
	default:
		fmt.Println("You are not booking for any of the cities we have tickets for")
	}

}

func sendTicket(userTickets uint, Email string, username string) {
	ticket := fmt.Sprintf("%v tickets for %v", userTickets, username)
	fmt.Printf("Ticket sent to %v\n at %v", ticket, Email)
}
