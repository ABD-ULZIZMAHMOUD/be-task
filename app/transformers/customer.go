package transformers

import (
	"be-task/app/models"
	"be-task/helpers"
	"regexp"
)

type CustomerResponseRow struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Country string `json:"country"`
	State   string `json:"state"`
}

func CustomerResponse(row models.Customer) CustomerResponseRow {
	country := getCountry(row.Phone)
	stata := "in-valid"
	if country != "" {
		stata = "valid"
	}
	return CustomerResponseRow{
		ID:      row.ID,
		Name:    row.Name,
		Phone:   row.Phone,
		Country: country,
		State:   stata,
	}
}

func CustomersResponses(rows []models.Customer) []CustomerResponseRow {
	result := make([]CustomerResponseRow, 0)
	for _, value := range rows {
		result = append(result, CustomerResponse(value))
	}
	return result
}

func getCountry(phone string) string {
	for value, regex := range helpers.Countries {
		match, _ := regexp.MatchString(regex, phone)
		if match {
			return value
		}
	}
	return ""
}
