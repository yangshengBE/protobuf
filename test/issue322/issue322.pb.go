// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: issue322.proto

package test

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/yangshengBE/protobuf/gogoproto"
	proto "github.com/yangshengBE/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

type OneofTest struct {
	// Types that are valid to be assigned to Union:
	//	*OneofTest_I
	Union                isOneofTest_Union `protobuf_oneof:"union"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *OneofTest) Reset()         { *m = OneofTest{} }
func (m *OneofTest) String() string { return proto.CompactTextString(m) }
func (*OneofTest) ProtoMessage()    {}
func (*OneofTest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fbf4e8d164dccde1, []int{0}
}
func (m *OneofTest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OneofTest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OneofTest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OneofTest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OneofTest.Merge(m, src)
}
func (m *OneofTest) XXX_Size() int {
	return m.Size()
}
func (m *OneofTest) XXX_DiscardUnknown() {
	xxx_messageInfo_OneofTest.DiscardUnknown(m)
}

var xxx_messageInfo_OneofTest proto.InternalMessageInfo

const Default_OneofTest_I int32 = 4

type isOneofTest_Union interface {
	isOneofTest_Union()
	Equal(interface{}) bool
	MarshalTo([]byte) (int, error)
	Size() int
	Compare(interface{}) int
}

type OneofTest_I struct {
	I int32 `protobuf:"varint,1,opt,name=i,oneof,def=4" json:"i,omitempty"`
}

func (*OneofTest_I) isOneofTest_Union() {}

func (m *OneofTest) GetUnion() isOneofTest_Union {
	if m != nil {
		return m.Union
	}
	return nil
}

func (m *OneofTest) GetI() int32 {
	if x, ok := m.GetUnion().(*OneofTest_I); ok {
		return x.I
	}
	return Default_OneofTest_I
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*OneofTest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*OneofTest_I)(nil),
	}
}

func init() {
	proto.RegisterType((*OneofTest)(nil), "test.OneofTest")
}

func init() { proto.RegisterFile("issue322.proto", fileDescriptor_fbf4e8d164dccde1) }

var fileDescriptor_fbf4e8d164dccde1 = []byte{
	// 153 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0x2c, 0x2e, 0x2e,
	0x4d, 0x35, 0x36, 0x32, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x29, 0x49, 0x2d, 0x2e,
	0x91, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0x4f,
	0xcf, 0xd7, 0x07, 0x4b, 0x26, 0x95, 0xa6, 0x81, 0x79, 0x60, 0x0e, 0x98, 0x05, 0xd1, 0xa4, 0xa4,
	0xce, 0xc5, 0xe9, 0x9f, 0x97, 0x9a, 0x9f, 0x16, 0x92, 0x5a, 0x5c, 0x22, 0x24, 0xc8, 0xc5, 0x98,
	0x29, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x6a, 0xc5, 0x68, 0xe2, 0xc1, 0x10, 0xc4, 0x98, 0xe9, 0xc4,
	0xce, 0xc5, 0x5a, 0x9a, 0x97, 0x99, 0x9f, 0xe7, 0xa4, 0xf0, 0xe1, 0xa1, 0x1c, 0xe3, 0x8f, 0x87,
	0x72, 0x8c, 0x2b, 0x1e, 0xc9, 0x31, 0xee, 0x78, 0x24, 0xc7, 0x78, 0xe2, 0x91, 0x1c, 0xe3, 0x85,
	0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0xae, 0x78, 0x2c, 0xc7, 0x08, 0x08, 0x00, 0x00, 0xff,
	0xff, 0x42, 0xe9, 0xae, 0x78, 0x90, 0x00, 0x00, 0x00,
}

func (this *OneofTest) Compare(that interface{}) int {
	if that == nil {
		if this == nil {
			return 0
		}
		return 1
	}

	that1, ok := that.(*OneofTest)
	if !ok {
		that2, ok := that.(OneofTest)
		if ok {
			that1 = &that2
		} else {
			return 1
		}
	}
	if that1 == nil {
		if this == nil {
			return 0
		}
		return 1
	} else if this == nil {
		return -1
	}
	if that1.Union == nil {
		if this.Union != nil {
			return 1
		}
	} else if this.Union == nil {
		return -1
	} else {
		thisType := -1
		switch this.Union.(type) {
		case *OneofTest_I:
			thisType = 0
		default:
			panic(fmt.Sprintf("compare: unexpected type %T in oneof", this.Union))
		}
		that1Type := -1
		switch that1.Union.(type) {
		case *OneofTest_I:
			that1Type = 0
		default:
			panic(fmt.Sprintf("compare: unexpected type %T in oneof", that1.Union))
		}
		if thisType == that1Type {
			if c := this.Union.Compare(that1.Union); c != 0 {
				return c
			}
		} else if thisType < that1Type {
			return -1
		} else if thisType > that1Type {
			return 1
		}
	}
	if c := bytes.Compare(this.XXX_unrecognized, that1.XXX_unrecognized); c != 0 {
		return c
	}
	return 0
}
func (this *OneofTest_I) Compare(that interface{}) int {
	if that == nil {
		if this == nil {
			return 0
		}
		return 1
	}

	that1, ok := that.(*OneofTest_I)
	if !ok {
		that2, ok := that.(OneofTest_I)
		if ok {
			that1 = &that2
		} else {
			return 1
		}
	}
	if that1 == nil {
		if this == nil {
			return 0
		}
		return 1
	} else if this == nil {
		return -1
	}
	if this.I != that1.I {
		if this.I < that1.I {
			return -1
		}
		return 1
	}
	return 0
}
func (this *OneofTest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*OneofTest)
	if !ok {
		that2, ok := that.(OneofTest)
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
	if that1.Union == nil {
		if this.Union != nil {
			return false
		}
	} else if this.Union == nil {
		return false
	} else if !this.Union.Equal(that1.Union) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *OneofTest_I) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*OneofTest_I)
	if !ok {
		that2, ok := that.(OneofTest_I)
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
	if this.I != that1.I {
		return false
	}
	return true
}
func (this *OneofTest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&test.OneofTest{")
	if this.Union != nil {
		s = append(s, "Union: "+fmt.Sprintf("%#v", this.Union)+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *OneofTest_I) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&test.OneofTest_I{` +
		`I:` + fmt.Sprintf("%#v", this.I) + `}`}, ", ")
	return s
}
func valueToGoStringIssue322(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *OneofTest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OneofTest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OneofTest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Union != nil {
		{
			size := m.Union.Size()
			i -= size
			if _, err := m.Union.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *OneofTest_I) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OneofTest_I) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i = encodeVarintIssue322(dAtA, i, uint64(m.I))
	i--
	dAtA[i] = 0x8
	return len(dAtA) - i, nil
}
func encodeVarintIssue322(dAtA []byte, offset int, v uint64) int {
	offset -= sovIssue322(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedOneofTest(r randyIssue322, easy bool) *OneofTest {
	this := &OneofTest{}
	oneofNumber_Union := []int32{1}[r.Intn(1)]
	switch oneofNumber_Union {
	case 1:
		this.Union = NewPopulatedOneofTest_I(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedIssue322(r, 2)
	}
	return this
}

func NewPopulatedOneofTest_I(r randyIssue322, easy bool) *OneofTest_I {
	this := &OneofTest_I{}
	this.I = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.I *= -1
	}
	return this
}

type randyIssue322 interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneIssue322(r randyIssue322) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringIssue322(r randyIssue322) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneIssue322(r)
	}
	return string(tmps)
}
func randUnrecognizedIssue322(r randyIssue322, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldIssue322(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldIssue322(dAtA []byte, r randyIssue322, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateIssue322(dAtA, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		dAtA = encodeVarintPopulateIssue322(dAtA, uint64(v2))
	case 1:
		dAtA = encodeVarintPopulateIssue322(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateIssue322(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateIssue322(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateIssue322(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateIssue322(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *OneofTest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Union != nil {
		n += m.Union.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *OneofTest_I) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovIssue322(uint64(m.I))
	return n
}

func sovIssue322(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIssue322(x uint64) (n int) {
	return sovIssue322(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OneofTest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIssue322
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
			return fmt.Errorf("proto: OneofTest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OneofTest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field I", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIssue322
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Union = &OneofTest_I{v}
		default:
			iNdEx = preIndex
			skippy, err := skipIssue322(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIssue322
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
func skipIssue322(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIssue322
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
					return 0, ErrIntOverflowIssue322
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
					return 0, ErrIntOverflowIssue322
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
				return 0, ErrInvalidLengthIssue322
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIssue322
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIssue322
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIssue322        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIssue322          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIssue322 = fmt.Errorf("proto: unexpected end of group")
)
