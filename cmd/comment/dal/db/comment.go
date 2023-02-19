package db

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        int64     `json:"id"`
	UserId    int64     `json:"user_id"`
	VideoId   int64     `json:"video_id"`
	Content   string    `json:"comment_text"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

func (Comment) TableName() string {
	return "comment"
}

// 创建评论
func CreateComment(ctx context.Context, comment *Comment) error {
	DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 1. 新增评论数据
		err := tx.Create(comment).Error
		if err != nil {
			return err
		}
		// 2. 改变 video 表中的 comment_count
		err = tx.Table("video").Where("id = ?", comment.VideoId).Update("comment_count", gorm.Expr("comment_count + ?", 1)).Error
		if err != nil {
			return err
		}
		return nil
	})
	return nil
}

// 根据id删除评论，返回该评论
func DeleteComment(ctx context.Context, commentId int64) (*Comment, error) {
	var commentRaw *Comment
	DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Where("id = ?", commentId).First(commentRaw).Error
		if err == gorm.ErrRecordNotFound {
			return err
		}
		if err != nil {
			return err
		}
		err = tx.Where("id = ?", commentId).Delete(&Comment{}).Error
		if err != nil {
			return err
		}
		err = tx.Table("video").Where("id = ?", commentRaw.VideoId).Update("comment_count", gorm.Expr("comment_count - ?", 1)).Error
		if err != nil {
			return err
		}
		return nil
	})
	return commentRaw, nil
}

// 根据视频id返回一组评论，按更新时间倒序。
func QueryCommentByVideoId(ctx context.Context, videoId int64) ([]*Comment, error) {
	var comments []*Comment
	err := DB.WithContext(ctx).Order("updated_at desc").Where("video_id = ?", videoId).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}
