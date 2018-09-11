// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/admin/event.proto

package admin // import "github.com/lyft/flyteidl/gen/pb-go/flyteidl/admin"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import core "github.com/lyft/flyteidl/gen/pb-go/flyteidl/core"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type WorkflowExecutionEventRequest struct {
	// Unique ID for this request that can be traced between services
	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// Workflow execution id
	ExecutionId string `protobuf:"bytes,2,opt,name=execution_id,json=executionId,proto3" json:"execution_id,omitempty"`
	// This represents the timestamp of when the event occured, not the request timestamp
	OccuredAt            *timestamp.Timestamp         `protobuf:"bytes,3,opt,name=occured_at,json=occuredAt,proto3" json:"occured_at,omitempty"`
	Event                *core.WorkflowExecutionEvent `protobuf:"bytes,4,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *WorkflowExecutionEventRequest) Reset()         { *m = WorkflowExecutionEventRequest{} }
func (m *WorkflowExecutionEventRequest) String() string { return proto.CompactTextString(m) }
func (*WorkflowExecutionEventRequest) ProtoMessage()    {}
func (*WorkflowExecutionEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_52d642ffdbff3fb0, []int{0}
}
func (m *WorkflowExecutionEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowExecutionEventRequest.Unmarshal(m, b)
}
func (m *WorkflowExecutionEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowExecutionEventRequest.Marshal(b, m, deterministic)
}
func (dst *WorkflowExecutionEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowExecutionEventRequest.Merge(dst, src)
}
func (m *WorkflowExecutionEventRequest) XXX_Size() int {
	return xxx_messageInfo_WorkflowExecutionEventRequest.Size(m)
}
func (m *WorkflowExecutionEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowExecutionEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowExecutionEventRequest proto.InternalMessageInfo

func (m *WorkflowExecutionEventRequest) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *WorkflowExecutionEventRequest) GetExecutionId() string {
	if m != nil {
		return m.ExecutionId
	}
	return ""
}

func (m *WorkflowExecutionEventRequest) GetOccuredAt() *timestamp.Timestamp {
	if m != nil {
		return m.OccuredAt
	}
	return nil
}

func (m *WorkflowExecutionEventRequest) GetEvent() *core.WorkflowExecutionEvent {
	if m != nil {
		return m.Event
	}
	return nil
}

type WorkflowExecutionEventResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkflowExecutionEventResponse) Reset()         { *m = WorkflowExecutionEventResponse{} }
func (m *WorkflowExecutionEventResponse) String() string { return proto.CompactTextString(m) }
func (*WorkflowExecutionEventResponse) ProtoMessage()    {}
func (*WorkflowExecutionEventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_52d642ffdbff3fb0, []int{1}
}
func (m *WorkflowExecutionEventResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowExecutionEventResponse.Unmarshal(m, b)
}
func (m *WorkflowExecutionEventResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowExecutionEventResponse.Marshal(b, m, deterministic)
}
func (dst *WorkflowExecutionEventResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowExecutionEventResponse.Merge(dst, src)
}
func (m *WorkflowExecutionEventResponse) XXX_Size() int {
	return xxx_messageInfo_WorkflowExecutionEventResponse.Size(m)
}
func (m *WorkflowExecutionEventResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowExecutionEventResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowExecutionEventResponse proto.InternalMessageInfo

type NodeExecutionEventRequest struct {
	// Unique ID for this request that can be traced between services
	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// Workflow execution id
	ExecutionId string `protobuf:"bytes,2,opt,name=execution_id,json=executionId,proto3" json:"execution_id,omitempty"`
	// This represents the timestamp of when the event occured, not the request timestamp
	OccuredAt            *timestamp.Timestamp     `protobuf:"bytes,3,opt,name=occured_at,json=occuredAt,proto3" json:"occured_at,omitempty"`
	Event                *core.NodeExecutionEvent `protobuf:"bytes,4,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *NodeExecutionEventRequest) Reset()         { *m = NodeExecutionEventRequest{} }
func (m *NodeExecutionEventRequest) String() string { return proto.CompactTextString(m) }
func (*NodeExecutionEventRequest) ProtoMessage()    {}
func (*NodeExecutionEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_52d642ffdbff3fb0, []int{2}
}
func (m *NodeExecutionEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeExecutionEventRequest.Unmarshal(m, b)
}
func (m *NodeExecutionEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeExecutionEventRequest.Marshal(b, m, deterministic)
}
func (dst *NodeExecutionEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeExecutionEventRequest.Merge(dst, src)
}
func (m *NodeExecutionEventRequest) XXX_Size() int {
	return xxx_messageInfo_NodeExecutionEventRequest.Size(m)
}
func (m *NodeExecutionEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeExecutionEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NodeExecutionEventRequest proto.InternalMessageInfo

func (m *NodeExecutionEventRequest) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *NodeExecutionEventRequest) GetExecutionId() string {
	if m != nil {
		return m.ExecutionId
	}
	return ""
}

func (m *NodeExecutionEventRequest) GetOccuredAt() *timestamp.Timestamp {
	if m != nil {
		return m.OccuredAt
	}
	return nil
}

func (m *NodeExecutionEventRequest) GetEvent() *core.NodeExecutionEvent {
	if m != nil {
		return m.Event
	}
	return nil
}

type NodeExecutionEventResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeExecutionEventResponse) Reset()         { *m = NodeExecutionEventResponse{} }
func (m *NodeExecutionEventResponse) String() string { return proto.CompactTextString(m) }
func (*NodeExecutionEventResponse) ProtoMessage()    {}
func (*NodeExecutionEventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_52d642ffdbff3fb0, []int{3}
}
func (m *NodeExecutionEventResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeExecutionEventResponse.Unmarshal(m, b)
}
func (m *NodeExecutionEventResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeExecutionEventResponse.Marshal(b, m, deterministic)
}
func (dst *NodeExecutionEventResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeExecutionEventResponse.Merge(dst, src)
}
func (m *NodeExecutionEventResponse) XXX_Size() int {
	return xxx_messageInfo_NodeExecutionEventResponse.Size(m)
}
func (m *NodeExecutionEventResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeExecutionEventResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NodeExecutionEventResponse proto.InternalMessageInfo

type TaskExecutionEventRequest struct {
	// Unique ID for this request that can be traced between services
	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// Workflow execution id
	ExecutionId string `protobuf:"bytes,2,opt,name=execution_id,json=executionId,proto3" json:"execution_id,omitempty"`
	// This represents the timestamp of when the event occured, not the request timestamp
	OccuredAt            *timestamp.Timestamp     `protobuf:"bytes,3,opt,name=occured_at,json=occuredAt,proto3" json:"occured_at,omitempty"`
	Event                *core.TaskExecutionEvent `protobuf:"bytes,4,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *TaskExecutionEventRequest) Reset()         { *m = TaskExecutionEventRequest{} }
func (m *TaskExecutionEventRequest) String() string { return proto.CompactTextString(m) }
func (*TaskExecutionEventRequest) ProtoMessage()    {}
func (*TaskExecutionEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_52d642ffdbff3fb0, []int{4}
}
func (m *TaskExecutionEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskExecutionEventRequest.Unmarshal(m, b)
}
func (m *TaskExecutionEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskExecutionEventRequest.Marshal(b, m, deterministic)
}
func (dst *TaskExecutionEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskExecutionEventRequest.Merge(dst, src)
}
func (m *TaskExecutionEventRequest) XXX_Size() int {
	return xxx_messageInfo_TaskExecutionEventRequest.Size(m)
}
func (m *TaskExecutionEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskExecutionEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TaskExecutionEventRequest proto.InternalMessageInfo

func (m *TaskExecutionEventRequest) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *TaskExecutionEventRequest) GetExecutionId() string {
	if m != nil {
		return m.ExecutionId
	}
	return ""
}

func (m *TaskExecutionEventRequest) GetOccuredAt() *timestamp.Timestamp {
	if m != nil {
		return m.OccuredAt
	}
	return nil
}

func (m *TaskExecutionEventRequest) GetEvent() *core.TaskExecutionEvent {
	if m != nil {
		return m.Event
	}
	return nil
}

type TaskExecutionEventResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskExecutionEventResponse) Reset()         { *m = TaskExecutionEventResponse{} }
func (m *TaskExecutionEventResponse) String() string { return proto.CompactTextString(m) }
func (*TaskExecutionEventResponse) ProtoMessage()    {}
func (*TaskExecutionEventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_52d642ffdbff3fb0, []int{5}
}
func (m *TaskExecutionEventResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskExecutionEventResponse.Unmarshal(m, b)
}
func (m *TaskExecutionEventResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskExecutionEventResponse.Marshal(b, m, deterministic)
}
func (dst *TaskExecutionEventResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskExecutionEventResponse.Merge(dst, src)
}
func (m *TaskExecutionEventResponse) XXX_Size() int {
	return xxx_messageInfo_TaskExecutionEventResponse.Size(m)
}
func (m *TaskExecutionEventResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskExecutionEventResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TaskExecutionEventResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*WorkflowExecutionEventRequest)(nil), "flyteidl.admin.WorkflowExecutionEventRequest")
	proto.RegisterType((*WorkflowExecutionEventResponse)(nil), "flyteidl.admin.WorkflowExecutionEventResponse")
	proto.RegisterType((*NodeExecutionEventRequest)(nil), "flyteidl.admin.NodeExecutionEventRequest")
	proto.RegisterType((*NodeExecutionEventResponse)(nil), "flyteidl.admin.NodeExecutionEventResponse")
	proto.RegisterType((*TaskExecutionEventRequest)(nil), "flyteidl.admin.TaskExecutionEventRequest")
	proto.RegisterType((*TaskExecutionEventResponse)(nil), "flyteidl.admin.TaskExecutionEventResponse")
}

func init() { proto.RegisterFile("flyteidl/admin/event.proto", fileDescriptor_event_52d642ffdbff3fb0) }

var fileDescriptor_event_52d642ffdbff3fb0 = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x93, 0xbf, 0x4b, 0xc3, 0x40,
	0x14, 0xc7, 0x89, 0xbf, 0x20, 0x57, 0x71, 0xc8, 0x94, 0x06, 0xab, 0x69, 0x40, 0xe8, 0xe2, 0x1d,
	0xda, 0x41, 0xc4, 0x49, 0xa1, 0x43, 0x17, 0x87, 0x50, 0x10, 0x5c, 0x4a, 0x92, 0x7b, 0x89, 0x47,
	0x93, 0xbc, 0x98, 0x5c, 0xd4, 0xfe, 0x99, 0x8e, 0xfe, 0x37, 0x92, 0xcb, 0x0f, 0x5b, 0x49, 0x77,
	0xdd, 0x8e, 0x7b, 0xdf, 0xef, 0xe3, 0xf3, 0x19, 0x1e, 0xb1, 0xc2, 0x78, 0x2d, 0x41, 0xf0, 0x98,
	0x79, 0x3c, 0x11, 0x29, 0x83, 0x37, 0x48, 0x25, 0xcd, 0x72, 0x94, 0x68, 0x9c, 0xb4, 0x33, 0xaa,
	0x66, 0xd6, 0x79, 0x84, 0x18, 0xc5, 0xc0, 0xd4, 0xd4, 0x2f, 0x43, 0x26, 0x45, 0x02, 0x85, 0xf4,
	0x92, 0xac, 0x2e, 0x58, 0xc3, 0x6e, 0x59, 0x80, 0x39, 0x6c, 0xee, 0x72, 0xbe, 0x34, 0x32, 0x7a,
	0xc2, 0x7c, 0x15, 0xc6, 0xf8, 0x3e, 0xfb, 0x80, 0xa0, 0x94, 0x02, 0xd3, 0x59, 0x15, 0x70, 0xe1,
	0xb5, 0x84, 0x42, 0x1a, 0x23, 0x42, 0xf2, 0xfa, 0xb9, 0x14, 0xdc, 0xd4, 0x6c, 0x6d, 0xa2, 0xbb,
	0x7a, 0xf3, 0x33, 0xe7, 0xc6, 0x98, 0x1c, 0x43, 0xdb, 0xab, 0x02, 0x7b, 0x2a, 0x30, 0xe8, 0xfe,
	0xe6, 0xdc, 0xb8, 0x25, 0x04, 0x83, 0xa0, 0xcc, 0x81, 0x2f, 0x3d, 0x69, 0xee, 0xdb, 0xda, 0x64,
	0x70, 0x6d, 0xd1, 0x1a, 0x9a, 0xb6, 0xd0, 0x74, 0xd1, 0x42, 0xbb, 0x7a, 0x93, 0xbe, 0x97, 0xc6,
	0x1d, 0x39, 0x54, 0xb4, 0xe6, 0x81, 0x6a, 0x5d, 0xd0, 0x4e, 0xbd, 0x32, 0xa1, 0x3b, 0xc8, 0xeb,
	0x8e, 0x63, 0x93, 0xb3, 0x5d, 0x6a, 0x45, 0x86, 0x69, 0x01, 0xce, 0xa7, 0x46, 0x86, 0x8f, 0xc8,
	0xe1, 0x0f, 0x9a, 0xdf, 0x6c, 0x9b, 0x8f, 0x7f, 0x99, 0xf7, 0x50, 0x37, 0xd6, 0xa7, 0xc4, 0xea,
	0x53, 0xda, 0x30, 0x5e, 0x78, 0xc5, 0xea, 0xff, 0x19, 0xf7, 0x50, 0xff, 0x18, 0xf7, 0x29, 0xd5,
	0xc6, 0x0f, 0xd3, 0xe7, 0xab, 0x48, 0xc8, 0x97, 0xd2, 0xa7, 0x01, 0x26, 0x2c, 0x5e, 0x87, 0x92,
	0x75, 0xe7, 0x10, 0x41, 0xca, 0x32, 0xff, 0x32, 0x42, 0xb6, 0x7d, 0x6e, 0xfe, 0x91, 0x42, 0x9d,
	0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0xcc, 0x89, 0x06, 0xe1, 0x87, 0x03, 0x00, 0x00,
}