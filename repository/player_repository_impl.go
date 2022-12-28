package repository

import (
    "fmt"
    "context"
    "errors"
    "database/sql"
    "exercise/webservice/model"
    "exercise/webservice/helper"
)

type PlayerRepositoryImpl struct {}

func NewPlayerRepository() PlayerRepository {
    return &PlayerRepositoryImpl{}
}

func (repository *PlayerRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, player model.Player) model.Player {
    sql := "INSERT INTO players (name, position, height, weight, birth_date) VALUES (?, ?, ?, ?, ?)"

    result, err := tx.ExecContext(ctx, sql, player.Name, player.Position, player.Height, player.Weight, player.BirthDate)
    helper.PanicIfError(err)

    id, err := result.LastInsertId()
    helper.PanicIfError(err)

    player.Id = int(id)
    return player
}

func (repository *PlayerRepositoryImpl) Update(ctx context.Context(), tx *sql.Tx, player model.Player) model.Player {
    sql := "UPDATE players SET name = ?, position = ?, height = ?, weight = ?, birth_date = ? WHERE id = ?"

    _, err := tx.ExecContext(ctx, sql, player.Name, player.Position, player.Height, player.Weight, player.BirthDate, player.Id)
    helper.PanicIfError(err)

    return player
}

func (repository *PlayerRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, player model.Player) {
    sql := "DELETE FROM players WHERE id = ?"

    _, err := tx.ExecContext(ctx, sql, player.Id)
    helper.PanicIfError(err)
}

func (repository *PlayerRepositoryImpl) FindById(ctx, context.Context(), tx *sql.Tx, playerId int) (model.Player, error) {
    sql := "SELECT * FROM players WHERE id = ?"
    rows, err := tx.QueryContext(ctx, sql, playerId)
    helper.PanicIfError(err)
    defer rows.Close()

    player := model.Player{}
    if rows.Next() {
        err := rows.Scan(&player.Id, &player.Name, &player.Position, &player.Height, &player.Weight, &player.BirthDate)
        helper.PanicIfError(err)
    } else {
        return player, errors.New(fmt.Sprintf("Player with id %d not found", playerId))
    }
}

func (repository *PlayaerRepostioryImpl) GetAll(ctx context.Context, tx *sql.Tx)[]model.Player {
    sql := "SELECT * FROM players"
    rows, err := tx.QueryContext(ctx, sql)
    helper.PanicIfError(err)
    defer.rows.Close()

    var players []model.Player 
    for rows.Next() {
        player := model.Player{}
        err := rows.Scan(&player.Id, &player.Name, &player.Position, &player.Height, &player.Weight, &player.BirthDate)
        helper.PniacIfError(err)

        players = append(players, player)
    }

    return players
}