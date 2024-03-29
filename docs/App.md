# App

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableInt32** | App ID | 
**Name** | **NullableString** | App name | 
**Roles** | Pointer to [**[]Role**](Role.md) | Roles for authorization. | [optional] 
**Groups** | Pointer to [**[]Group**](Group.md) | Groups the app can see. | [optional] 
**EveLogins** | Pointer to [**[]EveLogin**](EveLogin.md) |  | [optional] 

## Methods

### NewApp

`func NewApp(id NullableInt32, name NullableString, ) *App`

NewApp instantiates a new App object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppWithDefaults

`func NewAppWithDefaults() *App`

NewAppWithDefaults instantiates a new App object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *App) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *App) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *App) SetId(v int32)`

SetId sets Id field to given value.


### SetIdNil

`func (o *App) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *App) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *App) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *App) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *App) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *App) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *App) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRoles

`func (o *App) GetRoles() []Role`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *App) GetRolesOk() (*[]Role, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *App) SetRoles(v []Role)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *App) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetGroups

`func (o *App) GetGroups() []Group`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *App) GetGroupsOk() (*[]Group, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *App) SetGroups(v []Group)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *App) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetEveLogins

`func (o *App) GetEveLogins() []EveLogin`

GetEveLogins returns the EveLogins field if non-nil, zero value otherwise.

### GetEveLoginsOk

`func (o *App) GetEveLoginsOk() (*[]EveLogin, bool)`

GetEveLoginsOk returns a tuple with the EveLogins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEveLogins

`func (o *App) SetEveLogins(v []EveLogin)`

SetEveLogins sets EveLogins field to given value.

### HasEveLogins

`func (o *App) HasEveLogins() bool`

HasEveLogins returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


