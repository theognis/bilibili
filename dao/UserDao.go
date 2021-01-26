package dao

import (
	"bilibili/model"
	"database/sql"
)

type UserDao struct {
	*sql.DB
}

//根据邮箱查询
func (dao *UserDao) QueryByEmail(email string) (model.Userinfo, error) {
	userinfo := model.Userinfo{}

	stmt, err := dao.DB.Prepare(`SELECT uid, username, phone, salt, password, reg_date, email, statement FROM userinfo WHERE email = ?`)
	defer stmt.Close()

	if err != nil {
		return userinfo, err
	}

	row := stmt.QueryRow(email)

	err = row.Scan(&userinfo.Uid, &userinfo.Username, &userinfo.Phone, &userinfo.Salt, &userinfo.Password, &userinfo.RegDate, &userinfo.Email, &userinfo.Statement)
	if err != nil {
		return userinfo, err
	}

	return userinfo, nil
}

//根据电话查询
func (dao *UserDao) QueryByPhone(phone string) (model.Userinfo, error) {
	userinfo := model.Userinfo{}

	stmt, err := dao.DB.Prepare(`SELECT uid, username, phone, salt, password, reg_date, email, statement FROM userinfo WHERE phone = ?`)
	defer stmt.Close()

	if err != nil {
		return userinfo, err
	}

	row := stmt.QueryRow(phone)

	err = row.Scan(&userinfo.Uid, &userinfo.Username, &userinfo.Phone, &userinfo.Salt, &userinfo.Password, &userinfo.RegDate, &userinfo.Email, &userinfo.Statement)
	if err != nil {
		return userinfo, err
	}

	return userinfo, nil
}

//根据用户名查询
func (dao *UserDao) QueryByUsername(username string) (model.Userinfo, error) {
	userinfo := model.Userinfo{}

	stmt, err := dao.DB.Prepare(`SELECT uid, username, phone, salt, password, reg_date, email, statement FROM userinfo WHERE username = ?`)
	defer stmt.Close()

	if err != nil {
		return userinfo, err
	}

	row := stmt.QueryRow(username)

	err = row.Scan(&userinfo.Uid, &userinfo.Username, &userinfo.Phone, &userinfo.Salt, &userinfo.Password, &userinfo.RegDate, &userinfo.Email, &userinfo.Statement)
	if err != nil {
		return userinfo, err
	}

	return userinfo, nil
}

//插入用户信息
func (dao *UserDao) InsertUser(userinfo model.Userinfo) error {
	stmt, err := dao.DB.Prepare(`INSERT INTO userinfo (username, password, reg_date, phone, salt) VALUES (?, ?, ?, ?, ?)`)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(userinfo.Username, userinfo.Password, userinfo.RegDate, userinfo.Phone, userinfo.Salt)

	stmt.Close()

	return err
}
