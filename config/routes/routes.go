package routes

import (
	"app/controller/pingController"
	"app/controller/postsController"
	"net/http"
)

func Route() {
	http.HandleFunc("/ping", pingController.Ping)
	http.HandleFunc("/posts", postsController.Posts)
}
