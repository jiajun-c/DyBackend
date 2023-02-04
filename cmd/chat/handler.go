package main

import (
	"context"
	"tiktok/cmd/chat/service"
	"tiktok/internal/app"
	chatpart "tiktok/kitex_gen/chatpart"
)

// ChatServiceImpl implements the last service interface defined in the IDL.
type ChatServiceImpl struct{}

// GetChatHistory implements the ChatServiceImpl interface.
func (s *ChatServiceImpl) GetChatHistory(ctx context.Context, req *chatpart.DouyinMessageChatRequest) (resp *chatpart.DouyinMessageChatResponse, err error) {
	// TODO: Your code here...
	resp = new(chatpart.DouyinMessageChatResponse)
	msgs, err := service.NewQueryChatHistoryService(ctx).QueryChatHistory(req)
	if err != nil {
		resp.BaseResp = app.BuildBaseResp(err)
		return resp, nil
	}
	resp.MessageList = msgs
	return resp, nil
}

// DoMessageAction implements the ChatServiceImpl interface.
func (s *ChatServiceImpl) DoMessageAction(ctx context.Context, req *chatpart.DouyinMessageActionRequest) (resp *chatpart.DouyinMessageActionResponse, err error) {
	// TODO: Your code here...
	resp = new(chatpart.DouyinMessageActionResponse)
	err = service.NewDoMsgActionService(ctx).SendMsg(req)
	if err != nil {
		resp.BaseResp = app.BuildBaseResp(err)
		return resp, nil
	}
	return resp, nil
}
