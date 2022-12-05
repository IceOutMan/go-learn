// 这个就是protobuf的中间文件

// 指定的当前proto语法的版本，有2和3

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.4
// source: helloworld.proto

// 指定等会文件生成出来的package

package service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type StringMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ss []*StringSingle `protobuf:"bytes,1,rep,name=ss,proto3" json:"ss,omitempty"`
}

func (x *StringMessage) Reset() {
	*x = StringMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringMessage) ProtoMessage() {}

func (x *StringMessage) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringMessage.ProtoReflect.Descriptor instead.
func (*StringMessage) Descriptor() ([]byte, []int) {
	return file_helloworld_proto_rawDescGZIP(), []int{0}
}

func (x *StringMessage) GetSs() []*StringSingle {
	if x != nil {
		return x.Ss
	}
	return nil
}

type StringSingle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *StringSingle) Reset() {
	*x = StringSingle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringSingle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringSingle) ProtoMessage() {}

func (x *StringSingle) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringSingle.ProtoReflect.Descriptor instead.
func (*StringSingle) Descriptor() ([]byte, []int) {
	return file_helloworld_proto_rawDescGZIP(), []int{1}
}

func (x *StringSingle) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *StringSingle) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_helloworld_proto_rawDescGZIP(), []int{2}
}

var File_helloworld_proto protoreflect.FileDescriptor

var file_helloworld_proto_rawDesc = []byte{
	0x0a, 0x10, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x36, 0x0a, 0x0d, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x25, 0x0a, 0x02,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x52,
	0x02, 0x73, 0x73, 0x22, 0x32, 0x0a, 0x0c, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x53, 0x69, 0x6e,
	0x67, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x32, 0x3d, 0x0a, 0x07, 0x4d, 0x61, 0x78, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x32, 0x0a, 0x04, 0x45,
	0x63, 0x68, 0x6f, 0x12, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42,
	0x0c, 0x5a, 0x0a, 0x2e, 0x2e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_helloworld_proto_rawDescOnce sync.Once
	file_helloworld_proto_rawDescData = file_helloworld_proto_rawDesc
)

func file_helloworld_proto_rawDescGZIP() []byte {
	file_helloworld_proto_rawDescOnce.Do(func() {
		file_helloworld_proto_rawDescData = protoimpl.X.CompressGZIP(file_helloworld_proto_rawDescData)
	})
	return file_helloworld_proto_rawDescData
}

var file_helloworld_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_helloworld_proto_goTypes = []interface{}{
	(*StringMessage)(nil), // 0: service.StringMessage
	(*StringSingle)(nil),  // 1: service.StringSingle
	(*Empty)(nil),         // 2: service.Empty
}
var file_helloworld_proto_depIdxs = []int32{
	1, // 0: service.StringMessage.ss:type_name -> service.StringSingle
	2, // 1: service.MaxSize.Echo:input_type -> service.Empty
	0, // 2: service.MaxSize.Echo:output_type -> service.StringMessage
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_helloworld_proto_init() }
func file_helloworld_proto_init() {
	if File_helloworld_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_helloworld_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StringMessage); i {
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
		file_helloworld_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StringSingle); i {
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
		file_helloworld_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_helloworld_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_helloworld_proto_goTypes,
		DependencyIndexes: file_helloworld_proto_depIdxs,
		MessageInfos:      file_helloworld_proto_msgTypes,
	}.Build()
	File_helloworld_proto = out.File
	file_helloworld_proto_rawDesc = nil
	file_helloworld_proto_goTypes = nil
	file_helloworld_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MaxSizeClient is the client API for MaxSize service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MaxSizeClient interface {
	Echo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (MaxSize_EchoClient, error)
}

type maxSizeClient struct {
	cc grpc.ClientConnInterface
}

func NewMaxSizeClient(cc grpc.ClientConnInterface) MaxSizeClient {
	return &maxSizeClient{cc}
}

func (c *maxSizeClient) Echo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (MaxSize_EchoClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MaxSize_serviceDesc.Streams[0], "/service.MaxSize/Echo", opts...)
	if err != nil {
		return nil, err
	}
	x := &maxSizeEchoClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MaxSize_EchoClient interface {
	Recv() (*StringMessage, error)
	grpc.ClientStream
}

type maxSizeEchoClient struct {
	grpc.ClientStream
}

func (x *maxSizeEchoClient) Recv() (*StringMessage, error) {
	m := new(StringMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MaxSizeServer is the server API for MaxSize service.
type MaxSizeServer interface {
	Echo(*Empty, MaxSize_EchoServer) error
}

// UnimplementedMaxSizeServer can be embedded to have forward compatible implementations.
type UnimplementedMaxSizeServer struct {
}

func (*UnimplementedMaxSizeServer) Echo(*Empty, MaxSize_EchoServer) error {
	return status.Errorf(codes.Unimplemented, "method Echo not implemented")
}

func RegisterMaxSizeServer(s *grpc.Server, srv MaxSizeServer) {
	s.RegisterService(&_MaxSize_serviceDesc, srv)
}

func _MaxSize_Echo_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MaxSizeServer).Echo(m, &maxSizeEchoServer{stream})
}

type MaxSize_EchoServer interface {
	Send(*StringMessage) error
	grpc.ServerStream
}

type maxSizeEchoServer struct {
	grpc.ServerStream
}

func (x *maxSizeEchoServer) Send(m *StringMessage) error {
	return x.ServerStream.SendMsg(m)
}

var _MaxSize_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.MaxSize",
	HandlerType: (*MaxSizeServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Echo",
			Handler:       _MaxSize_Echo_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "helloworld.proto",
}