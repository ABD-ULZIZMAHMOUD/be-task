package test

import (
	"be-task/app/models"
	"be-task/app/transformers"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var customerUri = "/customers/all"

type Cases struct {
	QueryParam  string `json:"query_param"`
	TotalRecord int    `json:"total_record"`
	TotalPage   int    `json:"total_page"`
	Offset      int    `json:"offset"`
	Code        int    `json:"code"`
}

func cases() []Cases {
	return []Cases{
		{
			QueryParam:  "/wrong",
			Code: 404,
		},
		{
			QueryParam:  "",
			TotalRecord: 41,
			TotalPage:   5,
			Offset:      0,
			Code:        200,
		},
		{
			QueryParam:  "?limit=50",
			TotalRecord: 41,
			TotalPage:   1,
			Offset:      0,
			Code:        200,
		},
		{
			QueryParam:  "?limit=10",
			TotalRecord: 41,
			TotalPage:   5,
			Offset:      0,
			Code:        200,
		},
		{
			QueryParam:  "?state=valid",
			TotalRecord: 27,
			TotalPage:   3,
			Offset:      0,
			Code:        200,
		},
		{
			QueryParam:  "?state=not_valid",
			TotalRecord: 14,
			TotalPage:   2,
			Offset:      0,
			Code:        200,
		},
		{
			QueryParam:  "?country=Cameroon",
			TotalRecord: 7,
			TotalPage:   1,
			Offset:      0,
			Code:        200,
		},
		{
			QueryParam:  "?country=Cameroon&state=valid",
			TotalRecord: 7,
			TotalPage:   1,
			Offset:      0,
			Code:        200,
		},
		{
			QueryParam:  "?country=Cameroon&state=not_valid",
			TotalRecord: 0,
			TotalPage:   0,
			Offset:      0,
			Code:        200,
		},

	}
}

//integration test example
func TestCustomers(t *testing.T) {
	cases := cases()
	for i, row := range cases {
		fmt.Println("test case number ", i)
		k := get(customerUri + row.QueryParam)
		assert.Equal(t, row.Code, k.Code)
		if k.Code==200{
			json := returnResponseJson(k)
			assert.EqualValues(t, row.TotalRecord, returnResponseKey(json, "payload.total_record"))
			assert.EqualValues(t, row.Offset, returnResponseKey(json, "payload.offset"))
			assert.EqualValues(t, row.TotalPage, returnResponseKey(json, "payload.total_page"))
		}
	}
}

// unit test example
func TestCustomersModel(t *testing.T) {
	customers:=make([]models.Customer,5)
	rows:=transformers.CustomersResponses(customers)
	assert.EqualValues(t, len(customers),len(rows))
}
