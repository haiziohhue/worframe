package types

type NormalListQuery struct {
	Page     int `form:"page,default=1" json:"page"`
	PageSize int `form:"page_size,default=10" json:"page_size"`
}
