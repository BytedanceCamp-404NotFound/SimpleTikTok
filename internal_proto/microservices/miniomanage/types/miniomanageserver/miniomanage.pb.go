// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: proto/miniomanage.proto

// import "common.proto"; // 导入失败了

package miniomanageserver

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

type RpcErrorCode int32

const (
	RpcErrorCode_UNKNOW_ERROR RpcErrorCode = 0   // 未知错误, proto编译必须定义，禁止使用此值
	RpcErrorCode_STATUS_OK    RpcErrorCode = 200 // 正常
)

// Enum value maps for RpcErrorCode.
var (
	RpcErrorCode_name = map[int32]string{
		0:   "UNKNOW_ERROR",
		200: "STATUS_OK",
	}
	RpcErrorCode_value = map[string]int32{
		"UNKNOW_ERROR": 0,
		"STATUS_OK":    200,
	}
)

func (x RpcErrorCode) Enum() *RpcErrorCode {
	p := new(RpcErrorCode)
	*p = x
	return p
}

func (x RpcErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RpcErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_miniomanage_proto_enumTypes[0].Descriptor()
}

func (RpcErrorCode) Type() protoreflect.EnumType {
	return &file_proto_miniomanage_proto_enumTypes[0]
}

func (x RpcErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RpcErrorCode.Descriptor instead.
func (RpcErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_proto_miniomanage_proto_rawDescGZIP(), []int{0}
}

type PutFileUploaderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BucketName string `protobuf:"bytes,1,opt,name=bucketName,proto3" json:"bucketName,omitempty"`
	ObjectPre  string `protobuf:"bytes,2,opt,name=objectPre,proto3" json:"objectPre,omitempty"`
	FilePath   string `protobuf:"bytes,3,opt,name=filePath,proto3" json:"filePath,omitempty"`
}

func (x *PutFileUploaderRequest) Reset() {
	*x = PutFileUploaderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_miniomanage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutFileUploaderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutFileUploaderRequest) ProtoMessage() {}

func (x *PutFileUploaderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_miniomanage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutFileUploaderRequest.ProtoReflect.Descriptor instead.
func (*PutFileUploaderRequest) Descriptor() ([]byte, []int) {
	return file_proto_miniomanage_proto_rawDescGZIP(), []int{0}
}

func (x *PutFileUploaderRequest) GetBucketName() string {
	if x != nil {
		return x.BucketName
	}
	return ""
}

func (x *PutFileUploaderRequest) GetObjectPre() string {
	if x != nil {
		return x.ObjectPre
	}
	return ""
}

func (x *PutFileUploaderRequest) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

type PutFileUploaderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error          RpcErrorCode `protobuf:"varint,1,opt,name=error,proto3,enum=miniomanageserver.RpcErrorCode" json:"error,omitempty"`
	BucketFilepath string       `protobuf:"bytes,2,opt,name=bucket_filepath,json=bucketFilepath,proto3" json:"bucket_filepath,omitempty"`
}

func (x *PutFileUploaderResponse) Reset() {
	*x = PutFileUploaderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_miniomanage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutFileUploaderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutFileUploaderResponse) ProtoMessage() {}

func (x *PutFileUploaderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_miniomanage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutFileUploaderResponse.ProtoReflect.Descriptor instead.
func (*PutFileUploaderResponse) Descriptor() ([]byte, []int) {
	return file_proto_miniomanage_proto_rawDescGZIP(), []int{1}
}

func (x *PutFileUploaderResponse) GetError() RpcErrorCode {
	if x != nil {
		return x.Error
	}
	return RpcErrorCode_UNKNOW_ERROR
}

func (x *PutFileUploaderResponse) GetBucketFilepath() string {
	if x != nil {
		return x.BucketFilepath
	}
	return ""
}

type PutFileUploaderByteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BucketName string `protobuf:"bytes,1,opt,name=bucketName,proto3" json:"bucketName,omitempty"`
	ObjectPre  string `protobuf:"bytes,2,opt,name=objectPre,proto3" json:"objectPre,omitempty"`
	FilePath   string `protobuf:"bytes,3,opt,name=filePath,proto3" json:"filePath,omitempty"`
}

func (x *PutFileUploaderByteRequest) Reset() {
	*x = PutFileUploaderByteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_miniomanage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutFileUploaderByteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutFileUploaderByteRequest) ProtoMessage() {}

func (x *PutFileUploaderByteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_miniomanage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutFileUploaderByteRequest.ProtoReflect.Descriptor instead.
func (*PutFileUploaderByteRequest) Descriptor() ([]byte, []int) {
	return file_proto_miniomanage_proto_rawDescGZIP(), []int{2}
}

func (x *PutFileUploaderByteRequest) GetBucketName() string {
	if x != nil {
		return x.BucketName
	}
	return ""
}

func (x *PutFileUploaderByteRequest) GetObjectPre() string {
	if x != nil {
		return x.ObjectPre
	}
	return ""
}

func (x *PutFileUploaderByteRequest) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

type PutFileUploaderByteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error          RpcErrorCode `protobuf:"varint,1,opt,name=error,proto3,enum=miniomanageserver.RpcErrorCode" json:"error,omitempty"`
	BucketFilepath string       `protobuf:"bytes,2,opt,name=bucket_filepath,json=bucketFilepath,proto3" json:"bucket_filepath,omitempty"`
}

func (x *PutFileUploaderByteResponse) Reset() {
	*x = PutFileUploaderByteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_miniomanage_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutFileUploaderByteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutFileUploaderByteResponse) ProtoMessage() {}

func (x *PutFileUploaderByteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_miniomanage_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutFileUploaderByteResponse.ProtoReflect.Descriptor instead.
func (*PutFileUploaderByteResponse) Descriptor() ([]byte, []int) {
	return file_proto_miniomanage_proto_rawDescGZIP(), []int{3}
}

func (x *PutFileUploaderByteResponse) GetError() RpcErrorCode {
	if x != nil {
		return x.Error
	}
	return RpcErrorCode_UNKNOW_ERROR
}

func (x *PutFileUploaderByteResponse) GetBucketFilepath() string {
	if x != nil {
		return x.BucketFilepath
	}
	return ""
}

type GetFileDownloaderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// 用户名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// 用户性别
	Gender string `protobuf:"bytes,3,opt,name=gender,proto3" json:"gender,omitempty"`
}

func (x *GetFileDownloaderRequest) Reset() {
	*x = GetFileDownloaderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_miniomanage_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFileDownloaderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFileDownloaderRequest) ProtoMessage() {}

func (x *GetFileDownloaderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_miniomanage_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFileDownloaderRequest.ProtoReflect.Descriptor instead.
func (*GetFileDownloaderRequest) Descriptor() ([]byte, []int) {
	return file_proto_miniomanage_proto_rawDescGZIP(), []int{4}
}

func (x *GetFileDownloaderRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetFileDownloaderRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetFileDownloaderRequest) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

type GetFileDownloaderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error          RpcErrorCode `protobuf:"varint,1,opt,name=error,proto3,enum=miniomanageserver.RpcErrorCode" json:"error,omitempty"`
	BucketFilepath string       `protobuf:"bytes,2,opt,name=bucket_filepath,json=bucketFilepath,proto3" json:"bucket_filepath,omitempty"`
}

func (x *GetFileDownloaderResponse) Reset() {
	*x = GetFileDownloaderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_miniomanage_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFileDownloaderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFileDownloaderResponse) ProtoMessage() {}

func (x *GetFileDownloaderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_miniomanage_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFileDownloaderResponse.ProtoReflect.Descriptor instead.
func (*GetFileDownloaderResponse) Descriptor() ([]byte, []int) {
	return file_proto_miniomanage_proto_rawDescGZIP(), []int{5}
}

func (x *GetFileDownloaderResponse) GetError() RpcErrorCode {
	if x != nil {
		return x.Error
	}
	return RpcErrorCode_UNKNOW_ERROR
}

func (x *GetFileDownloaderResponse) GetBucketFilepath() string {
	if x != nil {
		return x.BucketFilepath
	}
	return ""
}

var File_proto_miniomanage_proto protoreflect.FileDescriptor

var file_proto_miniomanage_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6d, 0x69, 0x6e, 0x69, 0x6f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x72, 0x0a, 0x16,
	0x50, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x50, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x50, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68,
	0x22, 0x79, 0x0a, 0x17, 0x50, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x6d, 0x69, 0x6e,
	0x69, 0x6f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x52,
	0x70, 0x63, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x66, 0x69, 0x6c,
	0x65, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x62, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x70, 0x61, 0x74, 0x68, 0x22, 0x76, 0x0a, 0x1a, 0x50,
	0x75, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x42, 0x79,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x50, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x50, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50,
	0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50,
	0x61, 0x74, 0x68, 0x22, 0x7d, 0x0a, 0x1b, 0x50, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x42, 0x79, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x35, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1f, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x52, 0x70, 0x63, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x62, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x70, 0x61,
	0x74, 0x68, 0x22, 0x56, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x22, 0x7b, 0x0a, 0x19, 0x47, 0x65,
	0x74, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x52, 0x70, 0x63, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x27,
	0x0a, 0x0f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x46,
	0x69, 0x6c, 0x65, 0x70, 0x61, 0x74, 0x68, 0x2a, 0x30, 0x0a, 0x0c, 0x52, 0x70, 0x63, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x09, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x4f, 0x4b, 0x10, 0xc8, 0x01, 0x32, 0xe3, 0x02, 0x0a, 0x11, 0x4d, 0x69,
	0x6e, 0x69, 0x6f, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12,
	0x68, 0x0a, 0x0f, 0x50, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x65, 0x72, 0x12, 0x29, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e,
	0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x50, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x74, 0x0a, 0x13, 0x50, 0x75, 0x74,
	0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x42, 0x79, 0x74, 0x65,
	0x12, 0x2d, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x65, 0x72, 0x42, 0x79, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2e, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x50, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x65, 0x72, 0x42, 0x79, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x6e, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f,
	0x61, 0x64, 0x65, 0x72, 0x12, 0x2b, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65,
	0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2c, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x15, 0x5a, 0x13, 0x2e, 0x2f, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_miniomanage_proto_rawDescOnce sync.Once
	file_proto_miniomanage_proto_rawDescData = file_proto_miniomanage_proto_rawDesc
)

func file_proto_miniomanage_proto_rawDescGZIP() []byte {
	file_proto_miniomanage_proto_rawDescOnce.Do(func() {
		file_proto_miniomanage_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_miniomanage_proto_rawDescData)
	})
	return file_proto_miniomanage_proto_rawDescData
}

var file_proto_miniomanage_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_miniomanage_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_miniomanage_proto_goTypes = []interface{}{
	(RpcErrorCode)(0),                   // 0: miniomanageserver.RpcErrorCode
	(*PutFileUploaderRequest)(nil),      // 1: miniomanageserver.PutFileUploaderRequest
	(*PutFileUploaderResponse)(nil),     // 2: miniomanageserver.PutFileUploaderResponse
	(*PutFileUploaderByteRequest)(nil),  // 3: miniomanageserver.PutFileUploaderByteRequest
	(*PutFileUploaderByteResponse)(nil), // 4: miniomanageserver.PutFileUploaderByteResponse
	(*GetFileDownloaderRequest)(nil),    // 5: miniomanageserver.GetFileDownloaderRequest
	(*GetFileDownloaderResponse)(nil),   // 6: miniomanageserver.GetFileDownloaderResponse
}
var file_proto_miniomanage_proto_depIdxs = []int32{
	0, // 0: miniomanageserver.PutFileUploaderResponse.error:type_name -> miniomanageserver.RpcErrorCode
	0, // 1: miniomanageserver.PutFileUploaderByteResponse.error:type_name -> miniomanageserver.RpcErrorCode
	0, // 2: miniomanageserver.GetFileDownloaderResponse.error:type_name -> miniomanageserver.RpcErrorCode
	1, // 3: miniomanageserver.MinioManageServer.PutFileUploader:input_type -> miniomanageserver.PutFileUploaderRequest
	3, // 4: miniomanageserver.MinioManageServer.PutFileUploaderByte:input_type -> miniomanageserver.PutFileUploaderByteRequest
	5, // 5: miniomanageserver.MinioManageServer.GetFileDownloader:input_type -> miniomanageserver.GetFileDownloaderRequest
	2, // 6: miniomanageserver.MinioManageServer.PutFileUploader:output_type -> miniomanageserver.PutFileUploaderResponse
	4, // 7: miniomanageserver.MinioManageServer.PutFileUploaderByte:output_type -> miniomanageserver.PutFileUploaderByteResponse
	6, // 8: miniomanageserver.MinioManageServer.GetFileDownloader:output_type -> miniomanageserver.GetFileDownloaderResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_miniomanage_proto_init() }
func file_proto_miniomanage_proto_init() {
	if File_proto_miniomanage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_miniomanage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutFileUploaderRequest); i {
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
		file_proto_miniomanage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutFileUploaderResponse); i {
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
		file_proto_miniomanage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutFileUploaderByteRequest); i {
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
		file_proto_miniomanage_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutFileUploaderByteResponse); i {
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
		file_proto_miniomanage_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFileDownloaderRequest); i {
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
		file_proto_miniomanage_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFileDownloaderResponse); i {
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
			RawDescriptor: file_proto_miniomanage_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_miniomanage_proto_goTypes,
		DependencyIndexes: file_proto_miniomanage_proto_depIdxs,
		EnumInfos:         file_proto_miniomanage_proto_enumTypes,
		MessageInfos:      file_proto_miniomanage_proto_msgTypes,
	}.Build()
	File_proto_miniomanage_proto = out.File
	file_proto_miniomanage_proto_rawDesc = nil
	file_proto_miniomanage_proto_goTypes = nil
	file_proto_miniomanage_proto_depIdxs = nil
}
