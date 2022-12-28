package controller 

import (
    // "context"
    "net/http"
    "strconv"
    "github.com/julienschmidt/httprouter"
    "exercise/webservice/web/webRequest"
    "exercise/webservice/web/webResponse"
    "exercise/webservice/helper"
    "exercise/webservice/service"
)

type PlayerControllerImpl struct {
    PlayerService service.PlayerService
}

func NewPlayerController(playerService service.PlayerService) PlayerController {
    return &PlayerControllerImpl {
        PlayerService: playerService,
    }
}

func (controller *PlayerControllerImpl) Create(writer http.ResponseWriter, request*http.Request, params httprouter.Params) {
    playerCreateRequest := webRequest.PlayerCreateRequest{}
    helper.ReadFromRequestBody(request, &playerCreateRequest)

    playerResponse := controller.PlayerService.Create(request.Context(), playerCreateRequest)
    webResponse := webResponse.WebResponse{
        Code: 200,
        Status: "Player created succesfully",
        Data: playerResponse,
    }

    helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PlayerControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
    playerUpdateRequest := webRequest.PlayerUpdateRequest{}
    helper.ReadFromRequestBody(request, &playerUpdateRequest)

    playerId, err := strconv.Atoi(params.ByName("id"))
    helper.PanicIfError(err)
    playerUpdateRequest.Id = playerId

    playerResponse := controller.PlayerService.Update(request.Context(), playerUpdateRequest)
    webResponse := webResponse.WebResponse{
        Code: 200,
        Status: "Player updated successfully",
        Data: playerResponse,
    }

    helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PlayerControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
    playerId, err := strconv.Atoi(params.ByName("id"))
    helper.PanicIfError(err)

    controller.PlayerService.Delete(request.Context(), playerId)
    webResponse := webResponse.WebResponse{
        Code: 200,
        Status: "Player deleted successfully",
    }

    helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PlayerControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
    playerId, err := strconv.Atoi(params.ByName("id"))
    helper.PanicIfError(err)

    playerResponse := controller.PlayerService.FindById(request.Context(), playerId)
    webResponse := webResponse.WebResponse{
        Code: 200,
        Status: "Player found successfully",
        Data: playerResponse,
    }

    helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PlayerControllerImpl) GetAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
    playerResponses := controller.PlayerService.GetAll(request.Context())
    webResponse := webResponse.WebResponse{
        Code: 200,
        Status: "OK",
        Data: playerResponses,
    }

    helper.WriteToResponseBody(writer, webResponse)
}