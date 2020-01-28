// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api.proto

// package 命名使用 {appid}.{version} 的方式, version 形如 v1, v2 ..

package api

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type BasicInfo struct {
	Nickname             string   `protobuf:"bytes,1,opt,name=nickname,proto3" json:"nickname,omitempty" json:"nickname"`
	Sign                 string   `protobuf:"bytes,2,opt,name=sign,proto3" json:"sign,omitempty" json:"sign"`
	ProfilePicUrl        string   `protobuf:"bytes,3,opt,name=profilePicUrl,proto3" json:"profilePicUrl,omitempty" json:"profile_pic_url"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty" validate:"required"`
	Uid                  int64    `protobuf:"varint,5,opt,name=uid,proto3" json:"uid,omitempty" json:"uid"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BasicInfo) Reset()         { *m = BasicInfo{} }
func (m *BasicInfo) String() string { return proto.CompactTextString(m) }
func (*BasicInfo) ProtoMessage()    {}
func (*BasicInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}
func (m *BasicInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BasicInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BasicInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BasicInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BasicInfo.Merge(m, src)
}
func (m *BasicInfo) XXX_Size() int {
	return m.Size()
}
func (m *BasicInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_BasicInfo.DiscardUnknown(m)
}

var xxx_messageInfo_BasicInfo proto.InternalMessageInfo

type BasicInfoRequest struct {
	Uid                  int64    `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty" form:"uid"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BasicInfoRequest) Reset()         { *m = BasicInfoRequest{} }
func (m *BasicInfoRequest) String() string { return proto.CompactTextString(m) }
func (*BasicInfoRequest) ProtoMessage()    {}
func (*BasicInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}
func (m *BasicInfoRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BasicInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BasicInfoRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BasicInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BasicInfoRequest.Merge(m, src)
}
func (m *BasicInfoRequest) XXX_Size() int {
	return m.Size()
}
func (m *BasicInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BasicInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BasicInfoRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*BasicInfo)(nil), "account.interface.v1.BasicInfo")
	proto.RegisterType((*BasicInfoRequest)(nil), "account.interface.v1.BasicInfoRequest")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 389 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xbb, 0x8e, 0xd4, 0x30,
	0x14, 0x86, 0xd7, 0x9b, 0x5d, 0x60, 0xcc, 0x65, 0x91, 0x77, 0x05, 0x21, 0x42, 0xc9, 0xc8, 0x48,
	0x68, 0x9b, 0x4d, 0xc4, 0xa5, 0xda, 0x0a, 0xd2, 0x20, 0x3a, 0x14, 0x89, 0x86, 0x66, 0xe5, 0x38,
	0x4e, 0x38, 0x90, 0xd8, 0x19, 0xc7, 0x9e, 0x92, 0x82, 0x57, 0xa0, 0xe1, 0x91, 0xa6, 0x44, 0xa2,
	0x8f, 0x60, 0xe0, 0x09, 0x22, 0x1e, 0x00, 0x8d, 0x33, 0x17, 0x40, 0x48, 0xdb, 0xe5, 0xe4, 0xfb,
	0x7e, 0x1f, 0xd9, 0x3f, 0x9e, 0xb0, 0x16, 0xe2, 0x56, 0x2b, 0xa3, 0xc8, 0x09, 0xe3, 0x5c, 0x59,
	0x69, 0x62, 0x90, 0x46, 0xe8, 0x92, 0x71, 0x11, 0xcf, 0x1f, 0x05, 0x67, 0x15, 0x98, 0xb7, 0x36,
	0x8f, 0xb9, 0x6a, 0x92, 0x4a, 0x55, 0x2a, 0x71, 0x72, 0x6e, 0x4b, 0x37, 0xb9, 0xc1, 0x7d, 0x8d,
	0x87, 0x04, 0xf7, 0x2b, 0xa5, 0xaa, 0x5a, 0x24, 0xac, 0x85, 0x84, 0x49, 0xa9, 0x0c, 0x33, 0xa0,
	0x64, 0x37, 0x52, 0xfa, 0x0b, 0xe1, 0x49, 0xca, 0x3a, 0xe0, 0x2f, 0x65, 0xa9, 0x48, 0x82, 0xaf,
	0x49, 0xe0, 0xef, 0x25, 0x6b, 0x84, 0x8f, 0xa6, 0xe8, 0x74, 0x92, 0x1e, 0x0f, 0x7d, 0x74, 0xf4,
	0xae, 0x53, 0xf2, 0x9c, 0x6e, 0x08, 0xcd, 0xb6, 0x12, 0x79, 0x80, 0x0f, 0x3a, 0xa8, 0xa4, 0xbf,
	0xef, 0xe4, 0xa3, 0xa1, 0x8f, 0xae, 0x8f, 0xf2, 0xea, 0x2f, 0xcd, 0x1c, 0x24, 0xcf, 0xf0, 0xcd,
	0x56, 0xab, 0x12, 0x6a, 0xf1, 0x0a, 0xf8, 0x6b, 0x5d, 0xfb, 0x9e, 0xb3, 0x83, 0xa1, 0x8f, 0xee,
	0x8c, 0xf6, 0x1a, 0x5f, 0xb4, 0xc0, 0x2f, 0xac, 0xae, 0x69, 0xf6, 0x77, 0x80, 0x9c, 0xe1, 0x43,
	0xd1, 0x30, 0xa8, 0xfd, 0x03, 0x97, 0xbc, 0x3b, 0xf4, 0xd1, 0xf1, 0x9c, 0xd5, 0x50, 0x30, 0x23,
	0xce, 0xa9, 0x16, 0x33, 0x0b, 0x5a, 0x14, 0x34, 0x1b, 0x2d, 0x32, 0xc5, 0x9e, 0x85, 0xc2, 0x3f,
	0x9c, 0xa2, 0x53, 0x2f, 0xbd, 0x35, 0xf4, 0x11, 0x1e, 0xd7, 0x58, 0x28, 0x68, 0xb6, 0x42, 0xf4,
	0x29, 0xbe, 0xbd, 0xbd, 0x75, 0x26, 0x66, 0x56, 0x74, 0x66, 0x93, 0x42, 0xbb, 0x54, 0xa9, 0x74,
	0xf3, 0x47, 0xea, 0xf1, 0x07, 0x7c, 0xf5, 0xf9, 0xd8, 0x08, 0xe9, 0xf0, 0x8d, 0x17, 0xc2, 0xec,
	0x5e, 0xee, 0x61, 0xfc, 0xbf, 0xae, 0xe2, 0x7f, 0x97, 0x04, 0xd1, 0x25, 0x1e, 0x0d, 0x3e, 0x7e,
	0xfd, 0xf9, 0x69, 0xff, 0x84, 0x90, 0x64, 0x2d, 0x26, 0xf9, 0x86, 0xa5, 0xf7, 0x16, 0xdf, 0xc3,
	0xbd, 0xc5, 0x32, 0x44, 0x5f, 0x96, 0x21, 0xfa, 0xb6, 0x0c, 0xd1, 0xe7, 0x1f, 0xe1, 0xde, 0x1b,
	0x8f, 0xb5, 0x90, 0x5f, 0x71, 0x75, 0x3e, 0xf9, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x60, 0x1d, 0x1f,
	0x84, 0x3e, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccountClient is the client API for Account service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountClient interface {
	GetBasicInfo(ctx context.Context, in *BasicInfoRequest, opts ...grpc.CallOption) (*BasicInfo, error)
}

type accountClient struct {
	cc *grpc.ClientConn
}

func NewAccountClient(cc *grpc.ClientConn) AccountClient {
	return &accountClient{cc}
}

func (c *accountClient) GetBasicInfo(ctx context.Context, in *BasicInfoRequest, opts ...grpc.CallOption) (*BasicInfo, error) {
	out := new(BasicInfo)
	err := c.cc.Invoke(ctx, "/account.interface.v1.Account/GetBasicInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServer is the server API for Account service.
type AccountServer interface {
	GetBasicInfo(context.Context, *BasicInfoRequest) (*BasicInfo, error)
}

// UnimplementedAccountServer can be embedded to have forward compatible implementations.
type UnimplementedAccountServer struct {
}

func (*UnimplementedAccountServer) GetBasicInfo(ctx context.Context, req *BasicInfoRequest) (*BasicInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBasicInfo not implemented")
}

func RegisterAccountServer(s *grpc.Server, srv AccountServer) {
	s.RegisterService(&_Account_serviceDesc, srv)
}

func _Account_GetBasicInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BasicInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).GetBasicInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.interface.v1.Account/GetBasicInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).GetBasicInfo(ctx, req.(*BasicInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Account_serviceDesc = grpc.ServiceDesc{
	ServiceName: "account.interface.v1.Account",
	HandlerType: (*AccountServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBasicInfo",
			Handler:    _Account_GetBasicInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func (m *BasicInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BasicInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BasicInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Uid != 0 {
		i = encodeVarintApi(dAtA, i, uint64(m.Uid))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Email) > 0 {
		i -= len(m.Email)
		copy(dAtA[i:], m.Email)
		i = encodeVarintApi(dAtA, i, uint64(len(m.Email)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ProfilePicUrl) > 0 {
		i -= len(m.ProfilePicUrl)
		copy(dAtA[i:], m.ProfilePicUrl)
		i = encodeVarintApi(dAtA, i, uint64(len(m.ProfilePicUrl)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Sign) > 0 {
		i -= len(m.Sign)
		copy(dAtA[i:], m.Sign)
		i = encodeVarintApi(dAtA, i, uint64(len(m.Sign)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Nickname) > 0 {
		i -= len(m.Nickname)
		copy(dAtA[i:], m.Nickname)
		i = encodeVarintApi(dAtA, i, uint64(len(m.Nickname)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BasicInfoRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BasicInfoRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BasicInfoRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Uid != 0 {
		i = encodeVarintApi(dAtA, i, uint64(m.Uid))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintApi(dAtA []byte, offset int, v uint64) int {
	offset -= sovApi(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BasicInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Nickname)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	l = len(m.Sign)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	l = len(m.ProfilePicUrl)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	l = len(m.Email)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	if m.Uid != 0 {
		n += 1 + sovApi(uint64(m.Uid))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *BasicInfoRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Uid != 0 {
		n += 1 + sovApi(uint64(m.Uid))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovApi(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozApi(x uint64) (n int) {
	return sovApi(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BasicInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
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
			return fmt.Errorf("proto: BasicInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BasicInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nickname", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
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
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nickname = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sign", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
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
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sign = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProfilePicUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
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
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProfilePicUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Email", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
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
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Email = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			m.Uid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Uid |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthApi
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
func (m *BasicInfoRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
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
			return fmt.Errorf("proto: BasicInfoRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BasicInfoRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			m.Uid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Uid |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthApi
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
func skipApi(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowApi
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
					return 0, ErrIntOverflowApi
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
					return 0, ErrIntOverflowApi
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
				return 0, ErrInvalidLengthApi
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupApi
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthApi
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthApi        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowApi          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupApi = fmt.Errorf("proto: unexpected end of group")
)
