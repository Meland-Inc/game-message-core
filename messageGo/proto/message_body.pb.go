// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: message_body.proto

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

// 消息交互体.
type Envelope struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type         EnvelopeType `protobuf:"varint,1,opt,name=type,proto3,enum=gameMessageCore.EnvelopeType" json:"type,omitempty"`
	SeqId        int32        `protobuf:"varint,2,opt,name=seq_id,json=seqId,proto3" json:"seq_id,omitempty"`
	ErrorCode    int32        `protobuf:"varint,3,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	ErrorMessage string       `protobuf:"bytes,4,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	// Types that are assignable to Payload:
	//	*Envelope_QueryPlayerRequest
	//	*Envelope_QueryPlayerResponse
	//	*Envelope_PingRequest
	//	*Envelope_PingResponse
	Payload isEnvelope_Payload `protobuf_oneof:"payload"`
}

func (x *Envelope) Reset() {
	*x = Envelope{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_body_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Envelope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Envelope) ProtoMessage() {}

func (x *Envelope) ProtoReflect() protoreflect.Message {
	mi := &file_message_body_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Envelope.ProtoReflect.Descriptor instead.
func (*Envelope) Descriptor() ([]byte, []int) {
	return file_message_body_proto_rawDescGZIP(), []int{0}
}

func (x *Envelope) GetType() EnvelopeType {
	if x != nil {
		return x.Type
	}
	return EnvelopeType_Unknown
}

func (x *Envelope) GetSeqId() int32 {
	if x != nil {
		return x.SeqId
	}
	return 0
}

func (x *Envelope) GetErrorCode() int32 {
	if x != nil {
		return x.ErrorCode
	}
	return 0
}

func (x *Envelope) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (m *Envelope) GetPayload() isEnvelope_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *Envelope) GetQueryPlayerRequest() *QueryPlayerRequest {
	if x, ok := x.GetPayload().(*Envelope_QueryPlayerRequest); ok {
		return x.QueryPlayerRequest
	}
	return nil
}

func (x *Envelope) GetQueryPlayerResponse() *QueryPlayerResponse {
	if x, ok := x.GetPayload().(*Envelope_QueryPlayerResponse); ok {
		return x.QueryPlayerResponse
	}
	return nil
}

func (x *Envelope) GetPingRequest() *PingRequest {
	if x, ok := x.GetPayload().(*Envelope_PingRequest); ok {
		return x.PingRequest
	}
	return nil
}

func (x *Envelope) GetPingResponse() *PingResponse {
	if x, ok := x.GetPayload().(*Envelope_PingResponse); ok {
		return x.PingResponse
	}
	return nil
}

type isEnvelope_Payload interface {
	isEnvelope_Payload()
}

type Envelope_QueryPlayerRequest struct {
	//accountSer协议 : 0x02ZZZZ  账号服务 协议 ************************************
	QueryPlayerRequest *QueryPlayerRequest `protobuf:"bytes,131073,opt,name=query_player_request,json=queryPlayerRequest,proto3,oneof"`
}

type Envelope_QueryPlayerResponse struct {
	QueryPlayerResponse *QueryPlayerResponse `protobuf:"bytes,131074,opt,name=query_player_response,json=queryPlayerResponse,proto3,oneof"`
}

type Envelope_PingRequest struct {
	//getaway   协议 : 0x06ZZZZ  网关服务 协议 ************************************
	PingRequest *PingRequest `protobuf:"bytes,393217,opt,name=ping_request,json=pingRequest,proto3,oneof"`
}

type Envelope_PingResponse struct {
	PingResponse *PingResponse `protobuf:"bytes,393218,opt,name=ping_response,json=pingResponse,proto3,oneof"`
}

func (*Envelope_QueryPlayerRequest) isEnvelope_Payload() {}

func (*Envelope_QueryPlayerResponse) isEnvelope_Payload() {}

func (*Envelope_PingRequest) isEnvelope_Payload() {}

func (*Envelope_PingResponse) isEnvelope_Payload() {}

//accountSer协议 : 0x02ZZZZ  账号服务 协议 ************************************
type QueryPlayerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *QueryPlayerRequest) Reset() {
	*x = QueryPlayerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_body_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryPlayerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryPlayerRequest) ProtoMessage() {}

func (x *QueryPlayerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_message_body_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryPlayerRequest.ProtoReflect.Descriptor instead.
func (*QueryPlayerRequest) Descriptor() ([]byte, []int) {
	return file_message_body_proto_rawDescGZIP(), []int{1}
}

func (x *QueryPlayerRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type QueryPlayerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Player *PlayerBaseData `protobuf:"bytes,1,opt,name=player,proto3" json:"player,omitempty"` // 玩家信息
}

func (x *QueryPlayerResponse) Reset() {
	*x = QueryPlayerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_body_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryPlayerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryPlayerResponse) ProtoMessage() {}

func (x *QueryPlayerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_message_body_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryPlayerResponse.ProtoReflect.Descriptor instead.
func (*QueryPlayerResponse) Descriptor() ([]byte, []int) {
	return file_message_body_proto_rawDescGZIP(), []int{2}
}

func (x *QueryPlayerResponse) GetPlayer() *PlayerBaseData {
	if x != nil {
		return x.Player
	}
	return nil
}

//getaway   协议 : 0x06ZZZZ  网关服务 协议 ************************************
type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_body_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_message_body_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_message_body_proto_rawDescGZIP(), []int{3}
}

type PingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingResponse) Reset() {
	*x = PingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_body_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResponse) ProtoMessage() {}

func (x *PingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_message_body_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResponse.ProtoReflect.Descriptor instead.
func (*PingResponse) Descriptor() ([]byte, []int) {
	return file_message_body_proto_rawDescGZIP(), []int{4}
}

var File_message_body_proto protoreflect.FileDescriptor

var file_message_body_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x67, 0x61, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x43, 0x6f, 0x72, 0x65, 0x1a, 0x11, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63,
	0x6d, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe9, 0x03, 0x0a, 0x08, 0x45, 0x6e, 0x76, 0x65, 0x6c,
	0x6f, 0x70, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1d, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43,
	0x6f, 0x72, 0x65, 0x2e, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x65, 0x71, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x65, 0x71, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x59, 0x0a, 0x14, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x81, 0x80, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43,
	0x6f, 0x72, 0x65, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x12, 0x71, 0x75, 0x65, 0x72, 0x79, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x5c, 0x0a, 0x15,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x82, 0x80, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e,
	0x67, 0x61, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x72, 0x65, 0x2e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x13, 0x71, 0x75, 0x65, 0x72, 0x79, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0c, 0x70, 0x69,
	0x6e, 0x67, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x81, 0x80, 0x18, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x43, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x48, 0x00, 0x52, 0x0b, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x46, 0x0a, 0x0d, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x18, 0x82, 0x80, 0x18, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x0c, 0x70, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x22, 0x2a, 0x0a, 0x12, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x4e,
	0x0a, 0x13, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x42, 0x61,
	0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x22, 0x0d,
	0x0a, 0x0b, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0e, 0x0a,
	0x0c, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x13, 0x5a,
	0x11, 0x2e, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x47, 0x6f, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_body_proto_rawDescOnce sync.Once
	file_message_body_proto_rawDescData = file_message_body_proto_rawDesc
)

func file_message_body_proto_rawDescGZIP() []byte {
	file_message_body_proto_rawDescOnce.Do(func() {
		file_message_body_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_body_proto_rawDescData)
	})
	return file_message_body_proto_rawDescData
}

var file_message_body_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_message_body_proto_goTypes = []interface{}{
	(*Envelope)(nil),            // 0: gameMessageCore.Envelope
	(*QueryPlayerRequest)(nil),  // 1: gameMessageCore.QueryPlayerRequest
	(*QueryPlayerResponse)(nil), // 2: gameMessageCore.QueryPlayerResponse
	(*PingRequest)(nil),         // 3: gameMessageCore.PingRequest
	(*PingResponse)(nil),        // 4: gameMessageCore.PingResponse
	(EnvelopeType)(0),           // 5: gameMessageCore.EnvelopeType
	(*PlayerBaseData)(nil),      // 6: gameMessageCore.PlayerBaseData
}
var file_message_body_proto_depIdxs = []int32{
	5, // 0: gameMessageCore.Envelope.type:type_name -> gameMessageCore.EnvelopeType
	1, // 1: gameMessageCore.Envelope.query_player_request:type_name -> gameMessageCore.QueryPlayerRequest
	2, // 2: gameMessageCore.Envelope.query_player_response:type_name -> gameMessageCore.QueryPlayerResponse
	3, // 3: gameMessageCore.Envelope.ping_request:type_name -> gameMessageCore.PingRequest
	4, // 4: gameMessageCore.Envelope.ping_response:type_name -> gameMessageCore.PingResponse
	6, // 5: gameMessageCore.QueryPlayerResponse.player:type_name -> gameMessageCore.PlayerBaseData
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_message_body_proto_init() }
func file_message_body_proto_init() {
	if File_message_body_proto != nil {
		return
	}
	file_message_cmd_proto_init()
	file_player_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_message_body_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Envelope); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_message_body_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryPlayerRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_message_body_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryPlayerResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_message_body_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_message_body_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_message_body_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Envelope_QueryPlayerRequest)(nil),
		(*Envelope_QueryPlayerResponse)(nil),
		(*Envelope_PingRequest)(nil),
		(*Envelope_PingResponse)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_message_body_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_body_proto_goTypes,
		DependencyIndexes: file_message_body_proto_depIdxs,
		MessageInfos:      file_message_body_proto_msgTypes,
	}.Build()
	File_message_body_proto = out.File
	file_message_body_proto_rawDesc = nil
	file_message_body_proto_goTypes = nil
	file_message_body_proto_depIdxs = nil
}
