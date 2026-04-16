package model

type SendWay string

func (S SendWay) String() string {
	return string(S)
}

var PrivateAction SendWay = "send_private_msg"
var GroupAction SendWay = "send_group_msg"
