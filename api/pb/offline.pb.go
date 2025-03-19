// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v6.30.1
// source: offline.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetOfflineMsgReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Uid           int32                  `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetOfflineMsgReq) Reset() {
	*x = GetOfflineMsgReq{}
	mi := &file_offline_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOfflineMsgReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOfflineMsgReq) ProtoMessage() {}

func (x *GetOfflineMsgReq) ProtoReflect() protoreflect.Message {
	mi := &file_offline_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOfflineMsgReq.ProtoReflect.Descriptor instead.
func (*GetOfflineMsgReq) Descriptor() ([]byte, []int) {
	return file_offline_proto_rawDescGZIP(), []int{0}
}

func (x *GetOfflineMsgReq) GetUid() int32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type GetOfflineMsgResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Msgs          []*ChatMsg             `protobuf:"bytes,1,rep,name=msgs,proto3" json:"msgs,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetOfflineMsgResp) Reset() {
	*x = GetOfflineMsgResp{}
	mi := &file_offline_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOfflineMsgResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOfflineMsgResp) ProtoMessage() {}

func (x *GetOfflineMsgResp) ProtoReflect() protoreflect.Message {
	mi := &file_offline_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOfflineMsgResp.ProtoReflect.Descriptor instead.
func (*GetOfflineMsgResp) Descriptor() ([]byte, []int) {
	return file_offline_proto_rawDescGZIP(), []int{1}
}

func (x *GetOfflineMsgResp) GetMsgs() []*ChatMsg {
	if x != nil {
		return x.Msgs
	}
	return nil
}

type PutMsgReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int32                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Msg           *ChatMsg               `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PutMsgReq) Reset() {
	*x = PutMsgReq{}
	mi := &file_offline_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PutMsgReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutMsgReq) ProtoMessage() {}

func (x *PutMsgReq) ProtoReflect() protoreflect.Message {
	mi := &file_offline_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutMsgReq.ProtoReflect.Descriptor instead.
func (*PutMsgReq) Descriptor() ([]byte, []int) {
	return file_offline_proto_rawDescGZIP(), []int{2}
}

func (x *PutMsgReq) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *PutMsgReq) GetMsg() *ChatMsg {
	if x != nil {
		return x.Msg
	}
	return nil
}

type PutMsgResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Succ          bool                   `protobuf:"varint,1,opt,name=succ,proto3" json:"succ,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PutMsgResp) Reset() {
	*x = PutMsgResp{}
	mi := &file_offline_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PutMsgResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutMsgResp) ProtoMessage() {}

func (x *PutMsgResp) ProtoReflect() protoreflect.Message {
	mi := &file_offline_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutMsgResp.ProtoReflect.Descriptor instead.
func (*PutMsgResp) Descriptor() ([]byte, []int) {
	return file_offline_proto_rawDescGZIP(), []int{3}
}

func (x *PutMsgResp) GetSucc() bool {
	if x != nil {
		return x.Succ
	}
	return false
}

type RemoveMsgReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	MsgIds        []int64                `protobuf:"varint,1,rep,packed,name=msg_ids,json=msgIds,proto3" json:"msg_ids,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveMsgReq) Reset() {
	*x = RemoveMsgReq{}
	mi := &file_offline_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveMsgReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveMsgReq) ProtoMessage() {}

func (x *RemoveMsgReq) ProtoReflect() protoreflect.Message {
	mi := &file_offline_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveMsgReq.ProtoReflect.Descriptor instead.
func (*RemoveMsgReq) Descriptor() ([]byte, []int) {
	return file_offline_proto_rawDescGZIP(), []int{4}
}

func (x *RemoveMsgReq) GetMsgIds() []int64 {
	if x != nil {
		return x.MsgIds
	}
	return nil
}

type RemoveMsgResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Succ          bool                   `protobuf:"varint,1,opt,name=succ,proto3" json:"succ,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveMsgResp) Reset() {
	*x = RemoveMsgResp{}
	mi := &file_offline_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveMsgResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveMsgResp) ProtoMessage() {}

func (x *RemoveMsgResp) ProtoReflect() protoreflect.Message {
	mi := &file_offline_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveMsgResp.ProtoReflect.Descriptor instead.
func (*RemoveMsgResp) Descriptor() ([]byte, []int) {
	return file_offline_proto_rawDescGZIP(), []int{5}
}

func (x *RemoveMsgResp) GetSucc() bool {
	if x != nil {
		return x.Succ
	}
	return false
}

var File_offline_proto protoreflect.FileDescriptor

var file_offline_proto_rawDesc = string([]byte{
	0x0a, 0x0d, 0x6f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x24,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x4d, 0x73, 0x67, 0x52,
	0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x75, 0x69, 0x64, 0x22, 0x36, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4f, 0x66, 0x66, 0x6c, 0x69,
	0x6e, 0x65, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12, 0x21, 0x0a, 0x04, 0x6d, 0x73, 0x67,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x63, 0x68,
	0x61, 0x74, 0x5f, 0x6d, 0x73, 0x67, 0x52, 0x04, 0x6d, 0x73, 0x67, 0x73, 0x22, 0x45, 0x0a, 0x09,
	0x50, 0x75, 0x74, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1f, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x6d, 0x73, 0x67, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x22, 0x20, 0x0a, 0x0a, 0x50, 0x75, 0x74, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x75, 0x63, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x04, 0x73, 0x75, 0x63, 0x63, 0x22, 0x27, 0x0a, 0x0c, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4d,
	0x73, 0x67, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x73, 0x67, 0x5f, 0x69, 0x64, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x06, 0x6d, 0x73, 0x67, 0x49, 0x64, 0x73, 0x22, 0x23,
	0x0a, 0x0d, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x75, 0x63, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x73,
	0x75, 0x63, 0x63, 0x32, 0x90, 0x01, 0x0a, 0x07, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x12,
	0x36, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x4d, 0x73, 0x67,
	0x12, 0x11, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x4d, 0x73, 0x67,
	0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65,
	0x4d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12, 0x21, 0x0a, 0x06, 0x50, 0x75, 0x74, 0x4d, 0x73,
	0x67, 0x12, 0x0a, 0x2e, 0x50, 0x75, 0x74, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x0b, 0x2e,
	0x50, 0x75, 0x74, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2a, 0x0a, 0x09, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x4d, 0x73, 0x67, 0x12, 0x0d, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4d,
	0x73, 0x67, 0x52, 0x65, 0x73, 0x70, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_offline_proto_rawDescOnce sync.Once
	file_offline_proto_rawDescData []byte
)

func file_offline_proto_rawDescGZIP() []byte {
	file_offline_proto_rawDescOnce.Do(func() {
		file_offline_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_offline_proto_rawDesc), len(file_offline_proto_rawDesc)))
	})
	return file_offline_proto_rawDescData
}

var file_offline_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_offline_proto_goTypes = []any{
	(*GetOfflineMsgReq)(nil),  // 0: GetOfflineMsgReq
	(*GetOfflineMsgResp)(nil), // 1: GetOfflineMsgResp
	(*PutMsgReq)(nil),         // 2: PutMsgReq
	(*PutMsgResp)(nil),        // 3: PutMsgResp
	(*RemoveMsgReq)(nil),      // 4: RemoveMsgReq
	(*RemoveMsgResp)(nil),     // 5: RemoveMsgResp
	(*ChatMsg)(nil),           // 6: msg.chat_msg
}
var file_offline_proto_depIdxs = []int32{
	6, // 0: GetOfflineMsgResp.msgs:type_name -> msg.chat_msg
	6, // 1: PutMsgReq.msg:type_name -> msg.chat_msg
	0, // 2: Offline.GetOfflineMsg:input_type -> GetOfflineMsgReq
	2, // 3: Offline.PutMsg:input_type -> PutMsgReq
	4, // 4: Offline.RemoveMsg:input_type -> RemoveMsgReq
	1, // 5: Offline.GetOfflineMsg:output_type -> GetOfflineMsgResp
	3, // 6: Offline.PutMsg:output_type -> PutMsgResp
	5, // 7: Offline.RemoveMsg:output_type -> RemoveMsgResp
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_offline_proto_init() }
func file_offline_proto_init() {
	if File_offline_proto != nil {
		return
	}
	file_message_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_offline_proto_rawDesc), len(file_offline_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_offline_proto_goTypes,
		DependencyIndexes: file_offline_proto_depIdxs,
		MessageInfos:      file_offline_proto_msgTypes,
	}.Build()
	File_offline_proto = out.File
	file_offline_proto_goTypes = nil
	file_offline_proto_depIdxs = nil
}
