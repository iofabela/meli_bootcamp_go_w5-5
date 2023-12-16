package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
)

func CreateRequestTest(method string, url string, body interface{}) (*http.Request, *httptest.ResponseRecorder) {
	bodyString, _ := json.Marshal(body)
	req := httptest.NewRequest(method, url, bytes.NewBuffer(bodyString))
	req.Header.Add("Content-Type", "application/json")
	return req, httptest.NewRecorder()
}
