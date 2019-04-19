package hygienic
type Middleware func(next EndpointFunc) EndpointFunc
