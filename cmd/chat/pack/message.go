package pack

import (
	"tiktok/cmd/chat/dal/db"
	"tiktok/kitex_gen/chatpart"
)

func Message(m *db.Message) *chatpart.Message {
	if m == nil {
		return nil
	}
	return &chatpart.Message{
		Id:         m.ID,
		ToUserId:   m.ToUserID,
		FromUserId: m.FromUserID,
		Content:    m.Content,
		CreateTime: m.CreatedAt.Unix(),
	}
}

func Messages(ms []*db.Message) []*chatpart.Message {
	msgs := make([]*chatpart.Message, len(ms))
	for i, m := range ms {
		if n := Message(m); n != nil {
			msgs[i] = n
		}
	}
	return msgs
}
