package domain

import "time"

type UserStruct struct {
	Id          string
	Username    string
	Email       string
	Phone       string
	DateOfBirth *time.Time
}

type UserDb interface {
	GetAllDB() ([]UserStruct, error)
	GetByIdDB(id string) (*UserStruct, error)
	InsertUserDB(user *UserStruct) (string, error)
	UpdateUserDB(user *UserStruct) (string, error)
	DeleteUserDB(id string) (string, error)
}
