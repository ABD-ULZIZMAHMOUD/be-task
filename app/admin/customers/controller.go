package customers

import (
	"be-task/app/transformers"
	"be-task/helpers"
	"github.com/gin-gonic/gin"
)

// Index godoc
// @Summary Retrieve all customers based on country and state filter
// @Produce json
// @Success 200 {object} transformers.CustomerResponseRow
// @Router /customers/all [get]
func Index(g *gin.Context) {
	paginator, rows := GetAllCustomer(g)
	paginator.Records = transformers.CustomersResponses(rows)
	helpers.OkResponse(g, "Done get Items", paginator)
}
