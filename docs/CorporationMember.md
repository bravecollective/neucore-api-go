# CorporationMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Player** | Pointer to [**Player**](Player.md) |  | [optional] 
**Id** | **NullableInt64** | EVE Character ID. | 
**Name** | **NullableString** | EVE Character name. | 
**Location** | Pointer to [**EsiLocation**](EsiLocation.md) |  | [optional] 
**LogoffDate** | Pointer to **NullableTime** |  | [optional] 
**LogonDate** | Pointer to **NullableTime** |  | [optional] 
**ShipType** | Pointer to [**EsiType**](EsiType.md) |  | [optional] 
**StartDate** | Pointer to **NullableTime** |  | [optional] 
**Character** | Pointer to [**Character**](Character.md) |  | [optional] 
**MissingCharacterMailSentDate** | Pointer to **NullableTime** | Date and time of the last sent mail. | [optional] 
**MissingCharacterMailSentResult** | Pointer to **NullableString** | Result of the last sent mail (OK, Blocked, CSPA charge &gt; 0) | [optional] 
**MissingCharacterMailSentNumber** | Pointer to **int32** | Number of mails sent, is reset when the character is added. | [optional] 

## Methods

### NewCorporationMember

`func NewCorporationMember(id NullableInt64, name NullableString, ) *CorporationMember`

NewCorporationMember instantiates a new CorporationMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorporationMemberWithDefaults

`func NewCorporationMemberWithDefaults() *CorporationMember`

NewCorporationMemberWithDefaults instantiates a new CorporationMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlayer

`func (o *CorporationMember) GetPlayer() Player`

GetPlayer returns the Player field if non-nil, zero value otherwise.

### GetPlayerOk

`func (o *CorporationMember) GetPlayerOk() (*Player, bool)`

GetPlayerOk returns a tuple with the Player field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayer

`func (o *CorporationMember) SetPlayer(v Player)`

SetPlayer sets Player field to given value.

### HasPlayer

`func (o *CorporationMember) HasPlayer() bool`

HasPlayer returns a boolean if a field has been set.

### GetId

`func (o *CorporationMember) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CorporationMember) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CorporationMember) SetId(v int64)`

SetId sets Id field to given value.


### SetIdNil

`func (o *CorporationMember) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CorporationMember) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *CorporationMember) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CorporationMember) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CorporationMember) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *CorporationMember) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CorporationMember) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetLocation

`func (o *CorporationMember) GetLocation() EsiLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *CorporationMember) GetLocationOk() (*EsiLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *CorporationMember) SetLocation(v EsiLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *CorporationMember) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetLogoffDate

`func (o *CorporationMember) GetLogoffDate() time.Time`

GetLogoffDate returns the LogoffDate field if non-nil, zero value otherwise.

### GetLogoffDateOk

`func (o *CorporationMember) GetLogoffDateOk() (*time.Time, bool)`

GetLogoffDateOk returns a tuple with the LogoffDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoffDate

`func (o *CorporationMember) SetLogoffDate(v time.Time)`

SetLogoffDate sets LogoffDate field to given value.

### HasLogoffDate

`func (o *CorporationMember) HasLogoffDate() bool`

HasLogoffDate returns a boolean if a field has been set.

### SetLogoffDateNil

`func (o *CorporationMember) SetLogoffDateNil(b bool)`

 SetLogoffDateNil sets the value for LogoffDate to be an explicit nil

### UnsetLogoffDate
`func (o *CorporationMember) UnsetLogoffDate()`

UnsetLogoffDate ensures that no value is present for LogoffDate, not even an explicit nil
### GetLogonDate

`func (o *CorporationMember) GetLogonDate() time.Time`

GetLogonDate returns the LogonDate field if non-nil, zero value otherwise.

### GetLogonDateOk

`func (o *CorporationMember) GetLogonDateOk() (*time.Time, bool)`

GetLogonDateOk returns a tuple with the LogonDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogonDate

`func (o *CorporationMember) SetLogonDate(v time.Time)`

SetLogonDate sets LogonDate field to given value.

### HasLogonDate

`func (o *CorporationMember) HasLogonDate() bool`

HasLogonDate returns a boolean if a field has been set.

### SetLogonDateNil

`func (o *CorporationMember) SetLogonDateNil(b bool)`

 SetLogonDateNil sets the value for LogonDate to be an explicit nil

### UnsetLogonDate
`func (o *CorporationMember) UnsetLogonDate()`

UnsetLogonDate ensures that no value is present for LogonDate, not even an explicit nil
### GetShipType

`func (o *CorporationMember) GetShipType() EsiType`

GetShipType returns the ShipType field if non-nil, zero value otherwise.

### GetShipTypeOk

`func (o *CorporationMember) GetShipTypeOk() (*EsiType, bool)`

GetShipTypeOk returns a tuple with the ShipType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipType

`func (o *CorporationMember) SetShipType(v EsiType)`

SetShipType sets ShipType field to given value.

### HasShipType

`func (o *CorporationMember) HasShipType() bool`

HasShipType returns a boolean if a field has been set.

### GetStartDate

`func (o *CorporationMember) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CorporationMember) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CorporationMember) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CorporationMember) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *CorporationMember) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *CorporationMember) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetCharacter

`func (o *CorporationMember) GetCharacter() Character`

GetCharacter returns the Character field if non-nil, zero value otherwise.

### GetCharacterOk

`func (o *CorporationMember) GetCharacterOk() (*Character, bool)`

GetCharacterOk returns a tuple with the Character field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacter

`func (o *CorporationMember) SetCharacter(v Character)`

SetCharacter sets Character field to given value.

### HasCharacter

`func (o *CorporationMember) HasCharacter() bool`

HasCharacter returns a boolean if a field has been set.

### GetMissingCharacterMailSentDate

`func (o *CorporationMember) GetMissingCharacterMailSentDate() time.Time`

GetMissingCharacterMailSentDate returns the MissingCharacterMailSentDate field if non-nil, zero value otherwise.

### GetMissingCharacterMailSentDateOk

`func (o *CorporationMember) GetMissingCharacterMailSentDateOk() (*time.Time, bool)`

GetMissingCharacterMailSentDateOk returns a tuple with the MissingCharacterMailSentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingCharacterMailSentDate

`func (o *CorporationMember) SetMissingCharacterMailSentDate(v time.Time)`

SetMissingCharacterMailSentDate sets MissingCharacterMailSentDate field to given value.

### HasMissingCharacterMailSentDate

`func (o *CorporationMember) HasMissingCharacterMailSentDate() bool`

HasMissingCharacterMailSentDate returns a boolean if a field has been set.

### SetMissingCharacterMailSentDateNil

`func (o *CorporationMember) SetMissingCharacterMailSentDateNil(b bool)`

 SetMissingCharacterMailSentDateNil sets the value for MissingCharacterMailSentDate to be an explicit nil

### UnsetMissingCharacterMailSentDate
`func (o *CorporationMember) UnsetMissingCharacterMailSentDate()`

UnsetMissingCharacterMailSentDate ensures that no value is present for MissingCharacterMailSentDate, not even an explicit nil
### GetMissingCharacterMailSentResult

`func (o *CorporationMember) GetMissingCharacterMailSentResult() string`

GetMissingCharacterMailSentResult returns the MissingCharacterMailSentResult field if non-nil, zero value otherwise.

### GetMissingCharacterMailSentResultOk

`func (o *CorporationMember) GetMissingCharacterMailSentResultOk() (*string, bool)`

GetMissingCharacterMailSentResultOk returns a tuple with the MissingCharacterMailSentResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingCharacterMailSentResult

`func (o *CorporationMember) SetMissingCharacterMailSentResult(v string)`

SetMissingCharacterMailSentResult sets MissingCharacterMailSentResult field to given value.

### HasMissingCharacterMailSentResult

`func (o *CorporationMember) HasMissingCharacterMailSentResult() bool`

HasMissingCharacterMailSentResult returns a boolean if a field has been set.

### SetMissingCharacterMailSentResultNil

`func (o *CorporationMember) SetMissingCharacterMailSentResultNil(b bool)`

 SetMissingCharacterMailSentResultNil sets the value for MissingCharacterMailSentResult to be an explicit nil

### UnsetMissingCharacterMailSentResult
`func (o *CorporationMember) UnsetMissingCharacterMailSentResult()`

UnsetMissingCharacterMailSentResult ensures that no value is present for MissingCharacterMailSentResult, not even an explicit nil
### GetMissingCharacterMailSentNumber

`func (o *CorporationMember) GetMissingCharacterMailSentNumber() int32`

GetMissingCharacterMailSentNumber returns the MissingCharacterMailSentNumber field if non-nil, zero value otherwise.

### GetMissingCharacterMailSentNumberOk

`func (o *CorporationMember) GetMissingCharacterMailSentNumberOk() (*int32, bool)`

GetMissingCharacterMailSentNumberOk returns a tuple with the MissingCharacterMailSentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingCharacterMailSentNumber

`func (o *CorporationMember) SetMissingCharacterMailSentNumber(v int32)`

SetMissingCharacterMailSentNumber sets MissingCharacterMailSentNumber field to given value.

### HasMissingCharacterMailSentNumber

`func (o *CorporationMember) HasMissingCharacterMailSentNumber() bool`

HasMissingCharacterMailSentNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


