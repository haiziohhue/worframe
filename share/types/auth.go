package types

type LoginReqBody struct {
	username string `json:"username"`
	password string `json:"password"`
}
type LoginRespBody struct {
	token string
}
