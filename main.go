package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("we have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println(("Get your tickets here to attend"))

	bookings := []string{}
	for remainingTickets > 0 && len(bookings) < 50 {

		var firstName string
		var lastName string
		var email string
		var userTickets uint
		//ask user for their name
		fmt.Println(("Enter your first name: "))
		fmt.Scan(&firstName)

		fmt.Println(("Enter your last name: "))
		fmt.Scan(&lastName)

		fmt.Println(("Enter your email: "))
		fmt.Scan(&email)

		fmt.Println(("Enter the no of ticket: "))
		fmt.Scan(&userTickets)

		isValidName := len(firstName) > 2 && len(lastName) > 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNo := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNo {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for  booking %v tickets. Your confirmation mail will be sent on %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets are remaining ", remainingTickets)
			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("The First name of all our bookings are: %v\n", firstNames)
			if remainingTickets == 0 {
				fmt.Println("our confrence is booked out come back next year")
				break
			}

		} else {
			fmt.Println("your input data is invalid, try again")

		}

	}
}
