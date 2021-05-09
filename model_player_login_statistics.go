/*
 * Neucore API
 *
 * Client library of Neucore API
 *
 * API version: 1.21.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package neucoreapi

import (
	"encoding/json"
)

// PlayerLoginStatistics struct for PlayerLoginStatistics
type PlayerLoginStatistics struct {
	UniqueLogins int32 `json:"unique_logins"`
	TotalLogins int32 `json:"total_logins"`
	Year int32 `json:"year"`
	Month int32 `json:"month"`
}

// NewPlayerLoginStatistics instantiates a new PlayerLoginStatistics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlayerLoginStatistics(uniqueLogins int32, totalLogins int32, year int32, month int32) *PlayerLoginStatistics {
	this := PlayerLoginStatistics{}
	this.UniqueLogins = uniqueLogins
	this.TotalLogins = totalLogins
	this.Year = year
	this.Month = month
	return &this
}

// NewPlayerLoginStatisticsWithDefaults instantiates a new PlayerLoginStatistics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlayerLoginStatisticsWithDefaults() *PlayerLoginStatistics {
	this := PlayerLoginStatistics{}
	return &this
}

// GetUniqueLogins returns the UniqueLogins field value
func (o *PlayerLoginStatistics) GetUniqueLogins() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UniqueLogins
}

// GetUniqueLoginsOk returns a tuple with the UniqueLogins field value
// and a boolean to check if the value has been set.
func (o *PlayerLoginStatistics) GetUniqueLoginsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UniqueLogins, true
}

// SetUniqueLogins sets field value
func (o *PlayerLoginStatistics) SetUniqueLogins(v int32) {
	o.UniqueLogins = v
}

// GetTotalLogins returns the TotalLogins field value
func (o *PlayerLoginStatistics) GetTotalLogins() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalLogins
}

// GetTotalLoginsOk returns a tuple with the TotalLogins field value
// and a boolean to check if the value has been set.
func (o *PlayerLoginStatistics) GetTotalLoginsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalLogins, true
}

// SetTotalLogins sets field value
func (o *PlayerLoginStatistics) SetTotalLogins(v int32) {
	o.TotalLogins = v
}

// GetYear returns the Year field value
func (o *PlayerLoginStatistics) GetYear() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Year
}

// GetYearOk returns a tuple with the Year field value
// and a boolean to check if the value has been set.
func (o *PlayerLoginStatistics) GetYearOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Year, true
}

// SetYear sets field value
func (o *PlayerLoginStatistics) SetYear(v int32) {
	o.Year = v
}

// GetMonth returns the Month field value
func (o *PlayerLoginStatistics) GetMonth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Month
}

// GetMonthOk returns a tuple with the Month field value
// and a boolean to check if the value has been set.
func (o *PlayerLoginStatistics) GetMonthOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Month, true
}

// SetMonth sets field value
func (o *PlayerLoginStatistics) SetMonth(v int32) {
	o.Month = v
}

func (o PlayerLoginStatistics) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["unique_logins"] = o.UniqueLogins
	}
	if true {
		toSerialize["total_logins"] = o.TotalLogins
	}
	if true {
		toSerialize["year"] = o.Year
	}
	if true {
		toSerialize["month"] = o.Month
	}
	return json.Marshal(toSerialize)
}

type NullablePlayerLoginStatistics struct {
	value *PlayerLoginStatistics
	isSet bool
}

func (v NullablePlayerLoginStatistics) Get() *PlayerLoginStatistics {
	return v.value
}

func (v *NullablePlayerLoginStatistics) Set(val *PlayerLoginStatistics) {
	v.value = val
	v.isSet = true
}

func (v NullablePlayerLoginStatistics) IsSet() bool {
	return v.isSet
}

func (v *NullablePlayerLoginStatistics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlayerLoginStatistics(val *PlayerLoginStatistics) *NullablePlayerLoginStatistics {
	return &NullablePlayerLoginStatistics{value: val, isSet: true}
}

func (v NullablePlayerLoginStatistics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlayerLoginStatistics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


