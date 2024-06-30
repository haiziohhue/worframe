package iface

type IDomainService[D IDao, E IEntity] interface {
	// GetOne 获取一个实体
	GetOne(id uint) (*E, error)
	// GetAll 获取筛选实体
	GetAll(page, pageSize int) ([]E, error)
	// Create 根据实体插入数据库
	Create(entity E) error
	// Update 更新相应 id 的实体
	Update(id uint, entity E) error
	// Delete 删除相应 id 的数据
	Delete(id uint) error
	// ToE 将 Dao 转化为 Entity
	ToE(D) (*E, error)
	// ToDao 将 Entity 转化为 Dao
	ToDao(E) (*D, error)
}
