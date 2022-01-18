package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceTkts int = 50

var conferenceName = "Go Conference"
var remainingTkts uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName string
	lastName  string
	email     string
	noOfTkts  uint
}

var waitGrp = sync.WaitGroup{}

func main() {

	greetUsers()
	for {
		firstName, lastName, email, userTkts := getUserInput()
		isValidName, isValidEmail, isValidTktNo := helper.ValidateUserInput(firstName, lastName, email, userTkts, remainingTkts)

		if isValidName && isValidEmail && isValidTktNo {

			bookTicket(userTkts, firstName, lastName, email)

			waitGrp.Add(1)
			go sendTicket(userTkts, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTkts == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				waitGrp.Wait()
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isValidTktNo {
				fmt.Println("number of tickets you entered is invalid")
			}
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTkts, remainingTkts)
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
	var userTkts uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTkts)

	return firstName, lastName, email, userTkts
}

func bookTicket(userTkts uint, firstName string, lastName string, email string) {
	remainingTkts = remainingTkts - userTkts

	var userData = UserData{
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		noOfTkts:  userTkts,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTkts, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTkts, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("--------------")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("---------------")
	waitGrp.Done()
}
