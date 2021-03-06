package secscan

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetRepoImageSecurityParams creates a new GetRepoImageSecurityParams object
// with the default values initialized.
func NewGetRepoImageSecurityParams() *GetRepoImageSecurityParams {
	var ()
	return &GetRepoImageSecurityParams{}
}

/*GetRepoImageSecurityParams contains all the parameters to send to the API endpoint
for the get repo image security operation typically these are written to a http.Request
*/
type GetRepoImageSecurityParams struct {

	/*Imageid
	  The image ID

	*/
	Imageid string
	/*Repository
	  The full path of the repository. e.g. namespace/name

	*/
	Repository string
	/*Vulnerabilities
	  Include vulnerabilities informations

	*/
	Vulnerabilities *bool
}

// WithImageid adds the imageid to the get repo image security params
func (o *GetRepoImageSecurityParams) WithImageid(Imageid string) *GetRepoImageSecurityParams {
	o.Imageid = Imageid
	return o
}

// WithRepository adds the repository to the get repo image security params
func (o *GetRepoImageSecurityParams) WithRepository(Repository string) *GetRepoImageSecurityParams {
	o.Repository = Repository
	return o
}

// WithVulnerabilities adds the vulnerabilities to the get repo image security params
func (o *GetRepoImageSecurityParams) WithVulnerabilities(Vulnerabilities *bool) *GetRepoImageSecurityParams {
	o.Vulnerabilities = Vulnerabilities
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetRepoImageSecurityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	var res []error

	// path param imageid
	if err := r.SetPathParam("imageid", o.Imageid); err != nil {
		return err
	}

	// path param repository
	if err := r.SetPathParam("repository", o.Repository); err != nil {
		return err
	}

	if o.Vulnerabilities != nil {

		// query param vulnerabilities
		var qrVulnerabilities bool
		if o.Vulnerabilities != nil {
			qrVulnerabilities = *o.Vulnerabilities
		}
		qVulnerabilities := swag.FormatBool(qrVulnerabilities)
		if qVulnerabilities != "" {
			if err := r.SetQueryParam("vulnerabilities", qVulnerabilities); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
