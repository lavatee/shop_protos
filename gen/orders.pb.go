// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.1
// source: orders.proto

package proto

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

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	ProductId   int64  `protobuf:"varint,2,opt,name=ProductId,proto3" json:"ProductId,omitempty"`
	UserId      int64  `protobuf:"varint,3,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Price       int64  `protobuf:"varint,4,opt,name=Price,proto3" json:"Price,omitempty"`
	ProductName string `protobuf:"bytes,5,opt,name=ProductName,proto3" json:"ProductName,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orders_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_orders_proto_msgTypes[0]
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
	return file_orders_proto_rawDescGZIP(), []int{0}
}

func (x *Order) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Order) GetProductId() int64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *Order) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Order) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Order) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

type GetOrdersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

func (x *GetOrdersRequest) Reset() {
	*x = GetOrdersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orders_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrdersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrdersRequest) ProtoMessage() {}

func (x *GetOrdersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orders_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrdersRequest.ProtoReflect.Descriptor instead.
func (*GetOrdersRequest) Descriptor() ([]byte, []int) {
	return file_orders_proto_rawDescGZIP(), []int{1}
}

func (x *GetOrdersRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetOrdersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orders []*Order `protobuf:"bytes,1,rep,name=Orders,proto3" json:"Orders,omitempty"`
}

func (x *GetOrdersResponse) Reset() {
	*x = GetOrdersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orders_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrdersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrdersResponse) ProtoMessage() {}

func (x *GetOrdersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orders_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrdersResponse.ProtoReflect.Descriptor instead.
func (*GetOrdersResponse) Descriptor() ([]byte, []int) {
	return file_orders_proto_rawDescGZIP(), []int{2}
}

func (x *GetOrdersResponse) GetOrders() []*Order {
	if x != nil {
		return x.Orders
	}
	return nil
}

var File_orders_proto protoreflect.FileDescriptor

var file_orders_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x22, 0x85, 0x01, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x2a,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x3a, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x25, 0x0a, 0x06, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x06,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x32, 0x4a, 0x0a, 0x06, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73,
	0x12, 0x40, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x18, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_orders_proto_rawDescOnce sync.Once
	file_orders_proto_rawDescData = file_orders_proto_rawDesc
)

func file_orders_proto_rawDescGZIP() []byte {
	file_orders_proto_rawDescOnce.Do(func() {
		file_orders_proto_rawDescData = protoimpl.X.CompressGZIP(file_orders_proto_rawDescData)
	})
	return file_orders_proto_rawDescData
}

var file_orders_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_orders_proto_goTypes = []any{
	(*Order)(nil),             // 0: orders.Order
	(*GetOrdersRequest)(nil),  // 1: orders.GetOrdersRequest
	(*GetOrdersResponse)(nil), // 2: orders.GetOrdersResponse
}
var file_orders_proto_depIdxs = []int32{
	0, // 0: orders.GetOrdersResponse.Orders:type_name -> orders.Order
	1, // 1: orders.Orders.GetOrders:input_type -> orders.GetOrdersRequest
	2, // 2: orders.Orders.GetOrders:output_type -> orders.GetOrdersResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_orders_proto_init() }
func file_orders_proto_init() {
	if File_orders_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_orders_proto_msgTypes[0].Exporter = func(v any, i int) any {
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
		file_orders_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetOrdersRequest); i {
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
		file_orders_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetOrdersResponse); i {
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
			RawDescriptor: file_orders_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_orders_proto_goTypes,
		DependencyIndexes: file_orders_proto_depIdxs,
		MessageInfos:      file_orders_proto_msgTypes,
	}.Build()
	File_orders_proto = out.File
	file_orders_proto_rawDesc = nil
	file_orders_proto_goTypes = nil
	file_orders_proto_depIdxs = nil
}
