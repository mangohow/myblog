package service

import (
	"blog_web/db/dao"
	"blog_web/model"
	"blog_web/utils"
)

type EssayService struct {
	essayDao *dao.EssayDao
}

func NewEssayService() *EssayService {
	return &EssayService{
		essayDao: dao.NewEssayDao(),
	}
}

func (e *EssayService) GetAll() []model.Essay {
	essays, err := e.essayDao.FindAll()
	if err != nil {
		utils.Logger().Warning("Get essays error:%v", err)
		return nil
	}
	return essays
}

func (e *EssayService) GetLimited(pageNum, pageSize int) []model.Essay {
	pageStart := (pageNum - 1) * pageSize
	essays, err := e.essayDao.FindLimited(pageStart, pageSize)
	if err != nil {
		utils.Logger().Warning("get essays error:%v", err)
		return nil
	}

	return essays
}

func (e *EssayService) GetCount() int {
	count, err := e.essayDao.FindTotalCount()
	if err != nil {
		utils.Logger().Warning("get essay count error:%v", err)
		return 0
	}

	return count
}

func (e *EssayService) AddEssay(essay *model.Essay) error {
	return e.essayDao.Add(essay)
}

func (e *EssayService) DeleteEssay(id int) error {
	return e.essayDao.Delete(id)
}

func (e *EssayService) UpdateEssay(essay *model.Essay) error {
	return e.essayDao.Update(essay)
}