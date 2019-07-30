// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: src/carnotpb/carnot.proto

package carnotpb

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	proto1 "pixielabs.ai/pixielabs/src/common/uuid/proto"
	proto2 "pixielabs.ai/pixielabs/src/table_store/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type QueuedRowBatch struct {
	Address       string               `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	QueryID       *proto1.UUID         `protobuf:"bytes,2,opt,name=query_id,json=queryId,proto3" json:"query_id,omitempty"`
	DestinationId int64                `protobuf:"varint,3,opt,name=destination_id,json=destinationId,proto3" json:"destination_id,omitempty"`
	RowBatch      *proto2.RowBatchData `protobuf:"bytes,4,opt,name=row_batch,json=rowBatch,proto3" json:"row_batch,omitempty"`
}

func (m *QueuedRowBatch) Reset()      { *m = QueuedRowBatch{} }
func (*QueuedRowBatch) ProtoMessage() {}
func (*QueuedRowBatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_9054907422204f4e, []int{0}
}
func (m *QueuedRowBatch) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueuedRowBatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueuedRowBatch.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueuedRowBatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueuedRowBatch.Merge(m, src)
}
func (m *QueuedRowBatch) XXX_Size() int {
	return m.Size()
}
func (m *QueuedRowBatch) XXX_DiscardUnknown() {
	xxx_messageInfo_QueuedRowBatch.DiscardUnknown(m)
}

var xxx_messageInfo_QueuedRowBatch proto.InternalMessageInfo

func (m *QueuedRowBatch) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *QueuedRowBatch) GetQueryID() *proto1.UUID {
	if m != nil {
		return m.QueryID
	}
	return nil
}

func (m *QueuedRowBatch) GetDestinationId() int64 {
	if m != nil {
		return m.DestinationId
	}
	return 0
}

func (m *QueuedRowBatch) GetRowBatch() *proto2.RowBatchData {
	if m != nil {
		return m.RowBatch
	}
	return nil
}

func init() {
	proto.RegisterType((*QueuedRowBatch)(nil), "pl.carnotpb.QueuedRowBatch")
}

func init() { proto.RegisterFile("src/carnotpb/carnot.proto", fileDescriptor_9054907422204f4e) }

var fileDescriptor_9054907422204f4e = []byte{
	// 338 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0xb1, 0x6e, 0xea, 0x30,
	0x18, 0x85, 0xe3, 0xcb, 0x55, 0x81, 0xa0, 0x52, 0x29, 0x53, 0xca, 0xf0, 0x37, 0x42, 0x42, 0x62,
	0xa9, 0x23, 0xb5, 0x43, 0xb7, 0x0e, 0x11, 0x4b, 0x46, 0x22, 0xb1, 0x74, 0x41, 0x76, 0xec, 0x42,
	0x24, 0x12, 0xa7, 0x8e, 0x23, 0xd4, 0xad, 0x8f, 0xd0, 0xc7, 0xe8, 0xa3, 0x74, 0xaa, 0x18, 0x99,
	0xaa, 0x62, 0x96, 0x8e, 0x3c, 0x42, 0x15, 0x27, 0x91, 0xd8, 0xce, 0xb1, 0xbf, 0xdf, 0xe7, 0xd8,
	0xb6, 0xaf, 0x0b, 0x19, 0xfb, 0x31, 0x91, 0x99, 0x50, 0x39, 0x6d, 0x04, 0xce, 0xa5, 0x50, 0xc2,
	0x19, 0xe4, 0x1b, 0xdc, 0xee, 0x8c, 0x6e, 0x57, 0x89, 0x5a, 0x97, 0x14, 0xc7, 0x22, 0xf5, 0x57,
	0x62, 0x25, 0x7c, 0xc3, 0xd0, 0xf2, 0xd9, 0x38, 0x63, 0x8c, 0xaa, 0x67, 0x47, 0x9e, 0x39, 0x56,
	0xa4, 0xa9, 0xc8, 0xfc, 0xb2, 0x4c, 0x58, 0x8d, 0x1b, 0xd9, 0x10, 0xe3, 0x8a, 0x50, 0x84, 0x6e,
	0xf8, 0xb2, 0x50, 0x42, 0xf2, 0x86, 0x28, 0xe2, 0x35, 0x4f, 0x49, 0xcd, 0x8c, 0xbf, 0x90, 0x3d,
	0x9c, 0x97, 0xbc, 0xe4, 0x2c, 0x12, 0xdb, 0x80, 0xa8, 0x78, 0xed, 0xb8, 0x76, 0x97, 0x30, 0x26,
	0x79, 0x51, 0xb8, 0xc8, 0x43, 0xd3, 0x7e, 0xd4, 0x5a, 0xe7, 0xc1, 0xee, 0xbd, 0x94, 0x5c, 0xbe,
	0x2e, 0x13, 0xe6, 0xfe, 0xf3, 0xd0, 0x74, 0x70, 0x77, 0x85, 0xf3, 0x0d, 0xae, 0x22, 0x73, 0x8a,
	0x17, 0x8b, 0x70, 0x16, 0x0c, 0xf4, 0xf7, 0x4d, 0x77, 0x5e, 0x41, 0xe1, 0x2c, 0xea, 0x1a, 0x3a,
	0x64, 0xce, 0xc4, 0x1e, 0x32, 0x5e, 0xa8, 0x24, 0x23, 0x2a, 0x11, 0x59, 0x35, 0xde, 0xf1, 0xd0,
	0xb4, 0x13, 0x5d, 0x9e, 0xad, 0x86, 0xcc, 0x09, 0xec, 0xbe, 0x14, 0xdb, 0x25, 0xad, 0x6a, 0xb8,
	0xff, 0x4d, 0xc0, 0xa4, 0x0a, 0x38, 0xbb, 0x03, 0xae, 0xdb, 0xe7, 0x14, 0xb7, 0x7d, 0x67, 0x44,
	0x91, 0xa8, 0x27, 0x1b, 0x17, 0x3c, 0xee, 0x0e, 0x60, 0xed, 0x0f, 0x60, 0x9d, 0x0e, 0x80, 0xde,
	0x34, 0xa0, 0x0f, 0x0d, 0xe8, 0x53, 0x03, 0xda, 0x69, 0x40, 0x3f, 0x1a, 0xd0, 0xaf, 0x06, 0xeb,
	0xa4, 0x01, 0xbd, 0x1f, 0xc1, 0xda, 0x1d, 0xc1, 0xda, 0x1f, 0xc1, 0x7a, 0xea, 0xb5, 0xbf, 0x40,
	0x2f, 0xcc, 0xbb, 0xdc, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x1d, 0xa5, 0x5f, 0x5f, 0xb6, 0x01,
	0x00, 0x00,
}

func (this *QueuedRowBatch) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*QueuedRowBatch)
	if !ok {
		that2, ok := that.(QueuedRowBatch)
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
	if this.Address != that1.Address {
		return false
	}
	if !this.QueryID.Equal(that1.QueryID) {
		return false
	}
	if this.DestinationId != that1.DestinationId {
		return false
	}
	if !this.RowBatch.Equal(that1.RowBatch) {
		return false
	}
	return true
}
func (this *QueuedRowBatch) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&carnotpb.QueuedRowBatch{")
	s = append(s, "Address: "+fmt.Sprintf("%#v", this.Address)+",\n")
	if this.QueryID != nil {
		s = append(s, "QueryID: "+fmt.Sprintf("%#v", this.QueryID)+",\n")
	}
	s = append(s, "DestinationId: "+fmt.Sprintf("%#v", this.DestinationId)+",\n")
	if this.RowBatch != nil {
		s = append(s, "RowBatch: "+fmt.Sprintf("%#v", this.RowBatch)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringCarnot(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *QueuedRowBatch) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueuedRowBatch) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCarnot(dAtA, i, uint64(len(m.Address)))
		i += copy(dAtA[i:], m.Address)
	}
	if m.QueryID != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCarnot(dAtA, i, uint64(m.QueryID.Size()))
		n1, err := m.QueryID.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.DestinationId != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintCarnot(dAtA, i, uint64(m.DestinationId))
	}
	if m.RowBatch != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintCarnot(dAtA, i, uint64(m.RowBatch.Size()))
		n2, err := m.RowBatch.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func encodeVarintCarnot(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *QueuedRowBatch) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovCarnot(uint64(l))
	}
	if m.QueryID != nil {
		l = m.QueryID.Size()
		n += 1 + l + sovCarnot(uint64(l))
	}
	if m.DestinationId != 0 {
		n += 1 + sovCarnot(uint64(m.DestinationId))
	}
	if m.RowBatch != nil {
		l = m.RowBatch.Size()
		n += 1 + l + sovCarnot(uint64(l))
	}
	return n
}

func sovCarnot(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCarnot(x uint64) (n int) {
	return sovCarnot(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *QueuedRowBatch) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&QueuedRowBatch{`,
		`Address:` + fmt.Sprintf("%v", this.Address) + `,`,
		`QueryID:` + strings.Replace(fmt.Sprintf("%v", this.QueryID), "UUID", "proto1.UUID", 1) + `,`,
		`DestinationId:` + fmt.Sprintf("%v", this.DestinationId) + `,`,
		`RowBatch:` + strings.Replace(fmt.Sprintf("%v", this.RowBatch), "RowBatchData", "proto2.RowBatchData", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringCarnot(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *QueuedRowBatch) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCarnot
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
			return fmt.Errorf("proto: QueuedRowBatch: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueuedRowBatch: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCarnot
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
				return ErrInvalidLengthCarnot
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCarnot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueryID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCarnot
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCarnot
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCarnot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.QueryID == nil {
				m.QueryID = &proto1.UUID{}
			}
			if err := m.QueryID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestinationId", wireType)
			}
			m.DestinationId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCarnot
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DestinationId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RowBatch", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCarnot
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCarnot
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCarnot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RowBatch == nil {
				m.RowBatch = &proto2.RowBatchData{}
			}
			if err := m.RowBatch.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCarnot(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCarnot
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCarnot
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
func skipCarnot(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCarnot
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
					return 0, ErrIntOverflowCarnot
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCarnot
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
				return 0, ErrInvalidLengthCarnot
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthCarnot
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCarnot
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipCarnot(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthCarnot
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthCarnot = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCarnot   = fmt.Errorf("proto: integer overflow")
)
