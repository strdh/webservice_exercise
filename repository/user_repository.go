package repository

import (
    "context"
    "database/sql"
    "exercise/webservice/model"
)

type UserRepository interface {
    Create(ctx context.Context, tx *sql.Tx, user model.User) model.User
    FindByUsername(ctx context.Context, tx *sql.Tx, username string) (model.User, error)
}