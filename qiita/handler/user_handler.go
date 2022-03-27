package handler

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"user/api/gen/api"
)

type userHandler struct{}

func NewUserHandle() *userHandler {
	return &userHandler{}
}

func (h *userHandler) Get(ctx context.Context, r *api.UserRequest) (*api.UserResponse, error) {
	if r.Id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "ユーザーの ID を指定してください")
	}

	res := &api.UserResponse{
		User: &api.User{
			Id:   24,
			Name: "John",
		},
	}

	return res, nil
}
