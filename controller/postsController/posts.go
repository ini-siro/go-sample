package postsController

import (
	"fmt"
	"net/http"
)

func Posts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		Create(w, r)
		break
	case http.MethodGet:
		Index(w, r)
		break
	}
}

func Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Create Posts!!")
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Posts!!")
}
