package main

import (
	"fmt"
	"strings"
)

func ifElse() { // must be always main

	var remainingTickets uint = 50
	conferenceName := "Go Conference"
	bookings := []string{}

	for remainingTickets > 0 && len(bookings) < 50 {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		if remainingTickets <= 0 {
			fmt.Println("Sorry, all tickets have been sold out.")
			break
		}

		// asking for user input
		fmt.Println("Enter Your First Name: ")
		fmt.Scanln(&firstName)

		fmt.Println("Enter Your Last Name: ")
		fmt.Scanln(&lastName)

		fmt.Println("Enter Your Email: ")
		fmt.Scanln(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scanln(&userTickets)

		if userTickets > remainingTickets {
			fmt.Printf("Sorry, you can not book more than %v tickets.\n", remainingTickets)
			break
		}

		// book ticket in system
		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n",
			firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
		fmt.Printf("These are all our bookings %v\n", bookings)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("The first names %v\n", firstNames)
	}
}
