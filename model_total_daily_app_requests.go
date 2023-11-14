/*
Neucore API

Client library of Neucore API

API version: 2.4.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package neucoreapi

import (
	"encoding/json"
)

// checks if the TotalDailyAppRequests type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TotalDailyAppRequests{}

// TotalDailyAppRequests struct for TotalDailyAppRequests
type TotalDailyAppRequests struct {
	Requests int32 `json:"requests"`
	Year int32 `json:"year"`
	Month int32 `json:"month"`
	DayOfMonth int32 `json:"day_of_month"`
}

// NewTotalDailyAppRequests instantiates a new TotalDailyAppRequests object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTotalDailyAppRequests(requests int32, year int32, month int32, dayOfMonth int32) *TotalDailyAppRequests {
	this := TotalDailyAppRequests{}
	this.Requests = requests
	this.Year = year
	this.Month = month
	this.DayOfMonth = dayOfMonth
	return &this
}

// NewTotalDailyAppRequestsWithDefaults instantiates a new TotalDailyAppRequests object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTotalDailyAppRequestsWithDefaults() *TotalDailyAppRequests {
	this := TotalDailyAppRequests{}
	return &this
}

// GetRequests returns the Requests field value
func (o *TotalDailyAppRequests) GetRequests() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value
// and a boolean to check if the value has been set.
func (o *TotalDailyAppRequests) GetRequestsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Requests, true
}

// SetRequests sets field value
func (o *TotalDailyAppRequests) SetRequests(v int32) {
	o.Requests = v
}

// GetYear returns the Year field value
func (o *TotalDailyAppRequests) GetYear() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Year
}

// GetYearOk returns a tuple with the Year field value
// and a boolean to check if the value has been set.
func (o *TotalDailyAppRequests) GetYearOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Year, true
}

// SetYear sets field value
func (o *TotalDailyAppRequests) SetYear(v int32) {
	o.Year = v
}

// GetMonth returns the Month field value
func (o *TotalDailyAppRequests) GetMonth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Month
}

// GetMonthOk returns a tuple with the Month field value
// and a boolean to check if the value has been set.
func (o *TotalDailyAppRequests) GetMonthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Month, true
}

// SetMonth sets field value
func (o *TotalDailyAppRequests) SetMonth(v int32) {
	o.Month = v
}

// GetDayOfMonth returns the DayOfMonth field value
func (o *TotalDailyAppRequests) GetDayOfMonth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DayOfMonth
}

// GetDayOfMonthOk returns a tuple with the DayOfMonth field value
// and a boolean to check if the value has been set.
func (o *TotalDailyAppRequests) GetDayOfMonthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DayOfMonth, true
}

// SetDayOfMonth sets field value
func (o *TotalDailyAppRequests) SetDayOfMonth(v int32) {
	o.DayOfMonth = v
}

func (o TotalDailyAppRequests) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TotalDailyAppRequests) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["requests"] = o.Requests
	toSerialize["year"] = o.Year
	toSerialize["month"] = o.Month
	toSerialize["day_of_month"] = o.DayOfMonth
	return toSerialize, nil
}

type NullableTotalDailyAppRequests struct {
	value *TotalDailyAppRequests
	isSet bool
}

func (v NullableTotalDailyAppRequests) Get() *TotalDailyAppRequests {
	return v.value
}

func (v *NullableTotalDailyAppRequests) Set(val *TotalDailyAppRequests) {
	v.value = val
	v.isSet = true
}

func (v NullableTotalDailyAppRequests) IsSet() bool {
	return v.isSet
}

func (v *NullableTotalDailyAppRequests) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTotalDailyAppRequests(val *TotalDailyAppRequests) *NullableTotalDailyAppRequests {
	return &NullableTotalDailyAppRequests{value: val, isSet: true}
}

func (v NullableTotalDailyAppRequests) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTotalDailyAppRequests) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


