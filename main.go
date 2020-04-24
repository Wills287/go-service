package main

import (
	"fmt"

	"github.com/wills287/go-service/model"
)

func main() {
	user := model.User{
		ID:        2,
		FirstName: "Person",
		LastName:  "Here",
	}
	fmt.Println(user)
}
