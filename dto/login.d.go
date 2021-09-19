package dto

type Jwt struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type JWTSource struct {
	Id                string `json:"id"`
	Name              string `json:"name"`
	School            string `json:"school"`
	Nickname          string `json:"nickname"`
	HashedAccessToken string `json:"hashedAccessToken"`
}

type LoginReq struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}

type LoginRes struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
