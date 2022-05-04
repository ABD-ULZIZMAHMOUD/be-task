package test

import (
	"gopkg.in/thedevsaddam/gojsonq.v2"
	"net/http/httptest"
	_ "testing"
)

func returnResponseJson(k *httptest.ResponseRecorder) gojsonq.JSONQ {
	responseData := responseData(k.Result().Body)
	return *gojsonq.New().JSONString(responseData)

}

func returnResponseKey(recoverResponse gojsonq.JSONQ, key string) interface{} {
	return recoverResponse.Find(key)
}
