package service 

import (
    "context"
    // "fmt"
    "database/sql"
    "encoding/base64"
     _ "github.com/go-sql-driver/mysql"
    "exercise/webservice/app/config"
    "exercise/webservice/web/webRequest"
    "exercise/webservice/web/webResponse"
    "exercise/webservice/model"
    "exercise/webservice/helper"
    "exercise/webservice/repository"
    "exercise/webservice/exception"
    "github.com/go-playground/validator/v10"
    "golang.org/x/crypto/bcrypt"    
)

type UserServiceImpl struct {
    UserRepository repository.UserRepository
    DB *sql.DB 
    Validate *validator.Validate
}

func NewUserService() UserService {
    userRepository := repository.NewUserRepository()
    var db *sql.DB = app.NewDB()
    var validate *validator.Validate = validator.New()
    return &UserServiceImpl{
        UserRepository: userRepository,
        DB: db,
        Validate: validate,
    }
}

func hashPassword(password string) string {
    salt, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    helper.PanicIfError(err)

    return base64.StdEncoding.EncodeToString(salt)
}

func verifyPassword(hashedPassword string, plainPassword string) (bool, error) {
    hashedPasswordBytes, err := base64.StdEncoding.DecodeString(hashedPassword)
    helper.PanicIfError(err)

    err = bcrypt.CompareHashAndPassword(hashedPasswordBytes, []byte(plainPassword))
    if err != nil {
        return false, err
    }

    return true, nil
}

func (service *UserServiceImpl) Register(ctx context.Context, request webRequest.UserCreateRequest) webResponse.UserResponse {
    err := service.Validate.Struct(request)
    helper.PanicIfError(err)

    tx, err := service.DB.Begin()
    helper.PanicIfError(err)
    defer helper.CommitOrRollback(tx)

    user := model.User{
        Name: request.Name,
        Username: request.Username,
        Password: hashPassword(request.Password),
    }

    user = service.UserRepository.Create(ctx, tx, user)
    return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) Login(ctx context.Context, username string, password string) string {
    tx, err := service.DB.Begin()
    helper.PanicIfError(err)
    defer helper.CommitOrRollback(tx)

    user, err := service.UserRepository.FindByUsername(ctx, tx, username)
    if err != nil {
        panic(exception.NewNotFoundError(err.Error()))
    }

    result, err := verifyPassword(user.Password, password)
    if err != nil || !result {
        panic(exception.NewNotFoundError(err.Error()))
    }

    token, err := helper.GenerateToken(user.Id, user.Username, user.Password)
    helper.PanicIfError(err)

    return token
}