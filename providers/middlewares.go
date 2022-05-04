package providers

import (
	"be-task/app/middleware"
	"github.com/gin-gonic/gin"
)

func middlewares(r *gin.Engine) *gin.Engine {
	/// run cors middleware
	r.Use(middleware.CORSMiddleware())
	return r
}
