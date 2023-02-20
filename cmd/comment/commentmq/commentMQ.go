package commentmq

import (
	"tiktok/middleware/rabbitmq"
	"time"
)

var commentMq *rabbitmq.RabbitMQ

type CommentRmqMessage struct {
	CommentId  int64
	UserId     int64
	VideoId    int64
	Content    string
	CreateTime time.Time
	ActionType int32
}

func InitCommentMq() {
	rabbitmq.InitRabbitMQ()
	commentMq = rabbitmq.NewRabbitMq("comment_queue", "comment_exchange", "comment")
}
