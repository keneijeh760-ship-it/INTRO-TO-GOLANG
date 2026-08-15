package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50
	fmt.Printf("Welcome to the %v Booking application\n", conferenceName)
	fmt.Println("Get your tickets here to attend!")
	fmt.Printf("There are %v tickets in total and %v tickets still available for purchase.\n", conferenceTickets, remainingTickets)

	var username string
	var userTickets int = 2

	fmt.Println("What is your name? ")
	fmt.Scan(&username)

	fmt.Printf("Hello %v, you have %v tickets!", username, userTickets)
}
