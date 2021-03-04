package main

import (
	"net/http"

	"github.com/pluralsight/go_getting_started/controllers"
)

func main() {
	controllers.RegisterController()
	http.ListenAndServe(":3000", nil)
}
