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

// Watchlist struct for Watchlist
type Watchlist struct {
	Id NullableInt32 `json:"id"`
	Name NullableString `json:"name"`
	LockWatchlistSettings *bool `json:"lockWatchlistSettings,omitempty"`
}

// NewWatchlist instantiates a new Watchlist object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWatchlist(id NullableInt32, name NullableString) *Watchlist {
	this := Watchlist{}
	this.Id = id
	this.Name = name
	return &this
}

// NewWatchlistWithDefaults instantiates a new Watchlist object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWatchlistWithDefaults() *Watchlist {
	this := Watchlist{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *Watchlist) GetId() int32 {
	if o == nil || o.Id.Get() == nil {
		var ret int32
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Watchlist) GetIdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *Watchlist) SetId(v int32) {
	o.Id.Set(&v)
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Watchlist) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Watchlist) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *Watchlist) SetName(v string) {
	o.Name.Set(&v)
}

// GetLockWatchlistSettings returns the LockWatchlistSettings field value if set, zero value otherwise.
func (o *Watchlist) GetLockWatchlistSettings() bool {
	if o == nil || isNil(o.LockWatchlistSettings) {
		var ret bool
		return ret
	}
	return *o.LockWatchlistSettings
}

// GetLockWatchlistSettingsOk returns a tuple with the LockWatchlistSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Watchlist) GetLockWatchlistSettingsOk() (*bool, bool) {
	if o == nil || isNil(o.LockWatchlistSettings) {
    return nil, false
	}
	return o.LockWatchlistSettings, true
}

// HasLockWatchlistSettings returns a boolean if a field has been set.
func (o *Watchlist) HasLockWatchlistSettings() bool {
	if o != nil && !isNil(o.LockWatchlistSettings) {
		return true
	}

	return false
}

// SetLockWatchlistSettings gets a reference to the given bool and assigns it to the LockWatchlistSettings field.
func (o *Watchlist) SetLockWatchlistSettings(v bool) {
	o.LockWatchlistSettings = &v
}

func (o Watchlist) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id.Get()
	}
	if true {
		toSerialize["name"] = o.Name.Get()
	}
	if !isNil(o.LockWatchlistSettings) {
		toSerialize["lockWatchlistSettings"] = o.LockWatchlistSettings
	}
	return json.Marshal(toSerialize)
}

type NullableWatchlist struct {
	value *Watchlist
	isSet bool
}

func (v NullableWatchlist) Get() *Watchlist {
	return v.value
}

func (v *NullableWatchlist) Set(val *Watchlist) {
	v.value = val
	v.isSet = true
}

func (v NullableWatchlist) IsSet() bool {
	return v.isSet
}

func (v *NullableWatchlist) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatchlist(val *Watchlist) *NullableWatchlist {
	return &NullableWatchlist{value: val, isSet: true}
}

func (v NullableWatchlist) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatchlist) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


