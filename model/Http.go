package model

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResponseJSON struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
