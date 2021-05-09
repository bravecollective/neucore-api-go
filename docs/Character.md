# Character

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | EVE character ID. | 
**Name** | **string** | EVE character name. | 
**Main** | Pointer to **bool** |  | [optional] 
**ValidToken** | Pointer to **NullableBool** | Shows if character&#39;s refresh token is valid or not.  This is null if there is no refresh token (EVE SSOv1 only) or a valid token but without scopes (SSOv2). | [optional] 
**ValidTokenTime** | Pointer to **NullableTime** | Date and time when that valid token property was last changed. | [optional] 
**Created** | Pointer to **NullableTime** |  | [optional] 
**LastUpdate** | Pointer to **NullableTime** | Last ESI update. | [optional] 
**Corporation** | Pointer to [**Corporation**](Corporation.md) |  | [optional] 
**CharacterNameChanges** | Pointer to [**[]CharacterNameChange**](CharacterNameChange.md) | List of previous character names (API: not included by default). | [optional] 

## Methods

### NewCharacter

`func NewCharacter(id int64, name string, ) *Character`

NewCharacter instantiates a new Character object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCharacterWithDefaults

`func NewCharacterWithDefaults() *Character`

NewCharacterWithDefaults instantiates a new Character object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Character) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Character) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Character) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *Character) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Character) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Character) SetName(v string)`

SetName sets Name field to given value.


### GetMain

`func (o *Character) GetMain() bool`

GetMain returns the Main field if non-nil, zero value otherwise.

### GetMainOk

`func (o *Character) GetMainOk() (*bool, bool)`

GetMainOk returns a tuple with the Main field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMain

`func (o *Character) SetMain(v bool)`

SetMain sets Main field to given value.

### HasMain

`func (o *Character) HasMain() bool`

HasMain returns a boolean if a field has been set.

### GetValidToken

`func (o *Character) GetValidToken() bool`

GetValidToken returns the ValidToken field if non-nil, zero value otherwise.

### GetValidTokenOk

`func (o *Character) GetValidTokenOk() (*bool, bool)`

GetValidTokenOk returns a tuple with the ValidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidToken

`func (o *Character) SetValidToken(v bool)`

SetValidToken sets ValidToken field to given value.

### HasValidToken

`func (o *Character) HasValidToken() bool`

HasValidToken returns a boolean if a field has been set.

### SetValidTokenNil

`func (o *Character) SetValidTokenNil(b bool)`

 SetValidTokenNil sets the value for ValidToken to be an explicit nil

### UnsetValidToken
`func (o *Character) UnsetValidToken()`

UnsetValidToken ensures that no value is present for ValidToken, not even an explicit nil
### GetValidTokenTime

`func (o *Character) GetValidTokenTime() time.Time`

GetValidTokenTime returns the ValidTokenTime field if non-nil, zero value otherwise.

### GetValidTokenTimeOk

`func (o *Character) GetValidTokenTimeOk() (*time.Time, bool)`

GetValidTokenTimeOk returns a tuple with the ValidTokenTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidTokenTime

`func (o *Character) SetValidTokenTime(v time.Time)`

SetValidTokenTime sets ValidTokenTime field to given value.

### HasValidTokenTime

`func (o *Character) HasValidTokenTime() bool`

HasValidTokenTime returns a boolean if a field has been set.

### SetValidTokenTimeNil

`func (o *Character) SetValidTokenTimeNil(b bool)`

 SetValidTokenTimeNil sets the value for ValidTokenTime to be an explicit nil

### UnsetValidTokenTime
`func (o *Character) UnsetValidTokenTime()`

UnsetValidTokenTime ensures that no value is present for ValidTokenTime, not even an explicit nil
### GetCreated

`func (o *Character) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Character) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Character) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Character) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *Character) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *Character) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdate

`func (o *Character) GetLastUpdate() time.Time`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *Character) GetLastUpdateOk() (*time.Time, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *Character) SetLastUpdate(v time.Time)`

SetLastUpdate sets LastUpdate field to given value.

### HasLastUpdate

`func (o *Character) HasLastUpdate() bool`

HasLastUpdate returns a boolean if a field has been set.

### SetLastUpdateNil

`func (o *Character) SetLastUpdateNil(b bool)`

 SetLastUpdateNil sets the value for LastUpdate to be an explicit nil

### UnsetLastUpdate
`func (o *Character) UnsetLastUpdate()`

UnsetLastUpdate ensures that no value is present for LastUpdate, not even an explicit nil
### GetCorporation

`func (o *Character) GetCorporation() Corporation`

GetCorporation returns the Corporation field if non-nil, zero value otherwise.

### GetCorporationOk

`func (o *Character) GetCorporationOk() (*Corporation, bool)`

GetCorporationOk returns a tuple with the Corporation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorporation

`func (o *Character) SetCorporation(v Corporation)`

SetCorporation sets Corporation field to given value.

### HasCorporation

`func (o *Character) HasCorporation() bool`

HasCorporation returns a boolean if a field has been set.

### GetCharacterNameChanges

`func (o *Character) GetCharacterNameChanges() []CharacterNameChange`

GetCharacterNameChanges returns the CharacterNameChanges field if non-nil, zero value otherwise.

### GetCharacterNameChangesOk

`func (o *Character) GetCharacterNameChangesOk() (*[]CharacterNameChange, bool)`

GetCharacterNameChangesOk returns a tuple with the CharacterNameChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacterNameChanges

`func (o *Character) SetCharacterNameChanges(v []CharacterNameChange)`

SetCharacterNameChanges sets CharacterNameChanges field to given value.

### HasCharacterNameChanges

`func (o *Character) HasCharacterNameChanges() bool`

HasCharacterNameChanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


