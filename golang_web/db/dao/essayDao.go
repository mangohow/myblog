package dao

import "blog_web/model"

type EssayDao struct {
	sql []string
}

func NewEssayDao() *EssayDao {
	return &EssayDao{
		sql: []string{
			`SELECT id, create_time, content FROM t_essay ORDER BY create_time DESC;`,
			`SELECT id, create_time, content FROM t_essay ORDER BY create_time DESC LIMIT ?, ?;`,
			`SELECT COUNT(*) FROM t_essay`,
			`INSERT INTO t_essay (create_time, content) VALUES (?, ?);`,
			`DELETE FROM t_essay WHERE id = ?;`,
			`UPDATE t_essay SET content = ? WHERE id = ?;`,
		},
	}
}

func (e *EssayDao) FindAll() (essays []model.Essay, err error) {
	err = sqldb.Select(&essays, e.sql[0])
	return
}

func (e *EssayDao) FindLimited(pageStart, pageSize int) (essays []model.Essay, err error) {
	err = sqldb.Select(&essays, e.sql[1], pageStart, pageSize)
	return
}

func (e *EssayDao) FindTotalCount() (count int, err error) {
	err = sqldb.Get(&count, e.sql[2])
	return
}

func (e *EssayDao) Add(essay *model.Essay) error {
	_, err := sqldb.Exec(e.sql[3], essay.CreateTime, essay.Content)
	return err
}

func (e *EssayDao) Delete(id int) error {
	_, err := sqldb.Exec(e.sql[4], id)
	return err
}

func (e *EssayDao) Update(essay *model.Essay) error {
	_, err := sqldb.Exec(e.sql[5], essay.Content, essay.Id)
	return err
}