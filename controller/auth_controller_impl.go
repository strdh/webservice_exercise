package controller

import (
    "net/http"
    "exercise/webservice/helper"
    "exercise/webservice/web/webRequest"
    "exercise/webservice/web/webResponse"
    "exercise/webservice/service"
    "github.com/julienschmidt/httprouter"
)

type AuthControllerImpl struct {
    UserService service.UserService
}

func NewAuthController() AuthController {
    userService := service.NewUserService()
    return &AuthControllerImpl{
        UserService: userService,
    }
}

func (controller *AuthControllerImpl) Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
    userCreateRequest := webRequest.UserCreateRequest{}
    helper.ReadFromRequestBody(request, &userCreateRequest)

    userResponse := controller.UserService.Register(request.Context(), userCreateRequest)
    webResponse := webResponse.WebResponse{
        Code:200,
        Status: "User registered successfully",
        Data: userResponse,
    }

    helper.WriteToResponseBody(writer, webResponse)
}

func (controller *AuthControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
    userLoginRequest := webRequest.UserLoginRequest{}
    helper.ReadFromRequestBody(request, &userLoginRequest)

    userResponse := controller.UserService.Login(request.Context(), userLoginRequest.Username, userLoginRequest.Password)
    webResponse := webResponse.WebResponse{
        Code:200,
        Status: "User logged in successfully",
        Data: userResponse,
    }

    helper.WriteToResponseBody(writer, webResponse)
}

