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

	BucketFilepath string `protobuf:"bytes,2,opt,name=bucket_filepath,json=bucketFilepath,proto3" json:"bucket_filepath,omitempty"`
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

	BucketFilepath string `protobuf:"bytes,2,opt,name=bucket_filepath,json=bucketFilepath,proto3" json:"bucket_filepath,omitempty"`
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

type GetFileDownloaderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BucketFilepath string `protobuf:"bytes,2,opt,name=bucket_filepath,json=bucketFilepath,proto3" json:"bucket_filepath,omitempty"`
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

func (x *GetFileDownloaderResponse) GetBucketFilepath() string {
	if x != nil {
		return x.BucketFilepath
	}
	return ""
}

// 获取Minio视频播放的URLRequest
type GetPlayUrlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayUrl string `protobuf:"bytes,1,opt,name=playUrl,proto3" json:"playUrl,omitempty"` // 想要转换的URL
}

func (x *GetPlayUrlRequest) Reset() {
	*x = GetPlayUrlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_miniomanage_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayUrlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayUrlRequest) ProtoMessage() {}

func (x *GetPlayUrlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_miniomanage_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayUrlRequest.ProtoReflect.Descriptor instead.
func (*GetPlayUrlRequest) Descriptor() ([]byte, []int) {
	return file_proto_miniomanage_proto_rawDescGZIP(), []int{6}
}

func (x *GetPlayUrlRequest) GetPlayUrl() string {
	if x != nil {
		return x.PlayUrl
	}
	return ""
}

// 获取Minio视频播放的URLResponse
type GetPlayUrlResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResPlayUrl string `protobuf:"bytes,2,opt,name=resPlayUrl,proto3" json:"resPlayUrl,omitempty"` // 实际的URL
}

func (x *GetPlayUrlResponse) Reset() {
	*x = GetPlayUrlResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_miniomanage_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayUrlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayUrlResponse) ProtoMessage() {}

func (x *GetPlayUrlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_miniomanage_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayUrlResponse.ProtoReflect.Descriptor instead.
func (*GetPlayUrlResponse) Descriptor() ([]byte, []int) {
	return file_proto_miniomanage_proto_rawDescGZIP(), []int{7}
}

func (x *GetPlayUrlResponse) GetResPlayUrl() string {
	if x != nil {
		return x.ResPlayUrl
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
	0x22, 0x42, 0x0a, 0x17, 0x50, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x62,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65,
	0x70, 0x61, 0x74, 0x68, 0x22, 0x76, 0x0a, 0x1a, 0x50, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x42, 0x79, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x50, 0x72, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x50, 0x72, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x22, 0x46, 0x0a, 0x1b,
	0x50, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x42,
	0x79, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x62,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65,
	0x70, 0x61, 0x74, 0x68, 0x22, 0x1a, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x44,
	0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x44, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c,
	0x6f, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a,
	0x0f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x46, 0x69,
	0x6c, 0x65, 0x70, 0x61, 0x74, 0x68, 0x22, 0x2d, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61,
	0x79, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70,
	0x6c, 0x61, 0x79, 0x55, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6c,
	0x61, 0x79, 0x55, 0x72, 0x6c, 0x22, 0x34, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79,
	0x55, 0x72, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x72,
	0x65, 0x73, 0x50, 0x6c, 0x61, 0x79, 0x55, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x72, 0x65, 0x73, 0x50, 0x6c, 0x61, 0x79, 0x55, 0x72, 0x6c, 0x32, 0xbe, 0x03, 0x0a, 0x11,
	0x4d, 0x69, 0x6e, 0x69, 0x6f, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x12, 0x68, 0x0a, 0x0f, 0x50, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x65, 0x72, 0x12, 0x29, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x65,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2a, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x50, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x74, 0x0a, 0x13, 0x50,
	0x75, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x42, 0x79,
	0x74, 0x65, 0x12, 0x2d, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x42, 0x79, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2e, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x65, 0x72, 0x42, 0x79, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x6e, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x6f, 0x77, 0x6e,
	0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x12, 0x2b, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69,
	0x6c, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x44,
	0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x59, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x55, 0x72, 0x6c, 0x12,
	0x24, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x55, 0x72, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61,
	0x79, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x15, 0x5a, 0x13,
	0x2e, 0x2f, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_proto_miniomanage_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_miniomanage_proto_goTypes = []interface{}{
	(*PutFileUploaderRequest)(nil),      // 0: miniomanageserver.PutFileUploaderRequest
	(*PutFileUploaderResponse)(nil),     // 1: miniomanageserver.PutFileUploaderResponse
	(*PutFileUploaderByteRequest)(nil),  // 2: miniomanageserver.PutFileUploaderByteRequest
	(*PutFileUploaderByteResponse)(nil), // 3: miniomanageserver.PutFileUploaderByteResponse
	(*GetFileDownloaderRequest)(nil),    // 4: miniomanageserver.GetFileDownloaderRequest
	(*GetFileDownloaderResponse)(nil),   // 5: miniomanageserver.GetFileDownloaderResponse
	(*GetPlayUrlRequest)(nil),           // 6: miniomanageserver.GetPlayUrlRequest
	(*GetPlayUrlResponse)(nil),          // 7: miniomanageserver.GetPlayUrlResponse
}
var file_proto_miniomanage_proto_depIdxs = []int32{
	0, // 0: miniomanageserver.MinioManageServer.PutFileUploader:input_type -> miniomanageserver.PutFileUploaderRequest
	2, // 1: miniomanageserver.MinioManageServer.PutFileUploaderByte:input_type -> miniomanageserver.PutFileUploaderByteRequest
	4, // 2: miniomanageserver.MinioManageServer.GetFileDownloader:input_type -> miniomanageserver.GetFileDownloaderRequest
	6, // 3: miniomanageserver.MinioManageServer.GetPlayUrl:input_type -> miniomanageserver.GetPlayUrlRequest
	1, // 4: miniomanageserver.MinioManageServer.PutFileUploader:output_type -> miniomanageserver.PutFileUploaderResponse
	3, // 5: miniomanageserver.MinioManageServer.PutFileUploaderByte:output_type -> miniomanageserver.PutFileUploaderByteResponse
	5, // 6: miniomanageserver.MinioManageServer.GetFileDownloader:output_type -> miniomanageserver.GetFileDownloaderResponse
	7, // 7: miniomanageserver.MinioManageServer.GetPlayUrl:output_type -> miniomanageserver.GetPlayUrlResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
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
		file_proto_miniomanage_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayUrlRequest); i {
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
		file_proto_miniomanage_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayUrlResponse); i {
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
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_miniomanage_proto_goTypes,
		DependencyIndexes: file_proto_miniomanage_proto_depIdxs,
		MessageInfos:      file_proto_miniomanage_proto_msgTypes,
	}.Build()
	File_proto_miniomanage_proto = out.File
	file_proto_miniomanage_proto_rawDesc = nil
	file_proto_miniomanage_proto_goTypes = nil
	file_proto_miniomanage_proto_depIdxs = nil
}
