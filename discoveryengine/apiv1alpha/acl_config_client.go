// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package discoveryengine

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"math"
	"net/http"
	"net/url"
	"time"

	discoveryenginepb "cloud.google.com/go/discoveryengine/apiv1alpha/discoveryenginepb"
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var newAclConfigClientHook clientHook

// AclConfigCallOptions contains the retry settings for each method of AclConfigClient.
type AclConfigCallOptions struct {
	UpdateAclConfig []gax.CallOption
	GetAclConfig    []gax.CallOption
	CancelOperation []gax.CallOption
	GetOperation    []gax.CallOption
	ListOperations  []gax.CallOption
}

func defaultAclConfigGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("discoveryengine.googleapis.com:443"),
		internaloption.WithDefaultEndpointTemplate("discoveryengine.UNIVERSE_DOMAIN:443"),
		internaloption.WithDefaultMTLSEndpoint("discoveryengine.mtls.googleapis.com:443"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://discoveryengine.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		internaloption.EnableNewAuthLibrary(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultAclConfigCallOptions() *AclConfigCallOptions {
	return &AclConfigCallOptions{
		UpdateAclConfig: []gax.CallOption{},
		GetAclConfig:    []gax.CallOption{},
		CancelOperation: []gax.CallOption{
			gax.WithTimeout(30000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetOperation: []gax.CallOption{
			gax.WithTimeout(30000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ListOperations: []gax.CallOption{
			gax.WithTimeout(300000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        30000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

func defaultAclConfigRESTCallOptions() *AclConfigCallOptions {
	return &AclConfigCallOptions{
		UpdateAclConfig: []gax.CallOption{},
		GetAclConfig:    []gax.CallOption{},
		CancelOperation: []gax.CallOption{
			gax.WithTimeout(30000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
		GetOperation: []gax.CallOption{
			gax.WithTimeout(30000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
		ListOperations: []gax.CallOption{
			gax.WithTimeout(300000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        30000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
	}
}

// internalAclConfigClient is an interface that defines the methods available from Discovery Engine API.
type internalAclConfigClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	UpdateAclConfig(context.Context, *discoveryenginepb.UpdateAclConfigRequest, ...gax.CallOption) (*discoveryenginepb.AclConfig, error)
	GetAclConfig(context.Context, *discoveryenginepb.GetAclConfigRequest, ...gax.CallOption) (*discoveryenginepb.AclConfig, error)
	CancelOperation(context.Context, *longrunningpb.CancelOperationRequest, ...gax.CallOption) error
	GetOperation(context.Context, *longrunningpb.GetOperationRequest, ...gax.CallOption) (*longrunningpb.Operation, error)
	ListOperations(context.Context, *longrunningpb.ListOperationsRequest, ...gax.CallOption) *OperationIterator
}

// AclConfigClient is a client for interacting with Discovery Engine API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service for managing Acl Configuration.
type AclConfigClient struct {
	// The internal transport-dependent client.
	internalClient internalAclConfigClient

	// The call options for this service.
	CallOptions *AclConfigCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *AclConfigClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *AclConfigClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *AclConfigClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// UpdateAclConfig default ACL configuration for use in a location of a customer’s project.
// Updates will only reflect to new data stores. Existing data stores will
// still use the old value.
func (c *AclConfigClient) UpdateAclConfig(ctx context.Context, req *discoveryenginepb.UpdateAclConfigRequest, opts ...gax.CallOption) (*discoveryenginepb.AclConfig, error) {
	return c.internalClient.UpdateAclConfig(ctx, req, opts...)
}

// GetAclConfig gets the AclConfig.
func (c *AclConfigClient) GetAclConfig(ctx context.Context, req *discoveryenginepb.GetAclConfigRequest, opts ...gax.CallOption) (*discoveryenginepb.AclConfig, error) {
	return c.internalClient.GetAclConfig(ctx, req, opts...)
}

// CancelOperation is a utility method from google.longrunning.Operations.
func (c *AclConfigClient) CancelOperation(ctx context.Context, req *longrunningpb.CancelOperationRequest, opts ...gax.CallOption) error {
	return c.internalClient.CancelOperation(ctx, req, opts...)
}

// GetOperation is a utility method from google.longrunning.Operations.
func (c *AclConfigClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	return c.internalClient.GetOperation(ctx, req, opts...)
}

// ListOperations is a utility method from google.longrunning.Operations.
func (c *AclConfigClient) ListOperations(ctx context.Context, req *longrunningpb.ListOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	return c.internalClient.ListOperations(ctx, req, opts...)
}

// aclConfigGRPCClient is a client for interacting with Discovery Engine API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type aclConfigGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing AclConfigClient
	CallOptions **AclConfigCallOptions

	// The gRPC API client.
	aclConfigClient discoveryenginepb.AclConfigServiceClient

	operationsClient longrunningpb.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string

	logger *slog.Logger
}

// NewAclConfigClient creates a new acl config service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service for managing Acl Configuration.
func NewAclConfigClient(ctx context.Context, opts ...option.ClientOption) (*AclConfigClient, error) {
	clientOpts := defaultAclConfigGRPCClientOptions()
	if newAclConfigClientHook != nil {
		hookOpts, err := newAclConfigClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := AclConfigClient{CallOptions: defaultAclConfigCallOptions()}

	c := &aclConfigGRPCClient{
		connPool:         connPool,
		aclConfigClient:  discoveryenginepb.NewAclConfigServiceClient(connPool),
		CallOptions:      &client.CallOptions,
		logger:           internaloption.GetLogger(opts),
		operationsClient: longrunningpb.NewOperationsClient(connPool),
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *aclConfigGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *aclConfigGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version, "pb", protoVersion)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *aclConfigGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type aclConfigRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing AclConfigClient
	CallOptions **AclConfigCallOptions

	logger *slog.Logger
}

// NewAclConfigRESTClient creates a new acl config service rest client.
//
// Service for managing Acl Configuration.
func NewAclConfigRESTClient(ctx context.Context, opts ...option.ClientOption) (*AclConfigClient, error) {
	clientOpts := append(defaultAclConfigRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultAclConfigRESTCallOptions()
	c := &aclConfigRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
		logger:      internaloption.GetLogger(opts),
	}
	c.setGoogleClientInfo()

	return &AclConfigClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultAclConfigRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://discoveryengine.googleapis.com"),
		internaloption.WithDefaultEndpointTemplate("https://discoveryengine.UNIVERSE_DOMAIN"),
		internaloption.WithDefaultMTLSEndpoint("https://discoveryengine.mtls.googleapis.com"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://discoveryengine.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableNewAuthLibrary(),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *aclConfigRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN", "pb", protoVersion)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *aclConfigRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *aclConfigRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *aclConfigGRPCClient) UpdateAclConfig(ctx context.Context, req *discoveryenginepb.UpdateAclConfigRequest, opts ...gax.CallOption) (*discoveryenginepb.AclConfig, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "acl_config.name", url.QueryEscape(req.GetAclConfig().GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).UpdateAclConfig[0:len((*c.CallOptions).UpdateAclConfig):len((*c.CallOptions).UpdateAclConfig)], opts...)
	var resp *discoveryenginepb.AclConfig
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.aclConfigClient.UpdateAclConfig, req, settings.GRPC, c.logger, "UpdateAclConfig")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *aclConfigGRPCClient) GetAclConfig(ctx context.Context, req *discoveryenginepb.GetAclConfigRequest, opts ...gax.CallOption) (*discoveryenginepb.AclConfig, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetAclConfig[0:len((*c.CallOptions).GetAclConfig):len((*c.CallOptions).GetAclConfig)], opts...)
	var resp *discoveryenginepb.AclConfig
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.aclConfigClient.GetAclConfig, req, settings.GRPC, c.logger, "GetAclConfig")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *aclConfigGRPCClient) CancelOperation(ctx context.Context, req *longrunningpb.CancelOperationRequest, opts ...gax.CallOption) error {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).CancelOperation[0:len((*c.CallOptions).CancelOperation):len((*c.CallOptions).CancelOperation)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = executeRPC(ctx, c.operationsClient.CancelOperation, req, settings.GRPC, c.logger, "CancelOperation")
		return err
	}, opts...)
	return err
}

func (c *aclConfigGRPCClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetOperation[0:len((*c.CallOptions).GetOperation):len((*c.CallOptions).GetOperation)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.operationsClient.GetOperation, req, settings.GRPC, c.logger, "GetOperation")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *aclConfigGRPCClient) ListOperations(ctx context.Context, req *longrunningpb.ListOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListOperations[0:len((*c.CallOptions).ListOperations):len((*c.CallOptions).ListOperations)], opts...)
	it := &OperationIterator{}
	req = proto.Clone(req).(*longrunningpb.ListOperationsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*longrunningpb.Operation, string, error) {
		resp := &longrunningpb.ListOperationsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = executeRPC(ctx, c.operationsClient.ListOperations, req, settings.GRPC, c.logger, "ListOperations")
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetOperations(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// UpdateAclConfig default ACL configuration for use in a location of a customer’s project.
// Updates will only reflect to new data stores. Existing data stores will
// still use the old value.
func (c *aclConfigRESTClient) UpdateAclConfig(ctx context.Context, req *discoveryenginepb.UpdateAclConfigRequest, opts ...gax.CallOption) (*discoveryenginepb.AclConfig, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetAclConfig()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1alpha/%v", req.GetAclConfig().GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "acl_config.name", url.QueryEscape(req.GetAclConfig().GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).UpdateAclConfig[0:len((*c.CallOptions).UpdateAclConfig):len((*c.CallOptions).UpdateAclConfig)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &discoveryenginepb.AclConfig{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("PATCH", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, jsonReq, "UpdateAclConfig")
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// GetAclConfig gets the AclConfig.
func (c *aclConfigRESTClient) GetAclConfig(ctx context.Context, req *discoveryenginepb.GetAclConfigRequest, opts ...gax.CallOption) (*discoveryenginepb.AclConfig, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1alpha/%v", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).GetAclConfig[0:len((*c.CallOptions).GetAclConfig):len((*c.CallOptions).GetAclConfig)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &discoveryenginepb.AclConfig{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "GetAclConfig")
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// CancelOperation is a utility method from google.longrunning.Operations.
func (c *aclConfigRESTClient) CancelOperation(ctx context.Context, req *longrunningpb.CancelOperationRequest, opts ...gax.CallOption) error {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return err
	}
	baseUrl.Path += fmt.Sprintf("/v1alpha/%v:cancel", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	return gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		_, err = executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, jsonReq, "CancelOperation")
		return err
	}, opts...)
}

// GetOperation is a utility method from google.longrunning.Operations.
func (c *aclConfigRESTClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1alpha/%v", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).GetOperation[0:len((*c.CallOptions).GetOperation):len((*c.CallOptions).GetOperation)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &longrunningpb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "GetOperation")
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// ListOperations is a utility method from google.longrunning.Operations.
func (c *aclConfigRESTClient) ListOperations(ctx context.Context, req *longrunningpb.ListOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	it := &OperationIterator{}
	req = proto.Clone(req).(*longrunningpb.ListOperationsRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*longrunningpb.Operation, string, error) {
		resp := &longrunningpb.ListOperationsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		baseUrl, err := url.Parse(c.endpoint)
		if err != nil {
			return nil, "", err
		}
		baseUrl.Path += fmt.Sprintf("/v1alpha/%v/operations", req.GetName())

		params := url.Values{}
		params.Add("$alt", "json;enum-encoding=int")
		if req.GetFilter() != "" {
			params.Add("filter", fmt.Sprintf("%v", req.GetFilter()))
		}
		if req.GetPageSize() != 0 {
			params.Add("pageSize", fmt.Sprintf("%v", req.GetPageSize()))
		}
		if req.GetPageToken() != "" {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}

		baseUrl.RawQuery = params.Encode()

		// Build HTTP headers from client and context metadata.
		hds := append(c.xGoogHeaders, "Content-Type", "application/json")
		headers := gax.BuildHeaders(ctx, hds...)
		e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			if settings.Path != "" {
				baseUrl.Path = settings.Path
			}
			httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
			if err != nil {
				return err
			}
			httpReq.Header = headers

			buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "ListOperations")
			if err != nil {
				return err
			}
			if err := unm.Unmarshal(buf, resp); err != nil {
				return err
			}

			return nil
		}, opts...)
		if e != nil {
			return nil, "", e
		}
		it.Response = resp
		return resp.GetOperations(), resp.GetNextPageToken(), nil
	}

	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}
