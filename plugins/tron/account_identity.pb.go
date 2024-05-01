//    MIT License
//
//    Copyright (c) 2022-2024 Aleksei Kotelnikov(gudron2s@gmail.com)
//
//    Permission is hereby granted, free of charge, to any person obtaining a copy
//    of this software and associated documentation files (the "Software"), to deal
//    in the Software without restriction, including without limitation the rights
//    to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//    copies of the Software, and to permit persons to whom the Software is
//    furnished to do so, subject to the following conditions:
//
//    The above copyright notice and this permission notice shall be included in all
//    copies or substantial portions of the Software.
//
//    THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//    IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//    FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//    AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//    LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//    OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
//    SOFTWARE.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.23.2
// source: account_identity.proto

package main

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

type AccountIdentity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountIndex  uint32 `protobuf:"varint,1,opt,name=AccountIndex,proto3" json:"AccountIndex,omitempty"`
	InternalIndex uint32 `protobuf:"varint,2,opt,name=InternalIndex,proto3" json:"InternalIndex,omitempty"`
	AddressIndex  uint32 `protobuf:"varint,3,opt,name=AddressIndex,proto3" json:"AddressIndex,omitempty"`
	Address       string `protobuf:"bytes,4,opt,name=Address,proto3" json:"Address,omitempty"`
}

func (x *AccountIdentity) Reset() {
	*x = AccountIdentity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_identity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountIdentity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountIdentity) ProtoMessage() {}

func (x *AccountIdentity) ProtoReflect() protoreflect.Message {
	mi := &file_account_identity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountIdentity.ProtoReflect.Descriptor instead.
func (*AccountIdentity) Descriptor() ([]byte, []int) {
	return file_account_identity_proto_rawDescGZIP(), []int{0}
}

func (x *AccountIdentity) GetAccountIndex() uint32 {
	if x != nil {
		return x.AccountIndex
	}
	return 0
}

func (x *AccountIdentity) GetInternalIndex() uint32 {
	if x != nil {
		return x.InternalIndex
	}
	return 0
}

func (x *AccountIdentity) GetAddressIndex() uint32 {
	if x != nil {
		return x.AddressIndex
	}
	return 0
}

func (x *AccountIdentity) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type RangeAccountRequestUnit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountIndex     uint32 `protobuf:"varint,3,opt,name=AccountIndex,proto3" json:"AccountIndex,omitempty"`
	InternalIndex    uint32 `protobuf:"varint,4,opt,name=InternalIndex,proto3" json:"InternalIndex,omitempty"`
	AddressIndexFrom uint32 `protobuf:"varint,5,opt,name=AddressIndexFrom,proto3" json:"AddressIndexFrom,omitempty"`
	AddressIndexTo   uint32 `protobuf:"varint,6,opt,name=AddressIndexTo,proto3" json:"AddressIndexTo,omitempty"`
}

func (x *RangeAccountRequestUnit) Reset() {
	*x = RangeAccountRequestUnit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_identity_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RangeAccountRequestUnit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RangeAccountRequestUnit) ProtoMessage() {}

func (x *RangeAccountRequestUnit) ProtoReflect() protoreflect.Message {
	mi := &file_account_identity_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RangeAccountRequestUnit.ProtoReflect.Descriptor instead.
func (*RangeAccountRequestUnit) Descriptor() ([]byte, []int) {
	return file_account_identity_proto_rawDescGZIP(), []int{1}
}

func (x *RangeAccountRequestUnit) GetAccountIndex() uint32 {
	if x != nil {
		return x.AccountIndex
	}
	return 0
}

func (x *RangeAccountRequestUnit) GetInternalIndex() uint32 {
	if x != nil {
		return x.InternalIndex
	}
	return 0
}

func (x *RangeAccountRequestUnit) GetAddressIndexFrom() uint32 {
	if x != nil {
		return x.AddressIndexFrom
	}
	return 0
}

func (x *RangeAccountRequestUnit) GetAddressIndexTo() uint32 {
	if x != nil {
		return x.AddressIndexTo
	}
	return 0
}

var File_account_identity_proto protoreflect.FileDescriptor

var file_account_identity_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x99,
	0x01, 0x0a, 0x0f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x24, 0x0a, 0x0d, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x22, 0x0a, 0x0c,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0xb7, 0x01, 0x0a, 0x17, 0x72,
	0x61, 0x6e, 0x67, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x24, 0x0a, 0x0d, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0d, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x2a, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x46, 0x72, 0x6f, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x26, 0x0a, 0x0e,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x54, 0x6f, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x54, 0x6f, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2d, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65,
	0x2f, 0x62, 0x63, 0x2d, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2d, 0x74, 0x72, 0x6f, 0x6e, 0x2d,
	0x68, 0x64, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73,
	0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_account_identity_proto_rawDescOnce sync.Once
	file_account_identity_proto_rawDescData = file_account_identity_proto_rawDesc
)

func file_account_identity_proto_rawDescGZIP() []byte {
	file_account_identity_proto_rawDescOnce.Do(func() {
		file_account_identity_proto_rawDescData = protoimpl.X.CompressGZIP(file_account_identity_proto_rawDescData)
	})
	return file_account_identity_proto_rawDescData
}

var file_account_identity_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_account_identity_proto_goTypes = []interface{}{
	(*AccountIdentity)(nil),         // 0: main.accountIdentity
	(*RangeAccountRequestUnit)(nil), // 1: main.rangeAccountRequestUnit
}
var file_account_identity_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_account_identity_proto_init() }
func file_account_identity_proto_init() {
	if File_account_identity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_account_identity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountIdentity); i {
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
		file_account_identity_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RangeAccountRequestUnit); i {
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
			RawDescriptor: file_account_identity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_account_identity_proto_goTypes,
		DependencyIndexes: file_account_identity_proto_depIdxs,
		MessageInfos:      file_account_identity_proto_msgTypes,
	}.Build()
	File_account_identity_proto = out.File
	file_account_identity_proto_rawDesc = nil
	file_account_identity_proto_goTypes = nil
	file_account_identity_proto_depIdxs = nil
}
