package providers

import (
	"github.com/gin-gonic/gin"
)

/**
* start gin framework
* add cross origin middleware
 */
func Gin() *gin.Engine {
	/// init gin
	r := gin.Default()
	middlewares(r)
	return r
}
