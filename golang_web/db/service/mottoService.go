package service

import (
	"blog_web/db/dao"
	"blog_web/model"
)

/*
* @Author: mgh
* @Date: 2022/3/19 11:41
* @Desc:
 */

type MottoService struct {
	mottoDao *dao.MottoDao
}

func NewMottoService() *MottoService {
	return &MottoService{
		mottoDao: dao.NewMottoDao(),
	}
}

func (m *MottoService) GetAllMotto() ([]model.Motto, error) {
	return m.mottoDao.FindAll()
}

func (m *MottoService) AddOne(motto *model.Motto) error {
	return m.mottoDao.AddOne(motto)
}

func (m *MottoService) DeleteOne(id uint32) error {
	return m.mottoDao.DeleteOne(id)
}

func (m *MottoService) UpdateOne(motto *model.Motto) error {
	return m.mottoDao.UpdateById(motto)
}

func (m *MottoService) GetAllMottoWithCreateTime() ([]model.Motto, error) {
	return m.mottoDao.FindAllWithCreateTime()
}


