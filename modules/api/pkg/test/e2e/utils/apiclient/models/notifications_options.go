// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NotificationsOptions notifications options
//
// swagger:model NotificationsOptions
type NotificationsOptions struct {

	// HideErrorEvents will silence error events for the dashboard.
	HideErrorEvents bool `json:"hideErrorEvents,omitempty"`

	// HideErrors will silence error notifications for the dashboard.
	HideErrors bool `json:"hideErrors,omitempty"`
}

// Validate validates this notifications options
func (m *NotificationsOptions) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this notifications options based on context it is used
func (m *NotificationsOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NotificationsOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NotificationsOptions) UnmarshalBinary(b []byte) error {
	var res NotificationsOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}