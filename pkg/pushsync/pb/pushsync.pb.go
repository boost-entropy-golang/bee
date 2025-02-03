// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pushsync.proto

package pb

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Delivery struct {
	Address []byte `protobuf:"bytes,1,opt,name=Address,proto3" json:"Address,omitempty"`
	Data    []byte `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
	Stamp   []byte `protobuf:"bytes,3,opt,name=Stamp,proto3" json:"Stamp,omitempty"`
}

func (m *Delivery) Reset()         { *m = Delivery{} }
func (m *Delivery) String() string { return proto.CompactTextString(m) }
func (*Delivery) ProtoMessage()    {}
func (*Delivery) Descriptor() ([]byte, []int) {
	return fileDescriptor_723cf31bfc02bfd6, []int{0}
}
func (m *Delivery) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Delivery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Delivery.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Delivery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Delivery.Merge(m, src)
}
func (m *Delivery) XXX_Size() int {
	return m.Size()
}
func (m *Delivery) XXX_DiscardUnknown() {
	xxx_messageInfo_Delivery.DiscardUnknown(m)
}

var xxx_messageInfo_Delivery proto.InternalMessageInfo

func (m *Delivery) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Delivery) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Delivery) GetStamp() []byte {
	if m != nil {
		return m.Stamp
	}
	return nil
}

type Receipt struct {
	Address       []byte `protobuf:"bytes,1,opt,name=Address,proto3" json:"Address,omitempty"`
	Signature     []byte `protobuf:"bytes,2,opt,name=Signature,proto3" json:"Signature,omitempty"`
	Nonce         []byte `protobuf:"bytes,3,opt,name=Nonce,proto3" json:"Nonce,omitempty"`
	Err           string `protobuf:"bytes,4,opt,name=Err,proto3" json:"Err,omitempty"`
	StorageRadius uint32 `protobuf:"varint,5,opt,name=StorageRadius,proto3" json:"StorageRadius,omitempty"`
}

func (m *Receipt) Reset()         { *m = Receipt{} }
func (m *Receipt) String() string { return proto.CompactTextString(m) }
func (*Receipt) ProtoMessage()    {}
func (*Receipt) Descriptor() ([]byte, []int) {
	return fileDescriptor_723cf31bfc02bfd6, []int{1}
}
func (m *Receipt) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Receipt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Receipt.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Receipt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Receipt.Merge(m, src)
}
func (m *Receipt) XXX_Size() int {
	return m.Size()
}
func (m *Receipt) XXX_DiscardUnknown() {
	xxx_messageInfo_Receipt.DiscardUnknown(m)
}

var xxx_messageInfo_Receipt proto.InternalMessageInfo

func (m *Receipt) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Receipt) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *Receipt) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

func (m *Receipt) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func (m *Receipt) GetStorageRadius() uint32 {
	if m != nil {
		return m.StorageRadius
	}
	return 0
}

func init() {
	proto.RegisterType((*Delivery)(nil), "pushsync.Delivery")
	proto.RegisterType((*Receipt)(nil), "pushsync.Receipt")
}

func init() { proto.RegisterFile("pushsync.proto", fileDescriptor_723cf31bfc02bfd6) }

var fileDescriptor_723cf31bfc02bfd6 = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x28, 0x2d, 0xce,
	0x28, 0xae, 0xcc, 0x4b, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0x95, 0xfc,
	0xb8, 0x38, 0x5c, 0x52, 0x73, 0x32, 0xcb, 0x52, 0x8b, 0x2a, 0x85, 0x24, 0xb8, 0xd8, 0x1d, 0x53,
	0x52, 0x8a, 0x52, 0x8b, 0x8b, 0x25, 0x18, 0x15, 0x18, 0x35, 0x78, 0x82, 0x60, 0x5c, 0x21, 0x21,
	0x2e, 0x16, 0x97, 0xc4, 0x92, 0x44, 0x09, 0x26, 0xb0, 0x30, 0x98, 0x2d, 0x24, 0xc2, 0xc5, 0x1a,
	0x5c, 0x92, 0x98, 0x5b, 0x20, 0xc1, 0x0c, 0x16, 0x84, 0x70, 0x94, 0xfa, 0x19, 0xb9, 0xd8, 0x83,
	0x52, 0x93, 0x53, 0x33, 0x0b, 0x4a, 0xf0, 0x98, 0x27, 0xc3, 0xc5, 0x19, 0x9c, 0x99, 0x9e, 0x97,
	0x58, 0x52, 0x5a, 0x94, 0x0a, 0x35, 0x14, 0x21, 0x00, 0x32, 0xd9, 0x2f, 0x3f, 0x2f, 0x39, 0x15,
	0x66, 0x32, 0x98, 0x23, 0x24, 0xc0, 0xc5, 0xec, 0x5a, 0x54, 0x24, 0xc1, 0xa2, 0xc0, 0xa8, 0xc1,
	0x19, 0x04, 0x62, 0x0a, 0xa9, 0x70, 0xf1, 0x06, 0x97, 0xe4, 0x17, 0x25, 0xa6, 0xa7, 0x06, 0x25,
	0xa6, 0x64, 0x96, 0x16, 0x4b, 0xb0, 0x2a, 0x30, 0x6a, 0xf0, 0x06, 0xa1, 0x0a, 0x3a, 0xc9, 0x9c,
	0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31,
	0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x14, 0x53, 0x41, 0x52, 0x12, 0x1b, 0x38,
	0x40, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x80, 0xc6, 0x9e, 0x01, 0x22, 0x01, 0x00, 0x00,
}

func (m *Delivery) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Delivery) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Delivery) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Stamp) > 0 {
		i -= len(m.Stamp)
		copy(dAtA[i:], m.Stamp)
		i = encodeVarintPushsync(dAtA, i, uint64(len(m.Stamp)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintPushsync(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintPushsync(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Receipt) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Receipt) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Receipt) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.StorageRadius != 0 {
		i = encodeVarintPushsync(dAtA, i, uint64(m.StorageRadius))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Err) > 0 {
		i -= len(m.Err)
		copy(dAtA[i:], m.Err)
		i = encodeVarintPushsync(dAtA, i, uint64(len(m.Err)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Nonce) > 0 {
		i -= len(m.Nonce)
		copy(dAtA[i:], m.Nonce)
		i = encodeVarintPushsync(dAtA, i, uint64(len(m.Nonce)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintPushsync(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintPushsync(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPushsync(dAtA []byte, offset int, v uint64) int {
	offset -= sovPushsync(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Delivery) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovPushsync(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovPushsync(uint64(l))
	}
	l = len(m.Stamp)
	if l > 0 {
		n += 1 + l + sovPushsync(uint64(l))
	}
	return n
}

func (m *Receipt) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovPushsync(uint64(l))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovPushsync(uint64(l))
	}
	l = len(m.Nonce)
	if l > 0 {
		n += 1 + l + sovPushsync(uint64(l))
	}
	l = len(m.Err)
	if l > 0 {
		n += 1 + l + sovPushsync(uint64(l))
	}
	if m.StorageRadius != 0 {
		n += 1 + sovPushsync(uint64(m.StorageRadius))
	}
	return n
}

func sovPushsync(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPushsync(x uint64) (n int) {
	return sovPushsync(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Delivery) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPushsync
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Delivery: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Delivery: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPushsync
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPushsync
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPushsync
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPushsync
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPushsync
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPushsync
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stamp", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPushsync
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPushsync
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPushsync
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Stamp = append(m.Stamp[:0], dAtA[iNdEx:postIndex]...)
			if m.Stamp == nil {
				m.Stamp = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPushsync(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPushsync
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPushsync
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Receipt) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPushsync
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Receipt: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Receipt: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPushsync
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPushsync
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPushsync
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPushsync
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPushsync
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPushsync
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = append(m.Signature[:0], dAtA[iNdEx:postIndex]...)
			if m.Signature == nil {
				m.Signature = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPushsync
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPushsync
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPushsync
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nonce = append(m.Nonce[:0], dAtA[iNdEx:postIndex]...)
			if m.Nonce == nil {
				m.Nonce = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Err", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPushsync
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPushsync
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPushsync
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Err = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StorageRadius", wireType)
			}
			m.StorageRadius = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPushsync
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StorageRadius |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPushsync(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPushsync
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPushsync
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipPushsync(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPushsync
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPushsync
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPushsync
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthPushsync
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPushsync
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPushsync
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPushsync        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPushsync          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPushsync = fmt.Errorf("proto: unexpected end of group")
)
