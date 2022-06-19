/*
Neucore API

Client library of Neucore API

API version: 1.33.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package neucoreapi

import (
	"encoding/json"
	"time"
)

// GroupApplication The player property contains only id and name.
type GroupApplication struct {
	Id int32 `json:"id"`
	Player Player `json:"player"`
	Group Group `json:"group"`
	Created NullableTime `json:"created"`
	// Group application status.
	Status *string `json:"status,omitempty"`
}

// NewGroupApplication instantiates a new GroupApplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupApplication(id int32, player Player, group Group, created NullableTime) *GroupApplication {
	this := GroupApplication{}
	this.Id = id
	this.Player = player
	this.Group = group
	this.Created = created
	return &this
}

// NewGroupApplicationWithDefaults instantiates a new GroupApplication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupApplicationWithDefaults() *GroupApplication {
	this := GroupApplication{}
	return &this
}

// GetId returns the Id field value
func (o *GroupApplication) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GroupApplication) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GroupApplication) SetId(v int32) {
	o.Id = v
}

// GetPlayer returns the Player field value
func (o *GroupApplication) GetPlayer() Player {
	if o == nil {
		var ret Player
		return ret
	}

	return o.Player
}

// GetPlayerOk returns a tuple with the Player field value
// and a boolean to check if the value has been set.
func (o *GroupApplication) GetPlayerOk() (*Player, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Player, true
}

// SetPlayer sets field value
func (o *GroupApplication) SetPlayer(v Player) {
	o.Player = v
}

// GetGroup returns the Group field value
func (o *GroupApplication) GetGroup() Group {
	if o == nil {
		var ret Group
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *GroupApplication) GetGroupOk() (*Group, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *GroupApplication) SetGroup(v Group) {
	o.Group = v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *GroupApplication) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupApplication) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *GroupApplication) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GroupApplication) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupApplication) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GroupApplication) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GroupApplication) SetStatus(v string) {
	o.Status = &v
}

func (o GroupApplication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["player"] = o.Player
	}
	if true {
		toSerialize["group"] = o.Group
	}
	if true {
		toSerialize["created"] = o.Created.Get()
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableGroupApplication struct {
	value *GroupApplication
	isSet bool
}

func (v NullableGroupApplication) Get() *GroupApplication {
	return v.value
}

func (v *NullableGroupApplication) Set(val *GroupApplication) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupApplication(val *GroupApplication) *NullableGroupApplication {
	return &NullableGroupApplication{value: val, isSet: true}
}

func (v NullableGroupApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


