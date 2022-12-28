package repository

import (
    "context"
    "database/sql"
    "exercise/webservice/model"
)

type PlayerRepository interface {
    Create(ctx context.Context, tx *sql.Tx, player model.Player) model.Player
    Update(ctx context.Context, tx *sql.Tx, player model.Player) model.Player
    Delete(ctx context.Context, tx *sql.Tx, player model.Player)
    FindById(ctx context.Context, tx *sql.Tx, playerId int) (model.Player, error)
    GetAll(ctx context.Context, tx *sql.Tx) []model.Player
}