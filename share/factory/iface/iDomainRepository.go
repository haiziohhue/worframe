package iface

type IDomainRepo[T IDao] interface {
	GetOne(id uint) (*T, error)
	GetAll(page, pageSize int) ([]T, error)
	Update(id uint, model T) error
	Delete(id uint) error
	Create(model T) error
}
