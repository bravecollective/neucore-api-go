# Corporation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableInt64** | EVE corporation ID. | 
**Name** | **NullableString** | EVE corporation name. | 
**Ticker** | **NullableString** | Corporation ticker. | 
**Alliance** | Pointer to [**Alliance**](Alliance.md) |  | [optional] 
**Groups** | Pointer to [**[]Group**](Group.md) | Groups for automatic assignment (API: not included by default). | [optional] 
**TrackingLastUpdate** | Pointer to **NullableTime** | Last update of corporation member tracking data (API: not included by default). | [optional] 
**AutoAllowlist** | Pointer to **bool** | True if this corporation was automatically placed on the allowlist of a watchlist (API: not included by default). | [optional] 

## Methods

### NewCorporation

`func NewCorporation(id NullableInt64, name NullableString, ticker NullableString, ) *Corporation`

NewCorporation instantiates a new Corporation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorporationWithDefaults

`func NewCorporationWithDefaults() *Corporation`

NewCorporationWithDefaults instantiates a new Corporation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Corporation) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Corporation) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Corporation) SetId(v int64)`

SetId sets Id field to given value.


### SetIdNil

`func (o *Corporation) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Corporation) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *Corporation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Corporation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Corporation) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *Corporation) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Corporation) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTicker

`func (o *Corporation) GetTicker() string`

GetTicker returns the Ticker field if non-nil, zero value otherwise.

### GetTickerOk

`func (o *Corporation) GetTickerOk() (*string, bool)`

GetTickerOk returns a tuple with the Ticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicker

`func (o *Corporation) SetTicker(v string)`

SetTicker sets Ticker field to given value.


### SetTickerNil

`func (o *Corporation) SetTickerNil(b bool)`

 SetTickerNil sets the value for Ticker to be an explicit nil

### UnsetTicker
`func (o *Corporation) UnsetTicker()`

UnsetTicker ensures that no value is present for Ticker, not even an explicit nil
### GetAlliance

`func (o *Corporation) GetAlliance() Alliance`

GetAlliance returns the Alliance field if non-nil, zero value otherwise.

### GetAllianceOk

`func (o *Corporation) GetAllianceOk() (*Alliance, bool)`

GetAllianceOk returns a tuple with the Alliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlliance

`func (o *Corporation) SetAlliance(v Alliance)`

SetAlliance sets Alliance field to given value.

### HasAlliance

`func (o *Corporation) HasAlliance() bool`

HasAlliance returns a boolean if a field has been set.

### GetGroups

`func (o *Corporation) GetGroups() []Group`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *Corporation) GetGroupsOk() (*[]Group, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *Corporation) SetGroups(v []Group)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *Corporation) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetTrackingLastUpdate

`func (o *Corporation) GetTrackingLastUpdate() time.Time`

GetTrackingLastUpdate returns the TrackingLastUpdate field if non-nil, zero value otherwise.

### GetTrackingLastUpdateOk

`func (o *Corporation) GetTrackingLastUpdateOk() (*time.Time, bool)`

GetTrackingLastUpdateOk returns a tuple with the TrackingLastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingLastUpdate

`func (o *Corporation) SetTrackingLastUpdate(v time.Time)`

SetTrackingLastUpdate sets TrackingLastUpdate field to given value.

### HasTrackingLastUpdate

`func (o *Corporation) HasTrackingLastUpdate() bool`

HasTrackingLastUpdate returns a boolean if a field has been set.

### SetTrackingLastUpdateNil

`func (o *Corporation) SetTrackingLastUpdateNil(b bool)`

 SetTrackingLastUpdateNil sets the value for TrackingLastUpdate to be an explicit nil

### UnsetTrackingLastUpdate
`func (o *Corporation) UnsetTrackingLastUpdate()`

UnsetTrackingLastUpdate ensures that no value is present for TrackingLastUpdate, not even an explicit nil
### GetAutoAllowlist

`func (o *Corporation) GetAutoAllowlist() bool`

GetAutoAllowlist returns the AutoAllowlist field if non-nil, zero value otherwise.

### GetAutoAllowlistOk

`func (o *Corporation) GetAutoAllowlistOk() (*bool, bool)`

GetAutoAllowlistOk returns a tuple with the AutoAllowlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoAllowlist

`func (o *Corporation) SetAutoAllowlist(v bool)`

SetAutoAllowlist sets AutoAllowlist field to given value.

### HasAutoAllowlist

`func (o *Corporation) HasAutoAllowlist() bool`

HasAutoAllowlist returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


