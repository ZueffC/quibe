package app

import (
	"gorm.io/gorm"
)

func AddNewMessage(msg *MessageForm, topic int, db *gorm.DB) error {
	b_id := GetBoardByName(msg.Board, db)

	message := Messages{
		Author:  msg.Name,
		Text:    msg.Text,
		BoardId: b_id.ID,
		Content: msg.File,
	}

	if msg.IsReply == true && topic != 0 {
		message.TopicId = uint(topic)
	}

	db.Create(&message)

	return nil
}
