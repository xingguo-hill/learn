package x

import (
	"strings"
	"testing"
	"time"

	stan "github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
)

func getNatStremingServices() ([]string, string) {
	//服务列表，集群唯一标志
	return []string{"nats://192.168.56.29:4222", "nats://192.168.56.30:4222"}, "eventbus"
}

/**
 * @description: Regular 模式
 * 订阅者的生命周期只在一个连接内，nats-streaming server 不会保存消息
 */
func TestBaseSubscribe(t *testing.T) {
	servers, clusterID := getNatStremingServices()
	//客户端唯一标志,特殊符号只支持 “-” 或 “_”
	clientID := "subscribe"

	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL(strings.Join(servers, ",")))
	if err != nil {
		logrus.Fatalln(err)
	}
	defer sc.Close()
	// 异步订阅
	chMsg := make(chan *stan.Msg)
	sub, err := sc.Subscribe("foo", func(m *stan.Msg) {
		chMsg <- m
	})
	if err != nil {
		logrus.Fatalln(err)
	}
	defer sub.Close()

	// listen message
	for {
		select {
		case msg := <-chMsg:
			logrus.Infof("Received a message: %s\n", string(msg.Data))
		}
	}
}

/**
 * @description: Durable 模式
 * 订阅者会创建一个持久订阅，除非显示的取消订阅，即使订阅者关闭连接，订阅依然生效，订阅者会从上次消费到的位置继续消费
 */
func TestDurableSubscribe(t *testing.T) {
	servers, clusterID := getNatStremingServices()
	//客户端唯一标志,特殊符号只支持 “-” 或 “_”
	clientID := "subscriber"
	DurableName := "durable"
	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL(strings.Join(servers, ",")))
	if err != nil {
		logrus.Fatalln(err)
	}
	defer sc.Close()
	// 异步订阅
	chMsg := make(chan *stan.Msg)
	sub, err := sc.Subscribe("foo", func(m *stan.Msg) {
		chMsg <- m
	}, stan.DurableName(DurableName), stan.SetManualAckMode())
	if err != nil {
		logrus.Fatalln(err)
	}
	defer sub.Close()

	// listen message
	for {
		select {
		case msg := <-chMsg:
			logrus.Infof("Received a message: %s", string(msg.Data))
			msg.Ack()
		}
	}
}

/**
 * @description: quque durable 模式
 * 订阅者会创建一个持久订阅，发布者向 channel 发布一条消息后，同一个 queue 中的多个订阅者只有一个能消费到这条消息,消费顺序不一定是有序的，若保障有序续费，只能有一个消费者，并且需回退至durable模式
 */
func TestQueueDurableSubscribe(t *testing.T) {
	servers, clusterID := getNatStremingServices()
	//客户端唯一标志,特殊符号只支持 “-” 或 “_”
	clientID := "subscriber"
	durableName := "durable"
	groupName := "group"
	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL(strings.Join(servers, ",")))
	if err != nil {
		logrus.Fatalln(err)
	}
	defer sc.Close()
	// 异步订阅
	chMsg := make(chan *stan.Msg)
	sub, err := sc.QueueSubscribe("foo", groupName, func(m *stan.Msg) {
		chMsg <- m
	}, stan.DurableName(durableName), stan.SetManualAckMode())
	if err != nil {
		logrus.Fatalln(err)
	}
	defer sub.Close()

	// listen message
	for {
		select {
		case msg := <-chMsg:
			logrus.Infof("Received a message: %s", string(msg.Data))
			msg.Ack()
		}
	}
}

func TestPublish(t *testing.T) {
	servers, clusterID := getNatStremingServices()
	//客户端唯一标志,特殊符号只支持 “-” 或 “_”
	clientID := "publisher"
	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL(strings.Join(servers, ",")))
	if err != nil {
		logrus.Fatalln(err)
	}
	defer sc.Close()
	// 同步带阻塞发布

	err = sc.Publish("foo", []byte(time.Now().Format("2006-01-02 15:04:05"))) // does not return until an ack has been received from NATS Streaming
	if err != nil {
		logrus.Fatalln(err)
	}
}
