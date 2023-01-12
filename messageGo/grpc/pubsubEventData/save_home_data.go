package pubsubEventData

import (
	base_data "game-message-core/grpc/baseData"
)

type SaveHomeEvent struct {
	MsgVersion int64                  `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	UserId     int64                  `json:"userId"`
	Data       base_data.GrpcHomeData `json:"data"`
}
