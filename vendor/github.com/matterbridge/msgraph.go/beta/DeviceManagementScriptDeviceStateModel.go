// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// DeviceManagementScriptDeviceState Contains properties for device run state of the device management script.
type DeviceManagementScriptDeviceState struct {
	// Entity is the base model of DeviceManagementScriptDeviceState
	Entity
	// RunState State of latest run of the device management script.
	RunState *RunState `json:"runState,omitempty"`
	// ResultMessage Details of execution output.
	ResultMessage *string `json:"resultMessage,omitempty"`
	// LastStateUpdateDateTime Latest time the device management script executes.
	LastStateUpdateDateTime *time.Time `json:"lastStateUpdateDateTime,omitempty"`
	// ErrorCode Error code corresponding to erroneous execution of the device management script.
	ErrorCode *int `json:"errorCode,omitempty"`
	// ErrorDescription Error description corresponding to erroneous execution of the device management script.
	ErrorDescription *string `json:"errorDescription,omitempty"`
	// ManagedDevice undocumented
	ManagedDevice *ManagedDevice `json:"managedDevice,omitempty"`
}