# GroupApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Player** | [**Player**](Player.md) |  | 
**Group** | [**Group**](Group.md) |  | 
**Created** | **NullableTime** |  | 
**Status** | Pointer to **string** | Group application status. | [optional] 

## Methods

### NewGroupApplication

`func NewGroupApplication(id int32, player Player, group Group, created NullableTime, ) *GroupApplication`

NewGroupApplication instantiates a new GroupApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupApplicationWithDefaults

`func NewGroupApplicationWithDefaults() *GroupApplication`

NewGroupApplicationWithDefaults instantiates a new GroupApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroupApplication) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupApplication) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupApplication) SetId(v int32)`

SetId sets Id field to given value.


### GetPlayer

`func (o *GroupApplication) GetPlayer() Player`

GetPlayer returns the Player field if non-nil, zero value otherwise.

### GetPlayerOk

`func (o *GroupApplication) GetPlayerOk() (*Player, bool)`

GetPlayerOk returns a tuple with the Player field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayer

`func (o *GroupApplication) SetPlayer(v Player)`

SetPlayer sets Player field to given value.


### GetGroup

`func (o *GroupApplication) GetGroup() Group`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *GroupApplication) GetGroupOk() (*Group, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *GroupApplication) SetGroup(v Group)`

SetGroup sets Group field to given value.


### GetCreated

`func (o *GroupApplication) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GroupApplication) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GroupApplication) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *GroupApplication) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *GroupApplication) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetStatus

`func (o *GroupApplication) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GroupApplication) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GroupApplication) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GroupApplication) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


