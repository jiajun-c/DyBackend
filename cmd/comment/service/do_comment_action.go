package service

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"tiktok/cmd/comment/commentmq"
	"tiktok/cmd/comment/dal/db"
	"tiktok/cmd/comment/pack"
	userdb "tiktok/cmd/user/dal/db"
	"tiktok/kitex_gen/commentpart"

	"github.com/cloudwego/kitex/pkg/klog"
)

type DoCmtActionService struct {
	ctx context.Context
}

func NewDoCmtActionService(ctx context.Context) *DoCmtActionService {
	return &DoCmtActionService{ctx: ctx}
}

func (s *DoCmtActionService) DoCmtAction(req *commentpart.CommentActionRequest) (*commentpart.Comment, error) {
	var wg sync.WaitGroup
	wg.Add(2)
	var comment *db.Comment
	var user *userdb.User
	var commentErr, userErr error

	// 发布评论
	if req.ActionType == 1 {
		comment = &db.Comment{
			UserId:  req.UserId,
			VideoId: req.VideoId,
			Content: req.CommentText,
		}
		// 创建评论
		go func() {
			defer wg.Done()
			// err := db.CreateComment(s.ctx, comment)
			// if err != nil {
			// 	commentErr = err
			// 	return
			// }

			// 发送消息给MQ
			comment.CreatedAt = time.Now()
			commentMessage := commentmq.CommentRmqMessage{
				UserId:     comment.UserId,
				VideoId:    comment.VideoId,
				Content:    comment.Content,
				CreateTime: comment.CreatedAt,
				ActionType: req.ActionType,
			}
			msg, err := json.Marshal(commentMessage)
			if err != nil {
				klog.Fatalf("序列化添加评论请求参数失败")
				return
			}
			commentmq.CommentActionMqSend(msg)
		}()
		// 获取当前用户信息
		go func() {
			defer wg.Done()
			users, err := userdb.QueryUserByIds(s.ctx, []int64{req.UserId})
			if err != nil {
				userErr = err
				return
			}
			user = users[0]
		}()
		wg.Wait()
		if commentErr != nil {
			return nil, commentErr
		}
		if userErr != nil {
			return nil, userErr
		}
		return pack.CommentInfo(comment, user), nil
	}

	// 删除评论
	if req.ActionType == 2 {
		//删除评论
		go func() {
			defer wg.Done()
			// var err error
			// comment, err = db.DeleteComment(s.ctx, req.CommentId)
			// if err != nil {
			// 	commentErr = err
			// 	return
			// }

			// 发送消息给MQ
			commentMessage := commentmq.CommentRmqMessage{
				CommentId:  req.CommentId,
				ActionType: req.ActionType,
			}
			msg, err := json.Marshal(commentMessage)
			if err != nil {
				klog.Fatalf("序列化删除评论请求参数失败")
				return
			}
			commentmq.CommentActionMqSend([]byte(msg))
		}()
		//获取用户信息
		go func() {
			defer wg.Done()
			users, err := db.QueryUserByIds(s.ctx, []int64{req.UserId})
			if err != nil {
				userErr = err
				return
			}
			user = users[0]
		}()
		wg.Wait()
		if commentErr != nil {
			return nil, commentErr
		}
		if userErr != nil {
			return nil, userErr
		}
		return pack.CommentInfo(comment, user), nil
	}

	return nil, nil
}
