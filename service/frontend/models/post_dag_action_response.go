// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostDagActionResponse post dag action response
//
// swagger:model postDagActionResponse
type PostDagActionResponse struct {

	// new dag ID
	NewDagID string `json:"NewDagID,omitempty"`
}

// Validate validates this post dag action response
func (m *PostDagActionResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post dag action response based on context it is used
func (m *PostDagActionResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostDagActionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostDagActionResponse) UnmarshalBinary(b []byte) error {
	var res PostDagActionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
