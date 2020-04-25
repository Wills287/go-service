package model

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
	user.ID = nextID
	nextID++
	users = append(users, &user)
	return user, nil
}
