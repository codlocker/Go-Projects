package main

import (
	"booking-app/helper"
	"booking-app/structs"
	"fmt"
)

var conferenceName = "Go Conference"

const conferenceTickets int = 50

var remainingTickets uint = 50

func main() {
	// Greet Users.
	helper.GreetUsers(conferenceName, conferenceTickets)

	// Map to store bookings
	bookings := make([]structs.UserData, 0)
	// Array to store firstNames
	firstNames := []string{}
	for {
		checkTicketState := remainingTickets == 0

		// Validate if tickets are enough.
		if checkTicketState {
			fmt.Printf("%s is booked out. Now buzz off\n", conferenceName)
			break
		}

		// Get User Input
		firstName, lastName, email, userTickets := helper.GetuserInput()

		// Validate User Input
		isNameValid, isEmailValid, isTicketValid := helper.ValidateInput(
			firstName,
			lastName,
			email,
			userTickets,
			remainingTickets)

		// Print validation failures
		if !isNameValid {
			fmt.Printf("Names %v %v are not valid\n", firstName, lastName)
			continue
		}

		if !isEmailValid {
			fmt.Printf("Email %v is not valid\n", email)
			continue
		}

		if !isTicketValid {
			fmt.Printf(
				"Tickets required(%v) need to be less than remaining tickets(%v)\n",
				userTickets,
				remainingTickets)
		}

		// print Valid result
		fmt.Printf(
			"User %v %v booked %v tickets. Confirmation email will be sent to %v\n",
			firstName,
			lastName,
			userTickets,
			email)

		remainingTickets -= userTickets

		booking := structs.UserData{
			FirstName:   firstName,
			LastName:    lastName,
			Email:       email,
			UserTickets: userTickets}

		bookings = append(bookings, booking)

		for _, booking := range bookings {
			firstNames = append(firstNames, booking.FirstName)
		}

		fmt.Printf("Remaining Tickets: %v\n", remainingTickets)
	}

	fmt.Println(bookings)
	fmt.Println(firstNames)
}
