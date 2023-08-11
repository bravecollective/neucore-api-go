/*
Neucore API

Client library of Neucore API

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package neucoreapi

import (
	"encoding/json"
)

// checks if the HourlyAppRequests type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HourlyAppRequests{}

// HourlyAppRequests struct for HourlyAppRequests
type HourlyAppRequests struct {
	AppId int32 `json:"app_id"`
	AppName string `json:"app_name"`
	Requests int32 `json:"requests"`
	Year int32 `json:"year"`
	Month int32 `json:"month"`
	DayOfMonth int32 `json:"day_of_month"`
	Hour int32 `json:"hour"`
}

// NewHourlyAppRequests instantiates a new HourlyAppRequests object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHourlyAppRequests(appId int32, appName string, requests int32, year int32, month int32, dayOfMonth int32, hour int32) *HourlyAppRequests {
	this := HourlyAppRequests{}
	this.AppId = appId
	this.AppName = appName
	this.Requests = requests
	this.Year = year
	this.Month = month
	this.DayOfMonth = dayOfMonth
	this.Hour = hour
	return &this
}

// NewHourlyAppRequestsWithDefaults instantiates a new HourlyAppRequests object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHourlyAppRequestsWithDefaults() *HourlyAppRequests {
	this := HourlyAppRequests{}
	return &this
}

// GetAppId returns the AppId field value
func (o *HourlyAppRequests) GetAppId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value
// and a boolean to check if the value has been set.
func (o *HourlyAppRequests) GetAppIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppId, true
}

// SetAppId sets field value
func (o *HourlyAppRequests) SetAppId(v int32) {
	o.AppId = v
}

// GetAppName returns the AppName field value
func (o *HourlyAppRequests) GetAppName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppName
}

// GetAppNameOk returns a tuple with the AppName field value
// and a boolean to check if the value has been set.
func (o *HourlyAppRequests) GetAppNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppName, true
}

// SetAppName sets field value
func (o *HourlyAppRequests) SetAppName(v string) {
	o.AppName = v
}

// GetRequests returns the Requests field value
func (o *HourlyAppRequests) GetRequests() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value
// and a boolean to check if the value has been set.
func (o *HourlyAppRequests) GetRequestsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Requests, true
}

// SetRequests sets field value
func (o *HourlyAppRequests) SetRequests(v int32) {
	o.Requests = v
}

// GetYear returns the Year field value
func (o *HourlyAppRequests) GetYear() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Year
}

// GetYearOk returns a tuple with the Year field value
// and a boolean to check if the value has been set.
func (o *HourlyAppRequests) GetYearOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Year, true
}

// SetYear sets field value
func (o *HourlyAppRequests) SetYear(v int32) {
	o.Year = v
}

// GetMonth returns the Month field value
func (o *HourlyAppRequests) GetMonth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Month
}

// GetMonthOk returns a tuple with the Month field value
// and a boolean to check if the value has been set.
func (o *HourlyAppRequests) GetMonthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Month, true
}

// SetMonth sets field value
func (o *HourlyAppRequests) SetMonth(v int32) {
	o.Month = v
}

// GetDayOfMonth returns the DayOfMonth field value
func (o *HourlyAppRequests) GetDayOfMonth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DayOfMonth
}

// GetDayOfMonthOk returns a tuple with the DayOfMonth field value
// and a boolean to check if the value has been set.
func (o *HourlyAppRequests) GetDayOfMonthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DayOfMonth, true
}

// SetDayOfMonth sets field value
func (o *HourlyAppRequests) SetDayOfMonth(v int32) {
	o.DayOfMonth = v
}

// GetHour returns the Hour field value
func (o *HourlyAppRequests) GetHour() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Hour
}

// GetHourOk returns a tuple with the Hour field value
// and a boolean to check if the value has been set.
func (o *HourlyAppRequests) GetHourOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hour, true
}

// SetHour sets field value
func (o *HourlyAppRequests) SetHour(v int32) {
	o.Hour = v
}

func (o HourlyAppRequests) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HourlyAppRequests) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["app_id"] = o.AppId
	toSerialize["app_name"] = o.AppName
	toSerialize["requests"] = o.Requests
	toSerialize["year"] = o.Year
	toSerialize["month"] = o.Month
	toSerialize["day_of_month"] = o.DayOfMonth
	toSerialize["hour"] = o.Hour
	return toSerialize, nil
}

type NullableHourlyAppRequests struct {
	value *HourlyAppRequests
	isSet bool
}

func (v NullableHourlyAppRequests) Get() *HourlyAppRequests {
	return v.value
}

func (v *NullableHourlyAppRequests) Set(val *HourlyAppRequests) {
	v.value = val
	v.isSet = true
}

func (v NullableHourlyAppRequests) IsSet() bool {
	return v.isSet
}

func (v *NullableHourlyAppRequests) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHourlyAppRequests(val *HourlyAppRequests) *NullableHourlyAppRequests {
	return &NullableHourlyAppRequests{value: val, isSet: true}
}

func (v NullableHourlyAppRequests) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHourlyAppRequests) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


