package pubsubEventData

import (
	base_data "game-message-core/grpc/baseData"
	"game-message-core/proto"
)

type TaskFinishEvent struct {
	MsgVersion   int64                        `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	UserId       int64                        `json:"userId"`
	TaskListType proto.TaskListType           `json:"taskListType"`
	TaskId       int32                        `json:"taskId"`
	RewardItem   []base_data.GrpcItemBaseInfo `json:"rewardItem"`
}
