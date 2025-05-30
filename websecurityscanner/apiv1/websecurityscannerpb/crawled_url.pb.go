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
// source: google/cloud/websecurityscanner/v1/crawled_url.proto

package websecurityscannerpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A CrawledUrl resource represents a URL that was crawled during a ScanRun. Web
// Security Scanner Service crawls the web applications, following all links
// within the scope of sites, to find the URLs to test against.
type CrawledUrl struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The http method of the request that was used to visit the URL, in
	// uppercase.
	HttpMethod string `protobuf:"bytes,1,opt,name=http_method,json=httpMethod,proto3" json:"http_method,omitempty"`
	// Output only. The URL that was crawled.
	Url string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	// Output only. The body of the request that was used to visit the URL.
	Body string `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *CrawledUrl) Reset() {
	*x = CrawledUrl{}
	mi := &file_google_cloud_websecurityscanner_v1_crawled_url_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CrawledUrl) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CrawledUrl) ProtoMessage() {}

func (x *CrawledUrl) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_websecurityscanner_v1_crawled_url_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CrawledUrl.ProtoReflect.Descriptor instead.
func (*CrawledUrl) Descriptor() ([]byte, []int) {
	return file_google_cloud_websecurityscanner_v1_crawled_url_proto_rawDescGZIP(), []int{0}
}

func (x *CrawledUrl) GetHttpMethod() string {
	if x != nil {
		return x.HttpMethod
	}
	return ""
}

func (x *CrawledUrl) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *CrawledUrl) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

var File_google_cloud_websecurityscanner_v1_crawled_url_proto protoreflect.FileDescriptor

var file_google_cloud_websecurityscanner_v1_crawled_url_proto_rawDesc = []byte{
	0x0a, 0x34, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x77,
	0x65, 0x62, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x64, 0x5f, 0x75, 0x72, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x22, 0x53, 0x0a, 0x0a, 0x43, 0x72,
	0x61, 0x77, 0x6c, 0x65, 0x64, 0x55, 0x72, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x68, 0x74, 0x74, 0x70,
	0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x68,
	0x74, 0x74, 0x70, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x62,
	0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x42,
	0x85, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x43, 0x72, 0x61, 0x77,
	0x6c, 0x65, 0x64, 0x55, 0x72, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x56, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x6f, 0x2f, 0x77, 0x65, 0x62, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x73, 0x63,
	0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x77, 0x65, 0x62, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x70, 0x62,
	0x3b, 0x77, 0x65, 0x62, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x73, 0x63, 0x61, 0x6e,
	0x6e, 0x65, 0x72, 0x70, 0x62, 0xaa, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x57, 0x65, 0x62, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x53, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x22, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x57, 0x65, 0x62, 0x53, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x53, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xea,
	0x02, 0x25, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a,
	0x3a, 0x57, 0x65, 0x62, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x53, 0x63, 0x61, 0x6e,
	0x6e, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_websecurityscanner_v1_crawled_url_proto_rawDescOnce sync.Once
	file_google_cloud_websecurityscanner_v1_crawled_url_proto_rawDescData = file_google_cloud_websecurityscanner_v1_crawled_url_proto_rawDesc
)

func file_google_cloud_websecurityscanner_v1_crawled_url_proto_rawDescGZIP() []byte {
	file_google_cloud_websecurityscanner_v1_crawled_url_proto_rawDescOnce.Do(func() {
		file_google_cloud_websecurityscanner_v1_crawled_url_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_websecurityscanner_v1_crawled_url_proto_rawDescData)
	})
	return file_google_cloud_websecurityscanner_v1_crawled_url_proto_rawDescData
}

var file_google_cloud_websecurityscanner_v1_crawled_url_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_websecurityscanner_v1_crawled_url_proto_goTypes = []any{
	(*CrawledUrl)(nil), // 0: google.cloud.websecurityscanner.v1.CrawledUrl
}
var file_google_cloud_websecurityscanner_v1_crawled_url_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_cloud_websecurityscanner_v1_crawled_url_proto_init() }
func file_google_cloud_websecurityscanner_v1_crawled_url_proto_init() {
	if File_google_cloud_websecurityscanner_v1_crawled_url_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_websecurityscanner_v1_crawled_url_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_websecurityscanner_v1_crawled_url_proto_goTypes,
		DependencyIndexes: file_google_cloud_websecurityscanner_v1_crawled_url_proto_depIdxs,
		MessageInfos:      file_google_cloud_websecurityscanner_v1_crawled_url_proto_msgTypes,
	}.Build()
	File_google_cloud_websecurityscanner_v1_crawled_url_proto = out.File
	file_google_cloud_websecurityscanner_v1_crawled_url_proto_rawDesc = nil
	file_google_cloud_websecurityscanner_v1_crawled_url_proto_goTypes = nil
	file_google_cloud_websecurityscanner_v1_crawled_url_proto_depIdxs = nil
}
