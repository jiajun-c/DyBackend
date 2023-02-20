package main

import (
	"context"
	"tiktok/cmd/comment/pack"
	"tiktok/cmd/comment/service"
	"tiktok/internal/errno"
	commentpart "tiktok/kitex_gen/commentpart"
	"unicode/utf8"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// CommentAction implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentAction(ctx context.Context, req *commentpart.CommentActionRequest) (resp *commentpart.CommentActionResponse, err error) {
	resp = new(commentpart.CommentActionResponse)

	if req.VideoId == 0 || utf8.RuneCountInString(req.CommentText) > 20 {
		resp.BaseResp = pack.BuildCommentBaseResp(errno.ParamErr)
		return resp, nil
	}

	comment, err := service.NewDoCmtActionService(ctx).DoCmtAction(req)
	if err != nil {
		resp.BaseResp = pack.BuildCommentBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildCommentBaseResp(errno.Success)
	resp.Comment = comment
	return resp, nil
}

// CommentList implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentList(ctx context.Context, req *commentpart.CommentListRequest) (resp *commentpart.CommentListResponse, err error) {
	resp = new(commentpart.CommentListResponse)

	if req.VideoId == 0 {
		resp.BaseResp = pack.BuildCommentBaseResp(errno.ParamErr)
		return resp, nil
	}

	commentList, err := service.NewCommentListService(ctx).CommentList(req)
	if err != nil {
		resp.BaseResp = pack.BuildCommentBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildCommentBaseResp(errno.Success)
	resp.CommentList = commentList
	return resp, nil
}
