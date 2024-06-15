package types

// NormalListQuery get all所带参数
type NormalListQuery struct {
	Page     int `form:"page,default=1" json:"page"`
	PageSize int `form:"page_size,default=10" json:"page_size"`
}

// IdParam 单Id返回
type IdParam struct {
	Id uint `uri:"id" binding:"required"`
}

// EmptyRes 无数据返回结构
type EmptyRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// BaseRes 带数据返回结构
type BaseRes[T interface{}] struct {
	EmptyRes
	Data T `json:"data"`
}
type JwtPayload struct {
	UUID string   `json:"uuid"`
	Role []string `json:"role,omitempty"`
}
type LoginBody struct {
	Username   string `form:"username" json:"username" binding:"required"`
	Password   string `form:"password" json:"password" binding:"required"`
	Email      string `form:"email" json:"email" binding:"required"`
	Captcha    string `form:"captcha" json:"captcha" binding:"required"`
	VerifyCode string `form:"verify_code" json:"verify_code" binding:"required"`
}
