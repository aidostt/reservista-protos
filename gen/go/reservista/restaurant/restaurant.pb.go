// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: reservista/restaurant/restaurant.proto

package proto_restaurant

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservista_restaurant_restaurant_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_reservista_restaurant_restaurant_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_reservista_restaurant_restaurant_proto_rawDescGZIP(), []int{0}
}

type RestaurantObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Address string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Contact string `protobuf:"bytes,4,opt,name=contact,proto3" json:"contact,omitempty"`
}

func (x *RestaurantObject) Reset() {
	*x = RestaurantObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservista_restaurant_restaurant_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RestaurantObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestaurantObject) ProtoMessage() {}

func (x *RestaurantObject) ProtoReflect() protoreflect.Message {
	mi := &file_reservista_restaurant_restaurant_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestaurantObject.ProtoReflect.Descriptor instead.
func (*RestaurantObject) Descriptor() ([]byte, []int) {
	return file_reservista_restaurant_restaurant_proto_rawDescGZIP(), []int{1}
}

func (x *RestaurantObject) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RestaurantObject) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RestaurantObject) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *RestaurantObject) GetContact() string {
	if x != nil {
		return x.Contact
	}
	return ""
}

type IDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IDRequest) Reset() {
	*x = IDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservista_restaurant_restaurant_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDRequest) ProtoMessage() {}

func (x *IDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reservista_restaurant_restaurant_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDRequest.ProtoReflect.Descriptor instead.
func (*IDRequest) Descriptor() ([]byte, []int) {
	return file_reservista_restaurant_restaurant_proto_rawDescGZIP(), []int{2}
}

func (x *IDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type StatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *StatusResponse) Reset() {
	*x = StatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservista_restaurant_restaurant_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusResponse) ProtoMessage() {}

func (x *StatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_reservista_restaurant_restaurant_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusResponse.ProtoReflect.Descriptor instead.
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return file_reservista_restaurant_restaurant_proto_rawDescGZIP(), []int{3}
}

func (x *StatusResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type RestaurantListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Restaurants []*RestaurantObject `protobuf:"bytes,1,rep,name=restaurants,proto3" json:"restaurants,omitempty"`
}

func (x *RestaurantListResponse) Reset() {
	*x = RestaurantListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservista_restaurant_restaurant_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RestaurantListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestaurantListResponse) ProtoMessage() {}

func (x *RestaurantListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_reservista_restaurant_restaurant_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestaurantListResponse.ProtoReflect.Descriptor instead.
func (*RestaurantListResponse) Descriptor() ([]byte, []int) {
	return file_reservista_restaurant_restaurant_proto_rawDescGZIP(), []int{4}
}

func (x *RestaurantListResponse) GetRestaurants() []*RestaurantObject {
	if x != nil {
		return x.Restaurants
	}
	return nil
}

var File_reservista_restaurant_restaurant_proto protoreflect.FileDescriptor

var file_reservista_restaurant_restaurant_proto_rawDesc = []byte{
	0x0a, 0x26, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x73, 0x74, 0x61, 0x2f, 0x72, 0x65, 0x73,
	0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75,
	0x72, 0x61, 0x6e, 0x74, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x6a, 0x0a,
	0x10, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x22, 0x1b, 0x0a, 0x09, 0x49, 0x44, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x28, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x58, 0x0a, 0x16, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0b, 0x72, 0x65,
	0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x73,
	0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x0b, 0x72,
	0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x73, 0x32, 0x80, 0x03, 0x0a, 0x0a, 0x52,
	0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x12, 0x49, 0x0a, 0x0d, 0x41, 0x64, 0x64,
	0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x12, 0x1c, 0x2e, 0x72, 0x65, 0x73,
	0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61,
	0x6e, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x1a, 0x1a, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61,
	0x75, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65,
	0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x11, 0x2e, 0x72, 0x65, 0x73, 0x74,
	0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x22, 0x2e, 0x72,
	0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75,
	0x72, 0x61, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x44, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e,
	0x74, 0x12, 0x15, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x49,
	0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61,
	0x75, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x4a, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1c, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61,
	0x75, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x1a, 0x1a, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72,
	0x61, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x49, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x74,
	0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x15, 0x2e, 0x72, 0x65, 0x73,
	0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x39, 0x5a,
	0x37, 0x61, 0x69, 0x64, 0x6f, 0x73, 0x74, 0x74, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x73, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75,
	0x72, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x72, 0x65,
	0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_reservista_restaurant_restaurant_proto_rawDescOnce sync.Once
	file_reservista_restaurant_restaurant_proto_rawDescData = file_reservista_restaurant_restaurant_proto_rawDesc
)

func file_reservista_restaurant_restaurant_proto_rawDescGZIP() []byte {
	file_reservista_restaurant_restaurant_proto_rawDescOnce.Do(func() {
		file_reservista_restaurant_restaurant_proto_rawDescData = protoimpl.X.CompressGZIP(file_reservista_restaurant_restaurant_proto_rawDescData)
	})
	return file_reservista_restaurant_restaurant_proto_rawDescData
}

var file_reservista_restaurant_restaurant_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_reservista_restaurant_restaurant_proto_goTypes = []interface{}{
	(*Empty)(nil),                  // 0: restaurant.Empty
	(*RestaurantObject)(nil),       // 1: restaurant.RestaurantObject
	(*IDRequest)(nil),              // 2: restaurant.IDRequest
	(*StatusResponse)(nil),         // 3: restaurant.StatusResponse
	(*RestaurantListResponse)(nil), // 4: restaurant.RestaurantListResponse
}
var file_reservista_restaurant_restaurant_proto_depIdxs = []int32{
	1, // 0: restaurant.RestaurantListResponse.restaurants:type_name -> restaurant.RestaurantObject
	1, // 1: restaurant.Restaurant.AddRestaurant:input_type -> restaurant.RestaurantObject
	0, // 2: restaurant.Restaurant.GetAllRestaurants:input_type -> restaurant.Empty
	2, // 3: restaurant.Restaurant.GetRestaurant:input_type -> restaurant.IDRequest
	1, // 4: restaurant.Restaurant.UpdateRestById:input_type -> restaurant.RestaurantObject
	2, // 5: restaurant.Restaurant.DeleteRestaurantById:input_type -> restaurant.IDRequest
	3, // 6: restaurant.Restaurant.AddRestaurant:output_type -> restaurant.StatusResponse
	4, // 7: restaurant.Restaurant.GetAllRestaurants:output_type -> restaurant.RestaurantListResponse
	1, // 8: restaurant.Restaurant.GetRestaurant:output_type -> restaurant.RestaurantObject
	3, // 9: restaurant.Restaurant.UpdateRestById:output_type -> restaurant.StatusResponse
	3, // 10: restaurant.Restaurant.DeleteRestaurantById:output_type -> restaurant.StatusResponse
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_reservista_restaurant_restaurant_proto_init() }
func file_reservista_restaurant_restaurant_proto_init() {
	if File_reservista_restaurant_restaurant_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_reservista_restaurant_restaurant_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_reservista_restaurant_restaurant_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RestaurantObject); i {
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
		file_reservista_restaurant_restaurant_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDRequest); i {
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
		file_reservista_restaurant_restaurant_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusResponse); i {
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
		file_reservista_restaurant_restaurant_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RestaurantListResponse); i {
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
			RawDescriptor: file_reservista_restaurant_restaurant_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_reservista_restaurant_restaurant_proto_goTypes,
		DependencyIndexes: file_reservista_restaurant_restaurant_proto_depIdxs,
		MessageInfos:      file_reservista_restaurant_restaurant_proto_msgTypes,
	}.Build()
	File_reservista_restaurant_restaurant_proto = out.File
	file_reservista_restaurant_restaurant_proto_rawDesc = nil
	file_reservista_restaurant_restaurant_proto_goTypes = nil
	file_reservista_restaurant_restaurant_proto_depIdxs = nil
}