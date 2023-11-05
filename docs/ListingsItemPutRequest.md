# ListingsItemPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductType** | **string** | The Amazon product type of the listings item. | 
**Requirements** | Pointer to **string** | The name of the requirements set for the provided data. | [optional] 
**Attributes** | **map[string]interface{}** | JSON object containing structured listings item attribute data keyed by attribute name. | 

## Methods

### NewListingsItemPutRequest

`func NewListingsItemPutRequest(productType string, attributes map[string]interface{}, ) *ListingsItemPutRequest`

NewListingsItemPutRequest instantiates a new ListingsItemPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListingsItemPutRequestWithDefaults

`func NewListingsItemPutRequestWithDefaults() *ListingsItemPutRequest`

NewListingsItemPutRequestWithDefaults instantiates a new ListingsItemPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductType

`func (o *ListingsItemPutRequest) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *ListingsItemPutRequest) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *ListingsItemPutRequest) SetProductType(v string)`

SetProductType sets ProductType field to given value.


### GetRequirements

`func (o *ListingsItemPutRequest) GetRequirements() string`

GetRequirements returns the Requirements field if non-nil, zero value otherwise.

### GetRequirementsOk

`func (o *ListingsItemPutRequest) GetRequirementsOk() (*string, bool)`

GetRequirementsOk returns a tuple with the Requirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirements

`func (o *ListingsItemPutRequest) SetRequirements(v string)`

SetRequirements sets Requirements field to given value.

### HasRequirements

`func (o *ListingsItemPutRequest) HasRequirements() bool`

HasRequirements returns a boolean if a field has been set.

### GetAttributes

`func (o *ListingsItemPutRequest) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ListingsItemPutRequest) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ListingsItemPutRequest) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


