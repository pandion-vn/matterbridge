// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// ConditionalAccessPolicy undocumented
type ConditionalAccessPolicy struct {
	// Entity is the base model of ConditionalAccessPolicy
	Entity
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// ModifiedDateTime undocumented
	ModifiedDateTime *time.Time `json:"modifiedDateTime,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// State undocumented
	State *ConditionalAccessPolicyState `json:"state,omitempty"`
	// Conditions undocumented
	Conditions *ConditionalAccessConditionSet `json:"conditions,omitempty"`
	// GrantControls undocumented
	GrantControls *ConditionalAccessGrantControls `json:"grantControls,omitempty"`
	// SessionControls undocumented
	SessionControls *ConditionalAccessSessionControls `json:"sessionControls,omitempty"`
}