// Copyright 2023 - MinIO, Inc. All rights reserved.
// Use of this source code is governed by the AGPLv3
// license that can be found in the LICENSE file.

// Generate the Go protobuf code by running the protobuf compiler
// from the repository root:
//
//   $ protoc -I=./kms/protobuf --go_out=. ./kms/protobuf/*.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: request.proto

package protobuf

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

type EditClusterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// RemoveIDs is a list of KMS server IDs that are removed from
	// the cluster definition of the KMS server that receives the
	// request.
	RemoveIDs []uint32 `protobuf:"varint,1,rep,packed,name=RemoveIDs,json=remove,proto3" json:"RemoveIDs,omitempty"`
}

func (x *EditClusterRequest) Reset() {
	*x = EditClusterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditClusterRequest) ProtoMessage() {}

func (x *EditClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditClusterRequest.ProtoReflect.Descriptor instead.
func (*EditClusterRequest) Descriptor() ([]byte, []int) {
	return file_request_proto_rawDescGZIP(), []int{0}
}

func (x *EditClusterRequest) GetRemoveIDs() []uint32 {
	if x != nil {
		return x.RemoveIDs
	}
	return nil
}

type CreateKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Type is the type of the created key. If not set, the KMS server picks
	// a key type.
	Type *string `protobuf:"bytes,1,opt,name=Type,json=type,proto3,oneof" json:"Type,omitempty"`
}

func (x *CreateKeyRequest) Reset() {
	*x = CreateKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateKeyRequest) ProtoMessage() {}

func (x *CreateKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateKeyRequest.ProtoReflect.Descriptor instead.
func (*CreateKeyRequest) Descriptor() ([]byte, []int) {
	return file_request_proto_rawDescGZIP(), []int{1}
}

func (x *CreateKeyRequest) GetType() string {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return ""
}

type AddKeyVersionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Type is the type of the key version added. If not set, the KMS server
	// picks a key type.
	Type *string `protobuf:"bytes,1,opt,name=Type,json=type,proto3,oneof" json:"Type,omitempty"`
}

func (x *AddKeyVersionRequest) Reset() {
	*x = AddKeyVersionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddKeyVersionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddKeyVersionRequest) ProtoMessage() {}

func (x *AddKeyVersionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_request_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddKeyVersionRequest.ProtoReflect.Descriptor instead.
func (*AddKeyVersionRequest) Descriptor() ([]byte, []int) {
	return file_request_proto_rawDescGZIP(), []int{2}
}

func (x *AddKeyVersionRequest) GetType() string {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return ""
}

type RemoveKeyVersionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Version identifies the key version within the key ring to remove.
	// If zero, the latest key version is removed.
	Version uint32 `protobuf:"varint,1,opt,name=Version,json=version,proto3" json:"Version,omitempty"`
}

func (x *RemoveKeyVersionRequest) Reset() {
	*x = RemoveKeyVersionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveKeyVersionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveKeyVersionRequest) ProtoMessage() {}

func (x *RemoveKeyVersionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_request_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveKeyVersionRequest.ProtoReflect.Descriptor instead.
func (*RemoveKeyVersionRequest) Descriptor() ([]byte, []int) {
	return file_request_proto_rawDescGZIP(), []int{3}
}

func (x *RemoveKeyVersionRequest) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

type EncryptRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Version identifies the key version within the key ring used to encrypt
	// the plaintext. If zero, the latest key version is used.
	Version uint32 `protobuf:"varint,1,opt,name=Version,json=version,proto3" json:"Version,omitempty"`
	// Plaintext is the plain message that is encrypted.
	Plaintext []byte `protobuf:"bytes,2,opt,name=Plaintext,json=plaintext,proto3" json:"Plaintext,omitempty"`
	// AssociatedData is additional data that is not encrypted but crypto. bound
	// to the ciphertext. The same associated data must be provided when decrypting
	// the ciphertext.
	//
	// Associated data should describe the context of the plaintext data. For example,
	// the name of the file that gets encrypted.
	AssociatedData []byte `protobuf:"bytes,3,opt,name=AssociatedData,json=associated_data,proto3" json:"AssociatedData,omitempty"`
}

func (x *EncryptRequest) Reset() {
	*x = EncryptRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncryptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncryptRequest) ProtoMessage() {}

func (x *EncryptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_request_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncryptRequest.ProtoReflect.Descriptor instead.
func (*EncryptRequest) Descriptor() ([]byte, []int) {
	return file_request_proto_rawDescGZIP(), []int{4}
}

func (x *EncryptRequest) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *EncryptRequest) GetPlaintext() []byte {
	if x != nil {
		return x.Plaintext
	}
	return nil
}

func (x *EncryptRequest) GetAssociatedData() []byte {
	if x != nil {
		return x.AssociatedData
	}
	return nil
}

type GenerateKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Version identifies the key version within the key ring used to generate
	// the data encryption key. If zero, the latest key version is used.
	Version uint32 `protobuf:"varint,1,opt,name=Version,json=version,proto3" json:"Version,omitempty"`
	// AssociatedData is additional data that is not encrypted but crypto. bound
	// to the ciphertext of the data encryption key. The same associated data must
	// be provided when decrypting the ciphertext.
	//
	// Associated data should describe the context within the data encryption key
	// is used. For example, the name of the file that gets encrypted with the
	// data encryption key.
	AssociatedData []byte `protobuf:"bytes,2,opt,name=AssociatedData,json=associated_data,proto3" json:"AssociatedData,omitempty"`
	// Length is the length of the generated plaintext data encryption key in bytes.
	// At most 1024 (8192 bits). If not present, defaults to 32 (256 bits).
	Length *uint32 `protobuf:"varint,3,opt,name=Length,json=length,proto3,oneof" json:"Length,omitempty"`
}

func (x *GenerateKeyRequest) Reset() {
	*x = GenerateKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateKeyRequest) ProtoMessage() {}

func (x *GenerateKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_request_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateKeyRequest.ProtoReflect.Descriptor instead.
func (*GenerateKeyRequest) Descriptor() ([]byte, []int) {
	return file_request_proto_rawDescGZIP(), []int{5}
}

func (x *GenerateKeyRequest) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *GenerateKeyRequest) GetAssociatedData() []byte {
	if x != nil {
		return x.AssociatedData
	}
	return nil
}

func (x *GenerateKeyRequest) GetLength() uint32 {
	if x != nil && x.Length != nil {
		return *x.Length
	}
	return 0
}

type DecryptRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Version identifies the key version within the key ring that should be
	// used to decrypt the ciphertext. Must not be zero.
	Version uint32 `protobuf:"varint,1,opt,name=Version,json=version,proto3" json:"Version,omitempty"`
	// Ciphertext is the encrypted message that gets decrypted.
	Ciphertext []byte `protobuf:"bytes,2,opt,name=Ciphertext,json=ciphertext,proto3" json:"Ciphertext,omitempty"`
	// AssociatedData is additional data that has been crypto. bound to the
	// ciphertext.
	AssociatedData []byte `protobuf:"bytes,3,opt,name=AssociatedData,json=associated_data,proto3" json:"AssociatedData,omitempty"`
}

func (x *DecryptRequest) Reset() {
	*x = DecryptRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecryptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecryptRequest) ProtoMessage() {}

func (x *DecryptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_request_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecryptRequest.ProtoReflect.Descriptor instead.
func (*DecryptRequest) Descriptor() ([]byte, []int) {
	return file_request_proto_rawDescGZIP(), []int{6}
}

func (x *DecryptRequest) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *DecryptRequest) GetCiphertext() []byte {
	if x != nil {
		return x.Ciphertext
	}
	return nil
}

func (x *DecryptRequest) GetAssociatedData() []byte {
	if x != nil {
		return x.AssociatedData
	}
	return nil
}

var File_request_proto protoreflect.FileDescriptor

var file_request_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x09, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x2e, 0x6b, 0x6d, 0x73, 0x22, 0x2f, 0x0a, 0x12, 0x45, 0x64,
	0x69, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x19, 0x0a, 0x09, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x49, 0x44, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x22, 0x34, 0x0a, 0x10, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x54, 0x79, 0x70,
	0x65, 0x22, 0x38, 0x0a, 0x14, 0x41, 0x64, 0x64, 0x4b, 0x65, 0x79, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x04, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x88,
	0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x54, 0x79, 0x70, 0x65, 0x22, 0x33, 0x0a, 0x17, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4b, 0x65, 0x79, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x22, 0x71, 0x0a, 0x0e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09,
	0x50, 0x6c, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x09, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x27, 0x0a, 0x0e, 0x41, 0x73,
	0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0f, 0x61, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x7f, 0x0a, 0x12, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4b,
	0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x0e, 0x41, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x65,
	0x64, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x61, 0x73, 0x73,
	0x6f, 0x63, 0x69, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x06,
	0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x06,
	0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x4c, 0x65,
	0x6e, 0x67, 0x74, 0x68, 0x22, 0x73, 0x0a, 0x0e, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x74, 0x65, 0x78, 0x74,
	0x12, 0x27, 0x0a, 0x0e, 0x41, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x65, 0x64, 0x44, 0x61,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x61, 0x73, 0x73, 0x6f, 0x63, 0x69,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x42, 0x0e, 0x5a, 0x0c, 0x6b, 0x6d, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_request_proto_rawDescOnce sync.Once
	file_request_proto_rawDescData = file_request_proto_rawDesc
)

func file_request_proto_rawDescGZIP() []byte {
	file_request_proto_rawDescOnce.Do(func() {
		file_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_request_proto_rawDescData)
	})
	return file_request_proto_rawDescData
}

var file_request_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_request_proto_goTypes = []interface{}{
	(*EditClusterRequest)(nil),      // 0: minio.kms.EditClusterRequest
	(*CreateKeyRequest)(nil),        // 1: minio.kms.CreateKeyRequest
	(*AddKeyVersionRequest)(nil),    // 2: minio.kms.AddKeyVersionRequest
	(*RemoveKeyVersionRequest)(nil), // 3: minio.kms.RemoveKeyVersionRequest
	(*EncryptRequest)(nil),          // 4: minio.kms.EncryptRequest
	(*GenerateKeyRequest)(nil),      // 5: minio.kms.GenerateKeyRequest
	(*DecryptRequest)(nil),          // 6: minio.kms.DecryptRequest
}
var file_request_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_request_proto_init() }
func file_request_proto_init() {
	if File_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditClusterRequest); i {
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
		file_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateKeyRequest); i {
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
		file_request_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddKeyVersionRequest); i {
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
		file_request_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveKeyVersionRequest); i {
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
		file_request_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncryptRequest); i {
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
		file_request_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateKeyRequest); i {
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
		file_request_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecryptRequest); i {
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
	file_request_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_request_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_request_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_request_proto_goTypes,
		DependencyIndexes: file_request_proto_depIdxs,
		MessageInfos:      file_request_proto_msgTypes,
	}.Build()
	File_request_proto = out.File
	file_request_proto_rawDesc = nil
	file_request_proto_goTypes = nil
	file_request_proto_depIdxs = nil
}
