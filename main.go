package main

import (
	"fmt"
	"github.com/alexsantossilva/go-app/routes"
	"net/http"
)

func main() {
	routes.LoadingRoutes()
	fmt.Println("")
	fmt.Println("Running server in: localhost - Port: 8000...")
	http.ListenAndServe(":8000", nil)
}
