// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api.proto

package grpc

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

type AddNewWalletRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title   string `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	Purpose string `protobuf:"bytes,2,opt,name=Purpose,proto3" json:"Purpose,omitempty"`
	IsHot   bool   `protobuf:"varint,3,opt,name=IsHot,proto3" json:"IsHot,omitempty"`
}

func (x *AddNewWalletRequest) Reset() {
	*x = AddNewWalletRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddNewWalletRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddNewWalletRequest) ProtoMessage() {}

func (x *AddNewWalletRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddNewWalletRequest.ProtoReflect.Descriptor instead.
func (*AddNewWalletRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *AddNewWalletRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *AddNewWalletRequest) GetPurpose() string {
	if x != nil {
		return x.Purpose
	}
	return ""
}

func (x *AddNewWalletRequest) GetIsHot() bool {
	if x != nil {
		return x.IsHot
	}
	return false
}

type AddNewWalletResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WalletUUID string `protobuf:"bytes,1,opt,name=WalletUUID,proto3" json:"WalletUUID,omitempty"`
}

func (x *AddNewWalletResponse) Reset() {
	*x = AddNewWalletResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddNewWalletResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddNewWalletResponse) ProtoMessage() {}

func (x *AddNewWalletResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddNewWalletResponse.ProtoReflect.Descriptor instead.
func (*AddNewWalletResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *AddNewWalletResponse) GetWalletUUID() string {
	if x != nil {
		return x.WalletUUID
	}
	return ""
}

type DerivationAddressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountIndex  uint32 `protobuf:"varint,2,opt,name=AccountIndex,proto3" json:"AccountIndex,omitempty"`
	InternalIndex uint32 `protobuf:"varint,3,opt,name=InternalIndex,proto3" json:"InternalIndex,omitempty"`
	AddressIndex  uint32 `protobuf:"varint,4,opt,name=AddressIndex,proto3" json:"AddressIndex,omitempty"`
}

func (x *DerivationAddressRequest) Reset() {
	*x = DerivationAddressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DerivationAddressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DerivationAddressRequest) ProtoMessage() {}

func (x *DerivationAddressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DerivationAddressRequest.ProtoReflect.Descriptor instead.
func (*DerivationAddressRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

func (x *DerivationAddressRequest) GetAccountIndex() uint32 {
	if x != nil {
		return x.AccountIndex
	}
	return 0
}

func (x *DerivationAddressRequest) GetInternalIndex() uint32 {
	if x != nil {
		return x.InternalIndex
	}
	return 0
}

func (x *DerivationAddressRequest) GetAddressIndex() uint32 {
	if x != nil {
		return x.AddressIndex
	}
	return 0
}

type DerivationAddressResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=Address,proto3" json:"Address,omitempty"`
}

func (x *DerivationAddressResponse) Reset() {
	*x = DerivationAddressResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DerivationAddressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DerivationAddressResponse) ProtoMessage() {}

func (x *DerivationAddressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DerivationAddressResponse.ProtoReflect.Descriptor instead.
func (*DerivationAddressResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

func (x *DerivationAddressResponse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x68, 0x64, 0x77,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x22, 0x5b, 0x0a, 0x13, 0x41, 0x64, 0x64,
	0x4e, 0x65, 0x77, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x75, 0x72, 0x70, 0x6f, 0x73,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x50, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x49, 0x73, 0x48, 0x6f, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x05, 0x49, 0x73, 0x48, 0x6f, 0x74, 0x22, 0x36, 0x0a, 0x14, 0x41, 0x64, 0x64, 0x4e, 0x65, 0x77,
	0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x55, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x55, 0x55, 0x49, 0x44, 0x22, 0x88,
	0x01, 0x0a, 0x18, 0x44, 0x65, 0x72, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12,
	0x24, 0x0a, 0x0d, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x22, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x35, 0x0a, 0x19, 0x44, 0x65, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x32, 0xd1, 0x01, 0x0a, 0x0b, 0x48, 0x64, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x41, 0x70, 0x69,
	0x12, 0x57, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x4e, 0x65, 0x77, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x12, 0x21, 0x2e, 0x68, 0x64, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e,
	0x41, 0x64, 0x64, 0x4e, 0x65, 0x77, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x68, 0x64, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x41, 0x64, 0x64, 0x4e, 0x65, 0x77, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x69, 0x0a, 0x14, 0x47, 0x65, 0x74,
	0x44, 0x65, 0x72, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x26, 0x2e, 0x68, 0x64, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x44, 0x65, 0x72, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x68, 0x64, 0x77, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2d, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65,
	0x2f, 0x62, 0x63, 0x2d, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2d, 0x65, 0x74, 0x68, 0x2d, 0x68,
	0x64, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_proto_goTypes = []interface{}{
	(*AddNewWalletRequest)(nil),       // 0: hdwallet_api.AddNewWalletRequest
	(*AddNewWalletResponse)(nil),      // 1: hdwallet_api.AddNewWalletResponse
	(*DerivationAddressRequest)(nil),  // 2: hdwallet_api.DerivationAddressRequest
	(*DerivationAddressResponse)(nil), // 3: hdwallet_api.DerivationAddressResponse
}
var file_api_proto_depIdxs = []int32{
	0, // 0: hdwallet_api.HdWalletApi.AddNewWallet:input_type -> hdwallet_api.AddNewWalletRequest
	2, // 1: hdwallet_api.HdWalletApi.GetDerivationAddress:input_type -> hdwallet_api.DerivationAddressRequest
	1, // 2: hdwallet_api.HdWalletApi.AddNewWallet:output_type -> hdwallet_api.AddNewWalletResponse
	3, // 3: hdwallet_api.HdWalletApi.GetDerivationAddress:output_type -> hdwallet_api.DerivationAddressResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddNewWalletRequest); i {
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
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddNewWalletResponse); i {
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
		file_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DerivationAddressRequest); i {
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
		file_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DerivationAddressResponse); i {
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
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}
