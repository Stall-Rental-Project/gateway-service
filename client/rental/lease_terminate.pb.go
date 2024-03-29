// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        v3.21.8
// source: rental/lease_terminate.proto

package rental

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

type LeaseTermination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TerminationId string            `protobuf:"bytes,1,opt,name=termination_id,json=terminationId,proto3" json:"termination_id"`
	ApplicationId string            `protobuf:"bytes,2,opt,name=application_id,json=applicationId,proto3" json:"application_id"`
	Reason        string            `protobuf:"bytes,3,opt,name=reason,proto3" json:"reason"`
	EndDate       string            `protobuf:"bytes,4,opt,name=end_date,json=endDate,proto3" json:"end_date"`
	Accepted      bool              `protobuf:"varint,5,opt,name=accepted,proto3" json:"accepted"`
	CreatedBy     string            `protobuf:"bytes,6,opt,name=created_by,json=createdBy,proto3" json:"created_by"`
	Status        TerminationStatus `protobuf:"varint,7,opt,name=status,proto3,enum=rental.TerminationStatus" json:"status"`
}

func (x *LeaseTermination) Reset() {
	*x = LeaseTermination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rental_lease_terminate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaseTermination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaseTermination) ProtoMessage() {}

func (x *LeaseTermination) ProtoReflect() protoreflect.Message {
	mi := &file_rental_lease_terminate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaseTermination.ProtoReflect.Descriptor instead.
func (*LeaseTermination) Descriptor() ([]byte, []int) {
	return file_rental_lease_terminate_proto_rawDescGZIP(), []int{0}
}

func (x *LeaseTermination) GetTerminationId() string {
	if x != nil {
		return x.TerminationId
	}
	return ""
}

func (x *LeaseTermination) GetApplicationId() string {
	if x != nil {
		return x.ApplicationId
	}
	return ""
}

func (x *LeaseTermination) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *LeaseTermination) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *LeaseTermination) GetAccepted() bool {
	if x != nil {
		return x.Accepted
	}
	return false
}

func (x *LeaseTermination) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *LeaseTermination) GetStatus() TerminationStatus {
	if x != nil {
		return x.Status
	}
	return TerminationStatus_T_REQUESTED
}

type GetLeaseTerminationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success"`
	// Types that are assignable to Response:
	//	*GetLeaseTerminationResponse_Data_
	//	*GetLeaseTerminationResponse_Error
	Response isGetLeaseTerminationResponse_Response `protobuf_oneof:"response"`
}

func (x *GetLeaseTerminationResponse) Reset() {
	*x = GetLeaseTerminationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rental_lease_terminate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLeaseTerminationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLeaseTerminationResponse) ProtoMessage() {}

func (x *GetLeaseTerminationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rental_lease_terminate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLeaseTerminationResponse.ProtoReflect.Descriptor instead.
func (*GetLeaseTerminationResponse) Descriptor() ([]byte, []int) {
	return file_rental_lease_terminate_proto_rawDescGZIP(), []int{1}
}

func (x *GetLeaseTerminationResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (m *GetLeaseTerminationResponse) GetResponse() isGetLeaseTerminationResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (x *GetLeaseTerminationResponse) GetData() *GetLeaseTerminationResponse_Data {
	if x, ok := x.GetResponse().(*GetLeaseTerminationResponse_Data_); ok {
		return x.Data
	}
	return nil
}

func (x *GetLeaseTerminationResponse) GetError() *common.Error {
	if x, ok := x.GetResponse().(*GetLeaseTerminationResponse_Error); ok {
		return x.Error
	}
	return nil
}

type isGetLeaseTerminationResponse_Response interface {
	isGetLeaseTerminationResponse_Response()
}

type GetLeaseTerminationResponse_Data_ struct {
	Data *GetLeaseTerminationResponse_Data `protobuf:"bytes,2,opt,name=data,proto3,oneof"`
}

type GetLeaseTerminationResponse_Error struct {
	Error *common.Error `protobuf:"bytes,3,opt,name=error,proto3,oneof"`
}

func (*GetLeaseTerminationResponse_Data_) isGetLeaseTerminationResponse_Response() {}

func (*GetLeaseTerminationResponse_Error) isGetLeaseTerminationResponse_Response() {}

type CreateLeaseTerminationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApplicationId string `protobuf:"bytes,1,opt,name=application_id,json=applicationId,proto3" json:"application_id"`
	Reason        string `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason"`
	EndDate       string `protobuf:"bytes,3,opt,name=end_date,json=endDate,proto3" json:"end_date"` // in form of yyyy-MM-ddTHH:mm:ssXXX
	Accepted      bool   `protobuf:"varint,4,opt,name=accepted,proto3" json:"accepted"`
}

func (x *CreateLeaseTerminationRequest) Reset() {
	*x = CreateLeaseTerminationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rental_lease_terminate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLeaseTerminationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLeaseTerminationRequest) ProtoMessage() {}

func (x *CreateLeaseTerminationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rental_lease_terminate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLeaseTerminationRequest.ProtoReflect.Descriptor instead.
func (*CreateLeaseTerminationRequest) Descriptor() ([]byte, []int) {
	return file_rental_lease_terminate_proto_rawDescGZIP(), []int{2}
}

func (x *CreateLeaseTerminationRequest) GetApplicationId() string {
	if x != nil {
		return x.ApplicationId
	}
	return ""
}

func (x *CreateLeaseTerminationRequest) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *CreateLeaseTerminationRequest) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *CreateLeaseTerminationRequest) GetAccepted() bool {
	if x != nil {
		return x.Accepted
	}
	return false
}

type ProceedLeaseTerminationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TerminationId string `protobuf:"bytes,1,opt,name=termination_id,json=terminationId,proto3" json:"termination_id"`
	Accepted      bool   `protobuf:"varint,2,opt,name=accepted,proto3" json:"accepted"`
}

func (x *ProceedLeaseTerminationRequest) Reset() {
	*x = ProceedLeaseTerminationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rental_lease_terminate_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProceedLeaseTerminationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProceedLeaseTerminationRequest) ProtoMessage() {}

func (x *ProceedLeaseTerminationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rental_lease_terminate_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProceedLeaseTerminationRequest.ProtoReflect.Descriptor instead.
func (*ProceedLeaseTerminationRequest) Descriptor() ([]byte, []int) {
	return file_rental_lease_terminate_proto_rawDescGZIP(), []int{3}
}

func (x *ProceedLeaseTerminationRequest) GetTerminationId() string {
	if x != nil {
		return x.TerminationId
	}
	return ""
}

func (x *ProceedLeaseTerminationRequest) GetAccepted() bool {
	if x != nil {
		return x.Accepted
	}
	return false
}

type GetLeaseTerminationResponse_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exist       bool              `protobuf:"varint,1,opt,name=exist,proto3" json:"exist"`
	Termination *LeaseTermination `protobuf:"bytes,2,opt,name=termination,proto3" json:"termination"`
}

func (x *GetLeaseTerminationResponse_Data) Reset() {
	*x = GetLeaseTerminationResponse_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rental_lease_terminate_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLeaseTerminationResponse_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLeaseTerminationResponse_Data) ProtoMessage() {}

func (x *GetLeaseTerminationResponse_Data) ProtoReflect() protoreflect.Message {
	mi := &file_rental_lease_terminate_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLeaseTerminationResponse_Data.ProtoReflect.Descriptor instead.
func (*GetLeaseTerminationResponse_Data) Descriptor() ([]byte, []int) {
	return file_rental_lease_terminate_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GetLeaseTerminationResponse_Data) GetExist() bool {
	if x != nil {
		return x.Exist
	}
	return false
}

func (x *GetLeaseTerminationResponse_Data) GetTermination() *LeaseTermination {
	if x != nil {
		return x.Termination
	}
	return nil
}

var File_rental_lease_terminate_proto protoreflect.FileDescriptor

var file_rental_lease_terminate_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2f, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x74,
	0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x72, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x72, 0x65, 0x6e,
	0x74, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x81, 0x02, 0x0a, 0x10, 0x4c, 0x65, 0x61, 0x73, 0x65, 0x54, 0x65, 0x72, 0x6d,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x65, 0x72, 0x6d, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x25,
	0x0a, 0x0e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x19, 0x0a,
	0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x65,
	0x70, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x61, 0x63, 0x63, 0x65,
	0x70, 0x74, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x62, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x42, 0x79, 0x12, 0x31, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x54, 0x65, 0x72,
	0x6d, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x84, 0x02, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x4c, 0x65,
	0x61, 0x73, 0x65, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x12, 0x3e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28,
	0x2e, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x61, 0x73, 0x65,
	0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x25, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x1a, 0x58, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x78, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05,
	0x65, 0x78, 0x69, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x0b, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x72, 0x65, 0x6e,
	0x74, 0x61, 0x6c, 0x2e, 0x4c, 0x65, 0x61, 0x73, 0x65, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x0a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x95, 0x01,
	0x0a, 0x1d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x65, 0x61, 0x73, 0x65, 0x54, 0x65, 0x72,
	0x6d, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x25, 0x0a, 0x0e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x19,
	0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x63, 0x63,
	0x65, 0x70, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x61, 0x63, 0x63,
	0x65, 0x70, 0x74, 0x65, 0x64, 0x22, 0x63, 0x0a, 0x1e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x65, 0x64,
	0x4c, 0x65, 0x61, 0x73, 0x65, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x65, 0x72, 0x6d, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x32, 0xae, 0x02, 0x0a, 0x17, 0x4c,
	0x65, 0x61, 0x73, 0x65, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x55, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x61,
	0x73, 0x65, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e,
	0x47, 0x65, 0x74, 0x4c, 0x65, 0x61, 0x73, 0x65, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5c, 0x0a,
	0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x65, 0x61, 0x73, 0x65, 0x54, 0x65, 0x72, 0x6d,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x2e, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x6c,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x65, 0x61, 0x73, 0x65, 0x54, 0x65, 0x72, 0x6d,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4e, 0x6f, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5e, 0x0a, 0x17, 0x50,
	0x72, 0x6f, 0x63, 0x65, 0x65, 0x64, 0x4c, 0x65, 0x61, 0x73, 0x65, 0x54, 0x65, 0x72, 0x6d, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x2e, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e,
	0x50, 0x72, 0x6f, 0x63, 0x65, 0x65, 0x64, 0x4c, 0x65, 0x61, 0x73, 0x65, 0x54, 0x65, 0x72, 0x6d,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4e, 0x6f, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x31, 0x0a, 0x0e, 0x63,
	0x6f, 0x6d, 0x2e, 0x73, 0x72, 0x73, 0x2e, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x50, 0x01, 0x5a,
	0x1d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rental_lease_terminate_proto_rawDescOnce sync.Once
	file_rental_lease_terminate_proto_rawDescData = file_rental_lease_terminate_proto_rawDesc
)

func file_rental_lease_terminate_proto_rawDescGZIP() []byte {
	file_rental_lease_terminate_proto_rawDescOnce.Do(func() {
		file_rental_lease_terminate_proto_rawDescData = protoimpl.X.CompressGZIP(file_rental_lease_terminate_proto_rawDescData)
	})
	return file_rental_lease_terminate_proto_rawDescData
}

var file_rental_lease_terminate_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_rental_lease_terminate_proto_goTypes = []interface{}{
	(*LeaseTermination)(nil),                 // 0: rental.LeaseTermination
	(*GetLeaseTerminationResponse)(nil),      // 1: rental.GetLeaseTerminationResponse
	(*CreateLeaseTerminationRequest)(nil),    // 2: rental.CreateLeaseTerminationRequest
	(*ProceedLeaseTerminationRequest)(nil),   // 3: rental.ProceedLeaseTerminationRequest
	(*GetLeaseTerminationResponse_Data)(nil), // 4: rental.GetLeaseTerminationResponse.Data
	(TerminationStatus)(0),                   // 5: rental.TerminationStatus
	(*common.Error)(nil),                     // 6: common.Error
	(*common.FindByIdRequest)(nil),           // 7: common.FindByIdRequest
	(*common.NoContentResponse)(nil),         // 8: common.NoContentResponse
}
var file_rental_lease_terminate_proto_depIdxs = []int32{
	5, // 0: rental.LeaseTermination.status:type_name -> rental.TerminationStatus
	4, // 1: rental.GetLeaseTerminationResponse.data:type_name -> rental.GetLeaseTerminationResponse.Data
	6, // 2: rental.GetLeaseTerminationResponse.error:type_name -> common.Error
	0, // 3: rental.GetLeaseTerminationResponse.Data.termination:type_name -> rental.LeaseTermination
	7, // 4: rental.LeaseTerminationService.GetLeaseTermination:input_type -> common.FindByIdRequest
	2, // 5: rental.LeaseTerminationService.CreateLeaseTermination:input_type -> rental.CreateLeaseTerminationRequest
	3, // 6: rental.LeaseTerminationService.ProceedLeaseTermination:input_type -> rental.ProceedLeaseTerminationRequest
	1, // 7: rental.LeaseTerminationService.GetLeaseTermination:output_type -> rental.GetLeaseTerminationResponse
	8, // 8: rental.LeaseTerminationService.CreateLeaseTermination:output_type -> common.NoContentResponse
	8, // 9: rental.LeaseTerminationService.ProceedLeaseTermination:output_type -> common.NoContentResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_rental_lease_terminate_proto_init() }
func file_rental_lease_terminate_proto_init() {
	if File_rental_lease_terminate_proto != nil {
		return
	}
	file_rental_constant_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rental_lease_terminate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaseTermination); i {
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
		file_rental_lease_terminate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLeaseTerminationResponse); i {
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
		file_rental_lease_terminate_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLeaseTerminationRequest); i {
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
		file_rental_lease_terminate_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProceedLeaseTerminationRequest); i {
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
		file_rental_lease_terminate_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLeaseTerminationResponse_Data); i {
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
	file_rental_lease_terminate_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*GetLeaseTerminationResponse_Data_)(nil),
		(*GetLeaseTerminationResponse_Error)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rental_lease_terminate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rental_lease_terminate_proto_goTypes,
		DependencyIndexes: file_rental_lease_terminate_proto_depIdxs,
		MessageInfos:      file_rental_lease_terminate_proto_msgTypes,
	}.Build()
	File_rental_lease_terminate_proto = out.File
	file_rental_lease_terminate_proto_rawDesc = nil
	file_rental_lease_terminate_proto_goTypes = nil
	file_rental_lease_terminate_proto_depIdxs = nil
}
