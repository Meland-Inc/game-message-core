package methodData

import (
	base_data "game-message-core/grpc/baseData"
)

// call main service get user home data
type MainServiceActionGetHomeDataInput struct {
	UserId int64 `json:"userId"`
}
type MainServiceActionGetHomeDataOutput struct {
	Success bool                   `json:"success"`
	ErrMsg  string                 `json:"errMsg"`
	UserId  int64                  `json:"userId"`
	Data    base_data.GrpcHomeData `json:"data"`
}
