package main
import "fmt"


	var firstName string
	var email string
	var lastName string
	var totalTickets uint = 50
	var remainingTickets uint = totalTickets
	var bookedTickets uint

	 var bookingDetails[] bookingDetail

type bookingDetail struct{
	 clientFirstName string 
	 ClientLastName string
	 ClientEmail string
	 bookedTickets	uint
}


func main()  {

	
	
	

	fmt.Println("XXXXXXXXXXXXXXX Welcome to concert ticket booking system XXXXXXXXXXXXXXX")

	for {
		if remainingTickets == 0 {
			break;
		}	

	fmt.Println("Enter your first name: ")
	fmt.Scanln(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scanln(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scanln(&email)

	

	fmt.Println("Enter the number of tickets you want to book: ")
	fmt.Scanln(&bookedTickets)

	

	if bookedTickets > totalTickets {
		fmt.Printf("Sorry. we only have %v tickets left \n", remainingTickets )
		continue;
	}

	
	TicketBooked := bookingDetail{clientFirstName: firstName, ClientLastName: lastName, ClientEmail: email, bookedTickets : bookedTickets}

	remainingTickets = remainingTickets - bookedTickets
	bookingDetails = append(bookingDetails, TicketBooked)
	fmt.Printf(" \n \n \nThank you %v, You have booked %v tickets. \nYou will receive an email containing the ticket details on %v. \nTotal remaining tickets are %v \n \n \n", TicketBooked.clientFirstName, TicketBooked.bookedTickets, TicketBooked.ClientEmail, remainingTickets)
	
	fmt.Printf("Booked tickets: %v \n \n", bookingDetails)
	}

	
	
}

