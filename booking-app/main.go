package main

import (
	"fmt"
	"sync"
	"time"
)

var bookings = make([]UserT, 0) // it is a list of maps
var remainingTickets uint = 50
var conferenceName string = "Go Conference"

const conferenceTickets int = 50

type UserT struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var waitGroup = sync.WaitGroup{}

func main() {

	greetUser()

	firstName, lastName, email, userTickets := getUser()
	isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		// book ticket in system
		bookTicket(firstName, lastName, email, userTickets)
		waitGroup.Add(1)
		go sendTicket(firstName, lastName, email, userTickets) // it is own thread

		// print only first names
		var firstNames []string = getFirstNames()

		fmt.Printf("First names %v\n", firstNames)

		// exit application if no tickets are left
		if remainingTickets == 0 {
			// end program
			fmt.Println("Our conference is booked out. Come back next year.")
			// break
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
		// continue
	}
	waitGroup.Wait()
}

func greetUser() {
	fmt.Printf("Welcome to %s booking application.\nWe have total of %v tickets and %v are still available.\n",
		conferenceName, conferenceTickets, remainingTickets)
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
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

func bookTicket(firstName string, lastName string, email string, userTickets uint) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserT{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("List of Bookings is %v\n type is %T", bookings, bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive name confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n bookings\n", remainingTickets, conferenceName)
}

func sendTicket(firstName string, lastName string, email string, userTickets uint) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("^^^^^^^^^^^^^^")
	fmt.Printf("Sending ticket\n %v \nto email address %v\n", ticket, email)
	fmt.Println("##############")
	waitGroup.Done()
}
