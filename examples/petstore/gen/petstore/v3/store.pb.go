// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: petstore/v3/store.proto

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

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	PetId    int64  `protobuf:"varint,2,opt,name=pet_id,json=petId,proto3" json:"pet_id,omitempty"`
	Quantity int32  `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	ShipDate string `protobuf:"bytes,4,opt,name=ship_date,json=shipDate,proto3" json:"ship_date,omitempty"`
	Status   string `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	Complete bool   `protobuf:"varint,6,opt,name=complete,proto3" json:"complete,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petstore_v3_store_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_petstore_v3_store_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_petstore_v3_store_proto_rawDescGZIP(), []int{0}
}

func (x *Order) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Order) GetPetId() int64 {
	if x != nil {
		return x.PetId
	}
	return 0
}

func (x *Order) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *Order) GetShipDate() string {
	if x != nil {
		return x.ShipDate
	}
	return ""
}

func (x *Order) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Order) GetComplete() bool {
	if x != nil {
		return x.Complete
	}
	return false
}

type AddStoreOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Order *Order `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *AddStoreOrderRequest) Reset() {
	*x = AddStoreOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petstore_v3_store_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddStoreOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddStoreOrderRequest) ProtoMessage() {}

func (x *AddStoreOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_petstore_v3_store_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddStoreOrderRequest.ProtoReflect.Descriptor instead.
func (*AddStoreOrderRequest) Descriptor() ([]byte, []int) {
	return file_petstore_v3_store_proto_rawDescGZIP(), []int{1}
}

func (x *AddStoreOrderRequest) GetOrder() *Order {
	if x != nil {
		return x.Order
	}
	return nil
}

type AddStoreOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Order *Order `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *AddStoreOrderResponse) Reset() {
	*x = AddStoreOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petstore_v3_store_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddStoreOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddStoreOrderResponse) ProtoMessage() {}

func (x *AddStoreOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_petstore_v3_store_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddStoreOrderResponse.ProtoReflect.Descriptor instead.
func (*AddStoreOrderResponse) Descriptor() ([]byte, []int) {
	return file_petstore_v3_store_proto_rawDescGZIP(), []int{2}
}

func (x *AddStoreOrderResponse) GetOrder() *Order {
	if x != nil {
		return x.Order
	}
	return nil
}

var File_petstore_v3_store_proto protoreflect.FileDescriptor

var file_petstore_v3_store_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x70, 0x65, 0x74, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9b, 0x01, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x15,
	0x0a, 0x06, 0x70, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x70, 0x65, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x69, 0x70, 0x44, 0x61, 0x74, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x65, 0x22, 0x40, 0x0a, 0x14, 0x41, 0x64, 0x64, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x05, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x65, 0x74, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x05, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x22, 0x41, 0x0a, 0x15, 0x41, 0x64, 0x64, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a,
	0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70,
	0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x32, 0x7f, 0x0a, 0x0c, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6f, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x21, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x41, 0x64, 0x64, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70, 0x65,
	0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x41, 0x64, 0x64, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a, 0x22, 0x0c, 0x2f, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x25, 0x5a, 0x23, 0x70, 0x65, 0x74, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2f, 0x76, 0x33, 0x3b, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x76, 0x33, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_petstore_v3_store_proto_rawDescOnce sync.Once
	file_petstore_v3_store_proto_rawDescData = file_petstore_v3_store_proto_rawDesc
)

func file_petstore_v3_store_proto_rawDescGZIP() []byte {
	file_petstore_v3_store_proto_rawDescOnce.Do(func() {
		file_petstore_v3_store_proto_rawDescData = protoimpl.X.CompressGZIP(file_petstore_v3_store_proto_rawDescData)
	})
	return file_petstore_v3_store_proto_rawDescData
}

var file_petstore_v3_store_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_petstore_v3_store_proto_goTypes = []interface{}{
	(*Order)(nil),                 // 0: petstore.v3.Order
	(*AddStoreOrderRequest)(nil),  // 1: petstore.v3.AddStoreOrderRequest
	(*AddStoreOrderResponse)(nil), // 2: petstore.v3.AddStoreOrderResponse
}
var file_petstore_v3_store_proto_depIdxs = []int32{
	0, // 0: petstore.v3.AddStoreOrderRequest.order:type_name -> petstore.v3.Order
	0, // 1: petstore.v3.AddStoreOrderResponse.order:type_name -> petstore.v3.Order
	1, // 2: petstore.v3.StoreService.AddStoreOrder:input_type -> petstore.v3.AddStoreOrderRequest
	2, // 3: petstore.v3.StoreService.AddStoreOrder:output_type -> petstore.v3.AddStoreOrderResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_petstore_v3_store_proto_init() }
func file_petstore_v3_store_proto_init() {
	if File_petstore_v3_store_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_petstore_v3_store_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
		file_petstore_v3_store_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddStoreOrderRequest); i {
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
		file_petstore_v3_store_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddStoreOrderResponse); i {
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
			RawDescriptor: file_petstore_v3_store_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_petstore_v3_store_proto_goTypes,
		DependencyIndexes: file_petstore_v3_store_proto_depIdxs,
		MessageInfos:      file_petstore_v3_store_proto_msgTypes,
	}.Build()
	File_petstore_v3_store_proto = out.File
	file_petstore_v3_store_proto_rawDesc = nil
	file_petstore_v3_store_proto_goTypes = nil
	file_petstore_v3_store_proto_depIdxs = nil
}
