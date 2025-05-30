// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.7
// source: google/shopping/merchant/promotions/v1beta/promotions.proto

package promotionspb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	PromotionsService_InsertPromotion_FullMethodName = "/google.shopping.merchant.promotions.v1beta.PromotionsService/InsertPromotion"
	PromotionsService_GetPromotion_FullMethodName    = "/google.shopping.merchant.promotions.v1beta.PromotionsService/GetPromotion"
	PromotionsService_ListPromotions_FullMethodName  = "/google.shopping.merchant.promotions.v1beta.PromotionsService/ListPromotions"
)

// PromotionsServiceClient is the client API for PromotionsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PromotionsServiceClient interface {
	// Inserts a promotion for your Merchant Center account. If the promotion
	// already exists, then it updates the promotion instead.
	InsertPromotion(ctx context.Context, in *InsertPromotionRequest, opts ...grpc.CallOption) (*Promotion, error)
	// Retrieves the promotion from your Merchant Center account.
	//
	// After inserting or updating a promotion input, it may take several
	// minutes before the updated promotion can be retrieved.
	GetPromotion(ctx context.Context, in *GetPromotionRequest, opts ...grpc.CallOption) (*Promotion, error)
	// Lists the promotions in your Merchant Center account. The
	// response might contain fewer items than specified by `pageSize`. Rely on
	// `pageToken` to determine if there are more items to be requested.
	//
	// After inserting or updating a promotion, it may take several minutes before
	// the updated processed promotion can be retrieved.
	ListPromotions(ctx context.Context, in *ListPromotionsRequest, opts ...grpc.CallOption) (*ListPromotionsResponse, error)
}

type promotionsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPromotionsServiceClient(cc grpc.ClientConnInterface) PromotionsServiceClient {
	return &promotionsServiceClient{cc}
}

func (c *promotionsServiceClient) InsertPromotion(ctx context.Context, in *InsertPromotionRequest, opts ...grpc.CallOption) (*Promotion, error) {
	out := new(Promotion)
	err := c.cc.Invoke(ctx, PromotionsService_InsertPromotion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *promotionsServiceClient) GetPromotion(ctx context.Context, in *GetPromotionRequest, opts ...grpc.CallOption) (*Promotion, error) {
	out := new(Promotion)
	err := c.cc.Invoke(ctx, PromotionsService_GetPromotion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *promotionsServiceClient) ListPromotions(ctx context.Context, in *ListPromotionsRequest, opts ...grpc.CallOption) (*ListPromotionsResponse, error) {
	out := new(ListPromotionsResponse)
	err := c.cc.Invoke(ctx, PromotionsService_ListPromotions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PromotionsServiceServer is the server API for PromotionsService service.
// All implementations should embed UnimplementedPromotionsServiceServer
// for forward compatibility
type PromotionsServiceServer interface {
	// Inserts a promotion for your Merchant Center account. If the promotion
	// already exists, then it updates the promotion instead.
	InsertPromotion(context.Context, *InsertPromotionRequest) (*Promotion, error)
	// Retrieves the promotion from your Merchant Center account.
	//
	// After inserting or updating a promotion input, it may take several
	// minutes before the updated promotion can be retrieved.
	GetPromotion(context.Context, *GetPromotionRequest) (*Promotion, error)
	// Lists the promotions in your Merchant Center account. The
	// response might contain fewer items than specified by `pageSize`. Rely on
	// `pageToken` to determine if there are more items to be requested.
	//
	// After inserting or updating a promotion, it may take several minutes before
	// the updated processed promotion can be retrieved.
	ListPromotions(context.Context, *ListPromotionsRequest) (*ListPromotionsResponse, error)
}

// UnimplementedPromotionsServiceServer should be embedded to have forward compatible implementations.
type UnimplementedPromotionsServiceServer struct {
}

func (UnimplementedPromotionsServiceServer) InsertPromotion(context.Context, *InsertPromotionRequest) (*Promotion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertPromotion not implemented")
}
func (UnimplementedPromotionsServiceServer) GetPromotion(context.Context, *GetPromotionRequest) (*Promotion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPromotion not implemented")
}
func (UnimplementedPromotionsServiceServer) ListPromotions(context.Context, *ListPromotionsRequest) (*ListPromotionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPromotions not implemented")
}

// UnsafePromotionsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PromotionsServiceServer will
// result in compilation errors.
type UnsafePromotionsServiceServer interface {
	mustEmbedUnimplementedPromotionsServiceServer()
}

func RegisterPromotionsServiceServer(s grpc.ServiceRegistrar, srv PromotionsServiceServer) {
	s.RegisterService(&PromotionsService_ServiceDesc, srv)
}

func _PromotionsService_InsertPromotion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertPromotionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromotionsServiceServer).InsertPromotion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PromotionsService_InsertPromotion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromotionsServiceServer).InsertPromotion(ctx, req.(*InsertPromotionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PromotionsService_GetPromotion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPromotionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromotionsServiceServer).GetPromotion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PromotionsService_GetPromotion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromotionsServiceServer).GetPromotion(ctx, req.(*GetPromotionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PromotionsService_ListPromotions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPromotionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromotionsServiceServer).ListPromotions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PromotionsService_ListPromotions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromotionsServiceServer).ListPromotions(ctx, req.(*ListPromotionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PromotionsService_ServiceDesc is the grpc.ServiceDesc for PromotionsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PromotionsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.shopping.merchant.promotions.v1beta.PromotionsService",
	HandlerType: (*PromotionsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InsertPromotion",
			Handler:    _PromotionsService_InsertPromotion_Handler,
		},
		{
			MethodName: "GetPromotion",
			Handler:    _PromotionsService_GetPromotion_Handler,
		},
		{
			MethodName: "ListPromotions",
			Handler:    _PromotionsService_ListPromotions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/shopping/merchant/promotions/v1beta/promotions.proto",
}
