// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test/service/rpc/proto/stream.proto

package stream

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

// Points are represented as latitude-longitude pairs in the E7 representation
// (degrees multiplied by 10**7 and rounded to the nearest integer).
// Latitudes should be in the range +/- 90 degrees and longitude should be in
// the range +/- 180 degrees (inclusive).
type Point struct {
	Latitude             int32    `protobuf:"varint,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude            int32    `protobuf:"varint,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Point) Reset()         { *m = Point{} }
func (m *Point) String() string { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()    {}
func (*Point) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a6280eeb8b3790e, []int{0}
}

func (m *Point) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Point.Unmarshal(m, b)
}
func (m *Point) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Point.Marshal(b, m, deterministic)
}
func (m *Point) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Point.Merge(m, src)
}
func (m *Point) XXX_Size() int {
	return xxx_messageInfo_Point.Size(m)
}
func (m *Point) XXX_DiscardUnknown() {
	xxx_messageInfo_Point.DiscardUnknown(m)
}

var xxx_messageInfo_Point proto.InternalMessageInfo

func (m *Point) GetLatitude() int32 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *Point) GetLongitude() int32 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

// A latitude-longitude rectangle, represented as two diagonally opposite
// points "lo" and "hi".
type Rectangle struct {
	// One corner of the rectangle.
	Lo *Point `protobuf:"bytes,1,opt,name=lo,proto3" json:"lo,omitempty"`
	// The other corner of the rectangle.
	Hi                   *Point   `protobuf:"bytes,2,opt,name=hi,proto3" json:"hi,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Rectangle) Reset()         { *m = Rectangle{} }
func (m *Rectangle) String() string { return proto.CompactTextString(m) }
func (*Rectangle) ProtoMessage()    {}
func (*Rectangle) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a6280eeb8b3790e, []int{1}
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

func (m *Rectangle) GetLo() *Point {
	if m != nil {
		return m.Lo
	}
	return nil
}

func (m *Rectangle) GetHi() *Point {
	if m != nil {
		return m.Hi
	}
	return nil
}

// A feature names something at a given point.
// If a feature could not be named, the name is empty.
type Feature struct {
	// The name of the feature.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The point where the feature is detected.
	Location             *Point   `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Feature) Reset()         { *m = Feature{} }
func (m *Feature) String() string { return proto.CompactTextString(m) }
func (*Feature) ProtoMessage()    {}
func (*Feature) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a6280eeb8b3790e, []int{2}
}

func (m *Feature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Feature.Unmarshal(m, b)
}
func (m *Feature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Feature.Marshal(b, m, deterministic)
}
func (m *Feature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Feature.Merge(m, src)
}
func (m *Feature) XXX_Size() int {
	return xxx_messageInfo_Feature.Size(m)
}
func (m *Feature) XXX_DiscardUnknown() {
	xxx_messageInfo_Feature.DiscardUnknown(m)
}

var xxx_messageInfo_Feature proto.InternalMessageInfo

func (m *Feature) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Feature) GetLocation() *Point {
	if m != nil {
		return m.Location
	}
	return nil
}

// A RouteNote is a message sent while at a given point.
type RouteNote struct {
	// The location from which the message is sent.
	Location *Point `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	// The message to be sent.
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteNote) Reset()         { *m = RouteNote{} }
func (m *RouteNote) String() string { return proto.CompactTextString(m) }
func (*RouteNote) ProtoMessage()    {}
func (*RouteNote) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a6280eeb8b3790e, []int{3}
}

func (m *RouteNote) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteNote.Unmarshal(m, b)
}
func (m *RouteNote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteNote.Marshal(b, m, deterministic)
}
func (m *RouteNote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteNote.Merge(m, src)
}
func (m *RouteNote) XXX_Size() int {
	return xxx_messageInfo_RouteNote.Size(m)
}
func (m *RouteNote) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteNote.DiscardUnknown(m)
}

var xxx_messageInfo_RouteNote proto.InternalMessageInfo

func (m *RouteNote) GetLocation() *Point {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *RouteNote) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// A RouteSummary is received in response to a RecordRoute rpc.
// It contains the number of individual points received, the number of
// detected features, and the total distance covered as the cumulative sum of
// the distance between each point.
type RouteSummary struct {
	// The number of points received.
	PointCount int32 `protobuf:"varint,1,opt,name=point_count,json=pointCount,proto3" json:"point_count,omitempty"`
	// The number of known features passed while traversing the route.
	FeatureCount int32 `protobuf:"varint,2,opt,name=feature_count,json=featureCount,proto3" json:"feature_count,omitempty"`
	// The distance covered in metres.
	Distance int32 `protobuf:"varint,3,opt,name=distance,proto3" json:"distance,omitempty"`
	// The duration of the traversal in seconds.
	ElapsedTime          int32    `protobuf:"varint,4,opt,name=elapsed_time,json=elapsedTime,proto3" json:"elapsed_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteSummary) Reset()         { *m = RouteSummary{} }
func (m *RouteSummary) String() string { return proto.CompactTextString(m) }
func (*RouteSummary) ProtoMessage()    {}
func (*RouteSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a6280eeb8b3790e, []int{4}
}

func (m *RouteSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteSummary.Unmarshal(m, b)
}
func (m *RouteSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteSummary.Marshal(b, m, deterministic)
}
func (m *RouteSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteSummary.Merge(m, src)
}
func (m *RouteSummary) XXX_Size() int {
	return xxx_messageInfo_RouteSummary.Size(m)
}
func (m *RouteSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteSummary.DiscardUnknown(m)
}

var xxx_messageInfo_RouteSummary proto.InternalMessageInfo

func (m *RouteSummary) GetPointCount() int32 {
	if m != nil {
		return m.PointCount
	}
	return 0
}

func (m *RouteSummary) GetFeatureCount() int32 {
	if m != nil {
		return m.FeatureCount
	}
	return 0
}

func (m *RouteSummary) GetDistance() int32 {
	if m != nil {
		return m.Distance
	}
	return 0
}

func (m *RouteSummary) GetElapsedTime() int32 {
	if m != nil {
		return m.ElapsedTime
	}
	return 0
}

func init() {
	proto.RegisterType((*Point)(nil), "stream.Point")
	proto.RegisterType((*Rectangle)(nil), "stream.Rectangle")
	proto.RegisterType((*Feature)(nil), "stream.Feature")
	proto.RegisterType((*RouteNote)(nil), "stream.RouteNote")
	proto.RegisterType((*RouteSummary)(nil), "stream.RouteSummary")
}

func init() {
	proto.RegisterFile("test/service/rpc/proto/stream.proto", fileDescriptor_3a6280eeb8b3790e)
}

var fileDescriptor_3a6280eeb8b3790e = []byte{
	// 384 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xdd, 0xca, 0xd3, 0x40,
	0x10, 0xfd, 0xb6, 0x7e, 0x3f, 0xcd, 0x24, 0x1f, 0xe2, 0xe0, 0x45, 0x08, 0x8a, 0x1a, 0x41, 0xea,
	0x4d, 0x5b, 0x2a, 0xd4, 0x6b, 0x29, 0x58, 0x05, 0x91, 0x12, 0xbd, 0x2f, 0x6b, 0x32, 0xb6, 0x0b,
	0x49, 0xb6, 0x64, 0x27, 0x82, 0xcf, 0xe1, 0x33, 0xfa, 0x1e, 0x92, 0xdd, 0x4d, 0x6c, 0x6d, 0xf1,
	0x2e, 0xe7, 0x9c, 0x39, 0x33, 0x3b, 0x67, 0x02, 0xaf, 0x98, 0x0c, 0xcf, 0x0c, 0x35, 0x3f, 0x54,
	0x4e, 0x33, 0xc3, 0x0d, 0xc9, 0x6a, 0x76, 0x68, 0x34, 0x6b, 0x0f, 0xa6, 0x16, 0xe0, 0xad, 0x43,
	0xe9, 0x3b, 0xb8, 0xd9, 0x68, 0x55, 0x33, 0x26, 0x30, 0x2e, 0x25, 0x2b, 0x6e, 0x0b, 0x8a, 0xc5,
	0x73, 0x31, 0xb9, 0xc9, 0x06, 0x8c, 0x4f, 0x20, 0x28, 0x75, 0xbd, 0x73, 0xe2, 0xc8, 0x8a, 0x7f,
	0x89, 0xf4, 0x23, 0x04, 0x19, 0xe5, 0x2c, 0xeb, 0x5d, 0x49, 0xf8, 0x14, 0x46, 0xa5, 0xb6, 0x0d,
	0xc2, 0xc5, 0xfd, 0xd4, 0x8f, 0xb4, 0x13, 0xb2, 0x51, 0xa9, 0x3b, 0x79, 0xaf, 0x6c, 0x8b, 0x73,
	0x79, 0xaf, 0xd2, 0x0f, 0x70, 0xf7, 0x9e, 0x24, 0xb7, 0x0d, 0x21, 0xc2, 0x75, 0x2d, 0x2b, 0xf7,
	0x96, 0x20, 0xb3, 0xdf, 0xf8, 0x1a, 0xc6, 0xa5, 0xce, 0x25, 0x2b, 0x5d, 0x5f, 0xee, 0x31, 0xc8,
	0xe9, 0x06, 0x82, 0x4c, 0xb7, 0x4c, 0x9f, 0x35, 0x9f, 0xfa, 0xc4, 0x7f, 0x7d, 0x18, 0xc3, 0x5d,
	0x45, 0xc6, 0xc8, 0x9d, 0x5b, 0x34, 0xc8, 0x7a, 0x98, 0xfe, 0x12, 0x10, 0xd9, 0x96, 0x5f, 0xda,
	0xaa, 0x92, 0xcd, 0x4f, 0x7c, 0x06, 0xe1, 0xa1, 0x73, 0x6f, 0x73, 0xdd, 0xd6, 0xec, 0x43, 0x03,
	0x4b, 0xad, 0x3a, 0x06, 0x5f, 0xc2, 0xfd, 0x77, 0xb7, 0x8d, 0x2f, 0x71, 0xd1, 0x45, 0x9e, 0x74,
	0x45, 0x09, 0x8c, 0x0b, 0x65, 0x58, 0xd6, 0x39, 0xc5, 0x0f, 0x5c, 0xee, 0x3d, 0xc6, 0x17, 0x10,
	0x51, 0x29, 0x0f, 0x86, 0x8a, 0x2d, 0xab, 0x8a, 0xe2, 0x6b, 0xab, 0x87, 0x9e, 0xfb, 0xaa, 0x2a,
	0x5a, 0xfc, 0x16, 0x00, 0xf6, 0x55, 0xeb, 0x56, 0x15, 0x84, 0x53, 0x80, 0x35, 0x71, 0x9f, 0xe1,
	0xe9, 0x96, 0xc9, 0xc3, 0x1e, 0x7a, 0x3d, 0xbd, 0xc2, 0x25, 0x44, 0x9f, 0x94, 0xe9, 0x0d, 0x06,
	0x1f, 0xf5, 0x25, 0xc3, 0x45, 0x2f, 0xb8, 0xe6, 0x02, 0x97, 0x10, 0x66, 0x94, 0xeb, 0xa6, 0xb0,
	0xb3, 0xff, 0x1d, 0xf4, 0x78, 0xe8, 0x72, 0x94, 0x57, 0x7a, 0x35, 0x11, 0xf8, 0xd6, 0x9f, 0x65,
	0xb5, 0x97, 0x7c, 0x34, 0xac, 0xbf, 0x54, 0x72, 0x4e, 0x75, 0xb6, 0xb9, 0xf8, 0x76, 0x6b, 0x7f,
	0xdb, 0x37, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x5d, 0x08, 0xd9, 0x0d, 0xe0, 0x02, 0x00, 0x00,
}
