// Copyright 2024 Google LLC
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

package cloudcontrolspartner

import (
	"context"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"time"

	cloudcontrolspartnerpb "cloud.google.com/go/cloudcontrolspartner/apiv1/cloudcontrolspartnerpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
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

var newCloudControlsPartnerMonitoringClientHook clientHook

// CloudControlsPartnerMonitoringCallOptions contains the retry settings for each method of CloudControlsPartnerMonitoringClient.
type CloudControlsPartnerMonitoringCallOptions struct {
	ListViolations []gax.CallOption
	GetViolation   []gax.CallOption
}

func defaultCloudControlsPartnerMonitoringGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("cloudcontrolspartner.googleapis.com:443"),
		internaloption.WithDefaultEndpointTemplate("cloudcontrolspartner.UNIVERSE_DOMAIN:443"),
		internaloption.WithDefaultMTLSEndpoint("cloudcontrolspartner.mtls.googleapis.com:443"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://cloudcontrolspartner.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultCloudControlsPartnerMonitoringCallOptions() *CloudControlsPartnerMonitoringCallOptions {
	return &CloudControlsPartnerMonitoringCallOptions{
		ListViolations: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
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
		GetViolation: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
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
	}
}

func defaultCloudControlsPartnerMonitoringRESTCallOptions() *CloudControlsPartnerMonitoringCallOptions {
	return &CloudControlsPartnerMonitoringCallOptions{
		ListViolations: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
		GetViolation: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
	}
}

// internalCloudControlsPartnerMonitoringClient is an interface that defines the methods available from Cloud Controls Partner API.
type internalCloudControlsPartnerMonitoringClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ListViolations(context.Context, *cloudcontrolspartnerpb.ListViolationsRequest, ...gax.CallOption) *ViolationIterator
	GetViolation(context.Context, *cloudcontrolspartnerpb.GetViolationRequest, ...gax.CallOption) (*cloudcontrolspartnerpb.Violation, error)
}

// CloudControlsPartnerMonitoringClient is a client for interacting with Cloud Controls Partner API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service describing handlers for resources
type CloudControlsPartnerMonitoringClient struct {
	// The internal transport-dependent client.
	internalClient internalCloudControlsPartnerMonitoringClient

	// The call options for this service.
	CallOptions *CloudControlsPartnerMonitoringCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *CloudControlsPartnerMonitoringClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *CloudControlsPartnerMonitoringClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *CloudControlsPartnerMonitoringClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ListViolations lists Violations for a workload
// Callers may also choose to read across multiple Customers or for a single
// customer as per
// AIP-159 (at https://google.aip.dev/159) by using ‘-’ (the hyphen or dash
// character) as a wildcard character instead of {customer} & {workload}.
// Format:
// organizations/{organization}/locations/{location}/customers/{customer}/workloads/{workload}
func (c *CloudControlsPartnerMonitoringClient) ListViolations(ctx context.Context, req *cloudcontrolspartnerpb.ListViolationsRequest, opts ...gax.CallOption) *ViolationIterator {
	return c.internalClient.ListViolations(ctx, req, opts...)
}

// GetViolation gets details of a single Violation.
func (c *CloudControlsPartnerMonitoringClient) GetViolation(ctx context.Context, req *cloudcontrolspartnerpb.GetViolationRequest, opts ...gax.CallOption) (*cloudcontrolspartnerpb.Violation, error) {
	return c.internalClient.GetViolation(ctx, req, opts...)
}

// cloudControlsPartnerMonitoringGRPCClient is a client for interacting with Cloud Controls Partner API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type cloudControlsPartnerMonitoringGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing CloudControlsPartnerMonitoringClient
	CallOptions **CloudControlsPartnerMonitoringCallOptions

	// The gRPC API client.
	cloudControlsPartnerMonitoringClient cloudcontrolspartnerpb.CloudControlsPartnerMonitoringClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewCloudControlsPartnerMonitoringClient creates a new cloud controls partner monitoring client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service describing handlers for resources
func NewCloudControlsPartnerMonitoringClient(ctx context.Context, opts ...option.ClientOption) (*CloudControlsPartnerMonitoringClient, error) {
	clientOpts := defaultCloudControlsPartnerMonitoringGRPCClientOptions()
	if newCloudControlsPartnerMonitoringClientHook != nil {
		hookOpts, err := newCloudControlsPartnerMonitoringClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := CloudControlsPartnerMonitoringClient{CallOptions: defaultCloudControlsPartnerMonitoringCallOptions()}

	c := &cloudControlsPartnerMonitoringGRPCClient{
		connPool:                             connPool,
		cloudControlsPartnerMonitoringClient: cloudcontrolspartnerpb.NewCloudControlsPartnerMonitoringClient(connPool),
		CallOptions:                          &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *cloudControlsPartnerMonitoringGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *cloudControlsPartnerMonitoringGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *cloudControlsPartnerMonitoringGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type cloudControlsPartnerMonitoringRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing CloudControlsPartnerMonitoringClient
	CallOptions **CloudControlsPartnerMonitoringCallOptions
}

// NewCloudControlsPartnerMonitoringRESTClient creates a new cloud controls partner monitoring rest client.
//
// Service describing handlers for resources
func NewCloudControlsPartnerMonitoringRESTClient(ctx context.Context, opts ...option.ClientOption) (*CloudControlsPartnerMonitoringClient, error) {
	clientOpts := append(defaultCloudControlsPartnerMonitoringRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultCloudControlsPartnerMonitoringRESTCallOptions()
	c := &cloudControlsPartnerMonitoringRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	return &CloudControlsPartnerMonitoringClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultCloudControlsPartnerMonitoringRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://cloudcontrolspartner.googleapis.com"),
		internaloption.WithDefaultEndpointTemplate("https://cloudcontrolspartner.UNIVERSE_DOMAIN"),
		internaloption.WithDefaultMTLSEndpoint("https://cloudcontrolspartner.mtls.googleapis.com"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://cloudcontrolspartner.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *cloudControlsPartnerMonitoringRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *cloudControlsPartnerMonitoringRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *cloudControlsPartnerMonitoringRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *cloudControlsPartnerMonitoringGRPCClient) ListViolations(ctx context.Context, req *cloudcontrolspartnerpb.ListViolationsRequest, opts ...gax.CallOption) *ViolationIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListViolations[0:len((*c.CallOptions).ListViolations):len((*c.CallOptions).ListViolations)], opts...)
	it := &ViolationIterator{}
	req = proto.Clone(req).(*cloudcontrolspartnerpb.ListViolationsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*cloudcontrolspartnerpb.Violation, string, error) {
		resp := &cloudcontrolspartnerpb.ListViolationsResponse{}
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
			resp, err = c.cloudControlsPartnerMonitoringClient.ListViolations(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetViolations(), resp.GetNextPageToken(), nil
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

func (c *cloudControlsPartnerMonitoringGRPCClient) GetViolation(ctx context.Context, req *cloudcontrolspartnerpb.GetViolationRequest, opts ...gax.CallOption) (*cloudcontrolspartnerpb.Violation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetViolation[0:len((*c.CallOptions).GetViolation):len((*c.CallOptions).GetViolation)], opts...)
	var resp *cloudcontrolspartnerpb.Violation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.cloudControlsPartnerMonitoringClient.GetViolation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ListViolations lists Violations for a workload
// Callers may also choose to read across multiple Customers or for a single
// customer as per
// AIP-159 (at https://google.aip.dev/159) by using ‘-’ (the hyphen or dash
// character) as a wildcard character instead of {customer} & {workload}.
// Format:
// organizations/{organization}/locations/{location}/customers/{customer}/workloads/{workload}
func (c *cloudControlsPartnerMonitoringRESTClient) ListViolations(ctx context.Context, req *cloudcontrolspartnerpb.ListViolationsRequest, opts ...gax.CallOption) *ViolationIterator {
	it := &ViolationIterator{}
	req = proto.Clone(req).(*cloudcontrolspartnerpb.ListViolationsRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*cloudcontrolspartnerpb.Violation, string, error) {
		resp := &cloudcontrolspartnerpb.ListViolationsResponse{}
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
		baseUrl.Path += fmt.Sprintf("/v1/%v/violations", req.GetParent())

		params := url.Values{}
		params.Add("$alt", "json;enum-encoding=int")
		if req.GetFilter() != "" {
			params.Add("filter", fmt.Sprintf("%v", req.GetFilter()))
		}
		if req.GetInterval().GetEndTime() != nil {
			endTime, err := protojson.Marshal(req.GetInterval().GetEndTime())
			if err != nil {
				return nil, "", err
			}
			params.Add("interval.endTime", string(endTime[1:len(endTime)-1]))
		}
		if req.GetInterval().GetStartTime() != nil {
			startTime, err := protojson.Marshal(req.GetInterval().GetStartTime())
			if err != nil {
				return nil, "", err
			}
			params.Add("interval.startTime", string(startTime[1:len(startTime)-1]))
		}
		if req.GetOrderBy() != "" {
			params.Add("orderBy", fmt.Sprintf("%v", req.GetOrderBy()))
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

			httpRsp, err := c.httpClient.Do(httpReq)
			if err != nil {
				return err
			}
			defer httpRsp.Body.Close()

			if err = googleapi.CheckResponse(httpRsp); err != nil {
				return err
			}

			buf, err := io.ReadAll(httpRsp.Body)
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
		return resp.GetViolations(), resp.GetNextPageToken(), nil
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

// GetViolation gets details of a single Violation.
func (c *cloudControlsPartnerMonitoringRESTClient) GetViolation(ctx context.Context, req *cloudcontrolspartnerpb.GetViolationRequest, opts ...gax.CallOption) (*cloudcontrolspartnerpb.Violation, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).GetViolation[0:len((*c.CallOptions).GetViolation):len((*c.CallOptions).GetViolation)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &cloudcontrolspartnerpb.Violation{}
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

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := io.ReadAll(httpRsp.Body)
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