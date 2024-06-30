package iface

type IWebService[T IDto, E IEntity, D IDao] interface {
	// GetOne 获取单个数据
	GetOne(id uint) (*T, error)
	// GetAll 获取所有数据
	GetAll(page, pageSize int) ([]*T, error)
	// Create 保存数据
	Create(model T) error
	// Update 更新数据
	Update(id uint, model T) error
	// Delete 删除数据
	Delete(id uint) error
	// ToE 将 Dto 转化为 Entity
	ToE(T) (*E, error)
	// ToDto 将 Entity 转化为 Dto
	ToDto(E) (*T, error)
}
