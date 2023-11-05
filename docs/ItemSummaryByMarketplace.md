# ItemSummaryByMarketplace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MarketplaceId** | **string** | A marketplace identifier. Identifies the Amazon marketplace for the listings item. | 
**Asin** | **string** | Amazon Standard Identification Number (ASIN) of the listings item. | 
**ProductType** | **string** | The Amazon product type of the listings item. | 
**ConditionType** | Pointer to **string** | Identifies the condition of the listings item. | [optional] 
**Status** | **[]string** | Statuses that apply to the listings item. | 
**FnSku** | Pointer to **string** | Fulfillment network stock keeping unit is an identifier used by Amazon fulfillment centers to identify each unique item. | [optional] 
**ItemName** | **string** | Name, or title, associated with an Amazon catalog item. | 
**CreatedDate** | **time.Time** | Date the listings item was created, in ISO 8601 format. | 
**LastUpdatedDate** | **time.Time** | Date the listings item was last updated, in ISO 8601 format. | 
**MainImage** | Pointer to [**ItemImage**](ItemImage.md) |  | [optional] 

## Methods

### NewItemSummaryByMarketplace

`func NewItemSummaryByMarketplace(marketplaceId string, asin string, productType string, status []string, itemName string, createdDate time.Time, lastUpdatedDate time.Time, ) *ItemSummaryByMarketplace`

NewItemSummaryByMarketplace instantiates a new ItemSummaryByMarketplace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemSummaryByMarketplaceWithDefaults

`func NewItemSummaryByMarketplaceWithDefaults() *ItemSummaryByMarketplace`

NewItemSummaryByMarketplaceWithDefaults instantiates a new ItemSummaryByMarketplace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarketplaceId

`func (o *ItemSummaryByMarketplace) GetMarketplaceId() string`

GetMarketplaceId returns the MarketplaceId field if non-nil, zero value otherwise.

### GetMarketplaceIdOk

`func (o *ItemSummaryByMarketplace) GetMarketplaceIdOk() (*string, bool)`

GetMarketplaceIdOk returns a tuple with the MarketplaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceId

`func (o *ItemSummaryByMarketplace) SetMarketplaceId(v string)`

SetMarketplaceId sets MarketplaceId field to given value.


### GetAsin

`func (o *ItemSummaryByMarketplace) GetAsin() string`

GetAsin returns the Asin field if non-nil, zero value otherwise.

### GetAsinOk

`func (o *ItemSummaryByMarketplace) GetAsinOk() (*string, bool)`

GetAsinOk returns a tuple with the Asin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsin

`func (o *ItemSummaryByMarketplace) SetAsin(v string)`

SetAsin sets Asin field to given value.


### GetProductType

`func (o *ItemSummaryByMarketplace) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *ItemSummaryByMarketplace) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *ItemSummaryByMarketplace) SetProductType(v string)`

SetProductType sets ProductType field to given value.


### GetConditionType

`func (o *ItemSummaryByMarketplace) GetConditionType() string`

GetConditionType returns the ConditionType field if non-nil, zero value otherwise.

### GetConditionTypeOk

`func (o *ItemSummaryByMarketplace) GetConditionTypeOk() (*string, bool)`

GetConditionTypeOk returns a tuple with the ConditionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionType

`func (o *ItemSummaryByMarketplace) SetConditionType(v string)`

SetConditionType sets ConditionType field to given value.

### HasConditionType

`func (o *ItemSummaryByMarketplace) HasConditionType() bool`

HasConditionType returns a boolean if a field has been set.

### GetStatus

`func (o *ItemSummaryByMarketplace) GetStatus() []string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ItemSummaryByMarketplace) GetStatusOk() (*[]string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ItemSummaryByMarketplace) SetStatus(v []string)`

SetStatus sets Status field to given value.


### GetFnSku

`func (o *ItemSummaryByMarketplace) GetFnSku() string`

GetFnSku returns the FnSku field if non-nil, zero value otherwise.

### GetFnSkuOk

`func (o *ItemSummaryByMarketplace) GetFnSkuOk() (*string, bool)`

GetFnSkuOk returns a tuple with the FnSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFnSku

`func (o *ItemSummaryByMarketplace) SetFnSku(v string)`

SetFnSku sets FnSku field to given value.

### HasFnSku

`func (o *ItemSummaryByMarketplace) HasFnSku() bool`

HasFnSku returns a boolean if a field has been set.

### GetItemName

`func (o *ItemSummaryByMarketplace) GetItemName() string`

GetItemName returns the ItemName field if non-nil, zero value otherwise.

### GetItemNameOk

`func (o *ItemSummaryByMarketplace) GetItemNameOk() (*string, bool)`

GetItemNameOk returns a tuple with the ItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemName

`func (o *ItemSummaryByMarketplace) SetItemName(v string)`

SetItemName sets ItemName field to given value.


### GetCreatedDate

`func (o *ItemSummaryByMarketplace) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ItemSummaryByMarketplace) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ItemSummaryByMarketplace) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetLastUpdatedDate

`func (o *ItemSummaryByMarketplace) GetLastUpdatedDate() time.Time`

GetLastUpdatedDate returns the LastUpdatedDate field if non-nil, zero value otherwise.

### GetLastUpdatedDateOk

`func (o *ItemSummaryByMarketplace) GetLastUpdatedDateOk() (*time.Time, bool)`

GetLastUpdatedDateOk returns a tuple with the LastUpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedDate

`func (o *ItemSummaryByMarketplace) SetLastUpdatedDate(v time.Time)`

SetLastUpdatedDate sets LastUpdatedDate field to given value.


### GetMainImage

`func (o *ItemSummaryByMarketplace) GetMainImage() ItemImage`

GetMainImage returns the MainImage field if non-nil, zero value otherwise.

### GetMainImageOk

`func (o *ItemSummaryByMarketplace) GetMainImageOk() (*ItemImage, bool)`

GetMainImageOk returns a tuple with the MainImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainImage

`func (o *ItemSummaryByMarketplace) SetMainImage(v ItemImage)`

SetMainImage sets MainImage field to given value.

### HasMainImage

`func (o *ItemSummaryByMarketplace) HasMainImage() bool`

HasMainImage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


