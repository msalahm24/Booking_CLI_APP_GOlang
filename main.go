package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

type UserInput struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {
	conferenceName := "Go conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var bookings []UserInput

	greatUsers(conferenceName, conferenceTickets, remainingTickets)

	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketsNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketsNumber {

		bookingTicket(&userTickets, &remainingTickets, &bookings, &firstName, &lastName, &email)
		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)
		var firstNames []string = returnFirstNames(bookings)
		fmt.Printf("The first name of the booking is %v\n", firstNames)

	} else {
		if !isValidName {
			fmt.Println("Please enter valied names more than 3 char")
		}
		if !isValidEmail {
			fmt.Println("This email not valid email")
		}
		if !isValidTicketsNumber {
			fmt.Println("Not valid ticket number")
		}
	}
	wg.Wait()
}

func greatUsers(conferenceName string, conferenceTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to our %v booking app\n", conferenceName)
	fmt.Printf("We have totaly %v tickets and %v are remaining\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func returnFirstNames(bookings []UserInput) []string {
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

	fmt.Println("please enter your first name")
	fmt.Scan(&firstName)

	fmt.Println("please enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("please enter your email")
	fmt.Scan(&email)

	fmt.Println("please enter the number of tickets you want to booking")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func bookingTicket(userTickets *uint, remainingTickets *uint, bookings *[]UserInput, firstName *string, lastName *string, email *string) {

	var userInputs = UserInput{
		firstName:       *firstName,
		lastName:        *lastName,
		email:           *email,
		numberOfTickets: *userTickets,
	}

	*bookings = append(*bookings, userInputs)

	*remainingTickets = *remainingTickets - *userTickets
	fmt.Printf("User %v booked %v tickets and you receive a comfirmation email at %v\n", *firstName, *userTickets, *email)
	fmt.Printf("%v tickets remaining\n", *remainingTickets)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#################################")
	fmt.Printf("Sending tickets to:\n %v into email address %v\n", ticket, email)
	fmt.Println("#################################")
	wg.Done()
}
