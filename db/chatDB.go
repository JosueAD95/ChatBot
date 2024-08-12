package db

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"

	model "github.com/keiko30/chatbot/model"
)

var db *gorm.DB

// Run this function must be run before using other functions on this file
func ConnectDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("Chat.db"), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Couldn't connect to Database: %s", err))
	}
}

func MigrateChatDB() {
	db.AutoMigrate(&model.Chat{})
}

func GetQuestion(id uint) model.Chat {
	var chats model.Chat
	db.First(&chats, id)
	return chats
}

func GetQuestions() ([]model.Chat, error) {
	var chats []model.Chat
	result := db.Find(&chats)
	return chats, result.Error
}

func CreateChats() {
	chats := []*model.Chat{
		{Question: "What is the coldest country in the world?", Answer: "Chili"},
		{Question: "What goes up but never ever comes down?", Answer: "Your age"},
		{Question: "What can one catch that is not thrown?", Answer: "A cold"},
		{Question: "How do you make the number one disappear?", Answer: "Add a ‘G’ and it’s gone!"},
		{Question: "What two keys can’t open any door?", Answer: "A monkey and a donkey"},
		{Question: "What has a head, a tail, but does not have a body?", Answer: "A coin"},
		{Question: "What gets sharper the more you use it?", Answer: "Your brain"},
		{Question: "What two words, when combined, hold the most letters?", Answer: "Post Office"},
		{Question: "What has 4 wheels and flies?", Answer: "A garbage truck"},
		{Question: "Which room has no walls?", Answer: "A mushroom"},
	}

	if result := db.Create(chats); result.Error != nil {
		panic("Error creating the users")
	}
}
