package main

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func (u *User) ModifyEmail(newEmail string) {
	u.Email = newEmail
}

func (u *User) ModifyAge(newAge int) {
	u.Age = newAge
}

func (u User) GetId() int {
	return u.Id
}

func (u User) GetName() string {
	return u.Name
}

func (u User) GetEmail() string {
	return u.Email
}

func (u User) GetAge() int {
	return u.Age
}
