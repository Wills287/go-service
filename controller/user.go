package controller

import "net/http"

type userController struct {
}

func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User Controller says hello."))
}
