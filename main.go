package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	bookings := []string{}
	for remainingTickets > 0 && len(bookings) < 50 {
		var firstName, lastName, email, userTickets = GetUserInput()

		isValidName, isValidEmail, isValidTicketNo := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNo {

			bookTicket(remainingTickets, userTickets, bookings, firstName, lastName, email, conferenceName)
			var firstNames = getFirstname(bookings)
			fmt.Printf("The First name of all our bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("our confrence is booked out come back next year")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("you have not entered valid name")
			}
			if !isValidEmail {
				fmt.Println("you have not entered valid Email")
			}
			if !isValidTicketNo {
				fmt.Println("you have not entered valid ticket number")
			}

		}

	}
}
func greetUsers(confName string, conferenceTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application", confName)
	fmt.Printf("we have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println(("Get your tickets here to attend"))

}
func getFirstname(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames

}
func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) > 2 && len(lastName) > 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNo := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNo
}
func bookTicket(remainingTickets uint, userTickets uint, bookings []string, firstName string, lastName string, email string, conferenceName string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for  booking %v tickets. Your confirmation mail will be sent on %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining %v \n", remainingTickets, conferenceName)

}
func GetUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	fmt.Println(("Enter your first name: "))
	fmt.Scan(&firstName)

	fmt.Println(("Enter your last name: "))
	fmt.Scan(&lastName)

	fmt.Println(("Enter your email: "))
	fmt.Scan(&email)

	fmt.Println(("Enter the no of ticket: "))
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets

}
