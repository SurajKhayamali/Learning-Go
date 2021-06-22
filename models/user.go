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
