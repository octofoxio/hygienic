package hygienic

import "net/http"

type HTTPRequestDecoder func(Context, *http.Request) (request interface{}, err error)
type HTTPResponseEncoder func(Context, interface{}) (response *HTTPResponse, err error)

type HTTPResponse struct {
	Code    int         `json:"-"` // if Code is status move permanently, kernel will use Result as string to redirect
	Result  interface{} `json:"result,omitempty"`
	Message string      `json:"message,omitempty"`
	Detail  []string `json:"detail,omitempty"`
	Debug   interface{} `json:"debug,omitempty"`
}

type HTTPRespose_PaginationResult struct {
	Count   int         `json:"count"`
	Limit   int         `json:"limit"`
	Offset  int         `json:"offset"`
	Records interface{} `json:"records"`
}


func NewHTTPTransports(endpoints EndpointFunc,reqDec HTTPRequestDecoder, resEnc HTTPResponseEncoder) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {



	}
}

