package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50
	fmt.Printf("Welcome to the %v Booking application", conferenceName)
	fmt.Println("Get your tickets here to attend!")
	fmt.Printf("There are %c tickets in total and %v tickets still available for purchase.", conferenceTickets, remainingTickets)
}
