package auth

type Auth struct {
	Token     string `json:"token"`
	UserName  string `json:"user_name"`
	ExpiresAt int64  `json:"expires_at"`
}
