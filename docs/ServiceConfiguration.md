# ServiceConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhpClass** | Pointer to **NullableString** |  | [optional] 
**Psr4Prefix** | Pointer to **NullableString** |  | [optional] 
**Psr4Path** | Pointer to **NullableString** |  | [optional] 
**OneAccount** | Pointer to **NullableBool** |  | [optional] 
**RequiredGroups** | Pointer to **[]int32** |  | [optional] 
**Properties** | **[]string** |  | 
**ShowPassword** | Pointer to **NullableBool** |  | [optional] 
**Actions** | **[]string** |  | 
**URLs** | [**[]ServiceConfigurationURL**](ServiceConfigurationURL.md) |  | 
**TextAccount** | **NullableString** |  | 
**TextTop** | **NullableString** |  | 
**TextRegister** | **NullableString** |  | 
**TextPending** | **NullableString** |  | 
**ConfigurationData** | **string** |  | 

## Methods

### NewServiceConfiguration

`func NewServiceConfiguration(properties []string, actions []string, uRLs []ServiceConfigurationURL, textAccount NullableString, textTop NullableString, textRegister NullableString, textPending NullableString, configurationData string, ) *ServiceConfiguration`

NewServiceConfiguration instantiates a new ServiceConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceConfigurationWithDefaults

`func NewServiceConfigurationWithDefaults() *ServiceConfiguration`

NewServiceConfigurationWithDefaults instantiates a new ServiceConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhpClass

`func (o *ServiceConfiguration) GetPhpClass() string`

GetPhpClass returns the PhpClass field if non-nil, zero value otherwise.

### GetPhpClassOk

`func (o *ServiceConfiguration) GetPhpClassOk() (*string, bool)`

GetPhpClassOk returns a tuple with the PhpClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhpClass

`func (o *ServiceConfiguration) SetPhpClass(v string)`

SetPhpClass sets PhpClass field to given value.

### HasPhpClass

`func (o *ServiceConfiguration) HasPhpClass() bool`

HasPhpClass returns a boolean if a field has been set.

### SetPhpClassNil

`func (o *ServiceConfiguration) SetPhpClassNil(b bool)`

 SetPhpClassNil sets the value for PhpClass to be an explicit nil

### UnsetPhpClass
`func (o *ServiceConfiguration) UnsetPhpClass()`

UnsetPhpClass ensures that no value is present for PhpClass, not even an explicit nil
### GetPsr4Prefix

`func (o *ServiceConfiguration) GetPsr4Prefix() string`

GetPsr4Prefix returns the Psr4Prefix field if non-nil, zero value otherwise.

### GetPsr4PrefixOk

`func (o *ServiceConfiguration) GetPsr4PrefixOk() (*string, bool)`

GetPsr4PrefixOk returns a tuple with the Psr4Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsr4Prefix

`func (o *ServiceConfiguration) SetPsr4Prefix(v string)`

SetPsr4Prefix sets Psr4Prefix field to given value.

### HasPsr4Prefix

`func (o *ServiceConfiguration) HasPsr4Prefix() bool`

HasPsr4Prefix returns a boolean if a field has been set.

### SetPsr4PrefixNil

`func (o *ServiceConfiguration) SetPsr4PrefixNil(b bool)`

 SetPsr4PrefixNil sets the value for Psr4Prefix to be an explicit nil

### UnsetPsr4Prefix
`func (o *ServiceConfiguration) UnsetPsr4Prefix()`

UnsetPsr4Prefix ensures that no value is present for Psr4Prefix, not even an explicit nil
### GetPsr4Path

`func (o *ServiceConfiguration) GetPsr4Path() string`

GetPsr4Path returns the Psr4Path field if non-nil, zero value otherwise.

### GetPsr4PathOk

`func (o *ServiceConfiguration) GetPsr4PathOk() (*string, bool)`

GetPsr4PathOk returns a tuple with the Psr4Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsr4Path

`func (o *ServiceConfiguration) SetPsr4Path(v string)`

SetPsr4Path sets Psr4Path field to given value.

### HasPsr4Path

`func (o *ServiceConfiguration) HasPsr4Path() bool`

HasPsr4Path returns a boolean if a field has been set.

### SetPsr4PathNil

`func (o *ServiceConfiguration) SetPsr4PathNil(b bool)`

 SetPsr4PathNil sets the value for Psr4Path to be an explicit nil

### UnsetPsr4Path
`func (o *ServiceConfiguration) UnsetPsr4Path()`

UnsetPsr4Path ensures that no value is present for Psr4Path, not even an explicit nil
### GetOneAccount

`func (o *ServiceConfiguration) GetOneAccount() bool`

GetOneAccount returns the OneAccount field if non-nil, zero value otherwise.

### GetOneAccountOk

`func (o *ServiceConfiguration) GetOneAccountOk() (*bool, bool)`

GetOneAccountOk returns a tuple with the OneAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneAccount

`func (o *ServiceConfiguration) SetOneAccount(v bool)`

SetOneAccount sets OneAccount field to given value.

### HasOneAccount

`func (o *ServiceConfiguration) HasOneAccount() bool`

HasOneAccount returns a boolean if a field has been set.

### SetOneAccountNil

`func (o *ServiceConfiguration) SetOneAccountNil(b bool)`

 SetOneAccountNil sets the value for OneAccount to be an explicit nil

### UnsetOneAccount
`func (o *ServiceConfiguration) UnsetOneAccount()`

UnsetOneAccount ensures that no value is present for OneAccount, not even an explicit nil
### GetRequiredGroups

`func (o *ServiceConfiguration) GetRequiredGroups() []int32`

GetRequiredGroups returns the RequiredGroups field if non-nil, zero value otherwise.

### GetRequiredGroupsOk

`func (o *ServiceConfiguration) GetRequiredGroupsOk() (*[]int32, bool)`

GetRequiredGroupsOk returns a tuple with the RequiredGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredGroups

`func (o *ServiceConfiguration) SetRequiredGroups(v []int32)`

SetRequiredGroups sets RequiredGroups field to given value.

### HasRequiredGroups

`func (o *ServiceConfiguration) HasRequiredGroups() bool`

HasRequiredGroups returns a boolean if a field has been set.

### SetRequiredGroupsNil

`func (o *ServiceConfiguration) SetRequiredGroupsNil(b bool)`

 SetRequiredGroupsNil sets the value for RequiredGroups to be an explicit nil

### UnsetRequiredGroups
`func (o *ServiceConfiguration) UnsetRequiredGroups()`

UnsetRequiredGroups ensures that no value is present for RequiredGroups, not even an explicit nil
### GetProperties

`func (o *ServiceConfiguration) GetProperties() []string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ServiceConfiguration) GetPropertiesOk() (*[]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ServiceConfiguration) SetProperties(v []string)`

SetProperties sets Properties field to given value.


### SetPropertiesNil

`func (o *ServiceConfiguration) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *ServiceConfiguration) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetShowPassword

`func (o *ServiceConfiguration) GetShowPassword() bool`

GetShowPassword returns the ShowPassword field if non-nil, zero value otherwise.

### GetShowPasswordOk

`func (o *ServiceConfiguration) GetShowPasswordOk() (*bool, bool)`

GetShowPasswordOk returns a tuple with the ShowPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowPassword

`func (o *ServiceConfiguration) SetShowPassword(v bool)`

SetShowPassword sets ShowPassword field to given value.

### HasShowPassword

`func (o *ServiceConfiguration) HasShowPassword() bool`

HasShowPassword returns a boolean if a field has been set.

### SetShowPasswordNil

`func (o *ServiceConfiguration) SetShowPasswordNil(b bool)`

 SetShowPasswordNil sets the value for ShowPassword to be an explicit nil

### UnsetShowPassword
`func (o *ServiceConfiguration) UnsetShowPassword()`

UnsetShowPassword ensures that no value is present for ShowPassword, not even an explicit nil
### GetActions

`func (o *ServiceConfiguration) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ServiceConfiguration) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ServiceConfiguration) SetActions(v []string)`

SetActions sets Actions field to given value.


### SetActionsNil

`func (o *ServiceConfiguration) SetActionsNil(b bool)`

 SetActionsNil sets the value for Actions to be an explicit nil

### UnsetActions
`func (o *ServiceConfiguration) UnsetActions()`

UnsetActions ensures that no value is present for Actions, not even an explicit nil
### GetURLs

`func (o *ServiceConfiguration) GetURLs() []ServiceConfigurationURL`

GetURLs returns the URLs field if non-nil, zero value otherwise.

### GetURLsOk

`func (o *ServiceConfiguration) GetURLsOk() (*[]ServiceConfigurationURL, bool)`

GetURLsOk returns a tuple with the URLs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetURLs

`func (o *ServiceConfiguration) SetURLs(v []ServiceConfigurationURL)`

SetURLs sets URLs field to given value.


### GetTextAccount

`func (o *ServiceConfiguration) GetTextAccount() string`

GetTextAccount returns the TextAccount field if non-nil, zero value otherwise.

### GetTextAccountOk

`func (o *ServiceConfiguration) GetTextAccountOk() (*string, bool)`

GetTextAccountOk returns a tuple with the TextAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextAccount

`func (o *ServiceConfiguration) SetTextAccount(v string)`

SetTextAccount sets TextAccount field to given value.


### SetTextAccountNil

`func (o *ServiceConfiguration) SetTextAccountNil(b bool)`

 SetTextAccountNil sets the value for TextAccount to be an explicit nil

### UnsetTextAccount
`func (o *ServiceConfiguration) UnsetTextAccount()`

UnsetTextAccount ensures that no value is present for TextAccount, not even an explicit nil
### GetTextTop

`func (o *ServiceConfiguration) GetTextTop() string`

GetTextTop returns the TextTop field if non-nil, zero value otherwise.

### GetTextTopOk

`func (o *ServiceConfiguration) GetTextTopOk() (*string, bool)`

GetTextTopOk returns a tuple with the TextTop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextTop

`func (o *ServiceConfiguration) SetTextTop(v string)`

SetTextTop sets TextTop field to given value.


### SetTextTopNil

`func (o *ServiceConfiguration) SetTextTopNil(b bool)`

 SetTextTopNil sets the value for TextTop to be an explicit nil

### UnsetTextTop
`func (o *ServiceConfiguration) UnsetTextTop()`

UnsetTextTop ensures that no value is present for TextTop, not even an explicit nil
### GetTextRegister

`func (o *ServiceConfiguration) GetTextRegister() string`

GetTextRegister returns the TextRegister field if non-nil, zero value otherwise.

### GetTextRegisterOk

`func (o *ServiceConfiguration) GetTextRegisterOk() (*string, bool)`

GetTextRegisterOk returns a tuple with the TextRegister field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextRegister

`func (o *ServiceConfiguration) SetTextRegister(v string)`

SetTextRegister sets TextRegister field to given value.


### SetTextRegisterNil

`func (o *ServiceConfiguration) SetTextRegisterNil(b bool)`

 SetTextRegisterNil sets the value for TextRegister to be an explicit nil

### UnsetTextRegister
`func (o *ServiceConfiguration) UnsetTextRegister()`

UnsetTextRegister ensures that no value is present for TextRegister, not even an explicit nil
### GetTextPending

`func (o *ServiceConfiguration) GetTextPending() string`

GetTextPending returns the TextPending field if non-nil, zero value otherwise.

### GetTextPendingOk

`func (o *ServiceConfiguration) GetTextPendingOk() (*string, bool)`

GetTextPendingOk returns a tuple with the TextPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextPending

`func (o *ServiceConfiguration) SetTextPending(v string)`

SetTextPending sets TextPending field to given value.


### SetTextPendingNil

`func (o *ServiceConfiguration) SetTextPendingNil(b bool)`

 SetTextPendingNil sets the value for TextPending to be an explicit nil

### UnsetTextPending
`func (o *ServiceConfiguration) UnsetTextPending()`

UnsetTextPending ensures that no value is present for TextPending, not even an explicit nil
### GetConfigurationData

`func (o *ServiceConfiguration) GetConfigurationData() string`

GetConfigurationData returns the ConfigurationData field if non-nil, zero value otherwise.

### GetConfigurationDataOk

`func (o *ServiceConfiguration) GetConfigurationDataOk() (*string, bool)`

GetConfigurationDataOk returns a tuple with the ConfigurationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationData

`func (o *ServiceConfiguration) SetConfigurationData(v string)`

SetConfigurationData sets ConfigurationData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


