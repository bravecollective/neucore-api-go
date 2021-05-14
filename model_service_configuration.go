/*
 * Neucore API
 *
 * Client library of Neucore API
 *
 * API version: 1.22.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package neucoreapi

import (
	"encoding/json"
)

// ServiceConfiguration struct for ServiceConfiguration
type ServiceConfiguration struct {
	PhpClass *string `json:"phpClass,omitempty"`
	Psr4Prefix *string `json:"psr4Prefix,omitempty"`
	Psr4Path *string `json:"psr4Path,omitempty"`
	RequiredGroups *[]int32 `json:"requiredGroups,omitempty"`
	Properties []string `json:"properties"`
	ShowPassword *bool `json:"showPassword,omitempty"`
	Actions []string `json:"actions"`
	URLs []ServiceConfigurationURL `json:"URLs"`
	TextAccount string `json:"textAccount"`
	TextTop string `json:"textTop"`
	TextRegister string `json:"textRegister"`
	TextPending string `json:"textPending"`
}

// NewServiceConfiguration instantiates a new ServiceConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceConfiguration(properties []string, actions []string, uRLs []ServiceConfigurationURL, textAccount string, textTop string, textRegister string, textPending string) *ServiceConfiguration {
	this := ServiceConfiguration{}
	this.Properties = properties
	this.Actions = actions
	this.URLs = uRLs
	this.TextAccount = textAccount
	this.TextTop = textTop
	this.TextRegister = textRegister
	this.TextPending = textPending
	return &this
}

// NewServiceConfigurationWithDefaults instantiates a new ServiceConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceConfigurationWithDefaults() *ServiceConfiguration {
	this := ServiceConfiguration{}
	return &this
}

// GetPhpClass returns the PhpClass field value if set, zero value otherwise.
func (o *ServiceConfiguration) GetPhpClass() string {
	if o == nil || o.PhpClass == nil {
		var ret string
		return ret
	}
	return *o.PhpClass
}

// GetPhpClassOk returns a tuple with the PhpClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceConfiguration) GetPhpClassOk() (*string, bool) {
	if o == nil || o.PhpClass == nil {
		return nil, false
	}
	return o.PhpClass, true
}

// HasPhpClass returns a boolean if a field has been set.
func (o *ServiceConfiguration) HasPhpClass() bool {
	if o != nil && o.PhpClass != nil {
		return true
	}

	return false
}

// SetPhpClass gets a reference to the given string and assigns it to the PhpClass field.
func (o *ServiceConfiguration) SetPhpClass(v string) {
	o.PhpClass = &v
}

// GetPsr4Prefix returns the Psr4Prefix field value if set, zero value otherwise.
func (o *ServiceConfiguration) GetPsr4Prefix() string {
	if o == nil || o.Psr4Prefix == nil {
		var ret string
		return ret
	}
	return *o.Psr4Prefix
}

// GetPsr4PrefixOk returns a tuple with the Psr4Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceConfiguration) GetPsr4PrefixOk() (*string, bool) {
	if o == nil || o.Psr4Prefix == nil {
		return nil, false
	}
	return o.Psr4Prefix, true
}

// HasPsr4Prefix returns a boolean if a field has been set.
func (o *ServiceConfiguration) HasPsr4Prefix() bool {
	if o != nil && o.Psr4Prefix != nil {
		return true
	}

	return false
}

// SetPsr4Prefix gets a reference to the given string and assigns it to the Psr4Prefix field.
func (o *ServiceConfiguration) SetPsr4Prefix(v string) {
	o.Psr4Prefix = &v
}

// GetPsr4Path returns the Psr4Path field value if set, zero value otherwise.
func (o *ServiceConfiguration) GetPsr4Path() string {
	if o == nil || o.Psr4Path == nil {
		var ret string
		return ret
	}
	return *o.Psr4Path
}

// GetPsr4PathOk returns a tuple with the Psr4Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceConfiguration) GetPsr4PathOk() (*string, bool) {
	if o == nil || o.Psr4Path == nil {
		return nil, false
	}
	return o.Psr4Path, true
}

// HasPsr4Path returns a boolean if a field has been set.
func (o *ServiceConfiguration) HasPsr4Path() bool {
	if o != nil && o.Psr4Path != nil {
		return true
	}

	return false
}

// SetPsr4Path gets a reference to the given string and assigns it to the Psr4Path field.
func (o *ServiceConfiguration) SetPsr4Path(v string) {
	o.Psr4Path = &v
}

// GetRequiredGroups returns the RequiredGroups field value if set, zero value otherwise.
func (o *ServiceConfiguration) GetRequiredGroups() []int32 {
	if o == nil || o.RequiredGroups == nil {
		var ret []int32
		return ret
	}
	return *o.RequiredGroups
}

// GetRequiredGroupsOk returns a tuple with the RequiredGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceConfiguration) GetRequiredGroupsOk() (*[]int32, bool) {
	if o == nil || o.RequiredGroups == nil {
		return nil, false
	}
	return o.RequiredGroups, true
}

// HasRequiredGroups returns a boolean if a field has been set.
func (o *ServiceConfiguration) HasRequiredGroups() bool {
	if o != nil && o.RequiredGroups != nil {
		return true
	}

	return false
}

// SetRequiredGroups gets a reference to the given []int32 and assigns it to the RequiredGroups field.
func (o *ServiceConfiguration) SetRequiredGroups(v []int32) {
	o.RequiredGroups = &v
}

// GetProperties returns the Properties field value
func (o *ServiceConfiguration) GetProperties() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *ServiceConfiguration) GetPropertiesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *ServiceConfiguration) SetProperties(v []string) {
	o.Properties = v
}

// GetShowPassword returns the ShowPassword field value if set, zero value otherwise.
func (o *ServiceConfiguration) GetShowPassword() bool {
	if o == nil || o.ShowPassword == nil {
		var ret bool
		return ret
	}
	return *o.ShowPassword
}

// GetShowPasswordOk returns a tuple with the ShowPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceConfiguration) GetShowPasswordOk() (*bool, bool) {
	if o == nil || o.ShowPassword == nil {
		return nil, false
	}
	return o.ShowPassword, true
}

// HasShowPassword returns a boolean if a field has been set.
func (o *ServiceConfiguration) HasShowPassword() bool {
	if o != nil && o.ShowPassword != nil {
		return true
	}

	return false
}

// SetShowPassword gets a reference to the given bool and assigns it to the ShowPassword field.
func (o *ServiceConfiguration) SetShowPassword(v bool) {
	o.ShowPassword = &v
}

// GetActions returns the Actions field value
func (o *ServiceConfiguration) GetActions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value
// and a boolean to check if the value has been set.
func (o *ServiceConfiguration) GetActionsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Actions, true
}

// SetActions sets field value
func (o *ServiceConfiguration) SetActions(v []string) {
	o.Actions = v
}

// GetURLs returns the URLs field value
func (o *ServiceConfiguration) GetURLs() []ServiceConfigurationURL {
	if o == nil {
		var ret []ServiceConfigurationURL
		return ret
	}

	return o.URLs
}

// GetURLsOk returns a tuple with the URLs field value
// and a boolean to check if the value has been set.
func (o *ServiceConfiguration) GetURLsOk() (*[]ServiceConfigurationURL, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.URLs, true
}

// SetURLs sets field value
func (o *ServiceConfiguration) SetURLs(v []ServiceConfigurationURL) {
	o.URLs = v
}

// GetTextAccount returns the TextAccount field value
func (o *ServiceConfiguration) GetTextAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TextAccount
}

// GetTextAccountOk returns a tuple with the TextAccount field value
// and a boolean to check if the value has been set.
func (o *ServiceConfiguration) GetTextAccountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TextAccount, true
}

// SetTextAccount sets field value
func (o *ServiceConfiguration) SetTextAccount(v string) {
	o.TextAccount = v
}

// GetTextTop returns the TextTop field value
func (o *ServiceConfiguration) GetTextTop() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TextTop
}

// GetTextTopOk returns a tuple with the TextTop field value
// and a boolean to check if the value has been set.
func (o *ServiceConfiguration) GetTextTopOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TextTop, true
}

// SetTextTop sets field value
func (o *ServiceConfiguration) SetTextTop(v string) {
	o.TextTop = v
}

// GetTextRegister returns the TextRegister field value
func (o *ServiceConfiguration) GetTextRegister() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TextRegister
}

// GetTextRegisterOk returns a tuple with the TextRegister field value
// and a boolean to check if the value has been set.
func (o *ServiceConfiguration) GetTextRegisterOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TextRegister, true
}

// SetTextRegister sets field value
func (o *ServiceConfiguration) SetTextRegister(v string) {
	o.TextRegister = v
}

// GetTextPending returns the TextPending field value
func (o *ServiceConfiguration) GetTextPending() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TextPending
}

// GetTextPendingOk returns a tuple with the TextPending field value
// and a boolean to check if the value has been set.
func (o *ServiceConfiguration) GetTextPendingOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TextPending, true
}

// SetTextPending sets field value
func (o *ServiceConfiguration) SetTextPending(v string) {
	o.TextPending = v
}

func (o ServiceConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PhpClass != nil {
		toSerialize["phpClass"] = o.PhpClass
	}
	if o.Psr4Prefix != nil {
		toSerialize["psr4Prefix"] = o.Psr4Prefix
	}
	if o.Psr4Path != nil {
		toSerialize["psr4Path"] = o.Psr4Path
	}
	if o.RequiredGroups != nil {
		toSerialize["requiredGroups"] = o.RequiredGroups
	}
	if true {
		toSerialize["properties"] = o.Properties
	}
	if o.ShowPassword != nil {
		toSerialize["showPassword"] = o.ShowPassword
	}
	if true {
		toSerialize["actions"] = o.Actions
	}
	if true {
		toSerialize["URLs"] = o.URLs
	}
	if true {
		toSerialize["textAccount"] = o.TextAccount
	}
	if true {
		toSerialize["textTop"] = o.TextTop
	}
	if true {
		toSerialize["textRegister"] = o.TextRegister
	}
	if true {
		toSerialize["textPending"] = o.TextPending
	}
	return json.Marshal(toSerialize)
}

type NullableServiceConfiguration struct {
	value *ServiceConfiguration
	isSet bool
}

func (v NullableServiceConfiguration) Get() *ServiceConfiguration {
	return v.value
}

func (v *NullableServiceConfiguration) Set(val *ServiceConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceConfiguration(val *ServiceConfiguration) *NullableServiceConfiguration {
	return &NullableServiceConfiguration{value: val, isSet: true}
}

func (v NullableServiceConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


