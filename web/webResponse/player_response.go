package webResponse

type PlayerResponse struct {
    Id string `json:"id"`
    Name string `json:"name"`
    Position string `json:"position"`
    Height float32 `json:"height"`
    Weight float32 `json:"weight"`
    BirthDate string `json:"birth_date"`
}