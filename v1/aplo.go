package aplo

import (
	"fmt"
	"net/http"
)

type contextFunc func(c Context) error

type Context struct {
	aplo    *Aplo
	writer  http.ResponseWriter
	request *http.Request
}

func (c *Context) String(i int, s string) error {
	c.writer.WriteHeader(i)
	_, err := c.writer.Write([]byte(s))
	return err
}

type Aplo struct {
	serveMux *http.ServeMux
}

func NewAplo() *Aplo {
	return &Aplo{
		serveMux: http.NewServeMux(),
	}
}

func makeFunc(c contextFunc, a *Aplo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		context := Context{
			aplo:    a,
			writer:  w,
			request: r,
		}

		c(context)
	}
}

func (a *Aplo) Serve(s string) error {
	return http.ListenAndServe(s, a.serveMux)
}

func (a *Aplo) GET(s string, c contextFunc) {
	a.serveMux.HandleFunc(fmt.Sprintf("GET %s", s), makeFunc(c, a))
}
