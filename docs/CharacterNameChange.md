# CharacterNameChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OldName** | **string** |  | 
**ChangeDate** | **NullableTime** |  | 

## Methods

### NewCharacterNameChange

`func NewCharacterNameChange(oldName string, changeDate NullableTime, ) *CharacterNameChange`

NewCharacterNameChange instantiates a new CharacterNameChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCharacterNameChangeWithDefaults

`func NewCharacterNameChangeWithDefaults() *CharacterNameChange`

NewCharacterNameChangeWithDefaults instantiates a new CharacterNameChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOldName

`func (o *CharacterNameChange) GetOldName() string`

GetOldName returns the OldName field if non-nil, zero value otherwise.

### GetOldNameOk

`func (o *CharacterNameChange) GetOldNameOk() (*string, bool)`

GetOldNameOk returns a tuple with the OldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldName

`func (o *CharacterNameChange) SetOldName(v string)`

SetOldName sets OldName field to given value.


### GetChangeDate

`func (o *CharacterNameChange) GetChangeDate() time.Time`

GetChangeDate returns the ChangeDate field if non-nil, zero value otherwise.

### GetChangeDateOk

`func (o *CharacterNameChange) GetChangeDateOk() (*time.Time, bool)`

GetChangeDateOk returns a tuple with the ChangeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeDate

`func (o *CharacterNameChange) SetChangeDate(v time.Time)`

SetChangeDate sets ChangeDate field to given value.


### SetChangeDateNil

`func (o *CharacterNameChange) SetChangeDateNil(b bool)`

 SetChangeDateNil sets the value for ChangeDate to be an explicit nil

### UnsetChangeDate
`func (o *CharacterNameChange) UnsetChangeDate()`

UnsetChangeDate ensures that no value is present for ChangeDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


