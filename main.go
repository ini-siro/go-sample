package main

import (
	"fmt"
	"go-sample/routes"
	"net/http"
)

func main() {
	routes.Route()
	fmt.Println("[INFO] Server listening")
	http.ListenAndServe(":3000", nil)
}
