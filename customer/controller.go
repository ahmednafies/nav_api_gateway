// Copyright 2019 Hem Design Studio. All rights reserved.
// Use of this source code is governed by a
// license that can be found in the LICENSE file.

package customer

import (
	"github.com/hem-nav-gateway/request"
)

//GetAll retrieves a List of all CustomerCards available Microsoft Navision .
//Function takes a list of fields to be returned by Microsoft Navision.
func GetAll(fields interface{}) (interface{}, error) {

	return request.GetAll(endpoint, fields, Response{})
}

//Filter retrieves a list of filtered CustomerCards based on a key-value pair
// added by the requester
//Function takes a list of fields to be returned by Microsoft Navision.
//Function takes filter arguments which are required for filtering results in Navision
func Filter(fields, args interface{}) (interface{}, error) {

	return request.Filter(endpoint, fields, args, Response{})

}

//Create creates a CustomerCard objects based on arguments
// added by the requester
//Function takes a list of fields to be returned by Microsoft Navision after creation.
func Create(fields, args interface{}) (interface{}, error) {

	return request.Create(endpoint, fields, args, Response{})

}

//Update modifies a certain CustomerCard Object Microsoft Navision.
//Function takes filter arguments which are required identifying
//the specific object to be updated/modified.
//Function returns a list of AssemblyBom values
func Update(fields, args interface{}) (interface{}, error) {

	return request.Update(endpoint, fields, args, nil, Response{})

}
