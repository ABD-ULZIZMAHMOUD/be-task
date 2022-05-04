package routes

import (
	"be-task/app/admin/customers"
	"github.com/gin-gonic/gin"
)

func Admin(r *gin.RouterGroup) *gin.RouterGroup {
	customers.Routes(r)
	return r
}
