package middleware

import (
    "net/http"
    "github.com/julienschmidt/httprouter"
    "exercise/webservice/helper"
    "exercise/webservice/web/webResponse"
)

func AuthMiddlewareGroup(next httprouter.Handle) httprouter.Handle {
    return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
        userToken := r.Header.Get("JWT-TOKEN")
        result := helper.VerifyJWT(userToken)

        if !result || userToken == ""{
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusUnauthorized)

            response := webResponse.WebResponse{
                Code:   http.StatusUnauthorized,
                Status: "UNAUTHORIZED",
                Data:   "You dont have access to this service",
            }

            helper.WriteToResponseBody(w, response)
            return
        }

        next(w, r, ps) // call original
    }
}