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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.25.7
// source: google/cloud/cloudcontrolspartner/v1/monitoring.proto

package cloudcontrolspartnerpb

import (
	context "context"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_google_cloud_cloudcontrolspartner_v1_monitoring_proto protoreflect.FileDescriptor

var file_google_cloud_cloudcontrolspartner_v1_monitoring_proto_rawDesc = []byte{
	0x0a, 0x35, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x70, 0x61, 0x72, 0x74,
	0x6e, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x73, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x35, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73,
	0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x69, 0x6f, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xbe, 0x04, 0x0a, 0x1e,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x50, 0x61, 0x72,
	0x74, 0x6e, 0x65, 0x72, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x12, 0xe9,
	0x01, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x3b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x70, 0x61,
	0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x69, 0x6f,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x70, 0x61, 0x72, 0x74, 0x6e,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x69, 0x6f, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5c, 0xda, 0x41,
	0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x4d, 0x12, 0x4b, 0x2f,
	0x76, 0x31, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x6f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73,
	0x2f, 0x2a, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x2f, 0x2a, 0x7d, 0x2f,
	0x76, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0xd6, 0x01, 0x0a, 0x0c, 0x47,
	0x65, 0x74, 0x56, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x73, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x69,
	0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x5a, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x4d, 0x12, 0x4b, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d,
	0x65, 0x3d, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x2a, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f,
	0x61, 0x64, 0x73, 0x2f, 0x2a, 0x2f, 0x76, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x2a, 0x7d, 0x1a, 0x57, 0xca, 0x41, 0x23, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x73, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x2e, 0x68, 0x74,
	0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x42, 0x93, 0x02, 0x0a,
	0x28, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x70,
	0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x4d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5c, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x70,
	0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65,
	0x72, 0x70, 0x62, 0x3b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x73, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x70, 0x62, 0xaa, 0x02, 0x24, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x24, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x50, 0x61,
	0x72, 0x74, 0x6e, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x27, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_google_cloud_cloudcontrolspartner_v1_monitoring_proto_goTypes = []any{
	(*ListViolationsRequest)(nil),  // 0: google.cloud.cloudcontrolspartner.v1.ListViolationsRequest
	(*GetViolationRequest)(nil),    // 1: google.cloud.cloudcontrolspartner.v1.GetViolationRequest
	(*ListViolationsResponse)(nil), // 2: google.cloud.cloudcontrolspartner.v1.ListViolationsResponse
	(*Violation)(nil),              // 3: google.cloud.cloudcontrolspartner.v1.Violation
}
var file_google_cloud_cloudcontrolspartner_v1_monitoring_proto_depIdxs = []int32{
	0, // 0: google.cloud.cloudcontrolspartner.v1.CloudControlsPartnerMonitoring.ListViolations:input_type -> google.cloud.cloudcontrolspartner.v1.ListViolationsRequest
	1, // 1: google.cloud.cloudcontrolspartner.v1.CloudControlsPartnerMonitoring.GetViolation:input_type -> google.cloud.cloudcontrolspartner.v1.GetViolationRequest
	2, // 2: google.cloud.cloudcontrolspartner.v1.CloudControlsPartnerMonitoring.ListViolations:output_type -> google.cloud.cloudcontrolspartner.v1.ListViolationsResponse
	3, // 3: google.cloud.cloudcontrolspartner.v1.CloudControlsPartnerMonitoring.GetViolation:output_type -> google.cloud.cloudcontrolspartner.v1.Violation
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_cloud_cloudcontrolspartner_v1_monitoring_proto_init() }
func file_google_cloud_cloudcontrolspartner_v1_monitoring_proto_init() {
	if File_google_cloud_cloudcontrolspartner_v1_monitoring_proto != nil {
		return
	}
	file_google_cloud_cloudcontrolspartner_v1_violations_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_cloudcontrolspartner_v1_monitoring_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_cloudcontrolspartner_v1_monitoring_proto_goTypes,
		DependencyIndexes: file_google_cloud_cloudcontrolspartner_v1_monitoring_proto_depIdxs,
	}.Build()
	File_google_cloud_cloudcontrolspartner_v1_monitoring_proto = out.File
	file_google_cloud_cloudcontrolspartner_v1_monitoring_proto_rawDesc = nil
	file_google_cloud_cloudcontrolspartner_v1_monitoring_proto_goTypes = nil
	file_google_cloud_cloudcontrolspartner_v1_monitoring_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CloudControlsPartnerMonitoringClient is the client API for CloudControlsPartnerMonitoring service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CloudControlsPartnerMonitoringClient interface {
	// Lists Violations for a workload
	// Callers may also choose to read across multiple Customers or for a single
	// customer as per
	// [AIP-159](https://google.aip.dev/159) by using '-' (the hyphen or dash
	// character) as a wildcard character instead of {customer} & {workload}.
	// Format:
	// `organizations/{organization}/locations/{location}/customers/{customer}/workloads/{workload}`
	ListViolations(ctx context.Context, in *ListViolationsRequest, opts ...grpc.CallOption) (*ListViolationsResponse, error)
	// Gets details of a single Violation.
	GetViolation(ctx context.Context, in *GetViolationRequest, opts ...grpc.CallOption) (*Violation, error)
}

type cloudControlsPartnerMonitoringClient struct {
	cc grpc.ClientConnInterface
}

func NewCloudControlsPartnerMonitoringClient(cc grpc.ClientConnInterface) CloudControlsPartnerMonitoringClient {
	return &cloudControlsPartnerMonitoringClient{cc}
}

func (c *cloudControlsPartnerMonitoringClient) ListViolations(ctx context.Context, in *ListViolationsRequest, opts ...grpc.CallOption) (*ListViolationsResponse, error) {
	out := new(ListViolationsResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.cloudcontrolspartner.v1.CloudControlsPartnerMonitoring/ListViolations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudControlsPartnerMonitoringClient) GetViolation(ctx context.Context, in *GetViolationRequest, opts ...grpc.CallOption) (*Violation, error) {
	out := new(Violation)
	err := c.cc.Invoke(ctx, "/google.cloud.cloudcontrolspartner.v1.CloudControlsPartnerMonitoring/GetViolation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CloudControlsPartnerMonitoringServer is the server API for CloudControlsPartnerMonitoring service.
type CloudControlsPartnerMonitoringServer interface {
	// Lists Violations for a workload
	// Callers may also choose to read across multiple Customers or for a single
	// customer as per
	// [AIP-159](https://google.aip.dev/159) by using '-' (the hyphen or dash
	// character) as a wildcard character instead of {customer} & {workload}.
	// Format:
	// `organizations/{organization}/locations/{location}/customers/{customer}/workloads/{workload}`
	ListViolations(context.Context, *ListViolationsRequest) (*ListViolationsResponse, error)
	// Gets details of a single Violation.
	GetViolation(context.Context, *GetViolationRequest) (*Violation, error)
}

// UnimplementedCloudControlsPartnerMonitoringServer can be embedded to have forward compatible implementations.
type UnimplementedCloudControlsPartnerMonitoringServer struct {
}

func (*UnimplementedCloudControlsPartnerMonitoringServer) ListViolations(context.Context, *ListViolationsRequest) (*ListViolationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListViolations not implemented")
}
func (*UnimplementedCloudControlsPartnerMonitoringServer) GetViolation(context.Context, *GetViolationRequest) (*Violation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetViolation not implemented")
}

func RegisterCloudControlsPartnerMonitoringServer(s *grpc.Server, srv CloudControlsPartnerMonitoringServer) {
	s.RegisterService(&_CloudControlsPartnerMonitoring_serviceDesc, srv)
}

func _CloudControlsPartnerMonitoring_ListViolations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListViolationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudControlsPartnerMonitoringServer).ListViolations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.cloudcontrolspartner.v1.CloudControlsPartnerMonitoring/ListViolations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudControlsPartnerMonitoringServer).ListViolations(ctx, req.(*ListViolationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudControlsPartnerMonitoring_GetViolation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetViolationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudControlsPartnerMonitoringServer).GetViolation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.cloudcontrolspartner.v1.CloudControlsPartnerMonitoring/GetViolation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudControlsPartnerMonitoringServer).GetViolation(ctx, req.(*GetViolationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CloudControlsPartnerMonitoring_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.cloudcontrolspartner.v1.CloudControlsPartnerMonitoring",
	HandlerType: (*CloudControlsPartnerMonitoringServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListViolations",
			Handler:    _CloudControlsPartnerMonitoring_ListViolations_Handler,
		},
		{
			MethodName: "GetViolation",
			Handler:    _CloudControlsPartnerMonitoring_GetViolation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/cloudcontrolspartner/v1/monitoring.proto",
}
