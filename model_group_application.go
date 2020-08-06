/*
 * Neucore API
 *
 * Client library of Neucore API
 *
 * API version: 1.14.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package neucoreapi
import (
	"time"
)
// GroupApplication The player property contains only id and name.
type GroupApplication struct {
	Id int32 `json:"id"`
	Player Player `json:"player"`
	Group Group `json:"group"`
	Created *time.Time `json:"created"`
	// Group application status.
	Status string `json:"status,omitempty"`
}
