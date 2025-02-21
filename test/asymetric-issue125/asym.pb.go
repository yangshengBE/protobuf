// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: asym.proto

package asym

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/yangshengBE/protobuf/gogoproto"
	proto "github.com/yangshengBE/protobuf/proto"
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

type M struct {
	Arr                  []MyType `protobuf:"bytes,1,rep,name=arr,customtype=MyType" json:"arr"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *M) Reset()         { *m = M{} }
func (m *M) String() string { return proto.CompactTextString(m) }
func (*M) ProtoMessage()    {}
func (*M) Descriptor() ([]byte, []int) {
	return fileDescriptor_72d286349de177ec, []int{0}
}
func (m *M) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *M) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_M.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *M) XXX_Merge(src proto.Message) {
	xxx_messageInfo_M.Merge(m, src)
}
func (m *M) XXX_Size() int {
	return m.Size()
}
func (m *M) XXX_DiscardUnknown() {
	xxx_messageInfo_M.DiscardUnknown(m)
}

var xxx_messageInfo_M proto.InternalMessageInfo

type MyType struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MyType) Reset()         { *m = MyType{} }
func (m *MyType) String() string { return proto.CompactTextString(m) }
func (*MyType) ProtoMessage()    {}
func (*MyType) Descriptor() ([]byte, []int) {
	return fileDescriptor_72d286349de177ec, []int{1}
}
func (m *MyType) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MyType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MyType.Marshal(b, m, deterministic)
}
func (m *MyType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MyType.Merge(m, src)
}
func (m *MyType) XXX_Size() int {
	return xxx_messageInfo_MyType.Size(m)
}
func (m *MyType) XXX_DiscardUnknown() {
	xxx_messageInfo_MyType.DiscardUnknown(m)
}

var xxx_messageInfo_MyType proto.InternalMessageInfo

func init() {
	proto.RegisterType((*M)(nil), "asym.M")
	proto.RegisterType((*MyType)(nil), "asym.MyType")
}

func init() { proto.RegisterFile("asym.proto", fileDescriptor_72d286349de177ec) }

var fileDescriptor_72d286349de177ec = []byte{
	// 158 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x2c, 0xae, 0xcc,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0xa5, 0x74, 0xd3, 0x33, 0x4b, 0x32,
	0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xd3, 0xf3, 0xd3, 0xf3, 0xf5, 0xc1, 0x92, 0x49, 0xa5,
	0x69, 0x60, 0x1e, 0x98, 0x03, 0x66, 0x41, 0x34, 0x29, 0xa9, 0x72, 0x31, 0xfa, 0x0a, 0x29, 0x70,
	0x31, 0x27, 0x16, 0x15, 0x49, 0x30, 0x2a, 0x30, 0x6b, 0xf0, 0x38, 0xf1, 0x9d, 0xb8, 0x27, 0xcf,
	0x70, 0xeb, 0x9e, 0x3c, 0x9b, 0x6f, 0x65, 0x48, 0x65, 0x41, 0x6a, 0x10, 0x48, 0x4a, 0x49, 0x8a,
	0x0b, 0xca, 0xb5, 0x12, 0xd8, 0xb1, 0x40, 0x9e, 0xe1, 0xc7, 0x02, 0x79, 0x86, 0x8e, 0x85, 0xf2,
	0x0c, 0x0b, 0x16, 0xca, 0x33, 0x38, 0x49, 0x3c, 0x78, 0x28, 0xc7, 0xf8, 0xe3, 0xa1, 0x1c, 0xe3,
	0x8a, 0x47, 0x72, 0x8c, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c,
	0x23, 0x20, 0x00, 0x00, 0xff, 0xff, 0x81, 0x2c, 0x72, 0xc1, 0x9e, 0x00, 0x00, 0x00,
}

func (this *M) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*M)
	if !ok {
		that2, ok := that.(M)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *M")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *M but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *M but is not nil && this == nil")
	}
	if len(this.Arr) != len(that1.Arr) {
		return fmt.Errorf("Arr this(%v) Not Equal that(%v)", len(this.Arr), len(that1.Arr))
	}
	for i := range this.Arr {
		if !this.Arr[i].Equal(that1.Arr[i]) {
			return fmt.Errorf("Arr this[%v](%v) Not Equal that[%v](%v)", i, this.Arr[i], i, that1.Arr[i])
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *M) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*M)
	if !ok {
		that2, ok := that.(M)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Arr) != len(that1.Arr) {
		return false
	}
	for i := range this.Arr {
		if !this.Arr[i].Equal(that1.Arr[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MyType) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*MyType)
	if !ok {
		that2, ok := that.(MyType)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *MyType")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *MyType but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *MyType but is not nil && this == nil")
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *MyType) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MyType)
	if !ok {
		that2, ok := that.(MyType)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (m *M) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *M) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *M) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Arr) > 0 {
		for iNdEx := len(m.Arr) - 1; iNdEx >= 0; iNdEx-- {
			{
				size := m.Arr[iNdEx].Size()
				i -= size
				if _, err := m.Arr[iNdEx].MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
				i = encodeVarintAsym(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintAsym(dAtA []byte, offset int, v uint64) int {
	offset -= sovAsym(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedM(r randyAsym, easy bool) *M {
	this := &M{}
	if r.Intn(5) != 0 {
		v1 := r.Intn(10)
		this.Arr = make([]MyType, v1)
		for i := 0; i < v1; i++ {
			v2 := NewPopulatedMyType(r)
			this.Arr[i] = *v2
		}
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedAsym(r, 2)
	}
	return this
}

type randyAsym interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneAsym(r randyAsym) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringAsym(r randyAsym) string {
	v3 := r.Intn(100)
	tmps := make([]rune, v3)
	for i := 0; i < v3; i++ {
		tmps[i] = randUTF8RuneAsym(r)
	}
	return string(tmps)
}
func randUnrecognizedAsym(r randyAsym, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldAsym(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldAsym(dAtA []byte, r randyAsym, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateAsym(dAtA, uint64(key))
		v4 := r.Int63()
		if r.Intn(2) == 0 {
			v4 *= -1
		}
		dAtA = encodeVarintPopulateAsym(dAtA, uint64(v4))
	case 1:
		dAtA = encodeVarintPopulateAsym(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateAsym(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateAsym(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateAsym(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateAsym(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *M) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Arr) > 0 {
		for _, e := range m.Arr {
			l = e.Size()
			n += 1 + l + sovAsym(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovAsym(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAsym(x uint64) (n int) {
	return sovAsym(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *M) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAsym
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
			return fmt.Errorf("proto: M: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: M: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Arr", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsym
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
				return ErrInvalidLengthAsym
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAsym
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v MyType
			m.Arr = append(m.Arr, v)
			if err := m.Arr[len(m.Arr)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAsym(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAsym
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MyType) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAsym
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
			return fmt.Errorf("proto: MyType: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MyType: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipAsym(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAsym
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipAsym(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAsym
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
					return 0, ErrIntOverflowAsym
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
					return 0, ErrIntOverflowAsym
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
				return 0, ErrInvalidLengthAsym
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAsym
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAsym
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAsym        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAsym          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAsym = fmt.Errorf("proto: unexpected end of group")
)
