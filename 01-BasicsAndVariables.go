// GO application must belong to a package. Everything is oragined in packages
package main

import "fmt"

// declare entry point of application
func main() {
	//var conferenceName = "Go Conference"
	//another way of defining variable
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("conferenceTicket is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	//ask user  for its info
	fmt.Printf("Enter your first name\n")
	fmt.Scan(&firstName)

	fmt.Printf("Enter your last name\n")
	fmt.Scan(&lastName)

	fmt.Printf("Enter your email\n")
	fmt.Scan(&email)

	fmt.Printf("Enter no. of tickets\n")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("User %v %v booked %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remining for %v\n", remainingTickets, conferenceTickets)
}
