package main

import (
    "fmt"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

var (
    firstName       string
    email           string
    lastName        string
    totalTickets    uint = 50
    remainingTickets uint = totalTickets
    bookedTickets   uint
    bookingDetails  []bookingDetail
)

type bookingDetail struct {
    clientFirstName string
    ClientLastName  string
    ClientEmail     string
    bookedTickets   uint
}

func main() {
    fmt.Println("XXXXXXXXXXXXXXX Welcome to concert ticket booking system XXXXXXXXXXXXXXX")

    // Open the SQLite database
    db, err := sql.Open("sqlite3", "tickets.db")
    if err != nil {
        panic(err)
    }
    defer db.Close()

    // Create the bookings table if it doesn't exist
    _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS bookings (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            first_name TEXT,
            last_name TEXT,
            email TEXT,
            booked_tickets INTEGER
        )
    `)
    if err != nil {
        panic(err)
    }

    for remainingTickets > 0 {
        fmt.Println("Enter your first name:")
        fmt.Scanln(&firstName)

        fmt.Println("Enter your last name:")
        fmt.Scanln(&lastName)

        fmt.Println("Enter your email address:")
        fmt.Scanln(&email)

        fmt.Println("Enter the number of tickets you want to book:")
        fmt.Scanln(&bookedTickets)

        if bookedTickets > remainingTickets {
            fmt.Printf("Sorry, we only have %v tickets left\n", remainingTickets)
            continue
        }

        // Insert booking details into the database
        _, err = db.Exec(`
            INSERT INTO bookings (first_name, last_name, email, booked_tickets)
            VALUES (?, ?, ?, ?)
        `, firstName, lastName, email, bookedTickets)
        if err != nil {
            panic(err)
        }

        remainingTickets -= bookedTickets

        fmt.Printf("\n\n\nThank you %v, You have booked %v tickets.\n"+
            "You will receive an email containing the ticket details on %v.\n"+
            "Total remaining tickets are %v\n\n\n",
            firstName, bookedTickets, email, remainingTickets)

        // Retrieve and display booking details
        rows, err := db.Query("SELECT first_name, last_name, email, booked_tickets FROM bookings")
        if err != nil {
            panic(err)
        }
        defer rows.Close()

        fmt.Println("Booking Details:")
        for rows.Next() {
            var firstName string
            var lastName string
            var email string
            var bookedTickets uint
            err := rows.Scan(&firstName, &lastName, &email, &bookedTickets)
            if err != nil {
                panic(err)
            }
            fmt.Printf("Name: %s %s, Email: %s, Booked Tickets: %d\n", firstName, lastName, email, bookedTickets)
        }
    }
}

