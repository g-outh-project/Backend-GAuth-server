package dto

type SignupReq struct {
	Id       string `json:"id"`
	Password string `json:"password"`
	Email    string `json:"email"`
	School   string `json:"school"`
	Birth    string `json:"birth"`
	Nickname string `json:"nickname"`
	Name     string `json:"name"`
}
