// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: song/songsLibrary.proto

package songsLibraryv1

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

// GetSongs
type GetSongsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	GroupName   string `protobuf:"bytes,3,opt,name=group_name,json=groupName,proto3" json:"group_name,omitempty"`
	ReleaseDate string `protobuf:"bytes,4,opt,name=release_date,json=releaseDate,proto3" json:"release_date,omitempty"`
	Text        string `protobuf:"bytes,5,opt,name=text,proto3" json:"text,omitempty"`
	Link        string `protobuf:"bytes,6,opt,name=link,proto3" json:"link,omitempty"`
	Page        int64  `protobuf:"varint,7,opt,name=page,proto3" json:"page,omitempty"`
	PageSize    int64  `protobuf:"varint,8,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *GetSongsRequest) Reset() {
	*x = GetSongsRequest{}
	mi := &file_song_songsLibrary_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSongsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSongsRequest) ProtoMessage() {}

func (x *GetSongsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_song_songsLibrary_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSongsRequest.ProtoReflect.Descriptor instead.
func (*GetSongsRequest) Descriptor() ([]byte, []int) {
	return file_song_songsLibrary_proto_rawDescGZIP(), []int{0}
}

func (x *GetSongsRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetSongsRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetSongsRequest) GetGroupName() string {
	if x != nil {
		return x.GroupName
	}
	return ""
}

func (x *GetSongsRequest) GetReleaseDate() string {
	if x != nil {
		return x.ReleaseDate
	}
	return ""
}

func (x *GetSongsRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *GetSongsRequest) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *GetSongsRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetSongsRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type GetSongsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	AuthorId    string `protobuf:"bytes,3,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	AuthorName  string `protobuf:"bytes,4,opt,name=author_name,json=authorName,proto3" json:"author_name,omitempty"`
	ReleaseDate string `protobuf:"bytes,5,opt,name=release_date,json=releaseDate,proto3" json:"release_date,omitempty"`
	Text        string `protobuf:"bytes,6,opt,name=text,proto3" json:"text,omitempty"`
	Link        string `protobuf:"bytes,7,opt,name=link,proto3" json:"link,omitempty"`
}

func (x *GetSongsResponse) Reset() {
	*x = GetSongsResponse{}
	mi := &file_song_songsLibrary_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSongsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSongsResponse) ProtoMessage() {}

func (x *GetSongsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_song_songsLibrary_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSongsResponse.ProtoReflect.Descriptor instead.
func (*GetSongsResponse) Descriptor() ([]byte, []int) {
	return file_song_songsLibrary_proto_rawDescGZIP(), []int{1}
}

func (x *GetSongsResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetSongsResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetSongsResponse) GetAuthorId() string {
	if x != nil {
		return x.AuthorId
	}
	return ""
}

func (x *GetSongsResponse) GetAuthorName() string {
	if x != nil {
		return x.AuthorName
	}
	return ""
}

func (x *GetSongsResponse) GetReleaseDate() string {
	if x != nil {
		return x.ReleaseDate
	}
	return ""
}

func (x *GetSongsResponse) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *GetSongsResponse) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

type GetSongsResponseList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Songs []*GetSongsResponse `protobuf:"bytes,1,rep,name=songs,proto3" json:"songs,omitempty"`
}

func (x *GetSongsResponseList) Reset() {
	*x = GetSongsResponseList{}
	mi := &file_song_songsLibrary_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSongsResponseList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSongsResponseList) ProtoMessage() {}

func (x *GetSongsResponseList) ProtoReflect() protoreflect.Message {
	mi := &file_song_songsLibrary_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSongsResponseList.ProtoReflect.Descriptor instead.
func (*GetSongsResponseList) Descriptor() ([]byte, []int) {
	return file_song_songsLibrary_proto_rawDescGZIP(), []int{2}
}

func (x *GetSongsResponseList) GetSongs() []*GetSongsResponse {
	if x != nil {
		return x.Songs
	}
	return nil
}

type DeleteSongsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteSongsRequest) Reset() {
	*x = DeleteSongsRequest{}
	mi := &file_song_songsLibrary_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteSongsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSongsRequest) ProtoMessage() {}

func (x *DeleteSongsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_song_songsLibrary_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSongsRequest.ProtoReflect.Descriptor instead.
func (*DeleteSongsRequest) Descriptor() ([]byte, []int) {
	return file_song_songsLibrary_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteSongsRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteSongsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	AuthorId    string `protobuf:"bytes,3,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	AuthorName  string `protobuf:"bytes,4,opt,name=author_name,json=authorName,proto3" json:"author_name,omitempty"`
	ReleaseDate string `protobuf:"bytes,5,opt,name=release_date,json=releaseDate,proto3" json:"release_date,omitempty"`
	Text        string `protobuf:"bytes,6,opt,name=text,proto3" json:"text,omitempty"`
	Link        string `protobuf:"bytes,7,opt,name=link,proto3" json:"link,omitempty"`
}

func (x *DeleteSongsResponse) Reset() {
	*x = DeleteSongsResponse{}
	mi := &file_song_songsLibrary_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteSongsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSongsResponse) ProtoMessage() {}

func (x *DeleteSongsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_song_songsLibrary_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSongsResponse.ProtoReflect.Descriptor instead.
func (*DeleteSongsResponse) Descriptor() ([]byte, []int) {
	return file_song_songsLibrary_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteSongsResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeleteSongsResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeleteSongsResponse) GetAuthorId() string {
	if x != nil {
		return x.AuthorId
	}
	return ""
}

func (x *DeleteSongsResponse) GetAuthorName() string {
	if x != nil {
		return x.AuthorName
	}
	return ""
}

func (x *DeleteSongsResponse) GetReleaseDate() string {
	if x != nil {
		return x.ReleaseDate
	}
	return ""
}

func (x *DeleteSongsResponse) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *DeleteSongsResponse) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

var File_song_songsLibrary_proto protoreflect.FileDescriptor

var file_song_songsLibrary_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x6f, 0x6e, 0x67, 0x2f, 0x73, 0x6f, 0x6e, 0x67, 0x73, 0x4c, 0x69, 0x62, 0x72,
	0x61, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x6f, 0x6e, 0x67, 0x73,
	0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x22, 0xd0, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53,
	0x6f, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0xbf, 0x01, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x53, 0x6f, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x22, 0x4c, 0x0a, 0x14,
	0x47, 0x65, 0x74, 0x53, 0x6f, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x05, 0x73, 0x6f, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x6f, 0x6e, 0x67, 0x73, 0x4c, 0x69, 0x62, 0x72, 0x61,
	0x72, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x6f, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x05, 0x73, 0x6f, 0x6e, 0x67, 0x73, 0x22, 0x24, 0x0a, 0x12, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x53, 0x6f, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0xc2, 0x01, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x6f, 0x6e, 0x67, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x32, 0xa8, 0x01, 0x0a, 0x04, 0x53, 0x6f, 0x6e, 0x67, 0x12, 0x4d,
	0x0a, 0x08, 0x47, 0x65, 0x74, 0x53, 0x6f, 0x6e, 0x67, 0x73, 0x12, 0x1d, 0x2e, 0x73, 0x6f, 0x6e,
	0x67, 0x73, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x6f, 0x6e,
	0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73, 0x6f, 0x6e, 0x67,
	0x73, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x6f, 0x6e, 0x67,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x51, 0x0a,
	0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x6f, 0x6e, 0x67, 0x12, 0x20, 0x2e, 0x73, 0x6f,
	0x6e, 0x67, 0x73, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x53, 0x6f, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e,
	0x73, 0x6f, 0x6e, 0x67, 0x73, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x53, 0x6f, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x2d, 0x5a, 0x2b, 0x6b, 0x6f, 0x7a, 0x6c, 0x79, 0x61, 0x6b, 0x6f, 0x76, 0x73, 0x6b, 0x79,
	0x2e, 0x73, 0x6f, 0x6e, 0x67, 0x73, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x76, 0x31,
	0x3b, 0x73, 0x6f, 0x6e, 0x67, 0x73, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_song_songsLibrary_proto_rawDescOnce sync.Once
	file_song_songsLibrary_proto_rawDescData = file_song_songsLibrary_proto_rawDesc
)

func file_song_songsLibrary_proto_rawDescGZIP() []byte {
	file_song_songsLibrary_proto_rawDescOnce.Do(func() {
		file_song_songsLibrary_proto_rawDescData = protoimpl.X.CompressGZIP(file_song_songsLibrary_proto_rawDescData)
	})
	return file_song_songsLibrary_proto_rawDescData
}

var file_song_songsLibrary_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_song_songsLibrary_proto_goTypes = []any{
	(*GetSongsRequest)(nil),      // 0: songsLibrary.GetSongsRequest
	(*GetSongsResponse)(nil),     // 1: songsLibrary.GetSongsResponse
	(*GetSongsResponseList)(nil), // 2: songsLibrary.GetSongsResponseList
	(*DeleteSongsRequest)(nil),   // 3: songsLibrary.DeleteSongsRequest
	(*DeleteSongsResponse)(nil),  // 4: songsLibrary.DeleteSongsResponse
}
var file_song_songsLibrary_proto_depIdxs = []int32{
	1, // 0: songsLibrary.GetSongsResponseList.songs:type_name -> songsLibrary.GetSongsResponse
	0, // 1: songsLibrary.Song.GetSongs:input_type -> songsLibrary.GetSongsRequest
	3, // 2: songsLibrary.Song.DeleteSong:input_type -> songsLibrary.DeleteSongsRequest
	2, // 3: songsLibrary.Song.GetSongs:output_type -> songsLibrary.GetSongsResponseList
	4, // 4: songsLibrary.Song.DeleteSong:output_type -> songsLibrary.DeleteSongsResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_song_songsLibrary_proto_init() }
func file_song_songsLibrary_proto_init() {
	if File_song_songsLibrary_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_song_songsLibrary_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_song_songsLibrary_proto_goTypes,
		DependencyIndexes: file_song_songsLibrary_proto_depIdxs,
		MessageInfos:      file_song_songsLibrary_proto_msgTypes,
	}.Build()
	File_song_songsLibrary_proto = out.File
	file_song_songsLibrary_proto_rawDesc = nil
	file_song_songsLibrary_proto_goTypes = nil
	file_song_songsLibrary_proto_depIdxs = nil
}
