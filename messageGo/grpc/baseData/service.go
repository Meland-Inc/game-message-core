package base_data

import "game-message-core/proto"

// service data
type ServiceData struct {
	AppId           string                    `json:"appId"`
	ServiceType     proto.ServiceType         `json:"serviceType"`
	SceneSerSubType proto.SceneServiceSubType `json:"sceneSerSubType"`
	Owner           int64                     `json:"owner"`
	Host            string                    `json:"host"`
	Port            int32                     `json:"port"`
	MapId           int32                     `json:"mapId"`
	Online          int32                     `json:"online"`
	MaxOnline       int32                     `json:"maxOnline"`
	CreatedAt       int64                     `json:"createdAt"`
	UpdatedAt       int64                     `json:"updatedAt"`
}
