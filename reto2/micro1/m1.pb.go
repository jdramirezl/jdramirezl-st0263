// Code generated by protoc-gen-go. DO NOT EDIT.
// source: m1.proto

package m1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetFilesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFilesRequest) Reset()         { *m = GetFilesRequest{} }
func (m *GetFilesRequest) String() string { return proto.CompactTextString(m) }
func (*GetFilesRequest) ProtoMessage()    {}
func (*GetFilesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bb41de714c835fe, []int{0}
}

func (m *GetFilesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFilesRequest.Unmarshal(m, b)
}
func (m *GetFilesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFilesRequest.Marshal(b, m, deterministic)
}
func (m *GetFilesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFilesRequest.Merge(m, src)
}
func (m *GetFilesRequest) XXX_Size() int {
	return xxx_messageInfo_GetFilesRequest.Size(m)
}
func (m *GetFilesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFilesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFilesRequest proto.InternalMessageInfo

type GetFileRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFileRequest) Reset()         { *m = GetFileRequest{} }
func (m *GetFileRequest) String() string { return proto.CompactTextString(m) }
func (*GetFileRequest) ProtoMessage()    {}
func (*GetFileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bb41de714c835fe, []int{1}
}

func (m *GetFileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFileRequest.Unmarshal(m, b)
}
func (m *GetFileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFileRequest.Marshal(b, m, deterministic)
}
func (m *GetFileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFileRequest.Merge(m, src)
}
func (m *GetFileRequest) XXX_Size() int {
	return xxx_messageInfo_GetFileRequest.Size(m)
}
func (m *GetFileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFileRequest proto.InternalMessageInfo

func (m *GetFileRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetFileResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFileResponse) Reset()         { *m = GetFileResponse{} }
func (m *GetFileResponse) String() string { return proto.CompactTextString(m) }
func (*GetFileResponse) ProtoMessage()    {}
func (*GetFileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bb41de714c835fe, []int{2}
}

func (m *GetFileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFileResponse.Unmarshal(m, b)
}
func (m *GetFileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFileResponse.Marshal(b, m, deterministic)
}
func (m *GetFileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFileResponse.Merge(m, src)
}
func (m *GetFileResponse) XXX_Size() int {
	return xxx_messageInfo_GetFileResponse.Size(m)
}
func (m *GetFileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetFileResponse proto.InternalMessageInfo

func (m *GetFileResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetFileResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetFilesResponse struct {
	Files                []*GetFileResponse `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GetFilesResponse) Reset()         { *m = GetFilesResponse{} }
func (m *GetFilesResponse) String() string { return proto.CompactTextString(m) }
func (*GetFilesResponse) ProtoMessage()    {}
func (*GetFilesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bb41de714c835fe, []int{3}
}

func (m *GetFilesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFilesResponse.Unmarshal(m, b)
}
func (m *GetFilesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFilesResponse.Marshal(b, m, deterministic)
}
func (m *GetFilesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFilesResponse.Merge(m, src)
}
func (m *GetFilesResponse) XXX_Size() int {
	return xxx_messageInfo_GetFilesResponse.Size(m)
}
func (m *GetFilesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFilesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetFilesResponse proto.InternalMessageInfo

func (m *GetFilesResponse) GetFiles() []*GetFileResponse {
	if m != nil {
		return m.Files
	}
	return nil
}

func init() {
	proto.RegisterType((*GetFilesRequest)(nil), "m1.GetFilesRequest")
	proto.RegisterType((*GetFileRequest)(nil), "m1.GetFileRequest")
	proto.RegisterType((*GetFileResponse)(nil), "m1.GetFileResponse")
	proto.RegisterType((*GetFilesResponse)(nil), "m1.GetFilesResponse")
}

func init() {
	proto.RegisterFile("m1.proto", fileDescriptor_9bb41de714c835fe)
}

var fileDescriptor_9bb41de714c835fe = []byte{
	// 203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xc8, 0x35, 0xd4, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xca, 0x35, 0x54, 0x12, 0xe4, 0xe2, 0x77, 0x4f, 0x2d, 0x71,
	0xcb, 0xcc, 0x49, 0x2d, 0x0e, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x51, 0x52, 0xe1, 0xe2, 0x83,
	0x0a, 0x41, 0x45, 0x84, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35,
	0x38, 0x83, 0xc0, 0x6c, 0x25, 0x53, 0xb8, 0xc6, 0xa0, 0xd4, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54,
	0x21, 0x3e, 0x2e, 0xa6, 0xcc, 0x14, 0xa8, 0x22, 0xa6, 0xcc, 0x14, 0xb8, 0x36, 0x26, 0x24, 0x6d,
	0xb6, 0x5c, 0x02, 0x08, 0xfb, 0xa0, 0xfa, 0x34, 0xb9, 0x58, 0xd3, 0x40, 0x02, 0x12, 0x8c, 0x0a,
	0xcc, 0x1a, 0xdc, 0x46, 0xc2, 0x7a, 0xb9, 0x86, 0x7a, 0x68, 0x66, 0x07, 0x41, 0x54, 0x18, 0xd5,
	0x70, 0x71, 0x83, 0x84, 0x83, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0x85, 0xcc, 0xb9, 0x38, 0x60,
	0xa6, 0x09, 0x21, 0x6b, 0x83, 0xf9, 0x45, 0x4a, 0x04, 0x55, 0x10, 0x62, 0x98, 0x12, 0x83, 0x90,
	0x09, 0x17, 0x3b, 0x54, 0x54, 0x48, 0x08, 0xc5, 0x3a, 0x88, 0x36, 0x6c, 0x4e, 0x50, 0x62, 0x70,
	0x62, 0x8f, 0x62, 0xd5, 0xd3, 0xb7, 0xce, 0x35, 0x4c, 0x62, 0x03, 0x07, 0xa0, 0x31, 0x20, 0x00,
	0x00, 0xff, 0xff, 0xc5, 0x69, 0x6c, 0x46, 0x4c, 0x01, 0x00, 0x00,
}
