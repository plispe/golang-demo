package middleware

import "github.com/urfave/negroni"

func NewDispatcher() *negroni.Negroni {
	return negroni.New()
}
