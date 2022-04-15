package service

import (
	"blog_web/db/dao"
	"blog_web/model"
)

/*
* @Author: mgh
* @Date: 2022/3/1 11:52
* @Desc:
 */

type TagService struct {
	tagDao *dao.TagDao
}

func NewTagService() *TagService {
	return &TagService{
		tagDao: dao.NewTagDao(),
	}
}

// 获取所有标签
func (t *TagService) GetAllTags() ([]model.Tag, error) {
	return t.tagDao.FindAllTags()
}

func (t *TagService) GetOnePageTags(pageNum, pageSize int) ([]model.Tag, int, error) {
	pageStart := (pageNum - 1) * pageSize
	tags, err := t.tagDao.GetOnePageTags(pageStart, pageSize)
	if err != nil {
		return nil, 0, err
	}
	count, _ := t.tagDao.GetTagsCount()
	return tags, count, nil
}

func (t *TagService) CheckTagExist(name string) bool {
	tag, err := t.tagDao.FindByName(name)
	if err != nil || tag.Id <= 0 {
		return false
	}

	return true
}

func (t *TagService) DeleteTagById(id int) error {
	return t.tagDao.DeleteById(id)
}

func (t *TagService) UpdateTagName(id int, name string) error {
	return t.tagDao.UpdateName(id, name)
}

func (t *TagService) AddTag(name string) error {
	return t.tagDao.AddTag(name)
}
