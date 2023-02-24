package main

import (
	"booking-app/helper"
	"fmt"
	"time"
)

var conferenceName = "Go Conference"

const conferenceTickets = 50

var remainingTickets uint = 50
var bookings = make([]userData, 0)

type userData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()

		isValidname, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidname && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)

			go sendTicket(userTickets,firstName,lastName,email)

			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			noTickets := remainingTickets == 0
			if noTickets {
				fmt.Println("conference is booked out")
				break
			}
		} else {
			if !isValidname {
				fmt.Println("first name or last name you entered in too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't have @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered in invalid")
			}

			fmt.Println("Your input data is invalid.")

		}

	
	}

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)

	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)

	fmt.Println("Get your tickets here to attend")

}
func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {

		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func getUserInput() (string, string, string, uint) {

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name-")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name-")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email-")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets you want to buy-")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}
func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	//CREATING A MAP for user
	var userData = userData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	// bookings = firstName+" "+lastName
	bookings = append(bookings, userData) //slice

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v.\n", firstName, lastName, userTickets, email)

	fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(50*time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("###########")
	fmt.Printf("Sending ticket %v to email address %v\n", ticket, email)
	fmt.Println("###########")
}
