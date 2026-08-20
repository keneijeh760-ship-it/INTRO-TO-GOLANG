package main

import (
	"strings"
)

func ValidInput(username string, Email string, userTickets int) (bool, bool, bool) {
	isValidUserName := len(username) >= 2
	isValidUserTickets := userTickets > 0
	isValidEmail := strings.Contains(Email, "@")
	return isValidUserName, isValidUserTickets, isValidEmail
}
