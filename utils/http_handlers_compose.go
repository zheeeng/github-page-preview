package utils

import (
	"net/http"
)

// Cont for handy using Continuation func
type Cont func(next func(), complete func(), err func())

// HTTPHandler is typed to a nextable  http.HandlerFunc
type HTTPHandler func(res http.ResponseWriter, req *http.Request) Cont

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

// NextHandler returns HTTPHandler which executes http.HandlerFunc and then call next automatically
func NextHandler(handler http.HandlerFunc) HTTPHandler {
	return func(res http.ResponseWriter, req *http.Request) Cont {
		return func(next func(), complete func(), err func()) {
			handler(res, req)
			next()
		}
	}
}

// CompleteHandler returns HTTPHandler which executes http.HandlerFunc and then call next automatically
func CompleteHandler(handler http.HandlerFunc) HTTPHandler {
	return func(res http.ResponseWriter, req *http.Request) Cont {
		return func(next func(), complete func(), err func()) {
			handler(res, req)
			complete()
		}
	}
}

// HTTPHandlersCompose compose multiple HTTPHandlers
func HTTPHandlersCompose(handlers ...HTTPHandler) http.HandlerFunc {
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
