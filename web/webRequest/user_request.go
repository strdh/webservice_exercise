package webRequest 

type UserCreateRequest struct {
    Name string `json:"name" validate:"required,min=3,max=50"` 
    Username string `json:"username" validate:"required,min=6,max=100"`
    Password string `json:"password" validate:"required,min=8,max=100"`
}

type UserLoginRequest struct {
    Username string `json:"username" validate:"required,min=6,max=100"`
    Password string `json:"password" validate:"required,min=8,max=100"`
}