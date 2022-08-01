// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: protos/peer/request.proto

package peer

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

// import "google/protobuf/timestamp.proto";
type SignedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The bytes of Request
	RequestBytes []byte `protobuf:"bytes,1,opt,name=Request_bytes,json=RequestBytes,proto3" json:"Request_bytes,omitempty"`
	// Signaure over RequestBytes; this signature is to be verified against
	// the creator identity contained in the header of the Request message
	// marshaled as RequestBytes
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *SignedRequest) Reset() {
	*x = SignedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_peer_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignedRequest) ProtoMessage() {}

func (x *SignedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_peer_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignedRequest.ProtoReflect.Descriptor instead.
func (*SignedRequest) Descriptor() ([]byte, []int) {
	return file_protos_peer_request_proto_rawDescGZIP(), []int{0}
}

func (x *SignedRequest) GetRequestBytes() []byte {
	if x != nil {
		return x.RequestBytes
	}
	return nil
}

func (x *SignedRequest) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The header of the request. It is the bytes of the Header
	Header []byte `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// The payload of the request as defined by the type in the Request
	// header.
	Payload []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_peer_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_protos_peer_request_proto_msgTypes[1]
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
	return file_protos_peer_request_proto_rawDescGZIP(), []int{1}
}

func (x *Request) GetHeader() []byte {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Request) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

var File_protos_peer_request_proto protoreflect.FileDescriptor

var file_protos_peer_request_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x65, 0x65, 0x72, 0x2f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x70, 0x65, 0x65,
	0x72, 0x22, 0x52, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x62, 0x79,
	0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x3b, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x42, 0x1f, 0x5a, 0x1d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x7a, 0x52, 0x69, 0x63, 0x68, 0x2f, 0x7a, 0x46, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x70,
	0x65, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_peer_request_proto_rawDescOnce sync.Once
	file_protos_peer_request_proto_rawDescData = file_protos_peer_request_proto_rawDesc
)

func file_protos_peer_request_proto_rawDescGZIP() []byte {
	file_protos_peer_request_proto_rawDescOnce.Do(func() {
		file_protos_peer_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_peer_request_proto_rawDescData)
	})
	return file_protos_peer_request_proto_rawDescData
}

var file_protos_peer_request_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_peer_request_proto_goTypes = []interface{}{
	(*SignedRequest)(nil), // 0: peer.SignedRequest
	(*Request)(nil),       // 1: peer.Request
}
var file_protos_peer_request_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_peer_request_proto_init() }
func file_protos_peer_request_proto_init() {
	if File_protos_peer_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_peer_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignedRequest); i {
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
		file_protos_peer_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_peer_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_peer_request_proto_goTypes,
		DependencyIndexes: file_protos_peer_request_proto_depIdxs,
		MessageInfos:      file_protos_peer_request_proto_msgTypes,
	}.Build()
	File_protos_peer_request_proto = out.File
	file_protos_peer_request_proto_rawDesc = nil
	file_protos_peer_request_proto_goTypes = nil
	file_protos_peer_request_proto_depIdxs = nil
}
