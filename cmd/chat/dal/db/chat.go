package db

import (
	"context"
	"time"
)

type Message struct {
	ID         int64     `json:"id"`
	FromUserID int64     `json:"from_user_id"`
	ToUserID   int64     `json:"to_user_id"`
	Content    string    `json:"content"`
	Status     int       `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
}

func (m *Message) TableName() string {
	return "message"
}

func GetChatHistory(ctx context.Context, fromUserId, toUserId int64) ([]*Message, error) {
	var res []*Message
	if err := DB.WithContext(ctx).Where("from_user_id=? and to_user_id=?", fromUserId, toUserId).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

func SendMessage(ctx context.Context, msg *Message) error {
	if err := DB.WithContext(ctx).Create(msg).Error; err != nil {
		return err
	}
	return nil
}
