// Code generated by protoc-gen-go. DO NOT EDIT.
// source: peer/chaincode_shim.proto

package peer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ChaincodeMessage_Type int32

const (
	ChaincodeMessage_UNDEFINED           ChaincodeMessage_Type = 0
	ChaincodeMessage_REGISTER            ChaincodeMessage_Type = 1
	ChaincodeMessage_REGISTERED          ChaincodeMessage_Type = 2
	ChaincodeMessage_INIT                ChaincodeMessage_Type = 3
	ChaincodeMessage_READY               ChaincodeMessage_Type = 4
	ChaincodeMessage_TRANSACTION         ChaincodeMessage_Type = 5
	ChaincodeMessage_COMPLETED           ChaincodeMessage_Type = 6
	ChaincodeMessage_ERROR               ChaincodeMessage_Type = 7
	ChaincodeMessage_GET_STATE           ChaincodeMessage_Type = 8
	ChaincodeMessage_PUT_STATE           ChaincodeMessage_Type = 9
	ChaincodeMessage_DEL_STATE           ChaincodeMessage_Type = 10
	ChaincodeMessage_INVOKE_CHAINCODE    ChaincodeMessage_Type = 11
	ChaincodeMessage_RESPONSE            ChaincodeMessage_Type = 13
	ChaincodeMessage_GET_STATE_BY_RANGE  ChaincodeMessage_Type = 14
	ChaincodeMessage_GET_QUERY_RESULT    ChaincodeMessage_Type = 15
	ChaincodeMessage_QUERY_STATE_NEXT    ChaincodeMessage_Type = 16
	ChaincodeMessage_QUERY_STATE_CLOSE   ChaincodeMessage_Type = 17
	ChaincodeMessage_KEEPALIVE           ChaincodeMessage_Type = 18
	ChaincodeMessage_GET_HISTORY_FOR_KEY ChaincodeMessage_Type = 19
	ChaincodeMessage_QUERY_BY_VIEW       ChaincodeMessage_Type = 20
)

var ChaincodeMessage_Type_name = map[int32]string{
	0:  "UNDEFINED",
	1:  "REGISTER",
	2:  "REGISTERED",
	3:  "INIT",
	4:  "READY",
	5:  "TRANSACTION",
	6:  "COMPLETED",
	7:  "ERROR",
	8:  "GET_STATE",
	9:  "PUT_STATE",
	10: "DEL_STATE",
	11: "INVOKE_CHAINCODE",
	13: "RESPONSE",
	14: "GET_STATE_BY_RANGE",
	15: "GET_QUERY_RESULT",
	16: "QUERY_STATE_NEXT",
	17: "QUERY_STATE_CLOSE",
	18: "KEEPALIVE",
	19: "GET_HISTORY_FOR_KEY",
	20: "QUERY_BY_VIEW",
}
var ChaincodeMessage_Type_value = map[string]int32{
	"UNDEFINED":           0,
	"REGISTER":            1,
	"REGISTERED":          2,
	"INIT":                3,
	"READY":               4,
	"TRANSACTION":         5,
	"COMPLETED":           6,
	"ERROR":               7,
	"GET_STATE":           8,
	"PUT_STATE":           9,
	"DEL_STATE":           10,
	"INVOKE_CHAINCODE":    11,
	"RESPONSE":            13,
	"GET_STATE_BY_RANGE":  14,
	"GET_QUERY_RESULT":    15,
	"QUERY_STATE_NEXT":    16,
	"QUERY_STATE_CLOSE":   17,
	"KEEPALIVE":           18,
	"GET_HISTORY_FOR_KEY": 19,
	"QUERY_BY_VIEW":       20,
}

func (x ChaincodeMessage_Type) String() string {
	return proto.EnumName(ChaincodeMessage_Type_name, int32(x))
}
func (ChaincodeMessage_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 0} }

type ChaincodeMessage struct {
	Type      ChaincodeMessage_Type       `protobuf:"varint,1,opt,name=type,enum=protos.ChaincodeMessage_Type" json:"type,omitempty"`
	Timestamp *google_protobuf1.Timestamp `protobuf:"bytes,2,opt,name=timestamp" json:"timestamp,omitempty"`
	Payload   []byte                      `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	Txid      string                      `protobuf:"bytes,4,opt,name=txid" json:"txid,omitempty"`
	Proposal  *SignedProposal             `protobuf:"bytes,5,opt,name=proposal" json:"proposal,omitempty"`
	// event emmited by chaincode. Used only with Init or Invoke.
	// This event is then stored (currently)
	// with Block.NonHashData.TransactionResult
	ChaincodeEvent *ChaincodeEvent `protobuf:"bytes,6,opt,name=chaincode_event,json=chaincodeEvent" json:"chaincode_event,omitempty"`
}

func (m *ChaincodeMessage) Reset()                    { *m = ChaincodeMessage{} }
func (m *ChaincodeMessage) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeMessage) ProtoMessage()               {}
func (*ChaincodeMessage) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *ChaincodeMessage) GetType() ChaincodeMessage_Type {
	if m != nil {
		return m.Type
	}
	return ChaincodeMessage_UNDEFINED
}

func (m *ChaincodeMessage) GetTimestamp() *google_protobuf1.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *ChaincodeMessage) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *ChaincodeMessage) GetTxid() string {
	if m != nil {
		return m.Txid
	}
	return ""
}

func (m *ChaincodeMessage) GetProposal() *SignedProposal {
	if m != nil {
		return m.Proposal
	}
	return nil
}

func (m *ChaincodeMessage) GetChaincodeEvent() *ChaincodeEvent {
	if m != nil {
		return m.ChaincodeEvent
	}
	return nil
}

type PutStateInfo struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *PutStateInfo) Reset()                    { *m = PutStateInfo{} }
func (m *PutStateInfo) String() string            { return proto.CompactTextString(m) }
func (*PutStateInfo) ProtoMessage()               {}
func (*PutStateInfo) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *PutStateInfo) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *PutStateInfo) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type GetStateByRange struct {
	StartKey string `protobuf:"bytes,1,opt,name=startKey" json:"startKey,omitempty"`
	EndKey   string `protobuf:"bytes,2,opt,name=endKey" json:"endKey,omitempty"`
}

func (m *GetStateByRange) Reset()                    { *m = GetStateByRange{} }
func (m *GetStateByRange) String() string            { return proto.CompactTextString(m) }
func (*GetStateByRange) ProtoMessage()               {}
func (*GetStateByRange) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *GetStateByRange) GetStartKey() string {
	if m != nil {
		return m.StartKey
	}
	return ""
}

func (m *GetStateByRange) GetEndKey() string {
	if m != nil {
		return m.EndKey
	}
	return ""
}

type GetQueryResult struct {
	Query string `protobuf:"bytes,1,opt,name=query" json:"query,omitempty"`
}

func (m *GetQueryResult) Reset()                    { *m = GetQueryResult{} }
func (m *GetQueryResult) String() string            { return proto.CompactTextString(m) }
func (*GetQueryResult) ProtoMessage()               {}
func (*GetQueryResult) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *GetQueryResult) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

type GetHistoryForKey struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
}

func (m *GetHistoryForKey) Reset()                    { *m = GetHistoryForKey{} }
func (m *GetHistoryForKey) String() string            { return proto.CompactTextString(m) }
func (*GetHistoryForKey) ProtoMessage()               {}
func (*GetHistoryForKey) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func (m *GetHistoryForKey) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type QueryStateNext struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *QueryStateNext) Reset()                    { *m = QueryStateNext{} }
func (m *QueryStateNext) String() string            { return proto.CompactTextString(m) }
func (*QueryStateNext) ProtoMessage()               {}
func (*QueryStateNext) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{5} }

func (m *QueryStateNext) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type QueryStateClose struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *QueryStateClose) Reset()                    { *m = QueryStateClose{} }
func (m *QueryStateClose) String() string            { return proto.CompactTextString(m) }
func (*QueryStateClose) ProtoMessage()               {}
func (*QueryStateClose) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{6} }

func (m *QueryStateClose) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type QueryResultBytes struct {
	ResultBytes []byte `protobuf:"bytes,1,opt,name=resultBytes,proto3" json:"resultBytes,omitempty"`
}

func (m *QueryResultBytes) Reset()                    { *m = QueryResultBytes{} }
func (m *QueryResultBytes) String() string            { return proto.CompactTextString(m) }
func (*QueryResultBytes) ProtoMessage()               {}
func (*QueryResultBytes) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{7} }

func (m *QueryResultBytes) GetResultBytes() []byte {
	if m != nil {
		return m.ResultBytes
	}
	return nil
}

type QueryResponse struct {
	Results []*QueryResultBytes `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
	HasMore bool                `protobuf:"varint,2,opt,name=has_more,json=hasMore" json:"has_more,omitempty"`
	Id      string              `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
}

func (m *QueryResponse) Reset()                    { *m = QueryResponse{} }
func (m *QueryResponse) String() string            { return proto.CompactTextString(m) }
func (*QueryResponse) ProtoMessage()               {}
func (*QueryResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{8} }

func (m *QueryResponse) GetResults() []*QueryResultBytes {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *QueryResponse) GetHasMore() bool {
	if m != nil {
		return m.HasMore
	}
	return false
}

func (m *QueryResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ViewOpt struct {
	Opt []byte `protobuf:"bytes,1,opt,name=opt,proto3" json:"opt,omitempty"`
}

func (m *ViewOpt) Reset()                    { *m = ViewOpt{} }
func (m *ViewOpt) String() string            { return proto.CompactTextString(m) }
func (*ViewOpt) ProtoMessage()               {}
func (*ViewOpt) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{9} }

func (m *ViewOpt) GetOpt() []byte {
	if m != nil {
		return m.Opt
	}
	return nil
}


func init() {
	proto.RegisterType((*ChaincodeMessage)(nil), "protos.ChaincodeMessage")
	proto.RegisterType((*PutStateInfo)(nil), "protos.PutStateInfo")
	proto.RegisterType((*GetStateByRange)(nil), "protos.GetStateByRange")
	proto.RegisterType((*GetQueryResult)(nil), "protos.GetQueryResult")
	proto.RegisterType((*GetHistoryForKey)(nil), "protos.GetHistoryForKey")
	proto.RegisterType((*QueryStateNext)(nil), "protos.QueryStateNext")
	proto.RegisterType((*QueryStateClose)(nil), "protos.QueryStateClose")
	proto.RegisterType((*QueryResultBytes)(nil), "protos.QueryResultBytes")
	proto.RegisterType((*QueryResponse)(nil), "protos.QueryResponse")
	proto.RegisterType((*ViewOpt)(nil), "protos.ViewOpt")
	proto.RegisterEnum("protos.ChaincodeMessage_Type", ChaincodeMessage_Type_name, ChaincodeMessage_Type_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ChaincodeSupport service

type ChaincodeSupportClient interface {
	Register(ctx context.Context, opts ...grpc.CallOption) (ChaincodeSupport_RegisterClient, error)
}

type chaincodeSupportClient struct {
	cc *grpc.ClientConn
}

func NewChaincodeSupportClient(cc *grpc.ClientConn) ChaincodeSupportClient {
	return &chaincodeSupportClient{cc}
}

func (c *chaincodeSupportClient) Register(ctx context.Context, opts ...grpc.CallOption) (ChaincodeSupport_RegisterClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ChaincodeSupport_serviceDesc.Streams[0], c.cc, "/protos.ChaincodeSupport/Register", opts...)
	if err != nil {
		return nil, err
	}
	x := &chaincodeSupportRegisterClient{stream}
	return x, nil
}

type ChaincodeSupport_RegisterClient interface {
	Send(*ChaincodeMessage) error
	Recv() (*ChaincodeMessage, error)
	grpc.ClientStream
}

type chaincodeSupportRegisterClient struct {
	grpc.ClientStream
}

func (x *chaincodeSupportRegisterClient) Send(m *ChaincodeMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chaincodeSupportRegisterClient) Recv() (*ChaincodeMessage, error) {
	m := new(ChaincodeMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for ChaincodeSupport service

type ChaincodeSupportServer interface {
	Register(ChaincodeSupport_RegisterServer) error
}

func RegisterChaincodeSupportServer(s *grpc.Server, srv ChaincodeSupportServer) {
	s.RegisterService(&_ChaincodeSupport_serviceDesc, srv)
}

func _ChaincodeSupport_Register_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChaincodeSupportServer).Register(&chaincodeSupportRegisterServer{stream})
}

type ChaincodeSupport_RegisterServer interface {
	Send(*ChaincodeMessage) error
	Recv() (*ChaincodeMessage, error)
	grpc.ServerStream
}

type chaincodeSupportRegisterServer struct {
	grpc.ServerStream
}

func (x *chaincodeSupportRegisterServer) Send(m *ChaincodeMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chaincodeSupportRegisterServer) Recv() (*ChaincodeMessage, error) {
	m := new(ChaincodeMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ChaincodeSupport_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.ChaincodeSupport",
	HandlerType: (*ChaincodeSupportServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Register",
			Handler:       _ChaincodeSupport_Register_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "peer/chaincode_shim.proto",
}

func init() { proto.RegisterFile("peer/chaincode_shim.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 771 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x94, 0x6f, 0x6f, 0xe2, 0x46,
	0x10, 0xc6, 0x8f, 0x3f, 0x49, 0x60, 0x20, 0xb0, 0xb7, 0xb9, 0xa6, 0x3e, 0xa4, 0xaa, 0xd4, 0xaa,
	0x2a, 0xfa, 0xc6, 0xb4, 0xb4, 0xaa, 0xfa, 0xae, 0x22, 0xb0, 0x21, 0x56, 0x88, 0xcd, 0xad, 0x9d,
	0xd3, 0xd1, 0x37, 0x96, 0x03, 0x1b, 0x63, 0x15, 0x58, 0x77, 0x77, 0x39, 0x9d, 0x3f, 0x61, 0xbf,
	0x4c, 0x3f, 0xc4, 0x69, 0x6d, 0x4c, 0xb8, 0x9c, 0xf2, 0xca, 0x7e, 0x66, 0x7e, 0xf3, 0xcc, 0xec,
	0x6a, 0x35, 0xf0, 0x36, 0x61, 0x4c, 0xf4, 0x17, 0xab, 0x30, 0xde, 0x2e, 0xf8, 0x92, 0x05, 0x72,
	0x15, 0x6f, 0xac, 0x44, 0x70, 0xc5, 0xf1, 0x69, 0xf6, 0x91, 0x9d, 0xce, 0x33, 0x84, 0x7d, 0x64,
	0x5b, 0x95, 0x33, 0x9d, 0x8b, 0x2c, 0x97, 0x08, 0x9e, 0x70, 0x19, 0xae, 0xf7, 0xc1, 0xef, 0x23,
	0xce, 0xa3, 0x35, 0xeb, 0x67, 0xea, 0x61, 0xf7, 0xd8, 0x57, 0xf1, 0x86, 0x49, 0x15, 0x6e, 0x92,
	0x1c, 0x30, 0xff, 0xaf, 0x02, 0x1a, 0x15, 0x7e, 0x77, 0x4c, 0xca, 0x30, 0x62, 0xf8, 0x57, 0xa8,
	0xaa, 0x34, 0x61, 0x46, 0xa9, 0x5b, 0xea, 0xb5, 0x06, 0xdf, 0xe5, 0xa8, 0xb4, 0x9e, 0x73, 0x96,
	0x9f, 0x26, 0x8c, 0x66, 0x28, 0xfe, 0x13, 0xea, 0x07, 0x6b, 0xa3, 0xdc, 0x2d, 0xf5, 0x1a, 0x83,
	0x8e, 0x95, 0x37, 0xb7, 0x8a, 0xe6, 0x96, 0x5f, 0x10, 0xf4, 0x09, 0xc6, 0x06, 0x9c, 0x25, 0x61,
	0xba, 0xe6, 0xe1, 0xd2, 0xa8, 0x74, 0x4b, 0xbd, 0x26, 0x2d, 0x24, 0xc6, 0x50, 0x55, 0x9f, 0xe2,
	0xa5, 0x51, 0xed, 0x96, 0x7a, 0x75, 0x9a, 0xfd, 0xe3, 0x01, 0xd4, 0x8a, 0x23, 0x1a, 0x27, 0x59,
	0x9b, 0xcb, 0x62, 0x3c, 0x2f, 0x8e, 0xb6, 0x6c, 0x39, 0xdb, 0x67, 0xe9, 0x81, 0xc3, 0x7f, 0x41,
	0xfb, 0xd9, 0x95, 0x19, 0xa7, 0x5f, 0x96, 0x1e, 0x4e, 0x46, 0x74, 0x96, 0xb6, 0x16, 0x5f, 0x68,
	0xf3, 0xbf, 0x32, 0x54, 0xf5, 0x59, 0xf1, 0x39, 0xd4, 0xef, 0x9d, 0x31, 0xb9, 0xb6, 0x1d, 0x32,
	0x46, 0xaf, 0x70, 0x13, 0x6a, 0x94, 0x4c, 0x6c, 0xcf, 0x27, 0x14, 0x95, 0x70, 0x0b, 0xa0, 0x50,
	0x64, 0x8c, 0xca, 0xb8, 0x06, 0x55, 0xdb, 0xb1, 0x7d, 0x54, 0xc1, 0x75, 0x38, 0xa1, 0x64, 0x38,
	0x9e, 0xa3, 0x2a, 0x6e, 0x43, 0xc3, 0xa7, 0x43, 0xc7, 0x1b, 0x8e, 0x7c, 0xdb, 0x75, 0xd0, 0x89,
	0xb6, 0x1c, 0xb9, 0x77, 0xb3, 0x29, 0xf1, 0xc9, 0x18, 0x9d, 0x6a, 0x94, 0x50, 0xea, 0x52, 0x74,
	0xa6, 0x33, 0x13, 0xe2, 0x07, 0x9e, 0x3f, 0xf4, 0x09, 0xaa, 0x69, 0x39, 0xbb, 0x2f, 0x64, 0x5d,
	0xcb, 0x31, 0x99, 0xee, 0x25, 0xe0, 0x37, 0x80, 0x6c, 0xe7, 0xbd, 0x7b, 0x4b, 0x82, 0xd1, 0xcd,
	0xd0, 0x76, 0x46, 0xee, 0x98, 0xa0, 0x46, 0x3e, 0xa0, 0x37, 0x73, 0x1d, 0x8f, 0xa0, 0x73, 0x7c,
	0x09, 0xf8, 0x60, 0x18, 0x5c, 0xcd, 0x03, 0x3a, 0x74, 0x26, 0x04, 0xb5, 0x74, 0xad, 0x8e, 0xbf,
	0xbb, 0x27, 0x74, 0x1e, 0x50, 0xe2, 0xdd, 0x4f, 0x7d, 0xd4, 0xd6, 0xd1, 0x3c, 0x92, 0xf3, 0x0e,
	0xf9, 0xe0, 0x23, 0x84, 0xbf, 0x81, 0xd7, 0xc7, 0xd1, 0xd1, 0xd4, 0xf5, 0x08, 0x7a, 0xad, 0xa7,
	0xb9, 0x25, 0x64, 0x36, 0x9c, 0xda, 0xef, 0x09, 0xc2, 0xf8, 0x5b, 0xb8, 0xd0, 0x8e, 0x37, 0xb6,
	0xe7, 0xbb, 0x74, 0x1e, 0x5c, 0xbb, 0x34, 0xb8, 0x25, 0x73, 0x74, 0x61, 0xfe, 0x01, 0xcd, 0xd9,
	0x4e, 0x79, 0x2a, 0x54, 0xcc, 0xde, 0x3e, 0x72, 0x8c, 0xa0, 0xf2, 0x0f, 0x4b, 0xb3, 0x87, 0x56,
	0xa7, 0xfa, 0x17, 0xbf, 0x81, 0x93, 0x8f, 0xe1, 0x7a, 0xc7, 0xb2, 0x47, 0xd4, 0xa4, 0xb9, 0x30,
	0x09, 0xb4, 0x27, 0x2c, 0xaf, 0xbb, 0x4a, 0x69, 0xb8, 0x8d, 0x18, 0xee, 0x40, 0x4d, 0xaa, 0x50,
	0xa8, 0xdb, 0x43, 0xfd, 0x41, 0xe3, 0x4b, 0x38, 0x65, 0xdb, 0xa5, 0xce, 0x94, 0xb3, 0xcc, 0x5e,
	0x99, 0x3f, 0x41, 0x6b, 0xc2, 0xd4, 0xbb, 0x1d, 0x13, 0x29, 0x65, 0x72, 0xb7, 0x56, 0xba, 0xdd,
	0xbf, 0x5a, 0xee, 0x2d, 0x72, 0x61, 0xfe, 0x08, 0x68, 0xc2, 0xd4, 0x4d, 0x2c, 0x15, 0x17, 0xe9,
	0x35, 0x17, 0xda, 0xf3, 0xab, 0x51, 0xcd, 0x2e, 0xb4, 0x32, 0xab, 0x6c, 0x2c, 0x87, 0x7d, 0x52,
	0xb8, 0x05, 0xe5, 0x78, 0xb9, 0x47, 0xca, 0xf1, 0xd2, 0xfc, 0x01, 0xda, 0x4f, 0xc4, 0x68, 0xcd,
	0x25, 0xfb, 0x0a, 0xf9, 0x1d, 0xd0, 0xd1, 0x3c, 0x57, 0xa9, 0x62, 0x12, 0x77, 0xa1, 0x21, 0x9e,
	0x64, 0x06, 0x37, 0xe9, 0x71, 0xc8, 0xdc, 0xc2, 0x79, 0x51, 0x95, 0xf0, 0xad, 0x64, 0x78, 0x00,
	0x67, 0x79, 0x5e, 0xe3, 0x95, 0x5e, 0x63, 0x60, 0x14, 0x6f, 0xfb, 0xb9, 0x3b, 0x2d, 0x40, 0xfc,
	0x16, 0x6a, 0xab, 0x50, 0x06, 0x1b, 0x2e, 0xf2, 0xdb, 0xae, 0xd1, 0xb3, 0x55, 0x28, 0xef, 0xb8,
	0x28, 0xa6, 0xac, 0x14, 0x53, 0x0e, 0x3e, 0x1c, 0x6d, 0x09, 0x6f, 0x97, 0x24, 0x5c, 0x28, 0x3c,
	0x86, 0x1a, 0x65, 0x51, 0x2c, 0x15, 0x13, 0xd8, 0x78, 0x69, 0x47, 0x74, 0x5e, 0xcc, 0x98, 0xaf,
	0x7a, 0xa5, 0x5f, 0x4a, 0x57, 0x2e, 0x98, 0x5c, 0x44, 0xd6, 0x2a, 0x4d, 0x98, 0x58, 0xb3, 0x65,
	0xc4, 0x84, 0xf5, 0x18, 0x3e, 0x88, 0x78, 0x51, 0xd4, 0xe9, 0xb5, 0xf6, 0xf7, 0xcf, 0x51, 0xac,
	0x56, 0xbb, 0x07, 0x6b, 0xc1, 0x37, 0xfd, 0x23, 0xb4, 0x9f, 0xa3, 0xf9, 0x7a, 0x93, 0x7d, 0x8d,
	0x3e, 0xe4, 0xbb, 0xf2, 0xb7, 0xcf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xad, 0x9f, 0x80, 0x68, 0x4f,
	0x05, 0x00, 0x00,
}
