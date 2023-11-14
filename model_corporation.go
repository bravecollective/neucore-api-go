/*
Neucore API

Client library of Neucore API

API version: 2.4.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package neucoreapi

import (
	"encoding/json"
	"time"
)

// checks if the Corporation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Corporation{}

// Corporation EVE corporation.
type Corporation struct {
	// EVE corporation ID.
	Id NullableInt64 `json:"id"`
	// EVE corporation name.
	Name NullableString `json:"name"`
	// Corporation ticker.
	Ticker NullableString `json:"ticker"`
	Alliance *Alliance `json:"alliance,omitempty"`
	// Groups for automatic assignment (API: not included by default).
	Groups []Group `json:"groups,omitempty"`
	// Last update of corporation member tracking data (API: not included by default).
	TrackingLastUpdate NullableTime `json:"trackingLastUpdate,omitempty"`
	// True if this corporation was automatically placed on the allowlist of a watchlist (API: not included by default).
	AutoAllowlist *bool `json:"autoAllowlist,omitempty"`
}

// NewCorporation instantiates a new Corporation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCorporation(id NullableInt64, name NullableString, ticker NullableString) *Corporation {
	this := Corporation{}
	this.Id = id
	this.Name = name
	this.Ticker = ticker
	return &this
}

// NewCorporationWithDefaults instantiates a new Corporation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCorporationWithDefaults() *Corporation {
	this := Corporation{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *Corporation) GetId() int64 {
	if o == nil || o.Id.Get() == nil {
		var ret int64
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Corporation) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *Corporation) SetId(v int64) {
	o.Id.Set(&v)
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Corporation) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Corporation) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *Corporation) SetName(v string) {
	o.Name.Set(&v)
}

// GetTicker returns the Ticker field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Corporation) GetTicker() string {
	if o == nil || o.Ticker.Get() == nil {
		var ret string
		return ret
	}

	return *o.Ticker.Get()
}

// GetTickerOk returns a tuple with the Ticker field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Corporation) GetTickerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ticker.Get(), o.Ticker.IsSet()
}

// SetTicker sets field value
func (o *Corporation) SetTicker(v string) {
	o.Ticker.Set(&v)
}

// GetAlliance returns the Alliance field value if set, zero value otherwise.
func (o *Corporation) GetAlliance() Alliance {
	if o == nil || IsNil(o.Alliance) {
		var ret Alliance
		return ret
	}
	return *o.Alliance
}

// GetAllianceOk returns a tuple with the Alliance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Corporation) GetAllianceOk() (*Alliance, bool) {
	if o == nil || IsNil(o.Alliance) {
		return nil, false
	}
	return o.Alliance, true
}

// HasAlliance returns a boolean if a field has been set.
func (o *Corporation) HasAlliance() bool {
	if o != nil && !IsNil(o.Alliance) {
		return true
	}

	return false
}

// SetAlliance gets a reference to the given Alliance and assigns it to the Alliance field.
func (o *Corporation) SetAlliance(v Alliance) {
	o.Alliance = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *Corporation) GetGroups() []Group {
	if o == nil || IsNil(o.Groups) {
		var ret []Group
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Corporation) GetGroupsOk() ([]Group, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *Corporation) HasGroups() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []Group and assigns it to the Groups field.
func (o *Corporation) SetGroups(v []Group) {
	o.Groups = v
}

// GetTrackingLastUpdate returns the TrackingLastUpdate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Corporation) GetTrackingLastUpdate() time.Time {
	if o == nil || IsNil(o.TrackingLastUpdate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.TrackingLastUpdate.Get()
}

// GetTrackingLastUpdateOk returns a tuple with the TrackingLastUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Corporation) GetTrackingLastUpdateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.TrackingLastUpdate.Get(), o.TrackingLastUpdate.IsSet()
}

// HasTrackingLastUpdate returns a boolean if a field has been set.
func (o *Corporation) HasTrackingLastUpdate() bool {
	if o != nil && o.TrackingLastUpdate.IsSet() {
		return true
	}

	return false
}

// SetTrackingLastUpdate gets a reference to the given NullableTime and assigns it to the TrackingLastUpdate field.
func (o *Corporation) SetTrackingLastUpdate(v time.Time) {
	o.TrackingLastUpdate.Set(&v)
}
// SetTrackingLastUpdateNil sets the value for TrackingLastUpdate to be an explicit nil
func (o *Corporation) SetTrackingLastUpdateNil() {
	o.TrackingLastUpdate.Set(nil)
}

// UnsetTrackingLastUpdate ensures that no value is present for TrackingLastUpdate, not even an explicit nil
func (o *Corporation) UnsetTrackingLastUpdate() {
	o.TrackingLastUpdate.Unset()
}

// GetAutoAllowlist returns the AutoAllowlist field value if set, zero value otherwise.
func (o *Corporation) GetAutoAllowlist() bool {
	if o == nil || IsNil(o.AutoAllowlist) {
		var ret bool
		return ret
	}
	return *o.AutoAllowlist
}

// GetAutoAllowlistOk returns a tuple with the AutoAllowlist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Corporation) GetAutoAllowlistOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoAllowlist) {
		return nil, false
	}
	return o.AutoAllowlist, true
}

// HasAutoAllowlist returns a boolean if a field has been set.
func (o *Corporation) HasAutoAllowlist() bool {
	if o != nil && !IsNil(o.AutoAllowlist) {
		return true
	}

	return false
}

// SetAutoAllowlist gets a reference to the given bool and assigns it to the AutoAllowlist field.
func (o *Corporation) SetAutoAllowlist(v bool) {
	o.AutoAllowlist = &v
}

func (o Corporation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Corporation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id.Get()
	toSerialize["name"] = o.Name.Get()
	toSerialize["ticker"] = o.Ticker.Get()
	if !IsNil(o.Alliance) {
		toSerialize["alliance"] = o.Alliance
	}
	if !IsNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	if o.TrackingLastUpdate.IsSet() {
		toSerialize["trackingLastUpdate"] = o.TrackingLastUpdate.Get()
	}
	if !IsNil(o.AutoAllowlist) {
		toSerialize["autoAllowlist"] = o.AutoAllowlist
	}
	return toSerialize, nil
}

type NullableCorporation struct {
	value *Corporation
	isSet bool
}

func (v NullableCorporation) Get() *Corporation {
	return v.value
}

func (v *NullableCorporation) Set(val *Corporation) {
	v.value = val
	v.isSet = true
}

func (v NullableCorporation) IsSet() bool {
	return v.isSet
}

func (v *NullableCorporation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCorporation(val *Corporation) *NullableCorporation {
	return &NullableCorporation{value: val, isSet: true}
}

func (v NullableCorporation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCorporation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


