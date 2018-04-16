// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protocol.proto

/*
Package protocol is a generated protocol buffer package.

It is generated from these files:
	protocol.proto

It has these top-level messages:
	Task
	Tasks
	Void
	ID
*/
package protocol

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

type Task struct {
	Id   string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *Task) Reset()                    { *m = Task{} }
func (m *Task) String() string            { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()               {}
func (*Task) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Task) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Task) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Tasks struct {
	Task []*Task `protobuf:"bytes,1,rep,name=task" json:"task,omitempty"`
}

func (m *Tasks) Reset()                    { *m = Tasks{} }
func (m *Tasks) String() string            { return proto.CompactTextString(m) }
func (*Tasks) ProtoMessage()               {}
func (*Tasks) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Tasks) GetTask() []*Task {
	if m != nil {
		return m.Task
	}
	return nil
}

type Void struct {
}

func (m *Void) Reset()                    { *m = Void{} }
func (m *Void) String() string            { return proto.CompactTextString(m) }
func (*Void) ProtoMessage()               {}
func (*Void) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type ID struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *ID) Reset()                    { *m = ID{} }
func (m *ID) String() string            { return proto.CompactTextString(m) }
func (*ID) ProtoMessage()               {}
func (*ID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ID) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*Task)(nil), "protocol.Task")
	proto.RegisterType((*Tasks)(nil), "protocol.Tasks")
	proto.RegisterType((*Void)(nil), "protocol.Void")
	proto.RegisterType((*ID)(nil), "protocol.ID")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Report service

type ReportClient interface {
	Create(ctx context.Context, in *Task, opts ...grpc.CallOption) (*ID, error)
	FindOne(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Task, error)
	FindAll(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Tasks, error)
	Update(ctx context.Context, in *Task, opts ...grpc.CallOption) (*ID, error)
	Delete(ctx context.Context, in *ID, opts ...grpc.CallOption) (*ID, error)
}

type reportClient struct {
	cc *grpc.ClientConn
}

func NewReportClient(cc *grpc.ClientConn) ReportClient {
	return &reportClient{cc}
}

func (c *reportClient) Create(ctx context.Context, in *Task, opts ...grpc.CallOption) (*ID, error) {
	out := new(ID)
	err := grpc.Invoke(ctx, "/protocol.Report/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportClient) FindOne(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := grpc.Invoke(ctx, "/protocol.Report/FindOne", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportClient) FindAll(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Tasks, error) {
	out := new(Tasks)
	err := grpc.Invoke(ctx, "/protocol.Report/FindAll", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportClient) Update(ctx context.Context, in *Task, opts ...grpc.CallOption) (*ID, error) {
	out := new(ID)
	err := grpc.Invoke(ctx, "/protocol.Report/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportClient) Delete(ctx context.Context, in *ID, opts ...grpc.CallOption) (*ID, error) {
	out := new(ID)
	err := grpc.Invoke(ctx, "/protocol.Report/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Report service

type ReportServer interface {
	Create(context.Context, *Task) (*ID, error)
	FindOne(context.Context, *ID) (*Task, error)
	FindAll(context.Context, *Void) (*Tasks, error)
	Update(context.Context, *Task) (*ID, error)
	Delete(context.Context, *ID) (*ID, error)
}

func RegisterReportServer(s *grpc.Server, srv ReportServer) {
	s.RegisterService(&_Report_serviceDesc, srv)
}

func _Report_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Task)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Report/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportServer).Create(ctx, req.(*Task))
	}
	return interceptor(ctx, in, info, handler)
}

func _Report_FindOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportServer).FindOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Report/FindOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportServer).FindOne(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Report_FindAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportServer).FindAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Report/FindAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportServer).FindAll(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _Report_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Task)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Report/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportServer).Update(ctx, req.(*Task))
	}
	return interceptor(ctx, in, info, handler)
}

func _Report_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Report/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportServer).Delete(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

var _Report_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.Report",
	HandlerType: (*ReportServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Report_Create_Handler,
		},
		{
			MethodName: "FindOne",
			Handler:    _Report_FindOne_Handler,
		},
		{
			MethodName: "FindAll",
			Handler:    _Report_FindAll_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Report_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Report_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protocol.proto",
}

func init() { proto.RegisterFile("protocol.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x8e, 0xcd, 0x4a, 0xc4, 0x30,
	0x14, 0x85, 0x9b, 0x18, 0xa3, 0x5e, 0xa5, 0xc2, 0xc5, 0x45, 0xe9, 0x6a, 0xb8, 0x0b, 0xa9, 0x3f,
	0xcc, 0x62, 0x7c, 0x02, 0xb1, 0x08, 0xb3, 0x12, 0x8a, 0xba, 0x8f, 0x26, 0x8b, 0x30, 0xb1, 0x29,
	0x4d, 0x5e, 0xd7, 0x77, 0x91, 0x44, 0xfc, 0x69, 0x56, 0xb3, 0xbb, 0x27, 0xe7, 0xe4, 0xe3, 0x83,
	0x7a, 0x9a, 0x7d, 0xf4, 0xef, 0xde, 0xad, 0xf3, 0x81, 0xc7, 0x3f, 0x99, 0xae, 0x41, 0x3c, 0xab,
	0xb0, 0xc3, 0x1a, 0xb8, 0xd5, 0x0d, 0x5b, 0xb1, 0xee, 0x64, 0xe0, 0x56, 0x23, 0x82, 0x18, 0xd5,
	0x87, 0x69, 0x78, 0x7e, 0xc9, 0x37, 0xdd, 0xc0, 0x61, 0xda, 0x06, 0x24, 0x10, 0x51, 0x85, 0x5d,
	0xc3, 0x56, 0x07, 0xdd, 0xe9, 0xa6, 0x5e, 0xff, 0xd2, 0x53, 0x3d, 0xe4, 0x8e, 0x24, 0x88, 0x57,
	0x6f, 0x35, 0x5d, 0x00, 0xdf, 0xf6, 0x25, 0x7e, 0xf3, 0xc9, 0x40, 0x0e, 0x66, 0xf2, 0x73, 0xc4,
	0x0e, 0xe4, 0xc3, 0x6c, 0x54, 0x34, 0x58, 0x80, 0xda, 0xb3, 0xbf, 0xbc, 0xed, 0xa9, 0xc2, 0x2b,
	0x38, 0x7a, 0xb4, 0xa3, 0x7e, 0x1a, 0x0d, 0x2e, 0xaa, 0xb6, 0xf8, 0x48, 0x15, 0xde, 0x7e, 0x4f,
	0xef, 0x9d, 0xfb, 0x4f, 0x4d, 0x42, 0xed, 0xf9, 0x72, 0x1c, 0xa8, 0x4a, 0x0a, 0x2f, 0x93, 0xde,
	0x47, 0xe1, 0x12, 0x64, 0x6f, 0x9c, 0x89, 0xa5, 0x41, 0xb1, 0x7b, 0x93, 0x39, 0xde, 0x7d, 0x05,
	0x00, 0x00, 0xff, 0xff, 0x4a, 0x58, 0xc1, 0xc5, 0x79, 0x01, 0x00, 0x00,
}
