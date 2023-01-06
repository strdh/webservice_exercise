package repository

import (
    "fmt"
    "context"
    "errors"
    "database/sql"
    "exercise/webservice/model"
    "exercise/webservice/helper"
)

type UserRepositoryImpl struct {}

func NewUserRepository() UserRepository {
    return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, user model.User) model.User{
    sql := "INSERT INTO users (name, username, password) VALUES (?, ?, ?)"
    result, err := tx.ExecContext(ctx, sql, user.Name, user.Username, user.Password)
    helper.PanicIfError(err)

    id, err := result.LastInsertId()
    helper.PanicIfError(err)

    user.Id = int(id)
    return user
}

func (repository *UserRepositoryImpl) FindByUsername(ctx context.Context, tx *sql.Tx, username string) (model.User, error) {
    sql := "SELECT * FROM users WHERE username = ?"
    rows, err := tx.QueryContext(ctx, sql, username)
    helper.PanicIfError(err)
    defer rows.Close()

    user := model.User{}
    if rows.Next() {
        err := rows.Scan(&user.Id, &user.Name, &user.Username, &user.Password)
        helper.PanicIfError(err)
        return user, nil
    }

    return user, errors.New(fmt.Sprintf("User with username %s not found", username))
}