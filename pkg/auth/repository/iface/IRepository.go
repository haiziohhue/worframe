package iface

import (
	"gorm.io/gorm/schema"
)

type IRepository[T schema.Tabler] interface {
	Create(T) error
	Update(T) error
	Delete(T) error
	Find(page, pageSize int) ([]T, error)
	FindOne(T) (T, error)
}
