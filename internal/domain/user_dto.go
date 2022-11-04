package domain

import "time"

type CreateUserDTO struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type SetRefreshTokenDTO struct {
	UserId                uint      `json:"user_id"`
	RefreshToken          string    `json:"refresh_token"`
	RefreshTokenExpiredAt time.Time `json:"refresh_token_expired_at"`
}
