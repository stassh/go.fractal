// Code generated by protoc-gen-go. DO NOT EDIT.
// source: entity.proto

package entity

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

type FractalInfo struct {
	Rectangle            *Rectangle `protobuf:"bytes,1,opt,name=rectangle,proto3" json:"rectangle,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *FractalInfo) Reset()         { *m = FractalInfo{} }
func (m *FractalInfo) String() string { return proto.CompactTextString(m) }
func (*FractalInfo) ProtoMessage()    {}
func (*FractalInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf50d946d740d100, []int{0}
}

func (m *FractalInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FractalInfo.Unmarshal(m, b)
}
func (m *FractalInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FractalInfo.Marshal(b, m, deterministic)
}
func (m *FractalInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FractalInfo.Merge(m, src)
}
func (m *FractalInfo) XXX_Size() int {
	return xxx_messageInfo_FractalInfo.Size(m)
}
func (m *FractalInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FractalInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FractalInfo proto.InternalMessageInfo

func (m *FractalInfo) GetRectangle() *Rectangle {
	if m != nil {
		return m.Rectangle
	}
	return nil
}

type Rectangle struct {
	Top                  float64  `protobuf:"fixed64,1,opt,name=top,proto3" json:"top,omitempty"`
	Bottom               float64  `protobuf:"fixed64,2,opt,name=bottom,proto3" json:"bottom,omitempty"`
	Left                 float64  `protobuf:"fixed64,3,opt,name=left,proto3" json:"left,omitempty"`
	Right                float64  `protobuf:"fixed64,4,opt,name=right,proto3" json:"right,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Rectangle) Reset()         { *m = Rectangle{} }
func (m *Rectangle) String() string { return proto.CompactTextString(m) }
func (*Rectangle) ProtoMessage()    {}
func (*Rectangle) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf50d946d740d100, []int{1}
}

func (m *Rectangle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rectangle.Unmarshal(m, b)
}
func (m *Rectangle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rectangle.Marshal(b, m, deterministic)
}
func (m *Rectangle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rectangle.Merge(m, src)
}
func (m *Rectangle) XXX_Size() int {
	return xxx_messageInfo_Rectangle.Size(m)
}
func (m *Rectangle) XXX_DiscardUnknown() {
	xxx_messageInfo_Rectangle.DiscardUnknown(m)
}

var xxx_messageInfo_Rectangle proto.InternalMessageInfo

func (m *Rectangle) GetTop() float64 {
	if m != nil {
		return m.Top
	}
	return 0
}

func (m *Rectangle) GetBottom() float64 {
	if m != nil {
		return m.Bottom
	}
	return 0
}

func (m *Rectangle) GetLeft() float64 {
	if m != nil {
		return m.Left
	}
	return 0
}

func (m *Rectangle) GetRight() float64 {
	if m != nil {
		return m.Right
	}
	return 0
}

func init() {
	proto.RegisterType((*FractalInfo)(nil), "entity.FractalInfo")
	proto.RegisterType((*Rectangle)(nil), "entity.Rectangle")
}

func init() { proto.RegisterFile("entity.proto", fileDescriptor_cf50d946d740d100) }

var fileDescriptor_cf50d946d740d100 = []byte{
	// 154 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xcd, 0x2b, 0xc9,
	0x2c, 0xa9, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0x94, 0xec, 0xb8, 0xb8,
	0xdd, 0x8a, 0x12, 0x93, 0x4b, 0x12, 0x73, 0x3c, 0xf3, 0xd2, 0xf2, 0x85, 0xf4, 0xb9, 0x38, 0x8b,
	0x52, 0x93, 0x4b, 0x12, 0xf3, 0xd2, 0x73, 0x52, 0x25, 0x18, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0x04,
	0xf5, 0xa0, 0x1a, 0x83, 0x60, 0x12, 0x41, 0x08, 0x35, 0x4a, 0xf1, 0x5c, 0x9c, 0x70, 0x71, 0x21,
	0x01, 0x2e, 0xe6, 0x92, 0xfc, 0x02, 0xb0, 0x3e, 0xc6, 0x20, 0x10, 0x53, 0x48, 0x8c, 0x8b, 0x2d,
	0x29, 0xbf, 0xa4, 0x24, 0x3f, 0x57, 0x82, 0x09, 0x2c, 0x08, 0xe5, 0x09, 0x09, 0x71, 0xb1, 0xe4,
	0xa4, 0xa6, 0x95, 0x48, 0x30, 0x83, 0x45, 0xc1, 0x6c, 0x21, 0x11, 0x2e, 0xd6, 0xa2, 0xcc, 0xf4,
	0x8c, 0x12, 0x09, 0x16, 0xb0, 0x20, 0x84, 0x93, 0xc4, 0x06, 0x76, 0xaf, 0x31, 0x20, 0x00, 0x00,
	0xff, 0xff, 0x56, 0x15, 0xb2, 0x7d, 0xbf, 0x00, 0x00, 0x00,
}
