package main

// func main() {
// 	// 创建一个计时器
// 	timeTicker := time.NewTicker(time.Second * 2)
// 	i := 0
// 	for {
// 		if i > 5 {
// 			break
// 		}
//
// 		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
// 		i++
// 		<-timeTicker.C
//
// 	}
// 	// 清理计时器
// 	timeTicker.Stop()
// }

// func main() {
// 	c := make(chan int)
// 	fmt.Println(time.Now())
// 	go func() {
// 		time.Sleep(time.Now().Add(5 * time.Second).Sub(time.Now()))
// 		c <- 0
// 		time.Sleep(time.Now().Add(10 * time.Second).Sub(time.Now()))
// 		c <- 1
// 	}()
//
// 	for {
// 		select {
// 		case p := <-c:
// 			fmt.Printf("p=%d\n", p)
// 		case <-time.After(20 * time.Second):
// 			fmt.Println(time.Now())
// 			fmt.Printf("timeout")
// 			return
// 		}
// 	}
// }


func main() {
	// c := cron.Cron{}
	//
	// c := cron.New(
	// )
	//
	// data, err := c.AddFunc("0 * * * * *", func() {
	// 	fmt.Println(time.Now())
	// })
	// c.Run()
	//
	//
	// fmt.Println("data:",data,"err:", err)
	// for  {
	// 	select {
	//
	// 	}
	// }
}

// func main() {
// 	now := time.Now()
// 	tamp := timestamppb.New(now)
// 	fmt.Println(now.UnixNano())
// 	fmt.Printf("%d%d\n", tamp.Seconds, tamp.Nanos)
// }
