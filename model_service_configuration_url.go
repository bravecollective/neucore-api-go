/*
Neucore API

Client library of Neucore API

API version: 1.33.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package neucoreapi

import (
	"encoding/json"
)

// ServiceConfigurationURL struct for ServiceConfigurationURL
type ServiceConfigurationURL struct {
	// placeholders: {username}, {password}, {email}
	Url string `json:"url"`
	Title string `json:"title"`
	Target string `json:"target"`
}

// NewServiceConfigurationURL instantiates a new ServiceConfigurationURL object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceConfigurationURL(url string, title string, target string) *ServiceConfigurationURL {
	this := ServiceConfigurationURL{}
	this.Url = url
	this.Title = title
	this.Target = target
	return &this
}

// NewServiceConfigurationURLWithDefaults instantiates a new ServiceConfigurationURL object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceConfigurationURLWithDefaults() *ServiceConfigurationURL {
	this := ServiceConfigurationURL{}
	return &this
}

// GetUrl returns the Url field value
func (o *ServiceConfigurationURL) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ServiceConfigurationURL) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ServiceConfigurationURL) SetUrl(v string) {
	o.Url = v
}

// GetTitle returns the Title field value
func (o *ServiceConfigurationURL) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ServiceConfigurationURL) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *ServiceConfigurationURL) SetTitle(v string) {
	o.Title = v
}

// GetTarget returns the Target field value
func (o *ServiceConfigurationURL) GetTarget() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *ServiceConfigurationURL) GetTargetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *ServiceConfigurationURL) SetTarget(v string) {
	o.Target = v
}

func (o ServiceConfigurationURL) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if true {
		toSerialize["target"] = o.Target
	}
	return json.Marshal(toSerialize)
}

type NullableServiceConfigurationURL struct {
	value *ServiceConfigurationURL
	isSet bool
}

func (v NullableServiceConfigurationURL) Get() *ServiceConfigurationURL {
	return v.value
}

func (v *NullableServiceConfigurationURL) Set(val *ServiceConfigurationURL) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceConfigurationURL) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceConfigurationURL) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceConfigurationURL(val *ServiceConfigurationURL) *NullableServiceConfigurationURL {
	return &NullableServiceConfigurationURL{value: val, isSet: true}
}

func (v NullableServiceConfigurationURL) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceConfigurationURL) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


