package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)

	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var bookings []string

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

		isValidname := len(firstName)>=2 && len(lastName)>=2

		isValidEmail := strings.Contains(email,"@")

		isValidTicketNumber  := userTickets>0 && userTickets<=remainingTickets


		if isValidname && isValidEmail && isValidTicketNumber {

			remainingTickets = remainingTickets - userTickets

			// bookings = firstName+" "+lastName
			bookings = append(bookings, firstName+" "+lastName) //slice

			fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v.\n", firstName, lastName, userTickets, email)

			fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)

				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("the first names of bookings are %v\n", firstNames)

			noTickets := remainingTickets == 0
			if noTickets {
				fmt.Println("conference is booked out")
				break
			}
		} else {
			if !isValidname{
				fmt.Println("first name or last name you entered in too short")
			}
			if !isValidEmail{
				fmt.Println("email address you entered doesn't have @ sign")
			}
			if !isValidTicketNumber{
				fmt.Println("number of tickets you entered in invalid")
			}

			fmt.Println("Your input data is invalid.")

		}

	}

}
