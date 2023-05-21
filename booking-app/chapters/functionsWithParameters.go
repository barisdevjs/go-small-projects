package main

import (
	"fmt"
	"strings"
)

func functions1() {

	var bookings = []string{}
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var conferenceName string = "Go Conference"

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for {

		firstName, lastName, email, userTickets := getUser()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			// book ticket in system
			bookTicket(firstName, lastName, email, userTickets, remainingTickets, bookings, conferenceName)

			// print only first names
			var firstNames []string = getFirstNames(bookings)

			fmt.Printf("First names %v\n", firstNames)

			// exit application if no tickets are left
			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
			continue
		}
	}
}

func greetUsers(name string, total int, tickets uint) {
	fmt.Printf("Welcome to %s booking application.\nWe have total of %v tickets and %v are still available.\n",
		name, total, tickets)
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserInput(name string, last string, email string, uTickets uint, rTickets uint) (bool, bool, bool) {
	isValidName := len(name) >= 2 && len(last) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := uTickets > 0 && uTickets <= rTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func getUser() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// asking for user input
	fmt.Println("Enter Your First Name: ")
	fmt.Scanln(&firstName)

	fmt.Println("Enter Your Last Name: ")
	fmt.Scanln(&lastName)

	fmt.Println("Enter Your Email: ")
	fmt.Scanln(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scanln(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(firstName string, lastName string, email string, userTickets uint, remainingTickets uint, bookings []string, conferenceName string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive name confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n bookings %v\n", remainingTickets, conferenceName, bookings)
}
