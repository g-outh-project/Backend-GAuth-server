package dto

import "time"

type Jwt struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type JWTSource struct {
	Uid               string    `json:"uid"`
	Id                string    `json:"id"`
	Email             string    `json:"email"`
	Name              string    `json:"name"`
	School            string    `json:"school"`
	Birth             string    `json:"birth"`
	Nickname          string    `json:"nickname"`
	HashedAccessToken string    `json:"hashedAccessToken"`
	CreatedAt         time.Time `json:"createdAt"`
}

type LoginReq struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}

type LoginRes struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
