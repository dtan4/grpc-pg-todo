// Code generated by protoc-gen-go.
// source: todo.proto
// DO NOT EDIT!

/*
Package todo is a generated protocol buffer package.

It is generated from these files:
	todo.proto

It has these top-level messages:
	ListRequest
	Task
	Tasks
*/
package todo

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
const _ = proto.ProtoPackageIsVersion1

type ListRequest struct {
	Title string `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Task struct {
	Title       string `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
}

func (m *Task) Reset()                    { *m = Task{} }
func (m *Task) String() string            { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()               {}
func (*Task) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Tasks struct {
	Tasks []*Task `protobuf:"bytes,1,rep,name=tasks" json:"tasks,omitempty"`
}

func (m *Tasks) Reset()                    { *m = Tasks{} }
func (m *Tasks) String() string            { return proto.CompactTextString(m) }
func (*Tasks) ProtoMessage()               {}
func (*Tasks) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Tasks) GetTasks() []*Task {
	if m != nil {
		return m.Tasks
	}
	return nil
}

func init() {
	proto.RegisterType((*ListRequest)(nil), "ListRequest")
	proto.RegisterType((*Task)(nil), "Task")
	proto.RegisterType((*Tasks)(nil), "Tasks")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Todo service

type TodoClient interface {
	AddTask(ctx context.Context, in *Task, opts ...grpc.CallOption) (*Task, error)
	ListTasks(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*Tasks, error)
}

type todoClient struct {
	cc *grpc.ClientConn
}

func NewTodoClient(cc *grpc.ClientConn) TodoClient {
	return &todoClient{cc}
}

func (c *todoClient) AddTask(ctx context.Context, in *Task, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := grpc.Invoke(ctx, "/Todo/AddTask", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoClient) ListTasks(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*Tasks, error) {
	out := new(Tasks)
	err := grpc.Invoke(ctx, "/Todo/ListTasks", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Todo service

type TodoServer interface {
	AddTask(context.Context, *Task) (*Task, error)
	ListTasks(context.Context, *ListRequest) (*Tasks, error)
}

func RegisterTodoServer(s *grpc.Server, srv TodoServer) {
	s.RegisterService(&_Todo_serviceDesc, srv)
}

func _Todo_AddTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Task)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(TodoServer).AddTask(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Todo_ListTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(TodoServer).ListTasks(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Todo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Todo",
	HandlerType: (*TodoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddTask",
			Handler:    _Todo_AddTask_Handler,
		},
		{
			MethodName: "ListTasks",
			Handler:    _Todo_ListTasks_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 173 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0xc9, 0x4f, 0xc9,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x52, 0xe6, 0xe2, 0xf6, 0xc9, 0x2c, 0x2e, 0x09, 0x4a,
	0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe1, 0x62, 0x2d, 0xc9, 0x2c, 0xc9, 0x49, 0x95, 0x60,
	0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x94, 0xec, 0xb8, 0x58, 0x42, 0x12, 0x8b, 0xb3, 0xb1,
	0xcb, 0x0a, 0x29, 0x70, 0x71, 0xa7, 0xa4, 0x16, 0x27, 0x17, 0x65, 0x16, 0x94, 0x64, 0xe6, 0xe7,
	0x49, 0x30, 0x81, 0xe5, 0x90, 0x85, 0x94, 0x54, 0xb8, 0x58, 0x41, 0xfa, 0x8b, 0x85, 0xa4, 0x81,
	0x06, 0x80, 0x18, 0x40, 0x03, 0x98, 0x35, 0xb8, 0x8d, 0x58, 0xf5, 0x40, 0xc2, 0x41, 0x10, 0x31,
	0x23, 0x37, 0xa0, 0x2d, 0x40, 0x87, 0x09, 0x49, 0x72, 0xb1, 0x3b, 0xa6, 0xa4, 0x80, 0x2d, 0x84,
	0x28, 0x90, 0x82, 0x50, 0x4a, 0x0c, 0x42, 0xca, 0x5c, 0x9c, 0x20, 0xd7, 0x42, 0x0c, 0xe3, 0xd1,
	0x43, 0x72, 0xb9, 0x14, 0x1b, 0x58, 0x4d, 0xb1, 0x12, 0x43, 0x12, 0x1b, 0xd8, 0x67, 0xc6, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x68, 0xfa, 0x95, 0xf3, 0xe7, 0x00, 0x00, 0x00,
}
