package routes

import (
	"go-sample/controller/pingController"
	"go-sample/controller/postsController"
	"net/http"
)

func Route() {
	http.HandleFunc("/ping", pingController.Ping)
	http.HandleFunc("/posts", postsController.Posts)
}
