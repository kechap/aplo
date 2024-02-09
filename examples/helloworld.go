package main

import (
	"net/http"

	"github.com/kechap/aplo/v1"
)

func main() {
	a := aplo.NewAplo()
	a.GET("/index", func(c aplo.Context) error {
		return c.String(http.StatusOK, "Hello world")
	})
	panic(a.Serve(":3000"))
}
