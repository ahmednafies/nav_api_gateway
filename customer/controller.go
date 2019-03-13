// Copyright 2019 Hem Design Studio. All rights reserved.
// Use of this source code is governed by a
// license that can be found in the LICENSE file.

package customer

import (
	"github.com/graphql-go/graphql"
	"github.com/hem-nav-gateway/request"
	"github.com/hem-nav-gateway/types"
)

var _type = createType()
var _args = createArgs()

type Request struct {
	Name    string
	Company string
	Object  types.RequestObject
}

func (*Request) CreateType() *graphql.Object {
	return _type
}

func (*Request) CreateArgs() map[string]*graphql.ArgumentConfig {
	return _args
}

func (r *Request) GetName() string {
	r.Name = "Customer"
	return r.Name
}

func (r *Request) GetCompany() string {
	return r.Company
}

func (r *Request) SetArgs(args map[string]interface{}) {
	r.Object.Args = args
}

func (r *Request) SetFields(fields []string) {
	r.Object.Fields = fields
}

// GetAll retrieves a List of all CustomerCards available Microsoft Navision .
// Function takes a list of fields to be returned by Microsoft Navision.
func (r *Request) GetAll() (interface{}, error) {
	r.Object.Endpoint = endpoint
	r.Object.Company = r.Company
	return request.GetAll(r.Object, Response{})
}

// Filter retrieves a list of filtered CustomerCards based on a key-value pair added by the requester
// Function takes a list of fields to be returned by Microsoft Navision.
// Function takes filter arguments which are required for filtering results in Navision
func (r *Request) Filter() (interface{}, error) {
	r.Object.Endpoint = endpoint
	r.Object.Company = r.Company
	return request.Filter(r.Object, Response{})

}

// Create creates a CustomerCard objects based on arguments added by the requester
// Function takes a list of fields to be returned by Microsoft Navision after creation.
func (r *Request) Create() (interface{}, error) {
	r.Object.Endpoint = endpoint
	r.Object.Company = r.Company
	return request.Create(r.Object, Response{})

}

// Update modifies a certain CustomerCard Object Microsoft Navision.
// Function takes filter arguments which are required identifying
// the specific object to be updated/modified.
// Function returns a list of AssemblyBom values
func (r *Request) Update() (interface{}, error) {
	r.Object.Endpoint = endpoint
	r.Object.Company = r.Company
	return request.Update(r.Object, Response{})

}
