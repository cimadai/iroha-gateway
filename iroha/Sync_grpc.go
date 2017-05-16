//Generated by gRPC Go plugin
//If you make any local changes, they will be lost
//source: endpoint

package iroha

import "github.com/google/flatbuffers/go"

import (
  context "golang.org/x/net/context"
  grpc "google.golang.org/grpc"
)

// Client API for Sync service
type SyncClient interface{
  CheckHash(ctx context.Context, in *flatbuffers.Builder, 
  	opts... grpc.CallOption) (* CheckHashResponse, error)  
  GetPeers(ctx context.Context, in *flatbuffers.Builder, 
  	opts... grpc.CallOption) (* PeersResponse, error)  
  GetTransactions(ctx context.Context, in *flatbuffers.Builder, 
  	opts... grpc.CallOption) (* TransactionResponse, error)  
  FetchStreamTransaction(ctx context.Context, in *flatbuffers.Builder, 
  	opts... grpc.CallOption) (Sync_FetchStreamTransactionClient, error)  
}

type syncClient struct {
  cc *grpc.ClientConn
}

func NewSyncClient(cc *grpc.ClientConn) SyncClient {
  return &syncClient{cc}
}

func (c *syncClient) CheckHash(ctx context.Context, in *flatbuffers.Builder, 
	opts... grpc.CallOption) (* CheckHashResponse, error) {
  out := new(CheckHashResponse)
  err := grpc.Invoke(ctx, "/iroha.Sync/CheckHash", in, out, c.cc, opts...)
  if err != nil { return nil, err }
  return out, nil
}

func (c *syncClient) GetPeers(ctx context.Context, in *flatbuffers.Builder, 
	opts... grpc.CallOption) (* PeersResponse, error) {
  out := new(PeersResponse)
  err := grpc.Invoke(ctx, "/iroha.Sync/GetPeers", in, out, c.cc, opts...)
  if err != nil { return nil, err }
  return out, nil
}

func (c *syncClient) GetTransactions(ctx context.Context, in *flatbuffers.Builder, 
	opts... grpc.CallOption) (* TransactionResponse, error) {
  out := new(TransactionResponse)
  err := grpc.Invoke(ctx, "/iroha.Sync/GetTransactions", in, out, c.cc, opts...)
  if err != nil { return nil, err }
  return out, nil
}

func (c *syncClient) FetchStreamTransaction(ctx context.Context, in *flatbuffers.Builder, 
	opts... grpc.CallOption) (Sync_FetchStreamTransactionClient, error) {
  stream, err := grpc.NewClientStream(ctx, &_Sync_serviceDesc.Streams[0], c.cc, "/iroha.Sync/FetchStreamTransaction", opts...)
  if err != nil { return nil, err }
  x := &syncFetchStreamTransactionClient{stream}
  if err := x.ClientStream.SendMsg(in); err != nil { return nil, err }
  if err := x.ClientStream.CloseSend(); err != nil { return nil, err }
  return x,nil
}

type Sync_FetchStreamTransactionClient interface {
  Recv() (*TxWithIndex, error)
  grpc.ClientStream
}

type syncFetchStreamTransactionClient struct{
  grpc.ClientStream
}

func (x *syncFetchStreamTransactionClient) Recv() (*TxWithIndex, error) {
  m := new(TxWithIndex)
  if err := x.ClientStream.RecvMsg(m); err != nil { return nil, err }
  return m, nil
}

// Server API for Sync service
type SyncServer interface {
  CheckHash(context.Context, *Ping) (*flatbuffers.Builder, error)  
  GetPeers(context.Context, *Ping) (*flatbuffers.Builder, error)  
  GetTransactions(context.Context, *Ping) (*flatbuffers.Builder, error)  
  FetchStreamTransaction(*TxRequest, Sync_FetchStreamTransactionServer) error  
}

func RegisterSyncServer(s *grpc.Server, srv SyncServer) {
  s.RegisterService(&_Sync_serviceDesc, srv)
}

func _Sync_CheckHash_Handler(srv interface{}, ctx context.Context,
	dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
  in := new(Ping)
  if err := dec(in); err != nil { return nil, err }
  if interceptor == nil { return srv.(SyncServer).CheckHash(ctx, in) }
  info := &grpc.UnaryServerInfo{
    Server: srv,
    FullMethod: "/iroha.Sync/CheckHash",
  }
  
  handler := func(ctx context.Context, req interface{}) (interface{}, error) {
    return srv.(SyncServer).CheckHash(ctx, req.(* Ping))
  }
  return interceptor(ctx, in, info, handler)
}


func _Sync_GetPeers_Handler(srv interface{}, ctx context.Context,
	dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
  in := new(Ping)
  if err := dec(in); err != nil { return nil, err }
  if interceptor == nil { return srv.(SyncServer).GetPeers(ctx, in) }
  info := &grpc.UnaryServerInfo{
    Server: srv,
    FullMethod: "/iroha.Sync/GetPeers",
  }
  
  handler := func(ctx context.Context, req interface{}) (interface{}, error) {
    return srv.(SyncServer).GetPeers(ctx, req.(* Ping))
  }
  return interceptor(ctx, in, info, handler)
}


func _Sync_GetTransactions_Handler(srv interface{}, ctx context.Context,
	dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
  in := new(Ping)
  if err := dec(in); err != nil { return nil, err }
  if interceptor == nil { return srv.(SyncServer).GetTransactions(ctx, in) }
  info := &grpc.UnaryServerInfo{
    Server: srv,
    FullMethod: "/iroha.Sync/GetTransactions",
  }
  
  handler := func(ctx context.Context, req interface{}) (interface{}, error) {
    return srv.(SyncServer).GetTransactions(ctx, req.(* Ping))
  }
  return interceptor(ctx, in, info, handler)
}


func _Sync_FetchStreamTransaction_Handler(srv interface{}, stream grpc.ServerStream) error {
  m := new(TxRequest)
  if err := stream.RecvMsg(m); err != nil { return err }
  return srv.(SyncServer).FetchStreamTransaction(m, &syncFetchStreamTransactionServer{stream})
}

type Sync_FetchStreamTransactionServer interface { 
  Send(* flatbuffers.Builder) error
  grpc.ServerStream
}

type syncFetchStreamTransactionServer struct {
  grpc.ServerStream
}

func (x *syncFetchStreamTransactionServer) Send(m *flatbuffers.Builder) error {
  return x.ServerStream.SendMsg(m)
}


var _Sync_serviceDesc = grpc.ServiceDesc{
  ServiceName: "iroha.Sync",
  HandlerType: (*SyncServer)(nil),
  Methods: []grpc.MethodDesc{
    {
      MethodName: "checkHash",
      Handler: _Sync_checkHash_Handler, 
    },
    {
      MethodName: "getPeers",
      Handler: _Sync_getPeers_Handler, 
    },
    {
      MethodName: "getTransactions",
      Handler: _Sync_getTransactions_Handler, 
    },
  },
  Streams: []grpc.StreamDesc{
    {
      StreamName: "fetchStreamTransaction",
      Handler: _Sync_fetchStreamTransaction_Handler, 
      ServerStreams: true,
    },
  },
}

