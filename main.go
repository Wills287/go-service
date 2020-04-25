package main

import (
	"net/http"

	"github.com/wills287/go-service/controller"
)

func main() {
	controller.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
