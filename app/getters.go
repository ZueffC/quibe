package app

import (
	"fmt"

	"gorm.io/gorm"
)

func GetBoardByName(name string, db *gorm.DB) Boards {
	var board Boards
	db.Where(map[string]interface{}{"index": name}).Find(&board)

	return board
}

func GetAllBoards(db *gorm.DB) []Boards {
	var AllBoards []Boards
	db.Find(&AllBoards)
	return AllBoards
}

func GetAllMessFromBoard(board_name string, db *gorm.DB) []Messages {
	var messages []Messages
	b_id := GetBoardByName(board_name, db).ID

	db.Where(map[string]interface{}{"board_id": b_id, "topic_id": 0}).Find(&messages)

	return messages
}

func GetMessByBoardAndTopic(b_id uint, t_id int, db *gorm.DB) []Messages {
	var messages []Messages
	db.Where(map[string]interface{}{"topic_id": t_id, "board_id": b_id}).Find(&messages)
	fmt.Println(messages)

	return messages
}

func GetMessById(id int, db *gorm.DB) Messages {
	var message Messages
	db.Where(map[string]interface{}{"id": id}).Find(&message)

	return message
}
