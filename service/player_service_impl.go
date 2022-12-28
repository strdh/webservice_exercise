package service

import (
    "context"
    "fmt"
    "database/sql"
    "exercise/webservice/web/webRequest"
    "exercise/webservice/web/webResponse"
    "exercise/webservice/model"
    "exercise/webservice/helper"
    "exercise/webservice/repository"
    "exercise/webservice/exception"
    "github.com/go-playground/validator/v10"
)

type PlayerServiceImpl struct {
    PlayerRepository repository.PlayerRepository
    DB *sql.DB 
    Validate *validator.Validate
}

func NewPlayerService(playerRepository repository.PlayerRepository, db *sql.DB, validate *validator.Validate) PlayerService {
    return &PlayerServiceImpl{
        PlayerRepository: playerRepository,
        DB: db,
        Validate: validate,
    }
}

func (service *PlayerServiceImpl) Create(ctx context.Context, request webRequest.PlayerCreateRequest) webResponse.PlayerResponse {
    err := service.Validate.Struct(request)
    helper.PanicIfError(err)

    tx, err := service.DB.Begin()
    helper.PanicIfError(err)
    defer helper.CommitOrRollback(tx)

    player := model.Player{
        Name: request.Name,
        Position: request.Position,
        Height: request.Height,
        Weight: request.Weight,
        BirthDate: request.BirthDate,
    }

    player = service.PlayerRepository.Create(ctx, tx, player)
    return helper.ToPlayerResponse(player)
}

func (service *PlayerServiceImpl) Update(ctx context.Context, request webRequest.PlayerUpdateRequest) webResponse.PlayerResponse {
    err := service.Validate.Struct(request)
    helper.PanicIfError(err)

    tx, err := service.DB.Begin()
    helper.PanicIfError(err)
    defer helper.CommitOrRollback(tx)

    player, err := service.PlayerRepository.FindById(ctx, tx, request.Id)
    if err != nil {
        panic(exception.NewNotFoundError(fmt.Sprintf("Player is not found with id %d", request.Id)))
    }

    player.Name = request.Name
    player.Position = request.Position
    player.Height = request.Height
    player.Weight = request.Weight
    player.BirthDate = request.BirthDate

    player = service.PlayerRepository.Update(ctx, tx, player)

    return helper.ToPlayerResponse(player)
}

func (service *PlayerServiceImpl) Delete(ctx context.Context, playerId int) {
    tx, err := service.DB.Begin()
    helper.PanicIfError(err)
    defer helper.CommitOrRollback(tx)

    player, err := service.PlayerRepository.FindById(ctx, tx, playerId)
    if err != nil {
        panic(exception.NewNotFoundError(fmt.Sprintf("Player is not found with id %d", playerId)))
    }

    service.PlayerRepository.Delete(ctx, tx, player)
}

func (service *PlayerServiceImpl) FindById(ctx context.Context, playerId int) webResponse.PlayerResponse {
    tx, err := service.DB.Begin()
    helper.PanicIfError(err)
    defer helper.CommitOrRollback(tx)

    player, err := service.PlayerRepository.FindById(ctx, tx, playerId)
    if err != nil {
        panic(exception.NewNotFoundError(fmt.Sprintf("Player is not found with id %d", playerId)))
    }

    return helper.ToPlayerResponse(player)
}

func (service *PlayerServiceImpl) GetAll(ctx context.Context) []webResponse.PlayerResponse {
    tx, err := service.DB.Begin()
    helper.PanicIfError(err)
    defer helper.CommitOrRollback(tx)

    players := service.PlayerRepository.GetAll(ctx, tx)
    
    var playerResponses []webResponse.PlayerResponse
    for _, player := range players {
        playerResponses = append(playerResponses, helper.ToPlayerResponse(player))
    }

    return playerResponses
}