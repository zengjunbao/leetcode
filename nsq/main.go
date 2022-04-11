package main

import (
	"encoding/json"
	"fmt"
	"github.com/nsqio/go-nsq"
)

// import (
// 	"context"
// 	"fmt"
// 	"google.golang.org/grpc"
// 	v1 "newgitlab.com/xg-pro/invitation/api/invitation/v1"
// )
//
// func main() {
// 	conn, err := grpc.Dial("127.0.0.1:30264", grpc.WithInsecure())
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer conn.Close()
// 	userRpc := v1.NewInvitationClient(conn)
// 	ctx := context.Background()
// 	// _, err = userRpc.CreateFans(ctx, &v1.CreateFansReq{
// 	// 	UserId:   33,
// 	// 	ParentId: 22,
// 	// })
// 	// if err != nil {
// 	// 	fmt.Println("register fail:", err.Error())
// 	// 	return
// 	// }
// 	data, err := userRpc.GetSumFansByUserId(ctx, &v1.GetSumFansByUserIdReq{
// 		UserId:   22,
// 	})
// 	if err != nil {
// 		fmt.Println("register fail:", err.Error())
// 		return
// 	}
// 	fmt.Println("data:",data)
//
// 	data2, err := userRpc.ListFansByUserId(ctx, &v1.ListFansByUserIdReq{
// 		UserId:   22,
// 		Page:     1,
// 		PerPage:  1,
// 		Paginate: true,
// 	})
// 	if err != nil {
// 		fmt.Println("register fail:", err.Error())
// 		return
// 	}
// 	fmt.Println("data2:",data2)
//
// 	fmt.Println("success")
// }

/**
* nsq demo - 消费者
* author: JetWu
* date: 2020.05.06
 */

// 消费者
type Consumer struct{}

type CheckFans struct {
	UserId   uint `json:"user_id,omitempty"`
	ParentId uint `json:"parent_id,omitempty"`
}

// 处理接收到的消息（实现nsq.Handler）
func (*Consumer) HandleMessage(msg *nsq.Message) error {
	fmt.Println("Receive from ", msg.NSQDAddress, ": ", string(msg.Body))
	return nil
}

func main() {
	// 创建生产者，连接nsqd
	var producer *nsq.Producer
	producer, err := nsq.NewProducer("127.0.0.1:4150", nsq.NewConfig())

	// 测试连接是否成功
	err = producer.Ping()
	if err != nil {
		producer.Stop()
		producer = nil
		panic(err)
	}

	// 指定topic，产生10条消息数据
	topic := "local-member-OnCreateFans"
	for i := 0; i < 1; i++ {
		message := fmt.Sprintf(`{"user_id":11,"parent_id":22}`)
		data, _ := json.Marshal(message)
		// fmt.Println(data, err)
		// var data2 string
		// err2 := json.Unmarshal(data,&data2)
		// fmt.Println(data2, err2)
		if producer != nil && message != "" {
			err = producer.Publish(topic, data)
			if err != nil {
				fmt.Printf("Producer publish fail: %v", err)
			}
			fmt.Println("Producer publish success:", message)
		}
	}
}
