// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dymension/streamer/gov.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type CreateStreamProposal struct {
	Title               string        `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description         string        `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	DistributeToRecords []DistrRecord `protobuf:"bytes,3,rep,name=distribute_to_records,json=distributeToRecords,proto3" json:"distribute_to_records"`
	// coins are coin(s) to be distributed by the stream
	Coins github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,4,rep,name=coins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"coins"`
	// start_time is the distribution start time
	StartTime            time.Time `protobuf:"bytes,5,opt,name=start_time,json=startTime,proto3,stdtime" json:"start_time" yaml:"timestamp"`
	DistrEpochIdentifier string    `protobuf:"bytes,6,opt,name=distr_epoch_identifier,json=distrEpochIdentifier,proto3" json:"distr_epoch_identifier,omitempty" yaml:"distr_epoch_identifier"`
	// num_epochs_paid_over is the number of epochs distribution will be completed
	// over
	NumEpochsPaidOver uint64 `protobuf:"varint,7,opt,name=num_epochs_paid_over,json=numEpochsPaidOver,proto3" json:"num_epochs_paid_over,omitempty"`
}

func (m *CreateStreamProposal) Reset()      { *m = CreateStreamProposal{} }
func (*CreateStreamProposal) ProtoMessage() {}
func (*CreateStreamProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_6da8318c7fb35093, []int{0}
}
func (m *CreateStreamProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreateStreamProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreateStreamProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreateStreamProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateStreamProposal.Merge(m, src)
}
func (m *CreateStreamProposal) XXX_Size() int {
	return m.Size()
}
func (m *CreateStreamProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateStreamProposal.DiscardUnknown(m)
}

var xxx_messageInfo_CreateStreamProposal proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CreateStreamProposal)(nil), "dymensionxyz.dymension.streamer.CreateStreamProposal")
}

func init() { proto.RegisterFile("dymension/streamer/gov.proto", fileDescriptor_6da8318c7fb35093) }

var fileDescriptor_6da8318c7fb35093 = []byte{
	// 508 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xbf, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0x6d, 0x9a, 0x16, 0xd5, 0x61, 0x00, 0x13, 0x90, 0x89, 0x8a, 0x2f, 0x84, 0x25, 0x03,
	0xdc, 0xd1, 0xb0, 0x75, 0x4c, 0x61, 0x40, 0x42, 0xa2, 0x32, 0x95, 0x2a, 0xb1, 0x58, 0x67, 0xfb,
	0xe2, 0x9e, 0x88, 0xfd, 0xac, 0xbb, 0x4b, 0xd4, 0xf0, 0x17, 0x20, 0xa6, 0x8e, 0x8c, 0x99, 0xf9,
	0x4b, 0x3a, 0x76, 0x64, 0x4a, 0x51, 0xb2, 0x30, 0xf7, 0x2f, 0x40, 0x77, 0xe7, 0xfc, 0x18, 0x2a,
	0x31, 0xd9, 0xef, 0xbd, 0xef, 0xf7, 0xfb, 0xac, 0xcf, 0xb3, 0x77, 0x90, 0x4d, 0x0b, 0x56, 0x4a,
	0x0e, 0x25, 0x91, 0x4a, 0x30, 0x5a, 0x30, 0x41, 0x72, 0x98, 0xe0, 0x4a, 0x80, 0x02, 0x1f, 0xad,
	0xa7, 0x17, 0xd3, 0x6f, 0x78, 0x5d, 0xe0, 0x95, 0xb4, 0xdd, 0xca, 0x21, 0x07, 0xa3, 0x25, 0xfa,
	0xcd, 0xda, 0xda, 0x61, 0x0a, 0xb2, 0x00, 0x49, 0x12, 0x2a, 0x19, 0x99, 0x1c, 0x26, 0x4c, 0xd1,
	0x43, 0x92, 0x02, 0x2f, 0xeb, 0x39, 0xca, 0x01, 0xf2, 0x11, 0x23, 0xa6, 0x4a, 0xc6, 0x43, 0xa2,
	0x78, 0xc1, 0xa4, 0xa2, 0x45, 0x55, 0x0b, 0x5e, 0xde, 0xf1, 0x55, 0x19, 0x97, 0x4a, 0xc4, 0xbc,
	0x1c, 0xd6, 0x5b, 0xba, 0x3f, 0x1a, 0x5e, 0xeb, 0x58, 0x30, 0xaa, 0xd8, 0x67, 0xa3, 0x39, 0x11,
	0x50, 0x81, 0xa4, 0x23, 0xbf, 0xe5, 0xed, 0x2a, 0xae, 0x46, 0x2c, 0x70, 0x3b, 0x6e, 0x6f, 0x3f,
	0xb2, 0x85, 0xdf, 0xf1, 0x9a, 0x19, 0x93, 0xa9, 0xe0, 0x95, 0xe2, 0x50, 0x06, 0xf7, 0xcc, 0x6c,
	0xbb, 0xe5, 0x0f, 0xbd, 0x27, 0x66, 0x09, 0x4f, 0xc6, 0x8a, 0xc5, 0x0a, 0x62, 0xc1, 0x52, 0x10,
	0x99, 0x0c, 0x76, 0x3a, 0x3b, 0xbd, 0x66, 0xff, 0x15, 0xfe, 0x0f, 0x0d, 0xfc, 0x4e, 0xbb, 0x23,
	0x63, 0x1a, 0x34, 0xae, 0xe6, 0xc8, 0x89, 0x1e, 0x6f, 0x02, 0x4f, 0xc1, 0x4e, 0xa4, 0x4f, 0xbd,
	0x5d, 0x0d, 0x43, 0x06, 0x0d, 0x93, 0xfb, 0x0c, 0x5b, 0x5c, 0x58, 0xe3, 0xc2, 0x35, 0x2e, 0x7c,
	0x0c, 0xbc, 0x1c, 0xbc, 0xd1, 0x21, 0xbf, 0x6e, 0x50, 0x2f, 0xe7, 0xea, 0x7c, 0x9c, 0xe0, 0x14,
	0x0a, 0x52, 0xb3, 0xb5, 0x8f, 0xd7, 0x32, 0xfb, 0x4a, 0xd4, 0xb4, 0x62, 0xd2, 0x18, 0x64, 0x64,
	0x93, 0xfd, 0x33, 0xcf, 0x93, 0x8a, 0x0a, 0x15, 0x6b, 0xb2, 0xc1, 0x6e, 0xc7, 0xed, 0x35, 0xfb,
	0x6d, 0x6c, 0xb1, 0xe3, 0x15, 0x76, 0x7c, 0xba, 0xc2, 0x3e, 0x38, 0xd0, 0x8b, 0x6e, 0xe7, 0xe8,
	0xe1, 0x94, 0x16, 0xa3, 0xa3, 0xee, 0xfa, 0x1e, 0xdd, 0xcb, 0x1b, 0xe4, 0x46, 0xfb, 0x26, 0x4b,
	0xab, 0xfd, 0x33, 0xef, 0xa9, 0x3d, 0x04, 0xab, 0x20, 0x3d, 0x8f, 0x79, 0xc6, 0x4a, 0xc5, 0x87,
	0x9c, 0x89, 0x60, 0x4f, 0x03, 0x1d, 0xbc, 0xb8, 0x9d, 0xa3, 0xe7, 0x36, 0xe4, 0x6e, 0x5d, 0x37,
	0x6a, 0x99, 0xc1, 0x7b, 0xdd, 0xff, 0xb0, 0x6e, 0xfb, 0xc4, 0x6b, 0x95, 0xe3, 0xc2, 0xca, 0x65,
	0x5c, 0x51, 0x9e, 0xc5, 0x30, 0x61, 0x22, 0xb8, 0xdf, 0x71, 0x7b, 0x8d, 0xe8, 0x51, 0x39, 0x2e,
	0x8c, 0x43, 0x9e, 0x50, 0x9e, 0x7d, 0x9a, 0x30, 0x71, 0xf4, 0xe0, 0xfb, 0x0c, 0x39, 0x3f, 0x67,
	0xc8, 0xf9, 0x3b, 0x43, 0xee, 0xe0, 0xe3, 0xd5, 0x22, 0x74, 0xaf, 0x17, 0xa1, 0xfb, 0x67, 0x11,
	0xba, 0x97, 0xcb, 0xd0, 0xb9, 0x5e, 0x86, 0xce, 0xef, 0x65, 0xe8, 0x7c, 0xe9, 0x6f, 0xb1, 0xdb,
	0x3e, 0xe0, 0xa6, 0x20, 0x17, 0x9b, 0xbf, 0xcc, 0xb0, 0x4c, 0xf6, 0x0c, 0xa2, 0xb7, 0xff, 0x02,
	0x00, 0x00, 0xff, 0xff, 0xd7, 0x89, 0x7b, 0xa7, 0x1e, 0x03, 0x00, 0x00,
}

func (this *CreateStreamProposal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CreateStreamProposal)
	if !ok {
		that2, ok := that.(CreateStreamProposal)
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
	if this.Title != that1.Title {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if len(this.DistributeToRecords) != len(that1.DistributeToRecords) {
		return false
	}
	for i := range this.DistributeToRecords {
		if !this.DistributeToRecords[i].Equal(&that1.DistributeToRecords[i]) {
			return false
		}
	}
	if len(this.Coins) != len(that1.Coins) {
		return false
	}
	for i := range this.Coins {
		if !this.Coins[i].Equal(&that1.Coins[i]) {
			return false
		}
	}
	if !this.StartTime.Equal(that1.StartTime) {
		return false
	}
	if this.DistrEpochIdentifier != that1.DistrEpochIdentifier {
		return false
	}
	if this.NumEpochsPaidOver != that1.NumEpochsPaidOver {
		return false
	}
	return true
}
func (m *CreateStreamProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateStreamProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreateStreamProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NumEpochsPaidOver != 0 {
		i = encodeVarintGov(dAtA, i, uint64(m.NumEpochsPaidOver))
		i--
		dAtA[i] = 0x38
	}
	if len(m.DistrEpochIdentifier) > 0 {
		i -= len(m.DistrEpochIdentifier)
		copy(dAtA[i:], m.DistrEpochIdentifier)
		i = encodeVarintGov(dAtA, i, uint64(len(m.DistrEpochIdentifier)))
		i--
		dAtA[i] = 0x32
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.StartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintGov(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x2a
	if len(m.Coins) > 0 {
		for iNdEx := len(m.Coins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Coins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGov(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.DistributeToRecords) > 0 {
		for iNdEx := len(m.DistributeToRecords) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DistributeToRecords[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGov(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGov(dAtA []byte, offset int, v uint64) int {
	offset -= sovGov(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CreateStreamProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	if len(m.DistributeToRecords) > 0 {
		for _, e := range m.DistributeToRecords {
			l = e.Size()
			n += 1 + l + sovGov(uint64(l))
		}
	}
	if len(m.Coins) > 0 {
		for _, e := range m.Coins {
			l = e.Size()
			n += 1 + l + sovGov(uint64(l))
		}
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime)
	n += 1 + l + sovGov(uint64(l))
	l = len(m.DistrEpochIdentifier)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	if m.NumEpochsPaidOver != 0 {
		n += 1 + sovGov(uint64(m.NumEpochsPaidOver))
	}
	return n
}

func sovGov(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGov(x uint64) (n int) {
	return sovGov(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CreateStreamProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: CreateStreamProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateStreamProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistributeToRecords", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DistributeToRecords = append(m.DistributeToRecords, DistrRecord{})
			if err := m.DistributeToRecords[len(m.DistributeToRecords)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coins = append(m.Coins, types.Coin{})
			if err := m.Coins[len(m.Coins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.StartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistrEpochIdentifier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DistrEpochIdentifier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumEpochsPaidOver", wireType)
			}
			m.NumEpochsPaidOver = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumEpochsPaidOver |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
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
func skipGov(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGov
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
					return 0, ErrIntOverflowGov
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
					return 0, ErrIntOverflowGov
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
				return 0, ErrInvalidLengthGov
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGov
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGov
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGov        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGov          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGov = fmt.Errorf("proto: unexpected end of group")
)
