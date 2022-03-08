package postModel

import (
	"app/config/database"
)

type Post struct {
	Id      int    `json:id`
	Message string `json:string`
}

func GetPosts() []Post {
	db := database.Connect()

	posts := []Post{}

	db.Find(&posts)

	return posts
}

func CreatePost(message string) {
	db := database.Connect()

	post := Post{}

	post.Message = message

	db.Create(&post)
}
