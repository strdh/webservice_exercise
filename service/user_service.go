package service

import (
    "context"
    "exercise/webservice/web/webRequest"
    "exercise/webservice/web/webResponse"
)

type UserService interface {
    Register(ctx context.Context, request webRequest.UserCreateRequest) webResponse.UserResponse
    Login(ctx context.Context, username string, password string) string
}