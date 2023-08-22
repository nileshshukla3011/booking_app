package main

import (
	"GOLANG/helping"
	"fmt"
	"strconv"
)

const conferenceTickets uint = 50

var conferenceName string = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]map[string]string, 0)

func main() {

	greetUsers()
	for remainingTickets > 0 && len(bookings) < 50 {
		var firstName, lastName, email, userTickets = GetUserInput()

		isValidName, isValidEmail, isValidTicketNo := helping.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNo {

			bookTicket(remainingTickets, userTickets, firstName, lastName, email, conferenceName)
			var firstNames = getFirstname()
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
func greetUsers() {
	fmt.Printf("Welcome to %v booking application", conferenceName)
	fmt.Printf("we have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println(("Get your tickets here to attend"))

}
func getFirstname() []string {
	firstNames := []string{}
	for _, booking := range bookings {

		firstNames = append(firstNames, booking["firstName"])
	}
	return firstNames

}

func bookTicket(remainingTickets uint, userTickets uint, firstName string, lastName string, email string, conferenceName string) {
	remainingTickets = remainingTickets - userTickets
	//create a map for a user
	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTicket"] = strconv.FormatUint(uint64(userTickets), 10)
	bookings = append(bookings, userData)
	fmt.Printf("list of booking is %v", bookings)

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
