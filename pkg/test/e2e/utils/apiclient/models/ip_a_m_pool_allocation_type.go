// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// IPAMPoolAllocationType +kubebuilder:validation:Enum=prefix;range
// IPAMPoolAllocationType defines the type of allocation to be used.
// Possible values are `prefix` and `range`.
//
// swagger:model IPAMPoolAllocationType
type IPAMPoolAllocationType string

// Validate validates this IP a m pool allocation type
func (m IPAMPoolAllocationType) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this IP a m pool allocation type based on context it is used
func (m IPAMPoolAllocationType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}