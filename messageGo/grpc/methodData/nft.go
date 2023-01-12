package methodData

import (
	base_data "game-message-core/grpc/baseData"
)

type TakeNftData struct {
	NftId   string `json:"nftId"`
	ItemCid int32  `json:"itemCid"`
	Num     int32  `json:"num"`
}

// 通知mainService 扣除User nft input
type MainServiceActionTakeNftInput struct {
	UserId   int64         `json:"userId"`
	TakeNfts []TakeNftData `json:"takeNfts"`
}
type MainServiceActionTakeNftOutput struct {
	Success   bool   `json:"success"`
	FailedMsg string `json:"failedMsg"`
}

// 通知mainService make User nft input
type MainServiceActionMintNftInput struct {
	UserId int64                      `json:"userId"`
	Item   base_data.GrpcItemBaseInfo `json:"item"`
}
type MainServiceActionMintNftOutput struct {
	Success   bool   `json:"success"`
	FailedMsg string `json:"failedMsg"`
}
