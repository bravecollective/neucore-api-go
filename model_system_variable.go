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

// checks if the SystemVariable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SystemVariable{}

// SystemVariable A system settings variable.  This is also used as a storage for Storage\\Variables with the prefix \"__storage__\" if APCu is not available.
type SystemVariable struct {
	// Variable name.
	Name string `json:"name"`
	// Variable value.
	Value NullableString `json:"value"`
}

// NewSystemVariable instantiates a new SystemVariable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemVariable(name string, value NullableString) *SystemVariable {
	this := SystemVariable{}
	this.Name = name
	this.Value = value
	return &this
}

// NewSystemVariableWithDefaults instantiates a new SystemVariable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemVariableWithDefaults() *SystemVariable {
	this := SystemVariable{}
	return &this
}

// GetName returns the Name field value
func (o *SystemVariable) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SystemVariable) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SystemVariable) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SystemVariable) GetValue() string {
	if o == nil || o.Value.Get() == nil {
		var ret string
		return ret
	}

	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SystemVariable) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// SetValue sets field value
func (o *SystemVariable) SetValue(v string) {
	o.Value.Set(&v)
}

func (o SystemVariable) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SystemVariable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["value"] = o.Value.Get()
	return toSerialize, nil
}

type NullableSystemVariable struct {
	value *SystemVariable
	isSet bool
}

func (v NullableSystemVariable) Get() *SystemVariable {
	return v.value
}

func (v *NullableSystemVariable) Set(val *SystemVariable) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemVariable) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemVariable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemVariable(val *SystemVariable) *NullableSystemVariable {
	return &NullableSystemVariable{value: val, isSet: true}
}

func (v NullableSystemVariable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemVariable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


