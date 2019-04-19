package hygienic

type GRPCRequestDecoder func(input interface{}) (request interface{}, err error)
type GRPCResponseEncoder func(output interface{}) (response interface{}, err error)
