package main

import (
	"errors"
	"fmt"
	"math/rand"
)

type UserService struct {
	UserRepository UserRepositoryJSON
}

var id int = rand.Int()

func (service *UserService) Initilized(fileName string) {
	service.UserRepository.Initilized(fileName)
}

func (service *UserService) RegisterUser(name string, age int, email string) error {
	err := IsEmpty(name)
	if err != nil {
		return err
	}

	err = AgeValidation(age)
	if err != nil {
		return err
	}

	err = EmailValidation(email)
	if err != nil {
		return err
	}
	id += 1
	newUser := User{id, name, email, age}
	service.UserRepository.AddUser(newUser)
	return nil
}

func (service *UserService) RemoveUserById(id int) error {
	err := service.UserRepository.RemoveUserByID(id)
	if err != nil {
		return err
	}
	return nil
}

func (service UserService) GetUserById(id int) (User, error) {
	for _, v := range service.UserRepository.Users {
		if v.GetId() == id {
			return v, nil
		}
	}
	return User{}, errors.New("Index not found in Users")
}

func (service UserService) SearchByName(name string) []User {
	matchedUsers := []User{}
	for _, v := range service.UserRepository.Users {
		if v.GetName() == name {
			matchedUsers = append(matchedUsers, v)
		}
	}

	return matchedUsers
}

func (service UserService) SearchByEmail(email string) []User {
	matchedUsers := []User{}
	for _, v := range service.UserRepository.Users {
		if v.GetEmail() == email {
			matchedUsers = append(matchedUsers, v)
		}
	}
	return matchedUsers
}

func (service UserService) SearchByAge(age int, operator string) []User {
	matchedUsers := []User{}
	for _, v := range service.UserRepository.Users {
		if v.GetAge() < age && operator == "low" {
			matchedUsers = append(matchedUsers, v)
		} else if v.GetAge() == age && operator == "eq" {
			matchedUsers = append(matchedUsers, v)
		} else if v.GetAge() > age && operator == "gre" {
			matchedUsers = append(matchedUsers, v)
		}
	}
	return matchedUsers
}

func (service *UserService) ModifyEmailById(id int, email string) error {
	for i, v := range service.UserRepository.Users {
		if v.GetId() == id {
			err := EmailValidation(email)
			if err != nil {
				return err
			}
			service.UserRepository.Users[i].ModifyEmail(email)
			return nil
		}
	}
	return errors.New("Index not found in Users")
}

func (service *UserService) ShowUsers(users []User) {
	fmt.Printf("%-5s %-15s %-25s %-5s\n", "ID", "Name", "Email", "Age")
	fmt.Println("-------------------------------------------------------")
	for _, v := range users {
		fmt.Printf("%-5d %-15s %-25s %-5d\n", v.Id, v.Name, v.Email, v.Age)
	}
}
