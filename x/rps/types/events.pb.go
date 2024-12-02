// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lb/rps/v1/events.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	proto "github.com/cosmos/gogoproto/proto"
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

// EventCreateGame
type EventCreateGame struct {
	// game_number
	GameNumber uint64 `protobuf:"varint,1,opt,name=game_number,json=gameNumber,proto3" json:"game_number,omitempty"`
	// player_a
	PlayerA string `protobuf:"bytes,2,opt,name=player_a,json=playerA,proto3" json:"player_a,omitempty"`
	// player_b
	PlayerB string `protobuf:"bytes,3,opt,name=player_b,json=playerB,proto3" json:"player_b,omitempty"`
}

func (m *EventCreateGame) Reset()         { *m = EventCreateGame{} }
func (m *EventCreateGame) String() string { return proto.CompactTextString(m) }
func (*EventCreateGame) ProtoMessage()    {}
func (*EventCreateGame) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fbcb81318aa27ce, []int{0}
}
func (m *EventCreateGame) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventCreateGame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventCreateGame.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventCreateGame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventCreateGame.Merge(m, src)
}
func (m *EventCreateGame) XXX_Size() int {
	return m.Size()
}
func (m *EventCreateGame) XXX_DiscardUnknown() {
	xxx_messageInfo_EventCreateGame.DiscardUnknown(m)
}

var xxx_messageInfo_EventCreateGame proto.InternalMessageInfo

func (m *EventCreateGame) GetGameNumber() uint64 {
	if m != nil {
		return m.GameNumber
	}
	return 0
}

func (m *EventCreateGame) GetPlayerA() string {
	if m != nil {
		return m.PlayerA
	}
	return ""
}

func (m *EventCreateGame) GetPlayerB() string {
	if m != nil {
		return m.PlayerB
	}
	return ""
}

// EventMakeMove
type EventMakeMove struct {
	// game_number
	GameNumber uint64 `protobuf:"varint,1,opt,name=game_number,json=gameNumber,proto3" json:"game_number,omitempty"`
	// player
	Player string `protobuf:"bytes,2,opt,name=player,proto3" json:"player,omitempty"`
	// move
	Move string `protobuf:"bytes,3,opt,name=move,proto3" json:"move,omitempty"`
}

func (m *EventMakeMove) Reset()         { *m = EventMakeMove{} }
func (m *EventMakeMove) String() string { return proto.CompactTextString(m) }
func (*EventMakeMove) ProtoMessage()    {}
func (*EventMakeMove) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fbcb81318aa27ce, []int{1}
}
func (m *EventMakeMove) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventMakeMove) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventMakeMove.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventMakeMove) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventMakeMove.Merge(m, src)
}
func (m *EventMakeMove) XXX_Size() int {
	return m.Size()
}
func (m *EventMakeMove) XXX_DiscardUnknown() {
	xxx_messageInfo_EventMakeMove.DiscardUnknown(m)
}

var xxx_messageInfo_EventMakeMove proto.InternalMessageInfo

func (m *EventMakeMove) GetGameNumber() uint64 {
	if m != nil {
		return m.GameNumber
	}
	return 0
}

func (m *EventMakeMove) GetPlayer() string {
	if m != nil {
		return m.Player
	}
	return ""
}

func (m *EventMakeMove) GetMove() string {
	if m != nil {
		return m.Move
	}
	return ""
}

// EventRevealMove
type EventRevealMove struct {
	// game_number
	GameNumber uint64 `protobuf:"varint,1,opt,name=game_number,json=gameNumber,proto3" json:"game_number,omitempty"`
	// player
	Player string `protobuf:"bytes,2,opt,name=player,proto3" json:"player,omitempty"`
	// revealed_move
	RevealedMove string `protobuf:"bytes,3,opt,name=revealed_move,json=revealedMove,proto3" json:"revealed_move,omitempty"`
}

func (m *EventRevealMove) Reset()         { *m = EventRevealMove{} }
func (m *EventRevealMove) String() string { return proto.CompactTextString(m) }
func (*EventRevealMove) ProtoMessage()    {}
func (*EventRevealMove) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fbcb81318aa27ce, []int{2}
}
func (m *EventRevealMove) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventRevealMove) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventRevealMove.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventRevealMove) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventRevealMove.Merge(m, src)
}
func (m *EventRevealMove) XXX_Size() int {
	return m.Size()
}
func (m *EventRevealMove) XXX_DiscardUnknown() {
	xxx_messageInfo_EventRevealMove.DiscardUnknown(m)
}

var xxx_messageInfo_EventRevealMove proto.InternalMessageInfo

func (m *EventRevealMove) GetGameNumber() uint64 {
	if m != nil {
		return m.GameNumber
	}
	return 0
}

func (m *EventRevealMove) GetPlayer() string {
	if m != nil {
		return m.Player
	}
	return ""
}

func (m *EventRevealMove) GetRevealedMove() string {
	if m != nil {
		return m.RevealedMove
	}
	return ""
}

// EventEndGame
type EventEndGame struct {
	// game_number
	GameNumber uint64 `protobuf:"varint,1,opt,name=game_number,json=gameNumber,proto3" json:"game_number,omitempty"`
	// status
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (m *EventEndGame) Reset()         { *m = EventEndGame{} }
func (m *EventEndGame) String() string { return proto.CompactTextString(m) }
func (*EventEndGame) ProtoMessage()    {}
func (*EventEndGame) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fbcb81318aa27ce, []int{3}
}
func (m *EventEndGame) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventEndGame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventEndGame.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventEndGame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventEndGame.Merge(m, src)
}
func (m *EventEndGame) XXX_Size() int {
	return m.Size()
}
func (m *EventEndGame) XXX_DiscardUnknown() {
	xxx_messageInfo_EventEndGame.DiscardUnknown(m)
}

var xxx_messageInfo_EventEndGame proto.InternalMessageInfo

func (m *EventEndGame) GetGameNumber() uint64 {
	if m != nil {
		return m.GameNumber
	}
	return 0
}

func (m *EventEndGame) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*EventCreateGame)(nil), "lb.rps.v1.EventCreateGame")
	proto.RegisterType((*EventMakeMove)(nil), "lb.rps.v1.EventMakeMove")
	proto.RegisterType((*EventRevealMove)(nil), "lb.rps.v1.EventRevealMove")
	proto.RegisterType((*EventEndGame)(nil), "lb.rps.v1.EventEndGame")
}

func init() { proto.RegisterFile("lb/rps/v1/events.proto", fileDescriptor_1fbcb81318aa27ce) }

var fileDescriptor_1fbcb81318aa27ce = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0xb1, 0x4e, 0xc2, 0x40,
	0x1c, 0xc6, 0x39, 0x25, 0x28, 0x27, 0xc4, 0xa4, 0x31, 0xa4, 0x3a, 0x54, 0x52, 0x1d, 0x58, 0x6c,
	0x21, 0xec, 0x26, 0x60, 0x08, 0x13, 0x0e, 0x75, 0x73, 0x69, 0xee, 0xe8, 0x3f, 0x40, 0xec, 0xf5,
	0x9a, 0xbb, 0xe3, 0x02, 0x8f, 0xe0, 0xe6, 0x03, 0xf8, 0x18, 0x3e, 0x84, 0x23, 0x71, 0x72, 0x34,
	0xf0, 0x22, 0x86, 0x3b, 0x48, 0xd8, 0xec, 0xe2, 0xd6, 0xff, 0xbf, 0xdf, 0x77, 0xdf, 0x2f, 0x77,
	0x1f, 0x6e, 0xa4, 0x34, 0x14, 0xb9, 0x0c, 0x75, 0x27, 0x04, 0x0d, 0x99, 0x92, 0x41, 0x2e, 0xb8,
	0xe2, 0x4e, 0x35, 0xa5, 0x81, 0xc8, 0x65, 0xa0, 0x3b, 0x57, 0x97, 0x63, 0x2e, 0x19, 0x97, 0xb1,
	0xf9, 0x11, 0xda, 0xc1, 0xaa, 0xfc, 0x77, 0x84, 0xcf, 0x07, 0x5b, 0xdb, 0x83, 0x00, 0xa2, 0x60,
	0x48, 0x18, 0x38, 0xd7, 0xf8, 0x6c, 0x42, 0x18, 0xc4, 0xd9, 0x9c, 0x51, 0x10, 0x2e, 0x6a, 0xa2,
	0x56, 0x39, 0xc2, 0xdb, 0xd5, 0xa3, 0xd9, 0x38, 0x5d, 0x7c, 0x9a, 0xa7, 0x64, 0x09, 0x22, 0x26,
	0xee, 0x51, 0x13, 0xb5, 0xaa, 0x7d, 0xf7, 0xeb, 0xe3, 0xee, 0x62, 0x77, 0x70, 0x2f, 0x49, 0x04,
	0x48, 0xf9, 0xa4, 0xc4, 0x2c, 0x9b, 0x44, 0x27, 0x56, 0xd9, 0x3b, 0x30, 0x51, 0xf7, 0xb8, 0x98,
	0xa9, 0xef, 0x6b, 0x5c, 0x37, 0x74, 0x23, 0xf2, 0x02, 0x23, 0xae, 0x0b, 0xb0, 0xb5, 0x71, 0xc5,
	0x9a, 0xff, 0x24, 0xdb, 0xe9, 0x1c, 0x07, 0x97, 0x19, 0xd7, 0x60, 0xa1, 0x22, 0xf3, 0xed, 0xbf,
	0xee, 0xaf, 0x25, 0x02, 0x0d, 0x24, 0xfd, 0xaf, 0xe8, 0x1b, 0x5c, 0x17, 0x26, 0x00, 0x92, 0xf8,
	0x80, 0xa1, 0xb6, 0x5f, 0x6e, 0x73, 0xfd, 0x21, 0xae, 0x19, 0x94, 0x41, 0x96, 0x14, 0x7b, 0x9e,
	0x06, 0xae, 0x48, 0x45, 0xd4, 0x5c, 0x5a, 0x8e, 0x68, 0x37, 0xf5, 0xef, 0x3f, 0xd7, 0x1e, 0x5a,
	0xad, 0x3d, 0xf4, 0xb3, 0xf6, 0xd0, 0xdb, 0xc6, 0x2b, 0xad, 0x36, 0x5e, 0xe9, 0x7b, 0xe3, 0x95,
	0x9e, 0x6f, 0x27, 0x33, 0x35, 0x9d, 0xd3, 0x60, 0xcc, 0x59, 0xd8, 0x5e, 0xd8, 0x42, 0x8d, 0xa7,
	0x64, 0x96, 0x85, 0x0b, 0xd3, 0x2d, 0xb5, 0xcc, 0x41, 0xd2, 0x8a, 0xa9, 0x4c, 0xf7, 0x37, 0x00,
	0x00, 0xff, 0xff, 0xc3, 0x18, 0x0d, 0x23, 0x72, 0x02, 0x00, 0x00,
}

func (m *EventCreateGame) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventCreateGame) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventCreateGame) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PlayerB) > 0 {
		i -= len(m.PlayerB)
		copy(dAtA[i:], m.PlayerB)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.PlayerB)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.PlayerA) > 0 {
		i -= len(m.PlayerA)
		copy(dAtA[i:], m.PlayerA)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.PlayerA)))
		i--
		dAtA[i] = 0x12
	}
	if m.GameNumber != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.GameNumber))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EventMakeMove) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventMakeMove) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventMakeMove) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Move) > 0 {
		i -= len(m.Move)
		copy(dAtA[i:], m.Move)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Move)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Player) > 0 {
		i -= len(m.Player)
		copy(dAtA[i:], m.Player)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Player)))
		i--
		dAtA[i] = 0x12
	}
	if m.GameNumber != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.GameNumber))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EventRevealMove) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventRevealMove) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventRevealMove) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RevealedMove) > 0 {
		i -= len(m.RevealedMove)
		copy(dAtA[i:], m.RevealedMove)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.RevealedMove)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Player) > 0 {
		i -= len(m.Player)
		copy(dAtA[i:], m.Player)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Player)))
		i--
		dAtA[i] = 0x12
	}
	if m.GameNumber != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.GameNumber))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EventEndGame) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventEndGame) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventEndGame) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0x12
	}
	if m.GameNumber != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.GameNumber))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventCreateGame) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GameNumber != 0 {
		n += 1 + sovEvents(uint64(m.GameNumber))
	}
	l = len(m.PlayerA)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.PlayerB)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventMakeMove) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GameNumber != 0 {
		n += 1 + sovEvents(uint64(m.GameNumber))
	}
	l = len(m.Player)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Move)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventRevealMove) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GameNumber != 0 {
		n += 1 + sovEvents(uint64(m.GameNumber))
	}
	l = len(m.Player)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.RevealedMove)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventEndGame) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GameNumber != 0 {
		n += 1 + sovEvents(uint64(m.GameNumber))
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventCreateGame) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventCreateGame: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventCreateGame: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GameNumber", wireType)
			}
			m.GameNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GameNumber |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlayerA", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PlayerA = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlayerB", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PlayerB = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventMakeMove) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventMakeMove: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventMakeMove: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GameNumber", wireType)
			}
			m.GameNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GameNumber |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Player", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Player = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Move", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Move = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventRevealMove) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventRevealMove: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventRevealMove: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GameNumber", wireType)
			}
			m.GameNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GameNumber |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Player", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Player = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RevealedMove", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RevealedMove = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventEndGame) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventEndGame: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventEndGame: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GameNumber", wireType)
			}
			m.GameNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GameNumber |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
