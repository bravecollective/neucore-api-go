# Alliance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | EVE alliance ID. | 
**Name** | **NullableString** | EVE alliance name. | 
**Ticker** | **NullableString** | Alliance ticker. | 
**Groups** | Pointer to [**[]Group**](Group.md) | Groups for automatic assignment (API: not included by default). | [optional] 

## Methods

### NewAlliance

`func NewAlliance(id int64, name NullableString, ticker NullableString, ) *Alliance`

NewAlliance instantiates a new Alliance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllianceWithDefaults

`func NewAllianceWithDefaults() *Alliance`

NewAllianceWithDefaults instantiates a new Alliance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Alliance) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Alliance) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Alliance) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *Alliance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Alliance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Alliance) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *Alliance) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Alliance) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTicker

`func (o *Alliance) GetTicker() string`

GetTicker returns the Ticker field if non-nil, zero value otherwise.

### GetTickerOk

`func (o *Alliance) GetTickerOk() (*string, bool)`

GetTickerOk returns a tuple with the Ticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicker

`func (o *Alliance) SetTicker(v string)`

SetTicker sets Ticker field to given value.


### SetTickerNil

`func (o *Alliance) SetTickerNil(b bool)`

 SetTickerNil sets the value for Ticker to be an explicit nil

### UnsetTicker
`func (o *Alliance) UnsetTicker()`

UnsetTicker ensures that no value is present for Ticker, not even an explicit nil
### GetGroups

`func (o *Alliance) GetGroups() []Group`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *Alliance) GetGroupsOk() (*[]Group, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *Alliance) SetGroups(v []Group)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *Alliance) HasGroups() bool`

HasGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


