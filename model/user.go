package model

import (
	"errors"
	"fmt"
)

/*
User object.
*/
type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

/*
GetUsers returns all users.
*/
func GetUsers() []*User {
	return users
}

/*
AddUser adds a user using the next available ID.
*/
func AddUser(user User) (User, error) {
	if user.ID != 0 {
		return User{}, errors.New("New User must not have an ID assigned")
	}
	user.ID = nextID
	nextID++
	users = append(users, &user)
	return user, nil
}

/*
GetUserByID retrieves a user by referencing the ID.
*/
func GetUserByID(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found", id)
}

/*
UpdateUser updates a user by matching the user ID.
*/
func UpdateUser(user User) (User, error) {
	for i, u := range users {
		if u.ID == user.ID {
			users[i] = &user
			return user, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found", user.ID)
}

/*
RemoveUserByID removes a user referenced by the given ID.
*/
func RemoveUserByID(id int) error {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User with ID '%v' not found", id)
}
