// Copyright 2023 Google LLC
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

package asset

import (
	"context"
	"time"

	assetpb "cloud.google.com/go/asset/apiv1/assetpb"
	"cloud.google.com/go/longrunning"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
)

// AnalyzeIamPolicyLongrunningOperation manages a long-running operation from AnalyzeIamPolicyLongrunning.
type AnalyzeIamPolicyLongrunningOperation struct {
	lro      *longrunning.Operation
	pollPath string
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *AnalyzeIamPolicyLongrunningOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*assetpb.AnalyzeIamPolicyLongrunningResponse, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp assetpb.AnalyzeIamPolicyLongrunningResponse
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *AnalyzeIamPolicyLongrunningOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*assetpb.AnalyzeIamPolicyLongrunningResponse, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp assetpb.AnalyzeIamPolicyLongrunningResponse
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *AnalyzeIamPolicyLongrunningOperation) Metadata() (*assetpb.AnalyzeIamPolicyLongrunningMetadata, error) {
	var meta assetpb.AnalyzeIamPolicyLongrunningMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *AnalyzeIamPolicyLongrunningOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *AnalyzeIamPolicyLongrunningOperation) Name() string {
	return op.lro.Name()
}

// ExportAssetsOperation manages a long-running operation from ExportAssets.
type ExportAssetsOperation struct {
	lro      *longrunning.Operation
	pollPath string
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *ExportAssetsOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*assetpb.ExportAssetsResponse, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp assetpb.ExportAssetsResponse
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *ExportAssetsOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*assetpb.ExportAssetsResponse, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp assetpb.ExportAssetsResponse
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *ExportAssetsOperation) Metadata() (*assetpb.ExportAssetsRequest, error) {
	var meta assetpb.ExportAssetsRequest
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *ExportAssetsOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *ExportAssetsOperation) Name() string {
	return op.lro.Name()
}

// AnalyzeOrgPoliciesResponse_OrgPolicyResultIterator manages a stream of *assetpb.AnalyzeOrgPoliciesResponse_OrgPolicyResult.
type AnalyzeOrgPoliciesResponse_OrgPolicyResultIterator struct {
	items    []*assetpb.AnalyzeOrgPoliciesResponse_OrgPolicyResult
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*assetpb.AnalyzeOrgPoliciesResponse_OrgPolicyResult, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *AnalyzeOrgPoliciesResponse_OrgPolicyResultIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *AnalyzeOrgPoliciesResponse_OrgPolicyResultIterator) Next() (*assetpb.AnalyzeOrgPoliciesResponse_OrgPolicyResult, error) {
	var item *assetpb.AnalyzeOrgPoliciesResponse_OrgPolicyResult
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *AnalyzeOrgPoliciesResponse_OrgPolicyResultIterator) bufLen() int {
	return len(it.items)
}

func (it *AnalyzeOrgPoliciesResponse_OrgPolicyResultIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// AnalyzeOrgPolicyGovernedAssetsResponse_GovernedAssetIterator manages a stream of *assetpb.AnalyzeOrgPolicyGovernedAssetsResponse_GovernedAsset.
type AnalyzeOrgPolicyGovernedAssetsResponse_GovernedAssetIterator struct {
	items    []*assetpb.AnalyzeOrgPolicyGovernedAssetsResponse_GovernedAsset
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*assetpb.AnalyzeOrgPolicyGovernedAssetsResponse_GovernedAsset, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *AnalyzeOrgPolicyGovernedAssetsResponse_GovernedAssetIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *AnalyzeOrgPolicyGovernedAssetsResponse_GovernedAssetIterator) Next() (*assetpb.AnalyzeOrgPolicyGovernedAssetsResponse_GovernedAsset, error) {
	var item *assetpb.AnalyzeOrgPolicyGovernedAssetsResponse_GovernedAsset
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *AnalyzeOrgPolicyGovernedAssetsResponse_GovernedAssetIterator) bufLen() int {
	return len(it.items)
}

func (it *AnalyzeOrgPolicyGovernedAssetsResponse_GovernedAssetIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// AnalyzeOrgPolicyGovernedContainersResponse_GovernedContainerIterator manages a stream of *assetpb.AnalyzeOrgPolicyGovernedContainersResponse_GovernedContainer.
type AnalyzeOrgPolicyGovernedContainersResponse_GovernedContainerIterator struct {
	items    []*assetpb.AnalyzeOrgPolicyGovernedContainersResponse_GovernedContainer
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*assetpb.AnalyzeOrgPolicyGovernedContainersResponse_GovernedContainer, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *AnalyzeOrgPolicyGovernedContainersResponse_GovernedContainerIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *AnalyzeOrgPolicyGovernedContainersResponse_GovernedContainerIterator) Next() (*assetpb.AnalyzeOrgPolicyGovernedContainersResponse_GovernedContainer, error) {
	var item *assetpb.AnalyzeOrgPolicyGovernedContainersResponse_GovernedContainer
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *AnalyzeOrgPolicyGovernedContainersResponse_GovernedContainerIterator) bufLen() int {
	return len(it.items)
}

func (it *AnalyzeOrgPolicyGovernedContainersResponse_GovernedContainerIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// AssetIterator manages a stream of *assetpb.Asset.
type AssetIterator struct {
	items    []*assetpb.Asset
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*assetpb.Asset, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *AssetIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *AssetIterator) Next() (*assetpb.Asset, error) {
	var item *assetpb.Asset
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *AssetIterator) bufLen() int {
	return len(it.items)
}

func (it *AssetIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// IamPolicySearchResultIterator manages a stream of *assetpb.IamPolicySearchResult.
type IamPolicySearchResultIterator struct {
	items    []*assetpb.IamPolicySearchResult
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*assetpb.IamPolicySearchResult, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *IamPolicySearchResultIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *IamPolicySearchResultIterator) Next() (*assetpb.IamPolicySearchResult, error) {
	var item *assetpb.IamPolicySearchResult
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *IamPolicySearchResultIterator) bufLen() int {
	return len(it.items)
}

func (it *IamPolicySearchResultIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// ResourceSearchResultIterator manages a stream of *assetpb.ResourceSearchResult.
type ResourceSearchResultIterator struct {
	items    []*assetpb.ResourceSearchResult
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*assetpb.ResourceSearchResult, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *ResourceSearchResultIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *ResourceSearchResultIterator) Next() (*assetpb.ResourceSearchResult, error) {
	var item *assetpb.ResourceSearchResult
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *ResourceSearchResultIterator) bufLen() int {
	return len(it.items)
}

func (it *ResourceSearchResultIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// SavedQueryIterator manages a stream of *assetpb.SavedQuery.
type SavedQueryIterator struct {
	items    []*assetpb.SavedQuery
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*assetpb.SavedQuery, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *SavedQueryIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *SavedQueryIterator) Next() (*assetpb.SavedQuery, error) {
	var item *assetpb.SavedQuery
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *SavedQueryIterator) bufLen() int {
	return len(it.items)
}

func (it *SavedQueryIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}