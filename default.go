package hygienic

import "github.com/gorilla/mux"

type Hygienic struct {
	Router *mux.Router
}



func New() *Hygienic {
	return  &Hygienic{
		Router: mux.NewRouter(),
	}
}

