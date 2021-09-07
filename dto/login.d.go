package dto

type LoginReq struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}

type LoginRes struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
