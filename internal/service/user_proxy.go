package service

import (
	"asocks-ws/internal/domain"
	"asocks-ws/internal/repository"
)

type UserProxyService struct {
	repository repository.UserProxy
}

type UserProxyServiceInterface interface {
	FindById(int) (domain.UserProxy, error)
	UpdateExtIp(message domain.MessageUserProxy, input domain.UpdateUserProxy) (domain.UserProxy, error)
}

func NewUserProxyService(repository repository.UserProxy) *UserProxyService {
	return &UserProxyService{
		repository: repository,
	}
}

func (u *UserProxyService) FindById(id int) (domain.UserProxy, error) {
	return u.repository.FindByColumn(id, "id")
}

func (u *UserProxyService) LoadAll() ([]domain.UserProxy, error) {
	return u.repository.GetAll()
}

func (u *UserProxyService) UpdateExtIp(message domain.MessageUserProxy, input domain.UpdateUserProxy) (domain.UserProxy, error) {
	return u.repository.UpdateOne(message, input)
}
