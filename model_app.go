/*
Neucore API

Client library of Neucore API

API version: 2.4.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package neucoreapi

import (
	"encoding/json"
)

// checks if the App type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &App{}

// App struct for App
type App struct {
	// App ID
	Id NullableInt32 `json:"id"`
	// App name
	Name NullableString `json:"name"`
	// Roles for authorization.
	Roles []Role `json:"roles,omitempty"`
	// Groups the app can see.
	Groups []Group `json:"groups,omitempty"`
	EveLogins []EveLogin `json:"eveLogins,omitempty"`
}

// NewApp instantiates a new App object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApp(id NullableInt32, name NullableString) *App {
	this := App{}
	this.Id = id
	this.Name = name
	return &this
}

// NewAppWithDefaults instantiates a new App object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppWithDefaults() *App {
	this := App{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *App) GetId() int32 {
	if o == nil || o.Id.Get() == nil {
		var ret int32
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *App) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *App) SetId(v int32) {
	o.Id.Set(&v)
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *App) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *App) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *App) SetName(v string) {
	o.Name.Set(&v)
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *App) GetRoles() []Role {
	if o == nil || IsNil(o.Roles) {
		var ret []Role
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetRolesOk() ([]Role, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *App) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []Role and assigns it to the Roles field.
func (o *App) SetRoles(v []Role) {
	o.Roles = v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *App) GetGroups() []Group {
	if o == nil || IsNil(o.Groups) {
		var ret []Group
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetGroupsOk() ([]Group, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *App) HasGroups() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []Group and assigns it to the Groups field.
func (o *App) SetGroups(v []Group) {
	o.Groups = v
}

// GetEveLogins returns the EveLogins field value if set, zero value otherwise.
func (o *App) GetEveLogins() []EveLogin {
	if o == nil || IsNil(o.EveLogins) {
		var ret []EveLogin
		return ret
	}
	return o.EveLogins
}

// GetEveLoginsOk returns a tuple with the EveLogins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetEveLoginsOk() ([]EveLogin, bool) {
	if o == nil || IsNil(o.EveLogins) {
		return nil, false
	}
	return o.EveLogins, true
}

// HasEveLogins returns a boolean if a field has been set.
func (o *App) HasEveLogins() bool {
	if o != nil && !IsNil(o.EveLogins) {
		return true
	}

	return false
}

// SetEveLogins gets a reference to the given []EveLogin and assigns it to the EveLogins field.
func (o *App) SetEveLogins(v []EveLogin) {
	o.EveLogins = v
}

func (o App) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o App) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id.Get()
	toSerialize["name"] = o.Name.Get()
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	if !IsNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	if !IsNil(o.EveLogins) {
		toSerialize["eveLogins"] = o.EveLogins
	}
	return toSerialize, nil
}

type NullableApp struct {
	value *App
	isSet bool
}

func (v NullableApp) Get() *App {
	return v.value
}

func (v *NullableApp) Set(val *App) {
	v.value = val
	v.isSet = true
}

func (v NullableApp) IsSet() bool {
	return v.isSet
}

func (v *NullableApp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApp(val *App) *NullableApp {
	return &NullableApp{value: val, isSet: true}
}

func (v NullableApp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


