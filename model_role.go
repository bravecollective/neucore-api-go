/*
 * Neucore API
 *
 * Client library of Neucore API
 *
 * API version: 1.14.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package neucoreapi
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
	ESI Role = "esi"
	SETTINGS Role = "settings"
	TRACKING Role = "tracking"
	TRACKING_ADMIN Role = "tracking-admin"
	WATCHLIST Role = "watchlist"
	WATCHLIST_MANAGER Role = "watchlist-manager"
	WATCHLIST_ADMIN Role = "watchlist-admin"
)
