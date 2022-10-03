package main
import "fmt"




func main()  {

	var firstName string
	var lastName string
	var totalTickets uint = 50
	var remainingTickets uint
	var bookedTickets uint
	

	fmt.Println("XXXXXXXXXXXXXXX Welcome to concert ticket booking system XXXXXXXXXXXXXXX")

	fmt.Println(firstName, "Enter your first name: ")
	fmt.Scanln(&firstName)

	fmt.Println(firstName, "Enter your last name: ")
	fmt.Scanln(&lastName)

	fmt.Println("Enter the number of tickets you want to book: ")
	fmt.Scanln(&bookedTickets)

	remainingTickets = totalTickets - bookedTickets
	fmt.Printf("Thank you %v %v, You have booked %v tickets. Total remaining tickets are %v \n", firstName, lastName, bookedTickets, remainingTickets)
	
}

func userInput(variableToScan string, messageToUser string){
	fmt.Println(messageToUser)
	fmt.Scan(&variableToScan)
	
}