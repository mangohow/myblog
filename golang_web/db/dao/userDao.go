package dao

import "blog_web/model"

/*
* @Author: mgh
* @Date: 2022/3/2 11:30
* @Desc:
 */

type UserDao struct {
	sql []string
}

func NewUserDao() *UserDao {
	return &UserDao{
		sql: []string{
			`SELECT * FROM t_user WHERE username = ? AND password = ?;`,
			`SELECT nickname, email, avatar, github, csdn FROM t_user Limit 1`,
		},
	}
}

func (u *UserDao) FindUserByUsernameAndPassword(username, password string) (*model.User, error) {
	var user model.User
	err := sqldb.Get(&user, u.sql[0], username, password)
	return &user, err
}

func (u *UserDao) FindInfo() (*model.User, error) {
	var user model.User
	err := sqldb.Get(&user, u.sql[1])
	return &user, err
}