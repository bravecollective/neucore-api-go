# Group

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableInt32** | Group ID. | 
**Name** | **NullableString** | A unique group name (can be changed). | 
**Visibility** | Pointer to **string** |  | [optional] 

## Methods

### NewGroup

`func NewGroup(id NullableInt32, name NullableString, ) *Group`

NewGroup instantiates a new Group object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupWithDefaults

`func NewGroupWithDefaults() *Group`

NewGroupWithDefaults instantiates a new Group object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Group) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Group) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Group) SetId(v int32)`

SetId sets Id field to given value.


### SetIdNil

`func (o *Group) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Group) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *Group) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Group) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Group) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *Group) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Group) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetVisibility

`func (o *Group) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *Group) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *Group) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *Group) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


