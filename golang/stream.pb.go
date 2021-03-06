// Code generated by protoc-gen-go. DO NOT EDIT.
// source: stream.proto

package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// StreamRequest represents the request to the stream service via GRPC
type StreamRequest struct {
	// The stream configuration for scanning streams such as Azure EventHub and Kafka
	StreamConfig *StreamConfig `protobuf:"bytes,1,opt,name=streamConfig,proto3" json:"streamConfig,omitempty"`
	// The analyzer template configures the fields to analyze
	AnalyzeTemplate *AnalyzeTemplate `protobuf:"bytes,2,opt,name=analyzeTemplate,proto3" json:"analyzeTemplate,omitempty"`
	// The anonymizer template configures how to anonymize the sensitive data [optional]
	AnonymizeTemplate *AnonymizeTemplate `protobuf:"bytes,3,opt,name=anonymizeTemplate,proto3" json:"anonymizeTemplate,omitempty"`
	// The datasinkTemplate configures the output destination of the analyzer/anonymizer results
	DatasinkTemplate     *DatasinkTemplate `protobuf:"bytes,4,opt,name=datasinkTemplate,proto3" json:"datasinkTemplate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *StreamRequest) Reset()         { *m = StreamRequest{} }
func (m *StreamRequest) String() string { return proto.CompactTextString(m) }
func (*StreamRequest) ProtoMessage()    {}
func (*StreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_stream_ee78de99a3818fe5, []int{0}
}
func (m *StreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamRequest.Unmarshal(m, b)
}
func (m *StreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamRequest.Marshal(b, m, deterministic)
}
func (dst *StreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamRequest.Merge(dst, src)
}
func (m *StreamRequest) XXX_Size() int {
	return xxx_messageInfo_StreamRequest.Size(m)
}
func (m *StreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamRequest proto.InternalMessageInfo

func (m *StreamRequest) GetStreamConfig() *StreamConfig {
	if m != nil {
		return m.StreamConfig
	}
	return nil
}

func (m *StreamRequest) GetAnalyzeTemplate() *AnalyzeTemplate {
	if m != nil {
		return m.AnalyzeTemplate
	}
	return nil
}

func (m *StreamRequest) GetAnonymizeTemplate() *AnonymizeTemplate {
	if m != nil {
		return m.AnonymizeTemplate
	}
	return nil
}

func (m *StreamRequest) GetDatasinkTemplate() *DatasinkTemplate {
	if m != nil {
		return m.DatasinkTemplate
	}
	return nil
}

func init() {
	proto.RegisterType((*StreamRequest)(nil), "types.StreamRequest")
}

func init() { proto.RegisterFile("stream.proto", fileDescriptor_stream_ee78de99a3818fe5) }

var fileDescriptor_stream_ee78de99a3818fe5 = []byte{
	// 184 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x2e, 0x29, 0x4a,
	0x4d, 0xcc, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2d, 0xa9, 0x2c, 0x48, 0x2d, 0x96,
	0xe2, 0x2b, 0x49, 0xcd, 0x2d, 0xc8, 0x49, 0x2c, 0x49, 0x85, 0x08, 0x2b, 0x4d, 0x67, 0xe2, 0xe2,
	0x0d, 0x06, 0xab, 0x0b, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x32, 0x87, 0x69, 0x74, 0xce,
	0xcf, 0x4b, 0xcb, 0x4c, 0x97, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x36, 0x12, 0xd6, 0x03, 0xeb, 0xd7,
	0x0b, 0x46, 0x92, 0x0a, 0x42, 0x51, 0x28, 0xe4, 0xc0, 0xc5, 0x9f, 0x98, 0x97, 0x98, 0x53, 0x59,
	0x95, 0x1a, 0x02, 0xb5, 0x43, 0x82, 0x09, 0xac, 0x57, 0x0c, 0xaa, 0xd7, 0x11, 0x55, 0x36, 0x08,
	0x5d, 0xb9, 0x90, 0x1b, 0x97, 0x60, 0x62, 0x5e, 0x7e, 0x5e, 0x65, 0x6e, 0x26, 0x92, 0x19, 0xcc,
	0x60, 0x33, 0x24, 0xe0, 0x66, 0xa0, 0xc9, 0x07, 0x61, 0x6a, 0x11, 0x72, 0xe6, 0x12, 0x48, 0x49,
	0x2c, 0x49, 0x2c, 0xce, 0xcc, 0xcb, 0x86, 0x1b, 0xc3, 0x02, 0x36, 0x46, 0x1c, 0x6a, 0x8c, 0x0b,
	0x9a, 0x74, 0x10, 0x86, 0x86, 0x24, 0x36, 0x70, 0x00, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x94, 0xb7, 0x27, 0xa6, 0x47, 0x01, 0x00, 0x00,
}
