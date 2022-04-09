package dao

import "blog_web/model"

/*
* @Author: mgh
* @Date: 2022/3/19 11:38
* @Desc:
 */

type MottoDao struct {
	sql []string
}

func NewMottoDao() *MottoDao {
	return &MottoDao{
		sql: []string{
			`SELECT ch, en FROM t_motto;`,
			`INSERT INTO t_motto(ch, en, create_time) VALUES(?, ?, ?);`,
			`DELETE FROM t_motto WHERE id = ?;`,
			`UPDATE t_motto SET ch = ?, en = ? WHERE id = ?;`,
			`SELECT id, ch, en, create_time FROM t_motto;`,
		},
	}
}

func (m *MottoDao) FindAll() (mottos []model.Motto, err error) {
	err = sqldb.Select(&mottos, m.sql[0])
	return
}

func (m *MottoDao) AddOne(motto *model.Motto) error {
	_, err := sqldb.Exec(m.sql[1], motto.Ch, motto.En, motto.CreateTime)
	return err
}

func (m *MottoDao) DeleteOne(id uint32) error {
	_, err := sqldb.Exec(m.sql[2], id)
	return err
}

func (m *MottoDao) UpdateById(motto *model.Motto) error {
	_, err := sqldb.Exec(m.sql[3], motto.Ch, motto.En, motto.Id)
	return err
}

func (m *MottoDao) FindAllWithCreateTime() (mottos []model.Motto, err error) {
	err = sqldb.Select(&mottos, m.sql[4])
	return
}



