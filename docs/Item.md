# Item

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sku** | **string** | A selling partner provided identifier for an Amazon listing. | 
**Summaries** | Pointer to [**[]ItemSummaryByMarketplace**](ItemSummaryByMarketplace.md) | Summary details of a listings item. | [optional] 
**Attributes** | Pointer to **map[string]interface{}** | JSON object containing structured listings item attribute data keyed by attribute name. | [optional] 
**Issues** | Pointer to [**[]Issue**](Issue.md) | Issues associated with the listings item. | [optional] 
**Offers** | Pointer to [**[]ItemOfferByMarketplace**](ItemOfferByMarketplace.md) | Offer details for the listings item. | [optional] 
**FulfillmentAvailability** | Pointer to [**[]FulfillmentAvailability**](FulfillmentAvailability.md) | Fulfillment availability for the listings item. | [optional] 
**Procurement** | Pointer to [**[]ItemProcurement**](ItemProcurement.md) | Vendor procurement information for the listings item. | [optional] 

## Methods

### NewItem

`func NewItem(sku string, ) *Item`

NewItem instantiates a new Item object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemWithDefaults

`func NewItemWithDefaults() *Item`

NewItemWithDefaults instantiates a new Item object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSku

`func (o *Item) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *Item) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *Item) SetSku(v string)`

SetSku sets Sku field to given value.


### GetSummaries

`func (o *Item) GetSummaries() []ItemSummaryByMarketplace`

GetSummaries returns the Summaries field if non-nil, zero value otherwise.

### GetSummariesOk

`func (o *Item) GetSummariesOk() (*[]ItemSummaryByMarketplace, bool)`

GetSummariesOk returns a tuple with the Summaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaries

`func (o *Item) SetSummaries(v []ItemSummaryByMarketplace)`

SetSummaries sets Summaries field to given value.

### HasSummaries

`func (o *Item) HasSummaries() bool`

HasSummaries returns a boolean if a field has been set.

### GetAttributes

`func (o *Item) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Item) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Item) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Item) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetIssues

`func (o *Item) GetIssues() []Issue`

GetIssues returns the Issues field if non-nil, zero value otherwise.

### GetIssuesOk

`func (o *Item) GetIssuesOk() (*[]Issue, bool)`

GetIssuesOk returns a tuple with the Issues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssues

`func (o *Item) SetIssues(v []Issue)`

SetIssues sets Issues field to given value.

### HasIssues

`func (o *Item) HasIssues() bool`

HasIssues returns a boolean if a field has been set.

### GetOffers

`func (o *Item) GetOffers() []ItemOfferByMarketplace`

GetOffers returns the Offers field if non-nil, zero value otherwise.

### GetOffersOk

`func (o *Item) GetOffersOk() (*[]ItemOfferByMarketplace, bool)`

GetOffersOk returns a tuple with the Offers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffers

`func (o *Item) SetOffers(v []ItemOfferByMarketplace)`

SetOffers sets Offers field to given value.

### HasOffers

`func (o *Item) HasOffers() bool`

HasOffers returns a boolean if a field has been set.

### GetFulfillmentAvailability

`func (o *Item) GetFulfillmentAvailability() []FulfillmentAvailability`

GetFulfillmentAvailability returns the FulfillmentAvailability field if non-nil, zero value otherwise.

### GetFulfillmentAvailabilityOk

`func (o *Item) GetFulfillmentAvailabilityOk() (*[]FulfillmentAvailability, bool)`

GetFulfillmentAvailabilityOk returns a tuple with the FulfillmentAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentAvailability

`func (o *Item) SetFulfillmentAvailability(v []FulfillmentAvailability)`

SetFulfillmentAvailability sets FulfillmentAvailability field to given value.

### HasFulfillmentAvailability

`func (o *Item) HasFulfillmentAvailability() bool`

HasFulfillmentAvailability returns a boolean if a field has been set.

### GetProcurement

`func (o *Item) GetProcurement() []ItemProcurement`

GetProcurement returns the Procurement field if non-nil, zero value otherwise.

### GetProcurementOk

`func (o *Item) GetProcurementOk() (*[]ItemProcurement, bool)`

GetProcurementOk returns a tuple with the Procurement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcurement

`func (o *Item) SetProcurement(v []ItemProcurement)`

SetProcurement sets Procurement field to given value.

### HasProcurement

`func (o *Item) HasProcurement() bool`

HasProcurement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


