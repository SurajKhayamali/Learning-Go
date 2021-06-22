package models

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

func GetUsers() []*User {
	return users
}

func AddUser(u User) (User, error) {
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

func Demo() {
	var u User
	u.ID = 1
	u.FirstName = "Arthur"
	u.LastName = "Gunn"
	fmt.Println(u)

	u2 := User{
		ID:        1,
		FirstName: "Arthur",
		LastName:  "Gunn",
	}
	fmt.Println(u2)
}
