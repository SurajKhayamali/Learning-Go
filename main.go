package main

import (
	"fmt"

	"github.com/SurajKhayamali/learning-go/models"
)

func main() {
	u := models.User{
		ID:        1,
		FirstName: "Arthur",
		LastName:  "Gunn",
	}

	fmt.Println(u)

	models.Demo()
}
