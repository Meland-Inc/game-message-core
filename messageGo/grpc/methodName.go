package grpc

type ManagerServiceAction string

const (
	ManagerServiceActionServiceTime        ManagerServiceAction = "QueryServerTime"
	ManagerServiceActionRegister           ManagerServiceAction = "ManagerActionServiceRegister"
	ManagerServiceActionStartService       ManagerServiceAction = "ManagerActionStartService"
	ManagerServiceActionSelectService      ManagerServiceAction = "ManagerActionSelectService"
	ManagerServiceActionMultiSelectService ManagerServiceAction = "ManagerActionMultiSelectService"
)

type MainServiceAction string

const (
	MainServiceActionTakeNFT     MainServiceAction = "MainServiceActionTakeNFT"
	MainServiceActionGetAllBuild MainServiceAction = "MainServiceActionGetAllBuild"
	MainServiceActionGetHomeData MainServiceAction = "MainServiceActionGetHomeData"
)

type ProtoMessageAction string

const (
	ProtoMessageActionPullClientMessage         ProtoMessageAction = "PullClientMessage"
	ProtoMessageActionBroadCastToClient         ProtoMessageAction = "BroadCastToClient"
	ProtoMessageActionMultipleBroadCastToClient ProtoMessageAction = "MultipleBroadCastToClient"
)

type UserAction string

const (
	UserActionGetUserData       UserAction = "GetUserData"
	UserActionUpdateUsedAvatar  UserAction = "UpdateUsedAvatar"
	UserActionUpdateUserProfile UserAction = "UpdateUserProfile"
)
