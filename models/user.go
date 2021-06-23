package models

import "fmt"

// public struct
type User struct {
	// public field
	ID        int
	FirstName string
	LastName  string
	// private field, only
	// available in models package
	password string
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

func (user User) show() {
	fmt.Println(user)
}

// this function is public as
// it begins with a capital letter
func Demo() {
	var u User
	u.ID = 1
	u.FirstName = "Arthur"
	u.LastName = "Gunn"
	fmt.Println(u)

	u2 := User{
		ID:        2,
		FirstName: "Arthur",
		LastName:  "Gunn",
		password:  "sensitive",
	}
	fmt.Println(u2)
	u2.show()

	if u == u2 {
		println("Same users!")
	} else if u.FirstName == u2.FirstName {
		println("Similar users.")
	} else {
		println("Different user!")
	}
}
