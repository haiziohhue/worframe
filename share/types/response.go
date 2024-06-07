package types

type BaseRes[T interface{}] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

type EmptyRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
