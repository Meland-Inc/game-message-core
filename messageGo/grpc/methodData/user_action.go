package methodData

import (
	base_data "game-message-core/grpc/baseData"
)

// 更新玩家使用的装备
type UpdateUsedAvatarInput struct {
	MsgVersion   int64                        `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	UserId       int64                        `json:"userId"`
	UsingAvatars []base_data.GrpcPlayerAvatar `json:"usingAvatars"`
	CurProfile   base_data.GrpcEntityProfile  `json:"curProfile"`
}
type UpdateUsedAvatarOutput struct {
	MsgVersion int64 `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	Success    bool  `json:"success"`
}

// 更新玩家战斗属性(升级/装备槽等级变更)
type UpdateUserProfileInput struct {
	MsgVersion int64                       `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	UserId     int64                       `json:"userId"`
	CurProfile base_data.GrpcEntityProfile `json:"curProfile"`
}
type UpdateUserProfileOutput struct {
	MsgVersion int64 `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	Success    bool  `json:"success"`
}

// 获取user 详细数据
type GetUserDataInput struct {
	MsgVersion int64 `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	UserId     int64 `json:"userId"`
}
type GetUserDataOutput struct {
	MsgVersion int64                        `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	Success    bool                         `json:"success"`
	ErrMsg     string                       `json:"errMsg"`
	BaseData   base_data.GrpcPlayerBaseData `json:"baseData"`
	Profile    base_data.GrpcEntityProfile  `json:"profile"`
	Feature    base_data.GrpcPlayerFeature  `json:"feature"`
	MapId      int32                        `json:"mapId"`
	Pos        base_data.GrpcVector3        `json:"pos"`
	Dir        base_data.GrpcVector3        `json:"dir"`
	Avatars    []base_data.GrpcPlayerAvatar `json:"avatars"`
	TalentData base_data.GrpcTalentData     `json:"talentData"`
}
