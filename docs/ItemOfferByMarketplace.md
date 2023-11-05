# ItemOfferByMarketplace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MarketplaceId** | **string** | Amazon marketplace identifier. | 
**OfferType** | **string** | Type of offer for the listings item. | 
**Price** | [**Money**](Money.md) |  | 
**Points** | Pointer to [**Points**](Points.md) |  | [optional] 

## Methods

### NewItemOfferByMarketplace

`func NewItemOfferByMarketplace(marketplaceId string, offerType string, price Money, ) *ItemOfferByMarketplace`

NewItemOfferByMarketplace instantiates a new ItemOfferByMarketplace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemOfferByMarketplaceWithDefaults

`func NewItemOfferByMarketplaceWithDefaults() *ItemOfferByMarketplace`

NewItemOfferByMarketplaceWithDefaults instantiates a new ItemOfferByMarketplace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarketplaceId

`func (o *ItemOfferByMarketplace) GetMarketplaceId() string`

GetMarketplaceId returns the MarketplaceId field if non-nil, zero value otherwise.

### GetMarketplaceIdOk

`func (o *ItemOfferByMarketplace) GetMarketplaceIdOk() (*string, bool)`

GetMarketplaceIdOk returns a tuple with the MarketplaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceId

`func (o *ItemOfferByMarketplace) SetMarketplaceId(v string)`

SetMarketplaceId sets MarketplaceId field to given value.


### GetOfferType

`func (o *ItemOfferByMarketplace) GetOfferType() string`

GetOfferType returns the OfferType field if non-nil, zero value otherwise.

### GetOfferTypeOk

`func (o *ItemOfferByMarketplace) GetOfferTypeOk() (*string, bool)`

GetOfferTypeOk returns a tuple with the OfferType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferType

`func (o *ItemOfferByMarketplace) SetOfferType(v string)`

SetOfferType sets OfferType field to given value.


### GetPrice

`func (o *ItemOfferByMarketplace) GetPrice() Money`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ItemOfferByMarketplace) GetPriceOk() (*Money, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ItemOfferByMarketplace) SetPrice(v Money)`

SetPrice sets Price field to given value.


### GetPoints

`func (o *ItemOfferByMarketplace) GetPoints() Points`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *ItemOfferByMarketplace) GetPointsOk() (*Points, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *ItemOfferByMarketplace) SetPoints(v Points)`

SetPoints sets Points field to given value.

### HasPoints

`func (o *ItemOfferByMarketplace) HasPoints() bool`

HasPoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


