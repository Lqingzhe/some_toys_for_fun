package api

import (
	"bytes"
	"context"
	"harassment/model"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/bytedance/sonic"
)

func (S *SendStruct) SetMessage() func([]string) (goalID string, returnString string, err error) {
	return func(messages []string) (goalID string, returnString string, err error) {
		messageLen := len(messages)
		goalID = S.goalID
		message := messages[rand.Intn(messageLen)]
		params := map[string]interface{}{
			"user_id": goalID,
			"message": message,
		}

		// 转为JSON
		jsonData, _ := sonic.Marshal(params)
		req, err := http.NewRequest("POST", S.path+S.sendWay, bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer 114514")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			returnString = ""
			return goalID, returnString, err
		}
		defer func() {
			closeErr := resp.Body.Close()
			if closeErr != nil && err == nil {
				err = closeErr
			}
		}()
		body, _ := io.ReadAll(resp.Body)
		returnString = string(body)
		return goalID, returnString, err
	}
}
func Send(ctx context.Context, cancel context.CancelFunc, MessageInfo []model.MessageInfo, do func([]string) (string, string, error)) {
	go func() {
		log.Println("begin")
		Len := len(MessageInfo)
		p := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
			}
			now := time.Now()
			today := time.Date(now.Year(), now.Month(), now.Day(), MessageInfo[p].Hour, MessageInfo[p].Minute, 0, 0, now.Location())
			for now.After(today) {
				p++
				if p == Len {
					p -= Len
					today = time.Date(now.Year(), now.Month(), now.Day(), MessageInfo[p].Hour, MessageInfo[p].Minute, 0, 0, now.Location()).AddDate(0, 0, 1)
					break
				}
				today = time.Date(now.Year(), now.Month(), now.Day(), MessageInfo[p].Hour, MessageInfo[p].Minute, 0, 0, now.Location())
			}
			wait := today.Sub(now)
			select {
			case <-ctx.Done():
				return
			case <-time.After(wait):
			}
			goalID, resp, err := do(MessageInfo[p].Message)
			if err == nil {
				log.Printf("success:\nGoalID:%s,resp:%s\n", goalID, resp)
			}
			for i, errTime := 1, 1*time.Second; err != nil; i++ {
				time.Sleep(errTime)
				_, resp, err = do(MessageInfo[p].Message)
				log.Printf("error:\nGoalID:%s,retry %d times, send result:%s\n", goalID, i, resp)
				if i == 3 && err != nil {
					cancel()
					return
				}
				errTime *= 2
			}
			p++
			if p == Len {
				p -= Len
			}
		}
	}()
}
