/*
Neucore API

Client library of Neucore API

API version: 1.26.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package neucoreapi

import (
	"encoding/json"
)

// TotalMonthlyAppRequests struct for TotalMonthlyAppRequests
type TotalMonthlyAppRequests struct {
	Requests int32 `json:"requests"`
	Year int32 `json:"year"`
	Month int32 `json:"month"`
}

// NewTotalMonthlyAppRequests instantiates a new TotalMonthlyAppRequests object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTotalMonthlyAppRequests(requests int32, year int32, month int32) *TotalMonthlyAppRequests {
	this := TotalMonthlyAppRequests{}
	this.Requests = requests
	this.Year = year
	this.Month = month
	return &this
}

// NewTotalMonthlyAppRequestsWithDefaults instantiates a new TotalMonthlyAppRequests object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTotalMonthlyAppRequestsWithDefaults() *TotalMonthlyAppRequests {
	this := TotalMonthlyAppRequests{}
	return &this
}

// GetRequests returns the Requests field value
func (o *TotalMonthlyAppRequests) GetRequests() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value
// and a boolean to check if the value has been set.
func (o *TotalMonthlyAppRequests) GetRequestsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Requests, true
}

// SetRequests sets field value
func (o *TotalMonthlyAppRequests) SetRequests(v int32) {
	o.Requests = v
}

// GetYear returns the Year field value
func (o *TotalMonthlyAppRequests) GetYear() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Year
}

// GetYearOk returns a tuple with the Year field value
// and a boolean to check if the value has been set.
func (o *TotalMonthlyAppRequests) GetYearOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Year, true
}

// SetYear sets field value
func (o *TotalMonthlyAppRequests) SetYear(v int32) {
	o.Year = v
}

// GetMonth returns the Month field value
func (o *TotalMonthlyAppRequests) GetMonth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Month
}

// GetMonthOk returns a tuple with the Month field value
// and a boolean to check if the value has been set.
func (o *TotalMonthlyAppRequests) GetMonthOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Month, true
}

// SetMonth sets field value
func (o *TotalMonthlyAppRequests) SetMonth(v int32) {
	o.Month = v
}

func (o TotalMonthlyAppRequests) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["requests"] = o.Requests
	}
	if true {
		toSerialize["year"] = o.Year
	}
	if true {
		toSerialize["month"] = o.Month
	}
	return json.Marshal(toSerialize)
}

type NullableTotalMonthlyAppRequests struct {
	value *TotalMonthlyAppRequests
	isSet bool
}

func (v NullableTotalMonthlyAppRequests) Get() *TotalMonthlyAppRequests {
	return v.value
}

func (v *NullableTotalMonthlyAppRequests) Set(val *TotalMonthlyAppRequests) {
	v.value = val
	v.isSet = true
}

func (v NullableTotalMonthlyAppRequests) IsSet() bool {
	return v.isSet
}

func (v *NullableTotalMonthlyAppRequests) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTotalMonthlyAppRequests(val *TotalMonthlyAppRequests) *NullableTotalMonthlyAppRequests {
	return &NullableTotalMonthlyAppRequests{value: val, isSet: true}
}

func (v NullableTotalMonthlyAppRequests) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTotalMonthlyAppRequests) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


