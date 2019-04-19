package hygienic

import (
	"mime/multipart"
	"net/http"
	"time"
)

type HTTPContext struct {
	req *http.Request
}

func (ctx *HTTPContext) Deadline() (deadline time.Time, ok bool) {
	panic("implement me")
}

func (ctx *HTTPContext) Done() <-chan struct{} {
	panic("implement me")
}

func (ctx *HTTPContext) Err() error {
	panic("implement me")
}

func (ctx *HTTPContext) Value(key interface{}) interface{} {
	panic("implement me")
}

func (ctx *HTTPContext) FormFile(key string) (*multipart.FileHeader, error) {
	panic("implement me")
}

func (ctx *HTTPContext) Set(key string, value interface{}) {
	panic("implement me")
}

func (ctx *HTTPContext) GetHeader(key string) string {
	panic("implement me")
}

func (ctx *HTTPContext) Param(key string) string {
	panic("implement me")
}

func (ctx *HTTPContext) Query(key string) string {
	panic("implement me")
}

func (ctx *HTTPContext) ValidateJSON(rules map[string][]string, data interface{}) error {
	if err := govalidator.New(govalidator.Options{
		Request: ctx.gin.Request,
		Data:    data,
		Rules:   rules,
	}).ValidateJSON(); len(err) > 0 {
		return entities.NewGlobErr(entities.SolarErrBadInput, "Invalid request body").WithDetail(err)
	} else {
		return nil
	}
}

func (ctx *HTTPContext) GetSession() (interface{}, error) {
	panic("implement me")
}

func (ctx *HTTPContext) GetSessionOrThrow() interface{} {
	panic("implement me")
}

func (ctx *HTTPContext) Request() *http.Request {
	return ctx.req
}

func newHTTPContext(req *http.Request) Context {
	return &HTTPContext{
		req: req,
	}
}
