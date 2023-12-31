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

// checks if the Issue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Issue{}

// Issue An issue with a listings item.
type Issue struct {
	// An issue code that identifies the type of issue.
	Code string `json:"code"`
	// A message that describes the issue.
	Message string `json:"message"`
	// The severity of the issue.
	Severity string `json:"severity"`
	// Names of the attributes associated with the issue, if applicable.
	AttributeNames []string `json:"attributeNames,omitempty"`
}

type _Issue Issue

// NewIssue instantiates a new Issue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssue(code string, message string, severity string) *Issue {
	this := Issue{}
	this.Code = code
	this.Message = message
	this.Severity = severity
	return &this
}

// NewIssueWithDefaults instantiates a new Issue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssueWithDefaults() *Issue {
	this := Issue{}
	return &this
}

// GetCode returns the Code field value
func (o *Issue) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *Issue) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *Issue) SetCode(v string) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *Issue) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *Issue) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *Issue) SetMessage(v string) {
	o.Message = v
}

// GetSeverity returns the Severity field value
func (o *Issue) GetSeverity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *Issue) GetSeverityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value
func (o *Issue) SetSeverity(v string) {
	o.Severity = v
}

// GetAttributeNames returns the AttributeNames field value if set, zero value otherwise.
func (o *Issue) GetAttributeNames() []string {
	if o == nil || IsNil(o.AttributeNames) {
		var ret []string
		return ret
	}
	return o.AttributeNames
}

// GetAttributeNamesOk returns a tuple with the AttributeNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Issue) GetAttributeNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.AttributeNames) {
		return nil, false
	}
	return o.AttributeNames, true
}

// HasAttributeNames returns a boolean if a field has been set.
func (o *Issue) HasAttributeNames() bool {
	if o != nil && !IsNil(o.AttributeNames) {
		return true
	}

	return false
}

// SetAttributeNames gets a reference to the given []string and assigns it to the AttributeNames field.
func (o *Issue) SetAttributeNames(v []string) {
	o.AttributeNames = v
}

func (o Issue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Issue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	toSerialize["severity"] = o.Severity
	if !IsNil(o.AttributeNames) {
		toSerialize["attributeNames"] = o.AttributeNames
	}
	return toSerialize, nil
}

func (o *Issue) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
		"message",
		"severity",
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

	varIssue := _Issue{}

	err = json.Unmarshal(bytes, &varIssue)

	if err != nil {
		return err
	}

	*o = Issue(varIssue)

	return err
}

type NullableIssue struct {
	value *Issue
	isSet bool
}

func (v NullableIssue) Get() *Issue {
	return v.value
}

func (v *NullableIssue) Set(val *Issue) {
	v.value = val
	v.isSet = true
}

func (v NullableIssue) IsSet() bool {
	return v.isSet
}

func (v *NullableIssue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssue(val *Issue) *NullableIssue {
	return &NullableIssue{value: val, isSet: true}
}

func (v NullableIssue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


