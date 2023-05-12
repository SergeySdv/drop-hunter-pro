// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: v1/settings.proto

package v1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ResetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ResetRequest) Reset() {
	*x = ResetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_settings_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetRequest) ProtoMessage() {}

func (x *ResetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_settings_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetRequest.ProtoReflect.Descriptor instead.
func (*ResetRequest) Descriptor() ([]byte, []int) {
	return file_v1_settings_proto_rawDescGZIP(), []int{0}
}

type ResetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Settings *Settings `protobuf:"bytes,1,opt,name=settings,proto3" json:"settings,omitempty"`
}

func (x *ResetResponse) Reset() {
	*x = ResetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_settings_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetResponse) ProtoMessage() {}

func (x *ResetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_settings_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetResponse.ProtoReflect.Descriptor instead.
func (*ResetResponse) Descriptor() ([]byte, []int) {
	return file_v1_settings_proto_rawDescGZIP(), []int{1}
}

func (x *ResetResponse) GetSettings() *Settings {
	if x != nil {
		return x.Settings
	}
	return nil
}

type GetNetworkByRPCResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Valid             bool    `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
	Id                int64   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	SuggestedGasPrice int64   `protobuf:"varint,3,opt,name=suggested_gas_price,json=suggestedGasPrice,proto3" json:"suggested_gas_price,omitempty"`
	Error             *string `protobuf:"bytes,4,opt,name=error,proto3,oneof" json:"error,omitempty"`
}

func (x *GetNetworkByRPCResponse) Reset() {
	*x = GetNetworkByRPCResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_settings_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNetworkByRPCResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNetworkByRPCResponse) ProtoMessage() {}

func (x *GetNetworkByRPCResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_settings_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNetworkByRPCResponse.ProtoReflect.Descriptor instead.
func (*GetNetworkByRPCResponse) Descriptor() ([]byte, []int) {
	return file_v1_settings_proto_rawDescGZIP(), []int{2}
}

func (x *GetNetworkByRPCResponse) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

func (x *GetNetworkByRPCResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetNetworkByRPCResponse) GetSuggestedGasPrice() int64 {
	if x != nil {
		return x.SuggestedGasPrice
	}
	return 0
}

func (x *GetNetworkByRPCResponse) GetError() string {
	if x != nil && x.Error != nil {
		return *x.Error
	}
	return ""
}

type GetNetworkByRPCRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoint string  `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	Network  Network `protobuf:"varint,2,opt,name=network,proto3,enum=shared.Network" json:"network,omitempty"`
}

func (x *GetNetworkByRPCRequest) Reset() {
	*x = GetNetworkByRPCRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_settings_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNetworkByRPCRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNetworkByRPCRequest) ProtoMessage() {}

func (x *GetNetworkByRPCRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_settings_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNetworkByRPCRequest.ProtoReflect.Descriptor instead.
func (*GetNetworkByRPCRequest) Descriptor() ([]byte, []int) {
	return file_v1_settings_proto_rawDescGZIP(), []int{3}
}

func (x *GetNetworkByRPCRequest) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *GetNetworkByRPCRequest) GetNetwork() Network {
	if x != nil {
		return x.Network
	}
	return Network_ARBITRUM
}

type UpdateSettingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Settings *Settings `protobuf:"bytes,1,opt,name=settings,proto3" json:"settings,omitempty"`
}

func (x *UpdateSettingsRequest) Reset() {
	*x = UpdateSettingsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_settings_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSettingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSettingsRequest) ProtoMessage() {}

func (x *UpdateSettingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_settings_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSettingsRequest.ProtoReflect.Descriptor instead.
func (*UpdateSettingsRequest) Descriptor() ([]byte, []int) {
	return file_v1_settings_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateSettingsRequest) GetSettings() *Settings {
	if x != nil {
		return x.Settings
	}
	return nil
}

type UpdateSettingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Settings *Settings `protobuf:"bytes,1,opt,name=settings,proto3" json:"settings,omitempty"`
}

func (x *UpdateSettingsResponse) Reset() {
	*x = UpdateSettingsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_settings_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSettingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSettingsResponse) ProtoMessage() {}

func (x *UpdateSettingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_settings_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSettingsResponse.ProtoReflect.Descriptor instead.
func (*UpdateSettingsResponse) Descriptor() ([]byte, []int) {
	return file_v1_settings_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateSettingsResponse) GetSettings() *Settings {
	if x != nil {
		return x.Settings
	}
	return nil
}

type GetSettingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetSettingsRequest) Reset() {
	*x = GetSettingsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_settings_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSettingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSettingsRequest) ProtoMessage() {}

func (x *GetSettingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_settings_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSettingsRequest.ProtoReflect.Descriptor instead.
func (*GetSettingsRequest) Descriptor() ([]byte, []int) {
	return file_v1_settings_proto_rawDescGZIP(), []int{6}
}

type GetSettingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Settings *Settings `protobuf:"bytes,1,opt,name=settings,proto3" json:"settings,omitempty"`
}

func (x *GetSettingsResponse) Reset() {
	*x = GetSettingsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_settings_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSettingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSettingsResponse) ProtoMessage() {}

func (x *GetSettingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_settings_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSettingsResponse.ProtoReflect.Descriptor instead.
func (*GetSettingsResponse) Descriptor() ([]byte, []int) {
	return file_v1_settings_proto_rawDescGZIP(), []int{7}
}

func (x *GetSettingsResponse) GetSettings() *Settings {
	if x != nil {
		return x.Settings
	}
	return nil
}

type Settings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string           `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Bnb       *SettingsNetwork `protobuf:"bytes,2,opt,name=bnb,proto3" json:"bnb,omitempty"`
	Optimism  *SettingsNetwork `protobuf:"bytes,3,opt,name=optimism,proto3" json:"optimism,omitempty"`
	Arbitrum  *SettingsNetwork `protobuf:"bytes,4,opt,name=arbitrum,proto3" json:"arbitrum,omitempty"`
	Etherium  *SettingsNetwork `protobuf:"bytes,5,opt,name=etherium,proto3" json:"etherium,omitempty"`
	Polygon   *SettingsNetwork `protobuf:"bytes,6,opt,name=polygon,proto3" json:"polygon,omitempty"`
	Avalanche *SettingsNetwork `protobuf:"bytes,7,opt,name=avalanche,proto3" json:"avalanche,omitempty"`
}

func (x *Settings) Reset() {
	*x = Settings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_settings_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Settings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Settings) ProtoMessage() {}

func (x *Settings) ProtoReflect() protoreflect.Message {
	mi := &file_v1_settings_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Settings.ProtoReflect.Descriptor instead.
func (*Settings) Descriptor() ([]byte, []int) {
	return file_v1_settings_proto_rawDescGZIP(), []int{8}
}

func (x *Settings) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Settings) GetBnb() *SettingsNetwork {
	if x != nil {
		return x.Bnb
	}
	return nil
}

func (x *Settings) GetOptimism() *SettingsNetwork {
	if x != nil {
		return x.Optimism
	}
	return nil
}

func (x *Settings) GetArbitrum() *SettingsNetwork {
	if x != nil {
		return x.Arbitrum
	}
	return nil
}

func (x *Settings) GetEtherium() *SettingsNetwork {
	if x != nil {
		return x.Etherium
	}
	return nil
}

func (x *Settings) GetPolygon() *SettingsNetwork {
	if x != nil {
		return x.Polygon
	}
	return nil
}

func (x *Settings) GetAvalanche() *SettingsNetwork {
	if x != nil {
		return x.Avalanche
	}
	return nil
}

type SettingsNetwork struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UseLimitGas bool   `protobuf:"varint,2,opt,name=use_limit_gas,json=useLimitGas,proto3" json:"use_limit_gas,omitempty"`
	RpcEndpoint string `protobuf:"bytes,3,opt,name=rpc_endpoint,json=rpcEndpoint,proto3" json:"rpc_endpoint,omitempty"`
	GasTotal    *int64 `protobuf:"varint,4,opt,name=gas_total,json=gasTotal,proto3,oneof" json:"gas_total,omitempty"`
}

func (x *SettingsNetwork) Reset() {
	*x = SettingsNetwork{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_settings_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SettingsNetwork) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SettingsNetwork) ProtoMessage() {}

func (x *SettingsNetwork) ProtoReflect() protoreflect.Message {
	mi := &file_v1_settings_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SettingsNetwork.ProtoReflect.Descriptor instead.
func (*SettingsNetwork) Descriptor() ([]byte, []int) {
	return file_v1_settings_proto_rawDescGZIP(), []int{9}
}

func (x *SettingsNetwork) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SettingsNetwork) GetUseLimitGas() bool {
	if x != nil {
		return x.UseLimitGas
	}
	return false
}

func (x *SettingsNetwork) GetRpcEndpoint() string {
	if x != nil {
		return x.RpcEndpoint
	}
	return ""
}

func (x *SettingsNetwork) GetGasTotal() int64 {
	if x != nil && x.GasTotal != nil {
		return *x.GasTotal
	}
	return 0
}

var File_v1_settings_proto protoreflect.FileDescriptor

var file_v1_settings_proto_rawDesc = []byte{
	0x0a, 0x11, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x76, 0x31,
	0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0e, 0x0a,
	0x0c, 0x52, 0x65, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x51, 0x0a,
	0x0d, 0x52, 0x65, 0x73, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e,
	0x0a, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x52, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x3a, 0x10,
	0x92, 0x41, 0x0d, 0x0a, 0x0b, 0xd2, 0x01, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x22, 0xbe, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x42,
	0x79, 0x52, 0x50, 0x43, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x2e, 0x0a, 0x13, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f,
	0x67, 0x61, 0x73, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x11, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x65, 0x64, 0x47, 0x61, 0x73, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x19, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x3a, 0x28, 0x92,
	0x41, 0x25, 0x0a, 0x23, 0xd2, 0x01, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0xd2, 0x01, 0x02, 0x69,
	0x64, 0xd2, 0x01, 0x13, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x67, 0x61,
	0x73, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x22, 0x7b, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x42,
	0x79, 0x52, 0x50, 0x43, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x3a, 0x1a, 0x92, 0x41, 0x17, 0x0a, 0x15, 0xd2, 0x01, 0x08, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0xd2, 0x01, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x22, 0x59,
	0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x08, 0x73,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x3a, 0x10, 0x92, 0x41, 0x0d, 0x0a, 0x0b, 0xd2, 0x01,
	0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x5a, 0x0a, 0x16, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x2e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x3a, 0x10, 0x92, 0x41, 0x0d, 0x0a, 0x0b, 0xd2, 0x01, 0x08, 0x73, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x14, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x57, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2e, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x3a, 0x10, 0x92, 0x41, 0x0d, 0x0a, 0x0b, 0xd2, 0x01, 0x08, 0x73, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x22, 0xa7, 0x03, 0x0a, 0x08, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x03, 0x62, 0x6e,
	0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x52, 0x03, 0x62, 0x6e, 0x62, 0x12, 0x35, 0x0a, 0x08, 0x6f, 0x70, 0x74, 0x69, 0x6d,
	0x69, 0x73, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x52, 0x08, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x73, 0x6d, 0x12, 0x35,
	0x0a, 0x08, 0x61, 0x72, 0x62, 0x69, 0x74, 0x72, 0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x08, 0x61, 0x72, 0x62,
	0x69, 0x74, 0x72, 0x75, 0x6d, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x74, 0x68, 0x65, 0x72, 0x69, 0x75,
	0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x52, 0x08, 0x65, 0x74, 0x68, 0x65, 0x72, 0x69, 0x75, 0x6d, 0x12, 0x33, 0x0a, 0x07,
	0x70, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x07, 0x70, 0x6f, 0x6c, 0x79, 0x67, 0x6f,
	0x6e, 0x12, 0x37, 0x0a, 0x09, 0x61, 0x76, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52,
	0x09, 0x61, 0x76, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x3a, 0x42, 0x92, 0x41, 0x3f, 0x0a,
	0x3d, 0xd2, 0x01, 0x03, 0x62, 0x6e, 0x62, 0xd2, 0x01, 0x08, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x69,
	0x73, 0x6d, 0xd2, 0x01, 0x08, 0x61, 0x72, 0x62, 0x69, 0x74, 0x72, 0x75, 0x6d, 0xd2, 0x01, 0x08,
	0x65, 0x74, 0x68, 0x65, 0x72, 0x69, 0x75, 0x6d, 0xd2, 0x01, 0x07, 0x70, 0x6f, 0x6c, 0x79, 0x67,
	0x6f, 0x6e, 0xd2, 0x01, 0x09, 0x61, 0x76, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x22, 0xc3,
	0x01, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f,
	0x67, 0x61, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x47, 0x61, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x70, 0x63, 0x5f, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x70,
	0x63, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x09, 0x67, 0x61, 0x73,
	0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x08,
	0x67, 0x61, 0x73, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x88, 0x01, 0x01, 0x3a, 0x29, 0x92, 0x41, 0x26,
	0x0a, 0x24, 0xd2, 0x01, 0x02, 0x69, 0x64, 0xd2, 0x01, 0x0d, 0x75, 0x73, 0x65, 0x5f, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x5f, 0x67, 0x61, 0x73, 0xd2, 0x01, 0x0c, 0x72, 0x70, 0x63, 0x5f, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x67, 0x61, 0x73, 0x5f, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x32, 0xfb, 0x03, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x62, 0x0a, 0x05, 0x52, 0x65, 0x73, 0x65,
	0x74, 0x12, 0x16, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x52, 0x65, 0x73,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x3a, 0x01, 0x2a, 0x22, 0x1d, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x67, 0x77, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x65, 0x74, 0x12, 0x8f, 0x01, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x42, 0x79, 0x52, 0x50, 0x43,
	0x12, 0x20, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x42, 0x79, 0x52, 0x50, 0x43, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x47, 0x65,
	0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x42, 0x79, 0x52, 0x50, 0x43, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x37, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x31, 0x3a, 0x01, 0x2a,
	0x22, 0x2c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x77, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12, 0x72,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x1c, 0x2e,
	0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x20, 0x3a, 0x01, 0x2a, 0x22, 0x1b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x77, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x67,
	0x65, 0x74, 0x12, 0x7e, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x12, 0x1f, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x3a,
	0x01, 0x2a, 0x22, 0x1e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x77, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_settings_proto_rawDescOnce sync.Once
	file_v1_settings_proto_rawDescData = file_v1_settings_proto_rawDesc
)

func file_v1_settings_proto_rawDescGZIP() []byte {
	file_v1_settings_proto_rawDescOnce.Do(func() {
		file_v1_settings_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_settings_proto_rawDescData)
	})
	return file_v1_settings_proto_rawDescData
}

var file_v1_settings_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_v1_settings_proto_goTypes = []interface{}{
	(*ResetRequest)(nil),            // 0: settings.ResetRequest
	(*ResetResponse)(nil),           // 1: settings.ResetResponse
	(*GetNetworkByRPCResponse)(nil), // 2: settings.GetNetworkByRPCResponse
	(*GetNetworkByRPCRequest)(nil),  // 3: settings.GetNetworkByRPCRequest
	(*UpdateSettingsRequest)(nil),   // 4: settings.UpdateSettingsRequest
	(*UpdateSettingsResponse)(nil),  // 5: settings.UpdateSettingsResponse
	(*GetSettingsRequest)(nil),      // 6: settings.GetSettingsRequest
	(*GetSettingsResponse)(nil),     // 7: settings.GetSettingsResponse
	(*Settings)(nil),                // 8: settings.Settings
	(*SettingsNetwork)(nil),         // 9: settings.SettingsNetwork
	(Network)(0),                    // 10: shared.Network
}
var file_v1_settings_proto_depIdxs = []int32{
	8,  // 0: settings.ResetResponse.settings:type_name -> settings.Settings
	10, // 1: settings.GetNetworkByRPCRequest.network:type_name -> shared.Network
	8,  // 2: settings.UpdateSettingsRequest.settings:type_name -> settings.Settings
	8,  // 3: settings.UpdateSettingsResponse.settings:type_name -> settings.Settings
	8,  // 4: settings.GetSettingsResponse.settings:type_name -> settings.Settings
	9,  // 5: settings.Settings.bnb:type_name -> settings.SettingsNetwork
	9,  // 6: settings.Settings.optimism:type_name -> settings.SettingsNetwork
	9,  // 7: settings.Settings.arbitrum:type_name -> settings.SettingsNetwork
	9,  // 8: settings.Settings.etherium:type_name -> settings.SettingsNetwork
	9,  // 9: settings.Settings.polygon:type_name -> settings.SettingsNetwork
	9,  // 10: settings.Settings.avalanche:type_name -> settings.SettingsNetwork
	0,  // 11: settings.SettingsService.Reset:input_type -> settings.ResetRequest
	3,  // 12: settings.SettingsService.GetNetworkByRPC:input_type -> settings.GetNetworkByRPCRequest
	6,  // 13: settings.SettingsService.GetSettings:input_type -> settings.GetSettingsRequest
	4,  // 14: settings.SettingsService.UpdateSettings:input_type -> settings.UpdateSettingsRequest
	1,  // 15: settings.SettingsService.Reset:output_type -> settings.ResetResponse
	2,  // 16: settings.SettingsService.GetNetworkByRPC:output_type -> settings.GetNetworkByRPCResponse
	7,  // 17: settings.SettingsService.GetSettings:output_type -> settings.GetSettingsResponse
	5,  // 18: settings.SettingsService.UpdateSettings:output_type -> settings.UpdateSettingsResponse
	15, // [15:19] is the sub-list for method output_type
	11, // [11:15] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_v1_settings_proto_init() }
func file_v1_settings_proto_init() {
	if File_v1_settings_proto != nil {
		return
	}
	file_v1_shared_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v1_settings_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetRequest); i {
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
		file_v1_settings_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetResponse); i {
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
		file_v1_settings_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNetworkByRPCResponse); i {
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
		file_v1_settings_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNetworkByRPCRequest); i {
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
		file_v1_settings_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSettingsRequest); i {
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
		file_v1_settings_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSettingsResponse); i {
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
		file_v1_settings_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSettingsRequest); i {
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
		file_v1_settings_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSettingsResponse); i {
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
		file_v1_settings_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Settings); i {
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
		file_v1_settings_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SettingsNetwork); i {
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
	file_v1_settings_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_v1_settings_proto_msgTypes[9].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_settings_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_settings_proto_goTypes,
		DependencyIndexes: file_v1_settings_proto_depIdxs,
		MessageInfos:      file_v1_settings_proto_msgTypes,
	}.Build()
	File_v1_settings_proto = out.File
	file_v1_settings_proto_rawDesc = nil
	file_v1_settings_proto_goTypes = nil
	file_v1_settings_proto_depIdxs = nil
}