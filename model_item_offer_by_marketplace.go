/*
Selling Partner API for Listings Items

The Selling Partner API for Listings Items (Listings Items API) provides programmatic access to selling partner listings on Amazon. Use this API in collaboration with the Selling Partner API for Product Type Definitions, which you use to retrieve the information about Amazon product types needed to use the Listings Items API.  For more information, see the [Listings Items API Use Case Guide](doc:listings-items-api-v2021-08-01-use-case-guide).

API version: 2021-08-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// checks if the ItemOfferByMarketplace type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemOfferByMarketplace{}

// ItemOfferByMarketplace Offer details of a listings item for an Amazon marketplace.
type ItemOfferByMarketplace struct {
	// Amazon marketplace identifier.
	MarketplaceId string `json:"marketplaceId"`
	// Type of offer for the listings item.
	OfferType string `json:"offerType"`
	Price Money `json:"price"`
	Points *Points `json:"points,omitempty"`
}

type _ItemOfferByMarketplace ItemOfferByMarketplace

// NewItemOfferByMarketplace instantiates a new ItemOfferByMarketplace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemOfferByMarketplace(marketplaceId string, offerType string, price Money) *ItemOfferByMarketplace {
	this := ItemOfferByMarketplace{}
	this.MarketplaceId = marketplaceId
	this.OfferType = offerType
	this.Price = price
	return &this
}

// NewItemOfferByMarketplaceWithDefaults instantiates a new ItemOfferByMarketplace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemOfferByMarketplaceWithDefaults() *ItemOfferByMarketplace {
	this := ItemOfferByMarketplace{}
	return &this
}

// GetMarketplaceId returns the MarketplaceId field value
func (o *ItemOfferByMarketplace) GetMarketplaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MarketplaceId
}

// GetMarketplaceIdOk returns a tuple with the MarketplaceId field value
// and a boolean to check if the value has been set.
func (o *ItemOfferByMarketplace) GetMarketplaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MarketplaceId, true
}

// SetMarketplaceId sets field value
func (o *ItemOfferByMarketplace) SetMarketplaceId(v string) {
	o.MarketplaceId = v
}

// GetOfferType returns the OfferType field value
func (o *ItemOfferByMarketplace) GetOfferType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OfferType
}

// GetOfferTypeOk returns a tuple with the OfferType field value
// and a boolean to check if the value has been set.
func (o *ItemOfferByMarketplace) GetOfferTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OfferType, true
}

// SetOfferType sets field value
func (o *ItemOfferByMarketplace) SetOfferType(v string) {
	o.OfferType = v
}

// GetPrice returns the Price field value
func (o *ItemOfferByMarketplace) GetPrice() Money {
	if o == nil {
		var ret Money
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *ItemOfferByMarketplace) GetPriceOk() (*Money, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *ItemOfferByMarketplace) SetPrice(v Money) {
	o.Price = v
}

// GetPoints returns the Points field value if set, zero value otherwise.
func (o *ItemOfferByMarketplace) GetPoints() Points {
	if o == nil || IsNil(o.Points) {
		var ret Points
		return ret
	}
	return *o.Points
}

// GetPointsOk returns a tuple with the Points field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemOfferByMarketplace) GetPointsOk() (*Points, bool) {
	if o == nil || IsNil(o.Points) {
		return nil, false
	}
	return o.Points, true
}

// HasPoints returns a boolean if a field has been set.
func (o *ItemOfferByMarketplace) HasPoints() bool {
	if o != nil && !IsNil(o.Points) {
		return true
	}

	return false
}

// SetPoints gets a reference to the given Points and assigns it to the Points field.
func (o *ItemOfferByMarketplace) SetPoints(v Points) {
	o.Points = &v
}

func (o ItemOfferByMarketplace) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ItemOfferByMarketplace) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["marketplaceId"] = o.MarketplaceId
	toSerialize["offerType"] = o.OfferType
	toSerialize["price"] = o.Price
	if !IsNil(o.Points) {
		toSerialize["points"] = o.Points
	}
	return toSerialize, nil
}

func (o *ItemOfferByMarketplace) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"marketplaceId",
		"offerType",
		"price",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varItemOfferByMarketplace := _ItemOfferByMarketplace{}

	err = json.Unmarshal(bytes, &varItemOfferByMarketplace)

	if err != nil {
		return err
	}

	*o = ItemOfferByMarketplace(varItemOfferByMarketplace)

	return err
}

type NullableItemOfferByMarketplace struct {
	value *ItemOfferByMarketplace
	isSet bool
}

func (v NullableItemOfferByMarketplace) Get() *ItemOfferByMarketplace {
	return v.value
}

func (v *NullableItemOfferByMarketplace) Set(val *ItemOfferByMarketplace) {
	v.value = val
	v.isSet = true
}

func (v NullableItemOfferByMarketplace) IsSet() bool {
	return v.isSet
}

func (v *NullableItemOfferByMarketplace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemOfferByMarketplace(val *ItemOfferByMarketplace) *NullableItemOfferByMarketplace {
	return &NullableItemOfferByMarketplace{value: val, isSet: true}
}

func (v NullableItemOfferByMarketplace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemOfferByMarketplace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

