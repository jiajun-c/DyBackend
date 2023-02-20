package commentmq

import (
	"context"
	"encoding/json"
	"tiktok/cmd/comment/dal/db"

	"github.com/cloudwego/kitex/pkg/klog"
)

func CommentConsumer() {
	_, err := commentMq.Channel.QueueDeclare(commentMq.QueueName, true, false, false, false, nil)
	if err != nil {
		klog.Info("Failed to declare a queue for comment action consumer")
		panic(err)
	}

	// 接收消息
	msgChanel, err := commentMq.Channel.Consume(
		commentMq.QueueName, // queue
		"",                  // consumer  用来区分多个消费者
		true,                // auto-ack  是否自动应答
		false,               // exclusive 是否具有排他性
		false,               // no-local  如果设置为true，表示不能将同一个connection中发送的消息传递给这个connection中的消费者
		false,               // no-wait   消息队列是否阻塞
		nil,                 // args
	)
	if err != nil {
		klog.Info("Failed to register a consumer for comment action")
		panic(err)
	}

	for msg := range msgChanel {
		// 这里写你的处理逻辑
		// 获取到的消息是amqp.Delivery对象，从中可以获取消息信息
		commentAction(string(msg.Body))
	}
}

func commentAction(msg string) {
	var req *CommentRmqMessage
	err := json.Unmarshal([]byte(msg), &req)
	if err != nil {
		klog.Fatalf("rabbitMq commentAdd消费时序列化失败")
		return
	}
	// 发布评论
	if req.ActionType == 1 {
		commentModel := &db.Comment{
			UserId:    req.UserId,
			VideoId:   req.VideoId,
			Content:   req.Content,
			CreatedAt: req.CreateTime,
		}
		// commentModel.ID = req.CommentId
		err := db.CreateComment(context.Background(), commentModel)
		if err != nil {
			klog.Fatalf("Failed to create a comment in rabbitmq")
			panic(err)
		}
	}
	// 删除评论
	if req.ActionType == 2 {
		_, err := db.DeleteComment(context.Background(), req.CommentId)
		if err != nil {
			klog.Fatalf("Failed to delete a comment in rabbitmq")
		}
	}
}
