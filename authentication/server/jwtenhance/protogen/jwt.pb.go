// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: jwt.proto

package protogen

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

type JwtHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Algorithm string `protobuf:"bytes,1,opt,name=algorithm,proto3" json:"algorithm,omitempty"`
	Issuer    string `protobuf:"bytes,2,opt,name=issuer,proto3" json:"issuer,omitempty"`
}

func (x *JwtHeader) Reset() {
	*x = JwtHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jwt_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JwtHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JwtHeader) ProtoMessage() {}

func (x *JwtHeader) ProtoReflect() protoreflect.Message {
	mi := &file_jwt_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JwtHeader.ProtoReflect.Descriptor instead.
func (*JwtHeader) Descriptor() ([]byte, []int) {
	return file_jwt_proto_rawDescGZIP(), []int{0}
}

func (x *JwtHeader) GetAlgorithm() string {
	if x != nil {
		return x.Algorithm
	}
	return ""
}

func (x *JwtHeader) GetIssuer() string {
	if x != nil {
		return x.Issuer
	}
	return ""
}

type JwtPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JwtId      string `protobuf:"bytes,1,opt,name=jwt_id,json=jwtId,proto3" json:"jwt_id,omitempty"`
	Ak         string `protobuf:"bytes,2,opt,name=ak,proto3" json:"ak,omitempty"`
	ServerSign string `protobuf:"bytes,3,opt,name=server_sign,json=serverSign,proto3" json:"server_sign,omitempty"`
	SecretId   int64  `protobuf:"varint,4,opt,name=secret_id,json=secretId,proto3" json:"secret_id,omitempty"`
	Nonce      string `protobuf:"bytes,5,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Ip         string `protobuf:"bytes,6,opt,name=ip,proto3" json:"ip,omitempty"`
	Timestamp  int64  `protobuf:"varint,7,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *JwtPayload) Reset() {
	*x = JwtPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jwt_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JwtPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JwtPayload) ProtoMessage() {}

func (x *JwtPayload) ProtoReflect() protoreflect.Message {
	mi := &file_jwt_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JwtPayload.ProtoReflect.Descriptor instead.
func (*JwtPayload) Descriptor() ([]byte, []int) {
	return file_jwt_proto_rawDescGZIP(), []int{1}
}

func (x *JwtPayload) GetJwtId() string {
	if x != nil {
		return x.JwtId
	}
	return ""
}

func (x *JwtPayload) GetAk() string {
	if x != nil {
		return x.Ak
	}
	return ""
}

func (x *JwtPayload) GetServerSign() string {
	if x != nil {
		return x.ServerSign
	}
	return ""
}

func (x *JwtPayload) GetSecretId() int64 {
	if x != nil {
		return x.SecretId
	}
	return 0
}

func (x *JwtPayload) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

func (x *JwtPayload) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *JwtPayload) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type Jwt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header    *JwtHeader  `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Payload   *JwtPayload `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	Signature []byte      `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *Jwt) Reset() {
	*x = Jwt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jwt_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Jwt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Jwt) ProtoMessage() {}

func (x *Jwt) ProtoReflect() protoreflect.Message {
	mi := &file_jwt_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Jwt.ProtoReflect.Descriptor instead.
func (*Jwt) Descriptor() ([]byte, []int) {
	return file_jwt_proto_rawDescGZIP(), []int{2}
}

func (x *Jwt) GetHeader() *JwtHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Jwt) GetPayload() *JwtPayload {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *Jwt) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

var File_jwt_proto protoreflect.FileDescriptor

var file_jwt_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6a, 0x77, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a, 0x09, 0x4a,
	0x77, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x6c, 0x67, 0x6f,
	0x72, 0x69, 0x74, 0x68, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x6c, 0x67,
	0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x22, 0xb5,
	0x01, 0x0a, 0x0a, 0x4a, 0x77, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x15, 0x0a,
	0x06, 0x6a, 0x77, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a,
	0x77, 0x74, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x61, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x61, 0x6b, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73,
	0x69, 0x67, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x53, 0x69, 0x67, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x6e, 0x0a, 0x03, 0x4a, 0x77, 0x74, 0x12, 0x22, 0x0a,
	0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x4a, 0x77, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x12, 0x25, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x4a, 0x77, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52,
	0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x42, 0x0b, 0x5a, 0x09, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_jwt_proto_rawDescOnce sync.Once
	file_jwt_proto_rawDescData = file_jwt_proto_rawDesc
)

func file_jwt_proto_rawDescGZIP() []byte {
	file_jwt_proto_rawDescOnce.Do(func() {
		file_jwt_proto_rawDescData = protoimpl.X.CompressGZIP(file_jwt_proto_rawDescData)
	})
	return file_jwt_proto_rawDescData
}

var file_jwt_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_jwt_proto_goTypes = []any{
	(*JwtHeader)(nil),  // 0: JwtHeader
	(*JwtPayload)(nil), // 1: JwtPayload
	(*Jwt)(nil),        // 2: Jwt
}
var file_jwt_proto_depIdxs = []int32{
	0, // 0: Jwt.header:type_name -> JwtHeader
	1, // 1: Jwt.payload:type_name -> JwtPayload
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_jwt_proto_init() }
func file_jwt_proto_init() {
	if File_jwt_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_jwt_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*JwtHeader); i {
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
		file_jwt_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*JwtPayload); i {
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
		file_jwt_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Jwt); i {
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
			RawDescriptor: file_jwt_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_jwt_proto_goTypes,
		DependencyIndexes: file_jwt_proto_depIdxs,
		MessageInfos:      file_jwt_proto_msgTypes,
	}.Build()
	File_jwt_proto = out.File
	file_jwt_proto_rawDesc = nil
	file_jwt_proto_goTypes = nil
	file_jwt_proto_depIdxs = nil
}
