package hygienic

type EndpointFunc func(ctx Context, request interface{}) (response interface{}, err error)

