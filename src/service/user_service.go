package service

import (
	"lingye-gin/src/dao"
	"lingye-gin/src/entity"
	"lingye-gin/src/service/dto"
	"lingye-gin/src/service/query"
	"lingye-gin/src/util"
)

type UserService struct{}

func (UserService) Save(resource *dto.UserDTO) {
	if resource.ID != 0 {
		panic("id is not zero")
	}
	record := &entity.User{}
	util.StructCopy(resource, record)
	new(dao.UserDAO).Insert(record)
}

func (UserService) RemoveById(id uint) {
	if id == 0 {
		panic("id is zero")
	}
	new(dao.UserDAO).DeleteById(id)
}

func (UserService) ModifyById(resource *dto.UserDTO) {
	if resource.ID == 0 {
		panic("id is zero")
	}
	record := &entity.User{}
	util.StructCopy(resource, record)
	new(dao.UserDAO).UpdateById(record)
}

func (UserService) DescribeUsers(condition query.UserQuery) ([]entity.User, int) {
	return new(dao.UserDAO).SelectPage(condition)
}

func (UserService) DescribeUserById(id uint) entity.User {
	return new(dao.UserDAO).SelectOne(id)
}
