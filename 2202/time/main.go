package main

import (
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func main() {
	now := time.Now()
	tamp := timestamppb.New(now)
	fmt.Println(now.UnixNano())
	fmt.Printf("%d%d\n", tamp.Seconds, tamp.Nanos)
}
