package main

import (
	"fmt"
	"strings"
)

var bookings = []string{}
var remainingTickets uint = 50
var conferenceName string = "Go Conference"

const conferenceTickets int = 50

func functions2() {

	greetUser()

	for {

		firstName, lastName, email, userTickets := getUser2()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput2(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			// book ticket in system
			bookTicket2(firstName, lastName, email, userTickets)

			// print only first names
			var firstNames []string = getFirstNames2()

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

func greetUser() {
	fmt.Printf("Welcome to %s booking application.\nWe have total of %v tickets and %v are still available.\n",
		conferenceName, conferenceTickets, remainingTickets)
}

func getFirstNames2() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserInput2(name string, last string, email string, uTickets uint) (bool, bool, bool) {
	isValidName := len(name) >= 2 && len(last) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := uTickets > 0 && uTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func getUser2() (string, string, string, uint) {
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

func bookTicket2(firstName string, lastName string, email string, userTickets uint) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive name confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n bookings\n", remainingTickets, conferenceName)
}
