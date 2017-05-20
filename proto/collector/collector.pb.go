// Code generated by protoc-gen-go. DO NOT EDIT.
// source: collector.proto

/*
Package collector is a generated protocol buffer package.

It is generated from these files:
	collector.proto

It has these top-level messages:
	Empty
	Repositories
	Repository
*/
package collector

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Repositories struct {
	Repos []*Repository `protobuf:"bytes,1,rep,name=repos" json:"repos,omitempty"`
}

func (m *Repositories) Reset()                    { *m = Repositories{} }
func (m *Repositories) String() string            { return proto.CompactTextString(m) }
func (*Repositories) ProtoMessage()               {}
func (*Repositories) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Repositories) GetRepos() []*Repository {
	if m != nil {
		return m.Repos
	}
	return nil
}

type Repository struct {
	Id        int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Url       string `protobuf:"bytes,3,opt,name=url" json:"url,omitempty"`
	Html      string `protobuf:"bytes,4,opt,name=html" json:"html,omitempty"`
	Desc      string `protobuf:"bytes,5,opt,name=desc" json:"desc,omitempty"`
	StarCount int32  `protobuf:"varint,6,opt,name=star_count,json=starCount" json:"star_count,omitempty"`
	Git       string `protobuf:"bytes,7,opt,name=git" json:"git,omitempty"`
	Clone     string `protobuf:"bytes,8,opt,name=clone" json:"clone,omitempty"`
}

func (m *Repository) Reset()                    { *m = Repository{} }
func (m *Repository) String() string            { return proto.CompactTextString(m) }
func (*Repository) ProtoMessage()               {}
func (*Repository) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Repository) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Repository) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Repository) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Repository) GetHtml() string {
	if m != nil {
		return m.Html
	}
	return ""
}

func (m *Repository) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *Repository) GetStarCount() int32 {
	if m != nil {
		return m.StarCount
	}
	return 0
}

func (m *Repository) GetGit() string {
	if m != nil {
		return m.Git
	}
	return ""
}

func (m *Repository) GetClone() string {
	if m != nil {
		return m.Clone
	}
	return ""
}

func init() {
	proto.RegisterType((*Empty)(nil), "collector.Empty")
	proto.RegisterType((*Repositories)(nil), "collector.Repositories")
	proto.RegisterType((*Repository)(nil), "collector.Repository")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Collector service

type CollectorClient interface {
	GetStarredRepositories(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Repositories, error)
}

type collectorClient struct {
	cc *grpc.ClientConn
}

func NewCollectorClient(cc *grpc.ClientConn) CollectorClient {
	return &collectorClient{cc}
}

func (c *collectorClient) GetStarredRepositories(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Repositories, error) {
	out := new(Repositories)
	err := grpc.Invoke(ctx, "/collector.Collector/GetStarredRepositories", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Collector service

type CollectorServer interface {
	GetStarredRepositories(context.Context, *Empty) (*Repositories, error)
}

func RegisterCollectorServer(s *grpc.Server, srv CollectorServer) {
	s.RegisterService(&_Collector_serviceDesc, srv)
}

func _Collector_GetStarredRepositories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectorServer).GetStarredRepositories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/collector.Collector/GetStarredRepositories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectorServer).GetStarredRepositories(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Collector_serviceDesc = grpc.ServiceDesc{
	ServiceName: "collector.Collector",
	HandlerType: (*CollectorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStarredRepositories",
			Handler:    _Collector_GetStarredRepositories_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "collector.proto",
}

func init() { proto.RegisterFile("collector.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x4f, 0x4b, 0x03, 0x31,
	0x10, 0xc5, 0xcd, 0x6e, 0xd3, 0xba, 0xa3, 0x68, 0x09, 0xfe, 0x09, 0x82, 0xb0, 0xec, 0x69, 0x41,
	0xe8, 0xa1, 0x1e, 0x3d, 0x96, 0xe2, 0x3d, 0x7e, 0x00, 0x59, 0xb3, 0xa1, 0x06, 0xb2, 0x9b, 0x25,
	0x99, 0x1e, 0xfa, 0xb1, 0xfc, 0x86, 0x32, 0x09, 0xb6, 0x0a, 0xde, 0xde, 0xfc, 0x78, 0xbc, 0xe1,
	0x3d, 0xb8, 0xd6, 0xde, 0x39, 0xa3, 0xd1, 0x87, 0xd5, 0x14, 0x3c, 0x7a, 0x51, 0x1d, 0x41, 0xb3,
	0x00, 0xbe, 0x1d, 0x26, 0x3c, 0x34, 0x2f, 0x70, 0xa9, 0xcc, 0xe4, 0xa3, 0x45, 0x1f, 0xac, 0x89,
	0xe2, 0x09, 0x78, 0xa0, 0x5b, 0xb2, 0xba, 0x6c, 0x2f, 0xd6, 0xb7, 0xab, 0x53, 0xc8, 0xd1, 0x77,
	0x50, 0xd9, 0xd3, 0x7c, 0x31, 0x80, 0x13, 0x15, 0x57, 0x50, 0xd8, 0x5e, 0xb2, 0x9a, 0xb5, 0x5c,
	0x15, 0xb6, 0x17, 0x02, 0x66, 0x63, 0x37, 0x18, 0x59, 0xd4, 0xac, 0xad, 0x54, 0xd2, 0x62, 0x09,
	0xe5, 0x3e, 0x38, 0x59, 0x26, 0x44, 0x92, 0x5c, 0x9f, 0x38, 0x38, 0x39, 0xcb, 0x2e, 0xd2, 0xc4,
	0x7a, 0x13, 0xb5, 0xe4, 0x99, 0x91, 0x16, 0x8f, 0x00, 0x11, 0xbb, 0xf0, 0xae, 0xfd, 0x7e, 0x44,
	0x39, 0x4f, 0x5f, 0x2a, 0x22, 0x1b, 0x02, 0x14, 0xbc, 0xb3, 0x28, 0x17, 0x39, 0x78, 0x67, 0x51,
	0xdc, 0x00, 0xd7, 0xce, 0x8f, 0x46, 0x9e, 0x27, 0x96, 0x8f, 0xb5, 0x82, 0x6a, 0xf3, 0x53, 0x49,
	0x6c, 0xe1, 0xee, 0xd5, 0xe0, 0x1b, 0x76, 0x21, 0x98, 0xfe, 0xcf, 0x0e, 0xcb, 0x5f, 0xc5, 0xd3,
	0x52, 0x0f, 0xf7, 0xff, 0x4d, 0x61, 0x4d, 0x6c, 0xce, 0x3e, 0xe6, 0x69, 0xdf, 0xe7, 0xef, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xe8, 0xc8, 0x7b, 0x7f, 0x72, 0x01, 0x00, 0x00,
}
