package base_data

import "game-message-core/proto"

type GrpcTalentLevel struct {
	TalentType proto.TalentType `json:"talentType"`
	Level      uint32           `json:"level"`
}

type GrpcTalentNodeData struct {
	NodeId uint32 `json:"nodeId"`
	Level  uint32 `json:"level"`
}

type GrpcTalentTree struct {
	TalentType proto.TalentType      `json:"talentType"`
	Nodes      []*GrpcTalentNodeData `json:"nodes"`
}

type GrpcTalentData struct {
	Levels []*GrpcTalentLevel `json:"levels"`
	Trees  []*GrpcTalentTree  `json:"trees"`
}
