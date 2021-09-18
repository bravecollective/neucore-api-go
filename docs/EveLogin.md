# EveLogin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableInt32** |  | 
**Name** | **string** | Names starting with &#39;core.&#39; are reserved for internal use. | 
**Description** | **string** |  | 
**EsiScopes** | **string** |  | 
**EveRoles** | **[]string** | Maximum length of all roles separated by comma: 1024. | 

## Methods

### NewEveLogin

`func NewEveLogin(id NullableInt32, name string, description string, esiScopes string, eveRoles []string, ) *EveLogin`

NewEveLogin instantiates a new EveLogin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEveLoginWithDefaults

`func NewEveLoginWithDefaults() *EveLogin`

NewEveLoginWithDefaults instantiates a new EveLogin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EveLogin) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EveLogin) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EveLogin) SetId(v int32)`

SetId sets Id field to given value.


### SetIdNil

`func (o *EveLogin) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *EveLogin) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *EveLogin) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EveLogin) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EveLogin) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *EveLogin) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EveLogin) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EveLogin) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEsiScopes

`func (o *EveLogin) GetEsiScopes() string`

GetEsiScopes returns the EsiScopes field if non-nil, zero value otherwise.

### GetEsiScopesOk

`func (o *EveLogin) GetEsiScopesOk() (*string, bool)`

GetEsiScopesOk returns a tuple with the EsiScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsiScopes

`func (o *EveLogin) SetEsiScopes(v string)`

SetEsiScopes sets EsiScopes field to given value.


### GetEveRoles

`func (o *EveLogin) GetEveRoles() []string`

GetEveRoles returns the EveRoles field if non-nil, zero value otherwise.

### GetEveRolesOk

`func (o *EveLogin) GetEveRolesOk() (*[]string, bool)`

GetEveRolesOk returns a tuple with the EveRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEveRoles

`func (o *EveLogin) SetEveRoles(v []string)`

SetEveRoles sets EveRoles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


