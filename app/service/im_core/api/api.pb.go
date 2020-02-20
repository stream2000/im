// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api.proto

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

type ConnectReq struct {
	Server string `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty" validate:"required"`
	// the jwt include uid, device id
	Jwt                  string   `protobuf:"bytes,2,opt,name=jwt,proto3" json:"jwt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnectReq) Reset()         { *m = ConnectReq{} }
func (m *ConnectReq) String() string { return proto.CompactTextString(m) }
func (*ConnectReq) ProtoMessage()    {}
func (*ConnectReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}
func (m *ConnectReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConnectReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConnectReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConnectReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectReq.Merge(m, src)
}
func (m *ConnectReq) XXX_Size() int {
	return m.Size()
}
func (m *ConnectReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectReq.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectReq proto.InternalMessageInfo

type ConnectResp struct {
	Uid                  int64    `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	DeviceId             string   `protobuf:"bytes,2,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	TokenId              string   `protobuf:"bytes,3,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnectResp) Reset()         { *m = ConnectResp{} }
func (m *ConnectResp) String() string { return proto.CompactTextString(m) }
func (*ConnectResp) ProtoMessage()    {}
func (*ConnectResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}
func (m *ConnectResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConnectResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConnectResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConnectResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectResp.Merge(m, src)
}
func (m *ConnectResp) XXX_Size() int {
	return m.Size()
}
func (m *ConnectResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectResp.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectResp proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ConnectReq)(nil), "imCore.service.v1.ConnectReq")
	proto.RegisterType((*ConnectResp)(nil), "imCore.service.v1.ConnectResp")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xbf, 0x4e, 0xfb, 0x30,
	0x18, 0x6c, 0x7e, 0x91, 0xfa, 0xc7, 0xbf, 0x05, 0xcc, 0x40, 0x5b, 0xc0, 0xa0, 0x4c, 0x2c, 0x38,
	0x02, 0x36, 0xc6, 0x76, 0xa1, 0x13, 0x52, 0x16, 0x24, 0x16, 0xe4, 0xd6, 0x1f, 0xc1, 0xd0, 0xfa,
	0x73, 0x1d, 0x27, 0xbc, 0x0a, 0x8f, 0xd4, 0x91, 0x27, 0x40, 0x10, 0xde, 0x80, 0x27, 0x40, 0x76,
	0x22, 0x18, 0x90, 0xd8, 0xee, 0x7c, 0xe7, 0xfb, 0x74, 0x47, 0x06, 0xc2, 0x28, 0x6e, 0x2c, 0x3a,
	0xa4, 0xdb, 0x6a, 0x35, 0x45, 0x0b, 0xbc, 0x00, 0x5b, 0xa9, 0x05, 0xf0, 0xea, 0x74, 0x7c, 0x92,
	0x2b, 0x77, 0x5f, 0xce, 0xf9, 0x02, 0x57, 0x69, 0x8e, 0x39, 0xa6, 0xc1, 0x39, 0x2f, 0xef, 0x02,
	0x0b, 0x24, 0xa0, 0x26, 0x61, 0xbc, 0x9f, 0x23, 0xe6, 0x4b, 0x48, 0x85, 0x51, 0xa9, 0xd0, 0x1a,
	0x9d, 0x70, 0x0a, 0x75, 0xd1, 0xa8, 0xc9, 0x15, 0x21, 0x53, 0xd4, 0x1a, 0x16, 0x2e, 0x83, 0x35,
	0x4d, 0x49, 0xd7, 0x1f, 0x02, 0x3b, 0x8c, 0x8e, 0xa2, 0xe3, 0xc1, 0x64, 0xf7, 0xf3, 0xf5, 0x70,
	0xa7, 0x12, 0x4b, 0x25, 0x85, 0x83, 0x8b, 0xc4, 0xc2, 0xba, 0x54, 0x16, 0x64, 0x92, 0xb5, 0x36,
	0xba, 0x45, 0xe2, 0x87, 0x27, 0x37, 0xfc, 0xe7, 0xdd, 0x99, 0x87, 0xc9, 0x35, 0xf9, 0xff, 0x1d,
	0x58, 0x18, 0x6f, 0x28, 0x95, 0x0c, 0x71, 0x71, 0xe6, 0x21, 0xdd, 0x23, 0x03, 0x09, 0xbe, 0xcb,
	0xad, 0x92, 0xed, 0xc7, 0x7e, 0xf3, 0x30, 0x93, 0x74, 0x44, 0xfa, 0x0e, 0x1f, 0x41, 0x7b, 0x2d,
	0x0e, 0x5a, 0x2f, 0xf0, 0x99, 0x3c, 0xcb, 0x48, 0xb7, 0xd9, 0x82, 0x5e, 0x92, 0x5e, 0x7b, 0x82,
	0x1e, 0xf0, 0x5f, 0xfb, 0xf0, 0x9f, 0x3e, 0x63, 0xf6, 0x97, 0x5c, 0x98, 0xc9, 0x68, 0xf3, 0xce,
	0x3a, 0x9b, 0x9a, 0x45, 0x2f, 0x35, 0x8b, 0xde, 0x6a, 0x16, 0x3d, 0x7f, 0xb0, 0xce, 0x4d, 0x2c,
	0x8c, 0x9a, 0x77, 0xc3, 0x3e, 0xe7, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x5b, 0xc5, 0xc6, 0x92,
	0x8c, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ImCoreClient is the client API for ImCore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ImCoreClient interface {
	Connect(ctx context.Context, in *ConnectReq, opts ...grpc.CallOption) (*ConnectResp, error)
}

type imCoreClient struct {
	cc *grpc.ClientConn
}

func NewImCoreClient(cc *grpc.ClientConn) ImCoreClient {
	return &imCoreClient{cc}
}

func (c *imCoreClient) Connect(ctx context.Context, in *ConnectReq, opts ...grpc.CallOption) (*ConnectResp, error) {
	out := new(ConnectResp)
	err := c.cc.Invoke(ctx, "/imCore.service.v1.imCore/Connect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImCoreServer is the server API for ImCore service.
type ImCoreServer interface {
	Connect(context.Context, *ConnectReq) (*ConnectResp, error)
}

// UnimplementedImCoreServer can be embedded to have forward compatible implementations.
type UnimplementedImCoreServer struct {
}

func (*UnimplementedImCoreServer) Connect(ctx context.Context, req *ConnectReq) (*ConnectResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connect not implemented")
}

func RegisterImCoreServer(s *grpc.Server, srv ImCoreServer) {
	s.RegisterService(&_ImCore_serviceDesc, srv)
}

func _ImCore_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImCoreServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/imCore.service.v1.imCore/Connect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImCoreServer).Connect(ctx, req.(*ConnectReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _ImCore_serviceDesc = grpc.ServiceDesc{
	ServiceName: "imCore.service.v1.imCore",
	HandlerType: (*ImCoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Connect",
			Handler:    _ImCore_Connect_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func (m *ConnectReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConnectReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConnectReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Jwt) > 0 {
		i -= len(m.Jwt)
		copy(dAtA[i:], m.Jwt)
		i = encodeVarintApi(dAtA, i, uint64(len(m.Jwt)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Server) > 0 {
		i -= len(m.Server)
		copy(dAtA[i:], m.Server)
		i = encodeVarintApi(dAtA, i, uint64(len(m.Server)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ConnectResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConnectResp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConnectResp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.TokenId) > 0 {
		i -= len(m.TokenId)
		copy(dAtA[i:], m.TokenId)
		i = encodeVarintApi(dAtA, i, uint64(len(m.TokenId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.DeviceId) > 0 {
		i -= len(m.DeviceId)
		copy(dAtA[i:], m.DeviceId)
		i = encodeVarintApi(dAtA, i, uint64(len(m.DeviceId)))
		i--
		dAtA[i] = 0x12
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
func (m *ConnectReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Server)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	l = len(m.Jwt)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ConnectResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Uid != 0 {
		n += 1 + sovApi(uint64(m.Uid))
	}
	l = len(m.DeviceId)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	l = len(m.TokenId)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
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
func (m *ConnectReq) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ConnectReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConnectReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Server", wireType)
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
			m.Server = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Jwt", wireType)
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
			m.Jwt = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
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
func (m *ConnectResp) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ConnectResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConnectResp: illegal tag %d (wire type %d)", fieldNum, wire)
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeviceId", wireType)
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
			m.DeviceId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenId", wireType)
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
			m.TokenId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
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
