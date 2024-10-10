// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: models/movie/movie.proto

package moviepb

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

type Movie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Year        int32    `protobuf:"varint,4,opt,name=year,proto3" json:"year,omitempty"`
	Country     string   `protobuf:"bytes,5,opt,name=country,proto3" json:"country,omitempty"`
	Genres      []string `protobuf:"bytes,6,rep,name=genres,proto3" json:"genres,omitempty"`
	PosterUrl   string   `protobuf:"bytes,7,opt,name=poster_url,json=posterUrl,proto3" json:"poster_url,omitempty"`
	Rating      float32  `protobuf:"fixed32,8,opt,name=rating,proto3" json:"rating,omitempty"`
}

func (x *Movie) Reset() {
	*x = Movie{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_movie_movie_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Movie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Movie) ProtoMessage() {}

func (x *Movie) ProtoReflect() protoreflect.Message {
	mi := &file_models_movie_movie_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Movie.ProtoReflect.Descriptor instead.
func (*Movie) Descriptor() ([]byte, []int) {
	return file_models_movie_movie_proto_rawDescGZIP(), []int{0}
}

func (x *Movie) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Movie) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Movie) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Movie) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *Movie) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Movie) GetGenres() []string {
	if x != nil {
		return x.Genres
	}
	return nil
}

func (x *Movie) GetPosterUrl() string {
	if x != nil {
		return x.PosterUrl
	}
	return ""
}

func (x *Movie) GetRating() float32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

type AddResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AddResponse) Reset() {
	*x = AddResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_movie_movie_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddResponse) ProtoMessage() {}

func (x *AddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_models_movie_movie_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddResponse.ProtoReflect.Descriptor instead.
func (*AddResponse) Descriptor() ([]byte, []int) {
	return file_models_movie_movie_proto_rawDescGZIP(), []int{1}
}

func (x *AddResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type AddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Movie *Movie `protobuf:"bytes,1,opt,name=movie,proto3" json:"movie,omitempty"`
}

func (x *AddRequest) Reset() {
	*x = AddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_movie_movie_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRequest) ProtoMessage() {}

func (x *AddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_models_movie_movie_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRequest.ProtoReflect.Descriptor instead.
func (*AddRequest) Descriptor() ([]byte, []int) {
	return file_models_movie_movie_proto_rawDescGZIP(), []int{2}
}

func (x *AddRequest) GetMovie() *Movie {
	if x != nil {
		return x.Movie
	}
	return nil
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_movie_movie_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_models_movie_movie_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_models_movie_movie_proto_rawDescGZIP(), []int{3}
}

func (x *GetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Movie *Movie `protobuf:"bytes,1,opt,name=movie,proto3" json:"movie,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_movie_movie_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_models_movie_movie_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_models_movie_movie_proto_rawDescGZIP(), []int{4}
}

func (x *GetResponse) GetMovie() *Movie {
	if x != nil {
		return x.Movie
	}
	return nil
}

type RemoveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RemoveRequest) Reset() {
	*x = RemoveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_movie_movie_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveRequest) ProtoMessage() {}

func (x *RemoveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_models_movie_movie_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveRequest.ProtoReflect.Descriptor instead.
func (*RemoveRequest) Descriptor() ([]byte, []int) {
	return file_models_movie_movie_proto_rawDescGZIP(), []int{5}
}

func (x *RemoveRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RemoveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *RemoveResponse) Reset() {
	*x = RemoveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_movie_movie_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveResponse) ProtoMessage() {}

func (x *RemoveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_models_movie_movie_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveResponse.ProtoReflect.Descriptor instead.
func (*RemoveResponse) Descriptor() ([]byte, []int) {
	return file_models_movie_movie_proto_rawDescGZIP(), []int{6}
}

func (x *RemoveResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Movie *Movie `protobuf:"bytes,1,opt,name=movie,proto3" json:"movie,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_movie_movie_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_models_movie_movie_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_models_movie_movie_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateRequest) GetMovie() *Movie {
	if x != nil {
		return x.Movie
	}
	return nil
}

type UpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UpdateResponse) Reset() {
	*x = UpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_movie_movie_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResponse) ProtoMessage() {}

func (x *UpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_models_movie_movie_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResponse.ProtoReflect.Descriptor instead.
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return file_models_movie_movie_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type GetManyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SortByYear   bool `protobuf:"varint,1,opt,name=sort_by_year,json=sortByYear,proto3" json:"sort_by_year,omitempty"`
	SortByTitle  bool `protobuf:"varint,2,opt,name=sort_by_title,json=sortByTitle,proto3" json:"sort_by_title,omitempty"`
	SortByRating bool `protobuf:"varint,3,opt,name=sort_by_rating,json=sortByRating,proto3" json:"sort_by_rating,omitempty"`
	IsAscending  bool `protobuf:"varint,4,opt,name=is_ascending,json=isAscending,proto3" json:"is_ascending,omitempty"`
}

func (x *GetManyRequest) Reset() {
	*x = GetManyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_movie_movie_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetManyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetManyRequest) ProtoMessage() {}

func (x *GetManyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_models_movie_movie_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetManyRequest.ProtoReflect.Descriptor instead.
func (*GetManyRequest) Descriptor() ([]byte, []int) {
	return file_models_movie_movie_proto_rawDescGZIP(), []int{9}
}

func (x *GetManyRequest) GetSortByYear() bool {
	if x != nil {
		return x.SortByYear
	}
	return false
}

func (x *GetManyRequest) GetSortByTitle() bool {
	if x != nil {
		return x.SortByTitle
	}
	return false
}

func (x *GetManyRequest) GetSortByRating() bool {
	if x != nil {
		return x.SortByRating
	}
	return false
}

func (x *GetManyRequest) GetIsAscending() bool {
	if x != nil {
		return x.IsAscending
	}
	return false
}

type GetManyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Movies []*Movie `protobuf:"bytes,1,rep,name=movies,proto3" json:"movies,omitempty"`
}

func (x *GetManyResponse) Reset() {
	*x = GetManyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_movie_movie_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetManyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetManyResponse) ProtoMessage() {}

func (x *GetManyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_models_movie_movie_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetManyResponse.ProtoReflect.Descriptor instead.
func (*GetManyResponse) Descriptor() ([]byte, []int) {
	return file_models_movie_movie_proto_rawDescGZIP(), []int{10}
}

func (x *GetManyResponse) GetMovies() []*Movie {
	if x != nil {
		return x.Movies
	}
	return nil
}

var File_models_movie_movie_proto protoreflect.FileDescriptor

var file_models_movie_movie_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2f, 0x6d,
	0x6f, 0x76, 0x69, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x76, 0x69,
	0x65, 0x22, 0xcc, 0x01, 0x0a, 0x05, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x6f, 0x73,
	0x74, 0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x6f, 0x73, 0x74, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x22, 0x1d, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x30, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a,
	0x05, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d,
	0x6f, 0x76, 0x69, 0x65, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x05, 0x6d, 0x6f, 0x76, 0x69,
	0x65, 0x22, 0x1c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x31, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22,
	0x0a, 0x05, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x05, 0x6d, 0x6f, 0x76,
	0x69, 0x65, 0x22, 0x1f, 0x0a, 0x0d, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x28, 0x0a, 0x0e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x33, 0x0a,
	0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22,
	0x0a, 0x05, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x05, 0x6d, 0x6f, 0x76,
	0x69, 0x65, 0x22, 0x28, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x9f, 0x01, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x20, 0x0a, 0x0c, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x62, 0x79, 0x5f, 0x79, 0x65, 0x61, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x59, 0x65, 0x61,
	0x72, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x62, 0x79, 0x5f, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x73, 0x6f, 0x72, 0x74, 0x42, 0x79,
	0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x62, 0x79,
	0x5f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x73,
	0x6f, 0x72, 0x74, 0x42, 0x79, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x21, 0x0a, 0x0c, 0x69,
	0x73, 0x5f, 0x61, 0x73, 0x63, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0b, 0x69, 0x73, 0x41, 0x73, 0x63, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x22, 0x37,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x24, 0x0a, 0x06, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52,
	0x06, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x32, 0x92, 0x02, 0x0a, 0x0c, 0x4d, 0x6f, 0x76, 0x69,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12,
	0x11, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x12, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x11, 0x2e,
	0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x12, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x06, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x14,
	0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x06, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6d, 0x6f,
	0x76, 0x69, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x38, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x12, 0x15, 0x2e,
	0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x4d, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1a, 0x5a, 0x18,
	0x67, 0x65, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65,
	0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_movie_movie_proto_rawDescOnce sync.Once
	file_models_movie_movie_proto_rawDescData = file_models_movie_movie_proto_rawDesc
)

func file_models_movie_movie_proto_rawDescGZIP() []byte {
	file_models_movie_movie_proto_rawDescOnce.Do(func() {
		file_models_movie_movie_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_movie_movie_proto_rawDescData)
	})
	return file_models_movie_movie_proto_rawDescData
}

var file_models_movie_movie_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_models_movie_movie_proto_goTypes = []any{
	(*Movie)(nil),           // 0: movie.Movie
	(*AddResponse)(nil),     // 1: movie.AddResponse
	(*AddRequest)(nil),      // 2: movie.AddRequest
	(*GetRequest)(nil),      // 3: movie.GetRequest
	(*GetResponse)(nil),     // 4: movie.GetResponse
	(*RemoveRequest)(nil),   // 5: movie.RemoveRequest
	(*RemoveResponse)(nil),  // 6: movie.RemoveResponse
	(*UpdateRequest)(nil),   // 7: movie.UpdateRequest
	(*UpdateResponse)(nil),  // 8: movie.UpdateResponse
	(*GetManyRequest)(nil),  // 9: movie.GetManyRequest
	(*GetManyResponse)(nil), // 10: movie.GetManyResponse
}
var file_models_movie_movie_proto_depIdxs = []int32{
	0,  // 0: movie.AddRequest.movie:type_name -> movie.Movie
	0,  // 1: movie.GetResponse.movie:type_name -> movie.Movie
	0,  // 2: movie.UpdateRequest.movie:type_name -> movie.Movie
	0,  // 3: movie.GetManyResponse.movies:type_name -> movie.Movie
	2,  // 4: movie.MovieService.Add:input_type -> movie.AddRequest
	3,  // 5: movie.MovieService.Get:input_type -> movie.GetRequest
	5,  // 6: movie.MovieService.Remove:input_type -> movie.RemoveRequest
	7,  // 7: movie.MovieService.Update:input_type -> movie.UpdateRequest
	9,  // 8: movie.MovieService.GetMany:input_type -> movie.GetManyRequest
	1,  // 9: movie.MovieService.Add:output_type -> movie.AddResponse
	4,  // 10: movie.MovieService.Get:output_type -> movie.GetResponse
	6,  // 11: movie.MovieService.Remove:output_type -> movie.RemoveResponse
	8,  // 12: movie.MovieService.Update:output_type -> movie.UpdateResponse
	10, // 13: movie.MovieService.GetMany:output_type -> movie.GetManyResponse
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_models_movie_movie_proto_init() }
func file_models_movie_movie_proto_init() {
	if File_models_movie_movie_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_models_movie_movie_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Movie); i {
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
		file_models_movie_movie_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AddResponse); i {
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
		file_models_movie_movie_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*AddRequest); i {
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
		file_models_movie_movie_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetRequest); i {
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
		file_models_movie_movie_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetResponse); i {
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
		file_models_movie_movie_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*RemoveRequest); i {
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
		file_models_movie_movie_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*RemoveResponse); i {
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
		file_models_movie_movie_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateRequest); i {
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
		file_models_movie_movie_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateResponse); i {
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
		file_models_movie_movie_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*GetManyRequest); i {
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
		file_models_movie_movie_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*GetManyResponse); i {
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
			RawDescriptor: file_models_movie_movie_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_models_movie_movie_proto_goTypes,
		DependencyIndexes: file_models_movie_movie_proto_depIdxs,
		MessageInfos:      file_models_movie_movie_proto_msgTypes,
	}.Build()
	File_models_movie_movie_proto = out.File
	file_models_movie_movie_proto_rawDesc = nil
	file_models_movie_movie_proto_goTypes = nil
	file_models_movie_movie_proto_depIdxs = nil
}
