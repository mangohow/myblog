package service

import (
	"blog_web/db/dao"
	"blog_web/model"
	"blog_web/utils"
)

/*
* @Author: mgh
* @Date: 2022/3/2 11:30
* @Desc:
 */

type UserService struct {
	userDao *dao.UserDao
	tagDao *dao.TagDao
}

func NewUserService() *UserService {
	return &UserService{
		userDao: dao.NewUserDao(),
		tagDao: dao.NewTagDao(),
	}
}

func (u *UserService) CheckUser(username, password string) *model.User {
	user, err := u.userDao.FindUserByUsernameAndPassword(username, password)
	if err != nil {
		utils.Logger().Warning("%v", err)
		return nil
	}

	return user
}

func (u *UserService) GetInfo() (*model.User, int) {
	info, err := u.userDao.FindInfo()
	if err != nil {
		utils.Logger().Warning("%v", err)
		return nil, 0
	}
	count, _ := u.tagDao.GetTagsCount()


	return info, count
}
