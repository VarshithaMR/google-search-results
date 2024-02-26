package writer

import (
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

const (
	contentType = "Content-Type"
	appType     = "application/json"
)

func WriteResponse(rw http.ResponseWriter, resp interface{}, responseCode int) {
	rw.WriteHeader(responseCode)
	rw.Header().Set(contentType, appType)
	bytes, _ := jsoniter.Marshal(resp)
	rw.Write(bytes)
}
