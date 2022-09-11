/*
Neucore API

Client library of Neucore API

API version: 1.36.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package neucoreapi

import (
	"encoding/json"
)

// ServiceConfiguration struct for ServiceConfiguration
type ServiceConfiguration struct {
	PhpClass NullableString `json:"phpClass,omitempty"`
	Psr4Prefix NullableString `json:"psr4Prefix,omitempty"`
	Psr4Path NullableString `json:"psr4Path,omitempty"`
	OneAccount NullableBool `json:"oneAccount,omitempty"`
	RequiredGroups []int32 `json:"requiredGroups,omitempty"`
	Properties []string `json:"properties"`
	ShowPassword NullableBool `json:"showPassword,omitempty"`
	Actions []string `json:"actions"`
	URLs []ServiceConfigurationURL `json:"URLs"`
	TextAccount NullableString `json:"textAccount"`
	TextTop NullableString `json:"textTop"`
	TextRegister NullableString `json:"textRegister"`
	TextPending NullableString `json:"textPending"`
	ConfigurationData string `json:"configurationData"`
}

// NewServiceConfiguration instantiates a new ServiceConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceConfiguration(properties []string, actions []string, uRLs []ServiceConfigurationURL, textAccount NullableString, textTop NullableString, textRegister NullableString, textPending NullableString, configurationData string) *ServiceConfiguration {
	this := ServiceConfiguration{}
	this.Properties = properties
	this.Actions = actions
	this.URLs = uRLs
	this.TextAccount = textAccount
	this.TextTop = textTop
	this.TextRegister = textRegister
	this.TextPending = textPending
	this.ConfigurationData = configurationData
	return &this
}

// NewServiceConfigurationWithDefaults instantiates a new ServiceConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceConfigurationWithDefaults() *ServiceConfiguration {
	this := ServiceConfiguration{}
	return &this
}

// GetPhpClass returns the PhpClass field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceConfiguration) GetPhpClass() string {
	if o == nil || o.PhpClass.Get() == nil {
		var ret string
		return ret
	}
	return *o.PhpClass.Get()
}

// GetPhpClassOk returns a tuple with the PhpClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceConfiguration) GetPhpClassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PhpClass.Get(), o.PhpClass.IsSet()
}

// HasPhpClass returns a boolean if a field has been set.
func (o *ServiceConfiguration) HasPhpClass() bool {
	if o != nil && o.PhpClass.IsSet() {
		return true
	}

	return false
}

// SetPhpClass gets a reference to the given NullableString and assigns it to the PhpClass field.
func (o *ServiceConfiguration) SetPhpClass(v string) {
	o.PhpClass.Set(&v)
}
// SetPhpClassNil sets the value for PhpClass to be an explicit nil
func (o *ServiceConfiguration) SetPhpClassNil() {
	o.PhpClass.Set(nil)
}

// UnsetPhpClass ensures that no value is present for PhpClass, not even an explicit nil
func (o *ServiceConfiguration) UnsetPhpClass() {
	o.PhpClass.Unset()
}

// GetPsr4Prefix returns the Psr4Prefix field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceConfiguration) GetPsr4Prefix() string {
	if o == nil || o.Psr4Prefix.Get() == nil {
		var ret string
		return ret
	}
	return *o.Psr4Prefix.Get()
}

// GetPsr4PrefixOk returns a tuple with the Psr4Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceConfiguration) GetPsr4PrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Psr4Prefix.Get(), o.Psr4Prefix.IsSet()
}

// HasPsr4Prefix returns a boolean if a field has been set.
func (o *ServiceConfiguration) HasPsr4Prefix() bool {
	if o != nil && o.Psr4Prefix.IsSet() {
		return true
	}

	return false
}

// SetPsr4Prefix gets a reference to the given NullableString and assigns it to the Psr4Prefix field.
func (o *ServiceConfiguration) SetPsr4Prefix(v string) {
	o.Psr4Prefix.Set(&v)
}
// SetPsr4PrefixNil sets the value for Psr4Prefix to be an explicit nil
func (o *ServiceConfiguration) SetPsr4PrefixNil() {
	o.Psr4Prefix.Set(nil)
}

// UnsetPsr4Prefix ensures that no value is present for Psr4Prefix, not even an explicit nil
func (o *ServiceConfiguration) UnsetPsr4Prefix() {
	o.Psr4Prefix.Unset()
}

// GetPsr4Path returns the Psr4Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceConfiguration) GetPsr4Path() string {
	if o == nil || o.Psr4Path.Get() == nil {
		var ret string
		return ret
	}
	return *o.Psr4Path.Get()
}

// GetPsr4PathOk returns a tuple with the Psr4Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceConfiguration) GetPsr4PathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Psr4Path.Get(), o.Psr4Path.IsSet()
}

// HasPsr4Path returns a boolean if a field has been set.
func (o *ServiceConfiguration) HasPsr4Path() bool {
	if o != nil && o.Psr4Path.IsSet() {
		return true
	}

	return false
}

// SetPsr4Path gets a reference to the given NullableString and assigns it to the Psr4Path field.
func (o *ServiceConfiguration) SetPsr4Path(v string) {
	o.Psr4Path.Set(&v)
}
// SetPsr4PathNil sets the value for Psr4Path to be an explicit nil
func (o *ServiceConfiguration) SetPsr4PathNil() {
	o.Psr4Path.Set(nil)
}

// UnsetPsr4Path ensures that no value is present for Psr4Path, not even an explicit nil
func (o *ServiceConfiguration) UnsetPsr4Path() {
	o.Psr4Path.Unset()
}

// GetOneAccount returns the OneAccount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceConfiguration) GetOneAccount() bool {
	if o == nil || o.OneAccount.Get() == nil {
		var ret bool
		return ret
	}
	return *o.OneAccount.Get()
}

// GetOneAccountOk returns a tuple with the OneAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceConfiguration) GetOneAccountOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.OneAccount.Get(), o.OneAccount.IsSet()
}

// HasOneAccount returns a boolean if a field has been set.
func (o *ServiceConfiguration) HasOneAccount() bool {
	if o != nil && o.OneAccount.IsSet() {
		return true
	}

	return false
}

// SetOneAccount gets a reference to the given NullableBool and assigns it to the OneAccount field.
func (o *ServiceConfiguration) SetOneAccount(v bool) {
	o.OneAccount.Set(&v)
}
// SetOneAccountNil sets the value for OneAccount to be an explicit nil
func (o *ServiceConfiguration) SetOneAccountNil() {
	o.OneAccount.Set(nil)
}

// UnsetOneAccount ensures that no value is present for OneAccount, not even an explicit nil
func (o *ServiceConfiguration) UnsetOneAccount() {
	o.OneAccount.Unset()
}

// GetRequiredGroups returns the RequiredGroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceConfiguration) GetRequiredGroups() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.RequiredGroups
}

// GetRequiredGroupsOk returns a tuple with the RequiredGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceConfiguration) GetRequiredGroupsOk() ([]int32, bool) {
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
	o.RequiredGroups = v
}

// GetProperties returns the Properties field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *ServiceConfiguration) GetProperties() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceConfiguration) GetPropertiesOk() ([]string, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// SetProperties sets field value
func (o *ServiceConfiguration) SetProperties(v []string) {
	o.Properties = v
}

// GetShowPassword returns the ShowPassword field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceConfiguration) GetShowPassword() bool {
	if o == nil || o.ShowPassword.Get() == nil {
		var ret bool
		return ret
	}
	return *o.ShowPassword.Get()
}

// GetShowPasswordOk returns a tuple with the ShowPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceConfiguration) GetShowPasswordOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShowPassword.Get(), o.ShowPassword.IsSet()
}

// HasShowPassword returns a boolean if a field has been set.
func (o *ServiceConfiguration) HasShowPassword() bool {
	if o != nil && o.ShowPassword.IsSet() {
		return true
	}

	return false
}

// SetShowPassword gets a reference to the given NullableBool and assigns it to the ShowPassword field.
func (o *ServiceConfiguration) SetShowPassword(v bool) {
	o.ShowPassword.Set(&v)
}
// SetShowPasswordNil sets the value for ShowPassword to be an explicit nil
func (o *ServiceConfiguration) SetShowPasswordNil() {
	o.ShowPassword.Set(nil)
}

// UnsetShowPassword ensures that no value is present for ShowPassword, not even an explicit nil
func (o *ServiceConfiguration) UnsetShowPassword() {
	o.ShowPassword.Unset()
}

// GetActions returns the Actions field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *ServiceConfiguration) GetActions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceConfiguration) GetActionsOk() ([]string, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return o.Actions, true
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
func (o *ServiceConfiguration) GetURLsOk() ([]ServiceConfigurationURL, bool) {
	if o == nil {
		return nil, false
	}
	return o.URLs, true
}

// SetURLs sets field value
func (o *ServiceConfiguration) SetURLs(v []ServiceConfigurationURL) {
	o.URLs = v
}

// GetTextAccount returns the TextAccount field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ServiceConfiguration) GetTextAccount() string {
	if o == nil || o.TextAccount.Get() == nil {
		var ret string
		return ret
	}

	return *o.TextAccount.Get()
}

// GetTextAccountOk returns a tuple with the TextAccount field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceConfiguration) GetTextAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TextAccount.Get(), o.TextAccount.IsSet()
}

// SetTextAccount sets field value
func (o *ServiceConfiguration) SetTextAccount(v string) {
	o.TextAccount.Set(&v)
}

// GetTextTop returns the TextTop field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ServiceConfiguration) GetTextTop() string {
	if o == nil || o.TextTop.Get() == nil {
		var ret string
		return ret
	}

	return *o.TextTop.Get()
}

// GetTextTopOk returns a tuple with the TextTop field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceConfiguration) GetTextTopOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TextTop.Get(), o.TextTop.IsSet()
}

// SetTextTop sets field value
func (o *ServiceConfiguration) SetTextTop(v string) {
	o.TextTop.Set(&v)
}

// GetTextRegister returns the TextRegister field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ServiceConfiguration) GetTextRegister() string {
	if o == nil || o.TextRegister.Get() == nil {
		var ret string
		return ret
	}

	return *o.TextRegister.Get()
}

// GetTextRegisterOk returns a tuple with the TextRegister field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceConfiguration) GetTextRegisterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TextRegister.Get(), o.TextRegister.IsSet()
}

// SetTextRegister sets field value
func (o *ServiceConfiguration) SetTextRegister(v string) {
	o.TextRegister.Set(&v)
}

// GetTextPending returns the TextPending field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ServiceConfiguration) GetTextPending() string {
	if o == nil || o.TextPending.Get() == nil {
		var ret string
		return ret
	}

	return *o.TextPending.Get()
}

// GetTextPendingOk returns a tuple with the TextPending field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceConfiguration) GetTextPendingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TextPending.Get(), o.TextPending.IsSet()
}

// SetTextPending sets field value
func (o *ServiceConfiguration) SetTextPending(v string) {
	o.TextPending.Set(&v)
}

// GetConfigurationData returns the ConfigurationData field value
func (o *ServiceConfiguration) GetConfigurationData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConfigurationData
}

// GetConfigurationDataOk returns a tuple with the ConfigurationData field value
// and a boolean to check if the value has been set.
func (o *ServiceConfiguration) GetConfigurationDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigurationData, true
}

// SetConfigurationData sets field value
func (o *ServiceConfiguration) SetConfigurationData(v string) {
	o.ConfigurationData = v
}

func (o ServiceConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PhpClass.IsSet() {
		toSerialize["phpClass"] = o.PhpClass.Get()
	}
	if o.Psr4Prefix.IsSet() {
		toSerialize["psr4Prefix"] = o.Psr4Prefix.Get()
	}
	if o.Psr4Path.IsSet() {
		toSerialize["psr4Path"] = o.Psr4Path.Get()
	}
	if o.OneAccount.IsSet() {
		toSerialize["oneAccount"] = o.OneAccount.Get()
	}
	if o.RequiredGroups != nil {
		toSerialize["requiredGroups"] = o.RequiredGroups
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if o.ShowPassword.IsSet() {
		toSerialize["showPassword"] = o.ShowPassword.Get()
	}
	if o.Actions != nil {
		toSerialize["actions"] = o.Actions
	}
	if true {
		toSerialize["URLs"] = o.URLs
	}
	if true {
		toSerialize["textAccount"] = o.TextAccount.Get()
	}
	if true {
		toSerialize["textTop"] = o.TextTop.Get()
	}
	if true {
		toSerialize["textRegister"] = o.TextRegister.Get()
	}
	if true {
		toSerialize["textPending"] = o.TextPending.Get()
	}
	if true {
		toSerialize["configurationData"] = o.ConfigurationData
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


