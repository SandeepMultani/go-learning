package main

import (
	"net/http"

	"github.com/sandeepmultani/go-learning/webservice/controllers"
)

func main() {
	controllers.RegisterControllers()

	http.ListenAndServe(":3000", nil)
}
