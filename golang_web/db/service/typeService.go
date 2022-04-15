package service

import (
	"blog_web/db/dao"
	"blog_web/model"
)

/*
* @Author: mgh
* @Date: 2022/3/1 11:09
* @Desc:
 */

type TypeService struct {
	typeDao *dao.TypeDao
}

func NewTypeService() *TypeService {
	return &TypeService{
		typeDao: dao.NewTypeDao(),
	}
}

// 获取所有博客类型
func (t *TypeService) FindAll() ([]model.TheType, error) {
	return t.typeDao.FindAll()
}

func (t *TypeService) GetOnePage(pageNum, pageSize int) ([]model.TheType, int, error) {
	pageStart := (pageNum - 1) * pageSize
	types, err := t.typeDao.FindOnePage(pageStart, pageSize)
	if err != nil {
		return nil, 0, err
	}
	count, _ := t.typeDao.GetTypesCount()

	return types, count, nil
}

func (t *TypeService) CheckTypeExist(name string) bool {
	tp, err := t.typeDao.FindTypeByName(name)
	if err != nil || tp.Id <= 0 {
		return false
	}

	return true
}

func (t *TypeService) DeleteById(id int) error {
	return t.typeDao.DeleteById(id)
}

func (t *TypeService) UpdateName(id int, name string) error {
	return t.typeDao.UpdateName(id, name)
}

func (t *TypeService) AddType(name string) error {
	return t.typeDao.AddType(name)
}
