// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.30.1
// source: oss.proto

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
	OSS_Upload_FullMethodName   = "/OSS/Upload"
	OSS_Download_FullMethodName = "/OSS/Download"
)

// OSSClient is the client API for OSS service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OSSClient interface {
	Upload(ctx context.Context, in *UploadReq, opts ...grpc.CallOption) (*UploadResp, error)
	Download(ctx context.Context, in *DownloadReq, opts ...grpc.CallOption) (*DownloadResp, error)
}

type oSSClient struct {
	cc grpc.ClientConnInterface
}

func NewOSSClient(cc grpc.ClientConnInterface) OSSClient {
	return &oSSClient{cc}
}

func (c *oSSClient) Upload(ctx context.Context, in *UploadReq, opts ...grpc.CallOption) (*UploadResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UploadResp)
	err := c.cc.Invoke(ctx, OSS_Upload_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oSSClient) Download(ctx context.Context, in *DownloadReq, opts ...grpc.CallOption) (*DownloadResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DownloadResp)
	err := c.cc.Invoke(ctx, OSS_Download_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OSSServer is the server API for OSS service.
// All implementations must embed UnimplementedOSSServer
// for forward compatibility.
type OSSServer interface {
	Upload(context.Context, *UploadReq) (*UploadResp, error)
	Download(context.Context, *DownloadReq) (*DownloadResp, error)
	mustEmbedUnimplementedOSSServer()
}

// UnimplementedOSSServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedOSSServer struct{}

func (UnimplementedOSSServer) Upload(context.Context, *UploadReq) (*UploadResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (UnimplementedOSSServer) Download(context.Context, *DownloadReq) (*DownloadResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Download not implemented")
}
func (UnimplementedOSSServer) mustEmbedUnimplementedOSSServer() {}
func (UnimplementedOSSServer) testEmbeddedByValue()             {}

// UnsafeOSSServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OSSServer will
// result in compilation errors.
type UnsafeOSSServer interface {
	mustEmbedUnimplementedOSSServer()
}

func RegisterOSSServer(s grpc.ServiceRegistrar, srv OSSServer) {
	// If the following call pancis, it indicates UnimplementedOSSServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&OSS_ServiceDesc, srv)
}

func _OSS_Upload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OSSServer).Upload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OSS_Upload_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OSSServer).Upload(ctx, req.(*UploadReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OSS_Download_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OSSServer).Download(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OSS_Download_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OSSServer).Download(ctx, req.(*DownloadReq))
	}
	return interceptor(ctx, in, info, handler)
}

// OSS_ServiceDesc is the grpc.ServiceDesc for OSS service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OSS_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "OSS",
	HandlerType: (*OSSServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Upload",
			Handler:    _OSS_Upload_Handler,
		},
		{
			MethodName: "Download",
			Handler:    _OSS_Download_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "oss.proto",
}
