package service

import (
	"context"
	"tiktok/cmd/chat/dal/db"
	"tiktok/kitex_gen/chatpart"
)

type DoMsgActionService struct {
	ctx context.Context
}

func NewDoMsgActionService(ctx context.Context) *DoMsgActionService {
	return &DoMsgActionService{ctx: ctx}
}

func (m *DoMsgActionService) SendMsg(req *chatpart.DouyinMessageActionRequest) error {
	msg := &db.Message{
		FromUserID: req.FromUserId,
		ToUserID:   req.ToUserId,
		Content:    req.Content,
	}
	return db.SendMessage(m.ctx, msg)
}
