// Code generated by protoc-gen-go. DO NOT EDIT.
// source: UserService.proto

package micro_service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RegResponse struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegResponse) Reset()         { *m = RegResponse{} }
func (m *RegResponse) String() string { return proto.CompactTextString(m) }
func (*RegResponse) ProtoMessage()    {}
func (*RegResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_UserService_476f150760f6f7e7, []int{0}
}
func (m *RegResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegResponse.Unmarshal(m, b)
}
func (m *RegResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegResponse.Marshal(b, m, deterministic)
}
func (dst *RegResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegResponse.Merge(dst, src)
}
func (m *RegResponse) XXX_Size() int {
	return xxx_messageInfo_RegResponse.Size(m)
}
func (m *RegResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegResponse proto.InternalMessageInfo

func (m *RegResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *RegResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*RegResponse)(nil), "micro.service.RegResponse")
}

func init() { proto.RegisterFile("UserService.proto", fileDescriptor_UserService_476f150760f6f7e7) }

var fileDescriptor_UserService_476f150760f6f7e7 = []byte{
	// 174 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8e, 0xb1, 0x0a, 0xc2, 0x30,
	0x10, 0x86, 0xa9, 0x43, 0x8b, 0xa9, 0x0e, 0x66, 0x90, 0x10, 0x1c, 0xc4, 0xc9, 0x29, 0x05, 0x7d,
	0x80, 0x3e, 0x81, 0x0e, 0x11, 0x1f, 0x20, 0xd6, 0x23, 0x04, 0xda, 0x5c, 0xc9, 0x45, 0x9f, 0x5f,
	0x9a, 0x56, 0xa8, 0x8e, 0xdf, 0xff, 0x1d, 0xdc, 0xc7, 0x36, 0x77, 0x82, 0x70, 0x83, 0xf0, 0x76,
	0x0d, 0xa8, 0x3e, 0x60, 0x44, 0xbe, 0xee, 0x5c, 0x13, 0x50, 0xd1, 0x38, 0xca, 0x9d, 0x45, 0xb4,
	0x2d, 0x54, 0xa6, 0x77, 0x95, 0xf1, 0x1e, 0xa3, 0x89, 0x0e, 0x3d, 0x8d, 0xc7, 0x72, 0x75, 0xc1,
	0x27, 0xb4, 0x13, 0x1d, 0x6a, 0x56, 0x6a, 0xb0, 0x1a, 0xa8, 0x47, 0x4f, 0xc0, 0xb7, 0x2c, 0xa7,
	0x68, 0xe2, 0x8b, 0x44, 0xb6, 0xcf, 0x8e, 0x4b, 0x3d, 0x11, 0x17, 0xac, 0xe8, 0x80, 0xc8, 0x58,
	0x10, 0x8b, 0x24, 0xbe, 0x78, 0xba, 0xb2, 0x72, 0x16, 0xc4, 0x6b, 0x56, 0x0c, 0xa8, 0xc1, 0x72,
	0xa1, 0x7e, 0xb2, 0xd4, 0xb0, 0xa7, 0xdf, 0x52, 0xfe, 0x99, 0x59, 0xc1, 0x23, 0x4f, 0x5d, 0xe7,
	0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe2, 0xb3, 0x68, 0xc8, 0xe7, 0x00, 0x00, 0x00,
}
