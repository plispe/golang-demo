package middleware

import "github.com/urfave/negroni"

// NewDispatcher - Constructor
func NewDispatcher() *negroni.Negroni {
	return negroni.New()
}
