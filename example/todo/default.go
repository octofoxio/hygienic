package todo

import "github.com/octofoxio/hygienic"

func NewTodoApp() {

	h := hygienic.New()
	h.Router.Methods("GET").PathPrefix("/todos").HandlerFunc()

}
