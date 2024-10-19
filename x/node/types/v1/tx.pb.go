// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: imhub/node/v1/tx.proto

package v1

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "github.com/dun-io/imhub/types/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

type MsgRegisterNodeRequest struct {
	From    string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Address string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	PubKey  string `protobuf:"bytes,4,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty"`
}

func (m *MsgRegisterNodeRequest) Reset()         { *m = MsgRegisterNodeRequest{} }
func (m *MsgRegisterNodeRequest) String() string { return proto.CompactTextString(m) }
func (*MsgRegisterNodeRequest) ProtoMessage()    {}
func (*MsgRegisterNodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_992faf421a3a872b, []int{0}
}
func (m *MsgRegisterNodeRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRegisterNodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRegisterNodeRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRegisterNodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRegisterNodeRequest.Merge(m, src)
}
func (m *MsgRegisterNodeRequest) XXX_Size() int {
	return m.Size()
}
func (m *MsgRegisterNodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRegisterNodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRegisterNodeRequest proto.InternalMessageInfo

type MsgRegisterNodeResponse struct {
}

func (m *MsgRegisterNodeResponse) Reset()         { *m = MsgRegisterNodeResponse{} }
func (m *MsgRegisterNodeResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRegisterNodeResponse) ProtoMessage()    {}
func (*MsgRegisterNodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_992faf421a3a872b, []int{1}
}
func (m *MsgRegisterNodeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRegisterNodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRegisterNodeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRegisterNodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRegisterNodeResponse.Merge(m, src)
}
func (m *MsgRegisterNodeResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRegisterNodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRegisterNodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRegisterNodeResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgRegisterNodeRequest)(nil), "imhub.node.v1.MsgRegisterNodeRequest")
	proto.RegisterType((*MsgRegisterNodeResponse)(nil), "imhub.node.v1.MsgRegisterNodeResponse")
}

func init() { proto.RegisterFile("imhub/node/v1/tx.proto", fileDescriptor_992faf421a3a872b) }

var fileDescriptor_992faf421a3a872b = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x51, 0xcd, 0xea, 0xd3, 0x40,
	0x10, 0x4f, 0xb4, 0xb4, 0xb8, 0x20, 0x62, 0x90, 0x36, 0x8d, 0xb0, 0x4a, 0xc1, 0x8f, 0x8b, 0x59,
	0xaa, 0x6f, 0xe0, 0x4d, 0xa4, 0x1e, 0xea, 0xcd, 0x4b, 0xdd, 0x24, 0xd3, 0x75, 0xd1, 0xcd, 0xc4,
	0xcc, 0xa6, 0xb4, 0x6f, 0xe1, 0x63, 0xf8, 0x28, 0x3d, 0xf6, 0xe8, 0x51, 0xd3, 0x17, 0x91, 0xdd,
	0x4d, 0x0f, 0xd5, 0x3f, 0xfc, 0x2f, 0xcb, 0xfc, 0xbe, 0x76, 0x86, 0x19, 0x36, 0xd5, 0xe6, 0x4b,
	0x57, 0x88, 0x1a, 0x2b, 0x10, 0xbb, 0xa5, 0xb0, 0xfb, 0xbc, 0x69, 0xd1, 0x62, 0x72, 0xdf, 0xf3,
	0xb9, 0xe3, 0xf3, 0xdd, 0x32, 0x7b, 0x28, 0x8d, 0xae, 0x51, 0xf8, 0x37, 0x38, 0xb2, 0x59, 0x89,
	0x64, 0x90, 0x84, 0x21, 0xe5, 0x92, 0x86, 0xd4, 0x20, 0xcc, 0x83, 0xb0, 0xf1, 0x48, 0x04, 0x30,
	0x48, 0x7c, 0xc8, 0x14, 0x92, 0x5c, 0xbb, 0x02, 0xac, 0x5c, 0x8a, 0x12, 0x75, 0x3d, 0xe8, 0x8f,
	0x14, 0x2a, 0x0c, 0x39, 0x57, 0x0d, 0x6c, 0x76, 0x3d, 0x63, 0x23, 0x5b, 0x69, 0x2e, 0x3f, 0x3e,
	0x51, 0x88, 0xea, 0x1b, 0x08, 0x8f, 0x8a, 0x6e, 0x2b, 0xac, 0x36, 0x40, 0x56, 0x9a, 0x66, 0x30,
	0x3c, 0x0e, 0x61, 0x7b, 0x68, 0x80, 0x5c, 0x9a, 0xac, 0xb4, 0xdd, 0x90, 0x5e, 0x10, 0x9b, 0xae,
	0x48, 0xad, 0x41, 0x69, 0xb2, 0xd0, 0x7e, 0xc0, 0x0a, 0xd6, 0xf0, 0xbd, 0x03, 0xb2, 0x49, 0xc2,
	0x46, 0xdb, 0x16, 0x4d, 0x1a, 0x3f, 0x8d, 0x5f, 0xde, 0x5b, 0xfb, 0xda, 0x71, 0xb5, 0x34, 0x90,
	0xde, 0x09, 0x9c, 0xab, 0x93, 0x94, 0x4d, 0x64, 0x55, 0xb5, 0x40, 0x94, 0xde, 0xf5, 0xf4, 0x05,
	0x26, 0x33, 0x36, 0x69, 0xba, 0x62, 0xf3, 0x15, 0x0e, 0xe9, 0xc8, 0x2b, 0xe3, 0xa6, 0x2b, 0xde,
	0xc3, 0x61, 0x31, 0x67, 0xb3, 0xff, 0x9a, 0x52, 0x83, 0x35, 0xc1, 0xeb, 0x9a, 0xb1, 0x15, 0xa9,
	0x8f, 0xd0, 0xee, 0x74, 0x09, 0xc9, 0x67, 0xf6, 0xe0, 0x1f, 0x63, 0xf2, 0x2c, 0xbf, 0xba, 0x4b,
	0x7e, 0xf3, 0xf4, 0xd9, 0xf3, 0xdb, 0x6c, 0xa1, 0xdf, 0xdb, 0x77, 0xc7, 0x3f, 0x3c, 0xfa, 0xd9,
	0xf3, 0xe8, 0xd8, 0xf3, 0xf8, 0xd4, 0xf3, 0xf8, 0x77, 0xcf, 0xe3, 0x1f, 0x67, 0x1e, 0x9d, 0xce,
	0x3c, 0xfa, 0x75, 0xe6, 0xd1, 0xa7, 0x17, 0x4a, 0x5b, 0xf7, 0x53, 0x89, 0x46, 0x54, 0x5d, 0xfd,
	0x4a, 0xa3, 0x08, 0x0b, 0xdd, 0x87, 0x7b, 0x5c, 0xf6, 0x5a, 0x8c, 0xfd, 0x46, 0xdf, 0xfc, 0x0d,
	0x00, 0x00, 0xff, 0xff, 0x08, 0xce, 0x90, 0xe8, 0x51, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgServiceClient is the client API for MsgService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgServiceClient interface {
	MsgRegisterNode(ctx context.Context, in *MsgRegisterNodeRequest, opts ...grpc.CallOption) (*MsgRegisterNodeResponse, error)
}

type msgServiceClient struct {
	cc grpc1.ClientConn
}

func NewMsgServiceClient(cc grpc1.ClientConn) MsgServiceClient {
	return &msgServiceClient{cc}
}

func (c *msgServiceClient) MsgRegisterNode(ctx context.Context, in *MsgRegisterNodeRequest, opts ...grpc.CallOption) (*MsgRegisterNodeResponse, error) {
	out := new(MsgRegisterNodeResponse)
	err := c.cc.Invoke(ctx, "/imhub.node.v1.MsgService/MsgRegisterNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServiceServer is the server API for MsgService service.
type MsgServiceServer interface {
	MsgRegisterNode(context.Context, *MsgRegisterNodeRequest) (*MsgRegisterNodeResponse, error)
}

// UnimplementedMsgServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServiceServer struct {
}

func (*UnimplementedMsgServiceServer) MsgRegisterNode(ctx context.Context, req *MsgRegisterNodeRequest) (*MsgRegisterNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MsgRegisterNode not implemented")
}

func RegisterMsgServiceServer(s grpc1.Server, srv MsgServiceServer) {
	s.RegisterService(&_MsgService_serviceDesc, srv)
}

func _MsgService_MsgRegisterNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRegisterNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).MsgRegisterNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/imhub.node.v1.MsgService/MsgRegisterNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).MsgRegisterNode(ctx, req.(*MsgRegisterNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var MsgService_serviceDesc = _MsgService_serviceDesc
var _MsgService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "imhub.node.v1.MsgService",
	HandlerType: (*MsgServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MsgRegisterNode",
			Handler:    _MsgService_MsgRegisterNode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "imhub/node/v1/tx.proto",
}

func (m *MsgRegisterNodeRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRegisterNodeRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRegisterNodeRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PubKey) > 0 {
		i -= len(m.PubKey)
		copy(dAtA[i:], m.PubKey)
		i = encodeVarintTx(dAtA, i, uint64(len(m.PubKey)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintTx(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgRegisterNodeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRegisterNodeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRegisterNodeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgRegisterNodeRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.PubKey)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgRegisterNodeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgRegisterNodeRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgRegisterNodeRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRegisterNodeRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PubKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgRegisterNodeResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgRegisterNodeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRegisterNodeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
