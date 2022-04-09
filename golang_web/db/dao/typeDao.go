package dao

import "blog_web/model"

/*
* @Author: mgh
* @Date: 2022/3/1 11:12
* @Desc: 操作t_type表的dao
 */

type TypeDao struct {
	sql []string
}

func NewTypeDao() *TypeDao {
	return &TypeDao{
		sql: []string{
			`SELECT * FROM t_type;`,
			`SELECT * FROM t_type LIMIT ?, ?;`,
			`SELECT * FROM t_type WHERE name = ?;`,
			`DELETE FROM t_type WHERE id = ?;`,
			`UPDATE t_type SET name = ? WHERE id = ?;`,
			`INSERT INTO t_type (name) values (?);`,
			`SELECT COUNT(*) FROM t_type;`,
		},
	}
}

// 查询所有的博客类型
func (t *TypeDao) FindAll() (types []model.TheType, err error) {
	err = sqldb.Select(&types, t.sql[0])
	return
}

func (t *TypeDao) FindOnePage(pageStart, pageSize int) (types []model.TheType, err error) {
	err = sqldb.Select(&types, t.sql[1], pageStart, pageSize)
	return types, err
}

func (t *TypeDao) FindTypeByName(name string) (*model.TheType, error) {
	var tp model.TheType
	err := sqldb.Get(&tp, t.sql[2], name)
	return &tp, err
}

func (t *TypeDao) DeleteById(id int) error {
	_, err := sqldb.Exec(t.sql[3], id)
	return err
}

func (t *TypeDao) UpdateName(id int, name string) error {
	_, err := sqldb.Exec(t.sql[4], name, id)
	return err
}

func (t *TypeDao) AddType(name string) error {
	_, err := sqldb.Exec(t.sql[5], name)
	return err
}

func (t *TypeDao) GetTypesCount() (int, error) {
	var count int
	err := sqldb.Get(&count, t.sql[6])
	return count, err
}
