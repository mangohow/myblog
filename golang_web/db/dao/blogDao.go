package dao

import (
	"blog_web/model"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strconv"
)

/*
* @Author: mgh
* @Date: 2022/2/28 19:08
* @Desc: 操作t_blog表的dao
 */

type BlogDao struct {
	sql []string
}

func NewBlogDao() *BlogDao {
	return &BlogDao{
		sql: []string{
			// 0、查询博客列表
			`SELECT b.id, b.title, b.description, b.first_picture, b.create_time,
			b.views, u.nickname nickname, t.name typename
			FROM t_blog b, t_user u, t_type t
			WHERE b.user_id = u.id and b.type_id = t.id and b.published = 1
			ORDER BY b.create_time DESC
			LIMIT ?, ?`,

			// 1、查询博客数量
			"SELECT COUNT(*) FROM t_blog;",

			// 2、更新浏览次数
			"UPDATE t_blog SET views = views + 1 WHERE id = ?;",

			// 3、查询博客详情
			`SELECT b.id, b.title, b.create_time, b.update_time, b.flag, b.appreciation, b.views, b.content, u.nickname nickname, t.name typename
			FROM t_blog b JOIN t_user u ON b.user_id = u.id 
			JOIN t_type t ON b.type_id = t.id 
			WHERE b.id = ?;`,

			// 4、查询所有的type
			`SELECT type_id FROM t_blog;`,

			// 5、根据博客类型id查询博客列表
			`SELECT b.id, b.title, b.description, b.first_picture, b.create_time,
			b.views, u.nickname nickname, t.name typename
        	FROM t_blog b, t_user u, t_type t
        	WHERE b.user_id = u.id AND b.type_id = t.id AND b.published = 1 AND b.type_id = ?
			ORDER BY b.create_time DESC
			LIMIT ?, ?`,

			// 6、查询一个博客类型下的博客数量
			`SELECT COUNT(*) FROM t_blog b, t_user u, t_type t WHERE 
			b.user_id = u.id AND b.type_id = t.id AND b.published = 1 AND b.type_id = ?;`,

			// 7、根据博客标签id查询博客列表
			`SELECT b.id, b.title, b.description, b.first_picture, b.create_time,
        	b.views, u.nickname nickname, t.name typename
        	FROM t_blog b, t_user u, t_type t
        	WHERE b.user_id = u.id AND b.type_id = t.id AND b.published = 1 AND b.id in (
            SELECT blog_id from t_blog_tags WHERE tag_id = ?)
			ORDER BY b.create_time DESC 
			LIMIT ?, ?;`,

			// 8、查询一个博客标签下的博客数量
			`SELECT COUNT(*) FROM t_blog b, t_user u, t_type t WHERE
			b.user_id = u.id AND b.type_id = t.id AND b.published = 1 AND b.id in (
            SELECT blog_id FROM t_blog_tags WHERE tag_id = ?)`,

			// 9、查询时间线
			`SELECT DISTINCT DATE_FORMAT(t_blog.create_time, '%Y') ut FROM t_blog ORDER BY ut DESC;`,

			// 10、根据某年查询该年发布的博客
			`SELECT b.title, b.create_time, b.id, b.flag
        	FROM t_blog b
        	WHERE DATE_FORMAT(b.create_time, "%Y") = ? ORDER BY b.create_time DESC;`,

			// 11、根据博客标题或博客内容搜索关键字
			`SELECT id, title
			FROM t_blog WHERE title LIKE ? OR content LIKE ?
			ORDER BY create_time DESC;`,

			// 12、获取搜索的博客的总条数
			`SELECT COUNT(*)
			FROM t_blog
			WHERE title LIKE ? OR content LIKE ?;`,

			// 13、根据博客标题、博客类型、是否推荐来查询博客信息，包括博客类型名
			// 在方法中拼接sql
			`SELECT b.id, b.title, b.create_time, b.recommend, b.published, b.type_id, t.name typename 
			FROM t_blog b ,t_type t 
			WHERE  b.type_id = t.id`,

			// 13、根据博客标题、博客类型、是否推荐来查询博客条数
			`SELECT COUNT(*) FROM t_blog b, t_type t WHERE  b.type_id = t.id `,

			// 15、根据博客id删除博客
			`DELETE FROM t_blog WHERE id = ?;`,

			// 16、根据博客id查询博客详情，管理界面修改博客使用
			`SELECT b.id, b.first_picture, b.flag, b.title, b.content, b.views,
        	b.update_time, b.commentabled, b.recommend, b.share_statement, b.appreciation, b.description,
        	t.name typename, t.id type_id
        	FROM t_blog b, t_type t
        	WHERE b.type_id = t.id AND  b.id = ?;`,

			// 17、根据博客id更新博客内容
			`UPDATE t_blog SET title=?, content=?, first_picture=?, flag=?, appreciation=?, share_statement=?,
			commentabled=?, published=?, recommend=?, update_time=?, type_id=?, user_id=?, description=? WHERE id=?`,

			// 18、新增博客
			`INSERT INTO t_blog(title, content, first_picture, flag, views,
        	appreciation, share_statement, commentabled, published, recommend,
        	create_time, update_time, type_id, user_id, description) VALUES(?, ?,
        	?, ?, ?, ?, ?,?, ?, ?, ?, ?, ?,?, ?)`,

			// 19、添加博客标签
			`INSERT INTO t_blog_tags (blog_id, tag_id) VALUES (:blog_id, :tag_id)`,

			// 20、删除博客
			`DELETE FROM t_blog_tags WHERE blog_id = ?;`,

			// 21、根据创建时间获取最新推荐的n篇博客
			`SELECT id, title FROM t_blog WHERE recommend = 1 ORDER BY create_time DESC LIMIT ?;`,

			// 22、根据浏览次数获取最热门的n篇博客
			`SELECT id, title FROM t_blog ORDER BY views DESC LIMIT ?;`,

			// 23、根据博客类型对博客分组、计算每个分组的数量
			`SELECT t.id, count(*) 'count',  t.name FROM t_blog b JOIN t_type t ON b.type_id = t.id GROUP BY type_id;`,
		},
	}
}

// 0、获取精简的博客信息，不包括博客内容
func (b *BlogDao) GetConciseBlogs(pageStart, pageSize int) (blogs []model.BlogUserType, err error) {
	err = sqldb.Select(&blogs, b.sql[0], pageStart, pageSize)
	return
}

// 1、获取博客数量
func (b *BlogDao) GetBolgCount() (count int, err error) {
	err = sqldb.Get(&count, b.sql[1])
	return
}

// 2、更新浏览次数
func (b *BlogDao) UpdateViews(id int) error {
	_, err := sqldb.Exec(b.sql[2], id)
	return err
}

// 3、查询详细的博客
func (b *BlogDao) FindDetailedBlog(id int) (*model.DetailedBlog, error) {
	var blog model.DetailedBlog
	err := sqldb.Get(&blog, b.sql[3], id)
	return &blog, err
}

// 4、查询所有的types
func (b *BlogDao) FindAllTypes() (typeIds []int, err error) {
	err = sqldb.Select(&typeIds, b.sql[4])
	return
}

// 5、根据类型id查询博客
func (b *BlogDao) FindByTypeId(id, pageStart, pageSize int) (blogs []model.BlogUserType, err error) {
	err = sqldb.Select(&blogs, b.sql[5], id, pageStart, pageSize)
	return
}

// 6、根据类型id获取博客
func (b *BlogDao) GetBolgCountByTypeId(id int) (count int, err error) {
	err = sqldb.Get(&count, b.sql[6], id)
	return
}

// 7、根据标签id获取博客
func (b *BlogDao) FindBlogsByTagId(id, pageStart, pageSize int) (blogs []model.BlogUserType, err error) {
	err = sqldb.Select(&blogs, b.sql[7], id, pageStart, pageSize)
	return
}

// 8、根据标签id获取博客数量
func (b *BlogDao) GetBlogCountByTagId(id int) (count int, err error) {
	err = sqldb.Get(&count, b.sql[8], id)
	return
}

// 9、获取去重后的博客发布时间
func (b *BlogDao) GetAllBlogPublishYear() (years []string, err error) {
	err = sqldb.Select(&years, b.sql[9])
	return
}

// 10、根据%Y类型的时间获取博客
func (b *BlogDao) GetBlogByFormatedYear(year string) (blogs []model.Blog, err error) {
	err = sqldb.Select(&blogs, b.sql[10], year)
	return
}

// 11、根据关键词查询博客
func (b *BlogDao) FindBlogsByKeyWord(keyWord string) (blogs []model.BlogSection, err error) {
	err = sqldb.Select(&blogs, b.sql[11], keyWord, keyWord)
	return
}

// 12、根据关键词查询博客数量
func (b *BlogDao) GetBlogCountByKeyWord(keyWord string) (count int, err error) {
	err = sqldb.Get(&count, b.sql[12], keyWord, keyWord)
	return
}

// 13、根据三个关键字的组合来获取博客，分别为博客标题、博客类型以及是否推荐
func (b *BlogDao) FindBlogsByTitleOrTypeOrRecommend(pageStart, pageSize int, title string, typeId int, recommend string) (blogs []model.BlogUserType, count int, err error) {
	sq1 := b.sql[13]
	sq2 := b.sql[14]
	param := make([]interface{}, 0, 5)
	if title != "" {
		title = "%" + title + "%"
		sq1 = fmt.Sprintf("%s %s", sq1, "AND b.title like ?")
		sq2 = fmt.Sprintf("%s %s", sq2, "AND b.title like ?")
		param = append(param, title)
	}
	if typeId > 0 {
		sq1 = fmt.Sprintf("%s %s", sq1, "AND b.type_id = ?")
		sq2 = fmt.Sprintf("%s %s", sq2, "AND b.type_id = ?")
		param = append(param, typeId)
	}
	if recommend != "" {
		sq1 = fmt.Sprintf("%s %s", sq1, "AND b.recommend = ?")
		sq2 = fmt.Sprintf("%s %s", sq2, "AND b.recommend = ?")
		rec, err := strconv.ParseBool(recommend)
		if err != nil {
			return nil, 0, err
		}

		param = append(param, rec)
	}

	sq1 = fmt.Sprintf("%s %s;", sq1, "ORDER BY b.create_time DESC LIMIT ?, ?")
	param = append(param, pageStart)
	param = append(param, pageSize)

	err = sqldb.Select(&blogs, sq1, param...)
	if err != nil {
		return nil, 0, err
	}

	par := param[:len(param)-2]
	err = sqldb.Get(&count, sq2, par...)
	if err != nil {
		return nil, 0, err
	}

	return
}

// 15、根据id删除博客
func (b *BlogDao) DeleteById(id int) error {
	_, err := sqldb.Exec(b.sql[15], id)
	return err
}

// 16、博客编辑页获取博客完整信息
func (b *BlogDao) FindFullBlog(id int) (*model.BlogUserType, error) {
	var blog model.BlogUserType
	err := sqldb.Get(&blog, b.sql[16], id)
	return &blog, err
}

// 17、更新博客
func (b *BlogDao) UpdateBlog(tx *sqlx.Tx, blog *model.FullBlog) error {
	_, err := tx.Exec(b.sql[17], blog.Title, blog.Content, blog.FirstPicture, blog.Flag, blog.Appreciation, blog.ShareStatement,
		blog.Commentabled, blog.Published, blog.Recommend, blog.UpdateTime, blog.TypeId, blog.UserId, blog.Description, blog.Id)
	return err
}

// 18、新增博客
func (b *BlogDao) AddBlog(tx *sqlx.Tx, blog *model.FullBlog) (int, error) {
	result, err := tx.Exec(b.sql[18], blog.Title, blog.Content, blog.FirstPicture, blog.Flag, blog.Views,
		blog.Appreciation, blog.ShareStatement, blog.Commentabled, blog.Published, blog.Recommend,
		blog.CreateTime, blog.UpdateTime, blog.TypeId, blog.UserId, blog.Description)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// 19、添加博客标签
func (b *BlogDao) InsertIntoBlogTags(tx *sqlx.Tx, bt []model.BlogTag) error {
	_, err := tx.NamedExec(b.sql[19], bt)
	return err
}

// 20、删除博客标签
func (b *BlogDao) DeleteBlogTagsByBlogId(tx *sqlx.Tx, id int) error {
	_, err := tx.Exec(b.sql[20], id)
	return err
}

// 21、根据创建时间获取最新的n篇博客
func (b *BlogDao) FindBlogsByCreateTime(n int) (blogs []model.BlogSection, err error) {
	err = sqldb.Select(&blogs, b.sql[21], n)
	return
}


// 22、根据浏览次数获取最热门的n篇博客
func (b *BlogDao) FindBlogsByViews(n int) (blogs []model.BlogSection, err error) {
	err = sqldb.Select(&blogs, b.sql[22], n)
	return
}

// 23、根据博客类型对博客分组、计算每个分组的数量
func (b *BlogDao) FindTypeAndBlogCount() (types []model.TheType, err error) {
	err = sqldb.Select(&types, b.sql[23])
	return
}