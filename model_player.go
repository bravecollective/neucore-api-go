/*
 * Neucore API
 *
 * Client library of Neucore API
 *
 * API version: 1.21.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package neucoreapi

import (
	"encoding/json"
)

// Player struct for Player
type Player struct {
	// External service accounts (API: not included by default)
	ServiceAccounts *[]ServiceAccount `json:"serviceAccounts,omitempty"`
	Id int32 `json:"id"`
	// A name for the player.  This is the EVE character name of the current main character or of the last main character if there is currently none.
	Name string `json:"name"`
	// Player account status.
	Status *string `json:"status,omitempty"`
	// Roles for authorization.
	Roles *[]Role `json:"roles,omitempty"`
	Characters *[]Character `json:"characters,omitempty"`
	// Group membership.
	Groups *[]Group `json:"groups,omitempty"`
	// Manager of groups.
	ManagerGroups *[]Group `json:"managerGroups,omitempty"`
	// Manager of apps.
	ManagerApps *[]App `json:"managerApps,omitempty"`
	// Characters that were removed from a player (API: not included by default).
	RemovedCharacters *[]RemovedCharacter `json:"removedCharacters,omitempty"`
	// Characters that were moved from another player account to this account (API: not included by default).
	IncomingCharacters *[]RemovedCharacter `json:"incomingCharacters,omitempty"`
}

// NewPlayer instantiates a new Player object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlayer(id int32, name string) *Player {
	this := Player{}
	this.Id = id
	this.Name = name
	return &this
}

// NewPlayerWithDefaults instantiates a new Player object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlayerWithDefaults() *Player {
	this := Player{}
	return &this
}

// GetServiceAccounts returns the ServiceAccounts field value if set, zero value otherwise.
func (o *Player) GetServiceAccounts() []ServiceAccount {
	if o == nil || o.ServiceAccounts == nil {
		var ret []ServiceAccount
		return ret
	}
	return *o.ServiceAccounts
}

// GetServiceAccountsOk returns a tuple with the ServiceAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Player) GetServiceAccountsOk() (*[]ServiceAccount, bool) {
	if o == nil || o.ServiceAccounts == nil {
		return nil, false
	}
	return o.ServiceAccounts, true
}

// HasServiceAccounts returns a boolean if a field has been set.
func (o *Player) HasServiceAccounts() bool {
	if o != nil && o.ServiceAccounts != nil {
		return true
	}

	return false
}

// SetServiceAccounts gets a reference to the given []ServiceAccount and assigns it to the ServiceAccounts field.
func (o *Player) SetServiceAccounts(v []ServiceAccount) {
	o.ServiceAccounts = &v
}

// GetId returns the Id field value
func (o *Player) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Player) GetIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Player) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Player) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Player) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Player) SetName(v string) {
	o.Name = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Player) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Player) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Player) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Player) SetStatus(v string) {
	o.Status = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *Player) GetRoles() []Role {
	if o == nil || o.Roles == nil {
		var ret []Role
		return ret
	}
	return *o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Player) GetRolesOk() (*[]Role, bool) {
	if o == nil || o.Roles == nil {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *Player) HasRoles() bool {
	if o != nil && o.Roles != nil {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []Role and assigns it to the Roles field.
func (o *Player) SetRoles(v []Role) {
	o.Roles = &v
}

// GetCharacters returns the Characters field value if set, zero value otherwise.
func (o *Player) GetCharacters() []Character {
	if o == nil || o.Characters == nil {
		var ret []Character
		return ret
	}
	return *o.Characters
}

// GetCharactersOk returns a tuple with the Characters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Player) GetCharactersOk() (*[]Character, bool) {
	if o == nil || o.Characters == nil {
		return nil, false
	}
	return o.Characters, true
}

// HasCharacters returns a boolean if a field has been set.
func (o *Player) HasCharacters() bool {
	if o != nil && o.Characters != nil {
		return true
	}

	return false
}

// SetCharacters gets a reference to the given []Character and assigns it to the Characters field.
func (o *Player) SetCharacters(v []Character) {
	o.Characters = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *Player) GetGroups() []Group {
	if o == nil || o.Groups == nil {
		var ret []Group
		return ret
	}
	return *o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Player) GetGroupsOk() (*[]Group, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *Player) HasGroups() bool {
	if o != nil && o.Groups != nil {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []Group and assigns it to the Groups field.
func (o *Player) SetGroups(v []Group) {
	o.Groups = &v
}

// GetManagerGroups returns the ManagerGroups field value if set, zero value otherwise.
func (o *Player) GetManagerGroups() []Group {
	if o == nil || o.ManagerGroups == nil {
		var ret []Group
		return ret
	}
	return *o.ManagerGroups
}

// GetManagerGroupsOk returns a tuple with the ManagerGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Player) GetManagerGroupsOk() (*[]Group, bool) {
	if o == nil || o.ManagerGroups == nil {
		return nil, false
	}
	return o.ManagerGroups, true
}

// HasManagerGroups returns a boolean if a field has been set.
func (o *Player) HasManagerGroups() bool {
	if o != nil && o.ManagerGroups != nil {
		return true
	}

	return false
}

// SetManagerGroups gets a reference to the given []Group and assigns it to the ManagerGroups field.
func (o *Player) SetManagerGroups(v []Group) {
	o.ManagerGroups = &v
}

// GetManagerApps returns the ManagerApps field value if set, zero value otherwise.
func (o *Player) GetManagerApps() []App {
	if o == nil || o.ManagerApps == nil {
		var ret []App
		return ret
	}
	return *o.ManagerApps
}

// GetManagerAppsOk returns a tuple with the ManagerApps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Player) GetManagerAppsOk() (*[]App, bool) {
	if o == nil || o.ManagerApps == nil {
		return nil, false
	}
	return o.ManagerApps, true
}

// HasManagerApps returns a boolean if a field has been set.
func (o *Player) HasManagerApps() bool {
	if o != nil && o.ManagerApps != nil {
		return true
	}

	return false
}

// SetManagerApps gets a reference to the given []App and assigns it to the ManagerApps field.
func (o *Player) SetManagerApps(v []App) {
	o.ManagerApps = &v
}

// GetRemovedCharacters returns the RemovedCharacters field value if set, zero value otherwise.
func (o *Player) GetRemovedCharacters() []RemovedCharacter {
	if o == nil || o.RemovedCharacters == nil {
		var ret []RemovedCharacter
		return ret
	}
	return *o.RemovedCharacters
}

// GetRemovedCharactersOk returns a tuple with the RemovedCharacters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Player) GetRemovedCharactersOk() (*[]RemovedCharacter, bool) {
	if o == nil || o.RemovedCharacters == nil {
		return nil, false
	}
	return o.RemovedCharacters, true
}

// HasRemovedCharacters returns a boolean if a field has been set.
func (o *Player) HasRemovedCharacters() bool {
	if o != nil && o.RemovedCharacters != nil {
		return true
	}

	return false
}

// SetRemovedCharacters gets a reference to the given []RemovedCharacter and assigns it to the RemovedCharacters field.
func (o *Player) SetRemovedCharacters(v []RemovedCharacter) {
	o.RemovedCharacters = &v
}

// GetIncomingCharacters returns the IncomingCharacters field value if set, zero value otherwise.
func (o *Player) GetIncomingCharacters() []RemovedCharacter {
	if o == nil || o.IncomingCharacters == nil {
		var ret []RemovedCharacter
		return ret
	}
	return *o.IncomingCharacters
}

// GetIncomingCharactersOk returns a tuple with the IncomingCharacters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Player) GetIncomingCharactersOk() (*[]RemovedCharacter, bool) {
	if o == nil || o.IncomingCharacters == nil {
		return nil, false
	}
	return o.IncomingCharacters, true
}

// HasIncomingCharacters returns a boolean if a field has been set.
func (o *Player) HasIncomingCharacters() bool {
	if o != nil && o.IncomingCharacters != nil {
		return true
	}

	return false
}

// SetIncomingCharacters gets a reference to the given []RemovedCharacter and assigns it to the IncomingCharacters field.
func (o *Player) SetIncomingCharacters(v []RemovedCharacter) {
	o.IncomingCharacters = &v
}

func (o Player) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ServiceAccounts != nil {
		toSerialize["serviceAccounts"] = o.ServiceAccounts
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Roles != nil {
		toSerialize["roles"] = o.Roles
	}
	if o.Characters != nil {
		toSerialize["characters"] = o.Characters
	}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	if o.ManagerGroups != nil {
		toSerialize["managerGroups"] = o.ManagerGroups
	}
	if o.ManagerApps != nil {
		toSerialize["managerApps"] = o.ManagerApps
	}
	if o.RemovedCharacters != nil {
		toSerialize["removedCharacters"] = o.RemovedCharacters
	}
	if o.IncomingCharacters != nil {
		toSerialize["incomingCharacters"] = o.IncomingCharacters
	}
	return json.Marshal(toSerialize)
}

type NullablePlayer struct {
	value *Player
	isSet bool
}

func (v NullablePlayer) Get() *Player {
	return v.value
}

func (v *NullablePlayer) Set(val *Player) {
	v.value = val
	v.isSet = true
}

func (v NullablePlayer) IsSet() bool {
	return v.isSet
}

func (v *NullablePlayer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlayer(val *Player) *NullablePlayer {
	return &NullablePlayer{value: val, isSet: true}
}

func (v NullablePlayer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlayer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


