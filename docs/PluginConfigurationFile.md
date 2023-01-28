# PluginConfigurationFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DirectoryName** | Pointer to **string** | Directory where the plugin.yml file is stored.  Only from database but always set when the data from the file is read. | [optional] 
**URLs** | Pointer to [**[]PluginConfigurationURL**](PluginConfigurationURL.md) |  | [optional] 
**TextTop** | Pointer to **string** |  | [optional] 
**TextAccount** | Pointer to **string** |  | [optional] 
**TextRegister** | Pointer to **string** |  | [optional] 
**TextPending** | Pointer to **string** |  | [optional] 
**ConfigurationData** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Types** | Pointer to **[]string** | Not part of the file but will be set when the plugin implementation is loaded. | [optional] 
**OneAccount** | Pointer to **bool** |  | [optional] 
**Properties** | **[]string** |  | 
**ShowPassword** | Pointer to **bool** |  | [optional] 
**Actions** | **[]string** |  | 

## Methods

### NewPluginConfigurationFile

`func NewPluginConfigurationFile(properties []string, actions []string, ) *PluginConfigurationFile`

NewPluginConfigurationFile instantiates a new PluginConfigurationFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginConfigurationFileWithDefaults

`func NewPluginConfigurationFileWithDefaults() *PluginConfigurationFile`

NewPluginConfigurationFileWithDefaults instantiates a new PluginConfigurationFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirectoryName

`func (o *PluginConfigurationFile) GetDirectoryName() string`

GetDirectoryName returns the DirectoryName field if non-nil, zero value otherwise.

### GetDirectoryNameOk

`func (o *PluginConfigurationFile) GetDirectoryNameOk() (*string, bool)`

GetDirectoryNameOk returns a tuple with the DirectoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryName

`func (o *PluginConfigurationFile) SetDirectoryName(v string)`

SetDirectoryName sets DirectoryName field to given value.

### HasDirectoryName

`func (o *PluginConfigurationFile) HasDirectoryName() bool`

HasDirectoryName returns a boolean if a field has been set.

### GetURLs

`func (o *PluginConfigurationFile) GetURLs() []PluginConfigurationURL`

GetURLs returns the URLs field if non-nil, zero value otherwise.

### GetURLsOk

`func (o *PluginConfigurationFile) GetURLsOk() (*[]PluginConfigurationURL, bool)`

GetURLsOk returns a tuple with the URLs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetURLs

`func (o *PluginConfigurationFile) SetURLs(v []PluginConfigurationURL)`

SetURLs sets URLs field to given value.

### HasURLs

`func (o *PluginConfigurationFile) HasURLs() bool`

HasURLs returns a boolean if a field has been set.

### GetTextTop

`func (o *PluginConfigurationFile) GetTextTop() string`

GetTextTop returns the TextTop field if non-nil, zero value otherwise.

### GetTextTopOk

`func (o *PluginConfigurationFile) GetTextTopOk() (*string, bool)`

GetTextTopOk returns a tuple with the TextTop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextTop

`func (o *PluginConfigurationFile) SetTextTop(v string)`

SetTextTop sets TextTop field to given value.

### HasTextTop

`func (o *PluginConfigurationFile) HasTextTop() bool`

HasTextTop returns a boolean if a field has been set.

### GetTextAccount

`func (o *PluginConfigurationFile) GetTextAccount() string`

GetTextAccount returns the TextAccount field if non-nil, zero value otherwise.

### GetTextAccountOk

`func (o *PluginConfigurationFile) GetTextAccountOk() (*string, bool)`

GetTextAccountOk returns a tuple with the TextAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextAccount

`func (o *PluginConfigurationFile) SetTextAccount(v string)`

SetTextAccount sets TextAccount field to given value.

### HasTextAccount

`func (o *PluginConfigurationFile) HasTextAccount() bool`

HasTextAccount returns a boolean if a field has been set.

### GetTextRegister

`func (o *PluginConfigurationFile) GetTextRegister() string`

GetTextRegister returns the TextRegister field if non-nil, zero value otherwise.

### GetTextRegisterOk

`func (o *PluginConfigurationFile) GetTextRegisterOk() (*string, bool)`

GetTextRegisterOk returns a tuple with the TextRegister field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextRegister

`func (o *PluginConfigurationFile) SetTextRegister(v string)`

SetTextRegister sets TextRegister field to given value.

### HasTextRegister

`func (o *PluginConfigurationFile) HasTextRegister() bool`

HasTextRegister returns a boolean if a field has been set.

### GetTextPending

`func (o *PluginConfigurationFile) GetTextPending() string`

GetTextPending returns the TextPending field if non-nil, zero value otherwise.

### GetTextPendingOk

`func (o *PluginConfigurationFile) GetTextPendingOk() (*string, bool)`

GetTextPendingOk returns a tuple with the TextPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextPending

`func (o *PluginConfigurationFile) SetTextPending(v string)`

SetTextPending sets TextPending field to given value.

### HasTextPending

`func (o *PluginConfigurationFile) HasTextPending() bool`

HasTextPending returns a boolean if a field has been set.

### GetConfigurationData

`func (o *PluginConfigurationFile) GetConfigurationData() string`

GetConfigurationData returns the ConfigurationData field if non-nil, zero value otherwise.

### GetConfigurationDataOk

`func (o *PluginConfigurationFile) GetConfigurationDataOk() (*string, bool)`

GetConfigurationDataOk returns a tuple with the ConfigurationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationData

`func (o *PluginConfigurationFile) SetConfigurationData(v string)`

SetConfigurationData sets ConfigurationData field to given value.

### HasConfigurationData

`func (o *PluginConfigurationFile) HasConfigurationData() bool`

HasConfigurationData returns a boolean if a field has been set.

### GetName

`func (o *PluginConfigurationFile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PluginConfigurationFile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PluginConfigurationFile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PluginConfigurationFile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTypes

`func (o *PluginConfigurationFile) GetTypes() []string`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *PluginConfigurationFile) GetTypesOk() (*[]string, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *PluginConfigurationFile) SetTypes(v []string)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *PluginConfigurationFile) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### GetOneAccount

`func (o *PluginConfigurationFile) GetOneAccount() bool`

GetOneAccount returns the OneAccount field if non-nil, zero value otherwise.

### GetOneAccountOk

`func (o *PluginConfigurationFile) GetOneAccountOk() (*bool, bool)`

GetOneAccountOk returns a tuple with the OneAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneAccount

`func (o *PluginConfigurationFile) SetOneAccount(v bool)`

SetOneAccount sets OneAccount field to given value.

### HasOneAccount

`func (o *PluginConfigurationFile) HasOneAccount() bool`

HasOneAccount returns a boolean if a field has been set.

### GetProperties

`func (o *PluginConfigurationFile) GetProperties() []string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *PluginConfigurationFile) GetPropertiesOk() (*[]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *PluginConfigurationFile) SetProperties(v []string)`

SetProperties sets Properties field to given value.


### GetShowPassword

`func (o *PluginConfigurationFile) GetShowPassword() bool`

GetShowPassword returns the ShowPassword field if non-nil, zero value otherwise.

### GetShowPasswordOk

`func (o *PluginConfigurationFile) GetShowPasswordOk() (*bool, bool)`

GetShowPasswordOk returns a tuple with the ShowPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowPassword

`func (o *PluginConfigurationFile) SetShowPassword(v bool)`

SetShowPassword sets ShowPassword field to given value.

### HasShowPassword

`func (o *PluginConfigurationFile) HasShowPassword() bool`

HasShowPassword returns a boolean if a field has been set.

### GetActions

`func (o *PluginConfigurationFile) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *PluginConfigurationFile) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *PluginConfigurationFile) SetActions(v []string)`

SetActions sets Actions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


