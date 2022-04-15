package service

import (
	"blog_web/db/dao"
	"blog_web/model"
)

type LinkService struct {
	linkDao *dao.LinkDao
}

func NewLinkService() *LinkService {
	return &LinkService{
		linkDao: dao.NewLinkDao(),
	}
}

func (l *LinkService) GetAllLinks() ([]model.Link, error) {
	return l.linkDao.FindAllLinks()
}

func (l *LinkService) GetAllCategory() ([]model.LinkCategory, error) {
	return l.linkDao.FindAllLinkCategory()
}

func (l *LinkService) GetLimitedLinks(pageNum, pageSize int) ([]model.Link, error) {
	pageStart := (pageNum - 1) * pageSize
	return l.linkDao.FindLimitedLinks(pageStart, pageSize)
}

func (l *LinkService) GetLinkCount() (int, error) {
	return l.linkDao.FindLinkCount()
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