/*
Neucore API

Client library of Neucore API

API version: 2.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package neucoreapi

import (
	"encoding/json"
)

// checks if the PluginConfigurationFile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PluginConfigurationFile{}

// PluginConfigurationFile Plugin configuration from YAML file.  API: The required properties are necessary for the service page where users register their account. The rest is necessary for the admin page.
type PluginConfigurationFile struct {
	// Directory where the plugin.yml file is stored.  Only from database but always set when the data from the file is read.
	DirectoryName *string `json:"directoryName,omitempty"`
	URLs []PluginConfigurationURL `json:"URLs,omitempty"`
	TextTop *string `json:"textTop,omitempty"`
	TextAccount *string `json:"textAccount,omitempty"`
	TextRegister *string `json:"textRegister,omitempty"`
	TextPending *string `json:"textPending,omitempty"`
	ConfigurationData *string `json:"configurationData,omitempty"`
	Name *string `json:"name,omitempty"`
	// Not part of the file but will be set when the plugin implementation is loaded.
	Types []string `json:"types,omitempty"`
	OneAccount *bool `json:"oneAccount,omitempty"`
	Properties []string `json:"properties"`
	ShowPassword *bool `json:"showPassword,omitempty"`
	Actions []string `json:"actions"`
}

// NewPluginConfigurationFile instantiates a new PluginConfigurationFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPluginConfigurationFile(properties []string, actions []string) *PluginConfigurationFile {
	this := PluginConfigurationFile{}
	this.Properties = properties
	this.Actions = actions
	return &this
}

// NewPluginConfigurationFileWithDefaults instantiates a new PluginConfigurationFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPluginConfigurationFileWithDefaults() *PluginConfigurationFile {
	this := PluginConfigurationFile{}
	return &this
}

// GetDirectoryName returns the DirectoryName field value if set, zero value otherwise.
func (o *PluginConfigurationFile) GetDirectoryName() string {
	if o == nil || IsNil(o.DirectoryName) {
		var ret string
		return ret
	}
	return *o.DirectoryName
}

// GetDirectoryNameOk returns a tuple with the DirectoryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginConfigurationFile) GetDirectoryNameOk() (*string, bool) {
	if o == nil || IsNil(o.DirectoryName) {
		return nil, false
	}
	return o.DirectoryName, true
}

// HasDirectoryName returns a boolean if a field has been set.
func (o *PluginConfigurationFile) HasDirectoryName() bool {
	if o != nil && !IsNil(o.DirectoryName) {
		return true
	}

	return false
}

// SetDirectoryName gets a reference to the given string and assigns it to the DirectoryName field.
func (o *PluginConfigurationFile) SetDirectoryName(v string) {
	o.DirectoryName = &v
}

// GetURLs returns the URLs field value if set, zero value otherwise.
func (o *PluginConfigurationFile) GetURLs() []PluginConfigurationURL {
	if o == nil || IsNil(o.URLs) {
		var ret []PluginConfigurationURL
		return ret
	}
	return o.URLs
}

// GetURLsOk returns a tuple with the URLs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginConfigurationFile) GetURLsOk() ([]PluginConfigurationURL, bool) {
	if o == nil || IsNil(o.URLs) {
		return nil, false
	}
	return o.URLs, true
}

// HasURLs returns a boolean if a field has been set.
func (o *PluginConfigurationFile) HasURLs() bool {
	if o != nil && !IsNil(o.URLs) {
		return true
	}

	return false
}

// SetURLs gets a reference to the given []PluginConfigurationURL and assigns it to the URLs field.
func (o *PluginConfigurationFile) SetURLs(v []PluginConfigurationURL) {
	o.URLs = v
}

// GetTextTop returns the TextTop field value if set, zero value otherwise.
func (o *PluginConfigurationFile) GetTextTop() string {
	if o == nil || IsNil(o.TextTop) {
		var ret string
		return ret
	}
	return *o.TextTop
}

// GetTextTopOk returns a tuple with the TextTop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginConfigurationFile) GetTextTopOk() (*string, bool) {
	if o == nil || IsNil(o.TextTop) {
		return nil, false
	}
	return o.TextTop, true
}

// HasTextTop returns a boolean if a field has been set.
func (o *PluginConfigurationFile) HasTextTop() bool {
	if o != nil && !IsNil(o.TextTop) {
		return true
	}

	return false
}

// SetTextTop gets a reference to the given string and assigns it to the TextTop field.
func (o *PluginConfigurationFile) SetTextTop(v string) {
	o.TextTop = &v
}

// GetTextAccount returns the TextAccount field value if set, zero value otherwise.
func (o *PluginConfigurationFile) GetTextAccount() string {
	if o == nil || IsNil(o.TextAccount) {
		var ret string
		return ret
	}
	return *o.TextAccount
}

// GetTextAccountOk returns a tuple with the TextAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginConfigurationFile) GetTextAccountOk() (*string, bool) {
	if o == nil || IsNil(o.TextAccount) {
		return nil, false
	}
	return o.TextAccount, true
}

// HasTextAccount returns a boolean if a field has been set.
func (o *PluginConfigurationFile) HasTextAccount() bool {
	if o != nil && !IsNil(o.TextAccount) {
		return true
	}

	return false
}

// SetTextAccount gets a reference to the given string and assigns it to the TextAccount field.
func (o *PluginConfigurationFile) SetTextAccount(v string) {
	o.TextAccount = &v
}

// GetTextRegister returns the TextRegister field value if set, zero value otherwise.
func (o *PluginConfigurationFile) GetTextRegister() string {
	if o == nil || IsNil(o.TextRegister) {
		var ret string
		return ret
	}
	return *o.TextRegister
}

// GetTextRegisterOk returns a tuple with the TextRegister field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginConfigurationFile) GetTextRegisterOk() (*string, bool) {
	if o == nil || IsNil(o.TextRegister) {
		return nil, false
	}
	return o.TextRegister, true
}

// HasTextRegister returns a boolean if a field has been set.
func (o *PluginConfigurationFile) HasTextRegister() bool {
	if o != nil && !IsNil(o.TextRegister) {
		return true
	}

	return false
}

// SetTextRegister gets a reference to the given string and assigns it to the TextRegister field.
func (o *PluginConfigurationFile) SetTextRegister(v string) {
	o.TextRegister = &v
}

// GetTextPending returns the TextPending field value if set, zero value otherwise.
func (o *PluginConfigurationFile) GetTextPending() string {
	if o == nil || IsNil(o.TextPending) {
		var ret string
		return ret
	}
	return *o.TextPending
}

// GetTextPendingOk returns a tuple with the TextPending field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginConfigurationFile) GetTextPendingOk() (*string, bool) {
	if o == nil || IsNil(o.TextPending) {
		return nil, false
	}
	return o.TextPending, true
}

// HasTextPending returns a boolean if a field has been set.
func (o *PluginConfigurationFile) HasTextPending() bool {
	if o != nil && !IsNil(o.TextPending) {
		return true
	}

	return false
}

// SetTextPending gets a reference to the given string and assigns it to the TextPending field.
func (o *PluginConfigurationFile) SetTextPending(v string) {
	o.TextPending = &v
}

// GetConfigurationData returns the ConfigurationData field value if set, zero value otherwise.
func (o *PluginConfigurationFile) GetConfigurationData() string {
	if o == nil || IsNil(o.ConfigurationData) {
		var ret string
		return ret
	}
	return *o.ConfigurationData
}

// GetConfigurationDataOk returns a tuple with the ConfigurationData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginConfigurationFile) GetConfigurationDataOk() (*string, bool) {
	if o == nil || IsNil(o.ConfigurationData) {
		return nil, false
	}
	return o.ConfigurationData, true
}

// HasConfigurationData returns a boolean if a field has been set.
func (o *PluginConfigurationFile) HasConfigurationData() bool {
	if o != nil && !IsNil(o.ConfigurationData) {
		return true
	}

	return false
}

// SetConfigurationData gets a reference to the given string and assigns it to the ConfigurationData field.
func (o *PluginConfigurationFile) SetConfigurationData(v string) {
	o.ConfigurationData = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PluginConfigurationFile) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginConfigurationFile) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PluginConfigurationFile) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PluginConfigurationFile) SetName(v string) {
	o.Name = &v
}

// GetTypes returns the Types field value if set, zero value otherwise.
func (o *PluginConfigurationFile) GetTypes() []string {
	if o == nil || IsNil(o.Types) {
		var ret []string
		return ret
	}
	return o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginConfigurationFile) GetTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.Types) {
		return nil, false
	}
	return o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *PluginConfigurationFile) HasTypes() bool {
	if o != nil && !IsNil(o.Types) {
		return true
	}

	return false
}

// SetTypes gets a reference to the given []string and assigns it to the Types field.
func (o *PluginConfigurationFile) SetTypes(v []string) {
	o.Types = v
}

// GetOneAccount returns the OneAccount field value if set, zero value otherwise.
func (o *PluginConfigurationFile) GetOneAccount() bool {
	if o == nil || IsNil(o.OneAccount) {
		var ret bool
		return ret
	}
	return *o.OneAccount
}

// GetOneAccountOk returns a tuple with the OneAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginConfigurationFile) GetOneAccountOk() (*bool, bool) {
	if o == nil || IsNil(o.OneAccount) {
		return nil, false
	}
	return o.OneAccount, true
}

// HasOneAccount returns a boolean if a field has been set.
func (o *PluginConfigurationFile) HasOneAccount() bool {
	if o != nil && !IsNil(o.OneAccount) {
		return true
	}

	return false
}

// SetOneAccount gets a reference to the given bool and assigns it to the OneAccount field.
func (o *PluginConfigurationFile) SetOneAccount(v bool) {
	o.OneAccount = &v
}

// GetProperties returns the Properties field value
func (o *PluginConfigurationFile) GetProperties() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *PluginConfigurationFile) GetPropertiesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties, true
}

// SetProperties sets field value
func (o *PluginConfigurationFile) SetProperties(v []string) {
	o.Properties = v
}

// GetShowPassword returns the ShowPassword field value if set, zero value otherwise.
func (o *PluginConfigurationFile) GetShowPassword() bool {
	if o == nil || IsNil(o.ShowPassword) {
		var ret bool
		return ret
	}
	return *o.ShowPassword
}

// GetShowPasswordOk returns a tuple with the ShowPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginConfigurationFile) GetShowPasswordOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowPassword) {
		return nil, false
	}
	return o.ShowPassword, true
}

// HasShowPassword returns a boolean if a field has been set.
func (o *PluginConfigurationFile) HasShowPassword() bool {
	if o != nil && !IsNil(o.ShowPassword) {
		return true
	}

	return false
}

// SetShowPassword gets a reference to the given bool and assigns it to the ShowPassword field.
func (o *PluginConfigurationFile) SetShowPassword(v bool) {
	o.ShowPassword = &v
}

// GetActions returns the Actions field value
func (o *PluginConfigurationFile) GetActions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value
// and a boolean to check if the value has been set.
func (o *PluginConfigurationFile) GetActionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Actions, true
}

// SetActions sets field value
func (o *PluginConfigurationFile) SetActions(v []string) {
	o.Actions = v
}

func (o PluginConfigurationFile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PluginConfigurationFile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DirectoryName) {
		toSerialize["directoryName"] = o.DirectoryName
	}
	if !IsNil(o.URLs) {
		toSerialize["URLs"] = o.URLs
	}
	if !IsNil(o.TextTop) {
		toSerialize["textTop"] = o.TextTop
	}
	if !IsNil(o.TextAccount) {
		toSerialize["textAccount"] = o.TextAccount
	}
	if !IsNil(o.TextRegister) {
		toSerialize["textRegister"] = o.TextRegister
	}
	if !IsNil(o.TextPending) {
		toSerialize["textPending"] = o.TextPending
	}
	if !IsNil(o.ConfigurationData) {
		toSerialize["configurationData"] = o.ConfigurationData
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Types) {
		toSerialize["types"] = o.Types
	}
	if !IsNil(o.OneAccount) {
		toSerialize["oneAccount"] = o.OneAccount
	}
	toSerialize["properties"] = o.Properties
	if !IsNil(o.ShowPassword) {
		toSerialize["showPassword"] = o.ShowPassword
	}
	toSerialize["actions"] = o.Actions
	return toSerialize, nil
}

type NullablePluginConfigurationFile struct {
	value *PluginConfigurationFile
	isSet bool
}

func (v NullablePluginConfigurationFile) Get() *PluginConfigurationFile {
	return v.value
}

func (v *NullablePluginConfigurationFile) Set(val *PluginConfigurationFile) {
	v.value = val
	v.isSet = true
}

func (v NullablePluginConfigurationFile) IsSet() bool {
	return v.isSet
}

func (v *NullablePluginConfigurationFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePluginConfigurationFile(val *PluginConfigurationFile) *NullablePluginConfigurationFile {
	return &NullablePluginConfigurationFile{value: val, isSet: true}
}

func (v NullablePluginConfigurationFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePluginConfigurationFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


