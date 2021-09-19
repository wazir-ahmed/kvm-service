// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: kvm.proto

package protobuf

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

type AgentIdentity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity string `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
}

func (x *AgentIdentity) Reset() {
	*x = AgentIdentity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kvm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentIdentity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentIdentity) ProtoMessage() {}

func (x *AgentIdentity) ProtoReflect() protoreflect.Message {
	mi := &file_kvm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentIdentity.ProtoReflect.Descriptor instead.
func (*AgentIdentity) Descriptor() ([]byte, []int) {
	return file_kvm_proto_rawDescGZIP(), []int{0}
}

func (x *AgentIdentity) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kvm_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_kvm_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_kvm_proto_rawDescGZIP(), []int{1}
}

func (x *Status) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type PolicyData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PolicyName string `protobuf:"bytes,1,opt,name=policyName,proto3" json:"policyName,omitempty"`
	PolicyData []byte `protobuf:"bytes,2,opt,name=policyData,proto3" json:"policyData,omitempty"`
}

func (x *PolicyData) Reset() {
	*x = PolicyData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kvm_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PolicyData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolicyData) ProtoMessage() {}

func (x *PolicyData) ProtoReflect() protoreflect.Message {
	mi := &file_kvm_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolicyData.ProtoReflect.Descriptor instead.
func (*PolicyData) Descriptor() ([]byte, []int) {
	return file_kvm_proto_rawDescGZIP(), []int{2}
}

func (x *PolicyData) GetPolicyName() string {
	if x != nil {
		return x.PolicyName
	}
	return ""
}

func (x *PolicyData) GetPolicyData() []byte {
	if x != nil {
		return x.PolicyData
	}
	return nil
}

var File_kvm_proto protoreflect.FileDescriptor

var file_kvm_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6b, 0x76, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6b, 0x76, 0x6d,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x2b, 0x0a, 0x0d, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x22, 0x20, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x4c, 0x0a, 0x0a, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x44, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x44, 0x61, 0x74, 0x61, 0x32, 0x8b, 0x01, 0x0a, 0x03, 0x4b, 0x56, 0x4d, 0x12, 0x46, 0x0a, 0x15,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x19, 0x2e, 0x6b, 0x76, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x1a, 0x12, 0x2e, 0x6b, 0x76, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x3c, 0x0a, 0x0a, 0x73, 0x65, 0x6e, 0x64, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x12, 0x12, 0x2e, 0x6b, 0x76, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x16, 0x2e, 0x6b, 0x76, 0x6d, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x44, 0x61, 0x74, 0x61, 0x28, 0x01,
	0x30, 0x01, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6b, 0x75, 0x62, 0x65, 0x61, 0x72, 0x6d, 0x6f, 0x72, 0x2f, 0x4b, 0x56, 0x4d, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kvm_proto_rawDescOnce sync.Once
	file_kvm_proto_rawDescData = file_kvm_proto_rawDesc
)

func file_kvm_proto_rawDescGZIP() []byte {
	file_kvm_proto_rawDescOnce.Do(func() {
		file_kvm_proto_rawDescData = protoimpl.X.CompressGZIP(file_kvm_proto_rawDescData)
	})
	return file_kvm_proto_rawDescData
}

var file_kvm_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_kvm_proto_goTypes = []interface{}{
	(*AgentIdentity)(nil), // 0: kvmService.agentIdentity
	(*Status)(nil),        // 1: kvmService.status
	(*PolicyData)(nil),    // 2: kvmService.policyData
}
var file_kvm_proto_depIdxs = []int32{
	0, // 0: kvmService.KVM.registerAgentIdentity:input_type -> kvmService.agentIdentity
	1, // 1: kvmService.KVM.sendPolicy:input_type -> kvmService.status
	1, // 2: kvmService.KVM.registerAgentIdentity:output_type -> kvmService.status
	2, // 3: kvmService.KVM.sendPolicy:output_type -> kvmService.policyData
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_kvm_proto_init() }
func file_kvm_proto_init() {
	if File_kvm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kvm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentIdentity); i {
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
		file_kvm_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_kvm_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PolicyData); i {
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
			RawDescriptor: file_kvm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kvm_proto_goTypes,
		DependencyIndexes: file_kvm_proto_depIdxs,
		MessageInfos:      file_kvm_proto_msgTypes,
	}.Build()
	File_kvm_proto = out.File
	file_kvm_proto_rawDesc = nil
	file_kvm_proto_goTypes = nil
	file_kvm_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// KVMClient is the client API for KVM service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KVMClient interface {
	RegisterAgentIdentity(ctx context.Context, in *AgentIdentity, opts ...grpc.CallOption) (*Status, error)
	SendPolicy(ctx context.Context, opts ...grpc.CallOption) (KVM_SendPolicyClient, error)
}

type kVMClient struct {
	cc grpc.ClientConnInterface
}

func NewKVMClient(cc grpc.ClientConnInterface) KVMClient {
	return &kVMClient{cc}
}

func (c *kVMClient) RegisterAgentIdentity(ctx context.Context, in *AgentIdentity, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/kvmService.KVM/registerAgentIdentity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVMClient) SendPolicy(ctx context.Context, opts ...grpc.CallOption) (KVM_SendPolicyClient, error) {
	stream, err := c.cc.NewStream(ctx, &_KVM_serviceDesc.Streams[0], "/kvmService.KVM/sendPolicy", opts...)
	if err != nil {
		return nil, err
	}
	x := &kVMSendPolicyClient{stream}
	return x, nil
}

type KVM_SendPolicyClient interface {
	Send(*Status) error
	Recv() (*PolicyData, error)
	grpc.ClientStream
}

type kVMSendPolicyClient struct {
	grpc.ClientStream
}

func (x *kVMSendPolicyClient) Send(m *Status) error {
	return x.ClientStream.SendMsg(m)
}

func (x *kVMSendPolicyClient) Recv() (*PolicyData, error) {
	m := new(PolicyData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// KVMServer is the server API for KVM service.
type KVMServer interface {
	RegisterAgentIdentity(context.Context, *AgentIdentity) (*Status, error)
	SendPolicy(KVM_SendPolicyServer) error
}

// UnimplementedKVMServer can be embedded to have forward compatible implementations.
type UnimplementedKVMServer struct {
}

func (*UnimplementedKVMServer) RegisterAgentIdentity(context.Context, *AgentIdentity) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterAgentIdentity not implemented")
}
func (*UnimplementedKVMServer) SendPolicy(KVM_SendPolicyServer) error {
	return status.Errorf(codes.Unimplemented, "method SendPolicy not implemented")
}

func RegisterKVMServer(s *grpc.Server, srv KVMServer) {
	s.RegisterService(&_KVM_serviceDesc, srv)
}

func _KVM_RegisterAgentIdentity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentIdentity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVMServer).RegisterAgentIdentity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kvmService.KVM/RegisterAgentIdentity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVMServer).RegisterAgentIdentity(ctx, req.(*AgentIdentity))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVM_SendPolicy_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(KVMServer).SendPolicy(&kVMSendPolicyServer{stream})
}

type KVM_SendPolicyServer interface {
	Send(*PolicyData) error
	Recv() (*Status, error)
	grpc.ServerStream
}

type kVMSendPolicyServer struct {
	grpc.ServerStream
}

func (x *kVMSendPolicyServer) Send(m *PolicyData) error {
	return x.ServerStream.SendMsg(m)
}

func (x *kVMSendPolicyServer) Recv() (*Status, error) {
	m := new(Status)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _KVM_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kvmService.KVM",
	HandlerType: (*KVMServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "registerAgentIdentity",
			Handler:    _KVM_RegisterAgentIdentity_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "sendPolicy",
			Handler:       _KVM_SendPolicy_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "kvm.proto",
}
