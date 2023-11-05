# ListingsItemSubmissionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sku** | **string** | A selling partner provided identifier for an Amazon listing. | 
**Status** | **string** | The status of the listings item submission. | 
**SubmissionId** | **string** | The unique identifier of the listings item submission. | 
**Issues** | Pointer to [**[]Issue**](Issue.md) | Listings item issues related to the listings item submission. | [optional] 

## Methods

### NewListingsItemSubmissionResponse

`func NewListingsItemSubmissionResponse(sku string, status string, submissionId string, ) *ListingsItemSubmissionResponse`

NewListingsItemSubmissionResponse instantiates a new ListingsItemSubmissionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListingsItemSubmissionResponseWithDefaults

`func NewListingsItemSubmissionResponseWithDefaults() *ListingsItemSubmissionResponse`

NewListingsItemSubmissionResponseWithDefaults instantiates a new ListingsItemSubmissionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSku

`func (o *ListingsItemSubmissionResponse) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *ListingsItemSubmissionResponse) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *ListingsItemSubmissionResponse) SetSku(v string)`

SetSku sets Sku field to given value.


### GetStatus

`func (o *ListingsItemSubmissionResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListingsItemSubmissionResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListingsItemSubmissionResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSubmissionId

`func (o *ListingsItemSubmissionResponse) GetSubmissionId() string`

GetSubmissionId returns the SubmissionId field if non-nil, zero value otherwise.

### GetSubmissionIdOk

`func (o *ListingsItemSubmissionResponse) GetSubmissionIdOk() (*string, bool)`

GetSubmissionIdOk returns a tuple with the SubmissionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionId

`func (o *ListingsItemSubmissionResponse) SetSubmissionId(v string)`

SetSubmissionId sets SubmissionId field to given value.


### GetIssues

`func (o *ListingsItemSubmissionResponse) GetIssues() []Issue`

GetIssues returns the Issues field if non-nil, zero value otherwise.

### GetIssuesOk

`func (o *ListingsItemSubmissionResponse) GetIssuesOk() (*[]Issue, bool)`

GetIssuesOk returns a tuple with the Issues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssues

`func (o *ListingsItemSubmissionResponse) SetIssues(v []Issue)`

SetIssues sets Issues field to given value.

### HasIssues

`func (o *ListingsItemSubmissionResponse) HasIssues() bool`

HasIssues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


