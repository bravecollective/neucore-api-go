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

// checks if the EsiAccessToken type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EsiAccessToken{}

// EsiAccessToken struct for EsiAccessToken
type EsiAccessToken struct {
	Token string `json:"token"`
	Scopes []string `json:"scopes"`
	Expires int32 `json:"expires"`
}

// NewEsiAccessToken instantiates a new EsiAccessToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEsiAccessToken(token string, scopes []string, expires int32) *EsiAccessToken {
	this := EsiAccessToken{}
	this.Token = token
	this.Scopes = scopes
	this.Expires = expires
	return &this
}

// NewEsiAccessTokenWithDefaults instantiates a new EsiAccessToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEsiAccessTokenWithDefaults() *EsiAccessToken {
	this := EsiAccessToken{}
	return &this
}

// GetToken returns the Token field value
func (o *EsiAccessToken) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *EsiAccessToken) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *EsiAccessToken) SetToken(v string) {
	o.Token = v
}

// GetScopes returns the Scopes field value
func (o *EsiAccessToken) GetScopes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value
// and a boolean to check if the value has been set.
func (o *EsiAccessToken) GetScopesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scopes, true
}

// SetScopes sets field value
func (o *EsiAccessToken) SetScopes(v []string) {
	o.Scopes = v
}

// GetExpires returns the Expires field value
func (o *EsiAccessToken) GetExpires() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Expires
}

// GetExpiresOk returns a tuple with the Expires field value
// and a boolean to check if the value has been set.
func (o *EsiAccessToken) GetExpiresOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expires, true
}

// SetExpires sets field value
func (o *EsiAccessToken) SetExpires(v int32) {
	o.Expires = v
}

func (o EsiAccessToken) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EsiAccessToken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["token"] = o.Token
	toSerialize["scopes"] = o.Scopes
	toSerialize["expires"] = o.Expires
	return toSerialize, nil
}

type NullableEsiAccessToken struct {
	value *EsiAccessToken
	isSet bool
}

func (v NullableEsiAccessToken) Get() *EsiAccessToken {
	return v.value
}

func (v *NullableEsiAccessToken) Set(val *EsiAccessToken) {
	v.value = val
	v.isSet = true
}

func (v NullableEsiAccessToken) IsSet() bool {
	return v.isSet
}

func (v *NullableEsiAccessToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEsiAccessToken(val *EsiAccessToken) *NullableEsiAccessToken {
	return &NullableEsiAccessToken{value: val, isSet: true}
}

func (v NullableEsiAccessToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEsiAccessToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


