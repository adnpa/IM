// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.0
// source: pb_user.proto

package pb_user

import (
	pb_ws "github.com/adnpa/IM/pkg/pb/pb_ws"
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

type GetUserInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserIDList  []string `protobuf:"bytes,1,rep,name=userIDList,proto3" json:"userIDList,omitempty"`
	Token       string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	OperationID string   `protobuf:"bytes,3,opt,name=OperationID,proto3" json:"OperationID,omitempty"`
}

func (x *GetUserInfoReq) Reset() {
	*x = GetUserInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfoReq) ProtoMessage() {}

func (x *GetUserInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfoReq.ProtoReflect.Descriptor instead.
func (*GetUserInfoReq) Descriptor() ([]byte, []int) {
	return file_pb_user_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserInfoReq) GetUserIDList() []string {
	if x != nil {
		return x.UserIDList
	}
	return nil
}

func (x *GetUserInfoReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *GetUserInfoReq) GetOperationID() string {
	if x != nil {
		return x.OperationID
	}
	return ""
}

type GetUserInfoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorCode int32             `protobuf:"varint,1,opt,name=errorCode,proto3" json:"errorCode,omitempty"`
	ErrorMsg  string            `protobuf:"bytes,2,opt,name=errorMsg,proto3" json:"errorMsg,omitempty"`
	Data      []*pb_ws.UserInfo `protobuf:"bytes,3,rep,name=Data,proto3" json:"Data,omitempty"`
}

func (x *GetUserInfoResp) Reset() {
	*x = GetUserInfoResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfoResp) ProtoMessage() {}

func (x *GetUserInfoResp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfoResp.ProtoReflect.Descriptor instead.
func (*GetUserInfoResp) Descriptor() ([]byte, []int) {
	return file_pb_user_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserInfoResp) GetErrorCode() int32 {
	if x != nil {
		return x.ErrorCode
	}
	return 0
}

func (x *GetUserInfoResp) GetErrorMsg() string {
	if x != nil {
		return x.ErrorMsg
	}
	return ""
}

func (x *GetUserInfoResp) GetData() []*pb_ws.UserInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetAllUsersUidReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token       string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	OperationID string `protobuf:"bytes,3,opt,name=operationID,proto3" json:"operationID,omitempty"`
}

func (x *GetAllUsersUidReq) Reset() {
	*x = GetAllUsersUidReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllUsersUidReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllUsersUidReq) ProtoMessage() {}

func (x *GetAllUsersUidReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllUsersUidReq.ProtoReflect.Descriptor instead.
func (*GetAllUsersUidReq) Descriptor() ([]byte, []int) {
	return file_pb_user_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllUsersUidReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *GetAllUsersUidReq) GetOperationID() string {
	if x != nil {
		return x.OperationID
	}
	return ""
}

type GetAllUsersUidResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommonResp *pb_ws.CommonResp `protobuf:"bytes,1,opt,name=commonResp,proto3" json:"commonResp,omitempty"`
	UidList    []string          `protobuf:"bytes,2,rep,name=uidList,proto3" json:"uidList,omitempty"`
}

func (x *GetAllUsersUidResp) Reset() {
	*x = GetAllUsersUidResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllUsersUidResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllUsersUidResp) ProtoMessage() {}

func (x *GetAllUsersUidResp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllUsersUidResp.ProtoReflect.Descriptor instead.
func (*GetAllUsersUidResp) Descriptor() ([]byte, []int) {
	return file_pb_user_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllUsersUidResp) GetCommonResp() *pb_ws.CommonResp {
	if x != nil {
		return x.CommonResp
	}
	return nil
}

func (x *GetAllUsersUidResp) GetUidList() []string {
	if x != nil {
		return x.UidList
	}
	return nil
}

type DeleteUsersReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeleteUidList []string `protobuf:"bytes,2,rep,name=deleteUidList,proto3" json:"deleteUidList,omitempty"`
	Token         string   `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	OperationID   string   `protobuf:"bytes,4,opt,name=OperationID,proto3" json:"OperationID,omitempty"`
}

func (x *DeleteUsersReq) Reset() {
	*x = DeleteUsersReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUsersReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUsersReq) ProtoMessage() {}

func (x *DeleteUsersReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUsersReq.ProtoReflect.Descriptor instead.
func (*DeleteUsersReq) Descriptor() ([]byte, []int) {
	return file_pb_user_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteUsersReq) GetDeleteUidList() []string {
	if x != nil {
		return x.DeleteUidList
	}
	return nil
}

func (x *DeleteUsersReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *DeleteUsersReq) GetOperationID() string {
	if x != nil {
		return x.OperationID
	}
	return ""
}

type DeleteUsersResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommonResp    *pb_ws.CommonResp `protobuf:"bytes,1,opt,name=commonResp,proto3" json:"commonResp,omitempty"`
	FailedUidList []string          `protobuf:"bytes,2,rep,name=failedUidList,proto3" json:"failedUidList,omitempty"`
}

func (x *DeleteUsersResp) Reset() {
	*x = DeleteUsersResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUsersResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUsersResp) ProtoMessage() {}

func (x *DeleteUsersResp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUsersResp.ProtoReflect.Descriptor instead.
func (*DeleteUsersResp) Descriptor() ([]byte, []int) {
	return file_pb_user_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteUsersResp) GetCommonResp() *pb_ws.CommonResp {
	if x != nil {
		return x.CommonResp
	}
	return nil
}

func (x *DeleteUsersResp) GetFailedUidList() []string {
	if x != nil {
		return x.FailedUidList
	}
	return nil
}

// 离线
type LogoutReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OperationID string `protobuf:"bytes,1,opt,name=OperationID,proto3" json:"OperationID,omitempty"`
	Token       string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *LogoutReq) Reset() {
	*x = LogoutReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogoutReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutReq) ProtoMessage() {}

func (x *LogoutReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutReq.ProtoReflect.Descriptor instead.
func (*LogoutReq) Descriptor() ([]byte, []int) {
	return file_pb_user_proto_rawDescGZIP(), []int{6}
}

func (x *LogoutReq) GetOperationID() string {
	if x != nil {
		return x.OperationID
	}
	return ""
}

func (x *LogoutReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

// 更新用户信息
type UpdateUserInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid         string `protobuf:"bytes,1,opt,name=Uid,proto3" json:"Uid,omitempty"`
	Icon        string `protobuf:"bytes,2,opt,name=icon,proto3" json:"icon,omitempty"`
	Name        string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Gender      int32  `protobuf:"varint,4,opt,name=gender,proto3" json:"gender,omitempty"`
	Mobile      string `protobuf:"bytes,5,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Birth       string `protobuf:"bytes,6,opt,name=birth,proto3" json:"birth,omitempty"`
	Email       string `protobuf:"bytes,7,opt,name=email,proto3" json:"email,omitempty"`
	Ex          string `protobuf:"bytes,8,opt,name=ex,proto3" json:"ex,omitempty"`
	Token       string `protobuf:"bytes,9,opt,name=token,proto3" json:"token,omitempty"`
	OperationID string `protobuf:"bytes,10,opt,name=OperationID,proto3" json:"OperationID,omitempty"`
}

func (x *UpdateUserInfoReq) Reset() {
	*x = UpdateUserInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserInfoReq) ProtoMessage() {}

func (x *UpdateUserInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserInfoReq.ProtoReflect.Descriptor instead.
func (*UpdateUserInfoReq) Descriptor() ([]byte, []int) {
	return file_pb_user_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateUserInfoReq) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *UpdateUserInfoReq) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *UpdateUserInfoReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateUserInfoReq) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *UpdateUserInfoReq) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *UpdateUserInfoReq) GetBirth() string {
	if x != nil {
		return x.Birth
	}
	return ""
}

func (x *UpdateUserInfoReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UpdateUserInfoReq) GetEx() string {
	if x != nil {
		return x.Ex
	}
	return ""
}

func (x *UpdateUserInfoReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *UpdateUserInfoReq) GetOperationID() string {
	if x != nil {
		return x.OperationID
	}
	return ""
}

var File_pb_user_proto protoreflect.FileDescriptor

var file_pb_user_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x62, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0b, 0x70, 0x62, 0x5f, 0x77, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x68, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x1e,
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x22, 0x7c, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x4d, 0x73, 0x67, 0x12, 0x2f, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04,
	0x44, 0x61, 0x74, 0x61, 0x22, 0x4b, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x55, 0x69, 0x64, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x20, 0x0a, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x44, 0x22, 0x6d, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x73,
	0x55, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x69, 0x64, 0x4c, 0x69, 0x73,
	0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x75, 0x69, 0x64, 0x4c, 0x69, 0x73, 0x74,
	0x22, 0x6e, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x71, 0x12, 0x24, 0x0a, 0x0d, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x69, 0x64, 0x4c,
	0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x55, 0x69, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x20,
	0x0a, 0x0b, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44,
	0x22, 0x76, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x3d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x5f, 0x61, 0x70, 0x69, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x24, 0x0a, 0x0d, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x55, 0x69, 0x64, 0x4c,
	0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x66, 0x61, 0x69, 0x6c, 0x65,
	0x64, 0x55, 0x69, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x43, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x6f,
	0x75, 0x74, 0x52, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x0b, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xf1, 0x01,
	0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x55, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x67,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x62, 0x69, 0x72, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x69,
	0x72, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x65, 0x78, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x65, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x20, 0x0a, 0x0b, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x44, 0x32, 0xea, 0x01, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0f, 0x2e, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x12, 0x39, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x73, 0x55, 0x69, 0x64, 0x12, 0x12,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x73, 0x55, 0x69, 0x64, 0x52,
	0x65, 0x71, 0x1a, 0x13, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x73,
	0x55, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x43, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x30, 0x0a, 0x0b,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x0f, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x42, 0x24,
	0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x64, 0x6e,
	0x70, 0x61, 0x2f, 0x49, 0x4d, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x62, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_user_proto_rawDescOnce sync.Once
	file_pb_user_proto_rawDescData = file_pb_user_proto_rawDesc
)

func file_pb_user_proto_rawDescGZIP() []byte {
	file_pb_user_proto_rawDescOnce.Do(func() {
		file_pb_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_user_proto_rawDescData)
	})
	return file_pb_user_proto_rawDescData
}

var file_pb_user_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_pb_user_proto_goTypes = []interface{}{
	(*GetUserInfoReq)(nil),     // 0: GetUserInfoReq
	(*GetUserInfoResp)(nil),    // 1: GetUserInfoResp
	(*GetAllUsersUidReq)(nil),  // 2: GetAllUsersUidReq
	(*GetAllUsersUidResp)(nil), // 3: GetAllUsersUidResp
	(*DeleteUsersReq)(nil),     // 4: DeleteUsersReq
	(*DeleteUsersResp)(nil),    // 5: DeleteUsersResp
	(*LogoutReq)(nil),          // 6: LogoutReq
	(*UpdateUserInfoReq)(nil),  // 7: UpdateUserInfoReq
	(*pb_ws.UserInfo)(nil),     // 8: server_api_params.UserInfo
	(*pb_ws.CommonResp)(nil),   // 9: server_api_params.CommonResp
}
var file_pb_user_proto_depIdxs = []int32{
	8, // 0: GetUserInfoResp.Data:type_name -> server_api_params.UserInfo
	9, // 1: GetAllUsersUidResp.commonResp:type_name -> server_api_params.CommonResp
	9, // 2: DeleteUsersResp.commonResp:type_name -> server_api_params.CommonResp
	0, // 3: user.GetUserInfo:input_type -> GetUserInfoReq
	2, // 4: user.GetAllUsersUid:input_type -> GetAllUsersUidReq
	7, // 5: user.UpdateUserInfo:input_type -> UpdateUserInfoReq
	4, // 6: user.DeleteUsers:input_type -> DeleteUsersReq
	1, // 7: user.GetUserInfo:output_type -> GetUserInfoResp
	3, // 8: user.GetAllUsersUid:output_type -> GetAllUsersUidResp
	9, // 9: user.UpdateUserInfo:output_type -> server_api_params.CommonResp
	5, // 10: user.DeleteUsers:output_type -> DeleteUsersResp
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pb_user_proto_init() }
func file_pb_user_proto_init() {
	if File_pb_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserInfoReq); i {
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
		file_pb_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserInfoResp); i {
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
		file_pb_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllUsersUidReq); i {
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
		file_pb_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllUsersUidResp); i {
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
		file_pb_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUsersReq); i {
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
		file_pb_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUsersResp); i {
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
		file_pb_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogoutReq); i {
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
		file_pb_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserInfoReq); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_user_proto_goTypes,
		DependencyIndexes: file_pb_user_proto_depIdxs,
		MessageInfos:      file_pb_user_proto_msgTypes,
	}.Build()
	File_pb_user_proto = out.File
	file_pb_user_proto_rawDesc = nil
	file_pb_user_proto_goTypes = nil
	file_pb_user_proto_depIdxs = nil
}
