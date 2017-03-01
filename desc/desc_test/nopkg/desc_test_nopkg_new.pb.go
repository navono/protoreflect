// Code generated by protoc-gen-go.
// source: nopkg/desc_test_nopkg_new.proto
// DO NOT EDIT!

package nopkg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type TopLevel struct {
	I                            *int32   `protobuf:"varint,1,opt,name=i" json:"i,omitempty"`
	J                            *int64   `protobuf:"varint,2,opt,name=j" json:"j,omitempty"`
	K                            *int32   `protobuf:"zigzag32,3,opt,name=k" json:"k,omitempty"`
	L                            *int64   `protobuf:"zigzag64,4,opt,name=l" json:"l,omitempty"`
	M                            *uint32  `protobuf:"varint,5,opt,name=m" json:"m,omitempty"`
	N                            *uint64  `protobuf:"varint,6,opt,name=n" json:"n,omitempty"`
	O                            *uint32  `protobuf:"fixed32,7,opt,name=o" json:"o,omitempty"`
	P                            *uint64  `protobuf:"fixed64,8,opt,name=p" json:"p,omitempty"`
	Q                            *int32   `protobuf:"fixed32,9,opt,name=q" json:"q,omitempty"`
	R                            *int64   `protobuf:"fixed64,10,opt,name=r" json:"r,omitempty"`
	S                            *float32 `protobuf:"fixed32,11,opt,name=s" json:"s,omitempty"`
	T                            *float64 `protobuf:"fixed64,12,opt,name=t" json:"t,omitempty"`
	U                            []byte   `protobuf:"bytes,13,opt,name=u" json:"u,omitempty"`
	V                            *string  `protobuf:"bytes,14,opt,name=v" json:"v,omitempty"`
	W                            *bool    `protobuf:"varint,15,opt,name=w" json:"w,omitempty"`
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
}

func (m *TopLevel) Reset()                    { *m = TopLevel{} }
func (m *TopLevel) String() string            { return proto.CompactTextString(m) }
func (*TopLevel) ProtoMessage()               {}
func (*TopLevel) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

var extRange_TopLevel = []proto.ExtensionRange{
	{100, 1000},
}

func (*TopLevel) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_TopLevel
}

func (m *TopLevel) GetI() int32 {
	if m != nil && m.I != nil {
		return *m.I
	}
	return 0
}

func (m *TopLevel) GetJ() int64 {
	if m != nil && m.J != nil {
		return *m.J
	}
	return 0
}

func (m *TopLevel) GetK() int32 {
	if m != nil && m.K != nil {
		return *m.K
	}
	return 0
}

func (m *TopLevel) GetL() int64 {
	if m != nil && m.L != nil {
		return *m.L
	}
	return 0
}

func (m *TopLevel) GetM() uint32 {
	if m != nil && m.M != nil {
		return *m.M
	}
	return 0
}

func (m *TopLevel) GetN() uint64 {
	if m != nil && m.N != nil {
		return *m.N
	}
	return 0
}

func (m *TopLevel) GetO() uint32 {
	if m != nil && m.O != nil {
		return *m.O
	}
	return 0
}

func (m *TopLevel) GetP() uint64 {
	if m != nil && m.P != nil {
		return *m.P
	}
	return 0
}

func (m *TopLevel) GetQ() int32 {
	if m != nil && m.Q != nil {
		return *m.Q
	}
	return 0
}

func (m *TopLevel) GetR() int64 {
	if m != nil && m.R != nil {
		return *m.R
	}
	return 0
}

func (m *TopLevel) GetS() float32 {
	if m != nil && m.S != nil {
		return *m.S
	}
	return 0
}

func (m *TopLevel) GetT() float64 {
	if m != nil && m.T != nil {
		return *m.T
	}
	return 0
}

func (m *TopLevel) GetU() []byte {
	if m != nil {
		return m.U
	}
	return nil
}

func (m *TopLevel) GetV() string {
	if m != nil && m.V != nil {
		return *m.V
	}
	return ""
}

func (m *TopLevel) GetW() bool {
	if m != nil && m.W != nil {
		return *m.W
	}
	return false
}

func init() {
	proto.RegisterType((*TopLevel)(nil), "TopLevel")
}

func init() { proto.RegisterFile("nopkg/desc_test_nopkg_new.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x44, 0xd0, 0xbb, 0x4e, 0xeb, 0x40,
	0x10, 0xc6, 0x71, 0x7d, 0xb9, 0x3a, 0x3e, 0xc9, 0x49, 0x48, 0x35, 0x1d, 0x23, 0xaa, 0x11, 0x45,
	0xdc, 0x22, 0xe8, 0xa8, 0xa9, 0x46, 0x54, 0x34, 0x91, 0x48, 0x96, 0x5c, 0xed, 0xdd, 0xd8, 0x6b,
	0xfb, 0x81, 0xa9, 0x78, 0x0b, 0x34, 0xdb, 0xd0, 0x8c, 0xf4, 0x2b, 0xa6, 0xf8, 0xfe, 0xf9, 0x7d,
	0xe5, 0xc3, 0xe5, 0x50, 0xec, 0x5d, 0xb3, 0xdb, 0x46, 0xd7, 0xc4, 0x6d, 0xf2, 0xb6, 0x72, 0xfd,
	0x26, 0xd4, 0x3e, 0xfa, 0x87, 0x6f, 0xe4, 0xd9, 0xbb, 0x0f, 0x6f, 0xae, 0x73, 0xd7, 0xf5, 0x3c,
	0xc7, 0x89, 0xc0, 0x90, 0xb1, 0xe2, 0x64, 0x3a, 0xd3, 0x80, 0x21, 0x43, 0xc5, 0xd9, 0x74, 0xa1,
	0x21, 0x43, 0xee, 0x14, 0x17, 0xd3, 0x95, 0x46, 0x0c, 0x59, 0x2b, 0xd2, 0x5f, 0x49, 0x63, 0x86,
	0x2c, 0x14, 0xa5, 0xa9, 0xa2, 0x09, 0x43, 0x46, 0x8a, 0xca, 0xe4, 0x69, 0xca, 0x90, 0xa9, 0xc2,
	0x9b, 0x02, 0x65, 0x0c, 0x99, 0x28, 0x82, 0xe9, 0x46, 0x33, 0x86, 0x2c, 0x15, 0x37, 0x53, 0x4d,
	0x39, 0x43, 0x56, 0x8a, 0xda, 0xd4, 0xd0, 0x3f, 0x86, 0x0c, 0x14, 0x8d, 0x29, 0xd2, 0x9c, 0x21,
	0x50, 0x44, 0x53, 0x4b, 0x0b, 0x86, 0xcc, 0x15, 0xad, 0xa9, 0xa3, 0xff, 0x0c, 0x99, 0x29, 0x3a,
	0x53, 0x4f, 0x4b, 0x86, 0x64, 0x8a, 0xfe, 0x71, 0x9c, 0xed, 0x57, 0x3f, 0xd3, 0xd7, 0xe7, 0x8f,
	0xa7, 0xc3, 0x29, 0x1e, 0xdb, 0xcf, 0xcd, 0xce, 0x97, 0xc5, 0xf9, 0xd8, 0x96, 0xa1, 0x48, 0xfb,
	0x6b, 0xf7, 0x75, 0x75, 0xbb, 0x98, 0x02, 0xfd, 0x55, 0x2a, 0x52, 0xa5, 0x97, 0x74, 0x7f, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x22, 0x24, 0x10, 0x87, 0x43, 0x01, 0x00, 0x00,
}
