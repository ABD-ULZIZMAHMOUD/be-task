package models

import "be-task/config"

type Customer struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Country string `json:"country"`
}

func (Customer) TableName() string {
	return "customer"
}

func MigrateCustomer() {
	config.DB.AutoMigrate(&Customer{})
}
