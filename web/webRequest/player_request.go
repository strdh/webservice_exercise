package webRequest 

type PlayerCreateRequest struct {
    Name string `json:"name" validate:"required,min=3,max=50"` 
    Position string `json:"position" validate:"required,min=1,max=50"`
    Height float32 `json:"height" validate:"required,gte=3,lte=8"`
    Weight float32 `json:"weight" validate:"required,gte=30,lte=300"`
    BirthDate string `json:"birth_date" validate:"required,datetime=2006-01-02"`
}

type PlayerUpdateRequest struct {
    Id int `json:"id"`
    Name string `json:"name" validate:"required,min=3,max=50"` 
    Position string `json:"position" validate:"required,min=1,max=50"`
    Height float32 `json:"height" validate:"required,gte=3,lte=8"`
    Weight float32 `json:"weight" validate:"required,gte=30,lte=300"`
    BirthDate string `json:"birth_date" validate:"required,datetime=2006-01-02"`
}