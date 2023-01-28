/*
Neucore API

Client library of Neucore API

API version: 1.43.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package neucoreapi

import (
	"encoding/json"
)

// EsiLocation An EVE location (System, Station, Structure, ...)
type EsiLocation struct {
	Id NullableInt64 `json:"id"`
	Category NullableString `json:"category"`
	Name NullableString `json:"name"`
}

// NewEsiLocation instantiates a new EsiLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEsiLocation(id NullableInt64, category NullableString, name NullableString) *EsiLocation {
	this := EsiLocation{}
	this.Id = id
	this.Category = category
	this.Name = name
	return &this
}

// NewEsiLocationWithDefaults instantiates a new EsiLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEsiLocationWithDefaults() *EsiLocation {
	this := EsiLocation{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *EsiLocation) GetId() int64 {
	if o == nil || o.Id.Get() == nil {
		var ret int64
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EsiLocation) GetIdOk() (*int64, bool) {
	if o == nil {
    return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *EsiLocation) SetId(v int64) {
	o.Id.Set(&v)
}

// GetCategory returns the Category field value
// If the value is explicit nil, the zero value for string will be returned
func (o *EsiLocation) GetCategory() string {
	if o == nil || o.Category.Get() == nil {
		var ret string
		return ret
	}

	return *o.Category.Get()
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EsiLocation) GetCategoryOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Category.Get(), o.Category.IsSet()
}

// SetCategory sets field value
func (o *EsiLocation) SetCategory(v string) {
	o.Category.Set(&v)
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *EsiLocation) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EsiLocation) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *EsiLocation) SetName(v string) {
	o.Name.Set(&v)
}

func (o EsiLocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id.Get()
	}
	if true {
		toSerialize["category"] = o.Category.Get()
	}
	if true {
		toSerialize["name"] = o.Name.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableEsiLocation struct {
	value *EsiLocation
	isSet bool
}

func (v NullableEsiLocation) Get() *EsiLocation {
	return v.value
}

func (v *NullableEsiLocation) Set(val *EsiLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableEsiLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableEsiLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEsiLocation(val *EsiLocation) *NullableEsiLocation {
	return &NullableEsiLocation{value: val, isSet: true}
}

func (v NullableEsiLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEsiLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


