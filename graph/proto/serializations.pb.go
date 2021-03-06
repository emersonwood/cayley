// Code generated by protoc-gen-gogo.
// source: serializations.proto
// DO NOT EDIT!

/*
	Package proto is a generated protocol buffer package.

	It is generated from these files:
		serializations.proto

	It has these top-level messages:
		LogDelta
		HistoryEntry
		NodeData
*/
package proto

import proto1 "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import pquads "github.com/codelingo/cayley/quad/pquads"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto1.GoGoProtoPackageIsVersion1

type LogDelta struct {
	ID        uint64       `protobuf:"varint,1,opt,name=ID,json=iD,proto3" json:"ID,omitempty"`
	Quad      *pquads.Quad `protobuf:"bytes,2,opt,name=Quad,json=quad" json:"Quad,omitempty"`
	Action    int32        `protobuf:"varint,3,opt,name=Action,json=action,proto3" json:"Action,omitempty"`
	Timestamp int64        `protobuf:"varint,4,opt,name=Timestamp,json=timestamp,proto3" json:"Timestamp,omitempty"`
}

func (m *LogDelta) Reset()                    { *m = LogDelta{} }
func (m *LogDelta) String() string            { return proto1.CompactTextString(m) }
func (*LogDelta) ProtoMessage()               {}
func (*LogDelta) Descriptor() ([]byte, []int) { return fileDescriptorSerializations, []int{0} }

func (m *LogDelta) GetQuad() *pquads.Quad {
	if m != nil {
		return m.Quad
	}
	return nil
}

type HistoryEntry struct {
	History []uint64 `protobuf:"varint,1,rep,name=History,json=history" json:"History,omitempty"`
}

func (m *HistoryEntry) Reset()                    { *m = HistoryEntry{} }
func (m *HistoryEntry) String() string            { return proto1.CompactTextString(m) }
func (*HistoryEntry) ProtoMessage()               {}
func (*HistoryEntry) Descriptor() ([]byte, []int) { return fileDescriptorSerializations, []int{1} }

type NodeData struct {
	Name  string        `protobuf:"bytes,1,opt,name=Name,json=name,proto3" json:"Name,omitempty"`
	Size  int64         `protobuf:"varint,2,opt,name=Size,json=size,proto3" json:"Size,omitempty"`
	Value *pquads.Value `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
}

func (m *NodeData) Reset()                    { *m = NodeData{} }
func (m *NodeData) String() string            { return proto1.CompactTextString(m) }
func (*NodeData) ProtoMessage()               {}
func (*NodeData) Descriptor() ([]byte, []int) { return fileDescriptorSerializations, []int{2} }

func (m *NodeData) GetValue() *pquads.Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto1.RegisterType((*LogDelta)(nil), "proto.LogDelta")
	proto1.RegisterType((*HistoryEntry)(nil), "proto.HistoryEntry")
	proto1.RegisterType((*NodeData)(nil), "proto.NodeData")
}
func (m *LogDelta) Marshal() (data []byte, err error) {
	size := m.ProtoSize()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *LogDelta) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ID != 0 {
		data[i] = 0x8
		i++
		i = encodeVarintSerializations(data, i, uint64(m.ID))
	}
	if m.Quad != nil {
		data[i] = 0x12
		i++
		i = encodeVarintSerializations(data, i, uint64(m.Quad.ProtoSize()))
		n1, err := m.Quad.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Action != 0 {
		data[i] = 0x18
		i++
		i = encodeVarintSerializations(data, i, uint64(m.Action))
	}
	if m.Timestamp != 0 {
		data[i] = 0x20
		i++
		i = encodeVarintSerializations(data, i, uint64(m.Timestamp))
	}
	return i, nil
}

func (m *HistoryEntry) Marshal() (data []byte, err error) {
	size := m.ProtoSize()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *HistoryEntry) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.History) > 0 {
		for _, num := range m.History {
			data[i] = 0x8
			i++
			i = encodeVarintSerializations(data, i, uint64(num))
		}
	}
	return i, nil
}

func (m *NodeData) Marshal() (data []byte, err error) {
	size := m.ProtoSize()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *NodeData) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintSerializations(data, i, uint64(len(m.Name)))
		i += copy(data[i:], m.Name)
	}
	if m.Size != 0 {
		data[i] = 0x10
		i++
		i = encodeVarintSerializations(data, i, uint64(m.Size))
	}
	if m.Value != nil {
		data[i] = 0x1a
		i++
		i = encodeVarintSerializations(data, i, uint64(m.Value.ProtoSize()))
		n2, err := m.Value.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func encodeFixed64Serializations(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Serializations(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintSerializations(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *LogDelta) ProtoSize() (n int) {
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovSerializations(uint64(m.ID))
	}
	if m.Quad != nil {
		l = m.Quad.ProtoSize()
		n += 1 + l + sovSerializations(uint64(l))
	}
	if m.Action != 0 {
		n += 1 + sovSerializations(uint64(m.Action))
	}
	if m.Timestamp != 0 {
		n += 1 + sovSerializations(uint64(m.Timestamp))
	}
	return n
}

func (m *HistoryEntry) ProtoSize() (n int) {
	var l int
	_ = l
	if len(m.History) > 0 {
		for _, e := range m.History {
			n += 1 + sovSerializations(uint64(e))
		}
	}
	return n
}

func (m *NodeData) ProtoSize() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovSerializations(uint64(l))
	}
	if m.Size != 0 {
		n += 1 + sovSerializations(uint64(m.Size))
	}
	if m.Value != nil {
		l = m.Value.ProtoSize()
		n += 1 + l + sovSerializations(uint64(l))
	}
	return n
}

func sovSerializations(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSerializations(x uint64) (n int) {
	return sovSerializations(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LogDelta) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSerializations
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LogDelta: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LogDelta: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSerializations
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.ID |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Quad", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSerializations
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSerializations
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Quad == nil {
				m.Quad = &pquads.Quad{}
			}
			if err := m.Quad.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Action", wireType)
			}
			m.Action = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSerializations
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Action |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSerializations
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Timestamp |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSerializations(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSerializations
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
func (m *HistoryEntry) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSerializations
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HistoryEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HistoryEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field History", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSerializations
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.History = append(m.History, v)
		default:
			iNdEx = preIndex
			skippy, err := skipSerializations(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSerializations
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
func (m *NodeData) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSerializations
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NodeData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NodeData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSerializations
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSerializations
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Size", wireType)
			}
			m.Size = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSerializations
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Size |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSerializations
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSerializations
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Value == nil {
				m.Value = &pquads.Value{}
			}
			if err := m.Value.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSerializations(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSerializations
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
func skipSerializations(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSerializations
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowSerializations
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowSerializations
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthSerializations
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSerializations
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
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
				next, err := skipSerializations(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthSerializations = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSerializations   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorSerializations = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x90, 0xdf, 0x4a, 0xc3, 0x30,
	0x14, 0xc6, 0xe9, 0x9a, 0xfd, 0x69, 0x56, 0xbd, 0x08, 0x22, 0x65, 0xc8, 0x18, 0xf5, 0x66, 0x37,
	0xb6, 0x30, 0xf1, 0x01, 0x94, 0x0a, 0x0a, 0x32, 0x30, 0x8a, 0x5e, 0xa7, 0x5d, 0x6c, 0x03, 0x6d,
	0x53, 0x93, 0x54, 0xe8, 0x9e, 0xc6, 0xc7, 0xf1, 0xd2, 0x67, 0xd0, 0x17, 0x31, 0x49, 0x3b, 0xf0,
	0x22, 0xc9, 0xf7, 0x7d, 0xe7, 0x1c, 0xce, 0x8f, 0xc0, 0x13, 0x49, 0x05, 0x23, 0x25, 0xdb, 0x13,
	0xc5, 0x78, 0x2d, 0xa3, 0x46, 0x70, 0xc5, 0xd1, 0xd8, 0x3e, 0x8b, 0x8b, 0x9c, 0xa9, 0xa2, 0x4d,
	0xa3, 0x8c, 0x57, 0x71, 0xce, 0x73, 0x1e, 0xdb, 0x38, 0x6d, 0xdf, 0xac, 0xb3, 0xc6, 0xaa, 0x7e,
	0x6a, 0x71, 0xf5, 0xaf, 0x3d, 0x23, 0x5d, 0x49, 0xbb, 0x5c, 0x90, 0xa6, 0x18, 0x74, 0xfc, 0xde,
	0x92, 0x5d, 0xdc, 0x98, 0x5b, 0x5a, 0x3d, 0x2c, 0x0b, 0x05, 0x9c, 0x3d, 0xf0, 0x3c, 0xa1, 0xa5,
	0x22, 0xe8, 0x18, 0x8e, 0xee, 0x93, 0xc0, 0x59, 0x39, 0x6b, 0x80, 0x47, 0x2c, 0x41, 0x2b, 0x08,
	0x1e, 0x75, 0x6b, 0x30, 0xd2, 0xc9, 0x7c, 0xe3, 0x47, 0xfd, 0x78, 0x64, 0x32, 0x0c, 0x8c, 0x46,
	0xa7, 0x70, 0x72, 0x9d, 0x19, 0xf6, 0xc0, 0xd5, 0x3d, 0x63, 0x3c, 0x21, 0xd6, 0xa1, 0x33, 0xe8,
	0x3d, 0xb3, 0x8a, 0x4a, 0x45, 0xaa, 0x26, 0x00, 0xba, 0xe4, 0x62, 0x4f, 0x1d, 0x82, 0x70, 0x0d,
	0xfd, 0x3b, 0x26, 0x15, 0x17, 0xdd, 0x6d, 0xad, 0x44, 0x87, 0x02, 0x38, 0x1d, 0xbc, 0x5e, 0xee,
	0xea, 0xe5, 0xd3, 0xa2, 0xb7, 0xe1, 0x2b, 0x9c, 0x6d, 0xf9, 0x8e, 0x26, 0x44, 0xd3, 0x21, 0x08,
	0xb6, 0xa4, 0xa2, 0x96, 0xcf, 0xc3, 0xa0, 0xd6, 0xda, 0x64, 0x4f, 0x6c, 0x4f, 0x2d, 0xa1, 0x8b,
	0x81, 0xd4, 0x1a, 0x9d, 0xc3, 0xf1, 0x07, 0x29, 0x5b, 0x6a, 0x91, 0xe6, 0x9b, 0xa3, 0x03, 0xf6,
	0x8b, 0x09, 0x71, 0x5f, 0xbb, 0xf1, 0xbf, 0x7e, 0x96, 0xce, 0xb7, 0x3e, 0x9f, 0xbf, 0x4b, 0x27,
	0x9d, 0xd8, 0xbf, 0xb8, 0xfc, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xc4, 0x7e, 0x15, 0x5d, 0x90, 0x01,
	0x00, 0x00,
}
