package test

import (
	"be-task/config"
	"be-task/helpers"
	"be-task/providers"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/http/httptest"
)

var IsSetup = false
var Router *gin.Engine

/**
* define struct that carry the arg of request
 */
type RequestData struct {
	RequestType string
	Url         string
	Data        interface{}
}

/**
* init gin and return gin engine
 */
func setupRouter() *gin.Engine {
	if !IsSetup {
		// connect database
		config.ConnectToDatabase()
		//get countries from memory
		helpers.LoadCountries()
		/// start gin
		Router = providers.Gin()
		IsSetup = true
		Router = providers.Routing(Router)
	}
	return Router
}

/**
* Get request
 */
func get(url string) *httptest.ResponseRecorder {
	return request(RequestData{
		Url:         url,
		RequestType: "GET",
	})
}

/**
* Make new request
 */
func request(request RequestData) *httptest.ResponseRecorder {
	router := setupRouter()
	w := httptest.NewRecorder()
	sendData, _ := json.Marshal(&request.Data)
	req, _ := http.NewRequest(request.RequestType, request.Url, bytes.NewReader(sendData))
	router.ServeHTTP(w, req)
	return w
}

/**
* return response as json
 */
func responseData(c io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(c)
	return buf.String()
}
