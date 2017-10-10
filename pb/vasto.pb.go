// Code generated by protoc-gen-go.
// source: vasto.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	vasto.proto

It has these top-level messages:
	BalanceRequest
	ClientMessage
	Topology
	Ring
	StoreResource
	Location
	StoreHeartbeat
	ClientHeartbeat
	Empty
	Requests
	Responses
	Request
	PutRequest
	PutResponse
	DeleteRequest
	DeleteResponse
	GetRequest
	GetResponse
	RangeGetRequest
	RangeGetResponse
	Response
	KeyValue
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// ////////////////////////////////////////////////
// 1. master received request to balance the data
type BalanceRequest struct {
	DataCenter string `protobuf:"bytes,1,opt,name=data_center,json=dataCenter" json:"data_center,omitempty"`
	StoreGroup string `protobuf:"bytes,2,opt,name=store_group,json=storeGroup" json:"store_group,omitempty"`
	StoreCount uint32 `protobuf:"varint,3,opt,name=store_count,json=storeCount" json:"store_count,omitempty"`
}

func (m *BalanceRequest) Reset()                    { *m = BalanceRequest{} }
func (m *BalanceRequest) String() string            { return proto.CompactTextString(m) }
func (*BalanceRequest) ProtoMessage()               {}
func (*BalanceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BalanceRequest) GetDataCenter() string {
	if m != nil {
		return m.DataCenter
	}
	return ""
}

func (m *BalanceRequest) GetStoreGroup() string {
	if m != nil {
		return m.StoreGroup
	}
	return ""
}

func (m *BalanceRequest) GetStoreCount() uint32 {
	if m != nil {
		return m.StoreCount
	}
	return 0
}

type ClientMessage struct {
	Store *StoreResource `protobuf:"bytes,1,opt,name=store" json:"store,omitempty"`
}

func (m *ClientMessage) Reset()                    { *m = ClientMessage{} }
func (m *ClientMessage) String() string            { return proto.CompactTextString(m) }
func (*ClientMessage) ProtoMessage()               {}
func (*ClientMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ClientMessage) GetStore() *StoreResource {
	if m != nil {
		return m.Store
	}
	return nil
}

type Topology struct {
	Rings []*Ring `protobuf:"bytes,1,rep,name=rings" json:"rings,omitempty"`
}

func (m *Topology) Reset()                    { *m = Topology{} }
func (m *Topology) String() string            { return proto.CompactTextString(m) }
func (*Topology) ProtoMessage()               {}
func (*Topology) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Topology) GetRings() []*Ring {
	if m != nil {
		return m.Rings
	}
	return nil
}

type Ring struct {
	DataCenter string           `protobuf:"bytes,1,opt,name=data_center,json=dataCenter" json:"data_center,omitempty"`
	Stores     []*StoreResource `protobuf:"bytes,2,rep,name=stores" json:"stores,omitempty"`
}

func (m *Ring) Reset()                    { *m = Ring{} }
func (m *Ring) String() string            { return proto.CompactTextString(m) }
func (*Ring) ProtoMessage()               {}
func (*Ring) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Ring) GetDataCenter() string {
	if m != nil {
		return m.DataCenter
	}
	return ""
}

func (m *Ring) GetStores() []*StoreResource {
	if m != nil {
		return m.Stores
	}
	return nil
}

type StoreResource struct {
	Id       int32     `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Location *Location `protobuf:"bytes,2,opt,name=location" json:"location,omitempty"`
}

func (m *StoreResource) Reset()                    { *m = StoreResource{} }
func (m *StoreResource) String() string            { return proto.CompactTextString(m) }
func (*StoreResource) ProtoMessage()               {}
func (*StoreResource) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *StoreResource) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *StoreResource) GetLocation() *Location {
	if m != nil {
		return m.Location
	}
	return nil
}

type Location struct {
	DataCenter string `protobuf:"bytes,1,opt,name=data_center,json=dataCenter" json:"data_center,omitempty"`
	Server     string `protobuf:"bytes,2,opt,name=server" json:"server,omitempty"`
	Port       int32  `protobuf:"varint,3,opt,name=port" json:"port,omitempty"`
}

func (m *Location) Reset()                    { *m = Location{} }
func (m *Location) String() string            { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()               {}
func (*Location) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Location) GetDataCenter() string {
	if m != nil {
		return m.DataCenter
	}
	return ""
}

func (m *Location) GetServer() string {
	if m != nil {
		return m.Server
	}
	return ""
}

func (m *Location) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

// ////////////////////////////////////////////////
type StoreHeartbeat struct {
	DataCenter string         `protobuf:"bytes,1,opt,name=data_center,json=dataCenter" json:"data_center,omitempty"`
	Store      *StoreResource `protobuf:"bytes,2,opt,name=store" json:"store,omitempty"`
}

func (m *StoreHeartbeat) Reset()                    { *m = StoreHeartbeat{} }
func (m *StoreHeartbeat) String() string            { return proto.CompactTextString(m) }
func (*StoreHeartbeat) ProtoMessage()               {}
func (*StoreHeartbeat) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *StoreHeartbeat) GetDataCenter() string {
	if m != nil {
		return m.DataCenter
	}
	return ""
}

func (m *StoreHeartbeat) GetStore() *StoreResource {
	if m != nil {
		return m.Store
	}
	return nil
}

type ClientHeartbeat struct {
	Location *Location `protobuf:"bytes,1,opt,name=location" json:"location,omitempty"`
}

func (m *ClientHeartbeat) Reset()                    { *m = ClientHeartbeat{} }
func (m *ClientHeartbeat) String() string            { return proto.CompactTextString(m) }
func (*ClientHeartbeat) ProtoMessage()               {}
func (*ClientHeartbeat) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ClientHeartbeat) GetLocation() *Location {
	if m != nil {
		return m.Location
	}
	return nil
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

// ////////////////////////////////////////////////
type Requests struct {
	Requests []*Request `protobuf:"bytes,1,rep,name=requests" json:"requests,omitempty"`
}

func (m *Requests) Reset()                    { *m = Requests{} }
func (m *Requests) String() string            { return proto.CompactTextString(m) }
func (*Requests) ProtoMessage()               {}
func (*Requests) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *Requests) GetRequests() []*Request {
	if m != nil {
		return m.Requests
	}
	return nil
}

type Responses struct {
	Responses []*Response `protobuf:"bytes,1,rep,name=responses" json:"responses,omitempty"`
}

func (m *Responses) Reset()                    { *m = Responses{} }
func (m *Responses) String() string            { return proto.CompactTextString(m) }
func (*Responses) ProtoMessage()               {}
func (*Responses) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *Responses) GetResponses() []*Response {
	if m != nil {
		return m.Responses
	}
	return nil
}

type Request struct {
	Put      *PutRequest      `protobuf:"bytes,1,opt,name=put" json:"put,omitempty"`
	Get      *GetRequest      `protobuf:"bytes,2,opt,name=get" json:"get,omitempty"`
	RangeGet *RangeGetRequest `protobuf:"bytes,3,opt,name=range_get,json=rangeGet" json:"range_get,omitempty"`
	Delete   *DeleteRequest   `protobuf:"bytes,4,opt,name=delete" json:"delete,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *Request) GetPut() *PutRequest {
	if m != nil {
		return m.Put
	}
	return nil
}

func (m *Request) GetGet() *GetRequest {
	if m != nil {
		return m.Get
	}
	return nil
}

func (m *Request) GetRangeGet() *RangeGetRequest {
	if m != nil {
		return m.RangeGet
	}
	return nil
}

func (m *Request) GetDelete() *DeleteRequest {
	if m != nil {
		return m.Delete
	}
	return nil
}

type PutRequest struct {
	KeyValue    *KeyValue `protobuf:"bytes,1,opt,name=key_value,json=keyValue" json:"key_value,omitempty"`
	TimestampNs int64     `protobuf:"varint,2,opt,name=timestamp_ns,json=timestampNs" json:"timestamp_ns,omitempty"`
	TtlMs       uint32    `protobuf:"varint,3,opt,name=ttl_ms,json=ttlMs" json:"ttl_ms,omitempty"`
}

func (m *PutRequest) Reset()                    { *m = PutRequest{} }
func (m *PutRequest) String() string            { return proto.CompactTextString(m) }
func (*PutRequest) ProtoMessage()               {}
func (*PutRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *PutRequest) GetKeyValue() *KeyValue {
	if m != nil {
		return m.KeyValue
	}
	return nil
}

func (m *PutRequest) GetTimestampNs() int64 {
	if m != nil {
		return m.TimestampNs
	}
	return 0
}

func (m *PutRequest) GetTtlMs() uint32 {
	if m != nil {
		return m.TtlMs
	}
	return 0
}

type PutResponse struct {
	Ok     bool   `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
	Status string `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
}

func (m *PutResponse) Reset()                    { *m = PutResponse{} }
func (m *PutResponse) String() string            { return proto.CompactTextString(m) }
func (*PutResponse) ProtoMessage()               {}
func (*PutResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *PutResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *PutResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type DeleteRequest struct {
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *DeleteRequest) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

type DeleteResponse struct {
	Ok     bool   `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
	Status string `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
}

func (m *DeleteResponse) Reset()                    { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()               {}
func (*DeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *DeleteResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *DeleteResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type GetRequest struct {
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *GetRequest) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

type GetResponse struct {
	Ok       bool      `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
	KeyValue *KeyValue `protobuf:"bytes,2,opt,name=key_value,json=keyValue" json:"key_value,omitempty"`
}

func (m *GetResponse) Reset()                    { *m = GetResponse{} }
func (m *GetResponse) String() string            { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()               {}
func (*GetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *GetResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *GetResponse) GetKeyValue() *KeyValue {
	if m != nil {
		return m.KeyValue
	}
	return nil
}

type RangeGetRequest struct {
	Prefix   []byte `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	StartKey []byte `protobuf:"bytes,2,opt,name=start_key,json=startKey,proto3" json:"start_key,omitempty"`
	Limit    uint32 `protobuf:"varint,3,opt,name=limit" json:"limit,omitempty"`
}

func (m *RangeGetRequest) Reset()                    { *m = RangeGetRequest{} }
func (m *RangeGetRequest) String() string            { return proto.CompactTextString(m) }
func (*RangeGetRequest) ProtoMessage()               {}
func (*RangeGetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

func (m *RangeGetRequest) GetPrefix() []byte {
	if m != nil {
		return m.Prefix
	}
	return nil
}

func (m *RangeGetRequest) GetStartKey() []byte {
	if m != nil {
		return m.StartKey
	}
	return nil
}

func (m *RangeGetRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type RangeGetResponse struct {
	KeyValues []*KeyValue `protobuf:"bytes,1,rep,name=key_values,json=keyValues" json:"key_values,omitempty"`
}

func (m *RangeGetResponse) Reset()                    { *m = RangeGetResponse{} }
func (m *RangeGetResponse) String() string            { return proto.CompactTextString(m) }
func (*RangeGetResponse) ProtoMessage()               {}
func (*RangeGetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

func (m *RangeGetResponse) GetKeyValues() []*KeyValue {
	if m != nil {
		return m.KeyValues
	}
	return nil
}

type Response struct {
	Put      *PutResponse      `protobuf:"bytes,1,opt,name=put" json:"put,omitempty"`
	Get      *GetResponse      `protobuf:"bytes,2,opt,name=get" json:"get,omitempty"`
	RangeGet *RangeGetResponse `protobuf:"bytes,3,opt,name=range_get,json=rangeGet" json:"range_get,omitempty"`
	Delete   *DeleteResponse   `protobuf:"bytes,4,opt,name=delete" json:"delete,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{20} }

func (m *Response) GetPut() *PutResponse {
	if m != nil {
		return m.Put
	}
	return nil
}

func (m *Response) GetGet() *GetResponse {
	if m != nil {
		return m.Get
	}
	return nil
}

func (m *Response) GetRangeGet() *RangeGetResponse {
	if m != nil {
		return m.RangeGet
	}
	return nil
}

func (m *Response) GetDelete() *DeleteResponse {
	if m != nil {
		return m.Delete
	}
	return nil
}

type KeyValue struct {
	Key   []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *KeyValue) Reset()                    { *m = KeyValue{} }
func (m *KeyValue) String() string            { return proto.CompactTextString(m) }
func (*KeyValue) ProtoMessage()               {}
func (*KeyValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{21} }

func (m *KeyValue) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *KeyValue) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*BalanceRequest)(nil), "pb.BalanceRequest")
	proto.RegisterType((*ClientMessage)(nil), "pb.ClientMessage")
	proto.RegisterType((*Topology)(nil), "pb.Topology")
	proto.RegisterType((*Ring)(nil), "pb.Ring")
	proto.RegisterType((*StoreResource)(nil), "pb.StoreResource")
	proto.RegisterType((*Location)(nil), "pb.Location")
	proto.RegisterType((*StoreHeartbeat)(nil), "pb.StoreHeartbeat")
	proto.RegisterType((*ClientHeartbeat)(nil), "pb.ClientHeartbeat")
	proto.RegisterType((*Empty)(nil), "pb.Empty")
	proto.RegisterType((*Requests)(nil), "pb.Requests")
	proto.RegisterType((*Responses)(nil), "pb.Responses")
	proto.RegisterType((*Request)(nil), "pb.Request")
	proto.RegisterType((*PutRequest)(nil), "pb.PutRequest")
	proto.RegisterType((*PutResponse)(nil), "pb.PutResponse")
	proto.RegisterType((*DeleteRequest)(nil), "pb.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "pb.DeleteResponse")
	proto.RegisterType((*GetRequest)(nil), "pb.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "pb.GetResponse")
	proto.RegisterType((*RangeGetRequest)(nil), "pb.RangeGetRequest")
	proto.RegisterType((*RangeGetResponse)(nil), "pb.RangeGetResponse")
	proto.RegisterType((*Response)(nil), "pb.Response")
	proto.RegisterType((*KeyValue)(nil), "pb.KeyValue")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for VastoMaster service

type VastoMasterClient interface {
	RegisterStore(ctx context.Context, opts ...grpc.CallOption) (VastoMaster_RegisterStoreClient, error)
	RegisterClient(ctx context.Context, opts ...grpc.CallOption) (VastoMaster_RegisterClientClient, error)
}

type vastoMasterClient struct {
	cc *grpc.ClientConn
}

func NewVastoMasterClient(cc *grpc.ClientConn) VastoMasterClient {
	return &vastoMasterClient{cc}
}

func (c *vastoMasterClient) RegisterStore(ctx context.Context, opts ...grpc.CallOption) (VastoMaster_RegisterStoreClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_VastoMaster_serviceDesc.Streams[0], c.cc, "/pb.VastoMaster/RegisterStore", opts...)
	if err != nil {
		return nil, err
	}
	x := &vastoMasterRegisterStoreClient{stream}
	return x, nil
}

type VastoMaster_RegisterStoreClient interface {
	Send(*StoreHeartbeat) error
	Recv() (*Ring, error)
	grpc.ClientStream
}

type vastoMasterRegisterStoreClient struct {
	grpc.ClientStream
}

func (x *vastoMasterRegisterStoreClient) Send(m *StoreHeartbeat) error {
	return x.ClientStream.SendMsg(m)
}

func (x *vastoMasterRegisterStoreClient) Recv() (*Ring, error) {
	m := new(Ring)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *vastoMasterClient) RegisterClient(ctx context.Context, opts ...grpc.CallOption) (VastoMaster_RegisterClientClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_VastoMaster_serviceDesc.Streams[1], c.cc, "/pb.VastoMaster/RegisterClient", opts...)
	if err != nil {
		return nil, err
	}
	x := &vastoMasterRegisterClientClient{stream}
	return x, nil
}

type VastoMaster_RegisterClientClient interface {
	Send(*ClientHeartbeat) error
	Recv() (*ClientMessage, error)
	grpc.ClientStream
}

type vastoMasterRegisterClientClient struct {
	grpc.ClientStream
}

func (x *vastoMasterRegisterClientClient) Send(m *ClientHeartbeat) error {
	return x.ClientStream.SendMsg(m)
}

func (x *vastoMasterRegisterClientClient) Recv() (*ClientMessage, error) {
	m := new(ClientMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for VastoMaster service

type VastoMasterServer interface {
	RegisterStore(VastoMaster_RegisterStoreServer) error
	RegisterClient(VastoMaster_RegisterClientServer) error
}

func RegisterVastoMasterServer(s *grpc.Server, srv VastoMasterServer) {
	s.RegisterService(&_VastoMaster_serviceDesc, srv)
}

func _VastoMaster_RegisterStore_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(VastoMasterServer).RegisterStore(&vastoMasterRegisterStoreServer{stream})
}

type VastoMaster_RegisterStoreServer interface {
	Send(*Ring) error
	Recv() (*StoreHeartbeat, error)
	grpc.ServerStream
}

type vastoMasterRegisterStoreServer struct {
	grpc.ServerStream
}

func (x *vastoMasterRegisterStoreServer) Send(m *Ring) error {
	return x.ServerStream.SendMsg(m)
}

func (x *vastoMasterRegisterStoreServer) Recv() (*StoreHeartbeat, error) {
	m := new(StoreHeartbeat)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _VastoMaster_RegisterClient_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(VastoMasterServer).RegisterClient(&vastoMasterRegisterClientServer{stream})
}

type VastoMaster_RegisterClientServer interface {
	Send(*ClientMessage) error
	Recv() (*ClientHeartbeat, error)
	grpc.ServerStream
}

type vastoMasterRegisterClientServer struct {
	grpc.ServerStream
}

func (x *vastoMasterRegisterClientServer) Send(m *ClientMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *vastoMasterRegisterClientServer) Recv() (*ClientHeartbeat, error) {
	m := new(ClientHeartbeat)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _VastoMaster_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.VastoMaster",
	HandlerType: (*VastoMasterServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RegisterStore",
			Handler:       _VastoMaster_RegisterStore_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "RegisterClient",
			Handler:       _VastoMaster_RegisterClient_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "vasto.proto",
}

// Client API for VastoStore service

type VastoStoreClient interface {
	Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error)
}

type vastoStoreClient struct {
	cc *grpc.ClientConn
}

func NewVastoStoreClient(cc *grpc.ClientConn) VastoStoreClient {
	return &vastoStoreClient{cc}
}

func (c *vastoStoreClient) Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error) {
	out := new(PutResponse)
	err := grpc.Invoke(ctx, "/pb.VastoStore/Put", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for VastoStore service

type VastoStoreServer interface {
	Put(context.Context, *PutRequest) (*PutResponse, error)
}

func RegisterVastoStoreServer(s *grpc.Server, srv VastoStoreServer) {
	s.RegisterService(&_VastoStore_serviceDesc, srv)
}

func _VastoStore_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VastoStoreServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.VastoStore/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VastoStoreServer).Put(ctx, req.(*PutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VastoStore_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.VastoStore",
	HandlerType: (*VastoStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Put",
			Handler:    _VastoStore_Put_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vasto.proto",
}

func init() { proto.RegisterFile("vasto.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 794 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x55, 0xdf, 0x6b, 0xe3, 0x46,
	0x10, 0x8e, 0xe4, 0xd8, 0x91, 0x47, 0xfe, 0x91, 0x6c, 0xd3, 0x62, 0x52, 0x48, 0xed, 0x7d, 0x89,
	0x93, 0x42, 0x48, 0x1d, 0xda, 0x06, 0x0a, 0x2d, 0x34, 0x2d, 0x49, 0x49, 0x53, 0xc2, 0xb6, 0xa4,
	0x50, 0x0a, 0x46, 0xb6, 0xa7, 0x42, 0x58, 0xd6, 0x2a, 0xbb, 0xab, 0x50, 0xbf, 0xde, 0x9f, 0x73,
	0x2f, 0xf7, 0x2f, 0x1e, 0xbb, 0x5a, 0x49, 0xb6, 0x93, 0x4b, 0xee, 0xde, 0x76, 0xbf, 0xf9, 0xe6,
	0xc7, 0xce, 0x7c, 0x1a, 0x81, 0xff, 0x18, 0x48, 0xc5, 0x4f, 0x53, 0xc1, 0x15, 0x27, 0x6e, 0x3a,
	0xa1, 0x12, 0x3a, 0x3f, 0x07, 0x71, 0x90, 0x4c, 0x91, 0xe1, 0x43, 0x86, 0x52, 0x91, 0xaf, 0xc0,
	0x9f, 0x05, 0x2a, 0x18, 0x4f, 0x31, 0x51, 0x28, 0x7a, 0x4e, 0xdf, 0x19, 0x36, 0x19, 0x68, 0xe8,
	0xd2, 0x20, 0x9a, 0x20, 0x15, 0x17, 0x38, 0x0e, 0x05, 0xcf, 0xd2, 0x9e, 0x9b, 0x13, 0x0c, 0x74,
	0xa5, 0x91, 0x8a, 0x30, 0xe5, 0x59, 0xa2, 0x7a, 0xb5, 0xbe, 0x33, 0x6c, 0x5b, 0xc2, 0xa5, 0x46,
	0xe8, 0x05, 0xb4, 0x2f, 0xe3, 0x08, 0x13, 0x75, 0x8b, 0x52, 0x06, 0x21, 0x92, 0x23, 0xa8, 0x1b,
	0xb3, 0xc9, 0xe6, 0x8f, 0xf6, 0x4e, 0xd3, 0xc9, 0xe9, 0x9f, 0x1a, 0x60, 0x28, 0x79, 0x26, 0xa6,
	0xc8, 0x72, 0x3b, 0x3d, 0x01, 0xef, 0x2f, 0x9e, 0xf2, 0x98, 0x87, 0x4b, 0x72, 0x08, 0x75, 0x11,
	0x25, 0xa1, 0xec, 0x39, 0xfd, 0xda, 0xd0, 0x1f, 0x79, 0xda, 0x89, 0x45, 0x49, 0xc8, 0x72, 0x98,
	0x32, 0xd8, 0xd6, 0xd7, 0xd7, 0x1f, 0x74, 0x0c, 0x0d, 0x13, 0x5d, 0xf6, 0x5c, 0x13, 0xe9, 0x99,
	0xf4, 0x96, 0x40, 0x7f, 0x83, 0xf6, 0x9a, 0x81, 0x74, 0xc0, 0x8d, 0x66, 0x26, 0x66, 0x9d, 0xb9,
	0xd1, 0x8c, 0x0c, 0xc1, 0x8b, 0xf9, 0x34, 0x50, 0x11, 0x4f, 0x4c, 0x67, 0xfc, 0x51, 0x4b, 0x47,
	0xfb, 0xdd, 0x62, 0xac, 0xb4, 0xd2, 0xbf, 0xc1, 0x2b, 0xd0, 0xd7, 0x4b, 0xfc, 0x02, 0x1a, 0x12,
	0xc5, 0x23, 0x0a, 0xdb, 0x6e, 0x7b, 0x23, 0x04, 0xb6, 0x53, 0x2e, 0xf2, 0x1e, 0xd7, 0x99, 0x39,
	0xd3, 0x7f, 0xa0, 0x63, 0x6a, 0xbc, 0xc6, 0x40, 0xa8, 0x09, 0x06, 0x1f, 0x31, 0xd2, 0xb2, 0xff,
	0xee, 0x2b, 0xfd, 0xff, 0x01, 0xba, 0xf9, 0xe4, 0xaa, 0xe0, 0xab, 0x2f, 0x76, 0x5e, 0x7c, 0xf1,
	0x0e, 0xd4, 0x7f, 0x5d, 0xa4, 0x6a, 0x49, 0xcf, 0xc1, 0xb3, 0x6a, 0x93, 0xe4, 0x08, 0x3c, 0x61,
	0xcf, 0x76, 0x90, 0xbe, 0x19, 0x64, 0x8e, 0xb1, 0xd2, 0x48, 0xbf, 0x87, 0x26, 0x43, 0x99, 0xf2,
	0x44, 0xa2, 0x24, 0x27, 0xd0, 0x14, 0xc5, 0xc5, 0xba, 0xb5, 0x72, 0xb7, 0x1c, 0x64, 0x95, 0x99,
	0xbe, 0x75, 0x60, 0xa7, 0x10, 0x77, 0x1f, 0x6a, 0x69, 0xa6, 0x6c, 0x9d, 0x1d, 0xed, 0x71, 0x97,
	0xa9, 0x22, 0x97, 0x36, 0x69, 0x46, 0x88, 0xca, 0x36, 0xc2, 0x30, 0xae, 0xb0, 0x62, 0x84, 0xa8,
	0xc8, 0x19, 0x34, 0x45, 0x90, 0x84, 0x38, 0xd6, 0xbc, 0x9a, 0xe1, 0x7d, 0x66, 0x72, 0x6b, 0x70,
	0x85, 0xec, 0x09, 0x0b, 0x68, 0x81, 0xcd, 0x30, 0x46, 0x85, 0xbd, 0xed, 0xaa, 0xbf, 0xbf, 0x18,
	0xa4, 0x20, 0x5b, 0x02, 0x7d, 0x00, 0xa8, 0x2a, 0x22, 0xc7, 0xd0, 0x9c, 0xe3, 0x72, 0xfc, 0x18,
	0xc4, 0x19, 0xae, 0x36, 0xf7, 0x06, 0x97, 0xf7, 0x1a, 0x63, 0xde, 0xdc, 0x9e, 0xc8, 0x00, 0x5a,
	0x2a, 0x5a, 0xa0, 0x54, 0xc1, 0x22, 0x1d, 0x27, 0xd2, 0x3c, 0xa0, 0xc6, 0xfc, 0x12, 0xfb, 0x43,
	0x92, 0xcf, 0xa1, 0xa1, 0x54, 0x3c, 0x5e, 0x48, 0xfb, 0x49, 0xd6, 0x95, 0x8a, 0x6f, 0x25, 0xfd,
	0x16, 0x7c, 0x93, 0x32, 0xef, 0x97, 0x56, 0x34, 0x9f, 0x9b, 0x64, 0x1e, 0x73, 0xf9, 0xdc, 0x48,
	0x4f, 0x05, 0x2a, 0x93, 0xa5, 0xf4, 0xcc, 0x8d, 0x0e, 0xa0, 0xbd, 0xf6, 0x04, 0xb2, 0x0b, 0xb5,
	0x39, 0x2e, 0x8d, 0x67, 0x8b, 0xe9, 0x23, 0xbd, 0x80, 0x4e, 0x41, 0xf9, 0xc4, 0xe0, 0x87, 0x00,
	0x55, 0x27, 0x9f, 0x89, 0x7c, 0x0d, 0xbe, 0xb1, 0x7f, 0x20, 0xec, 0x5a, 0xdf, 0xdc, 0x97, 0xfa,
	0x46, 0xff, 0x85, 0xee, 0xc6, 0xe0, 0x74, 0x51, 0xa9, 0xc0, 0xff, 0xa2, 0xff, 0x6d, 0x46, 0x7b,
	0x23, 0x5f, 0x42, 0x53, 0xaa, 0x40, 0xa8, 0xb1, 0x2e, 0xc6, 0x35, 0x26, 0xcf, 0x00, 0x37, 0xb8,
	0x24, 0xfb, 0x50, 0x8f, 0xa3, 0x45, 0x54, 0xac, 0xbb, 0xfc, 0x42, 0x7f, 0x82, 0xdd, 0x2a, 0xba,
	0x2d, 0xf6, 0x6b, 0x80, 0xb2, 0xb8, 0x35, 0xf1, 0x96, 0xd5, 0x35, 0x8b, 0xea, 0x24, 0x7d, 0xe7,
	0xe8, 0x6f, 0xc5, 0x7a, 0x0e, 0x56, 0xd5, 0xdb, 0x2d, 0xd5, 0x6b, 0x25, 0x6f, 0xe4, 0x3b, 0x58,
	0x95, 0x6f, 0xb7, 0x94, 0x6f, 0x41, 0xd1, 0xfa, 0xfd, 0xe6, 0xa9, 0x7e, 0xf7, 0xd7, 0xf5, 0x6b,
	0xd9, 0x95, 0x80, 0x4f, 0x36, 0x04, 0x4c, 0x56, 0x05, 0x6c, 0xd9, 0x85, 0x82, 0x47, 0xe0, 0x15,
	0x0f, 0x79, 0x3a, 0x38, 0xdd, 0xa6, 0x6a, 0x2a, 0x2d, 0x96, 0x5f, 0x46, 0x6f, 0x1c, 0xf0, 0xef,
	0xf5, 0x9f, 0xe9, 0x36, 0x90, 0x7a, 0x1f, 0x9d, 0x43, 0x9b, 0x61, 0x18, 0xe9, 0xb3, 0x59, 0x43,
	0x84, 0x94, 0x1b, 0xa9, 0x5c, 0x3c, 0x07, 0xe5, 0xc2, 0xa7, 0x5b, 0x43, 0xe7, 0xcc, 0x21, 0x3f,
	0x42, 0xa7, 0x70, 0xca, 0x77, 0x14, 0x31, 0x9f, 0xe5, 0xc6, 0xbe, 0x3a, 0xd8, 0xab, 0x40, 0xfb,
	0xfb, 0xc9, 0xfd, 0x47, 0xdf, 0x01, 0x98, 0x1a, 0xf2, 0x8c, 0x43, 0xa8, 0xdd, 0x65, 0x8a, 0x6c,
	0xec, 0x88, 0x83, 0xcd, 0xae, 0xd3, 0xad, 0x49, 0xc3, 0xfc, 0x4d, 0xcf, 0xdf, 0x07, 0x00, 0x00,
	0xff, 0xff, 0x3a, 0x98, 0x82, 0x8b, 0x5c, 0x07, 0x00, 0x00,
}
