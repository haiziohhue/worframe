package factory

import (
	"errors"
	"worframe/share/factory/iface"
)

type BaseDomainService[D iface.IDao, E iface.IEntity] struct {
	repo iface.IDomainRepo[D]
}

func (ds *BaseDomainService[D, E]) ToE(Dao D) (*E, error) {
	entity, ok := Dao.ToEntity().(E)
	if !ok {
		return nil, errors.New("entity does not implement E")
	}
	return &entity, nil
}
func (ds *BaseDomainService[D, E]) ToDao(Entity E) (*D, error) {
	dao, ok := Entity.ToDao().(D)
	if !ok {
		return nil, errors.New("entity does not implement E")
	}
	return &dao, nil
}
func NewBaseDomainService[D iface.IDao, E iface.IEntity](repo iface.IDomainRepo[D]) iface.IDomainService[D, E] {
	return &BaseDomainService[D, E]{
		repo: repo,
	}
}
func (ds *BaseDomainService[D, E]) GetOne(id uint) (*E, error) {
	dao, err := ds.repo.GetOne(id)
	if err != nil || dao == nil {
		return nil, err
	}
	entity, err := ds.ToE(*dao)
	if err != nil {
		return nil, err
	}
	return entity, nil
}
func (ds *BaseDomainService[D, E]) GetAll(page, pageSize int) ([]E, error) {
	list, err := ds.repo.GetAll(page, pageSize)
	if err != nil {
		return nil, err
	}
	var entities []E
	for _, v := range list {
		entity, err := ds.ToE(v)
		if err != nil {
			return nil, err
		}
		entities = append(entities, *entity)
	}
	return entities, nil
}
func (ds *BaseDomainService[D, E]) Create(entity E) error {
	dao, err := ds.ToDao(entity)
	if err != nil {
		return err
	}
	return ds.repo.Create(*dao)
}
func (ds *BaseDomainService[D, E]) Update(id uint, entity E) error {
	dao, err := ds.ToDao(entity)
	if err != nil {
		return err
	}
	return ds.repo.Update(id, *dao)
}
func (ds *BaseDomainService[D, E]) Delete(id uint) error {
	return ds.repo.Delete(id)
}
