package api

import (
	"errors"
	"harassment/model"
)

type SendStruct struct {
	goalID  string
	sendWay string
	path    string
}

type Operation func(*SendStruct)

func AddGoalID(goalID string) Operation {
	return func(send *SendStruct) {
		send.goalID = goalID
	}
}
func AddSendWay(bo bool) Operation {
	var sendWay model.SendWay
	if bo {
		sendWay = model.PrivateAction
	} else {
		sendWay = model.GroupAction
	}
	return func(send *SendStruct) {
		send.sendWay = sendWay.String()
	}
}
func AddPath(path string) Operation {
	return func(send *SendStruct) {
		send.path = path
	}
}

func NewSendGoal(input ...Operation) (*SendStruct, error) {
	newStruct := &SendStruct{}

	for _, op := range input {
		op(newStruct)
	}
	if newStruct.sendWay == "" {
		return nil, errors.New("without send way")
	}
	if newStruct.goalID == "" {
		return nil, errors.New("without goal ID")
	}
	return newStruct, nil
}
