package tools

import (
	"app/config/database"
	"app/model/postModel"
)

func AutoMigration() {
	db := database.Connect()

	db.AutoMigrate(&postModel.Post{})

	db.Close()
}
