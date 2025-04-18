// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.30.1
// source: group.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Group_GetUserGroups_FullMethodName          = "/Group/GetUserGroups"
	Group_GetGroupInfoById_FullMethodName       = "/Group/GetGroupInfoById"
	Group_CreateGroupInfo_FullMethodName        = "/Group/CreateGroupInfo"
	Group_UpdateGroupInfo_FullMethodName        = "/Group/UpdateGroupInfo"
	Group_DeleteGroupInfo_FullMethodName        = "/Group/DeleteGroupInfo"
	Group_GetGroupMemberById_FullMethodName     = "/Group/GetGroupMemberById"
	Group_CreateGroupMember_FullMethodName      = "/Group/CreateGroupMember"
	Group_UpdateGroupMember_FullMethodName      = "/Group/UpdateGroupMember"
	Group_DeleteGroupMember_FullMethodName      = "/Group/DeleteGroupMember"
	Group_GetGroupApplyByGroupId_FullMethodName = "/Group/GetGroupApplyByGroupId"
	Group_GetGroupApplyByUserId_FullMethodName  = "/Group/GetGroupApplyByUserId"
	Group_CreateGroupApply_FullMethodName       = "/Group/CreateGroupApply"
	Group_UpdateGroupApply_FullMethodName       = "/Group/UpdateGroupApply"
	Group_DeleteGroupApply_FullMethodName       = "/Group/DeleteGroupApply"
)

// GroupClient is the client API for Group service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GroupClient interface {
	GetUserGroups(ctx context.Context, in *GetUserGroupsReq, opts ...grpc.CallOption) (*GetUserGroupsResp, error)
	// 群聊基础信息管理
	GetGroupInfoById(ctx context.Context, in *GetGroupInfoByIdReq, opts ...grpc.CallOption) (*GetGroupInfoByIdResp, error)
	CreateGroupInfo(ctx context.Context, in *CreateGroupInfoReq, opts ...grpc.CallOption) (*CreateGroupInfoResp, error)
	UpdateGroupInfo(ctx context.Context, in *UpdateGroupInfoReq, opts ...grpc.CallOption) (*UpdateGroupInfoResp, error)
	DeleteGroupInfo(ctx context.Context, in *DeleteGroupInfoReq, opts ...grpc.CallOption) (*DeleteGroupInfoResp, error)
	// 成员管理
	GetGroupMemberById(ctx context.Context, in *GetGroupMemberByIdReq, opts ...grpc.CallOption) (*GetGroupMemberByIdResp, error)
	CreateGroupMember(ctx context.Context, in *CreateGroupMemberReq, opts ...grpc.CallOption) (*CreateGroupMemberResp, error)
	UpdateGroupMember(ctx context.Context, in *UpdateGroupMemberReq, opts ...grpc.CallOption) (*UpdateGroupMemberResp, error)
	DeleteGroupMember(ctx context.Context, in *DeleteGroupMemberReq, opts ...grpc.CallOption) (*DeleteGroupMemberResp, error)
	// 申请管理
	GetGroupApplyByGroupId(ctx context.Context, in *GetGroupApplyByGroupIdReq, opts ...grpc.CallOption) (*GetGroupApplyByGroupIdResp, error)
	GetGroupApplyByUserId(ctx context.Context, in *GetGroupApplyByUserIdReq, opts ...grpc.CallOption) (*GetGroupApplyByUserIdResp, error)
	CreateGroupApply(ctx context.Context, in *CreateGroupApplyReq, opts ...grpc.CallOption) (*CreateGroupApplyResp, error)
	UpdateGroupApply(ctx context.Context, in *UpdateGroupApplyReq, opts ...grpc.CallOption) (*UpdateGroupApplyResp, error)
	DeleteGroupApply(ctx context.Context, in *DeleteGroupApplyReq, opts ...grpc.CallOption) (*DeleteGroupApplyResp, error)
}

type groupClient struct {
	cc grpc.ClientConnInterface
}

func NewGroupClient(cc grpc.ClientConnInterface) GroupClient {
	return &groupClient{cc}
}

func (c *groupClient) GetUserGroups(ctx context.Context, in *GetUserGroupsReq, opts ...grpc.CallOption) (*GetUserGroupsResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserGroupsResp)
	err := c.cc.Invoke(ctx, Group_GetUserGroups_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) GetGroupInfoById(ctx context.Context, in *GetGroupInfoByIdReq, opts ...grpc.CallOption) (*GetGroupInfoByIdResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetGroupInfoByIdResp)
	err := c.cc.Invoke(ctx, Group_GetGroupInfoById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) CreateGroupInfo(ctx context.Context, in *CreateGroupInfoReq, opts ...grpc.CallOption) (*CreateGroupInfoResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateGroupInfoResp)
	err := c.cc.Invoke(ctx, Group_CreateGroupInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) UpdateGroupInfo(ctx context.Context, in *UpdateGroupInfoReq, opts ...grpc.CallOption) (*UpdateGroupInfoResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateGroupInfoResp)
	err := c.cc.Invoke(ctx, Group_UpdateGroupInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) DeleteGroupInfo(ctx context.Context, in *DeleteGroupInfoReq, opts ...grpc.CallOption) (*DeleteGroupInfoResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteGroupInfoResp)
	err := c.cc.Invoke(ctx, Group_DeleteGroupInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) GetGroupMemberById(ctx context.Context, in *GetGroupMemberByIdReq, opts ...grpc.CallOption) (*GetGroupMemberByIdResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetGroupMemberByIdResp)
	err := c.cc.Invoke(ctx, Group_GetGroupMemberById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) CreateGroupMember(ctx context.Context, in *CreateGroupMemberReq, opts ...grpc.CallOption) (*CreateGroupMemberResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateGroupMemberResp)
	err := c.cc.Invoke(ctx, Group_CreateGroupMember_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) UpdateGroupMember(ctx context.Context, in *UpdateGroupMemberReq, opts ...grpc.CallOption) (*UpdateGroupMemberResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateGroupMemberResp)
	err := c.cc.Invoke(ctx, Group_UpdateGroupMember_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) DeleteGroupMember(ctx context.Context, in *DeleteGroupMemberReq, opts ...grpc.CallOption) (*DeleteGroupMemberResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteGroupMemberResp)
	err := c.cc.Invoke(ctx, Group_DeleteGroupMember_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) GetGroupApplyByGroupId(ctx context.Context, in *GetGroupApplyByGroupIdReq, opts ...grpc.CallOption) (*GetGroupApplyByGroupIdResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetGroupApplyByGroupIdResp)
	err := c.cc.Invoke(ctx, Group_GetGroupApplyByGroupId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) GetGroupApplyByUserId(ctx context.Context, in *GetGroupApplyByUserIdReq, opts ...grpc.CallOption) (*GetGroupApplyByUserIdResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetGroupApplyByUserIdResp)
	err := c.cc.Invoke(ctx, Group_GetGroupApplyByUserId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) CreateGroupApply(ctx context.Context, in *CreateGroupApplyReq, opts ...grpc.CallOption) (*CreateGroupApplyResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateGroupApplyResp)
	err := c.cc.Invoke(ctx, Group_CreateGroupApply_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) UpdateGroupApply(ctx context.Context, in *UpdateGroupApplyReq, opts ...grpc.CallOption) (*UpdateGroupApplyResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateGroupApplyResp)
	err := c.cc.Invoke(ctx, Group_UpdateGroupApply_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) DeleteGroupApply(ctx context.Context, in *DeleteGroupApplyReq, opts ...grpc.CallOption) (*DeleteGroupApplyResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteGroupApplyResp)
	err := c.cc.Invoke(ctx, Group_DeleteGroupApply_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GroupServer is the server API for Group service.
// All implementations must embed UnimplementedGroupServer
// for forward compatibility.
type GroupServer interface {
	GetUserGroups(context.Context, *GetUserGroupsReq) (*GetUserGroupsResp, error)
	// 群聊基础信息管理
	GetGroupInfoById(context.Context, *GetGroupInfoByIdReq) (*GetGroupInfoByIdResp, error)
	CreateGroupInfo(context.Context, *CreateGroupInfoReq) (*CreateGroupInfoResp, error)
	UpdateGroupInfo(context.Context, *UpdateGroupInfoReq) (*UpdateGroupInfoResp, error)
	DeleteGroupInfo(context.Context, *DeleteGroupInfoReq) (*DeleteGroupInfoResp, error)
	// 成员管理
	GetGroupMemberById(context.Context, *GetGroupMemberByIdReq) (*GetGroupMemberByIdResp, error)
	CreateGroupMember(context.Context, *CreateGroupMemberReq) (*CreateGroupMemberResp, error)
	UpdateGroupMember(context.Context, *UpdateGroupMemberReq) (*UpdateGroupMemberResp, error)
	DeleteGroupMember(context.Context, *DeleteGroupMemberReq) (*DeleteGroupMemberResp, error)
	// 申请管理
	GetGroupApplyByGroupId(context.Context, *GetGroupApplyByGroupIdReq) (*GetGroupApplyByGroupIdResp, error)
	GetGroupApplyByUserId(context.Context, *GetGroupApplyByUserIdReq) (*GetGroupApplyByUserIdResp, error)
	CreateGroupApply(context.Context, *CreateGroupApplyReq) (*CreateGroupApplyResp, error)
	UpdateGroupApply(context.Context, *UpdateGroupApplyReq) (*UpdateGroupApplyResp, error)
	DeleteGroupApply(context.Context, *DeleteGroupApplyReq) (*DeleteGroupApplyResp, error)
	mustEmbedUnimplementedGroupServer()
}

// UnimplementedGroupServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGroupServer struct{}

func (UnimplementedGroupServer) GetUserGroups(context.Context, *GetUserGroupsReq) (*GetUserGroupsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserGroups not implemented")
}
func (UnimplementedGroupServer) GetGroupInfoById(context.Context, *GetGroupInfoByIdReq) (*GetGroupInfoByIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroupInfoById not implemented")
}
func (UnimplementedGroupServer) CreateGroupInfo(context.Context, *CreateGroupInfoReq) (*CreateGroupInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroupInfo not implemented")
}
func (UnimplementedGroupServer) UpdateGroupInfo(context.Context, *UpdateGroupInfoReq) (*UpdateGroupInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGroupInfo not implemented")
}
func (UnimplementedGroupServer) DeleteGroupInfo(context.Context, *DeleteGroupInfoReq) (*DeleteGroupInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGroupInfo not implemented")
}
func (UnimplementedGroupServer) GetGroupMemberById(context.Context, *GetGroupMemberByIdReq) (*GetGroupMemberByIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroupMemberById not implemented")
}
func (UnimplementedGroupServer) CreateGroupMember(context.Context, *CreateGroupMemberReq) (*CreateGroupMemberResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroupMember not implemented")
}
func (UnimplementedGroupServer) UpdateGroupMember(context.Context, *UpdateGroupMemberReq) (*UpdateGroupMemberResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGroupMember not implemented")
}
func (UnimplementedGroupServer) DeleteGroupMember(context.Context, *DeleteGroupMemberReq) (*DeleteGroupMemberResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGroupMember not implemented")
}
func (UnimplementedGroupServer) GetGroupApplyByGroupId(context.Context, *GetGroupApplyByGroupIdReq) (*GetGroupApplyByGroupIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroupApplyByGroupId not implemented")
}
func (UnimplementedGroupServer) GetGroupApplyByUserId(context.Context, *GetGroupApplyByUserIdReq) (*GetGroupApplyByUserIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroupApplyByUserId not implemented")
}
func (UnimplementedGroupServer) CreateGroupApply(context.Context, *CreateGroupApplyReq) (*CreateGroupApplyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroupApply not implemented")
}
func (UnimplementedGroupServer) UpdateGroupApply(context.Context, *UpdateGroupApplyReq) (*UpdateGroupApplyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGroupApply not implemented")
}
func (UnimplementedGroupServer) DeleteGroupApply(context.Context, *DeleteGroupApplyReq) (*DeleteGroupApplyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGroupApply not implemented")
}
func (UnimplementedGroupServer) mustEmbedUnimplementedGroupServer() {}
func (UnimplementedGroupServer) testEmbeddedByValue()               {}

// UnsafeGroupServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GroupServer will
// result in compilation errors.
type UnsafeGroupServer interface {
	mustEmbedUnimplementedGroupServer()
}

func RegisterGroupServer(s grpc.ServiceRegistrar, srv GroupServer) {
	// If the following call pancis, it indicates UnimplementedGroupServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Group_ServiceDesc, srv)
}

func _Group_GetUserGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserGroupsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).GetUserGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Group_GetUserGroups_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).GetUserGroups(ctx, req.(*GetUserGroupsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_GetGroupInfoById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupInfoByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).GetGroupInfoById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Group_GetGroupInfoById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).GetGroupInfoById(ctx, req.(*GetGroupInfoByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_CreateGroupInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).CreateGroupInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Group_CreateGroupInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).CreateGroupInfo(ctx, req.(*CreateGroupInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_UpdateGroupInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGroupInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).UpdateGroupInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Group_UpdateGroupInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).UpdateGroupInfo(ctx, req.(*UpdateGroupInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_DeleteGroupInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGroupInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).DeleteGroupInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Group_DeleteGroupInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).DeleteGroupInfo(ctx, req.(*DeleteGroupInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_GetGroupMemberById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupMemberByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).GetGroupMemberById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Group_GetGroupMemberById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).GetGroupMemberById(ctx, req.(*GetGroupMemberByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_CreateGroupMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupMemberReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).CreateGroupMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Group_CreateGroupMember_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).CreateGroupMember(ctx, req.(*CreateGroupMemberReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_UpdateGroupMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGroupMemberReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).UpdateGroupMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Group_UpdateGroupMember_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).UpdateGroupMember(ctx, req.(*UpdateGroupMemberReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_DeleteGroupMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGroupMemberReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).DeleteGroupMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Group_DeleteGroupMember_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).DeleteGroupMember(ctx, req.(*DeleteGroupMemberReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_GetGroupApplyByGroupId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupApplyByGroupIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).GetGroupApplyByGroupId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Group_GetGroupApplyByGroupId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).GetGroupApplyByGroupId(ctx, req.(*GetGroupApplyByGroupIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_GetGroupApplyByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupApplyByUserIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).GetGroupApplyByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Group_GetGroupApplyByUserId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).GetGroupApplyByUserId(ctx, req.(*GetGroupApplyByUserIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_CreateGroupApply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupApplyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).CreateGroupApply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Group_CreateGroupApply_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).CreateGroupApply(ctx, req.(*CreateGroupApplyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_UpdateGroupApply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGroupApplyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).UpdateGroupApply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Group_UpdateGroupApply_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).UpdateGroupApply(ctx, req.(*UpdateGroupApplyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_DeleteGroupApply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGroupApplyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).DeleteGroupApply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Group_DeleteGroupApply_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).DeleteGroupApply(ctx, req.(*DeleteGroupApplyReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Group_ServiceDesc is the grpc.ServiceDesc for Group service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Group_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Group",
	HandlerType: (*GroupServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserGroups",
			Handler:    _Group_GetUserGroups_Handler,
		},
		{
			MethodName: "GetGroupInfoById",
			Handler:    _Group_GetGroupInfoById_Handler,
		},
		{
			MethodName: "CreateGroupInfo",
			Handler:    _Group_CreateGroupInfo_Handler,
		},
		{
			MethodName: "UpdateGroupInfo",
			Handler:    _Group_UpdateGroupInfo_Handler,
		},
		{
			MethodName: "DeleteGroupInfo",
			Handler:    _Group_DeleteGroupInfo_Handler,
		},
		{
			MethodName: "GetGroupMemberById",
			Handler:    _Group_GetGroupMemberById_Handler,
		},
		{
			MethodName: "CreateGroupMember",
			Handler:    _Group_CreateGroupMember_Handler,
		},
		{
			MethodName: "UpdateGroupMember",
			Handler:    _Group_UpdateGroupMember_Handler,
		},
		{
			MethodName: "DeleteGroupMember",
			Handler:    _Group_DeleteGroupMember_Handler,
		},
		{
			MethodName: "GetGroupApplyByGroupId",
			Handler:    _Group_GetGroupApplyByGroupId_Handler,
		},
		{
			MethodName: "GetGroupApplyByUserId",
			Handler:    _Group_GetGroupApplyByUserId_Handler,
		},
		{
			MethodName: "CreateGroupApply",
			Handler:    _Group_CreateGroupApply_Handler,
		},
		{
			MethodName: "UpdateGroupApply",
			Handler:    _Group_UpdateGroupApply_Handler,
		},
		{
			MethodName: "DeleteGroupApply",
			Handler:    _Group_DeleteGroupApply_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "group.proto",
}
