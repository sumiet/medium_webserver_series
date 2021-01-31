package rpc

import (
	"net/rpc"
)

// Handler is the interface which exposes the User Server methods
type Handler struct {
}

// New returns the object for the RPC handler
func New() *Handler {
	h := &Handler{}
	err := rpc.Register(h)
	if err != nil {
		panic(err)
	}
	return h
}

// GetUsers function returns the list of users
func (rh *Handler) GetUsers(payload int, reply *string) error {
	// add logic to return users
	return nil
}
