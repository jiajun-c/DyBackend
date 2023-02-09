package main

import (
	"context"
	"tiktok/cmd/chat/service"
	"tiktok/internal/app"
	"tiktok/internal/errno"
	chatpart "tiktok/kitex_gen/chatpart"
)

// ChatServiceImpl implements the last service interface defined in the IDL.
type ChatServiceImpl struct{}

// GetChatHistory implements the ChatServiceImpl interface.
func (s *ChatServiceImpl) GetChatHistory(ctx context.Context, req *chatpart.DouyinMessageChatRequest) (resp *chatpart.DouyinMessageChatResponse, err error) {
	resp = new(chatpart.DouyinMessageChatResponse)
	if req.FromUserId <= 0 || req.ToUserId <= 0 {
		resp.BaseResp = app.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}
	msgs, err := service.NewQueryChatHistoryService(ctx).QueryChatHistory(req)
	if err != nil {
		resp.BaseResp = app.BuildBaseResp(err)
		return resp, nil
	}
	resp.MessageList = msgs
	resp.BaseResp = app.BuildBaseResp(errno.Success)
	return resp, nil
}

// DoMessageAction implements the ChatServiceImpl interface.
func (s *ChatServiceImpl) DoMessageAction(ctx context.Context, req *chatpart.DouyinMessageActionRequest) (resp *chatpart.DouyinMessageActionResponse, err error) {
	resp = new(chatpart.DouyinMessageActionResponse)
	if req.ActionType != 0 || req.FromUserId <= 0 || req.ToUserId <= 0 {
		resp.BaseResp = app.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}
	err = service.NewDoMsgActionService(ctx).SendMsg(req)
	if err != nil {
		resp.BaseResp = app.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = app.BuildBaseResp(errno.Success)
	return resp, nil
}
