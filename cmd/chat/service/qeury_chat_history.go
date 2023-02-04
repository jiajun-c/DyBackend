package service

import (
	"context"
	"tiktok/cmd/chat/dal/db"
	"tiktok/cmd/chat/pack"
	"tiktok/kitex_gen/chatpart"
)

type QueryChatHistory struct {
	ctx context.Context
}

func NewQueryChatHistoryService(ctx context.Context) *QueryChatHistory {
	return &QueryChatHistory{ctx: ctx}
}

func (c *QueryChatHistory) QueryChatHistory(req *chatpart.DouyinMessageChatRequest) ([]*chatpart.Message, error) {
	msgs, err := db.GetChatHistory(c.ctx, req.FromUserId, req.ToUserId)
	if err != nil {
		return nil, err
	}
	res := pack.Messages(msgs)
	return res, nil
}
