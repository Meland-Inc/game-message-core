// 应用交互消息定义.

//***************************************************************************
// 消息协议号 ： 服务端主动推送的以  BroadCast开头  如 xxx  ***********************
// 消息结构   ： server -> client 以 Response   如 BroadCastXXResponse结尾******
// 消息结构   ： server回复client  以 Response   如 xxxResponse 结尾  ***********
// 消息结构   ： client->server   以 Request    如 xxxRequest  结尾  ***********
//***************************************************************************

//***************************************************************************
//协议ID规则      : 0xXXZZZZ, 其中 XX 为服务ID头, ZZZZ为具体的协议ID, 示例如下 ****
//保留id协议      : 0x00ZZZZ   ***********************************************
//mainServer协议 : 0x01ZZZZ  主数据服务协议************************************
//accountSer协议 : 0x02ZZZZ  账号服务 协议 ************************************
//screenSer 协议 : 0x03ZZZZ  战斗服务 协议 ************************************
//taskServer协议 : 0x04ZZZZ  任务服务 协议 ************************************
//chatServer协议 : 0x05ZZZZ  聊天服务 协议 ************************************
//getaway   协议 : 0x06ZZZZ  网关服务 协议 ************************************
//ServiceMgr协议 : 0x09ZZZZ  服务管理 协议(客户端不能调用) **********************
//CTRlServer协议 : 0x0aZZZZ  控制服务 协议 ************************************
//**************************************************************************

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: message_cmd.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EnvelopeType int32

const (
	//保留id协议      : 0x00ZZZZ   ***********************************************
	EnvelopeType_Unknown EnvelopeType = 0
	//mainServer协议 : 0x01ZZZZ  主数据服务协议************************************
	EnvelopeType_ItemGet                 EnvelopeType = 65537
	EnvelopeType_ItemUse                 EnvelopeType = 65539
	EnvelopeType_ItemDrop                EnvelopeType = 65541
	EnvelopeType_UpdateAvatar            EnvelopeType = 65543
	EnvelopeType_UnloadAvatar            EnvelopeType = 65545
	EnvelopeType_BroadCastItemAdd        EnvelopeType = 65553 // 添加道具
	EnvelopeType_BroadCastItemUpdate     EnvelopeType = 65554 // 更新道具
	EnvelopeType_BroadCastItemDel        EnvelopeType = 65555 // del道具
	EnvelopeType_BroadCastUpdateItemSlot EnvelopeType = 65556 // 推送 玩家道具槽
	EnvelopeType_GetItemSlot             EnvelopeType = 65557 // 查询 玩家道具槽信息
	EnvelopeType_UpgradeItemSlot         EnvelopeType = 65559 // 升级 玩家道具槽
	EnvelopeType_SigninPlayer            EnvelopeType = 65561 // 登录角色
	EnvelopeType_SignOutPlayer           EnvelopeType = 65569 // 角色退出游戏
	//accountSer协议 : 0x02ZZZZ  账号服务 协议 ************************************
	EnvelopeType_QueryPlayer  EnvelopeType = 131073
	EnvelopeType_CreatePlayer EnvelopeType = 131075
	//screenSer 协议 : 0x03ZZZZ  战斗服务 协议 ************************************
	EnvelopeType_EnterMap                 EnvelopeType = 196609
	EnvelopeType_UpdateSelfLocation       EnvelopeType = 196611
	EnvelopeType_UseSkill                 EnvelopeType = 196613
	EnvelopeType_PickFallingObject        EnvelopeType = 196615
	EnvelopeType_RespawnPlayer            EnvelopeType = 196617
	EnvelopeType_BroadCastInitMapElement  EnvelopeType = 196625
	EnvelopeType_BroadCastMapEntityUpdate EnvelopeType = 196626
	EnvelopeType_BroadCastEntityDestroy   EnvelopeType = 196627
	EnvelopeType_BroadCastEntityMove      EnvelopeType = 196628
	EnvelopeType_BroadCastEntityCombat    EnvelopeType = 196629
	EnvelopeType_BroadCastRespawnPlayer   EnvelopeType = 196630
	//taskServer协议 : 0x04ZZZZ  任务服务 协议 ************************************
	EnvelopeType_SelfTasks               EnvelopeType = 262144 // 查询玩家已接的任务列表（进行中状态）
	EnvelopeType_BroadCastUpdateTaskList EnvelopeType = 262145 // 任务链进度更新(推送)
	EnvelopeType_AcceptTask              EnvelopeType = 262146 // 领取任务
	EnvelopeType_AbandonmentTask         EnvelopeType = 262147 // 放弃任务(任务有保护时间)
	EnvelopeType_TaskReward              EnvelopeType = 262148 // 获取任务奖励(附带提交任务功能)
	EnvelopeType_TaskListReward          EnvelopeType = 262149 // 获取任务链奖励
	EnvelopeType_UpgradeTaskProgress     EnvelopeType = 262150 // 上报更新任务进度
	EnvelopeType_BroadCastTaskReward     EnvelopeType = 262151 // 推送获取的任务奖励
	//getaway   协议 : 0x06ZZZZ  网关服务 协议 ************************************
	EnvelopeType_Ping EnvelopeType = 393217
)

// Enum value maps for EnvelopeType.
var (
	EnvelopeType_name = map[int32]string{
		0:      "Unknown",
		65537:  "ItemGet",
		65539:  "ItemUse",
		65541:  "ItemDrop",
		65543:  "UpdateAvatar",
		65545:  "UnloadAvatar",
		65553:  "BroadCastItemAdd",
		65554:  "BroadCastItemUpdate",
		65555:  "BroadCastItemDel",
		65556:  "BroadCastUpdateItemSlot",
		65557:  "GetItemSlot",
		65559:  "UpgradeItemSlot",
		65561:  "SigninPlayer",
		65569:  "SignOutPlayer",
		131073: "QueryPlayer",
		131075: "CreatePlayer",
		196609: "EnterMap",
		196611: "UpdateSelfLocation",
		196613: "UseSkill",
		196615: "PickFallingObject",
		196617: "RespawnPlayer",
		196625: "BroadCastInitMapElement",
		196626: "BroadCastMapEntityUpdate",
		196627: "BroadCastEntityDestroy",
		196628: "BroadCastEntityMove",
		196629: "BroadCastEntityCombat",
		196630: "BroadCastRespawnPlayer",
		262144: "SelfTasks",
		262145: "BroadCastUpdateTaskList",
		262146: "AcceptTask",
		262147: "AbandonmentTask",
		262148: "TaskReward",
		262149: "TaskListReward",
		262150: "UpgradeTaskProgress",
		262151: "BroadCastTaskReward",
		393217: "Ping",
	}
	EnvelopeType_value = map[string]int32{
		"Unknown":                  0,
		"ItemGet":                  65537,
		"ItemUse":                  65539,
		"ItemDrop":                 65541,
		"UpdateAvatar":             65543,
		"UnloadAvatar":             65545,
		"BroadCastItemAdd":         65553,
		"BroadCastItemUpdate":      65554,
		"BroadCastItemDel":         65555,
		"BroadCastUpdateItemSlot":  65556,
		"GetItemSlot":              65557,
		"UpgradeItemSlot":          65559,
		"SigninPlayer":             65561,
		"SignOutPlayer":            65569,
		"QueryPlayer":              131073,
		"CreatePlayer":             131075,
		"EnterMap":                 196609,
		"UpdateSelfLocation":       196611,
		"UseSkill":                 196613,
		"PickFallingObject":        196615,
		"RespawnPlayer":            196617,
		"BroadCastInitMapElement":  196625,
		"BroadCastMapEntityUpdate": 196626,
		"BroadCastEntityDestroy":   196627,
		"BroadCastEntityMove":      196628,
		"BroadCastEntityCombat":    196629,
		"BroadCastRespawnPlayer":   196630,
		"SelfTasks":                262144,
		"BroadCastUpdateTaskList":  262145,
		"AcceptTask":               262146,
		"AbandonmentTask":          262147,
		"TaskReward":               262148,
		"TaskListReward":           262149,
		"UpgradeTaskProgress":      262150,
		"BroadCastTaskReward":      262151,
		"Ping":                     393217,
	}
)

func (x EnvelopeType) Enum() *EnvelopeType {
	p := new(EnvelopeType)
	*p = x
	return p
}

func (x EnvelopeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EnvelopeType) Descriptor() protoreflect.EnumDescriptor {
	return file_message_cmd_proto_enumTypes[0].Descriptor()
}

func (EnvelopeType) Type() protoreflect.EnumType {
	return &file_message_cmd_proto_enumTypes[0]
}

func (x EnvelopeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EnvelopeType.Descriptor instead.
func (EnvelopeType) EnumDescriptor() ([]byte, []int) {
	return file_message_cmd_proto_rawDescGZIP(), []int{0}
}

var File_message_cmd_proto protoreflect.FileDescriptor

var file_message_cmd_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6d, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x67, 0x61, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x43, 0x6f, 0x72, 0x65, 0x2a, 0xb3, 0x06, 0x0a, 0x0c, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e,
	0x10, 0x00, 0x12, 0x0d, 0x0a, 0x07, 0x49, 0x74, 0x65, 0x6d, 0x47, 0x65, 0x74, 0x10, 0x81, 0x80,
	0x04, 0x12, 0x0d, 0x0a, 0x07, 0x49, 0x74, 0x65, 0x6d, 0x55, 0x73, 0x65, 0x10, 0x83, 0x80, 0x04,
	0x12, 0x0e, 0x0a, 0x08, 0x49, 0x74, 0x65, 0x6d, 0x44, 0x72, 0x6f, 0x70, 0x10, 0x85, 0x80, 0x04,
	0x12, 0x12, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x10, 0x87, 0x80, 0x04, 0x12, 0x12, 0x0a, 0x0c, 0x55, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x10, 0x89, 0x80, 0x04, 0x12, 0x16, 0x0a, 0x10, 0x42, 0x72, 0x6f, 0x61,
	0x64, 0x43, 0x61, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x41, 0x64, 0x64, 0x10, 0x91, 0x80, 0x04,
	0x12, 0x19, 0x0a, 0x13, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x73, 0x74, 0x49, 0x74, 0x65,
	0x6d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x10, 0x92, 0x80, 0x04, 0x12, 0x16, 0x0a, 0x10, 0x42,
	0x72, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x44, 0x65, 0x6c, 0x10,
	0x93, 0x80, 0x04, 0x12, 0x1d, 0x0a, 0x17, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x73, 0x74,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x6c, 0x6f, 0x74, 0x10, 0x94,
	0x80, 0x04, 0x12, 0x11, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x6c, 0x6f,
	0x74, 0x10, 0x95, 0x80, 0x04, 0x12, 0x15, 0x0a, 0x0f, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65,
	0x49, 0x74, 0x65, 0x6d, 0x53, 0x6c, 0x6f, 0x74, 0x10, 0x97, 0x80, 0x04, 0x12, 0x12, 0x0a, 0x0c,
	0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x10, 0x99, 0x80, 0x04,
	0x12, 0x13, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x4f, 0x75, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x10, 0xa1, 0x80, 0x04, 0x12, 0x11, 0x0a, 0x0b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x10, 0x81, 0x80, 0x08, 0x12, 0x12, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x10, 0x83, 0x80, 0x08, 0x12, 0x0e, 0x0a, 0x08,
	0x45, 0x6e, 0x74, 0x65, 0x72, 0x4d, 0x61, 0x70, 0x10, 0x81, 0x80, 0x0c, 0x12, 0x18, 0x0a, 0x12,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x6c, 0x66, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x10, 0x83, 0x80, 0x0c, 0x12, 0x0e, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x53, 0x6b, 0x69,
	0x6c, 0x6c, 0x10, 0x85, 0x80, 0x0c, 0x12, 0x17, 0x0a, 0x11, 0x50, 0x69, 0x63, 0x6b, 0x46, 0x61,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x10, 0x87, 0x80, 0x0c, 0x12,
	0x13, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x70, 0x61, 0x77, 0x6e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x10, 0x89, 0x80, 0x0c, 0x12, 0x1d, 0x0a, 0x17, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x73,
	0x74, 0x49, 0x6e, 0x69, 0x74, 0x4d, 0x61, 0x70, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x10,
	0x91, 0x80, 0x0c, 0x12, 0x1e, 0x0a, 0x18, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x73, 0x74,
	0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x10,
	0x92, 0x80, 0x0c, 0x12, 0x1c, 0x0a, 0x16, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x73, 0x74,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x10, 0x93, 0x80,
	0x0c, 0x12, 0x19, 0x0a, 0x13, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x73, 0x74, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x4d, 0x6f, 0x76, 0x65, 0x10, 0x94, 0x80, 0x0c, 0x12, 0x1b, 0x0a, 0x15,
	0x42, 0x72, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x43,
	0x6f, 0x6d, 0x62, 0x61, 0x74, 0x10, 0x95, 0x80, 0x0c, 0x12, 0x1c, 0x0a, 0x16, 0x42, 0x72, 0x6f,
	0x61, 0x64, 0x43, 0x61, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x61, 0x77, 0x6e, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x10, 0x96, 0x80, 0x0c, 0x12, 0x0f, 0x0a, 0x09, 0x53, 0x65, 0x6c, 0x66, 0x54,
	0x61, 0x73, 0x6b, 0x73, 0x10, 0x80, 0x80, 0x10, 0x12, 0x1d, 0x0a, 0x17, 0x42, 0x72, 0x6f, 0x61,
	0x64, 0x43, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x4c,
	0x69, 0x73, 0x74, 0x10, 0x81, 0x80, 0x10, 0x12, 0x10, 0x0a, 0x0a, 0x41, 0x63, 0x63, 0x65, 0x70,
	0x74, 0x54, 0x61, 0x73, 0x6b, 0x10, 0x82, 0x80, 0x10, 0x12, 0x15, 0x0a, 0x0f, 0x41, 0x62, 0x61,
	0x6e, 0x64, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x10, 0x83, 0x80, 0x10,
	0x12, 0x10, 0x0a, 0x0a, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x10, 0x84,
	0x80, 0x10, 0x12, 0x14, 0x0a, 0x0e, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x10, 0x85, 0x80, 0x10, 0x12, 0x19, 0x0a, 0x13, 0x55, 0x70, 0x67, 0x72,
	0x61, 0x64, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x10,
	0x86, 0x80, 0x10, 0x12, 0x19, 0x0a, 0x13, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x73, 0x74,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x10, 0x87, 0x80, 0x10, 0x12, 0x0a,
	0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x10, 0x81, 0x80, 0x18, 0x42, 0x13, 0x5a, 0x11, 0x2e, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x47, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_cmd_proto_rawDescOnce sync.Once
	file_message_cmd_proto_rawDescData = file_message_cmd_proto_rawDesc
)

func file_message_cmd_proto_rawDescGZIP() []byte {
	file_message_cmd_proto_rawDescOnce.Do(func() {
		file_message_cmd_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_cmd_proto_rawDescData)
	})
	return file_message_cmd_proto_rawDescData
}

var file_message_cmd_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_message_cmd_proto_goTypes = []interface{}{
	(EnvelopeType)(0), // 0: gameMessageCore.EnvelopeType
}
var file_message_cmd_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_message_cmd_proto_init() }
func file_message_cmd_proto_init() {
	if File_message_cmd_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_message_cmd_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_cmd_proto_goTypes,
		DependencyIndexes: file_message_cmd_proto_depIdxs,
		EnumInfos:         file_message_cmd_proto_enumTypes,
	}.Build()
	File_message_cmd_proto = out.File
	file_message_cmd_proto_rawDesc = nil
	file_message_cmd_proto_goTypes = nil
	file_message_cmd_proto_depIdxs = nil
}
