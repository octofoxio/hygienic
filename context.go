package hygienic

import (
	"mime/multipart"
	"net/http"
	"time"
)

type Context interface {
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key interface{}) interface{}

	FormFile(key string) (*multipart.FileHeader, error)
	Set(key string, value interface{})
	GetHeader(key string) string
	Param(key string) string
	Query(key string) string
	ValidateJSON(rules map[string][]string, data interface{}) error

	GetSession() (interface{}, error)
	GetSessionOrThrow() (interface{})

	Request() *http.Request
}
