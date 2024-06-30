package iface

import (
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type IEntity interface {
	ToDao() IDao
	ToDto() IDto
}
type IDao interface {
	schema.Tabler
	GetGormM() gorm.Model
	SetGormM(gorm.Model)
	ToEntity() IEntity
}
type IDto interface {
	ToEntity() IEntity
}
