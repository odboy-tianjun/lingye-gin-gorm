package service

import (
	"lingye-gin/src/dao"
	"lingye-gin/src/entity"
	"lingye-gin/src/service/dto"
	"lingye-gin/src/service/query"
	"lingye-gin/src/util"
)

type UserService struct {
	userDao dao.UserDAO
}

func (service UserService) Save(resource *dto.UserDTO) {
	if resource.ID != 0 {
		panic("id is not zero")
	}
	record := &entity.User{}
	util.StructCopy(resource, record)
	service.userDao.Insert(record)
}

func (service UserService) RemoveById(id uint) {
	if id == 0 {
		panic("id is zero")
	}
	service.userDao.DeleteById(id)
}

func (service UserService) ModifyById(resource *dto.UserDTO) {
	if resource.ID == 0 {
		panic("id is zero")
	}
	record := &entity.User{}
	util.StructCopy(resource, record)
	service.userDao.UpdateById(record)
}

func (service UserService) DescribeUsers(condition query.UserQuery) ([]entity.User, int) {
	return service.userDao.SelectPage(condition)
}

func (service UserService) DescribeUserById(id uint) entity.User {
	return service.userDao.SelectOne(id)
}
