package main

import (
	"encoding/json"
	"errors"
	"os"
)

type UserRepositoryJSON struct {
	Users    []User
	FileName string
}

func (repository *UserRepositoryJSON) Initilized(fileName string) {
	repository.FileName = fileName + ".json"
	repository.Users = []User{}
}

func (repository *UserRepositoryJSON) AddUser(newUser User) {
	repository.Users = append(repository.Users, newUser)
}

func (repository *UserRepositoryJSON) RemoveUserByID(id int) error {
	for i, v := range repository.Users {
		if v.GetId() == id {
			repository.Users = append(repository.Users[:i], repository.Users[i+1:]...)
			return nil
		}
	}
	return errors.New("User Not Found")

}

func (repository *UserRepositoryJSON) Save() error {
	file, err := os.Create(repository.FileName)
	if err != nil {
		return err
	}
	defer file.Close()

	userJSON, err := json.MarshalIndent(repository.Users, "", "\t")

	if err != nil {
		return err
	}

	file.WriteString(string(userJSON))
	return nil
}

func (repository *UserRepositoryJSON) Open() error {
	data, err := os.ReadFile(repository.FileName)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &repository.Users)
	if err != nil {
		return err
	}
	return nil
}

//func (repository *UserRepositoryJSON) GetAll() {
//	fmt.Printf("%-5s %-15s %-25s %-5s\n", "ID", "Name", "Email", "Age")
//	fmt.Println("-------------------------------------------------------")
//	for _, v := range repository.Users {
//		fmt.Printf("%-5d %-15s %-25s %-5d\n", v.Id, v.Name, v.Email, v.Age)
//	}
//	fmt.Println("-------------------------------------------------------")
//}

func (repository *UserRepositoryJSON) GetAll() []User {
	return repository.Users
}
