// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        v3.15.8
// source: market/market.proto

package market

import (
	common "gateway-service/client/common"
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

type MarketLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Province string `protobuf:"bytes,1,opt,name=province,proto3" json:"province"`
	City     string `protobuf:"bytes,2,opt,name=city,proto3" json:"city"`
	Ward     string `protobuf:"bytes,3,opt,name=ward,proto3" json:"ward"`
	Address  string `protobuf:"bytes,4,opt,name=address,proto3" json:"address"`
}

func (x *MarketLocation) Reset() {
	*x = MarketLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_market_market_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarketLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarketLocation) ProtoMessage() {}

func (x *MarketLocation) ProtoReflect() protoreflect.Message {
	mi := &file_market_market_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarketLocation.ProtoReflect.Descriptor instead.
func (*MarketLocation) Descriptor() ([]byte, []int) {
	return file_market_market_proto_rawDescGZIP(), []int{0}
}

func (x *MarketLocation) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *MarketLocation) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *MarketLocation) GetWard() string {
	if x != nil {
		return x.Ward
	}
	return ""
}

func (x *MarketLocation) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type Market struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MarketId      string       `protobuf:"bytes,1,opt,name=market_id,json=marketId,proto3" json:"market_id"`
	Name          string       `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	Location      *Location    `protobuf:"bytes,3,opt,name=location,proto3" json:"location"`
	Type          MarketType   `protobuf:"varint,4,opt,name=type,proto3,enum=market.MarketType" json:"type"`
	Status        MarketStatus `protobuf:"varint,5,opt,name=status,proto3,enum=market.MarketStatus" json:"status"`
	State         MarketState  `protobuf:"varint,6,opt,name=state,proto3,enum=market.MarketState" json:"state"`
	HasPublished  bool         `protobuf:"varint,7,opt,name=has_published,json=hasPublished,proto3" json:"has_published"`
	HasDraft      bool         `protobuf:"varint,8,opt,name=has_draft,json=hasDraft,proto3" json:"has_draft"`
	FullAddress   string       `protobuf:"bytes,13,opt,name=full_address,json=fullAddress,proto3" json:"full_address"`
	Supervisor    *Supervisor  `protobuf:"bytes,14,opt,name=supervisor,proto3" json:"supervisor"`
	AvailableDate int64        `protobuf:"varint,15,opt,name=available_date,json=availableDate,proto3" json:"available_date"`
	GoogleMap     string       `protobuf:"bytes,16,opt,name=google_map,json=googleMap,proto3" json:"google_map"`
	Code          string       `protobuf:"bytes,17,opt,name=code,proto3" json:"code"`
	Clazz         MarketClass  `protobuf:"varint,18,opt,name=clazz,proto3,enum=market.MarketClass" json:"clazz"`
	HasDeleted    bool         `protobuf:"varint,20,opt,name=has_deleted,json=hasDeleted,proto3" json:"has_deleted"`
}

func (x *Market) Reset() {
	*x = Market{}
	if protoimpl.UnsafeEnabled {
		mi := &file_market_market_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Market) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Market) ProtoMessage() {}

func (x *Market) ProtoReflect() protoreflect.Message {
	mi := &file_market_market_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Market.ProtoReflect.Descriptor instead.
func (*Market) Descriptor() ([]byte, []int) {
	return file_market_market_proto_rawDescGZIP(), []int{1}
}

func (x *Market) GetMarketId() string {
	if x != nil {
		return x.MarketId
	}
	return ""
}

func (x *Market) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Market) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *Market) GetType() MarketType {
	if x != nil {
		return x.Type
	}
	return MarketType_MARKET_TYPE_UNSPECIFIED
}

func (x *Market) GetStatus() MarketStatus {
	if x != nil {
		return x.Status
	}
	return MarketStatus_MARKET_STATUS_UNSPECIFIED
}

func (x *Market) GetState() MarketState {
	if x != nil {
		return x.State
	}
	return MarketState_MARKET_STATE_UNSPECIFIED
}

func (x *Market) GetHasPublished() bool {
	if x != nil {
		return x.HasPublished
	}
	return false
}

func (x *Market) GetHasDraft() bool {
	if x != nil {
		return x.HasDraft
	}
	return false
}

func (x *Market) GetFullAddress() string {
	if x != nil {
		return x.FullAddress
	}
	return ""
}

func (x *Market) GetSupervisor() *Supervisor {
	if x != nil {
		return x.Supervisor
	}
	return nil
}

func (x *Market) GetAvailableDate() int64 {
	if x != nil {
		return x.AvailableDate
	}
	return 0
}

func (x *Market) GetGoogleMap() string {
	if x != nil {
		return x.GoogleMap
	}
	return ""
}

func (x *Market) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Market) GetClazz() MarketClass {
	if x != nil {
		return x.Clazz
	}
	return MarketClass_MARKET_CLASS_UNSPECIFIED
}

func (x *Market) GetHasDeleted() bool {
	if x != nil {
		return x.HasDeleted
	}
	return false
}

///////////////////////////////////////////////////////////////////
// Get Market
///////////////////////////////////////////////////////////////////
type GetMarketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MarketId string `protobuf:"bytes,1,opt,name=market_id,json=marketId,proto3" json:"market_id"`
	Draft    *bool  `protobuf:"varint,2,opt,name=draft,proto3,oneof" json:"draft"`
	Code     string `protobuf:"bytes,3,opt,name=code,proto3" json:"code"`
}

func (x *GetMarketRequest) Reset() {
	*x = GetMarketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_market_market_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMarketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMarketRequest) ProtoMessage() {}

func (x *GetMarketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_market_market_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMarketRequest.ProtoReflect.Descriptor instead.
func (*GetMarketRequest) Descriptor() ([]byte, []int) {
	return file_market_market_proto_rawDescGZIP(), []int{2}
}

func (x *GetMarketRequest) GetMarketId() string {
	if x != nil {
		return x.MarketId
	}
	return ""
}

func (x *GetMarketRequest) GetDraft() bool {
	if x != nil && x.Draft != nil {
		return *x.Draft
	}
	return false
}

func (x *GetMarketRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type GetMarketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success"`
	// Types that are assignable to Response:
	//	*GetMarketResponse_Data_
	//	*GetMarketResponse_Error
	Response isGetMarketResponse_Response `protobuf_oneof:"response"`
}

func (x *GetMarketResponse) Reset() {
	*x = GetMarketResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_market_market_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMarketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMarketResponse) ProtoMessage() {}

func (x *GetMarketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_market_market_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMarketResponse.ProtoReflect.Descriptor instead.
func (*GetMarketResponse) Descriptor() ([]byte, []int) {
	return file_market_market_proto_rawDescGZIP(), []int{3}
}

func (x *GetMarketResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (m *GetMarketResponse) GetResponse() isGetMarketResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (x *GetMarketResponse) GetData() *GetMarketResponse_Data {
	if x, ok := x.GetResponse().(*GetMarketResponse_Data_); ok {
		return x.Data
	}
	return nil
}

func (x *GetMarketResponse) GetError() *common.Error {
	if x, ok := x.GetResponse().(*GetMarketResponse_Error); ok {
		return x.Error
	}
	return nil
}

type isGetMarketResponse_Response interface {
	isGetMarketResponse_Response()
}

type GetMarketResponse_Data_ struct {
	Data *GetMarketResponse_Data `protobuf:"bytes,2,opt,name=data,proto3,oneof"`
}

type GetMarketResponse_Error struct {
	Error *common.Error `protobuf:"bytes,3,opt,name=error,proto3,oneof"`
}

func (*GetMarketResponse_Data_) isGetMarketResponse_Response() {}

func (*GetMarketResponse_Error) isGetMarketResponse_Response() {}

///////////////////////////////////////////////////////////////////
// Update Market
///////////////////////////////////////////////////////////////////
type UpsertMarketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MarketId   string          `protobuf:"bytes,1,opt,name=market_id,json=marketId,proto3" json:"market_id"`
	Name       *string         `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name"`
	Location   *MarketLocation `protobuf:"bytes,3,opt,name=location,proto3,oneof" json:"location"`
	Supervisor *Supervisor     `protobuf:"bytes,4,opt,name=supervisor,proto3,oneof" json:"supervisor"`
	Type       *MarketType     `protobuf:"varint,5,opt,name=type,proto3,enum=market.MarketType,oneof" json:"type"`
	Clazz      *MarketClass    `protobuf:"varint,6,opt,name=clazz,proto3,enum=market.MarketClass,oneof" json:"clazz"`
	Status     *MarketStatus   `protobuf:"varint,7,opt,name=status,proto3,enum=market.MarketStatus,oneof" json:"status"`
	GoogleMap  *string         `protobuf:"bytes,8,opt,name=google_map,json=googleMap,proto3,oneof" json:"google_map"`
}

func (x *UpsertMarketRequest) Reset() {
	*x = UpsertMarketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_market_market_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertMarketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertMarketRequest) ProtoMessage() {}

func (x *UpsertMarketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_market_market_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertMarketRequest.ProtoReflect.Descriptor instead.
func (*UpsertMarketRequest) Descriptor() ([]byte, []int) {
	return file_market_market_proto_rawDescGZIP(), []int{4}
}

func (x *UpsertMarketRequest) GetMarketId() string {
	if x != nil {
		return x.MarketId
	}
	return ""
}

func (x *UpsertMarketRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *UpsertMarketRequest) GetLocation() *MarketLocation {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *UpsertMarketRequest) GetSupervisor() *Supervisor {
	if x != nil {
		return x.Supervisor
	}
	return nil
}

func (x *UpsertMarketRequest) GetType() MarketType {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return MarketType_MARKET_TYPE_UNSPECIFIED
}

func (x *UpsertMarketRequest) GetClazz() MarketClass {
	if x != nil && x.Clazz != nil {
		return *x.Clazz
	}
	return MarketClass_MARKET_CLASS_UNSPECIFIED
}

func (x *UpsertMarketRequest) GetStatus() MarketStatus {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return MarketStatus_MARKET_STATUS_UNSPECIFIED
}

func (x *UpsertMarketRequest) GetGoogleMap() string {
	if x != nil && x.GoogleMap != nil {
		return *x.GoogleMap
	}
	return ""
}

type UpdateMarketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success"`
	// Types that are assignable to Response:
	//	*UpdateMarketResponse_Data_
	//	*UpdateMarketResponse_Error
	Response isUpdateMarketResponse_Response `protobuf_oneof:"response"`
}

func (x *UpdateMarketResponse) Reset() {
	*x = UpdateMarketResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_market_market_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMarketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMarketResponse) ProtoMessage() {}

func (x *UpdateMarketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_market_market_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMarketResponse.ProtoReflect.Descriptor instead.
func (*UpdateMarketResponse) Descriptor() ([]byte, []int) {
	return file_market_market_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateMarketResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (m *UpdateMarketResponse) GetResponse() isUpdateMarketResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (x *UpdateMarketResponse) GetData() *UpdateMarketResponse_Data {
	if x, ok := x.GetResponse().(*UpdateMarketResponse_Data_); ok {
		return x.Data
	}
	return nil
}

func (x *UpdateMarketResponse) GetError() *common.Error {
	if x, ok := x.GetResponse().(*UpdateMarketResponse_Error); ok {
		return x.Error
	}
	return nil
}

type isUpdateMarketResponse_Response interface {
	isUpdateMarketResponse_Response()
}

type UpdateMarketResponse_Data_ struct {
	Data *UpdateMarketResponse_Data `protobuf:"bytes,2,opt,name=data,proto3,oneof"`
}

type UpdateMarketResponse_Error struct {
	Error *common.Error `protobuf:"bytes,3,opt,name=error,proto3,oneof"`
}

func (*UpdateMarketResponse_Data_) isUpdateMarketResponse_Response() {}

func (*UpdateMarketResponse_Error) isUpdateMarketResponse_Response() {}

type GetMarketResponse_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Market *Market `protobuf:"bytes,1,opt,name=market,proto3" json:"market"`
}

func (x *GetMarketResponse_Data) Reset() {
	*x = GetMarketResponse_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_market_market_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMarketResponse_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMarketResponse_Data) ProtoMessage() {}

func (x *GetMarketResponse_Data) ProtoReflect() protoreflect.Message {
	mi := &file_market_market_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMarketResponse_Data.ProtoReflect.Descriptor instead.
func (*GetMarketResponse_Data) Descriptor() ([]byte, []int) {
	return file_market_market_proto_rawDescGZIP(), []int{3, 0}
}

func (x *GetMarketResponse_Data) GetMarket() *Market {
	if x != nil {
		return x.Market
	}
	return nil
}

type UpdateMarketResponse_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MarketId string `protobuf:"bytes,1,opt,name=market_id,json=marketId,proto3" json:"market_id"`
}

func (x *UpdateMarketResponse_Data) Reset() {
	*x = UpdateMarketResponse_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_market_market_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMarketResponse_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMarketResponse_Data) ProtoMessage() {}

func (x *UpdateMarketResponse_Data) ProtoReflect() protoreflect.Message {
	mi := &file_market_market_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMarketResponse_Data.ProtoReflect.Descriptor instead.
func (*UpdateMarketResponse_Data) Descriptor() ([]byte, []int) {
	return file_market_market_proto_rawDescGZIP(), []int{5, 0}
}

func (x *UpdateMarketResponse_Data) GetMarketId() string {
	if x != nil {
		return x.MarketId
	}
	return ""
}

var File_market_market_proto protoreflect.FileDescriptor

var file_market_market_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x1a, 0x13, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x15, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x6d, 0x61, 0x72, 0x6b, 0x65,
	0x74, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2f, 0x73, 0x75, 0x70, 0x65, 0x72, 0x76, 0x69,
	0x73, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6e, 0x0a, 0x0e, 0x4d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x77,
	0x61, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x77, 0x61, 0x72, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0xa7, 0x04, 0x0a, 0x06, 0x4d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x12, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x6d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x29, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65,
	0x74, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x68, 0x61, 0x73, 0x5f, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x68, 0x61, 0x73,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x68, 0x61, 0x73,
	0x5f, 0x64, 0x72, 0x61, 0x66, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x68, 0x61,
	0x73, 0x44, 0x72, 0x61, 0x66, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x75,
	0x6c, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x32, 0x0a, 0x0a, 0x73, 0x75, 0x70,
	0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x53, 0x75, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f,
	0x72, 0x52, 0x0a, 0x73, 0x75, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x12, 0x25, 0x0a,
	0x0e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x6d,
	0x61, 0x70, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x4d, 0x61, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x29, 0x0a, 0x05, 0x63, 0x6c, 0x61, 0x7a, 0x7a,
	0x18, 0x12, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e,
	0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x05, 0x63, 0x6c, 0x61,
	0x7a, 0x7a, 0x12, 0x1f, 0x0a, 0x0b, 0x68, 0x61, 0x73, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x18, 0x14, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x68, 0x61, 0x73, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x22, 0x68, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x72, 0x6b, 0x65,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x05, 0x64, 0x72, 0x61, 0x66, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x05, 0x64, 0x72, 0x61, 0x66, 0x74, 0x88, 0x01, 0x01, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x64, 0x72, 0x61, 0x66, 0x74, 0x22, 0xc6, 0x01,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x34, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x25, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x1a, 0x2e, 0x0a, 0x04, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x26, 0x0a, 0x06, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x4d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x52, 0x06, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x42, 0x0a, 0x0a, 0x08, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xc3, 0x03, 0x0a, 0x13, 0x55, 0x70, 0x73, 0x65, 0x72,
	0x74, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x37, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e,
	0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x01,
	0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x37, 0x0a,
	0x0a, 0x73, 0x75, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x53, 0x75, 0x70, 0x65, 0x72,
	0x76, 0x69, 0x73, 0x6f, 0x72, 0x48, 0x02, 0x52, 0x0a, 0x73, 0x75, 0x70, 0x65, 0x72, 0x76, 0x69,
	0x73, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x12, 0x2b, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x4d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x48, 0x03, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x2e, 0x0a, 0x05, 0x63, 0x6c, 0x61, 0x7a, 0x7a, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x13, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x4d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x48, 0x04, 0x52, 0x05, 0x63, 0x6c, 0x61, 0x7a, 0x7a,
	0x88, 0x01, 0x01, 0x12, 0x31, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x4d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x5f, 0x6d, 0x61, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x06, 0x52, 0x09, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x4d, 0x61, 0x70, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x73, 0x75, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x42,
	0x07, 0x0a, 0x05, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x63, 0x6c, 0x61,
	0x7a, 0x7a, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x0d, 0x0a,
	0x0b, 0x5f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x6d, 0x61, 0x70, 0x22, 0xc1, 0x01, 0x0a,
	0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12,
	0x37, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x25, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x1a,
	0x23, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x72, 0x6b, 0x65,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x49, 0x64, 0x42, 0x0a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x32, 0xee, 0x02, 0x0a, 0x0d, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x41, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x73, 0x12, 0x1a, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x12, 0x1b, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x55,
	0x70, 0x73, 0x65, 0x72, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4f, 0x6e, 0x6c, 0x79,
	0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x09, 0x47, 0x65,
	0x74, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x12, 0x18, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4b,
	0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x12, 0x1b,
	0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x4d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x72, 0x6b, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0c, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x12, 0x17, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4e, 0x6f,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x31, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x72, 0x73, 0x2e, 0x6d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x50, 0x01, 0x5a, 0x1d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2d, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x6d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_market_market_proto_rawDescOnce sync.Once
	file_market_market_proto_rawDescData = file_market_market_proto_rawDesc
)

func file_market_market_proto_rawDescGZIP() []byte {
	file_market_market_proto_rawDescOnce.Do(func() {
		file_market_market_proto_rawDescData = protoimpl.X.CompressGZIP(file_market_market_proto_rawDescData)
	})
	return file_market_market_proto_rawDescData
}

var file_market_market_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_market_market_proto_goTypes = []interface{}{
	(*MarketLocation)(nil),            // 0: market.MarketLocation
	(*Market)(nil),                    // 1: market.Market
	(*GetMarketRequest)(nil),          // 2: market.GetMarketRequest
	(*GetMarketResponse)(nil),         // 3: market.GetMarketResponse
	(*UpsertMarketRequest)(nil),       // 4: market.UpsertMarketRequest
	(*UpdateMarketResponse)(nil),      // 5: market.UpdateMarketResponse
	(*GetMarketResponse_Data)(nil),    // 6: market.GetMarketResponse.Data
	(*UpdateMarketResponse_Data)(nil), // 7: market.UpdateMarketResponse.Data
	(*Location)(nil),                  // 8: market.Location
	(MarketType)(0),                   // 9: market.MarketType
	(MarketStatus)(0),                 // 10: market.MarketStatus
	(MarketState)(0),                  // 11: market.MarketState
	(*Supervisor)(nil),                // 12: market.Supervisor
	(MarketClass)(0),                  // 13: market.MarketClass
	(*common.Error)(nil),              // 14: common.Error
	(*ListMarketsRequest)(nil),        // 15: market.ListMarketsRequest
	(*common.FindByIdRequest)(nil),    // 16: common.FindByIdRequest
	(*common.PageResponse)(nil),       // 17: common.PageResponse
	(*common.OnlyIdResponse)(nil),     // 18: common.OnlyIdResponse
	(*common.NoContentResponse)(nil),  // 19: common.NoContentResponse
}
var file_market_market_proto_depIdxs = []int32{
	8,  // 0: market.Market.location:type_name -> market.Location
	9,  // 1: market.Market.type:type_name -> market.MarketType
	10, // 2: market.Market.status:type_name -> market.MarketStatus
	11, // 3: market.Market.state:type_name -> market.MarketState
	12, // 4: market.Market.supervisor:type_name -> market.Supervisor
	13, // 5: market.Market.clazz:type_name -> market.MarketClass
	6,  // 6: market.GetMarketResponse.data:type_name -> market.GetMarketResponse.Data
	14, // 7: market.GetMarketResponse.error:type_name -> common.Error
	0,  // 8: market.UpsertMarketRequest.location:type_name -> market.MarketLocation
	12, // 9: market.UpsertMarketRequest.supervisor:type_name -> market.Supervisor
	9,  // 10: market.UpsertMarketRequest.type:type_name -> market.MarketType
	13, // 11: market.UpsertMarketRequest.clazz:type_name -> market.MarketClass
	10, // 12: market.UpsertMarketRequest.status:type_name -> market.MarketStatus
	7,  // 13: market.UpdateMarketResponse.data:type_name -> market.UpdateMarketResponse.Data
	14, // 14: market.UpdateMarketResponse.error:type_name -> common.Error
	1,  // 15: market.GetMarketResponse.Data.market:type_name -> market.Market
	15, // 16: market.MarketService.ListMarkets:input_type -> market.ListMarketsRequest
	4,  // 17: market.MarketService.CreateMarket:input_type -> market.UpsertMarketRequest
	2,  // 18: market.MarketService.GetMarket:input_type -> market.GetMarketRequest
	4,  // 19: market.MarketService.UpdateMarket:input_type -> market.UpsertMarketRequest
	16, // 20: market.MarketService.DeleteMarket:input_type -> common.FindByIdRequest
	17, // 21: market.MarketService.ListMarkets:output_type -> common.PageResponse
	18, // 22: market.MarketService.CreateMarket:output_type -> common.OnlyIdResponse
	3,  // 23: market.MarketService.GetMarket:output_type -> market.GetMarketResponse
	5,  // 24: market.MarketService.UpdateMarket:output_type -> market.UpdateMarketResponse
	19, // 25: market.MarketService.DeleteMarket:output_type -> common.NoContentResponse
	21, // [21:26] is the sub-list for method output_type
	16, // [16:21] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_market_market_proto_init() }
func file_market_market_proto_init() {
	if File_market_market_proto != nil {
		return
	}
	file_market_constant_proto_init()
	file_market_location_proto_init()
	file_market_supervisor_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_market_market_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarketLocation); i {
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
		file_market_market_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Market); i {
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
		file_market_market_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMarketRequest); i {
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
		file_market_market_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMarketResponse); i {
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
		file_market_market_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpsertMarketRequest); i {
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
		file_market_market_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMarketResponse); i {
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
		file_market_market_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMarketResponse_Data); i {
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
		file_market_market_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMarketResponse_Data); i {
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
	file_market_market_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_market_market_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*GetMarketResponse_Data_)(nil),
		(*GetMarketResponse_Error)(nil),
	}
	file_market_market_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_market_market_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*UpdateMarketResponse_Data_)(nil),
		(*UpdateMarketResponse_Error)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_market_market_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_market_market_proto_goTypes,
		DependencyIndexes: file_market_market_proto_depIdxs,
		MessageInfos:      file_market_market_proto_msgTypes,
	}.Build()
	File_market_market_proto = out.File
	file_market_market_proto_rawDesc = nil
	file_market_market_proto_goTypes = nil
	file_market_market_proto_depIdxs = nil
}
