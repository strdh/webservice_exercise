package service

import (
    "context"
    "exercise/webservice/web/webRequest"
    "exercise/webservice/web/webResponse"
)

type PlayerService interface {
    Create(ctx context.Context(), request webRequest.PlayerCreateRequest) webResponse.PlayerResponse 
    Update(ctx context.Context(), request.webRequest.PlayerUpdateRequest) webResponse.PlayerResponse
    Delete(ctx context.Context(), playerId int)
    FindById(ctx context.Context(), playerId int) webResponse.PlayerResponse
    GetAll(ctx context.Context())[]webResponse.PlayerResponse
}