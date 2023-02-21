package main
import "fmt"

func main(){
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application\n",conferenceName)

	fmt.Printf("We have total of %v tickets and %v are still available\n",conferenceTickets,remainingTickets)
	fmt.Println("Get your tickets here to attend")

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

	remainingTickets = remainingTickets - userTickets

	// bookings = firstName+" "+lastName
	bookings = append(bookings,firstName+" "+lastName)  //slice

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v.\n",firstName,lastName,userTickets,email)

	fmt.Printf("%v tickets remaining for %v.\n",remainingTickets,conferenceName)
	


}


