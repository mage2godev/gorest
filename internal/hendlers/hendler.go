package hendlers

import "github.com/julienschmidt/httprouter"

type Hendler interface {
	Register(router *httprouter.Router)
}
