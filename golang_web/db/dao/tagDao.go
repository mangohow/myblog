package dao

import "blog_web/model"

/*
* @Author: mgh
* @Date: 2022/3/1 9:48
* @Desc: 操作t_tag表的dao
 */

type TagDao struct {
	sql []string
}

func NewTagDao() *TagDao {
	return &TagDao{
		sql: []string{
			`SELECT t.id, t.name FROM t_blog_tags bt JOIN t_tag t ON t.id = bt.tag_id WHERE bt.blog_id = ?;`,
			`SELECT * FROM t_tag ORDER BY id ASC;`,
			`SELECT * FROM t_tag LIMIT ?, ?;`,
			`SELECT * FROM t_tag WHERE name = ?;`,
			`DELETE FROM t_tag WHERE id = ?;`,
			`UPDATE t_tag SET name = ? WHERE id = ?;`,
			`INSERT INTO t_tag (name) VALUES(?);`,
			`SELECT COUNT(*) FROM t_tag;`,
		},
	}
}

// 根据博客ID获取所有的标签
func (t *TagDao) GetTagsByBlogId(id int) (tags []model.Tag, err error) {
	err = sqldb.Select(&tags, t.sql[0], id)
	return
}

// 查询所有标签
func (t *TagDao) FindAllTags() (tags []model.Tag, err error) {
	err = sqldb.Select(&tags, t.sql[1])
	return
}

func (t *TagDao) GetOnePageTags(pageStart, pageSize int) (tags []model.Tag, err error) {
	err = sqldb.Select(&tags, t.sql[2], pageStart, pageSize)
	return
}

func (t *TagDao) FindByName(name string) (*model.Tag, error) {
	var tag model.Tag
	err := sqldb.Get(&tag, t.sql[3], name)
	return &tag, err
}

func (t *TagDao) DeleteById(id int) error {
	_, err := sqldb.Exec(t.sql[4], id)
	return err
}

func (t *TagDao) UpdateName(id int, name string) error {
	_, err := sqldb.Exec(t.sql[5], name, id)
	return err
}

func (t *TagDao) AddTag(name string) error {
	_, err := sqldb.Exec(t.sql[6], name)
	return err
}

func (t *TagDao) GetTagsCount() (int, error) {
	var count int
	err := sqldb.Get(&count, t.sql[7])
	return count, err
}
