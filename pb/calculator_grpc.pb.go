// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: calculator.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Calculator_Add_FullMethodName    = "/calculator.Calculator/Add"
	Calculator_Divide_FullMethodName = "/calculator.Calculator/Divide"
	Calculator_Sum_FullMethodName    = "/calculator.Calculator/Sum"
)

// CalculatorClient is the client API for Calculator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalculatorClient interface {
	Add(ctx context.Context, in *CalculationRequest, opts ...grpc.CallOption) (*CalculationResponse, error)
	Divide(ctx context.Context, in *CalculationRequest, opts ...grpc.CallOption) (*CalculationResponse, error)
	Sum(ctx context.Context, in *NumberRequest, opts ...grpc.CallOption) (*CalculationResponse, error)
}

type calculatorClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorClient(cc grpc.ClientConnInterface) CalculatorClient {
	return &calculatorClient{cc}
}

func (c *calculatorClient) Add(ctx context.Context, in *CalculationRequest, opts ...grpc.CallOption) (*CalculationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CalculationResponse)
	err := c.cc.Invoke(ctx, Calculator_Add_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorClient) Divide(ctx context.Context, in *CalculationRequest, opts ...grpc.CallOption) (*CalculationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CalculationResponse)
	err := c.cc.Invoke(ctx, Calculator_Divide_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorClient) Sum(ctx context.Context, in *NumberRequest, opts ...grpc.CallOption) (*CalculationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CalculationResponse)
	err := c.cc.Invoke(ctx, Calculator_Sum_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculatorServer is the server API for Calculator service.
// All implementations must embed UnimplementedCalculatorServer
// for forward compatibility.
type CalculatorServer interface {
	Add(context.Context, *CalculationRequest) (*CalculationResponse, error)
	Divide(context.Context, *CalculationRequest) (*CalculationResponse, error)
	Sum(context.Context, *NumberRequest) (*CalculationResponse, error)
	mustEmbedUnimplementedCalculatorServer()
}

// UnimplementedCalculatorServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCalculatorServer struct{}

func (UnimplementedCalculatorServer) Add(context.Context, *CalculationRequest) (*CalculationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedCalculatorServer) Divide(context.Context, *CalculationRequest) (*CalculationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Divide not implemented")
}
func (UnimplementedCalculatorServer) Sum(context.Context, *NumberRequest) (*CalculationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (UnimplementedCalculatorServer) mustEmbedUnimplementedCalculatorServer() {}
func (UnimplementedCalculatorServer) testEmbeddedByValue()                    {}

// UnsafeCalculatorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalculatorServer will
// result in compilation errors.
type UnsafeCalculatorServer interface {
	mustEmbedUnimplementedCalculatorServer()
}

func RegisterCalculatorServer(s grpc.ServiceRegistrar, srv CalculatorServer) {
	// If the following call pancis, it indicates UnimplementedCalculatorServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Calculator_ServiceDesc, srv)
}

func _Calculator_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Calculator_Add_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Add(ctx, req.(*CalculationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calculator_Divide_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Divide(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Calculator_Divide_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Divide(ctx, req.(*CalculationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calculator_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Calculator_Sum_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Sum(ctx, req.(*NumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Calculator_ServiceDesc is the grpc.ServiceDesc for Calculator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Calculator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.Calculator",
	HandlerType: (*CalculatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Calculator_Add_Handler,
		},
		{
			MethodName: "Divide",
			Handler:    _Calculator_Divide_Handler,
		},
		{
			MethodName: "Sum",
			Handler:    _Calculator_Sum_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calculator.proto",
}
