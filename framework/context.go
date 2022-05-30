package framework

import (
	"context"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type Context struct {
	request        *http.Request
	responseWriter http.ResponseWriter
	ctx            context.Context
	handler        ControllerHandler
	hasTimeOut     bool
	writeMux       *sync.Mutex
}

func NewContext(r *http.Request, w http.ResponseWriter) *Context {
	return &Context{request: r, responseWriter: w, ctx: r.Context(), writeMux: &sync.Mutex{}}
}

func (c *Context) WriteMux() *sync.Mutex {
	return c.writeMux
}

func (c *Context) GetRequest() *http.Request {
	return c.request
}

func (c *Context) GetResponse() http.ResponseWriter {
	return c.responseWriter
}

func (c *Context) SetTimeOut() {
	c.hasTimeOut = true
}

func (c *Context) HasTimeOut() bool {
	return c.hasTimeOut
}

func (c *Context) BaseContext() context.Context {
	return c.request.Context()
}

func (c *Context) Deadline() (deadline time.Time, ok bool) {
	return c.BaseContext().Deadline()
}

func (c *Context) Done() <-chan struct{} {
	return c.BaseContext().Done()
}

func (c *Context) Error() error {
	return c.BaseContext().Err()
}

func (c *Context) Value(key interface{}) interface{} {
	return c.BaseContext().Value(key)
}

//query
func (c *Context) QueryInt(key string, def int) int {
	params := c.QueryAll()
	if vals, ok := params[key]; ok {
		length := len(vals)
		if length > 0 {
			interval, err := strconv.Atoi(vals[length-1])
			if err != nil {
				return def
			}
			return interval
		}
	}
	return def

}
func (c *Context) QueryString(key, def string) string {
	params := c.QueryAll()
	if vals, ok := params[key]; ok {
		length := len(vals)
		if length > 0 {
			return vals[length-1]
		}
		return def
	}
	return def
}

func (c *Context) QueryArray(key string, def []string) []string {
	params := c.QueryAll()
	if vals, ok := params[key]; ok {
		return vals
	}
	return def

}

//Form Post
func (c *Context) name() {

}

func (c *Context) QueryAll() map[string][]string {
	if c.request != nil {
		return c.request.URL.Query()
	}
	return map[string][]string{}
}
func (c *Context) FormAll() map[string][]string {
	if c.request != nil {
		return c.request.PostForm
	}
	return map[string][]string{}
}
