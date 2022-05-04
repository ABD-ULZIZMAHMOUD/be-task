package customers

import (
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup) {
	usersRoutes := r.Group("customers")
	{
		usersRoutes.GET("/all", Index)
	}
}
