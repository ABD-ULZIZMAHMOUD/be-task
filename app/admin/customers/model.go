package customers

import (
	"be-task/app/models"
	"be-task/helpers"
	"github.com/gin-gonic/gin"
)

func GetAllCustomer(g *gin.Context) (*helpers.Paginator, []models.Customer) {
	var rows []models.Customer
	paginator := helpers.Paging(&helpers.Param{
		DB:    filter(g),
		Page:  helpers.Page(g),
		Limit: helpers.Limit(g),
	}, &rows)
	return paginator, rows
}
