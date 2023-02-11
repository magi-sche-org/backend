// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: proto/event.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Answer_ProposedSchedule_Availability int32

const (
	Answer_ProposedSchedule_UNSPECIFIED Answer_ProposedSchedule_Availability = 0
	Answer_ProposedSchedule_AVAILABLE   Answer_ProposedSchedule_Availability = 1
	Answer_ProposedSchedule_MAYBE       Answer_ProposedSchedule_Availability = 2
	Answer_ProposedSchedule_UNAVAILABLE Answer_ProposedSchedule_Availability = 3
)

// Enum value maps for Answer_ProposedSchedule_Availability.
var (
	Answer_ProposedSchedule_Availability_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "AVAILABLE",
		2: "MAYBE",
		3: "UNAVAILABLE",
	}
	Answer_ProposedSchedule_Availability_value = map[string]int32{
		"UNSPECIFIED": 0,
		"AVAILABLE":   1,
		"MAYBE":       2,
		"UNAVAILABLE": 3,
	}
)

func (x Answer_ProposedSchedule_Availability) Enum() *Answer_ProposedSchedule_Availability {
	p := new(Answer_ProposedSchedule_Availability)
	*p = x
	return p
}

func (x Answer_ProposedSchedule_Availability) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Answer_ProposedSchedule_Availability) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_event_proto_enumTypes[0].Descriptor()
}

func (Answer_ProposedSchedule_Availability) Type() protoreflect.EnumType {
	return &file_proto_event_proto_enumTypes[0]
}

func (x Answer_ProposedSchedule_Availability) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Answer_ProposedSchedule_Availability.Descriptor instead.
func (Answer_ProposedSchedule_Availability) EnumDescriptor() ([]byte, []int) {
	return file_proto_event_proto_rawDescGZIP(), []int{6, 0, 0}
}

type GetEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`       // イベントid
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"` // 認証用トークン
}

func (x *GetEventRequest) Reset() {
	*x = GetEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventRequest) ProtoMessage() {}

func (x *GetEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventRequest.ProtoReflect.Descriptor instead.
func (*GetEventRequest) Descriptor() ([]byte, []int) {
	return file_proto_event_proto_rawDescGZIP(), []int{0}
}

func (x *GetEventRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetEventRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type GetEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string                   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`             // イベントid
	Name              string                   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`         // イベント名
	Owner             bool                     `protobuf:"varint,3,opt,name=owner,proto3" json:"owner,omitempty"`      // イベント所有者か
	TimeUnit          *durationpb.Duration     `protobuf:"bytes,4,opt,name=timeUnit,proto3" json:"timeUnit,omitempty"` // 時間単位(秒)
	Duration          *durationpb.Duration     `protobuf:"bytes,5,opt,name=duration,proto3" json:"duration,omitempty"` // 所要時間
	Answers           []*Answer                `protobuf:"bytes,6,rep,name=answers,proto3" json:"answers,omitempty"`   // 参加者の解答
	ProposedStartTime []*timestamppb.Timestamp `protobuf:"bytes,7,rep,name=proposedStartTime,proto3" json:"proposedStartTime,omitempty"`
}

func (x *GetEventResponse) Reset() {
	*x = GetEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventResponse) ProtoMessage() {}

func (x *GetEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventResponse.ProtoReflect.Descriptor instead.
func (*GetEventResponse) Descriptor() ([]byte, []int) {
	return file_proto_event_proto_rawDescGZIP(), []int{1}
}

func (x *GetEventResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetEventResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetEventResponse) GetOwner() bool {
	if x != nil {
		return x.Owner
	}
	return false
}

func (x *GetEventResponse) GetTimeUnit() *durationpb.Duration {
	if x != nil {
		return x.TimeUnit
	}
	return nil
}

func (x *GetEventResponse) GetDuration() *durationpb.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

func (x *GetEventResponse) GetAnswers() []*Answer {
	if x != nil {
		return x.Answers
	}
	return nil
}

func (x *GetEventResponse) GetProposedStartTime() []*timestamppb.Timestamp {
	if x != nil {
		return x.ProposedStartTime
	}
	return nil
}

type CreateEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`         // イベント名
	Token    string               `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`       // 認証用トークン
	TimeUnit *durationpb.Duration `protobuf:"bytes,3,opt,name=timeUnit,proto3" json:"timeUnit,omitempty"` // 時間単位
	Duration *durationpb.Duration `protobuf:"bytes,4,opt,name=duration,proto3" json:"duration,omitempty"` // 所要時間
	// 候補の開始時間の配列
	ProposedStartTime []*timestamppb.Timestamp `protobuf:"bytes,5,rep,name=proposedStartTime,proto3" json:"proposedStartTime,omitempty"`
}

func (x *CreateEventRequest) Reset() {
	*x = CreateEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEventRequest) ProtoMessage() {}

func (x *CreateEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEventRequest.ProtoReflect.Descriptor instead.
func (*CreateEventRequest) Descriptor() ([]byte, []int) {
	return file_proto_event_proto_rawDescGZIP(), []int{2}
}

func (x *CreateEventRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateEventRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *CreateEventRequest) GetTimeUnit() *durationpb.Duration {
	if x != nil {
		return x.TimeUnit
	}
	return nil
}

func (x *CreateEventRequest) GetDuration() *durationpb.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

func (x *CreateEventRequest) GetProposedStartTime() []*timestamppb.Timestamp {
	if x != nil {
		return x.ProposedStartTime
	}
	return nil
}

type CreateEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventId string `protobuf:"bytes,2,opt,name=eventId,proto3" json:"eventId,omitempty"` // イベントID
}

func (x *CreateEventResponse) Reset() {
	*x = CreateEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_event_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEventResponse) ProtoMessage() {}

func (x *CreateEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEventResponse.ProtoReflect.Descriptor instead.
func (*CreateEventResponse) Descriptor() ([]byte, []int) {
	return file_proto_event_proto_rawDescGZIP(), []int{3}
}

func (x *CreateEventResponse) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

type RegisterAnswerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventId string  `protobuf:"bytes,1,opt,name=eventId,proto3" json:"eventId,omitempty"` // イベントID
	Token   string  `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`     // 認証トークン
	Answer  *Answer `protobuf:"bytes,3,opt,name=answer,proto3" json:"answer,omitempty"`   // 名前、備考、回答
}

func (x *RegisterAnswerRequest) Reset() {
	*x = RegisterAnswerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_event_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterAnswerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterAnswerRequest) ProtoMessage() {}

func (x *RegisterAnswerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterAnswerRequest.ProtoReflect.Descriptor instead.
func (*RegisterAnswerRequest) Descriptor() ([]byte, []int) {
	return file_proto_event_proto_rawDescGZIP(), []int{4}
}

func (x *RegisterAnswerRequest) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

func (x *RegisterAnswerRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *RegisterAnswerRequest) GetAnswer() *Answer {
	if x != nil {
		return x.Answer
	}
	return nil
}

type RegisterAnswerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RegisterAnswerResponse) Reset() {
	*x = RegisterAnswerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_event_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterAnswerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterAnswerResponse) ProtoMessage() {}

func (x *RegisterAnswerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterAnswerResponse.ProtoReflect.Descriptor instead.
func (*RegisterAnswerResponse) Descriptor() ([]byte, []int) {
	return file_proto_event_proto_rawDescGZIP(), []int{5}
}

type Answer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string                     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Note     string                     `protobuf:"bytes,2,opt,name=note,proto3" json:"note,omitempty"`
	Schedule []*Answer_ProposedSchedule `protobuf:"bytes,3,rep,name=schedule,proto3" json:"schedule,omitempty"`
}

func (x *Answer) Reset() {
	*x = Answer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_event_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Answer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Answer) ProtoMessage() {}

func (x *Answer) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Answer.ProtoReflect.Descriptor instead.
func (*Answer) Descriptor() ([]byte, []int) {
	return file_proto_event_proto_rawDescGZIP(), []int{6}
}

func (x *Answer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Answer) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

func (x *Answer) GetSchedule() []*Answer_ProposedSchedule {
	if x != nil {
		return x.Schedule
	}
	return nil
}

type Answer_ProposedSchedule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartTime    *timestamppb.Timestamp               `protobuf:"bytes,1,opt,name=startTime,proto3" json:"startTime,omitempty"`                                                           // 開始時間
	Availability Answer_ProposedSchedule_Availability `protobuf:"varint,2,opt,name=availability,proto3,enum=geekCamp.Answer_ProposedSchedule_Availability" json:"availability,omitempty"` // 参加可否
}

func (x *Answer_ProposedSchedule) Reset() {
	*x = Answer_ProposedSchedule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_event_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Answer_ProposedSchedule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Answer_ProposedSchedule) ProtoMessage() {}

func (x *Answer_ProposedSchedule) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Answer_ProposedSchedule.ProtoReflect.Descriptor instead.
func (*Answer_ProposedSchedule) Descriptor() ([]byte, []int) {
	return file_proto_event_proto_rawDescGZIP(), []int{6, 0}
}

func (x *Answer_ProposedSchedule) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *Answer_ProposedSchedule) GetAvailability() Answer_ProposedSchedule_Availability {
	if x != nil {
		return x.Availability
	}
	return Answer_ProposedSchedule_UNSPECIFIED
}

var File_proto_event_proto protoreflect.FileDescriptor

var file_proto_event_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x67, 0x65, 0x65, 0x6b, 0x43, 0x61, 0x6d, 0x70, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x37,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xb0, 0x02, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x35, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x55, 0x6e,
	0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x35, 0x0a,
	0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x07, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x65, 0x65, 0x6b, 0x43, 0x61, 0x6d, 0x70,
	0x2e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x52, 0x07, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73,
	0x12, 0x48, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x64, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x11, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65,
	0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xf6, 0x01, 0x0a, 0x12, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x35, 0x0a, 0x08, 0x74,
	0x69, 0x6d, 0x65, 0x55, 0x6e, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x55, 0x6e,
	0x69, 0x74, 0x12, 0x35, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x48, 0x0a, 0x11, 0x70, 0x72, 0x6f,
	0x70, 0x6f, 0x73, 0x65, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x11, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x22, 0x2f, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x22, 0x71, 0x0a, 0x15, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x28, 0x0a,
	0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x67, 0x65, 0x65, 0x6b, 0x43, 0x61, 0x6d, 0x70, 0x2e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x52,
	0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x22, 0x18, 0x0a, 0x16, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0xde, 0x02, 0x0a, 0x06, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x6f, 0x74, 0x65, 0x12, 0x3d, 0x0a, 0x08, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67, 0x65, 0x65, 0x6b, 0x43, 0x61, 0x6d,
	0x70, 0x2e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65,
	0x64, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x08, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x1a, 0xec, 0x01, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x64,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x52, 0x0a, 0x0c, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x67, 0x65, 0x65, 0x6b, 0x43,
	0x61, 0x6d, 0x70, 0x2e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x6f,
	0x73, 0x65, 0x64, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x41, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x0c, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x22, 0x4a, 0x0a, 0x0c, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x56, 0x41, 0x49, 0x4c,
	0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x4d, 0x41, 0x59, 0x42, 0x45, 0x10,
	0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45,
	0x10, 0x03, 0x32, 0xf1, 0x01, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x43, 0x0a, 0x08,
	0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x2e, 0x67, 0x65, 0x65, 0x6b, 0x43,
	0x61, 0x6d, 0x70, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x67, 0x65, 0x65, 0x6b, 0x43, 0x61, 0x6d, 0x70, 0x2e, 0x47,
	0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x4c, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x1c, 0x2e, 0x67, 0x65, 0x65, 0x6b, 0x43, 0x61, 0x6d, 0x70, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d,
	0x2e, 0x67, 0x65, 0x65, 0x6b, 0x43, 0x61, 0x6d, 0x70, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x55, 0x0a, 0x0e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x6e, 0x73, 0x77, 0x65,
	0x72, 0x12, 0x1f, 0x2e, 0x67, 0x65, 0x65, 0x6b, 0x43, 0x61, 0x6d, 0x70, 0x2e, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x67, 0x65, 0x65, 0x6b, 0x43, 0x61, 0x6d, 0x70, 0x2e, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_event_proto_rawDescOnce sync.Once
	file_proto_event_proto_rawDescData = file_proto_event_proto_rawDesc
)

func file_proto_event_proto_rawDescGZIP() []byte {
	file_proto_event_proto_rawDescOnce.Do(func() {
		file_proto_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_event_proto_rawDescData)
	})
	return file_proto_event_proto_rawDescData
}

var file_proto_event_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_event_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_event_proto_goTypes = []interface{}{
	(Answer_ProposedSchedule_Availability)(0), // 0: geekCamp.Answer.ProposedSchedule.Availability
	(*GetEventRequest)(nil),                   // 1: geekCamp.GetEventRequest
	(*GetEventResponse)(nil),                  // 2: geekCamp.GetEventResponse
	(*CreateEventRequest)(nil),                // 3: geekCamp.CreateEventRequest
	(*CreateEventResponse)(nil),               // 4: geekCamp.CreateEventResponse
	(*RegisterAnswerRequest)(nil),             // 5: geekCamp.RegisterAnswerRequest
	(*RegisterAnswerResponse)(nil),            // 6: geekCamp.RegisterAnswerResponse
	(*Answer)(nil),                            // 7: geekCamp.Answer
	(*Answer_ProposedSchedule)(nil),           // 8: geekCamp.Answer.ProposedSchedule
	(*durationpb.Duration)(nil),               // 9: google.protobuf.Duration
	(*timestamppb.Timestamp)(nil),             // 10: google.protobuf.Timestamp
}
var file_proto_event_proto_depIdxs = []int32{
	9,  // 0: geekCamp.GetEventResponse.timeUnit:type_name -> google.protobuf.Duration
	9,  // 1: geekCamp.GetEventResponse.duration:type_name -> google.protobuf.Duration
	7,  // 2: geekCamp.GetEventResponse.answers:type_name -> geekCamp.Answer
	10, // 3: geekCamp.GetEventResponse.proposedStartTime:type_name -> google.protobuf.Timestamp
	9,  // 4: geekCamp.CreateEventRequest.timeUnit:type_name -> google.protobuf.Duration
	9,  // 5: geekCamp.CreateEventRequest.duration:type_name -> google.protobuf.Duration
	10, // 6: geekCamp.CreateEventRequest.proposedStartTime:type_name -> google.protobuf.Timestamp
	7,  // 7: geekCamp.RegisterAnswerRequest.answer:type_name -> geekCamp.Answer
	8,  // 8: geekCamp.Answer.schedule:type_name -> geekCamp.Answer.ProposedSchedule
	10, // 9: geekCamp.Answer.ProposedSchedule.startTime:type_name -> google.protobuf.Timestamp
	0,  // 10: geekCamp.Answer.ProposedSchedule.availability:type_name -> geekCamp.Answer.ProposedSchedule.Availability
	1,  // 11: geekCamp.Event.GetEvent:input_type -> geekCamp.GetEventRequest
	3,  // 12: geekCamp.Event.CreateEvent:input_type -> geekCamp.CreateEventRequest
	5,  // 13: geekCamp.Event.RegisterAnswer:input_type -> geekCamp.RegisterAnswerRequest
	2,  // 14: geekCamp.Event.GetEvent:output_type -> geekCamp.GetEventResponse
	4,  // 15: geekCamp.Event.CreateEvent:output_type -> geekCamp.CreateEventResponse
	6,  // 16: geekCamp.Event.RegisterAnswer:output_type -> geekCamp.RegisterAnswerResponse
	14, // [14:17] is the sub-list for method output_type
	11, // [11:14] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_proto_event_proto_init() }
func file_proto_event_proto_init() {
	if File_proto_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventRequest); i {
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
		file_proto_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventResponse); i {
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
		file_proto_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEventRequest); i {
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
		file_proto_event_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEventResponse); i {
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
		file_proto_event_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterAnswerRequest); i {
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
		file_proto_event_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterAnswerResponse); i {
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
		file_proto_event_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Answer); i {
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
		file_proto_event_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Answer_ProposedSchedule); i {
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
			RawDescriptor: file_proto_event_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_event_proto_goTypes,
		DependencyIndexes: file_proto_event_proto_depIdxs,
		EnumInfos:         file_proto_event_proto_enumTypes,
		MessageInfos:      file_proto_event_proto_msgTypes,
	}.Build()
	File_proto_event_proto = out.File
	file_proto_event_proto_rawDesc = nil
	file_proto_event_proto_goTypes = nil
	file_proto_event_proto_depIdxs = nil
}