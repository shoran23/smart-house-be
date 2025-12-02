package middleware

import "net/http"

// create middleware stack
// stack of middleware functions that can be applied to
// use to create different stacks of middleware which can be applied to

type Middleware func(http.HandlerFunc) *http.HandlerFunc

type MiddlewareStack struct {
	request    http.HandlerFunc
	middleware []Middleware
}

func NewMiddlewareStack() *MiddlewareStack {
	return &MiddlewareStack{
		request:    nil,
		middleware: []Middleware{},
	}
}

func (s *MiddlewareStack) AddMiddleware(m Middleware) {
	s.middleware = append(s.middleware, m)
}

// iterate through the stack and call the final handler func

//func (s *MiddlewareStack) ServeHTTP(last http.HandlerFunc) Middleware {
//	for _, m := range s.middleware {
//		next := m(last)
//	}
//	return next
//}
