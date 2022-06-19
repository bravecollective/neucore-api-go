# Watchlist

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableInt32** |  | 
**Name** | **NullableString** |  | 
**LockWatchlistSettings** | Pointer to **bool** |  | [optional] 

## Methods

### NewWatchlist

`func NewWatchlist(id NullableInt32, name NullableString, ) *Watchlist`

NewWatchlist instantiates a new Watchlist object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWatchlistWithDefaults

`func NewWatchlistWithDefaults() *Watchlist`

NewWatchlistWithDefaults instantiates a new Watchlist object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Watchlist) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Watchlist) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Watchlist) SetId(v int32)`

SetId sets Id field to given value.


### SetIdNil

`func (o *Watchlist) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Watchlist) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *Watchlist) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Watchlist) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Watchlist) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *Watchlist) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Watchlist) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetLockWatchlistSettings

`func (o *Watchlist) GetLockWatchlistSettings() bool`

GetLockWatchlistSettings returns the LockWatchlistSettings field if non-nil, zero value otherwise.

### GetLockWatchlistSettingsOk

`func (o *Watchlist) GetLockWatchlistSettingsOk() (*bool, bool)`

GetLockWatchlistSettingsOk returns a tuple with the LockWatchlistSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockWatchlistSettings

`func (o *Watchlist) SetLockWatchlistSettings(v bool)`

SetLockWatchlistSettings sets LockWatchlistSettings field to given value.

### HasLockWatchlistSettings

`func (o *Watchlist) HasLockWatchlistSettings() bool`

HasLockWatchlistSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


