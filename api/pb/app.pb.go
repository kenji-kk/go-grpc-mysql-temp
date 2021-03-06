// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: api/pb/app.proto

package pb

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

type Country struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Country) Reset() {
	*x = Country{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_pb_app_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Country) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Country) ProtoMessage() {}

func (x *Country) ProtoReflect() protoreflect.Message {
	mi := &file_api_pb_app_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Country.ProtoReflect.Descriptor instead.
func (*Country) Descriptor() ([]byte, []int) {
	return file_api_pb_app_proto_rawDescGZIP(), []int{0}
}

func (x *Country) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Country) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type AllCountryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AllCountryRequest) Reset() {
	*x = AllCountryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_pb_app_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllCountryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllCountryRequest) ProtoMessage() {}

func (x *AllCountryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_pb_app_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllCountryRequest.ProtoReflect.Descriptor instead.
func (*AllCountryRequest) Descriptor() ([]byte, []int) {
	return file_api_pb_app_proto_rawDescGZIP(), []int{1}
}

type AllCountryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AllCountry []*Country `protobuf:"bytes,1,rep,name=AllCountry,proto3" json:"AllCountry,omitempty"`
}

func (x *AllCountryResponse) Reset() {
	*x = AllCountryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_pb_app_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllCountryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllCountryResponse) ProtoMessage() {}

func (x *AllCountryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_pb_app_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllCountryResponse.ProtoReflect.Descriptor instead.
func (*AllCountryResponse) Descriptor() ([]byte, []int) {
	return file_api_pb_app_proto_rawDescGZIP(), []int{2}
}

func (x *AllCountryResponse) GetAllCountry() []*Country {
	if x != nil {
		return x.AllCountry
	}
	return nil
}

var File_api_pb_app_proto protoreflect.FileDescriptor

var file_api_pb_app_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62, 0x2f, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x2d, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x13, 0x0a, 0x11, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x42, 0x0a, 0x12, 0x41,
	0x6c, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2c, 0x0a, 0x0a, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x0a, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x32,
	0x50, 0x0a, 0x0a, 0x41, 0x70, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x16,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x6c, 0x6c,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_pb_app_proto_rawDescOnce sync.Once
	file_api_pb_app_proto_rawDescData = file_api_pb_app_proto_rawDesc
)

func file_api_pb_app_proto_rawDescGZIP() []byte {
	file_api_pb_app_proto_rawDescOnce.Do(func() {
		file_api_pb_app_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_pb_app_proto_rawDescData)
	})
	return file_api_pb_app_proto_rawDescData
}

var file_api_pb_app_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_pb_app_proto_goTypes = []interface{}{
	(*Country)(nil),            // 0: api.Country
	(*AllCountryRequest)(nil),  // 1: api.AllCountryRequest
	(*AllCountryResponse)(nil), // 2: api.AllCountryResponse
}
var file_api_pb_app_proto_depIdxs = []int32{
	0, // 0: api.AllCountryResponse.AllCountry:type_name -> api.Country
	1, // 1: api.AppService.GetAllCountry:input_type -> api.AllCountryRequest
	2, // 2: api.AppService.GetAllCountry:output_type -> api.AllCountryResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_pb_app_proto_init() }
func file_api_pb_app_proto_init() {
	if File_api_pb_app_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_pb_app_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Country); i {
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
		file_api_pb_app_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllCountryRequest); i {
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
		file_api_pb_app_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllCountryResponse); i {
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
			RawDescriptor: file_api_pb_app_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_pb_app_proto_goTypes,
		DependencyIndexes: file_api_pb_app_proto_depIdxs,
		MessageInfos:      file_api_pb_app_proto_msgTypes,
	}.Build()
	File_api_pb_app_proto = out.File
	file_api_pb_app_proto_rawDesc = nil
	file_api_pb_app_proto_goTypes = nil
	file_api_pb_app_proto_depIdxs = nil
}
