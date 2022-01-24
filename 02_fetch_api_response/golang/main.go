package main

import (
	"fmt"

	"github.com/nahidhasan98/care-t/02_fetch_api_response/golang/model"
	"github.com/nahidhasan98/care-t/02_fetch_api_response/golang/service"
)

func main() {
	fmt.Println("Program is running...")

	// displaying options
	fmt.Println("\n==============================")
	fmt.Println("1. Get all user list")
	fmt.Println("2. Create a new user")
	fmt.Println("==============================")

	fmt.Printf("\nPlease select an option: ")

	// taking user choice
	var option int
	fmt.Scan(&option)

	switch option {
	case 1: // for getting user list
		userList, err := service.GetUserList()
		if err != nil { // got error
			fmt.Println("Error occurred while fetching data: ", err)
			break
		}

		// no error, got the data
		fmt.Printf("Total %v user found. User List:", len(userList))
		for i := 0; i < len(userList); i++ {
			fmt.Printf("ID: %v, Name: %v, Email: %v, Gender: %v, Status: %v\n", userList[i].ID, userList[i].Name, userList[i].Email, userList[i].Gender, userList[i].Status)
		}
	case 2: // creating a new user
		// taking new user data
		var userData model.User
		fmt.Printf("Enter Name: ")
		fmt.Scan(&userData.Name)
		fmt.Printf("Enter Email: ")
		fmt.Scan(&userData.Email)
		fmt.Printf("Enter Gender: ")
		fmt.Scan(&userData.Gender)
		fmt.Printf("Enter Status (active/inactive): ")
		fmt.Scan(&userData.Status)

		user, err := service.CreateUser(&userData)
		if err != nil { // got error
			fmt.Println("Error occurred while creating user:", err)
			break
		}

		// no error, user created and got data
		fmt.Println("\nNew user created successfully. User Details:")
		fmt.Printf("ID: %v, Name: %v, Email: %v, Gender: %v, Status: %v\n", user.ID, user.Name, user.Email, user.Gender, user.Status)
	default: // wrong option choice
		fmt.Println("Wrong selection. Please try again!")
	}

	fmt.Println("\nProgram exiting...")
}
