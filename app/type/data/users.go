package data

type LoginByPasswordReq struct {
	Uid      string `json:"uid"`
	Password string `json:"password"`
}

type LoginByPasswordRes struct {
	Token string `json:"token"`
}

type RegisterReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

type RegisterRes struct {
	IsOk bool `json:"is_ok"`
}
