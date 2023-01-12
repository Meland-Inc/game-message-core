package methodData

import (
	base_data "game-message-core/grpc/baseData"
	"game-message-core/proto"
)

// 向manager service 查询服务信息
type ManagerActionSelectServiceInput struct {
	ServiceType     proto.ServiceType         `json:"serviceType"`
	SceneSerSubType proto.SceneServiceSubType `json:"sceneSerSubType"`
	OwnerId         int64                     `json:"ownerId"`
	MapId           int32                     `json:"mapId"`
}

type ManagerActionSelectServiceOutput struct {
	ErrorCode    int32                 `json:"errorCode"` //TODO: NEED USE GLOBAL ERROR CODE
	ErrorMessage string                `json:"errorMessage"`
	Service      base_data.ServiceData `json:"service"`
}

// 向manager service 查询服务信息
type MultiSelectServiceInput struct {
	ServiceType     proto.ServiceType         `json:"serviceType"`
	SceneSerSubType proto.SceneServiceSubType `json:"sceneSerSubType"`
	OwnerId         int64                     `json:"ownerId"`
	MapId           int32                     `json:"mapId"`
}

type MultiSelectServiceOutput struct {
	ErrorCode    int32                   `json:"errorCode"` //TODO: NEED USE GLOBAL ERROR CODE
	ErrorMessage string                  `json:"errorMessage"`
	Services     []base_data.ServiceData `json:"services"`
}
