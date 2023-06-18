# CharacterGroups

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Character** | [**Character**](Character.md) |  | 
**Groups** | [**[]Group**](Group.md) |  | 
**Deactivated** | **string** | Groups deactivation status. | 

## Methods

### NewCharacterGroups

`func NewCharacterGroups(character Character, groups []Group, deactivated string, ) *CharacterGroups`

NewCharacterGroups instantiates a new CharacterGroups object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCharacterGroupsWithDefaults

`func NewCharacterGroupsWithDefaults() *CharacterGroups`

NewCharacterGroupsWithDefaults instantiates a new CharacterGroups object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCharacter

`func (o *CharacterGroups) GetCharacter() Character`

GetCharacter returns the Character field if non-nil, zero value otherwise.

### GetCharacterOk

`func (o *CharacterGroups) GetCharacterOk() (*Character, bool)`

GetCharacterOk returns a tuple with the Character field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacter

`func (o *CharacterGroups) SetCharacter(v Character)`

SetCharacter sets Character field to given value.


### GetGroups

`func (o *CharacterGroups) GetGroups() []Group`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *CharacterGroups) GetGroupsOk() (*[]Group, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *CharacterGroups) SetGroups(v []Group)`

SetGroups sets Groups field to given value.


### GetDeactivated

`func (o *CharacterGroups) GetDeactivated() string`

GetDeactivated returns the Deactivated field if non-nil, zero value otherwise.

### GetDeactivatedOk

`func (o *CharacterGroups) GetDeactivatedOk() (*string, bool)`

GetDeactivatedOk returns a tuple with the Deactivated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeactivated

`func (o *CharacterGroups) SetDeactivated(v string)`

SetDeactivated sets Deactivated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


