# ServiceConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhpClass** | Pointer to **string** |  | [optional] 
**Psr4Prefix** | Pointer to **string** |  | [optional] 
**Psr4Path** | Pointer to **string** |  | [optional] 
**RequiredGroups** | Pointer to **[]int32** |  | [optional] 
**Properties** | **[]string** |  | 
**ShowPassword** | Pointer to **bool** |  | [optional] 
**Actions** | **[]string** |  | 
**URLs** | [**[]ServiceConfigurationURL**](ServiceConfigurationURL.md) |  | 
**TextAccount** | **string** |  | 
**TextTop** | **string** |  | 
**TextRegister** | **string** |  | 
**TextPending** | **string** |  | 

## Methods

### NewServiceConfiguration

`func NewServiceConfiguration(properties []string, actions []string, uRLs []ServiceConfigurationURL, textAccount string, textTop string, textRegister string, textPending string, ) *ServiceConfiguration`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


