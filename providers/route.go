package providers

import (
	"be-task/routes"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
)

func Routing(r *gin.Engine) *gin.Engine {
	admin := r.Group(os.Getenv("ADMIN_SLUG"))
	routes.Admin(admin)
	swaggerRouting(r)
	return r
}

//add route to swagger
func swaggerRouting(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
