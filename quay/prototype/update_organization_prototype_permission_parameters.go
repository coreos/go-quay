package prototype

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/coreos/go-quay/models"
)

// NewUpdateOrganizationPrototypePermissionParams creates a new UpdateOrganizationPrototypePermissionParams object
// with the default values initialized.
func NewUpdateOrganizationPrototypePermissionParams() *UpdateOrganizationPrototypePermissionParams {
	var ()
	return &UpdateOrganizationPrototypePermissionParams{}
}

/*UpdateOrganizationPrototypePermissionParams contains all the parameters to send to the API endpoint
for the update organization prototype permission operation typically these are written to a http.Request
*/
type UpdateOrganizationPrototypePermissionParams struct {

	/*Body
	  Request body contents.

	*/
	Body *models.PrototypeUpdate
	/*Orgname
	  The name of the organization

	*/
	Orgname string
	/*Prototypeid
	  The ID of the prototype

	*/
	Prototypeid string
}

// WithBody adds the body to the update organization prototype permission params
func (o *UpdateOrganizationPrototypePermissionParams) WithBody(Body *models.PrototypeUpdate) *UpdateOrganizationPrototypePermissionParams {
	o.Body = Body
	return o
}

// WithOrgname adds the orgname to the update organization prototype permission params
func (o *UpdateOrganizationPrototypePermissionParams) WithOrgname(Orgname string) *UpdateOrganizationPrototypePermissionParams {
	o.Orgname = Orgname
	return o
}

// WithPrototypeid adds the prototypeid to the update organization prototype permission params
func (o *UpdateOrganizationPrototypePermissionParams) WithPrototypeid(Prototypeid string) *UpdateOrganizationPrototypePermissionParams {
	o.Prototypeid = Prototypeid
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateOrganizationPrototypePermissionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	var res []error

	if o.Body == nil {
		o.Body = new(models.PrototypeUpdate)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param orgname
	if err := r.SetPathParam("orgname", o.Orgname); err != nil {
		return err
	}

	// path param prototypeid
	if err := r.SetPathParam("prototypeid", o.Prototypeid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
