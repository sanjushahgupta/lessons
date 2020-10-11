package handler

import (
	"github.com/sanjushahgupta/lessons/sample-rest-di/application"
)

// REST is a struct that handles request/response in a RESTful manner
type REST struct {
	logic *application.Logic
}

// NewREST creates new instance of REST handler
func NewREST(logic *application.Logic) *REST {
	return &REST{
		logic: logic,
	}
}
