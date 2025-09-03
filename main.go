package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func showMenu() int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("---------------User Management System-------------------")
	fmt.Println("------------------------Menu----------------------------")
	fmt.Println("1. Register User")
	fmt.Println("2. Get User")
	fmt.Println("3. Filter User By Name")
	fmt.Println("4. Filter User By Email")
	fmt.Println("5. Filter User By Age")
	fmt.Println("6. Modify User Email")
	fmt.Println("7. Get All Users")
	fmt.Println("8. Remove User")
	fmt.Println("9. Exit")
	fmt.Println("--------------------------------------------------------")

	fmt.Printf("Choose Options: ")
	tempInput, _ := reader.ReadString('\n')
	userChoice, _ := strconv.ParseInt(strings.TrimSpace(tempInput), 10, 64)
	return int(userChoice)
}

func checkRunning(running *bool) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("For Exit: Press Q else Press C:")
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	*running = strings.ToLower(strings.TrimSpace(input)) == "c"
}

func RegisterUser(service *UserService) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter Your name: ")
	tempInput, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	name := strings.ToTitle(strings.TrimSpace(tempInput))

	fmt.Printf("Enter Your email: ")
	tempInput, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	email := strings.ToLower(strings.TrimSpace(tempInput))

	fmt.Printf("Enter Your age: ")
	tempInput, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	age, _ := strconv.ParseInt(strings.TrimSpace(tempInput), 10, 64)
	err = service.RegisterUser(name, int(age), email)
	if err != nil {
		panic(err)
	}
}

func GetUser(service UserService) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter User ID: ")
	temInput, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	id, err := strconv.ParseInt(strings.TrimSpace(temInput), 10, 64)
	if err != nil {
		panic(err)
	}
	user, err := service.GetUserById(int(id))
	if err != nil {
		panic(err)
	}
	users := []User{}
	users = append(users, user)
	service.ShowUsers(users)
}

func FilterUserByName(service UserService) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter User Name: ")
	temInput, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	name := strings.ToTitle(strings.TrimSpace(temInput))
	matchedUsers := service.SearchByName(name)
	service.ShowUsers(matchedUsers)
}

func FilterUserByEmail(service UserService) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter User Email: ")
	temInput, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	email := strings.ToLower(strings.TrimSpace(temInput))
	matchedUsers := service.SearchByEmail(email)
	service.ShowUsers(matchedUsers)
}

func FilterUserByAge(service UserService) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter User Age: ")
	temInput, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	age, err := strconv.ParseInt(strings.TrimSpace(temInput), 10, 64)

	fmt.Printf("To filter user less than age %d Enter 'low'.", age)
	fmt.Printf("To filter user equal to age %d Enter 'eq'.", age)
	fmt.Printf("To filter user greater than age %d Enter 'gre'.", age)
	fmt.Printf("Enter Operator: ")

	temInput, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	operator := strings.ToLower(strings.TrimSpace(temInput))
	matchedUsers := service.SearchByAge(int(age), operator)
	service.ShowUsers(matchedUsers)
}

func ModifyUserEmail(service *UserService) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter User ID: ")
	temInput, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	id, err := strconv.ParseInt(strings.TrimSpace(temInput), 10, 64)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Enter Your email: ")
	temInput, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	newEmail := strings.ToLower(strings.TrimSpace(temInput))
	err = service.ModifyEmailById(int(id), newEmail)
	if err != nil {
		panic(err)
	}
}

func GetAllUsers(service UserService) {
	service.UserRepository.GetAll()
}

func RemoveUser(service *UserService) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter User ID: ")
	temInput, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	id, err := strconv.ParseInt(strings.TrimSpace(temInput), 10, 64)
	if err != nil {
		panic(err)
	}
	err = service.RemoveUserById(int(id))
	if err != nil {
		panic(err)
	}
}

func main() {
	userService := UserService{}
	userService.Initilized("Database")

	err := userService.UserRepository.Open()
	if err != nil {
		panic(err)
	}

	isRunning := true
	for isRunning {
		userOption := showMenu()
		switch userOption {
		case 1:
			RegisterUser(&userService)
		case 2:
			GetUser(userService)
		case 3:
			FilterUserByName(userService)
		case 4:
			FilterUserByEmail(userService)
		case 5:
			FilterUserByAge(userService)
		case 6:
			ModifyUserEmail(&userService)
		case 7:
			GetAllUsers(userService)
		case 8:
			RemoveUser(&userService)
		case 9:
			isRunning = false
		}
	}

	fmt.Println("Saving data to ", userService.UserRepository.FileName)
	err = userService.UserRepository.Save()

	if err != nil {
		panic(err)
	}
}
