package operationalinsights

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/operationalinsights/mgmt/2015-03-20/operationalinsights"

// PurgeState enumerates the values for purge state.
type PurgeState string

const (
	// Completed ...
	Completed PurgeState = "completed"
	// Pending ...
	Pending PurgeState = "pending"
)

// PossiblePurgeStateValues returns an array of possible values for the PurgeState const type.
func PossiblePurgeStateValues() []PurgeState {
	return []PurgeState{Completed, Pending}
}

// SearchSortEnum enumerates the values for search sort enum.
type SearchSortEnum string

const (
	// Asc ...
	Asc SearchSortEnum = "asc"
	// Desc ...
	Desc SearchSortEnum = "desc"
)

// PossibleSearchSortEnumValues returns an array of possible values for the SearchSortEnum const type.
func PossibleSearchSortEnumValues() []SearchSortEnum {
	return []SearchSortEnum{Asc, Desc}
}

// StorageInsightState enumerates the values for storage insight state.
type StorageInsightState string

const (
	// ERROR ...
	ERROR StorageInsightState = "ERROR"
	// OK ...
	OK StorageInsightState = "OK"
)

// PossibleStorageInsightStateValues returns an array of possible values for the StorageInsightState const type.
func PossibleStorageInsightStateValues() []StorageInsightState {
	return []StorageInsightState{ERROR, OK}
}

// CoreSummary the core summary of a search.
type CoreSummary struct {
	// Status - The status of a core summary.
	Status *string `json:"status,omitempty"`
	// NumberOfDocuments - The number of documents of a core summary.
	NumberOfDocuments *int64 `json:"numberOfDocuments,omitempty"`
}

// LinkTarget metadata for a workspace that isn't linked to an Azure subscription.
type LinkTarget struct {
	// CustomerID - The GUID that uniquely identifies the workspace.
	CustomerID *string `json:"customerId,omitempty"`
	// DisplayName - The display name of the workspace.
	DisplayName *string `json:"accountName,omitempty"`
	// WorkspaceName - The DNS valid workspace name.
	WorkspaceName *string `json:"workspaceName,omitempty"`
	// Location - The location of the workspace.
	Location *string `json:"location,omitempty"`
}

// ListLinkTarget ...
type ListLinkTarget struct {
	autorest.Response `json:"-"`
	Value             *[]LinkTarget `json:"value,omitempty"`
}

// Operation supported operation of OperationalInsights resource provider.
type Operation struct {
	// Name - Operation name: {provider}/{resource}/{operation}
	Name *string `json:"name,omitempty"`
	// Display - Display metadata associated with the operation.
	Display *OperationDisplay `json:"display,omitempty"`
}

// OperationDisplay display metadata associated with the operation.
type OperationDisplay struct {
	// Provider - Service provider: OperationalInsights.
	Provider *string `json:"provider,omitempty"`
	// Resource - Resource on which the operation is performed etc.
	Resource *string `json:"resource,omitempty"`
	// Operation - Type of operation: get, read, delete, etc.
	Operation *string `json:"operation,omitempty"`
}

// OperationListResult result of the request to list OperationalInsights operations.
type OperationListResult struct {
	autorest.Response `json:"-"`
	// Value - List of operations supported by the OperationalInsights resource provider.
	Value *[]Operation `json:"value,omitempty"`
}

// ProxyResource common properties of proxy resource.
type ProxyResource struct {
	// ID - READ-ONLY; Resource ID.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Resource name.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Resource type.
	Type *string `json:"type,omitempty"`
	// Tags - Resource tags
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for ProxyResource.
func (pr ProxyResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if pr.Tags != nil {
		objectMap["tags"] = pr.Tags
	}
	return json.Marshal(objectMap)
}

// Resource the resource definition.
type Resource struct {
	// ID - READ-ONLY; Resource Id
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Resource name
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Resource type
	Type *string `json:"type,omitempty"`
	// Location - Resource location
	Location *string `json:"location,omitempty"`
	// Tags - Resource tags
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if r.Location != nil {
		objectMap["location"] = r.Location
	}
	if r.Tags != nil {
		objectMap["tags"] = r.Tags
	}
	return json.Marshal(objectMap)
}

// SavedSearch value object for saved search results.
type SavedSearch struct {
	autorest.Response `json:"-"`
	// ID - READ-ONLY; The id of the saved search.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the saved search.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the saved search.
	Type *string `json:"type,omitempty"`
	// ETag - The ETag of the saved search.
	ETag *string `json:"eTag,omitempty"`
	// SavedSearchProperties - The properties of the saved search.
	*SavedSearchProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for SavedSearch.
func (ss SavedSearch) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ss.ETag != nil {
		objectMap["eTag"] = ss.ETag
	}
	if ss.SavedSearchProperties != nil {
		objectMap["properties"] = ss.SavedSearchProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for SavedSearch struct.
func (ss *SavedSearch) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				ss.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				ss.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				ss.Type = &typeVar
			}
		case "eTag":
			if v != nil {
				var eTag string
				err = json.Unmarshal(*v, &eTag)
				if err != nil {
					return err
				}
				ss.ETag = &eTag
			}
		case "properties":
			if v != nil {
				var savedSearchProperties SavedSearchProperties
				err = json.Unmarshal(*v, &savedSearchProperties)
				if err != nil {
					return err
				}
				ss.SavedSearchProperties = &savedSearchProperties
			}
		}
	}

	return nil
}

// SavedSearchesListResult the saved search list operation response.
type SavedSearchesListResult struct {
	autorest.Response `json:"-"`
	// Metadata - The metadata from search results.
	Metadata *SearchMetadata `json:"__metadata,omitempty"`
	// Value - The array of result values.
	Value *[]SavedSearch `json:"value,omitempty"`
}

// SavedSearchProperties value object for saved search results.
type SavedSearchProperties struct {
	// Category - The category of the saved search. This helps the user to find a saved search faster.
	Category *string `json:"category,omitempty"`
	// DisplayName - Saved search display name.
	DisplayName *string `json:"displayName,omitempty"`
	// Query - The query expression for the saved search. Please see https://docs.microsoft.com/en-us/azure/log-analytics/log-analytics-search-reference for reference.
	Query *string `json:"query,omitempty"`
	// Version - The version number of the query language. The current version is 2 and is the default.
	Version *int64 `json:"version,omitempty"`
	// Tags - The tags attached to the saved search.
	Tags *[]Tag `json:"tags,omitempty"`
}

// SearchError details for a search error.
type SearchError struct {
	// Type - The error type.
	Type *string `json:"type,omitempty"`
	// Message - The error message.
	Message *string `json:"message,omitempty"`
}

// SearchGetSchemaResponse the get schema operation response.
type SearchGetSchemaResponse struct {
	autorest.Response `json:"-"`
	// Metadata - The metadata from search results.
	Metadata *SearchMetadata `json:"metadata,omitempty"`
	// Value - The array of result values.
	Value *[]SearchSchemaValue `json:"value,omitempty"`
}

// SearchHighlight highlight details.
type SearchHighlight struct {
	// Pre - The string that is put before a matched result.
	Pre *string `json:"pre,omitempty"`
	// Post - The string that is put after a matched result.
	Post *string `json:"post,omitempty"`
}

// SearchMetadata metadata for search results.
type SearchMetadata struct {
	// SearchID - The request id of the search.
	SearchID *string `json:"requestId,omitempty"`
	// ResultType - The search result type.
	ResultType *string `json:"resultType,omitempty"`
	// Total - The total number of search results.
	Total *int64 `json:"total,omitempty"`
	// Top - The number of top search results.
	Top *int64 `json:"top,omitempty"`
	// ID - The id of the search results request.
	ID *string `json:"id,omitempty"`
	// CoreSummaries - The core summaries.
	CoreSummaries *[]CoreSummary `json:"coreSummaries,omitempty"`
	// Status - The status of the search results.
	Status *string `json:"status,omitempty"`
	// StartTime - The start time for the search.
	StartTime *date.Time `json:"startTime,omitempty"`
	// LastUpdated - The time of last update.
	LastUpdated *date.Time `json:"lastUpdated,omitempty"`
	// ETag - The ETag of the search results.
	ETag *string `json:"eTag,omitempty"`
	// Sort - How the results are sorted.
	Sort *[]SearchSort `json:"sort,omitempty"`
	// RequestTime - The request time.
	RequestTime *int64 `json:"requestTime,omitempty"`
	// AggregatedValueField - The aggregated value field.
	AggregatedValueField *string `json:"aggregatedValueField,omitempty"`
	// AggregatedGroupingFields - The aggregated grouping fields.
	AggregatedGroupingFields *string `json:"aggregatedGroupingFields,omitempty"`
	// Sum - The sum of all aggregates returned in the result set.
	Sum *int64 `json:"sum,omitempty"`
	// Max - The max of all aggregates returned in the result set.
	Max *int64 `json:"max,omitempty"`
	// Schema - The schema.
	Schema *SearchMetadataSchema `json:"schema,omitempty"`
}

// SearchMetadataSchema schema metadata for search.
type SearchMetadataSchema struct {
	// Name - The name of the metadata schema.
	Name *string `json:"name,omitempty"`
	// Version - The version of the metadata schema.
	Version *int32 `json:"version,omitempty"`
}

// SearchParameters parameters specifying the search query and range.
type SearchParameters struct {
	// Top - The number to get from the top.
	Top *int64 `json:"top,omitempty"`
	// Highlight - The highlight that looks for all occurrences of a string.
	Highlight *SearchHighlight `json:"highlight,omitempty"`
	// Query - The query to search.
	Query *string `json:"query,omitempty"`
	// Start - The start date filter, so the only query results returned are after this date.
	Start *date.Time `json:"start,omitempty"`
	// End - The end date filter, so the only query results returned are before this date.
	End *date.Time `json:"end,omitempty"`
}

// SearchResultsResponse the get search result operation response.
type SearchResultsResponse struct {
	autorest.Response `json:"-"`
	// ID - READ-ONLY; The id of the search, which includes the full url.
	ID *string `json:"id,omitempty"`
	// Metadata - The metadata from search results.
	Metadata *SearchMetadata `json:"metaData,omitempty"`
	// Value - The array of result values.
	Value *[]interface{} `json:"value,omitempty"`
	// Error - The error.
	Error *SearchError `json:"error,omitempty"`
}

// SearchSchemaValue value object for schema results.
type SearchSchemaValue struct {
	// Name - The name of the schema.
	Name *string `json:"name,omitempty"`
	// DisplayName - The display name of the schema.
	DisplayName *string `json:"displayName,omitempty"`
	// Type - The type.
	Type *string `json:"type,omitempty"`
	// Indexed - The boolean that indicates the field is searchable as free text.
	Indexed *bool `json:"indexed,omitempty"`
	// Stored - The boolean that indicates whether or not the field is stored.
	Stored *bool `json:"stored,omitempty"`
	// Facet - The boolean that indicates whether or not the field is a facet.
	Facet *bool `json:"facet,omitempty"`
	// OwnerType - The array of workflows containing the field.
	OwnerType *[]string `json:"ownerType,omitempty"`
}

// SearchSort the sort parameters for search.
type SearchSort struct {
	// Name - The name of the field the search query is sorted on.
	Name *string `json:"name,omitempty"`
	// Order - The sort order of the search. Possible values include: 'Asc', 'Desc'
	Order SearchSortEnum `json:"order,omitempty"`
}

// SharedKeys the shared keys for a workspace.
type SharedKeys struct {
	autorest.Response `json:"-"`
	// PrimarySharedKey - The primary shared key of a workspace.
	PrimarySharedKey *string `json:"primarySharedKey,omitempty"`
	// SecondarySharedKey - The secondary shared key of a workspace.
	SecondarySharedKey *string `json:"secondarySharedKey,omitempty"`
}

// StorageAccount describes a storage account connection.
type StorageAccount struct {
	// ID - The Azure Resource Manager ID of the storage account resource.
	ID *string `json:"id,omitempty"`
	// Key - The storage account key.
	Key *string `json:"key,omitempty"`
}

// StorageInsight the top level storage insight resource container.
type StorageInsight struct {
	autorest.Response `json:"-"`
	// StorageInsightProperties - Storage insight properties.
	*StorageInsightProperties `json:"properties,omitempty"`
	// ETag - The ETag of the storage insight.
	ETag *string `json:"eTag,omitempty"`
	// ID - READ-ONLY; Resource ID.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Resource name.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Resource type.
	Type *string `json:"type,omitempty"`
	// Tags - Resource tags
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for StorageInsight.
func (si StorageInsight) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if si.StorageInsightProperties != nil {
		objectMap["properties"] = si.StorageInsightProperties
	}
	if si.ETag != nil {
		objectMap["eTag"] = si.ETag
	}
	if si.Tags != nil {
		objectMap["tags"] = si.Tags
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for StorageInsight struct.
func (si *StorageInsight) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var storageInsightProperties StorageInsightProperties
				err = json.Unmarshal(*v, &storageInsightProperties)
				if err != nil {
					return err
				}
				si.StorageInsightProperties = &storageInsightProperties
			}
		case "eTag":
			if v != nil {
				var eTag string
				err = json.Unmarshal(*v, &eTag)
				if err != nil {
					return err
				}
				si.ETag = &eTag
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				si.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				si.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				si.Type = &typeVar
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				si.Tags = tags
			}
		}
	}

	return nil
}

// StorageInsightListResult the list storage insights operation response.
type StorageInsightListResult struct {
	autorest.Response `json:"-"`
	// Value - A list of storage insight items.
	Value *[]StorageInsight `json:"value,omitempty"`
	// OdataNextLink - The link (url) to the next page of results.
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// StorageInsightListResultIterator provides access to a complete listing of StorageInsight values.
type StorageInsightListResultIterator struct {
	i    int
	page StorageInsightListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *StorageInsightListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StorageInsightListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *StorageInsightListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter StorageInsightListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter StorageInsightListResultIterator) Response() StorageInsightListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter StorageInsightListResultIterator) Value() StorageInsight {
	if !iter.page.NotDone() {
		return StorageInsight{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the StorageInsightListResultIterator type.
func NewStorageInsightListResultIterator(page StorageInsightListResultPage) StorageInsightListResultIterator {
	return StorageInsightListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (silr StorageInsightListResult) IsEmpty() bool {
	return silr.Value == nil || len(*silr.Value) == 0
}

// storageInsightListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (silr StorageInsightListResult) storageInsightListResultPreparer(ctx context.Context) (*http.Request, error) {
	if silr.OdataNextLink == nil || len(to.String(silr.OdataNextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(silr.OdataNextLink)))
}

// StorageInsightListResultPage contains a page of StorageInsight values.
type StorageInsightListResultPage struct {
	fn   func(context.Context, StorageInsightListResult) (StorageInsightListResult, error)
	silr StorageInsightListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *StorageInsightListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StorageInsightListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.silr)
	if err != nil {
		return err
	}
	page.silr = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *StorageInsightListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page StorageInsightListResultPage) NotDone() bool {
	return !page.silr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page StorageInsightListResultPage) Response() StorageInsightListResult {
	return page.silr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page StorageInsightListResultPage) Values() []StorageInsight {
	if page.silr.IsEmpty() {
		return nil
	}
	return *page.silr.Value
}

// Creates a new instance of the StorageInsightListResultPage type.
func NewStorageInsightListResultPage(getNextPage func(context.Context, StorageInsightListResult) (StorageInsightListResult, error)) StorageInsightListResultPage {
	return StorageInsightListResultPage{fn: getNextPage}
}

// StorageInsightProperties storage insight properties.
type StorageInsightProperties struct {
	// Containers - The names of the blob containers that the workspace should read
	Containers *[]string `json:"containers,omitempty"`
	// Tables - The names of the Azure tables that the workspace should read
	Tables *[]string `json:"tables,omitempty"`
	// StorageAccount - The storage account connection details
	StorageAccount *StorageAccount `json:"storageAccount,omitempty"`
	// Status - READ-ONLY; The status of the storage insight
	Status *StorageInsightStatus `json:"status,omitempty"`
}

// StorageInsightStatus the status of the storage insight.
type StorageInsightStatus struct {
	// State - The state of the storage insight connection to the workspace. Possible values include: 'OK', 'ERROR'
	State StorageInsightState `json:"state,omitempty"`
	// Description - Description of the state of the storage insight.
	Description *string `json:"description,omitempty"`
}

// Tag a tag of a saved search.
type Tag struct {
	// Name - The tag name.
	Name *string `json:"name,omitempty"`
	// Value - The tag value.
	Value *string `json:"value,omitempty"`
}

// WorkspacePurgeBody describes the body of a purge request for an App Insights Workspace
type WorkspacePurgeBody struct {
	// Table - Table from which to purge data.
	Table *string `json:"table,omitempty"`
	// Filters - The set of columns and filters (queries) to run over them to purge the resulting data.
	Filters *[]WorkspacePurgeBodyFilters `json:"filters,omitempty"`
}

// WorkspacePurgeBodyFilters user-defined filters to return data which will be purged from the table.
type WorkspacePurgeBodyFilters struct {
	// Column - The column of the table over which the given query should run
	Column *string `json:"column,omitempty"`
	// Operator - A query operator to evaluate over the provided column and value(s).
	Operator *string `json:"operator,omitempty"`
	// Value - the value for the operator to function over. This can be a number (e.g., > 100), a string (timestamp >= '2017-09-01') or array of values.
	Value interface{} `json:"value,omitempty"`
	// Key - When filtering over custom dimensions, this key will be used as the name of the custom dimension.
	Key *string `json:"key,omitempty"`
}

// WorkspacePurgeResponse response containing operationId for a specific purge action.
type WorkspacePurgeResponse struct {
	autorest.Response `json:"-"`
	// OperationID - Id to use when querying for status for a particular purge operation.
	OperationID *string `json:"operationId,omitempty"`
}

// WorkspacePurgeStatusResponse response containing status for a specific purge operation.
type WorkspacePurgeStatusResponse struct {
	autorest.Response `json:"-"`
	// Status - Status of the operation represented by the requested Id. Possible values include: 'Pending', 'Completed'
	Status PurgeState `json:"status,omitempty"`
}

// WorkspacesGetSearchResultsFuture an abstraction for monitoring and retrieving the results of a
// long-running operation.
type WorkspacesGetSearchResultsFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *WorkspacesGetSearchResultsFuture) Result(client WorkspacesClient) (srr SearchResultsResponse, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationalinsights.WorkspacesGetSearchResultsFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("operationalinsights.WorkspacesGetSearchResultsFuture")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if srr.Response.Response, err = future.GetResult(sender); err == nil && srr.Response.Response.StatusCode != http.StatusNoContent {
		srr, err = client.GetSearchResultsResponder(srr.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "operationalinsights.WorkspacesGetSearchResultsFuture", "Result", srr.Response.Response, "Failure responding to request")
		}
	}
	return
}
