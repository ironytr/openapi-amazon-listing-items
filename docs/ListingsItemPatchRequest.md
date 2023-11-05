# ListingsItemPatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductType** | **string** | The Amazon product type of the listings item. | 
**Patches** | [**[]PatchOperation**](PatchOperation.md) | One or more JSON Patch operations to perform on the listings item. | 

## Methods

### NewListingsItemPatchRequest

`func NewListingsItemPatchRequest(productType string, patches []PatchOperation, ) *ListingsItemPatchRequest`

NewListingsItemPatchRequest instantiates a new ListingsItemPatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListingsItemPatchRequestWithDefaults

`func NewListingsItemPatchRequestWithDefaults() *ListingsItemPatchRequest`

NewListingsItemPatchRequestWithDefaults instantiates a new ListingsItemPatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductType

`func (o *ListingsItemPatchRequest) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *ListingsItemPatchRequest) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *ListingsItemPatchRequest) SetProductType(v string)`

SetProductType sets ProductType field to given value.


### GetPatches

`func (o *ListingsItemPatchRequest) GetPatches() []PatchOperation`

GetPatches returns the Patches field if non-nil, zero value otherwise.

### GetPatchesOk

`func (o *ListingsItemPatchRequest) GetPatchesOk() (*[]PatchOperation, bool)`

GetPatchesOk returns a tuple with the Patches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatches

`func (o *ListingsItemPatchRequest) SetPatches(v []PatchOperation)`

SetPatches sets Patches field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


