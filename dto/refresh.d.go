package dto

type RefreshReq struct {
	RefreshToken string `json:"refreshToken"`
}

type RefreshRes struct {
	AccessToken string `json:"accessToken"`
}
