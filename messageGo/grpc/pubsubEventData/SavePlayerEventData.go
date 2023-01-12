package pubsubEventData

import (
	base_data "game-message-core/grpc/baseData"
)

type SavePlayerEventData struct {
	MsgVersion     int64                 `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	FormService    base_data.ServiceData `json:"formService"`
	UserId         int64                 `json:"userId"`
	PosX           float32               `json:"posX"`
	PosY           float32               `json:"posY"`
	PosZ           float32               `json:"posZ"`
	CurHP          int32                 `json:"curHP"`
	DirX           float32               `json:"dirX"`
	DirY           float32               `json:"dirY"`
	DirZ           float32               `json:"dirZ"`
}
