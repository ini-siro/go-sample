package main

import (
	"app/config/routes"
	"app/config/tools"
	"fmt"
	"net/http"
)

func main() {
	routes.Route()
	tools.AutoMigration()
	fmt.Println("[INFO] Server listening")
	http.ListenAndServe(":3000", nil)
}
