/*
 * Neucore API
 *
 * Client library of Neucore API
 *
 * API version: 1.22.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package neucoreapi

import (
	"encoding/json"
	"fmt"
)

// Role the model 'Role'
type Role string

// List of Role
const (
	APP Role = "app"
	APP_GROUPS Role = "app-groups"
	APP_CHARS Role = "app-chars"
	APP_TRACKING Role = "app-tracking"
	APP_ESI Role = "app-esi"
	USER Role = "user"
	USER_ADMIN Role = "user-admin"
	USER_MANAGER Role = "user-manager"
	USER_CHARS Role = "user-chars"
	GROUP_ADMIN Role = "group-admin"
	GROUP_MANAGER Role = "group-manager"
	APP_ADMIN Role = "app-admin"
	APP_MANAGER Role = "app-manager"
	SERVICE_ADMIN Role = "service-admin"
	STATISTICS Role = "statistics"
	ESI Role = "esi"
	SETTINGS Role = "settings"
	TRACKING Role = "tracking"
	TRACKING_ADMIN Role = "tracking-admin"
	WATCHLIST Role = "watchlist"
	WATCHLIST_MANAGER Role = "watchlist-manager"
	WATCHLIST_ADMIN Role = "watchlist-admin"
)

var allowedRoleEnumValues = []Role{
	"app",
	"app-groups",
	"app-chars",
	"app-tracking",
	"app-esi",
	"user",
	"user-admin",
	"user-manager",
	"user-chars",
	"group-admin",
	"group-manager",
	"app-admin",
	"app-manager",
	"service-admin",
	"statistics",
	"esi",
	"settings",
	"tracking",
	"tracking-admin",
	"watchlist",
	"watchlist-manager",
	"watchlist-admin",
}

func (v *Role) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Role(value)
	for _, existing := range allowedRoleEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Role", value)
}

// NewRoleFromValue returns a pointer to a valid Role
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRoleFromValue(v string) (*Role, error) {
	ev := Role(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Role: valid values are %v", v, allowedRoleEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Role) IsValid() bool {
	for _, existing := range allowedRoleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Role value
func (v Role) Ptr() *Role {
	return &v
}

type NullableRole struct {
	value *Role
	isSet bool
}

func (v NullableRole) Get() *Role {
	return v.value
}

func (v *NullableRole) Set(val *Role) {
	v.value = val
	v.isSet = true
}

func (v NullableRole) IsSet() bool {
	return v.isSet
}

func (v *NullableRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRole(val *Role) *NullableRole {
	return &NullableRole{value: val, isSet: true}
}

func (v NullableRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

