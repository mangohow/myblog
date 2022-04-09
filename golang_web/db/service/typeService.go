package service

import (
	"blog_web/db/dao"
	"blog_web/model"
	"blog_web/utils"
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
func (t *TypeService) FindAll() []model.TheType {
	types, err := t.typeDao.FindAll()
	if err != nil {
		return nil
	}

	return types
}

func (t *TypeService) GetOnePage(pageNum, pageSize int) ([]model.TheType, int) {
	pageStart := (pageNum - 1) * pageSize
	types, err := t.typeDao.FindOnePage(pageStart, pageSize)
	if err != nil {
		utils.Logger().Warning("GetOnePage error:%v", err)
		return nil, 0
	}
	count, _ := t.typeDao.GetTypesCount()
	return types, count
}

func (t *TypeService) CheckTypeExist(name string) bool {
	tp, err := t.typeDao.FindTypeByName(name)
	utils.Logger().Debug("typename:%s", tp)
	if err != nil || tp.Id <= 0 {
		return false
	}

	return true
}

func (t *TypeService) DeleteById(id int) bool {
	err := t.typeDao.DeleteById(id)
	if err != nil {
		utils.Logger().Warning("DeleteById error:%v", err)
		return false
	}

	return true
}

func (t *TypeService) UpdateName(id int, name string) bool {
	err := t.typeDao.UpdateName(id, name)
	if err != nil {
		utils.Logger().Warning("UpdateName error:%v", err)
		return false
	}

	return true
}

func (t *TypeService) AddType(name string) bool {
	err := t.typeDao.AddType(name)
	if err != nil {
		utils.Logger().Warning("AddType error:%v", err)
		return false
	}

	return true
}
