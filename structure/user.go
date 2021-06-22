package structure

import "fmt"

func User() {
	type user struct {
		ID        int
		FirstName string
		LastName  string
	}

	var u user
	u.ID = 1
	u.FirstName = "Arthur"
	u.LastName = "Gunn"
	fmt.Println(u)

	u2 := user{
		ID:        1,
		FirstName: "Arthur",
		LastName:  "Gunn",
	}
	fmt.Println(u2)
}
