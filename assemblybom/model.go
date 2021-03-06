// Copyright 2019 Hem Design Studio. All rights reserved.
// Use of this source code is governed by a
// license that can be found in the LICENSE file.

/*
Package assemblybom implements a simple package for handling all graphql
operations related to Microsoft Navision Assembly_Bom page.

Package has a type "AssemblyBom" where all the fields related to Assembly_Bom are defined.

	Example:

		'''
		type AssemblyBom struct {
			No                string  `json:"No"`
			ParentItemNo      string  `json:"Parent_Item_No"`
			Type              string  `json:"Type"`
			...
		}
		'''

GraphQl Object Type along with its fields, arguments and attributes are generated
from the AssemblyBom type when "CreateType" method is invoked.
*/
package assemblybom

import (
	"github.com/graphql-go/graphql"
	"github.com/hem-nav-gateway/config"
	"github.com/hem-nav-gateway/types"
)

// Microsoft Navision endpoint path for Assembly_Bom page
var endpoint = config.AssemblyBomEndpoint

/*
Response is utilized as Microsoft Navision returns a list of objects
when requesting Assembly_Bom, It is utilized for JSON decoding

Example response from Navision

	Example:

		'''
		{
			"value": [
				{
					"No": "1234"
					"Parent_Item_No": "10005",
					"Line_No": 10000,
					"Type": "Item",
					...
				},
				{
					"No": "2345"
					"Parent_Item_No": "10005",
					"Line_No": 20000,
					"Type": "Item",
					...
				},
				{
				...

				},
			}
	'''
*/
type Response struct {
	Value []Model `json:"value"`
}

type Model struct {
	No                string  `json:"No"`
	ParentItemNo      string  `json:"Parent_Item_No"`
	Type              string  `json:"Type"`
	QuantityPer       float64 `json:"Quantity_per"`
	UnitOfMeasureCode string  `json:"Unit_of_Measure_Code"`
}

/*
CreateType function creates a GraphQl Object Type from the
'AssemblyBom' above

Example of GraphQl Object

	Example:
		'''
		graphql.NewObject(graphql.ObjectConfig{
				Name: "AssemblyBom",
				Fields: graphql.Fields{
					"Parent_Item_No":       &graphql.Field{Type: graphql.String},
					"No":                   &graphql.Field{Type: graphql.String},
					"Type":                 &graphql.Field{Type: graphql.String},
					...
				},
			})
		'''

The returned GraphQl arguments will be used as a part of the main mutation

*/
func createType() *graphql.Object {
	return types.GenerateGraphQlType("AssemblyBom", Model{}, nil)
}
