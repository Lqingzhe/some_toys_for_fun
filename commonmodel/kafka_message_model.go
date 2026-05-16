package commonmodel

type KafkaGroupNotice struct {
	GoalUserID  []int64
	SessionID   int64
	Data        map[string]any
	MessageType MessageType
	MessageCode MessageCode
}

type KafkaNewMessageNotice struct {
	GoalUserID  int64
	SessionID   int64
	MessageType MessageType
	MessageCode MessageCode
}
type KafkaSystemMessage struct {
	GoalUserID  []int64
	Data        map[string]any
	MessageType MessageType
	MessageCode MessageCode
}
