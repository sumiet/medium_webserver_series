package rpc

import (
	"context"

	pb "medium/medium_webserver_series/5/proto"
)

type Handler = pb.UsersServer

// handler is the interface which exposes the User Server methods
type handler struct {
	pb.UnimplementedUsersServer
}

// New returns the object for the RPC handler
func New() (Handler, error) {
	return &handler{}, nil
}

// GetUsers function returns the list of users
func (h *handler) GetUsers(ctx context.Context, r *pb.EmptyReq) (*pb.GetUsersResponse, error) {
	return &pb.GetUsersResponse{
		Users: []*pb.User{
			{
				Name: "test user",
				Age:  10,
			},
		},
	}, nil
}
