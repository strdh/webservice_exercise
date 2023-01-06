package helper 

import (
    "exercise/webservice/model"
    "exercise/webservice/web/webResponse"
)

func ToPlayerResponse(player model.Player) webResponse.PlayerResponse {
    return webResponse.PlayerResponse{
        Id: player.Id,
        Name: player.Name,
        Position: player.Position,
        Height: player.Height,
        Weight: player.Weight,
        BirthDate: player.BirthDate,
    }
}

func ToUserResponse(user model.User) webResponse.UserResponse {
    return webResponse.UserResponse{
        Id: user.Id,
        Name: user.Name,
        Username: user.Username,
        Password: user.Password,
    }
}