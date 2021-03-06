// Copyright 2019 Hem Design Studio. All rights reserved.
// Use of this source code is governed by a
// license that can be found in the LICENSE file.

package rest

import (
	"encoding/json"
	"github.com/hem-nav-gateway/errorhandler"
	"github.com/hem-nav-gateway/types"
)

func newController() *controller {
	return &controller{}
}

type Service struct {
}

// GetAll handles getting all entities for a specified object types (customer, item ... etc)
// takes fields --> fields to be returned from Navision
// returns a list of requests object values
func (*Service) GetAll(obj types.RequestObject) (interface{}, error) {
	c := newController()
	resValue := c.getAllEntities(obj)
	err := json.Unmarshal(resValue.([]byte), &obj.Response)
	if err != nil {
		return nil, errorhandler.CouldNotUnmarshalData()
	}
	res := obj.Response.(map[string]interface{})
	return res["value"], nil
}

// Filter handles getting specific entities of specified object types (customer, item ... etc)
// takes fields --> fields to be returned from Navision
// takes args --> filter arguments
// returns a list of filtered object values
func (*Service) Filter(obj types.RequestObject) (interface{}, error) {
	c := newController()
	resValue, resError := c.filterByArgs(obj)
	if resError != nil {
		return nil, resError
	}

	err := json.Unmarshal(resValue.([]byte), &obj.Response)
	if err != nil {
		return nil, errorhandler.CouldNotUnmarshalData()
	}
	res := obj.Response.(map[string]interface{})
	values := res["value"].([]interface{})

	if len(values) == 0 {
		return nil, errorhandler.ValueIsNotCorrect(obj.Args)
	}
	return values, nil
}

// Create a specific entity of a specified object type (customer, item ... etc)
// takes fields --> fields to be returned from Navision
// takes args --> arguments are object values to be created
// returns the created object with its return fields
func (*Service) Create(obj types.RequestObject) (interface{}, error) {
	c := newController()
	body, _ := json.Marshal(obj.Args)
	resValue, resError := c.createEntity(obj, body)
	if resError != nil {
		return nil, resError
	}
	err := json.Unmarshal(resValue.([]byte), &obj.Response)
	if err != nil {
		return nil, errorhandler.CouldNotUnmarshalData()
	}
	return obj.Response, nil
}

// Update a specific entity of a specified object type (customer, item ... etc)
// takes fields --> fields to be returned from Navision
// takes args --> arguments are object values to be updated
// returns the created object with its return fields
func (*Service) Update(obj types.RequestObject) (interface{}, error) {
	c := newController()
	var resValue interface{}
	var resError error

	body, _ := json.Marshal(obj.Args)

	if _, ok := obj.Properties["docType"]; ok {

		// In Order to update SalesLines for example
		// It is required to specify "Line_No", "Document_type", "Document_No"
		// In this specific case "Document_No" acts as id
		// this is related to how Microsoft Navision works
		if _, ok := obj.Args["Line_No"]; ok {
			resValue, resError = c.updateEntitybyDocumentTypeAndIDAndLineNo(obj, body)

			// In order to update SalesOrder or SalesInvoice
			// it is required to specify its "No" which acts as its id
			// and "Document_type" which specifies if it is "Order" or "Invoice"
		} else {
			resValue, resError = c.updateEntitybyDocumentTypeAndID(obj, body)
		}

		// This is the case for most entities to be updated
		// "No" which acts as id is all what it required to update an entity
	} else {
		resValue, resError = c.updateEntitybyId(obj, body)
	}

	if resError != nil {
		return nil, resError
	}
	err := json.Unmarshal(resValue.([]byte), &obj.Response)
	if err != nil {
		return nil, errorhandler.CouldNotUnmarshalData()
	}
	return obj.Response, nil
}

// Delete a specific entity of a specified object type (customer, item ... etc)
// takes args --> arguments used to get entity to be deleted
// returns the response code
func Delete(endpoint string, args map[string]interface{}, docType interface{}) (interface{}, error) {
	c := newController()

	var resCode int
	var resError error

	if docType != nil {
		docType := docType.(string)

		// In Order to delete and entity like "SalesLines" for example
		// It is required to specify "Line_No", "Document_type", "Document_No"
		// In this specific case "Document_No" acts as id
		// this is related to how Microsoft Navision works
		if lineNo, ok := args["Line_No"]; ok {
			id := args["Document_No"].(string)
			lineNo := lineNo.(int)
			resCode, resError = c.deleteEntitybyDocumentTypeAndIDAndLineNo(endpoint, id, docType, lineNo)

			// In order to delete SalesOrder or SalesInvoice
			// it is required to specify its "No" which acts as its id
			// and "Document_type" which specifies if it is "Order" or "Invoice"
		} else {
			id := args["No"].(string)
			resCode, resError = c.deleteEntitybyDocumentTypeAndID(endpoint, id, docType)
		}

		// This is the case for most entities to be deleted
		// "No" which acts as id is all what it required to delete an entity
	} else {
		id := args["No"].(string)
		resCode, resError = c.deleteEntitybyId(endpoint, id)
	}

	if resError != nil {
		return nil, resError
	}

	return resCode, nil
}
