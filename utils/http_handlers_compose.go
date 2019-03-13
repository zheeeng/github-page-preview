package utils

import (
	"net/http"
)

// Cont for handy using Continuation func
type Cont func(next func(), complete func(), err func())

// HTTPHandler is typed to a nextable  http.HandlerFunc
type HTTPHandler func(res http.ResponseWriter, req *http.Request) Cont

// HTTPHandlerCondition returns a bool result by http response and request
type HTTPHandlerCondition func(res http.ResponseWriter, req *http.Request) bool

type signal int

var (
	i = signal(0)
	n = signal(1)
	c = signal(2)
	e = signal(3)
	s = i
)

func reset() {
	s = i
}

func next() {
	if s == i {
		s = n
	}
}

func complete() {
	if s == i {
		s = c
	}
}

func err() {
	if s == i {
		s = e
	}
}

func isNext() bool {
	return s == n
}

func isComplete() bool {
	return s == c
}

func isErr() bool {
	return s == e
}

// EmptyHandlerFunc do nothing
func EmptyHandlerFunc(res http.ResponseWriter, req *http.Request) {}

// Of returns HTTPHandler which executes http.HandlerFunc
func Of(handler http.HandlerFunc) HTTPHandler {
	return func(res http.ResponseWriter, req *http.Request) Cont {
		return func(next func(), complete func(), err func()) {
			handler(res, req)
		}
	}
}

// Next returns HTTPHandler which executes http.HandlerFunc and then call next automatically
func Next(handler http.HandlerFunc) HTTPHandler {
	return func(res http.ResponseWriter, req *http.Request) Cont {
		return func(next func(), complete func(), err func()) {
			handler(res, req)
			next()
		}
	}
}

// Complete returns HTTPHandler which executes http.HandlerFunc and then call next automatically
func Complete(handler http.HandlerFunc) HTTPHandler {
	return func(res http.ResponseWriter, req *http.Request) Cont {
		return func(next func(), complete func(), err func()) {
			handler(res, req)
			complete()
		}
	}
}

// Err returns HTTPHandler which executes http.HandlerFunc and then call err automatically
func Err(handler http.HandlerFunc) HTTPHandler {
	return func(res http.ResponseWriter, req *http.Request) Cont {
		return func(next func(), complete func(), err func()) {
			handler(res, req)
			err()
		}
	}
}

// IfElse call a pair HTTPHandler by condition
func IfElse(condFunc HTTPHandlerCondition, handlerFunc1, handlerFunc2 HTTPHandler) HTTPHandler {
	return func(res http.ResponseWriter, req *http.Request) Cont {
		return func(next func(), complete func(), err func()) {
			if condFunc(res, req) {
				handlerFunc1(res, req)(next, complete, err)
			} else {
				handlerFunc2(res, req)(next, complete, err)
			}
		}
	}
}

// Compose compose multiple HTTPHandlers
func Compose(handlers ...HTTPHandler) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		for _, h := range handlers {
			reset()
			h(res, req)(next, complete, err)
			if isNext() {
				continue
			}

			break
		}
	}
}
