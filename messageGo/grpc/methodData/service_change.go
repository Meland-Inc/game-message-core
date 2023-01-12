package methodData

import (
	base_data "game-message-core/grpc/baseData"
)

/*
 * @Author: alex
 * @Date: 2022-12-9 10:00
 * @Description: user 切换 scene service 需要的grpc calls 相关交换协议
 * @FilePath: game-message-core/messageGo/grpc/methodData/service_change.go
 */

// user请求进入目标scene service request（预进入阶段）
type UserApplyEnterServiceInput struct {
	ApplyUser    int64                 `json:"applyUser"`
	UserPosition base_data.GrpcVector3 `json:"userPosition"`
	UserDir      base_data.GrpcVector3 `json:"userDir"`
	FromSer      base_data.ServiceData `json:"fromSer"`
}

// user 请求进入目标scene service  response （预进入阶段）
type UserApplyEnterServiceOutput struct {
	Success bool   `json:"success"`
	ErrMsg  string `json:"errMsg"`
}

// user正式进入目标scene service request（正式进入阶段）
type UserJoinServiceInput struct {
	UserId         int64                 `json:"userId"`
	UserAgentAppId string                `json:"userAgentAppId"`
	UserSocketId   string                `json:"userSocketId"`
	UserPosition   base_data.GrpcVector3 `json:"userPosition"`
	UserDir        base_data.GrpcVector3 `json:"userDir"`
	FromSer        base_data.ServiceData `json:"fromSer"`
}

// user正式进入目标scene service request（正式进入阶段）
type UserJoinServiceOutput struct {
	Success bool   `json:"success"`
	ErrMsg  string `json:"errMsg"`
}
