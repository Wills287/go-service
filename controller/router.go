package controller

import "net/http"

/*
RegisterControllers initialises controllers.
*/
func RegisterControllers() {
	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}
