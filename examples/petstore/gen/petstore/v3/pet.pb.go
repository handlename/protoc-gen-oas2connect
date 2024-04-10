// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: petstore/v3/pet.proto

package petstorev3

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type UpdatePetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Category  *Category `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
	PhotoUrls []string  `protobuf:"bytes,4,rep,name=photo_urls,json=photoUrls,proto3" json:"photo_urls,omitempty"`
	Tags      []*Tag    `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`
	Status    string    `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UpdatePetRequest) Reset() {
	*x = UpdatePetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petstore_v3_pet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePetRequest) ProtoMessage() {}

func (x *UpdatePetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_petstore_v3_pet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePetRequest.ProtoReflect.Descriptor instead.
func (*UpdatePetRequest) Descriptor() ([]byte, []int) {
	return file_petstore_v3_pet_proto_rawDescGZIP(), []int{0}
}

func (x *UpdatePetRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdatePetRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdatePetRequest) GetCategory() *Category {
	if x != nil {
		return x.Category
	}
	return nil
}

func (x *UpdatePetRequest) GetPhotoUrls() []string {
	if x != nil {
		return x.PhotoUrls
	}
	return nil
}

func (x *UpdatePetRequest) GetTags() []*Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *UpdatePetRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type UpdatePetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pet *Pet `protobuf:"bytes,1,opt,name=pet,proto3" json:"pet,omitempty"`
}

func (x *UpdatePetResponse) Reset() {
	*x = UpdatePetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petstore_v3_pet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePetResponse) ProtoMessage() {}

func (x *UpdatePetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_petstore_v3_pet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePetResponse.ProtoReflect.Descriptor instead.
func (*UpdatePetResponse) Descriptor() ([]byte, []int) {
	return file_petstore_v3_pet_proto_rawDescGZIP(), []int{1}
}

func (x *UpdatePetResponse) GetPet() *Pet {
	if x != nil {
		return x.Pet
	}
	return nil
}

type AddPetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddPetRequest) Reset() {
	*x = AddPetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petstore_v3_pet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPetRequest) ProtoMessage() {}

func (x *AddPetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_petstore_v3_pet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPetRequest.ProtoReflect.Descriptor instead.
func (*AddPetRequest) Descriptor() ([]byte, []int) {
	return file_petstore_v3_pet_proto_rawDescGZIP(), []int{2}
}

type AddPetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pet *Pet `protobuf:"bytes,1,opt,name=pet,proto3" json:"pet,omitempty"`
}

func (x *AddPetResponse) Reset() {
	*x = AddPetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petstore_v3_pet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPetResponse) ProtoMessage() {}

func (x *AddPetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_petstore_v3_pet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPetResponse.ProtoReflect.Descriptor instead.
func (*AddPetResponse) Descriptor() ([]byte, []int) {
	return file_petstore_v3_pet_proto_rawDescGZIP(), []int{3}
}

func (x *AddPetResponse) GetPet() *Pet {
	if x != nil {
		return x.Pet
	}
	return nil
}

type FindPetsByStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *FindPetsByStatusRequest) Reset() {
	*x = FindPetsByStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petstore_v3_pet_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindPetsByStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindPetsByStatusRequest) ProtoMessage() {}

func (x *FindPetsByStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_petstore_v3_pet_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindPetsByStatusRequest.ProtoReflect.Descriptor instead.
func (*FindPetsByStatusRequest) Descriptor() ([]byte, []int) {
	return file_petstore_v3_pet_proto_rawDescGZIP(), []int{4}
}

func (x *FindPetsByStatusRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type FindPetsByStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pets []*Pet `protobuf:"bytes,1,rep,name=pets,proto3" json:"pets,omitempty"`
}

func (x *FindPetsByStatusResponse) Reset() {
	*x = FindPetsByStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petstore_v3_pet_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindPetsByStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindPetsByStatusResponse) ProtoMessage() {}

func (x *FindPetsByStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_petstore_v3_pet_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindPetsByStatusResponse.ProtoReflect.Descriptor instead.
func (*FindPetsByStatusResponse) Descriptor() ([]byte, []int) {
	return file_petstore_v3_pet_proto_rawDescGZIP(), []int{5}
}

func (x *FindPetsByStatusResponse) GetPets() []*Pet {
	if x != nil {
		return x.Pets
	}
	return nil
}

type FindPetsByTagsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tags []string `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *FindPetsByTagsRequest) Reset() {
	*x = FindPetsByTagsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petstore_v3_pet_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindPetsByTagsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindPetsByTagsRequest) ProtoMessage() {}

func (x *FindPetsByTagsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_petstore_v3_pet_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindPetsByTagsRequest.ProtoReflect.Descriptor instead.
func (*FindPetsByTagsRequest) Descriptor() ([]byte, []int) {
	return file_petstore_v3_pet_proto_rawDescGZIP(), []int{6}
}

func (x *FindPetsByTagsRequest) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type FindPetsByTagsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pets []*Pet `protobuf:"bytes,1,rep,name=pets,proto3" json:"pets,omitempty"`
}

func (x *FindPetsByTagsResponse) Reset() {
	*x = FindPetsByTagsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petstore_v3_pet_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindPetsByTagsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindPetsByTagsResponse) ProtoMessage() {}

func (x *FindPetsByTagsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_petstore_v3_pet_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindPetsByTagsResponse.ProtoReflect.Descriptor instead.
func (*FindPetsByTagsResponse) Descriptor() ([]byte, []int) {
	return file_petstore_v3_pet_proto_rawDescGZIP(), []int{7}
}

func (x *FindPetsByTagsResponse) GetPets() []*Pet {
	if x != nil {
		return x.Pets
	}
	return nil
}

type FindPetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PetId int64 `protobuf:"varint,1,opt,name=pet_id,json=petId,proto3" json:"pet_id,omitempty"`
}

func (x *FindPetRequest) Reset() {
	*x = FindPetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petstore_v3_pet_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindPetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindPetRequest) ProtoMessage() {}

func (x *FindPetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_petstore_v3_pet_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindPetRequest.ProtoReflect.Descriptor instead.
func (*FindPetRequest) Descriptor() ([]byte, []int) {
	return file_petstore_v3_pet_proto_rawDescGZIP(), []int{8}
}

func (x *FindPetRequest) GetPetId() int64 {
	if x != nil {
		return x.PetId
	}
	return 0
}

type FindPetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pet *Pet `protobuf:"bytes,1,opt,name=pet,proto3" json:"pet,omitempty"`
}

func (x *FindPetResponse) Reset() {
	*x = FindPetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petstore_v3_pet_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindPetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindPetResponse) ProtoMessage() {}

func (x *FindPetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_petstore_v3_pet_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindPetResponse.ProtoReflect.Descriptor instead.
func (*FindPetResponse) Descriptor() ([]byte, []int) {
	return file_petstore_v3_pet_proto_rawDescGZIP(), []int{9}
}

func (x *FindPetResponse) GetPet() *Pet {
	if x != nil {
		return x.Pet
	}
	return nil
}

type DeletePetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PetId int64 `protobuf:"varint,1,opt,name=pet_id,json=petId,proto3" json:"pet_id,omitempty"`
}

func (x *DeletePetRequest) Reset() {
	*x = DeletePetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petstore_v3_pet_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePetRequest) ProtoMessage() {}

func (x *DeletePetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_petstore_v3_pet_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePetRequest.ProtoReflect.Descriptor instead.
func (*DeletePetRequest) Descriptor() ([]byte, []int) {
	return file_petstore_v3_pet_proto_rawDescGZIP(), []int{10}
}

func (x *DeletePetRequest) GetPetId() int64 {
	if x != nil {
		return x.PetId
	}
	return 0
}

type DeletePetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeletePetResponse) Reset() {
	*x = DeletePetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petstore_v3_pet_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePetResponse) ProtoMessage() {}

func (x *DeletePetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_petstore_v3_pet_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePetResponse.ProtoReflect.Descriptor instead.
func (*DeletePetResponse) Descriptor() ([]byte, []int) {
	return file_petstore_v3_pet_proto_rawDescGZIP(), []int{11}
}

type Pet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Category  *Category `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
	PhotoUrls []string  `protobuf:"bytes,4,rep,name=photo_urls,json=photoUrls,proto3" json:"photo_urls,omitempty"`
	Tags      []*Tag    `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`
	Status    string    `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Pet) Reset() {
	*x = Pet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petstore_v3_pet_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pet) ProtoMessage() {}

func (x *Pet) ProtoReflect() protoreflect.Message {
	mi := &file_petstore_v3_pet_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pet.ProtoReflect.Descriptor instead.
func (*Pet) Descriptor() ([]byte, []int) {
	return file_petstore_v3_pet_proto_rawDescGZIP(), []int{12}
}

func (x *Pet) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Pet) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Pet) GetCategory() *Category {
	if x != nil {
		return x.Category
	}
	return nil
}

func (x *Pet) GetPhotoUrls() []string {
	if x != nil {
		return x.PhotoUrls
	}
	return nil
}

func (x *Pet) GetTags() []*Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Pet) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type Category struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Category) Reset() {
	*x = Category{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petstore_v3_pet_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Category) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Category) ProtoMessage() {}

func (x *Category) ProtoReflect() protoreflect.Message {
	mi := &file_petstore_v3_pet_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Category.ProtoReflect.Descriptor instead.
func (*Category) Descriptor() ([]byte, []int) {
	return file_petstore_v3_pet_proto_rawDescGZIP(), []int{13}
}

func (x *Category) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Category) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Tag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Tag) Reset() {
	*x = Tag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petstore_v3_pet_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tag) ProtoMessage() {}

func (x *Tag) ProtoReflect() protoreflect.Message {
	mi := &file_petstore_v3_pet_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tag.ProtoReflect.Descriptor instead.
func (*Tag) Descriptor() ([]byte, []int) {
	return file_petstore_v3_pet_proto_rawDescGZIP(), []int{14}
}

func (x *Tag) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Tag) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_petstore_v3_pet_proto protoreflect.FileDescriptor

var file_petstore_v3_pet_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x70, 0x65,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x33, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xc6, 0x01, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x08, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x55, 0x72, 0x6c, 0x73, 0x12, 0x24, 0x0a,
	0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x65,
	0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x37, 0x0a, 0x11, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x22, 0x0a, 0x03, 0x70, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x65, 0x74, 0x52,
	0x03, 0x70, 0x65, 0x74, 0x22, 0x0f, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x50, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x34, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x50, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x03, 0x70, 0x65, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x33, 0x2e, 0x50, 0x65, 0x74, 0x52, 0x03, 0x70, 0x65, 0x74, 0x22, 0x31, 0x0a, 0x17, 0x46,
	0x69, 0x6e, 0x64, 0x50, 0x65, 0x74, 0x73, 0x42, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x40,
	0x0a, 0x18, 0x46, 0x69, 0x6e, 0x64, 0x50, 0x65, 0x74, 0x73, 0x42, 0x79, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x70, 0x65,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x65, 0x74, 0x52, 0x04, 0x70, 0x65, 0x74, 0x73,
	0x22, 0x2b, 0x0a, 0x15, 0x46, 0x69, 0x6e, 0x64, 0x50, 0x65, 0x74, 0x73, 0x42, 0x79, 0x54, 0x61,
	0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x3e, 0x0a,
	0x16, 0x46, 0x69, 0x6e, 0x64, 0x50, 0x65, 0x74, 0x73, 0x42, 0x79, 0x54, 0x61, 0x67, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x70, 0x65, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x76, 0x33, 0x2e, 0x50, 0x65, 0x74, 0x52, 0x04, 0x70, 0x65, 0x74, 0x73, 0x22, 0x27, 0x0a,
	0x0e, 0x46, 0x69, 0x6e, 0x64, 0x50, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x15, 0x0a, 0x06, 0x70, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x70, 0x65, 0x74, 0x49, 0x64, 0x22, 0x35, 0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x64, 0x50, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x03, 0x70, 0x65, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x65, 0x74, 0x52, 0x03, 0x70, 0x65, 0x74, 0x22, 0x29, 0x0a,
	0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x15, 0x0a, 0x06, 0x70, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x70, 0x65, 0x74, 0x49, 0x64, 0x22, 0x13, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x50, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xb9, 0x01,
	0x0a, 0x03, 0x50, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x08, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x65,
	0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x68, 0x6f, 0x74, 0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x55, 0x72, 0x6c, 0x73, 0x12, 0x24, 0x0a, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x65, 0x74, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x04, 0x74, 0x61, 0x67,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x2e, 0x0a, 0x08, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x29, 0x0a, 0x03, 0x54, 0x61, 0x67,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x32, 0xed, 0x04, 0x0a, 0x0a, 0x50, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x5b, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x74,
	0x12, 0x1d, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x0f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x09, 0x3a, 0x01, 0x2a, 0x1a, 0x04, 0x2f, 0x70, 0x65, 0x74,
	0x12, 0x52, 0x0a, 0x06, 0x41, 0x64, 0x64, 0x50, 0x65, 0x74, 0x12, 0x1a, 0x2e, 0x70, 0x65, 0x74,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x33, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x0f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x09, 0x3a, 0x01, 0x2a, 0x22, 0x04,
	0x2f, 0x70, 0x65, 0x74, 0x12, 0x7a, 0x0a, 0x10, 0x46, 0x69, 0x6e, 0x64, 0x50, 0x65, 0x74, 0x73,
	0x42, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x24, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x50, 0x65, 0x74, 0x73, 0x42,
	0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25,
	0x2e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x46, 0x69, 0x6e,
	0x64, 0x50, 0x65, 0x74, 0x73, 0x42, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f,
	0x70, 0x65, 0x74, 0x2f, 0x66, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x72, 0x0a, 0x0e, 0x46, 0x69, 0x6e, 0x64, 0x50, 0x65, 0x74, 0x73, 0x42, 0x79, 0x54, 0x61,
	0x67, 0x73, 0x12, 0x22, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33,
	0x2e, 0x46, 0x69, 0x6e, 0x64, 0x50, 0x65, 0x74, 0x73, 0x42, 0x79, 0x54, 0x61, 0x67, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x33, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x50, 0x65, 0x74, 0x73, 0x42, 0x79, 0x54,
	0x61, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x70, 0x65, 0x74, 0x2f, 0x66, 0x69, 0x6e, 0x64, 0x42, 0x79,
	0x54, 0x61, 0x67, 0x73, 0x12, 0x5b, 0x0a, 0x07, 0x46, 0x69, 0x6e, 0x64, 0x50, 0x65, 0x74, 0x12,
	0x1b, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x46, 0x69,
	0x6e, 0x64, 0x50, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70,
	0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x50,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x70, 0x65, 0x74, 0x2f, 0x7b, 0x70, 0x65, 0x74, 0x5f, 0x69, 0x64,
	0x7d, 0x12, 0x61, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x65, 0x74, 0x12, 0x1d,
	0x2e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x50, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x50, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x2a, 0x0d, 0x2f, 0x70, 0x65, 0x74, 0x2f, 0x7b, 0x70, 0x65, 0x74,
	0x5f, 0x69, 0x64, 0x7d, 0x42, 0x25, 0x5a, 0x23, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33,
	0x3b, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_petstore_v3_pet_proto_rawDescOnce sync.Once
	file_petstore_v3_pet_proto_rawDescData = file_petstore_v3_pet_proto_rawDesc
)

func file_petstore_v3_pet_proto_rawDescGZIP() []byte {
	file_petstore_v3_pet_proto_rawDescOnce.Do(func() {
		file_petstore_v3_pet_proto_rawDescData = protoimpl.X.CompressGZIP(file_petstore_v3_pet_proto_rawDescData)
	})
	return file_petstore_v3_pet_proto_rawDescData
}

var file_petstore_v3_pet_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_petstore_v3_pet_proto_goTypes = []interface{}{
	(*UpdatePetRequest)(nil),         // 0: petstore.v3.UpdatePetRequest
	(*UpdatePetResponse)(nil),        // 1: petstore.v3.UpdatePetResponse
	(*AddPetRequest)(nil),            // 2: petstore.v3.AddPetRequest
	(*AddPetResponse)(nil),           // 3: petstore.v3.AddPetResponse
	(*FindPetsByStatusRequest)(nil),  // 4: petstore.v3.FindPetsByStatusRequest
	(*FindPetsByStatusResponse)(nil), // 5: petstore.v3.FindPetsByStatusResponse
	(*FindPetsByTagsRequest)(nil),    // 6: petstore.v3.FindPetsByTagsRequest
	(*FindPetsByTagsResponse)(nil),   // 7: petstore.v3.FindPetsByTagsResponse
	(*FindPetRequest)(nil),           // 8: petstore.v3.FindPetRequest
	(*FindPetResponse)(nil),          // 9: petstore.v3.FindPetResponse
	(*DeletePetRequest)(nil),         // 10: petstore.v3.DeletePetRequest
	(*DeletePetResponse)(nil),        // 11: petstore.v3.DeletePetResponse
	(*Pet)(nil),                      // 12: petstore.v3.Pet
	(*Category)(nil),                 // 13: petstore.v3.Category
	(*Tag)(nil),                      // 14: petstore.v3.Tag
}
var file_petstore_v3_pet_proto_depIdxs = []int32{
	13, // 0: petstore.v3.UpdatePetRequest.category:type_name -> petstore.v3.Category
	14, // 1: petstore.v3.UpdatePetRequest.tags:type_name -> petstore.v3.Tag
	12, // 2: petstore.v3.UpdatePetResponse.pet:type_name -> petstore.v3.Pet
	12, // 3: petstore.v3.AddPetResponse.pet:type_name -> petstore.v3.Pet
	12, // 4: petstore.v3.FindPetsByStatusResponse.pets:type_name -> petstore.v3.Pet
	12, // 5: petstore.v3.FindPetsByTagsResponse.pets:type_name -> petstore.v3.Pet
	12, // 6: petstore.v3.FindPetResponse.pet:type_name -> petstore.v3.Pet
	13, // 7: petstore.v3.Pet.category:type_name -> petstore.v3.Category
	14, // 8: petstore.v3.Pet.tags:type_name -> petstore.v3.Tag
	0,  // 9: petstore.v3.PetService.UpdatePet:input_type -> petstore.v3.UpdatePetRequest
	2,  // 10: petstore.v3.PetService.AddPet:input_type -> petstore.v3.AddPetRequest
	4,  // 11: petstore.v3.PetService.FindPetsByStatus:input_type -> petstore.v3.FindPetsByStatusRequest
	6,  // 12: petstore.v3.PetService.FindPetsByTags:input_type -> petstore.v3.FindPetsByTagsRequest
	8,  // 13: petstore.v3.PetService.FindPet:input_type -> petstore.v3.FindPetRequest
	10, // 14: petstore.v3.PetService.DeletePet:input_type -> petstore.v3.DeletePetRequest
	1,  // 15: petstore.v3.PetService.UpdatePet:output_type -> petstore.v3.UpdatePetResponse
	3,  // 16: petstore.v3.PetService.AddPet:output_type -> petstore.v3.AddPetResponse
	5,  // 17: petstore.v3.PetService.FindPetsByStatus:output_type -> petstore.v3.FindPetsByStatusResponse
	7,  // 18: petstore.v3.PetService.FindPetsByTags:output_type -> petstore.v3.FindPetsByTagsResponse
	9,  // 19: petstore.v3.PetService.FindPet:output_type -> petstore.v3.FindPetResponse
	11, // 20: petstore.v3.PetService.DeletePet:output_type -> petstore.v3.DeletePetResponse
	15, // [15:21] is the sub-list for method output_type
	9,  // [9:15] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_petstore_v3_pet_proto_init() }
func file_petstore_v3_pet_proto_init() {
	if File_petstore_v3_pet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_petstore_v3_pet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePetRequest); i {
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
		file_petstore_v3_pet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePetResponse); i {
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
		file_petstore_v3_pet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPetRequest); i {
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
		file_petstore_v3_pet_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPetResponse); i {
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
		file_petstore_v3_pet_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindPetsByStatusRequest); i {
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
		file_petstore_v3_pet_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindPetsByStatusResponse); i {
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
		file_petstore_v3_pet_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindPetsByTagsRequest); i {
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
		file_petstore_v3_pet_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindPetsByTagsResponse); i {
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
		file_petstore_v3_pet_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindPetRequest); i {
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
		file_petstore_v3_pet_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindPetResponse); i {
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
		file_petstore_v3_pet_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePetRequest); i {
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
		file_petstore_v3_pet_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePetResponse); i {
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
		file_petstore_v3_pet_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pet); i {
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
		file_petstore_v3_pet_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Category); i {
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
		file_petstore_v3_pet_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tag); i {
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
			RawDescriptor: file_petstore_v3_pet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_petstore_v3_pet_proto_goTypes,
		DependencyIndexes: file_petstore_v3_pet_proto_depIdxs,
		MessageInfos:      file_petstore_v3_pet_proto_msgTypes,
	}.Build()
	File_petstore_v3_pet_proto = out.File
	file_petstore_v3_pet_proto_rawDesc = nil
	file_petstore_v3_pet_proto_goTypes = nil
	file_petstore_v3_pet_proto_depIdxs = nil
}
