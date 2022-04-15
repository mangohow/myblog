package service

import (
	"blog_web/db/dao"
	"blog_web/model"
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
func (b *BlogService) GetHomePageBlogs(pageNum, pageSize int) ([]model.BlogUserType, error) {
	pageStart := (pageNum - 1) * pageSize
	return b.blogDao.GetConciseBlogs(pageStart, pageSize)
}

// 获取所有博客的数量
func (b *BlogService) GetBolgCount() (int, error) {
	return b.blogDao.GetBolgCount()
}

// 获取博客详情，也就是浏览博客
func (b *BlogService) GetDetailedBlog(id int) (*model.DetailedBlog, []model.Tag, error) {
	b.blogDao.UpdateViews(id)
	blog, err := b.blogDao.FindDetailedBlog(id)
	if err != nil {
		return nil, nil, err
	}
	tags, _ := b.tagDao.GetTagsByBlogId(id)

	return blog, tags, nil
}

// 获取所有的博客类型
func (b *BlogService) GetAllTypes() ([]int, error) {
	return b.blogDao.FindAllTypes()
}

// 根据博客类型ID获取博客列表
func (b *BlogService) GetBlogListByTypeId(typeid, pageNum, pageSize int) ([]model.BlogUserType, int, error) {
	pageStart := (pageNum - 1) * pageSize
	blogs, err := b.blogDao.FindByTypeId(typeid, pageStart, pageSize)
	if err != nil {
		return nil, 0, err
	}

	count, err := b.blogDao.GetBolgCountByTypeId(typeid)
	if err != nil {
		return nil, 0, err
	}

	return blogs, count, nil
}

// 根据标签ID获取博客列表
func (b *BlogService) GetBlogsByTagId(id, pageNum, pageSize int) ([]model.BlogUserType, int, error) {
	pageStart := (pageNum - 1) * pageSize
	blogs, err := b.blogDao.FindBlogsByTagId(id, pageStart, pageSize)
	if err != nil {
		return nil, 0, err
	}

	count, err := b.blogDao.GetBlogCountByTagId(id)
	if err != nil {
		return nil, 0, err
	}

	return blogs, count, nil
}

// 获取时间线页面的博客列表
func (b *BlogService) GetTimeLineBlogs() ([][]model.Blog, error) {
	years, err := b.blogDao.GetAllBlogPublishYear()
	if err != nil {
		return nil, err
	}

	groupBlogs := make([][]model.Blog, 0, 1)
	for _, year := range years {
		blogs, err := b.blogDao.GetBlogByFormatedYear(year)
		if err != nil {
			return nil, err
		}
		groupBlogs = append(groupBlogs, blogs)
	}

	return groupBlogs, nil
}

// 根据关键词查询博客列表
func (b *BlogService) GetBlogsByKeyWord(keyWord string) ([]model.BlogSection, error) {
	keyWord = "%" + keyWord + "%"
	blogs, err := b.blogDao.FindBlogsByKeyWord(keyWord)
	if err != nil {
		return nil, err
	}

	return blogs, nil
}

// 根据title、type、recommend来筛选博客
func (b *BlogService) GetBlogsByTitleOrTypeOrRecommend(pageNum, pageSize int, title string, typeId int,
										recommend string) ([]model.BlogUserType, int, error) {
	pageStart := (pageNum - 1) * pageSize
	blogs, count, err := b.blogDao.FindBlogsByTitleOrTypeOrRecommend(pageStart, pageSize, title, typeId, recommend)
	if err != nil {
		return nil, 0, err
	}

	return blogs, count, nil
}

func (b *BlogService) DeleteBlog(id int) error {
	return b.blogDao.DeleteById(id)
}

func (b *BlogService) GetFullBlog(id int) (*model.BlogUserType, []model.Tag, error) {
	blog, err := b.blogDao.FindFullBlog(id)
	if err != nil {
		return nil, nil, err
	}

	tags, err := b.tagDao.GetTagsByBlogId(id)
	if err != nil {
		return nil, nil, err
	}

	return blog, tags, nil
}

func (b *BlogService) UpdateBlog(blog *model.FullBlog) error {
	tx, err := dao.Sqldb.Beginx()
	if err != nil {
		return err
	}

	defer func() {
		if e := recover(); e != nil {
			tx.Rollback()
		}
	}()

	//  从t_blog_tags表中删除原来的数据
	err = b.blogDao.DeleteBlogTagsByBlogId(tx, blog.Id)
	if err != nil {
		tx.Rollback()
		return err
	}

	// 更新t_blog表
	err = b.blogDao.UpdateBlog(tx, blog)
	if err != nil {
		tx.Rollback()
		return err
	}

	// 更新t_blog_tags表
	blogTags := make([]model.BlogTag, len(blog.TagIds))
	for i, tid := range blog.TagIds {
		blogTags[i].TagId = tid
		blogTags[i].BlogId = blog.Id
	}
	err = b.blogDao.InsertIntoBlogTags(tx, blogTags)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (b *BlogService) AddBlog(blog *model.FullBlog) error {
	tx, err := dao.Sqldb.Beginx()
	if err != nil {
		return err
	}

	defer func() {
		if e := recover(); e != nil {
			tx.Rollback()
		}
	}()

	// 更新t_blog表
	id, err := b.blogDao.AddBlog(tx, blog)
	if err != nil {
		tx.Rollback()
		return err
	}
	// 更新t_blog_tags表
	blogTags := make([]model.BlogTag, len(blog.TagIds))
	for i, tid := range blog.TagIds {
		blogTags[i].TagId = tid
		blogTags[i].BlogId = id
	}
	err = b.blogDao.InsertIntoBlogTags(tx, blogTags)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

// 根据创建时间获取最新的n篇博客
func (b *BlogService) GetNewBlogs(n int) ([]model.BlogSection, error) {
	return b.blogDao.FindBlogsByCreateTime(n)
}


// 根据浏览次数获取最热门的n篇博客
func (b *BlogService) GetHotBlogs(n int) ([]model.BlogSection, error) {
	return b.blogDao.FindBlogsByViews(n)
}

// 获取每个类型以及该类型下的博客数量
func (b *BlogService) GetTypeAndBlogCount() ([]model.TheType, error) {
	return b.blogDao.FindTypeAndBlogCount()
}