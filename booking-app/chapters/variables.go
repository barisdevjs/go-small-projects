package main

import "fmt"

func variables() {
	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50 // var

	// fmt.Println("Welcome to the", strings.Split(conferenceName, ""), "booking application")
	// fmt.Println("We have a total of", conferenceTickets, " tickets and", remainingTickets, "are still available")

	fmt.Printf("conferenceTickets is %T\n remainingTickets is %T\n conferenceName is %T\n",
		conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n",
		conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets to attend here")

	var userName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your user name")
	fmt.Scan(&userName)
	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email")
	fmt.Scan(&email)
	fmt.Println("Enter how many tickets you want")
	fmt.Scan(&userTickets)

	remainingTickets -= userTickets

	fmt.Printf("Name %v\n LastName %v\n Email %v\n", userName, lastName, email)
	fmt.Printf("User %v has %v tickets and %v are still available\n", userName, userTickets, remainingTickets)

	// fmt.Println(&remainingTickets)
	// Prints the memory location of var conferenceName
	// calling another function outside of this
	// main1()
}

// func main1() {
// 	var name string = "John"
// 	var age int = 30
// 	var height float64 = 1.8
// 	var isStudent bool = true

// 	fmt.Printf("Variable name is of type %T\n", name)
// 	fmt.Printf("Variable age is of type %T\n", age)
// 	fmt.Printf("Variable height is of type %T\n", height)
// 	fmt.Printf("Variable isStudent is of type %T\n", isStudent)
// }
