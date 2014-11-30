package models

type Creds struct {
	Username string `json:"username" val:"email"`
	Password string `json:"password" val:"password"`
}
