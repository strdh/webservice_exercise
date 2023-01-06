package helper

import (
    "time"
    "github.com/golang-jwt/jwt"
)

type Claims struct {
    Id int `json:"id"`
    Username string `json:"username"`
    Password string `json:"password"`
    jwt.StandardClaims
}

const KEY string = "what-zhit-tooya"

func GenerateToken(id int, username string, password string) (string, error) {
    claims := Claims{
        id,
        username,
        password,
        jwt.StandardClaims{
            ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
            Issuer: "webservice",
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    result, err := token.SignedString([]byte(KEY))
    PanicIfError(err)

    return result, err
}

func VerifyJWT(jwtToken string) bool {
    token, err := jwt.ParseWithClaims(jwtToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte(KEY), nil
    })

    if err != nil || !token.Valid {
        return false
    }

    return true
}