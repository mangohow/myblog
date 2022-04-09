package service

import (
	"blog_web/db/dao"
	"blog_web/model"
	"blog_web/utils"
)

/*
* @Author: mgh
* @Date: 2022/2/28 19:08
* @Desc:
 */

type BlogService struct {
	blogDao *dao.BlogDao
	tagDao  *dao.TagDao
}

func NewBlogService() *BlogService {
	return &BlogService{
		blogDao: dao.NewBlogDao(),
		tagDao:  dao.NewTagDao(),
	}
}

// 获取主页的博客列表，根据pageNum和pageSize进行分页
func (b *BlogService) GetHomePageBlogs(pageNum, pageSize int) []model.BlogUserType {
	pageStart := (pageNum - 1) * pageSize
	blogs, err := b.blogDao.GetConciseBlogs(pageStart, pageSize)
	if err != nil {
		utils.Logger().Warning("GetHomePageBlogs error:%v", err)
		return nil
	}

	return blogs
}

// 获取所有博客的数量
func (b *BlogService) GetBolgCount() int {
	count, err := b.blogDao.GetBolgCount()
	if err != nil {
		return 0
	}

	return count
}

// 获取博客详情，也就是浏览博客
func (b *BlogService) GetDetailedBlog(id int) (*model.DetailedBlog, []model.Tag) {
	b.blogDao.UpdateViews(id)
	blog, err := b.blogDao.FindDetailedBlog(id)
	if err != nil {
		utils.Logger().Warning("FindDetailedBlog error:%v", err)
		return nil, nil
	}
	tags, _ := b.tagDao.GetTagsByBlogId(id)

	return blog, tags
}

// 获取所有的博客类型
func (b *BlogService) GetAllTypes() []int {
	typeIds, err := b.blogDao.FindAllTypes()
	if err != nil {
		return nil
	}
	return typeIds
}

// 根据博客类型ID获取博客列表
func (b *BlogService) GetBlogListByTypeId(typeid, pageNum, pageSize int) ([]model.BlogUserType, int) {
	pageStart := (pageNum - 1) * pageSize
	blogs, err := b.blogDao.FindByTypeId(typeid, pageStart, pageSize)
	if err != nil {
		utils.Logger().Warning("GetBlogListByTypeId error:%v", err)
		return nil, 0
	}

	count, err := b.blogDao.GetBolgCountByTypeId(typeid)
	if err != nil {
		return nil, 0
	}

	return blogs, count
}

// 根据标签ID获取博客列表
func (b *BlogService) GetBlogsByTagId(id, pageNum, pageSize int) ([]model.BlogUserType, int) {
	pageStart := (pageNum - 1) * pageSize
	blogs, err := b.blogDao.FindBlogsByTagId(id, pageStart, pageSize)
	if err != nil {
		utils.Logger().Warning("GetBlogsByTagId error:%v", err)
		return nil, 0
	}

	count, err := b.blogDao.GetBlogCountByTagId(id)
	if err != nil {
		utils.Logger().Warning("GetBlogsByTagId error:%v", err)
		return nil, 0
	}

	return blogs, count
}

// 获取时间线页面的博客列表
func (b *BlogService) GetTimeLineBlogs() [][]model.Blog {
	years, err := b.blogDao.GetAllBlogPublishYear()
	if err != nil {
		utils.Logger().Warning("GetTimeLineBlogs error:%v", err)
		return nil
	}

	groupBlogs := make([][]model.Blog, 0, 1)
	for _, year := range years {
		blogs, err := b.blogDao.GetBlogByFormatedYear(year)
		if err != nil {
			utils.Logger().Warning("GetTimeLineBlogs error:%v", err)
			return nil
		}
		groupBlogs = append(groupBlogs, blogs)
	}

	return groupBlogs
}

// 根据关键词查询博客列表
func (b *BlogService) GetBlogsByKeyWord(keyWord string) []model.BlogSection {
	keyWord = "%" + keyWord + "%"
	blogs, err := b.blogDao.FindBlogsByKeyWord(keyWord)
	if err != nil {
		utils.Logger().Warning("GetBlogsByKeyWord error:%v", err)
		return nil
	}

	return blogs
}

// 根据title、type、recommend来筛选博客
func (b *BlogService) GetBlogsByTitleOrTypeOrRecommend(pageNum, pageSize int, title string, typeId int, recommend string) ([]model.BlogUserType, int) {
	pageStart := (pageNum - 1) * pageSize
	blogs, count, err := b.blogDao.FindBlogsByTitleOrTypeOrRecommend(pageStart, pageSize, title, typeId, recommend)
	if err != nil {
		utils.Logger().Warning("GetBlogsByTitleOrTypeOrRecommend error:%v", err)
		return nil, 0
	}

	return blogs, count
}

func (b *BlogService) DeleteBlog(id int) bool {
	err := b.blogDao.DeleteById(id)
	if err != nil {
		utils.Logger().Warning("delete blog error:%v", err)
		return false
	}

	return true
}

func (b *BlogService) GetFullBlog(id int) (*model.BlogUserType, []model.Tag) {
	blog, err := b.blogDao.FindFullBlog(id)
	if err != nil {
		utils.Logger().Warning("GetFullBlog error:%v", err)
		return nil, nil
	}

	tags, err := b.tagDao.GetTagsByBlogId(id)
	if err != nil {
		utils.Logger().Warning("GetFullBlog error:%v", err)
		return nil, nil
	}

	return blog, tags
}

func (b *BlogService) UpdateBlog(blog *model.FullBlog) bool {
	tx, err := dao.Sqldb.Beginx()
	if err != nil {
		return false
	}

	defer func() {
		if e := recover(); e != nil {
			utils.Logger().Warning("UpdateBlog panic")
			tx.Rollback()
		}
	}()

	//  从t_blog_tags表中删除原来的数据
	err = b.blogDao.DeleteBlogTagsByBlogId(tx, blog.Id)
	if err != nil {
		utils.Logger().Warning("DeleteBlogTagsByBlogId error:%v", err)
		tx.Rollback()
		return false
	}

	// 更新t_blog表
	err = b.blogDao.UpdateBlog(tx, blog)
	if err != nil {
		utils.Logger().Warning("UpdateBlog error:%v", err)
		tx.Rollback()
		return false
	}

	// 更新t_blog_tags表
	blogTags := make([]model.BlogTag, len(blog.TagIds))
	for i, tid := range blog.TagIds {
		blogTags[i].TagId = tid
		blogTags[i].BlogId = blog.Id
	}
	err = b.blogDao.InsertIntoBlogTags(tx, blogTags)
	if err != nil {
		utils.Logger().Warning("InsertIntoBlogTags error:%v", err)
		tx.Rollback()
		return false
	}

	tx.Commit()
	return true

}

func (b *BlogService) AddBlog(blog *model.FullBlog) bool {
	tx, err := dao.Sqldb.Beginx()
	if err != nil {
		return false
	}

	defer func() {
		if e := recover(); e != nil {
			utils.Logger().Error("UpdateBlog panic")
			tx.Rollback()
		}
	}()

	// 更新t_blog表
	id, err := b.blogDao.AddBlog(tx, blog)
	if err != nil {
		utils.Logger().Error("AddBlog error:%v", err)
		tx.Rollback()
		return false
	}
	// 更新t_blog_tags表
	blogTags := make([]model.BlogTag, len(blog.TagIds))
	for i, tid := range blog.TagIds {
		blogTags[i].TagId = tid
		blogTags[i].BlogId = id
	}
	err = b.blogDao.InsertIntoBlogTags(tx, blogTags)
	if err != nil {
		utils.Logger().Warning("InsertIntoBlogTags error:%v", err)
		tx.Rollback()
		return false
	}

	tx.Commit()
	return true
}

// 根据创建时间获取最新的n篇博客
func (b *BlogService) GetNewBlogs(n int) []model.BlogSection {
	blogs, err := b.blogDao.FindBlogsByCreateTime(n)
	if err != nil {
		utils.Logger().Warning("%v", err)
		return nil
	}

	return blogs
}


// 根据浏览次数获取最热门的n篇博客
func (b *BlogService) GetHotBlogs(n int) []model.BlogSection {
	blogs, err := b.blogDao.FindBlogsByViews(n)
	if err != nil {
		utils.Logger().Warning("%v", err)
		return nil
	}

	return blogs
}

// 获取每个类型以及该类型下的博客数量
func (b *BlogService) GetTypeAndBlogCount() []model.TheType {
	types, err := b.blogDao.FindTypeAndBlogCount()
	if err != nil {
		utils.Logger().Warning("%v", err)
		return nil
	}

	return types
}