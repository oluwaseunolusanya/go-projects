package main

import (
	"fmt"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	fmt.Printf("Welcome to %v booking application.\n", conferenceName) //using %v,a placeholder for variable value
	fmt.Println("Welcome to", conferenceName, "booking application.")
	fmt.Println("We have total of", conferenceTickets, "seats with", remainingTickets, "still available.")
	fmt.Println("Get your tickets here to attend.")

	var firstname, lastname, email string
	var userTickets uint
	// ask user for their name
	fmt.Println("Enter your firstname.")
	fmt.Scan(&firstname)

	fmt.Println("Enter your lastname.")
	fmt.Scan(&lastname)

	fmt.Println("Enter the number of tickets required.")
	fmt.Scan(&userTickets)

	fmt.Println("Enter your email address.")
	fmt.Scan(&email)

	remainingTickets -= userTickets

	fmt.Printf("Thank you %v %v for buying %v, a confirmation will be sent to %v\n", firstname, lastname, userTickets, email)
	fmt.Printf("Only %v tickets left.", remainingTickets)

}
