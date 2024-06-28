// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StatusDetail status detail
//
// swagger:model statusDetail
type StatusDetail struct {

	// finished at
	FinishedAt string `json:"FinishedAt,omitempty"`

	// log
	Log string `json:"Log,omitempty"`

	// name
	Name string `json:"Name,omitempty"`

	// nodes
	Nodes []*StatusNode `json:"Nodes"`

	// on cancel
	OnCancel *StatusNode `json:"OnCancel,omitempty"`

	// on exit
	OnExit *StatusNode `json:"OnExit,omitempty"`

	// on failure
	OnFailure *StatusNode `json:"OnFailure,omitempty"`

	// on success
	OnSuccess *StatusNode `json:"OnSuccess,omitempty"`

	// params
	Params string `json:"Params,omitempty"`

	// pid
	Pid int64 `json:"Pid,omitempty"`

	// request Id
	RequestID string `json:"RequestId,omitempty"`

	// started at
	StartedAt string `json:"StartedAt,omitempty"`

	// status
	Status int64 `json:"Status,omitempty"`

	// status text
	StatusText string `json:"StatusText,omitempty"`
}

// Validate validates this status detail
func (m *StatusDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNodes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOnCancel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOnExit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOnFailure(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOnSuccess(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StatusDetail) validateNodes(formats strfmt.Registry) error {
	if swag.IsZero(m.Nodes) { // not required
		return nil
	}

	for i := 0; i < len(m.Nodes); i++ {
		if swag.IsZero(m.Nodes[i]) { // not required
			continue
		}

		if m.Nodes[i] != nil {
			if err := m.Nodes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Nodes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Nodes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StatusDetail) validateOnCancel(formats strfmt.Registry) error {
	if swag.IsZero(m.OnCancel) { // not required
		return nil
	}

	if m.OnCancel != nil {
		if err := m.OnCancel.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("OnCancel")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("OnCancel")
			}
			return err
		}
	}

	return nil
}

func (m *StatusDetail) validateOnExit(formats strfmt.Registry) error {
	if swag.IsZero(m.OnExit) { // not required
		return nil
	}

	if m.OnExit != nil {
		if err := m.OnExit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("OnExit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("OnExit")
			}
			return err
		}
	}

	return nil
}

func (m *StatusDetail) validateOnFailure(formats strfmt.Registry) error {
	if swag.IsZero(m.OnFailure) { // not required
		return nil
	}

	if m.OnFailure != nil {
		if err := m.OnFailure.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("OnFailure")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("OnFailure")
			}
			return err
		}
	}

	return nil
}

func (m *StatusDetail) validateOnSuccess(formats strfmt.Registry) error {
	if swag.IsZero(m.OnSuccess) { // not required
		return nil
	}

	if m.OnSuccess != nil {
		if err := m.OnSuccess.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("OnSuccess")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("OnSuccess")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this status detail based on the context it is used
func (m *StatusDetail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNodes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOnCancel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOnExit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOnFailure(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOnSuccess(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StatusDetail) contextValidateNodes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Nodes); i++ {

		if m.Nodes[i] != nil {

			if swag.IsZero(m.Nodes[i]) { // not required
				return nil
			}

			if err := m.Nodes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Nodes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Nodes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StatusDetail) contextValidateOnCancel(ctx context.Context, formats strfmt.Registry) error {

	if m.OnCancel != nil {

		if swag.IsZero(m.OnCancel) { // not required
			return nil
		}

		if err := m.OnCancel.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("OnCancel")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("OnCancel")
			}
			return err
		}
	}

	return nil
}

func (m *StatusDetail) contextValidateOnExit(ctx context.Context, formats strfmt.Registry) error {

	if m.OnExit != nil {

		if swag.IsZero(m.OnExit) { // not required
			return nil
		}

		if err := m.OnExit.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("OnExit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("OnExit")
			}
			return err
		}
	}

	return nil
}

func (m *StatusDetail) contextValidateOnFailure(ctx context.Context, formats strfmt.Registry) error {

	if m.OnFailure != nil {

		if swag.IsZero(m.OnFailure) { // not required
			return nil
		}

		if err := m.OnFailure.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("OnFailure")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("OnFailure")
			}
			return err
		}
	}

	return nil
}

func (m *StatusDetail) contextValidateOnSuccess(ctx context.Context, formats strfmt.Registry) error {

	if m.OnSuccess != nil {

		if swag.IsZero(m.OnSuccess) { // not required
			return nil
		}

		if err := m.OnSuccess.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("OnSuccess")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("OnSuccess")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StatusDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StatusDetail) UnmarshalBinary(b []byte) error {
	var res StatusDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}