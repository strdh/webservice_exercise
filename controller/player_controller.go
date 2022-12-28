package controller

import (
    "net/http"
    "github.com/julienschmidt/httprouter"
)

type PlayerController interface {
    Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
    Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
    Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
    FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
    GetAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}