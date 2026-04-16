package main

import (
	"context"
	"harassment/api"
	"harassment/model"
	"log"
	"time"
)

func main() {
	log.Println("waiting other service start")
	time.Sleep(1 * time.Minute)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	info := api.GetInfo()
	path := info.Global.Path
	for _, TaskInfo := range info.Tasks {
		goalUser, err := api.NewSendGoal(api.AddGoalID(TaskInfo.QQID), api.AddPath(path), api.AddSendWay(TaskInfo.Private))
		if err != nil {
			log.Fatal(err)
		}
		doFunc := goalUser.SetMessage()
		messageInfo := make([]model.MessageInfo, 0, len(TaskInfo.Info))
		for _, i := range TaskInfo.Info {
			messageInfo = append(messageInfo, i)
		}
		api.Send(ctx, cancel, messageInfo, doFunc)
	}
	<-ctx.Done()
	log.Println("程序终止，需人工介入")
}
