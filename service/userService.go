package service

import "mongo_rest_api_users/domain"

type UserServiceStruct struct {
	repo domain.UserDb
}

type UserRepository interface {
	GetAll() ([]domain.UserStruct, error)
	GetById(id string) (*domain.UserStruct, error)
	Insert(user *domain.UserStruct) (string, error)
	Update(user *domain.UserStruct) (string, error)
	Delete(id string) (string, error)
}

func (r UserServiceStruct) GetAll() ([]domain.UserStruct, error) {
	return r.repo.GetAllDB()
}

func (r UserServiceStruct) GetById(id string) (*domain.UserStruct, error) {
	return r.repo.GetByIdDB(id)
}

func (r UserServiceStruct) Insert(user *domain.UserStruct) (string, error) {
	return r.repo.InsertUserDB(user)
}

func (r UserServiceStruct) Update(user *domain.UserStruct) (string, error) {
	return r.repo.UpdateUserDB(user)
}

func (r UserServiceStruct) Delete(id string) (string, error) {
	return r.repo.DeleteUserDB(id)
}

func NewUserRepsitory(repository domain.UserDb) UserServiceStruct {
	return UserServiceStruct{repo: repository}
}
