package main

type UserRepositoryInterface interface {
	AddUser(newUser User)
	RemoveUserByID(id int) error
	Save() error
	Open() error
	GetAll()
}
