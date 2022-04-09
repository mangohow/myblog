package service

import (
	"blog_web/db/dao"
	"blog_web/model"
	"blog_web/utils"
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
func (t *TagService) GetAllTags() []model.Tag {
	tags, err := t.tagDao.FindAllTags()
	if err != nil {
		return nil
	}

	return tags
}

func (t *TagService) GetOnePageTags(pageNum, pageSize int) ([]model.Tag, int) {
	pageStart := (pageNum - 1) * pageSize
	tags, err := t.tagDao.GetOnePageTags(pageStart, pageSize)
	if err != nil {
		utils.Logger().Warning("GetOnePageTags error:%v", err)
		return nil, 0
	}
	count, _ := t.tagDao.GetTagsCount()
	return tags, count
}

func (t *TagService) CheckTagExist(name string) bool {
	tag, err := t.tagDao.FindByName(name)
	if err != nil || tag.Id <= 0 {
		return false
	}

	return true
}

func (t *TagService) DeleteTagById(id int) bool {
	err := t.tagDao.DeleteById(id)
	if err != nil {
		utils.Logger().Warning("DeleteTagById error:%v", err)
		return false
	}

	return true
}

func (t *TagService) UpdateTagName(id int, name string) bool {
	err := t.tagDao.UpdateName(id, name)
	if err != nil {
		utils.Logger().Warning("UpdateTagName error:%v", err)
		return false
	}

	return true
}

func (t *TagService) AddTag(name string) bool {
	err := t.tagDao.AddTag(name)
	if err != nil {
		utils.Logger().Warning("AddTag error:%v", err)
		return false
	}

	return true
}
