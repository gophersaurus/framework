package models

// Creds describes a user's credentials.
type Creds struct {
	Username string `json:"username" val:"email"`
	Password string `json:"password" val:"password"`
}
