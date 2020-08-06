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
// CorporationMember The player property contains only id and name, character does not contain corporation.
type CorporationMember struct {
	Player Player `json:"player,omitempty"`
	// EVE Character ID.
	Id int64 `json:"id"`
	// EVE Character name.
	Name *string `json:"name"`
	Location EsiLocation `json:"location,omitempty"`
	LogoffDate *time.Time `json:"logoffDate,omitempty"`
	LogonDate *time.Time `json:"logonDate,omitempty"`
	ShipType EsiType `json:"shipType,omitempty"`
	StartDate *time.Time `json:"startDate,omitempty"`
	Character Character `json:"character,omitempty"`
	// Date and time of the last sent mail.
	MissingCharacterMailSentDate *time.Time `json:"missingCharacterMailSentDate,omitempty"`
	// Result of the last sent mail (OK, Blocked, CSPA charge > 0)
	MissingCharacterMailSentResult *string `json:"missingCharacterMailSentResult,omitempty"`
	// Number of mails sent, is reset when the character is added.
	MissingCharacterMailSentNumber int32 `json:"missingCharacterMailSentNumber,omitempty"`
}
