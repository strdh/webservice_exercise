package helper

import (
    "encoding/json"
    "net/http"
)

func ReadFromRequestBody(request *http.Request, result interface{}) error {
    decoder := json.NewDecoder(request.Body) 
    err := decode.Decode(result)
    PanicIfError(err)
}

func WriteToResponseBody(writer http.ResponseWriter, response interface{}) {
    writer.Header().Set("Content-Type", "application/json")
    encoder := json.NewEncoder(writer)
    err := encoder.Encode(response)
    PanicIfError(err)
}