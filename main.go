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
	var Email string
	var city string
	isValidUserName := len(username) >= 2
	isValidUserTickets := userTickets > 0
	isValidEmail := strings.Contains(Email, "@")

	for remainingTickets > 0 && len(bookings) < 50 {

		fmt.Println("What is your name? ")

		fmt.Scan(&username)

		if !isValidUserName {
			fmt.Printf("%v is not a valid name broo", username)

		} else {
			continue
		}
		fmt.Println("Please enter your email: ")
		fmt.Scan(&Email)

		if !isValidEmail {
			fmt.Printf("%v is not a valid email address", Email)
		} else {
			continue
		}
		fmt.Println("How many tickets do you wanna get? ")
		fmt.Scan(&userTickets)

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

		if userTickets > remainingTickets {
			fmt.Printf("You can't do that nigga, there only %v tickets left! \n Basically you cannot buy %v tickets\n", remainingTickets, userTickets)
			break
		} else if !isValidUserTickets {
			fmt.Printf("You cannot do that bruh, are you an idiot? \n how do you wanna order %v tickets", userTickets)
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

	}

}

func greetingStatement(confName string, ticketNumber int, remainingTcikets int) {

	fmt.Println("Welcome to %v, we have about  %v available with v tickets remaining, just giv us your email and where you book atand you would be ready to go")

}
