package helpers

import (
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
)

func GetParamFromRequest(g *gin.Context, key string) string {
	return g.DefaultQuery(key, "")
}

func Page(g *gin.Context) int {
	page, _ := strconv.Atoi(g.DefaultQuery("page", "1"))
	return page
}

func Limit(g *gin.Context) int {
	page, _ := strconv.Atoi(g.DefaultQuery("limit", os.Getenv("LIMIT")))
	return page
}
