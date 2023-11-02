// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: x/grpc/helloworld/helloworld.proto

package helloworld

import (
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

type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Who     string `protobuf:"bytes,1,opt,name=who,proto3" json:"who,omitempty"`
	YourAge int32  `protobuf:"varint,2,opt,name=your_age,json=yourAge,proto3" json:"your_age,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_x_grpc_helloworld_helloworld_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_x_grpc_helloworld_helloworld_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_x_grpc_helloworld_helloworld_proto_rawDescGZIP(), []int{0}
}

func (x *HelloRequest) GetWho() string {
	if x != nil {
		return x.Who
	}
	return ""
}

func (x *HelloRequest) GetYourAge() int32 {
	if x != nil {
		return x.YourAge
	}
	return 0
}

type HelloReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age  int32  `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *HelloReply) Reset() {
	*x = HelloReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_x_grpc_helloworld_helloworld_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReply) ProtoMessage() {}

func (x *HelloReply) ProtoReflect() protoreflect.Message {
	mi := &file_x_grpc_helloworld_helloworld_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReply.ProtoReflect.Descriptor instead.
func (*HelloReply) Descriptor() ([]byte, []int) {
	return file_x_grpc_helloworld_helloworld_proto_rawDescGZIP(), []int{1}
}

func (x *HelloReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HelloReply) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

type KeyValueMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MyMap []*HelloReply `protobuf:"bytes,1,rep,name=my_map,json=myMap,proto3" json:"my_map,omitempty"`
}

func (x *KeyValueMap) Reset() {
	*x = KeyValueMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_x_grpc_helloworld_helloworld_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyValueMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyValueMap) ProtoMessage() {}

func (x *KeyValueMap) ProtoReflect() protoreflect.Message {
	mi := &file_x_grpc_helloworld_helloworld_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyValueMap.ProtoReflect.Descriptor instead.
func (*KeyValueMap) Descriptor() ([]byte, []int) {
	return file_x_grpc_helloworld_helloworld_proto_rawDescGZIP(), []int{2}
}

func (x *KeyValueMap) GetMyMap() []*HelloReply {
	if x != nil {
		return x.MyMap
	}
	return nil
}

var File_x_grpc_helloworld_helloworld_proto protoreflect.FileDescriptor

var file_x_grpc_helloworld_helloworld_proto_rawDesc = []byte{
	0x0a, 0x22, 0x78, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f,
	0x72, 0x6c, 0x64, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x74, 0x65, 0x73, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x22, 0x3b, 0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x77, 0x68, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x77, 0x68, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x79, 0x6f,
	0x75, 0x72, 0x5f, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x79, 0x6f,
	0x75, 0x72, 0x41, 0x67, 0x65, 0x22, 0x32, 0x0a, 0x0a, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x22, 0x44, 0x0a, 0x0b, 0x4b, 0x65, 0x79,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x4d, 0x61, 0x70, 0x12, 0x35, 0x0a, 0x06, 0x6d, 0x79, 0x5f, 0x6d,
	0x61, 0x70, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x72,
	0x70, 0x63, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x05, 0x6d, 0x79, 0x4d, 0x61, 0x70, 0x32,
	0xac, 0x01, 0x0a, 0x07, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x12, 0x4c, 0x0a, 0x08, 0x53,
	0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x20, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x72, 0x70,
	0x63, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x72, 0x70, 0x63, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x53, 0x0a, 0x0e, 0x53, 0x61, 0x79,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x4b, 0x65, 0x79, 0x4d, 0x61, 0x70, 0x12, 0x20, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4d, 0x61, 0x70, 0x42, 0x19,
	0x5a, 0x17, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x2f, 0x78, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_x_grpc_helloworld_helloworld_proto_rawDescOnce sync.Once
	file_x_grpc_helloworld_helloworld_proto_rawDescData = file_x_grpc_helloworld_helloworld_proto_rawDesc
)

func file_x_grpc_helloworld_helloworld_proto_rawDescGZIP() []byte {
	file_x_grpc_helloworld_helloworld_proto_rawDescOnce.Do(func() {
		file_x_grpc_helloworld_helloworld_proto_rawDescData = protoimpl.X.CompressGZIP(file_x_grpc_helloworld_helloworld_proto_rawDescData)
	})
	return file_x_grpc_helloworld_helloworld_proto_rawDescData
}

var file_x_grpc_helloworld_helloworld_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_x_grpc_helloworld_helloworld_proto_goTypes = []interface{}{
	(*HelloRequest)(nil), // 0: testrpc.helloworld.HelloRequest
	(*HelloReply)(nil),   // 1: testrpc.helloworld.HelloReply
	(*KeyValueMap)(nil),  // 2: testrpc.helloworld.KeyValueMap
}
var file_x_grpc_helloworld_helloworld_proto_depIdxs = []int32{
	1, // 0: testrpc.helloworld.KeyValueMap.my_map:type_name -> testrpc.helloworld.HelloReply
	0, // 1: testrpc.helloworld.Greeter.SayHello:input_type -> testrpc.helloworld.HelloRequest
	0, // 2: testrpc.helloworld.Greeter.SayHelloKeyMap:input_type -> testrpc.helloworld.HelloRequest
	1, // 3: testrpc.helloworld.Greeter.SayHello:output_type -> testrpc.helloworld.HelloReply
	2, // 4: testrpc.helloworld.Greeter.SayHelloKeyMap:output_type -> testrpc.helloworld.KeyValueMap
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_x_grpc_helloworld_helloworld_proto_init() }
func file_x_grpc_helloworld_helloworld_proto_init() {
	if File_x_grpc_helloworld_helloworld_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_x_grpc_helloworld_helloworld_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_x_grpc_helloworld_helloworld_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloReply); i {
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
		file_x_grpc_helloworld_helloworld_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyValueMap); i {
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
			RawDescriptor: file_x_grpc_helloworld_helloworld_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_x_grpc_helloworld_helloworld_proto_goTypes,
		DependencyIndexes: file_x_grpc_helloworld_helloworld_proto_depIdxs,
		MessageInfos:      file_x_grpc_helloworld_helloworld_proto_msgTypes,
	}.Build()
	File_x_grpc_helloworld_helloworld_proto = out.File
	file_x_grpc_helloworld_helloworld_proto_rawDesc = nil
	file_x_grpc_helloworld_helloworld_proto_goTypes = nil
	file_x_grpc_helloworld_helloworld_proto_depIdxs = nil
}