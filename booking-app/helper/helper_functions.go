package helper

import (
	"fmt"
	"strings"
)

/*
 * Greet User function
 */
func GreetUsers(confName string, confTicket int) {
	fmt.Printf("Welcome to out %v booking application\n", confName)
	fmt.Println("Get your tickets here to attend. Max tickets", confTicket)
}

/*
 * Validate Users input
 */
func ValidateInput(firstName string, lastName string, email string, tickets uint, remTickets uint) (bool, bool, bool) {
	isInputValid := len(firstName) >= 2 && len(lastName) >= 2
	isEmailValid := strings.Contains(email, "@")
	isTicketValid := tickets <= remTickets

	return isInputValid, isEmailValid, isTicketValid
}

/*
 * Get user input from multiple users.
 */
func GetuserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email")
	fmt.Scan(&email)
	fmt.Println("Enter the number of tickets to buy")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}
