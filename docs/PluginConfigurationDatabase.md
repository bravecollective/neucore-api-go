# PluginConfigurationDatabase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DirectoryName** | Pointer to **string** | Directory where the plugin.yml file is stored.  Only from database but always set when the data from the file is read. | [optional] 
**URLs** | [**[]PluginConfigurationURL**](PluginConfigurationURL.md) |  | 
**TextTop** | **string** |  | 
**TextAccount** | **string** |  | 
**TextRegister** | **string** |  | 
**TextPending** | **string** |  | 
**ConfigurationData** | **string** |  | 
**Active** | Pointer to **bool** | Inactive plugins are neither updated by the cron job nor displayed to the user.  From admin UI. | [optional] 
**RequiredGroups** | Pointer to **[]int32** | From admin UI. | [optional] 

## Methods

### NewPluginConfigurationDatabase

`func NewPluginConfigurationDatabase(uRLs []PluginConfigurationURL, textTop string, textAccount string, textRegister string, textPending string, configurationData string, ) *PluginConfigurationDatabase`

NewPluginConfigurationDatabase instantiates a new PluginConfigurationDatabase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginConfigurationDatabaseWithDefaults

`func NewPluginConfigurationDatabaseWithDefaults() *PluginConfigurationDatabase`

NewPluginConfigurationDatabaseWithDefaults instantiates a new PluginConfigurationDatabase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirectoryName

`func (o *PluginConfigurationDatabase) GetDirectoryName() string`

GetDirectoryName returns the DirectoryName field if non-nil, zero value otherwise.

### GetDirectoryNameOk

`func (o *PluginConfigurationDatabase) GetDirectoryNameOk() (*string, bool)`

GetDirectoryNameOk returns a tuple with the DirectoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryName

`func (o *PluginConfigurationDatabase) SetDirectoryName(v string)`

SetDirectoryName sets DirectoryName field to given value.

### HasDirectoryName

`func (o *PluginConfigurationDatabase) HasDirectoryName() bool`

HasDirectoryName returns a boolean if a field has been set.

### GetURLs

`func (o *PluginConfigurationDatabase) GetURLs() []PluginConfigurationURL`

GetURLs returns the URLs field if non-nil, zero value otherwise.

### GetURLsOk

`func (o *PluginConfigurationDatabase) GetURLsOk() (*[]PluginConfigurationURL, bool)`

GetURLsOk returns a tuple with the URLs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetURLs

`func (o *PluginConfigurationDatabase) SetURLs(v []PluginConfigurationURL)`

SetURLs sets URLs field to given value.


### GetTextTop

`func (o *PluginConfigurationDatabase) GetTextTop() string`

GetTextTop returns the TextTop field if non-nil, zero value otherwise.

### GetTextTopOk

`func (o *PluginConfigurationDatabase) GetTextTopOk() (*string, bool)`

GetTextTopOk returns a tuple with the TextTop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextTop

`func (o *PluginConfigurationDatabase) SetTextTop(v string)`

SetTextTop sets TextTop field to given value.


### GetTextAccount

`func (o *PluginConfigurationDatabase) GetTextAccount() string`

GetTextAccount returns the TextAccount field if non-nil, zero value otherwise.

### GetTextAccountOk

`func (o *PluginConfigurationDatabase) GetTextAccountOk() (*string, bool)`

GetTextAccountOk returns a tuple with the TextAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextAccount

`func (o *PluginConfigurationDatabase) SetTextAccount(v string)`

SetTextAccount sets TextAccount field to given value.


### GetTextRegister

`func (o *PluginConfigurationDatabase) GetTextRegister() string`

GetTextRegister returns the TextRegister field if non-nil, zero value otherwise.

### GetTextRegisterOk

`func (o *PluginConfigurationDatabase) GetTextRegisterOk() (*string, bool)`

GetTextRegisterOk returns a tuple with the TextRegister field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextRegister

`func (o *PluginConfigurationDatabase) SetTextRegister(v string)`

SetTextRegister sets TextRegister field to given value.


### GetTextPending

`func (o *PluginConfigurationDatabase) GetTextPending() string`

GetTextPending returns the TextPending field if non-nil, zero value otherwise.

### GetTextPendingOk

`func (o *PluginConfigurationDatabase) GetTextPendingOk() (*string, bool)`

GetTextPendingOk returns a tuple with the TextPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextPending

`func (o *PluginConfigurationDatabase) SetTextPending(v string)`

SetTextPending sets TextPending field to given value.


### GetConfigurationData

`func (o *PluginConfigurationDatabase) GetConfigurationData() string`

GetConfigurationData returns the ConfigurationData field if non-nil, zero value otherwise.

### GetConfigurationDataOk

`func (o *PluginConfigurationDatabase) GetConfigurationDataOk() (*string, bool)`

GetConfigurationDataOk returns a tuple with the ConfigurationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationData

`func (o *PluginConfigurationDatabase) SetConfigurationData(v string)`

SetConfigurationData sets ConfigurationData field to given value.


### GetActive

`func (o *PluginConfigurationDatabase) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PluginConfigurationDatabase) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PluginConfigurationDatabase) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PluginConfigurationDatabase) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetRequiredGroups

`func (o *PluginConfigurationDatabase) GetRequiredGroups() []int32`

GetRequiredGroups returns the RequiredGroups field if non-nil, zero value otherwise.

### GetRequiredGroupsOk

`func (o *PluginConfigurationDatabase) GetRequiredGroupsOk() (*[]int32, bool)`

GetRequiredGroupsOk returns a tuple with the RequiredGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredGroups

`func (o *PluginConfigurationDatabase) SetRequiredGroups(v []int32)`

SetRequiredGroups sets RequiredGroups field to given value.

### HasRequiredGroups

`func (o *PluginConfigurationDatabase) HasRequiredGroups() bool`

HasRequiredGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


