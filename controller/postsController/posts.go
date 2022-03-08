package postsController

import (
	"app/model/postModel"
	"encoding/json"
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
	postModel.CreatePost(r.URL.Query()["message"][0])
	fmt.Fprint(w, "Create Complete")
}

func Index(w http.ResponseWriter, r *http.Request) {
	posts := postModel.GetPosts()
	encoder := json.NewEncoder(w)
	encoder.Encode(posts)
}
