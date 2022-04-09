package service

import (
	"blog_web/db/dao"
	"blog_web/model"
	"blog_web/utils"
)

type LinkService struct {
	linkDao *dao.LinkDao
}

func NewLinkService() *LinkService {
	return &LinkService{
		linkDao: dao.NewLinkDao(),
	}
}

func (l *LinkService) GetAllLinks() []model.Link {
	if links, err := l.linkDao.FindAllLinks(); err == nil {
		return links
	}
	return nil
}

func (l *LinkService) GetAllCategory() []model.LinkCategory {
	if categories, err := l.linkDao.FindAllLinkCategory(); err == nil {
		return categories
	}
	return nil
}

func (l *LinkService) GetLimitedLinks(pageNum, pageSize int) []model.Link {
	pageStart := (pageNum - 1) * pageSize
	links, err := l.linkDao.FindLimitedLinks(pageStart, pageSize)
	if err != nil {
		utils.Logger().Warning("get links error:%v", err)
		return nil
	}

	return links
}

func (l *LinkService) GetLinkCount() int {
	count, err := l.linkDao.FindLinkCount()
	if err != nil {
		utils.Logger().Warning("get link count error:%v", err)
	}

	return count
}

func (l *LinkService) AddLink(link *model.Link) error {
	return l.linkDao.AddLink(link)
}

func (l *LinkService) DeleteLink(id int) error {
	return l.linkDao.DeleteLink(id)
}

func (l *LinkService) UpdateLink(link *model.Link) error {
	return l.linkDao.UpdateLink(link)
}

func (l *LinkService) AddCategory(category *model.LinkCategory) error {
	return l.linkDao.AddCategory(category)
}

func (l *LinkService) DeleteCategory(id int) error {
	return l.linkDao.DeleteCategory(id)
}

func (l *LinkService) UpdateCategory(category *model.LinkCategory) error {
	return l.linkDao.UpdateCategory(category)
}