package main

import "fmt"

func main() {

	var repository UserRepositoryJSON
	repository.Initilized("example.json")
	repository.AddUser(User{1, "Subash", "p@gmail.com", 100})
	repository.AddUser(User{2, "Phuldel", "pdh@gmail.com", 1000})
	repository.AddUser(User{3, "Phssuldel", "pdh@gmssail.com", 1000})
	repository.RemoveUserByID(3)
	repository.Save()
	repository.GetAll()
	fmt.Println(AgeValidation(10))
	fmt.Println(EmailValidation("s@gmail.com"))
}
