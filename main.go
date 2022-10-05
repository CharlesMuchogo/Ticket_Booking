package main
import "fmt"




func main()  {

	var firstName string
	var lastName string
	var totalTickets uint = 50
	var remainingTickets uint = totalTickets
	var bookedTickets uint

	 names := []string{} 
	
	

	fmt.Println("XXXXXXXXXXXXXXX Welcome to concert ticket booking system XXXXXXXXXXXXXXX")

	for {

	

	fmt.Println(firstName, "Enter your first name: ")
	fmt.Scanln(&firstName)

	fmt.Println(firstName, "Enter your last name: ")
	fmt.Scanln(&lastName)

	

	fmt.Println("Enter the number of tickets you want to book: ")
	fmt.Scanln(&bookedTickets)

	

	if bookedTickets > totalTickets {
		fmt.Printf("Sorry. we only have %v tickets left \n", remainingTickets )
		continue;
	}

	remainingTickets = remainingTickets - bookedTickets
	names = append(names, firstName +" "+ lastName)
	fmt.Printf("Thank you %v %v, You have booked %v tickets. Total remaining tickets are %v \n", firstName, lastName, bookedTickets, remainingTickets)
	fmt.Println(names)
	}

	
	
}

