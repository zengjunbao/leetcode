package main

//import (
//	"github.com/dipperin/go-ms-toolkit/qyenv"
//	"github.com/go-kratos/kratos/v2/log"
//	"newgitlab.com/xg-pro/superstorage/pkg/service/timer"
//
//	"newgitlab.com/xg-pro/mstoolkit/kratosplugins/cron"
//)
//
//type taskHandle struct {
//	log *log.Helper
//	srv *timer.Service
//}
//
//const (
//	CronSpecAt0000 = "0 0 * * *"
//	CronSpecAt0200 = "0 2 * * *"
//)
//
//
//func (t *taskHandle) GetTasks() []*cron.Task {
//	return []*cron.Task{
//		{
//			Name: "创建宝石发放记录",
//			Spec: func() string {
//				if qyenv.GetRunEnv() != "prod" {
//					corn, err := t.srv.GetCornSpec("RecordBeachReward")
//					if err == nil {
//						return corn
//					}
//				}
//				return CronSpecAt0000
//			}(),
//		},
//	}
//}
