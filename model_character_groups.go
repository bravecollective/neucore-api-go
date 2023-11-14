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

// checks if the CharacterGroups type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CharacterGroups{}

// CharacterGroups struct for CharacterGroups
type CharacterGroups struct {
	Character Character `json:"character"`
	Groups []Group `json:"groups"`
	// Groups deactivation status.
	Deactivated string `json:"deactivated"`
}

// NewCharacterGroups instantiates a new CharacterGroups object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCharacterGroups(character Character, groups []Group, deactivated string) *CharacterGroups {
	this := CharacterGroups{}
	this.Character = character
	this.Groups = groups
	this.Deactivated = deactivated
	return &this
}

// NewCharacterGroupsWithDefaults instantiates a new CharacterGroups object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCharacterGroupsWithDefaults() *CharacterGroups {
	this := CharacterGroups{}
	return &this
}

// GetCharacter returns the Character field value
func (o *CharacterGroups) GetCharacter() Character {
	if o == nil {
		var ret Character
		return ret
	}

	return o.Character
}

// GetCharacterOk returns a tuple with the Character field value
// and a boolean to check if the value has been set.
func (o *CharacterGroups) GetCharacterOk() (*Character, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Character, true
}

// SetCharacter sets field value
func (o *CharacterGroups) SetCharacter(v Character) {
	o.Character = v
}

// GetGroups returns the Groups field value
func (o *CharacterGroups) GetGroups() []Group {
	if o == nil {
		var ret []Group
		return ret
	}

	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value
// and a boolean to check if the value has been set.
func (o *CharacterGroups) GetGroupsOk() ([]Group, bool) {
	if o == nil {
		return nil, false
	}
	return o.Groups, true
}

// SetGroups sets field value
func (o *CharacterGroups) SetGroups(v []Group) {
	o.Groups = v
}

// GetDeactivated returns the Deactivated field value
func (o *CharacterGroups) GetDeactivated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Deactivated
}

// GetDeactivatedOk returns a tuple with the Deactivated field value
// and a boolean to check if the value has been set.
func (o *CharacterGroups) GetDeactivatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Deactivated, true
}

// SetDeactivated sets field value
func (o *CharacterGroups) SetDeactivated(v string) {
	o.Deactivated = v
}

func (o CharacterGroups) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CharacterGroups) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["character"] = o.Character
	toSerialize["groups"] = o.Groups
	toSerialize["deactivated"] = o.Deactivated
	return toSerialize, nil
}

type NullableCharacterGroups struct {
	value *CharacterGroups
	isSet bool
}

func (v NullableCharacterGroups) Get() *CharacterGroups {
	return v.value
}

func (v *NullableCharacterGroups) Set(val *CharacterGroups) {
	v.value = val
	v.isSet = true
}

func (v NullableCharacterGroups) IsSet() bool {
	return v.isSet
}

func (v *NullableCharacterGroups) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCharacterGroups(val *CharacterGroups) *NullableCharacterGroups {
	return &NullableCharacterGroups{value: val, isSet: true}
}

func (v NullableCharacterGroups) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCharacterGroups) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


