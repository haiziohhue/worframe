package factory

import (
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"
	"worframe/share/factory/iface"
)

type BaseDomainRepo[D iface.IDao] struct {
	db   *gorm.DB
	pool *redis.Pool
}

func NewBaseDomainRepo[D iface.IDao](db *gorm.DB, pool *redis.Pool) iface.IDomainRepo[D] {
	return &BaseDomainRepo[D]{db, pool}
}
func (repo *BaseDomainRepo[D]) GetOne(id uint) (*D, error) {
	var dao *D
	err := repo.db.Where(id).First(&dao).Error
	if err != nil {
		return nil, err
	}
	return dao, nil
}
func (repo *BaseDomainRepo[D]) GetAll(page, pageSize int) ([]D, error) {
	var dao []D
	err := repo.db.Limit(pageSize).Offset((page - 1) * pageSize).Find(&dao).Error
	if err != nil {
		return nil, err
	}
	return dao, nil
}
func (repo *BaseDomainRepo[D]) Create(dao D) error {
	return repo.db.Create(dao).Error
}
func (repo *BaseDomainRepo[D]) Delete(id uint) error {
	var dao *D
	return repo.db.Model(dao).Delete(id).Error
}
func (repo *BaseDomainRepo[D]) Update(id uint, model D) error {
	model.SetGormM(gorm.Model{ID: id})
	return repo.db.Updates(model).Error
}
