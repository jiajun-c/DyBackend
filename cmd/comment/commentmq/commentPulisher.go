package commentmq

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/streadway/amqp"
)

func CommentActionMqSend(message []byte) {
	_, err := commentMq.Channel.QueueDeclare(
		commentMq.QueueName,
		false, // durable            持久化
		false, // delete when unused 自动删除
		false, // exclusive          排他性
		false, // no-wait            阻塞
		nil,   // arguments          额外属性
	)
	if err != nil {
		klog.Info("Failed to declare a queue for comment action publisher")
		panic(err)
	}

	// 声明交换器
	err = commentMq.Channel.ExchangeDeclare(
		commentMq.Exchange, // 交换器名
		"topic",            // exchange type：一般用fanout、direct、topic
		true,               // 是否持久化
		false,              // 是否自动删除（自动删除的前提是至少有一个队列或者交换器与这和交换器绑定，之后所有与这个交换器绑定的队列或者交换器都与此解绑）
		false,              // 设置是否内置的。true表示是内置的交换器，客户端程序无法直接发送消息到这个交换器中，只能通过交换器路由到交换器这种方式
		false,              // 是否阻塞
		nil,                // 额外属性
	)
	if err != nil {
		klog.Info("Failed to declare a exchange for comment action", err)
		return
	}

	// 建立绑定关系（可以建立多个绑定关系）
	err = commentMq.Channel.QueueBind(
		commentMq.QueueName,  // 绑定的队列名称
		commentMq.RoutingKey, // bindkey 用于消息路由分发的key
		commentMq.Exchange,   // 绑定的exchange名
		false,                // 是否阻塞
		nil,                  // 额外属性
	)
	if err != nil {
		klog.Info("Failed to bind queue with exchange for comment action", err)
		return
	}

	err = commentMq.Channel.Publish(
		commentMq.Exchange,   // exchange
		commentMq.RoutingKey, // routing key
		false,                // mandatory
		false,                // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		})
	if err != nil {
		klog.Info("Failed to publish a message for comment action")
		panic(err)
	}
}
