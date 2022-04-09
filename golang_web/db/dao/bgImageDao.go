package dao

/*
* @Author: mgh
* @Date: 2022/3/19 11:28
* @Desc:
 */

type BGImageDao struct {
	sql []string
}

func NewBGImageDao() *BGImageDao {
	return &BGImageDao{
		sql: []string{
			`SELECT url FROM t_bgimage;`,
			`INSERT into t_bgimage(url) VALUES(?);`,
			`DELETE FROM t_bgimage WHERE id = ?;`,
			`UPDATE  t_bgimage SET url = ? WHERE id = ?;`,
		},
	}
}

func (bg *BGImageDao) FindAllURL() (imgUrls []string, err error) {
	err = sqldb.Select(&imgUrls, bg.sql[0])
	return imgUrls, err
}

func (bg *BGImageDao) AddUrl(url string) error {
	_, err := sqldb.Exec(bg.sql[1], url)
	return err
}

func (bg *BGImageDao) DeleteUrl(id uint32) error {
	_, err := sqldb.Exec(bg.sql[2], id)
	return err
}

func (bg *BGImageDao) UpdateUrl(id uint32, url string) error {
	_, err := sqldb.Exec(bg.sql[3], url, id)
	return err
}