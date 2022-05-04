package helpers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func BadRequestResponse(g *gin.Context, msg string, data interface{}) {
	var errors map[string]string
	response(g, msg, data, errors, http.StatusBadRequest, 400, false)
	return
}

func ReturnNotFound(g *gin.Context, msg string) {
	var errors map[string]string
	var data map[string]interface{}
	response(g, msg, data, errors, http.StatusNotFound, 404, false)
	return
}

func OkResponse(g *gin.Context, msg string, data interface{}) {
	var errors map[string]string
	response(g, msg, data, errors, http.StatusOK, 200, true)
	return
}

func response(g *gin.Context, msg string, data interface{}, errors map[string]string, httpStatus int, code int, status bool) {
	g.JSON(httpStatus, gin.H{
		"status":  status,
		"message": msg,
		"errors":  errors,
		"code":    code,
		"payload": data,
	})
	return
}
