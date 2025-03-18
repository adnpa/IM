// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v6.30.1
// source: message.proto

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

// todo 后续改为tcp 字节流协议
type CommonMsg struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Version       string                 `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Cmd           int32                  `protobuf:"varint,2,opt,name=cmd,proto3" json:"cmd,omitempty"`
	Size          int32                  `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	Body          []byte                 `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CommonMsg) Reset() {
	*x = CommonMsg{}
	mi := &file_message_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CommonMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonMsg) ProtoMessage() {}

func (x *CommonMsg) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonMsg.ProtoReflect.Descriptor instead.
func (*CommonMsg) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

func (x *CommonMsg) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *CommonMsg) GetCmd() int32 {
	if x != nil {
		return x.Cmd
	}
	return 0
}

func (x *CommonMsg) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *CommonMsg) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

type ChatMsg struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                              // 消息ID
	Typ           int32                  `protobuf:"varint,2,opt,name=typ,proto3" json:"typ,omitempty"`                            // 消息类型
	From          int64                  `protobuf:"varint,3,opt,name=from,proto3" json:"from,omitempty"`                          // 发送者ID
	To            int64                  `protobuf:"varint,4,opt,name=to,proto3" json:"to,omitempty"`                              // 接收者ID或群组ID
	Media         int32                  `protobuf:"varint,5,opt,name=media,proto3" json:"media,omitempty"`                        // 媒体类型
	Content       string                 `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`                     // 消息内容
	Pic           string                 `protobuf:"bytes,7,opt,name=pic,proto3" json:"pic,omitempty"`                             // 缩略图URL
	Url           string                 `protobuf:"bytes,8,opt,name=url,proto3" json:"url,omitempty"`                             // 服务URL
	Memo          string                 `protobuf:"bytes,9,opt,name=memo,proto3" json:"memo,omitempty"`                           // 备注
	Amount        int32                  `protobuf:"varint,10,opt,name=amount,proto3" json:"amount,omitempty"`                     // 数字相关，如语音长度等
	Seq           int64                  `protobuf:"varint,11,opt,name=seq,proto3" json:"seq,omitempty"`                           // 序列号
	RecverId      int64                  `protobuf:"varint,12,opt,name=recver_id,json=recverId,proto3" json:"recver_id,omitempty"` // 接收者ID
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ChatMsg) Reset() {
	*x = ChatMsg{}
	mi := &file_message_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChatMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatMsg) ProtoMessage() {}

func (x *ChatMsg) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatMsg.ProtoReflect.Descriptor instead.
func (*ChatMsg) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{1}
}

func (x *ChatMsg) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ChatMsg) GetTyp() int32 {
	if x != nil {
		return x.Typ
	}
	return 0
}

func (x *ChatMsg) GetFrom() int64 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *ChatMsg) GetTo() int64 {
	if x != nil {
		return x.To
	}
	return 0
}

func (x *ChatMsg) GetMedia() int32 {
	if x != nil {
		return x.Media
	}
	return 0
}

func (x *ChatMsg) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *ChatMsg) GetPic() string {
	if x != nil {
		return x.Pic
	}
	return ""
}

func (x *ChatMsg) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *ChatMsg) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *ChatMsg) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *ChatMsg) GetSeq() int64 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (x *ChatMsg) GetRecverId() int64 {
	if x != nil {
		return x.RecverId
	}
	return 0
}

var File_message_proto protoreflect.FileDescriptor

var file_message_proto_rawDesc = string([]byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x6d, 0x73, 0x67, 0x22, 0x60, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x6d,
	0x73, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03,
	0x63, 0x6d, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x63, 0x6d, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0xff, 0x01, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x74, 0x5f,
	0x6d, 0x73, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x79, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x74, 0x79, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x65, 0x64,
	0x69, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x63,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x69, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a,
	0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d,
	0x6f, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x71,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x73, 0x65, 0x71, 0x12, 0x1b, 0x0a, 0x09, 0x72,
	0x65, 0x63, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x72, 0x65, 0x63, 0x76, 0x65, 0x72, 0x49, 0x64, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_message_proto_rawDescOnce sync.Once
	file_message_proto_rawDescData []byte
)

func file_message_proto_rawDescGZIP() []byte {
	file_message_proto_rawDescOnce.Do(func() {
		file_message_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_message_proto_rawDesc), len(file_message_proto_rawDesc)))
	})
	return file_message_proto_rawDescData
}

var file_message_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_message_proto_goTypes = []any{
	(*CommonMsg)(nil), // 0: msg.common_msg
	(*ChatMsg)(nil),   // 1: msg.chat_msg
}
var file_message_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_message_proto_init() }
func file_message_proto_init() {
	if File_message_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_message_proto_rawDesc), len(file_message_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_proto_goTypes,
		DependencyIndexes: file_message_proto_depIdxs,
		MessageInfos:      file_message_proto_msgTypes,
	}.Build()
	File_message_proto = out.File
	file_message_proto_goTypes = nil
	file_message_proto_depIdxs = nil
}
