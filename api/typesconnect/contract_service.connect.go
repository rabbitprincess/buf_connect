// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: contract_service.proto

package typesconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	types "github.com/gokch/buf_connect/types"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// ContractServiceName is the fully-qualified name of the ContractService service.
	ContractServiceName = "types.ContractService"
)

// ContractServiceClient is a client for the types.ContractService service.
type ContractServiceClient interface {
	TxMake(context.Context, *connect_go.Request[types.Empty]) (*connect_go.Response[types.Empty], error)
	TxSign(context.Context, *connect_go.Request[types.Empty]) (*connect_go.Response[types.Empty], error)
	TxSend(context.Context, *connect_go.Request[types.Empty]) (*connect_go.Response[types.Empty], error)
	// all in one
	TxCommit(context.Context, *connect_go.Request[types.Empty]) (*connect_go.Response[types.Empty], error)
	Deposit(context.Context, *connect_go.Request[types.Empty]) (*connect_go.Response[types.Empty], error)
}

// NewContractServiceClient constructs a client for the types.ContractService service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewContractServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) ContractServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &contractServiceClient{
		txMake: connect_go.NewClient[types.Empty, types.Empty](
			httpClient,
			baseURL+"/types.ContractService/TxMake",
			opts...,
		),
		txSign: connect_go.NewClient[types.Empty, types.Empty](
			httpClient,
			baseURL+"/types.ContractService/TxSign",
			opts...,
		),
		txSend: connect_go.NewClient[types.Empty, types.Empty](
			httpClient,
			baseURL+"/types.ContractService/TxSend",
			opts...,
		),
		txCommit: connect_go.NewClient[types.Empty, types.Empty](
			httpClient,
			baseURL+"/types.ContractService/TxCommit",
			opts...,
		),
		deposit: connect_go.NewClient[types.Empty, types.Empty](
			httpClient,
			baseURL+"/types.ContractService/Deposit",
			opts...,
		),
	}
}

// contractServiceClient implements ContractServiceClient.
type contractServiceClient struct {
	txMake   *connect_go.Client[types.Empty, types.Empty]
	txSign   *connect_go.Client[types.Empty, types.Empty]
	txSend   *connect_go.Client[types.Empty, types.Empty]
	txCommit *connect_go.Client[types.Empty, types.Empty]
	deposit  *connect_go.Client[types.Empty, types.Empty]
}

// TxMake calls types.ContractService.TxMake.
func (c *contractServiceClient) TxMake(ctx context.Context, req *connect_go.Request[types.Empty]) (*connect_go.Response[types.Empty], error) {
	return c.txMake.CallUnary(ctx, req)
}

// TxSign calls types.ContractService.TxSign.
func (c *contractServiceClient) TxSign(ctx context.Context, req *connect_go.Request[types.Empty]) (*connect_go.Response[types.Empty], error) {
	return c.txSign.CallUnary(ctx, req)
}

// TxSend calls types.ContractService.TxSend.
func (c *contractServiceClient) TxSend(ctx context.Context, req *connect_go.Request[types.Empty]) (*connect_go.Response[types.Empty], error) {
	return c.txSend.CallUnary(ctx, req)
}

// TxCommit calls types.ContractService.TxCommit.
func (c *contractServiceClient) TxCommit(ctx context.Context, req *connect_go.Request[types.Empty]) (*connect_go.Response[types.Empty], error) {
	return c.txCommit.CallUnary(ctx, req)
}

// Deposit calls types.ContractService.Deposit.
func (c *contractServiceClient) Deposit(ctx context.Context, req *connect_go.Request[types.Empty]) (*connect_go.Response[types.Empty], error) {
	return c.deposit.CallUnary(ctx, req)
}

// ContractServiceHandler is an implementation of the types.ContractService service.
type ContractServiceHandler interface {
	TxMake(context.Context, *connect_go.Request[types.Empty]) (*connect_go.Response[types.Empty], error)
	TxSign(context.Context, *connect_go.Request[types.Empty]) (*connect_go.Response[types.Empty], error)
	TxSend(context.Context, *connect_go.Request[types.Empty]) (*connect_go.Response[types.Empty], error)
	// all in one
	TxCommit(context.Context, *connect_go.Request[types.Empty]) (*connect_go.Response[types.Empty], error)
	Deposit(context.Context, *connect_go.Request[types.Empty]) (*connect_go.Response[types.Empty], error)
}

// NewContractServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewContractServiceHandler(svc ContractServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/types.ContractService/TxMake", connect_go.NewUnaryHandler(
		"/types.ContractService/TxMake",
		svc.TxMake,
		opts...,
	))
	mux.Handle("/types.ContractService/TxSign", connect_go.NewUnaryHandler(
		"/types.ContractService/TxSign",
		svc.TxSign,
		opts...,
	))
	mux.Handle("/types.ContractService/TxSend", connect_go.NewUnaryHandler(
		"/types.ContractService/TxSend",
		svc.TxSend,
		opts...,
	))
	mux.Handle("/types.ContractService/TxCommit", connect_go.NewUnaryHandler(
		"/types.ContractService/TxCommit",
		svc.TxCommit,
		opts...,
	))
	mux.Handle("/types.ContractService/Deposit", connect_go.NewUnaryHandler(
		"/types.ContractService/Deposit",
		svc.Deposit,
		opts...,
	))
	return "/types.ContractService/", mux
}

// UnimplementedContractServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedContractServiceHandler struct{}

func (UnimplementedContractServiceHandler) TxMake(context.Context, *connect_go.Request[types.Empty]) (*connect_go.Response[types.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("types.ContractService.TxMake is not implemented"))
}

func (UnimplementedContractServiceHandler) TxSign(context.Context, *connect_go.Request[types.Empty]) (*connect_go.Response[types.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("types.ContractService.TxSign is not implemented"))
}

func (UnimplementedContractServiceHandler) TxSend(context.Context, *connect_go.Request[types.Empty]) (*connect_go.Response[types.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("types.ContractService.TxSend is not implemented"))
}

func (UnimplementedContractServiceHandler) TxCommit(context.Context, *connect_go.Request[types.Empty]) (*connect_go.Response[types.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("types.ContractService.TxCommit is not implemented"))
}

func (UnimplementedContractServiceHandler) Deposit(context.Context, *connect_go.Request[types.Empty]) (*connect_go.Response[types.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("types.ContractService.Deposit is not implemented"))
}