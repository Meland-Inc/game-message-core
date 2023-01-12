package pubsubEventData

import (
	base_data "game-message-core/grpc/baseData"
)

type ServiceStartedEvent struct {
	MsgVersion int64                 `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	Service    base_data.ServiceData `json:"service"`    //
}
