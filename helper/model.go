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