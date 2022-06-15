package main
//
//import (
//	"context"
//	"fmt"
//	"google.golang.org/grpc"
//	v1 "newgitlab.com/xg-pro/superstorage/api/superstorage/v1"
//)
//
//func main() {
//	// superstorage.pro-test:3000 48424
//	conn, err := grpc.Dial("10.8.0.18:48424", grpc.WithInsecure())
//	// conn, err := grpc.Dial("superstorage:3000", grpc.WithInsecure())
//	if err != nil {
//		panic(err)
//	}
//	defer conn.Close()
//
//	rpc := v1.NewSuperstorageClient(conn)
//	ctx := context.Background()
//
//	// order_no
//	var orderNo = []string{
//		"218010020211111140118cedPpo",
//	}
//
//	_, err = rpc.InterPledgeRollBack(ctx, &v1.InterPledgeRollBackRequest{OrderNo: orderNo})
//	if err != nil {
//		fmt.Println("pledge rollback fail:", err.Error())
//		return
//	}
//	fmt.Println("success")
//}
