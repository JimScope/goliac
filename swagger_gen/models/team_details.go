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
	"github.com/go-openapi/validate"
)

// TeamDetails team details
//
// swagger:model teamDetails
type TeamDetails struct {

	// members
	Members []*TeamDetailsMembersItems0 `json:"members"`

	// name
	Name string `json:"name,omitempty"`

	// owners
	Owners []*TeamDetailsOwnersItems0 `json:"owners"`

	// path
	Path string `json:"path,omitempty"`

	// repositories
	Repositories []*Repository `json:"repositories"`
}

// Validate validates this team details
func (m *TeamDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMembers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwners(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRepositories(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TeamDetails) validateMembers(formats strfmt.Registry) error {
	if swag.IsZero(m.Members) { // not required
		return nil
	}

	for i := 0; i < len(m.Members); i++ {
		if swag.IsZero(m.Members[i]) { // not required
			continue
		}

		if m.Members[i] != nil {
			if err := m.Members[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("members" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("members" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TeamDetails) validateOwners(formats strfmt.Registry) error {
	if swag.IsZero(m.Owners) { // not required
		return nil
	}

	for i := 0; i < len(m.Owners); i++ {
		if swag.IsZero(m.Owners[i]) { // not required
			continue
		}

		if m.Owners[i] != nil {
			if err := m.Owners[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("owners" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("owners" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TeamDetails) validateRepositories(formats strfmt.Registry) error {
	if swag.IsZero(m.Repositories) { // not required
		return nil
	}

	for i := 0; i < len(m.Repositories); i++ {
		if swag.IsZero(m.Repositories[i]) { // not required
			continue
		}

		if m.Repositories[i] != nil {
			if err := m.Repositories[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("repositories" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("repositories" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this team details based on the context it is used
func (m *TeamDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMembers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOwners(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRepositories(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TeamDetails) contextValidateMembers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Members); i++ {

		if m.Members[i] != nil {

			if swag.IsZero(m.Members[i]) { // not required
				return nil
			}

			if err := m.Members[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("members" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("members" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TeamDetails) contextValidateOwners(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Owners); i++ {

		if m.Owners[i] != nil {

			if swag.IsZero(m.Owners[i]) { // not required
				return nil
			}

			if err := m.Owners[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("owners" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("owners" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TeamDetails) contextValidateRepositories(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Repositories); i++ {

		if m.Repositories[i] != nil {

			if swag.IsZero(m.Repositories[i]) { // not required
				return nil
			}

			if err := m.Repositories[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("repositories" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("repositories" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TeamDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TeamDetails) UnmarshalBinary(b []byte) error {
	var res TeamDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TeamDetailsMembersItems0 team details members items0
//
// swagger:model TeamDetailsMembersItems0
type TeamDetailsMembersItems0 struct {

	// external
	External bool `json:"external"`

	// githubid
	Githubid string `json:"githubid,omitempty"`

	// name
	// Min Length: 1
	Name string `json:"name,omitempty"`
}

// Validate validates this team details members items0
func (m *TeamDetailsMembersItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TeamDetailsMembersItems0) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("name", "body", m.Name, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this team details members items0 based on context it is used
func (m *TeamDetailsMembersItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TeamDetailsMembersItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TeamDetailsMembersItems0) UnmarshalBinary(b []byte) error {
	var res TeamDetailsMembersItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TeamDetailsOwnersItems0 team details owners items0
//
// swagger:model TeamDetailsOwnersItems0
type TeamDetailsOwnersItems0 struct {

	// external
	External bool `json:"external"`

	// githubid
	Githubid string `json:"githubid,omitempty"`

	// name
	// Min Length: 1
	Name string `json:"name,omitempty"`
}

// Validate validates this team details owners items0
func (m *TeamDetailsOwnersItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TeamDetailsOwnersItems0) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("name", "body", m.Name, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this team details owners items0 based on context it is used
func (m *TeamDetailsOwnersItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TeamDetailsOwnersItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TeamDetailsOwnersItems0) UnmarshalBinary(b []byte) error {
	var res TeamDetailsOwnersItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
