package test

import (
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
			QueryParam:  "",
			TotalRecord: 41,
			TotalPage:   5,
			Offset:      0,
			Code:        200,
		},
		{
			QueryParam:  "/hbh",
			Code: 404,
		},
	}
}

///// show all case
func TestCustomers(t *testing.T) {
	cases := cases()
	for i, row := range cases {
		fmt.Println("test case number ", i)
		k := get(customerUri + row.QueryParam)
		if assert.Equal(t, row.Code, k.Code){
			continue
		}
		json := returnResponseJson(k)
		assert.EqualValues(t, row.TotalRecord, returnResponseKey(json, "payload.total_record"))
		assert.EqualValues(t, row.Offset, returnResponseKey(json, "payload.offset"))
		assert.EqualValues(t, row.TotalPage, returnResponseKey(json, "payload.total_page"))
		assert.Equal(t, row.Code, k.Code)
	}
}
