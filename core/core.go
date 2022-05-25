package core

import "net/http"

type Core struct {
	Server int
}

func NewCore() *Core {
	return &Core{}
}

func (c *Core) ServerHTTP(response http.ResponseWriter, request http.Request) {

}
