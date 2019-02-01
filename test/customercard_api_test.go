package test

import (
	"encoding/json"
	"fmt"
	"github.com/nav-api-gateway/customer"
	"github.com/nav-api-gateway/test/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

type CustomerCardResponseBody struct {
	Data   CustomerCardData `json:"data"`
	Errors []ErrorMessage   `json:"errors"`
}

type ErrorMessage struct {
	Message interface{} `json:"message"`
}

type CustomerCardData struct {
	CustomerCard       []customer.CustomerCard `json:"CustomerCard"`
	CreateCustomerCard customer.CustomerCard   `json:"CreateCustomerCard"`
}

func TestGetAllCustomerCard(t *testing.T) {
	resBody := CustomerCardResponseBody{}
	page := utils.Query.CustomerCard
	attrs := utils.GetCustomerCardAttrs()
	query := utils.GetAllQuery(page, attrs)
	resCode, resBodyInBytes := utils.Client("GET", query, nil)
	json.Unmarshal(resBodyInBytes, &resBody)
	element := resBody.Data.CustomerCard[0]
	assert.Equal(t, 200, resCode, "Response code is 200 as expected")
	assert.NotNil(t, element.No, "No should not be Nil")

}

func TestFilterCustomerCard(t *testing.T) {
	resBody := CustomerCardResponseBody{}
	page := utils.Query.CustomerCard
	attrs := utils.GetCustomerCardAttrs()
	args := utils.GetCustomerCardArgs().FilterArgs
	queryList := utils.GetQueryList(page, attrs, args)

	for _, query := range queryList {
		resCode, resBodyInBytes := utils.Client("GET", query, nil)
		json.Unmarshal(resBodyInBytes, &resBody)

		assert.Equal(t, 200, resCode, "Response code is 200 as expected")
		for _, element := range resBody.Data.CustomerCard {
			values := utils.Serialize(element)
			for _, val := range values {
				assert.NotNil(t, val)
			}
		}

	}
	navNo := args[0]["value"]
	assert.Equal(t, navNo, resBody.Data.CustomerCard[0].No, fmt.Sprintf("Expected No = %s", navNo))
}

func TestCreateCustomerCard(t *testing.T) {
	resBody := CustomerCardResponseBody{}
	page := utils.Mutation.CreateCustomerCard
	attrs := utils.GetCustomerCardAttrs()
	args := utils.GetCustomerCardArgs().CreateArgs
	body := utils.GetPOSTBody(page, attrs, args)
	resCode, resBodyInBytes := utils.Client("POST", "", body)
	json.Unmarshal(resBodyInBytes, &resBody)
	assert.Equal(t, 200, resCode, "Response code is 200 as expected")
	element := resBody.Data.CreateCustomerCard
	values := utils.Serialize(element)
	for _, val := range values {
		assert.NotNil(t, val)
	}

}
