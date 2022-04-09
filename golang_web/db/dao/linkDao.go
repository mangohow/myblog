package dao

import "blog_web/model"

type LinkDao struct {
	sql []string
}

func NewLinkDao() *LinkDao {
	return &LinkDao{
		sql: []string{
			"SELECT id, `name`, `desc`, url, category_id, icon FROM t_resource_link;",
			"SELECT id, `name` FROM t_resource_category;",
			"SELECT id, `name`, `desc`, url, category_id, icon FROM t_resource_link LIMIT ?, ?;",
			"SELECT COUNT(*) FROM t_resource_link;",
			"INSERT INTO t_resource_link (`name`, `desc`, url, category_id, icon) VALUES (?, ?, ?, ?, ?);",
			"DELETE FROM t_resource_link WHERE id = ?",
			"UPDATE t_resource_link SET `name` = ?, `desc` = ?, url = ?, category_id = ?, icon = ? WHERE id = ?;",
			"INSERT INTO t_resource_category (`name`) VALUES (?);",
			"DELETE FROM t_resource_category WHERE id = ?;",
			"UPDATE t_resource_category SET `name` = ? WHERE id = ?;",
		},
	}
}

func (l *LinkDao) FindAllLinks() (links []model.Link, err error) {
	err = sqldb.Select(&links, l.sql[0])
	return
}

func (l *LinkDao) FindAllLinkCategory() (categories []model.LinkCategory, err error) {
	err = sqldb.Select(&categories, l.sql[1])
	return
}

func (l *LinkDao) FindLimitedLinks(pageStart, pageSize int) (links []model.Link, err error) {
	err = sqldb.Select(&links, l.sql[2], pageStart, pageSize)
	return
}

func (l *LinkDao) FindLinkCount() (count int, err error) {
	err = sqldb.Get(&count, l.sql[3])
	return
}

func (l *LinkDao) AddLink(link *model.Link) error {
	_, err := sqldb.Exec(l.sql[4], link.Name, link.Desc, link.Url, link.CategoryId, link.Icon)
	return err
}

func (l *LinkDao) DeleteLink(id int) error {
	_, err := sqldb.Exec(l.sql[5], id)
	return err
}

func (l *LinkDao) UpdateLink(link *model.Link) error {
	_, err := sqldb.Exec(l.sql[6], link.Name, link.Desc, link.Url, link.CategoryId, link.Icon, link.Id)
	return err
}

func (l *LinkDao) AddCategory(category *model.LinkCategory) error {
	_, err := sqldb.Exec(l.sql[7], category.Name)
	return err
}

func (l *LinkDao) DeleteCategory(id int) error {
	_, err := sqldb.Exec(l.sql[8], id)
	return err
}

func (l *LinkDao) UpdateCategory(category *model.LinkCategory) error {
	_, err := sqldb.Exec(l.sql[9], category.Name, category.Id)
	return err
}