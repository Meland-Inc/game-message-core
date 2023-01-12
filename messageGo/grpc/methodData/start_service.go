package methodData

import (
	"game-message-core/proto"
)

// 向manager service 请求启动服务 request
type StartServiceInput struct {
	ServiceType     proto.ServiceType         `json:"serviceType"`
	SceneSerSubType proto.SceneServiceSubType `json:"sceneSerSubType"`
	OwnerId         int64                     `json:"ownerId"`
	MapId           int32                     `json:"mapId"`
}

type StartServiceOutput struct {
	Success bool   `json:"success"`
	ErrMsg  string `json:"errMsg"`
}
