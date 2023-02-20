package rabbitmq

import (
	"fmt"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/streadway/amqp"
)

const MQURL = ""

type RabbitMQ struct {
	conn       *amqp.Connection
	Channel    *amqp.Channel
	QueueName  string
	Exchange   string
	RoutingKey string
	mqurl      string
}

var Rmq *RabbitMQ

// 初始化RabbitMQ连接
func InitRabbitMQ() {
	Rmq = &RabbitMQ{
		mqurl: MQURL,
	}
	dial, err := amqp.Dial(Rmq.mqurl)
	if err != nil {
		Rmq.failOnError(err, "Failed to connect to RabbitMQ")
		return
	}
	Rmq.conn = dial
}

// 创建RabbitMQ channel
func NewRabbitMq(queueName, exchange, routingKey string) *RabbitMQ {
	rabbitMQ := RabbitMQ{
		QueueName:  queueName,
		conn:       Rmq.conn,
		Exchange:   exchange,
		RoutingKey: routingKey,
		mqurl:      MQURL,
	}
	var err error
	rabbitMQ.Channel, err = rabbitMQ.conn.Channel()
	if err != nil {
		Rmq.failOnError(err, "Failed to open a channel")
	}
	return &rabbitMQ
}

// 关闭RabbitMQ连接
func (r *RabbitMQ) ReleaseRes() {
	err := r.conn.Close()
	if err != nil {
		Rmq.failOnError(err, "Failed to close the connection")
	}
}

func (r *RabbitMQ) failOnError(err error, msg string) {
	if err != nil {
		klog.Fatalf("%s: %s", err, msg)
		panic(fmt.Sprintf("%s:%s\n", err, msg))
	}
}
