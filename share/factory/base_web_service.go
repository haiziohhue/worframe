package factory

import (
	"errors"
	"worframe/share/factory/iface"
)

type BaseWebService[T iface.IDto, E iface.IEntity, D iface.IDao] struct {
	ds iface.IDomainService[D, E]
}

func (ws BaseWebService[T, E, D]) GetOne(id uint) (*T, error) {
	entity, err := ws.ds.GetOne(id)
	if err != nil {
		return nil, err
	}
	dto, err := ws.ToDto(*entity)
	if err != nil {
		return nil, err
	}
	return dto, nil
}

func (ws BaseWebService[T, E, D]) GetAll(page, pageSize int) ([]*T, error) {
	list, err := ws.ds.GetAll(page, pageSize)
	if err != nil {
		return nil, err
	}
	var dtos []*T
	for _, entity := range list {
		entityDto, err := ws.ToDto(entity)
		if err != nil {
			return nil, err
		}
		dtos = append(dtos, entityDto)
	}
	return dtos, nil
}

func (ws BaseWebService[T, E, D]) Create(model T) error {
	modelEntity, err := ws.ToE(model)
	if err != nil {
		return err
	}
	return ws.ds.Create(*modelEntity)
}

func (ws BaseWebService[T, E, D]) Update(id uint, model T) error {
	modelEntity, err := ws.ToE(model)
	if err != nil {
		return err
	}
	return ws.ds.Update(id, *modelEntity)
}

func (ws BaseWebService[T, E, D]) Delete(id uint) error {
	return ws.ds.Delete(id)
}

func (ws BaseWebService[T, E, D]) ToE(t T) (*E, error) {
	entity, ok := t.ToEntity().(E)
	if !ok {
		return nil, errors.New("entity is not an E")
	}
	return &entity, nil
}

func (ws BaseWebService[T, E, D]) ToDto(e E) (*T, error) {
	dto, ok := e.ToDto().(T)
	if !ok {
		return nil, errors.New("entity is not an E")
	}
	return &dto, nil
}

func NewBaseWebService[T iface.IDto, E iface.IEntity, D iface.IDao](service iface.IDomainService[D, E]) iface.IWebService[T, E, D] {
	return BaseWebService[T, E, D]{
		service,
	}
}
