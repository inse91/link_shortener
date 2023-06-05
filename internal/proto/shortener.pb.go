// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: api/shortener.proto

package proto

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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Link    string `protobuf:"bytes,1,opt,name=Link,proto3" json:"Link,omitempty"`
	Success bool   `protobuf:"varint,2,opt,name=Success,proto3" json:"Success,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_shortener_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_shortener_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_api_shortener_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *Request) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Link    string `protobuf:"bytes,1,opt,name=Link,proto3" json:"Link,omitempty"`
	Success bool   `protobuf:"varint,2,opt,name=Success,proto3" json:"Success,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_shortener_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_shortener_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_api_shortener_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *Response) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_api_shortener_proto protoreflect.FileDescriptor

var file_api_shortener_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x37, 0x0a, 0x07, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4c, 0x69, 0x6e, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x53, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x22, 0x38, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x4c, 0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x5e,
	0x0a, 0x09, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x06, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0d, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x26, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0d, 0x2e,
	0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x10,
	0x5a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_shortener_proto_rawDescOnce sync.Once
	file_api_shortener_proto_rawDescData = file_api_shortener_proto_rawDesc
)

func file_api_shortener_proto_rawDescGZIP() []byte {
	file_api_shortener_proto_rawDescOnce.Do(func() {
		file_api_shortener_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_shortener_proto_rawDescData)
	})
	return file_api_shortener_proto_rawDescData
}

var file_api_shortener_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_shortener_proto_goTypes = []interface{}{
	(*Request)(nil),  // 0: main.Request
	(*Response)(nil), // 1: main.Response
}
var file_api_shortener_proto_depIdxs = []int32{
	0, // 0: main.Shortener.Create:input_type -> main.Request
	0, // 1: main.Shortener.Get:input_type -> main.Request
	1, // 2: main.Shortener.Create:output_type -> main.Response
	1, // 3: main.Shortener.Get:output_type -> main.Response
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_shortener_proto_init() }
func file_api_shortener_proto_init() {
	if File_api_shortener_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_shortener_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_api_shortener_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_api_shortener_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_shortener_proto_goTypes,
		DependencyIndexes: file_api_shortener_proto_depIdxs,
		MessageInfos:      file_api_shortener_proto_msgTypes,
	}.Build()
	File_api_shortener_proto = out.File
	file_api_shortener_proto_rawDesc = nil
	file_api_shortener_proto_goTypes = nil
	file_api_shortener_proto_depIdxs = nil
}
