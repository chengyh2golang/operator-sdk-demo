package controller

import (
	"operator-sdk-demo/pkg/controller/demo"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, demo.Add)
}
