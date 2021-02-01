package dao

import (
	"bilibili/model"
	"database/sql"
	"time"
)

type UserDao struct {
	*sql.DB
}

func (dao *UserDao) UpdateLastCheckInDate(username string) error {
	timeNow := time.Now()

	stmt, err := dao.DB.Prepare(`UPDATE userinfo SET last_check_in_date = ? WHERE username = ?`)
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(timeNow, username)
	if err != nil {
		return err
	}

	return nil
}

func (dao *UserDao) UpdateCoins(username string, num int64) error {
	stmt, err := dao.DB.Prepare(`UPDATE userinfo SET coins = coins + ? WHERE username = ?`)
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(num, username)
	if err != nil {
		return err
	}

	return nil
}

func (dao *UserDao) UpdateExp(username string, num int64) error {
	stmt, err := dao.DB.Prepare(`UPDATE userinfo SET exp = exp + ? WHERE username = ?`)
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(num, username)
	if err != nil {
		return err
	}

	return nil
}

func (dao *UserDao) UpdateStatement(username, newStatement string) error {
	stmt, err := dao.DB.Prepare(`UPDATE userinfo SET statement = ? WHERE username = ?`)
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(newStatement, username)
	if err != nil {
		return err
	}

	return nil
}

func (dao *UserDao) UpdatePhone(username, newPhone string) error {
	stmt, err := dao.DB.Prepare(`UPDATE userinfo SET phone = ? WHERE username = ?`)
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(newPhone, username)
	if err != nil {
		return err
	}

	return nil
}

func (dao *UserDao) UpdateEmail(username, newEmail string) error {
	stmt, err := dao.DB.Prepare(`UPDATE userinfo SET email = ? WHERE username = ?`)
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(newEmail, username)
	if err != nil {
		return err
	}

	return nil
}

//根据邮箱查询
func (dao *UserDao) QueryByEmail(email string) (model.Userinfo, error) {
	userinfo := model.Userinfo{}

	stmt, err := dao.DB.Prepare(`SELECT uid, username, phone, salt, password, reg_date, email, statement, coins, exp, last_check_in_date FROM userinfo WHERE email = ?`)
	defer stmt.Close()

	if err != nil {
		return userinfo, err
	}

	row := stmt.QueryRow(email)

	err = row.Scan(&userinfo.Uid, &userinfo.Username, &userinfo.Phone, &userinfo.Salt, &userinfo.Password, &userinfo.RegDate, &userinfo.Email, &userinfo.Statement, &userinfo.Coins, &userinfo.Exp, &userinfo.LastCheckInDate)
	if err != nil {
		return userinfo, err
	}

	return userinfo, nil
}

//根据电话查询
func (dao *UserDao) QueryByPhone(phone string) (model.Userinfo, error) {
	userinfo := model.Userinfo{}

	stmt, err := dao.DB.Prepare(`SELECT uid, username, phone, salt, password, reg_date, email, statement, coins, exp, last_check_in_date FROM userinfo WHERE phone = ?`)
	defer stmt.Close()

	if err != nil {
		return userinfo, err
	}

	row := stmt.QueryRow(phone)

	err = row.Scan(&userinfo.Uid, &userinfo.Username, &userinfo.Phone, &userinfo.Salt, &userinfo.Password, &userinfo.RegDate, &userinfo.Email, &userinfo.Statement, &userinfo.Coins, &userinfo.Exp, &userinfo.LastCheckInDate)
	if err != nil {
		return userinfo, err
	}

	return userinfo, nil
}

//根据用户名查询
func (dao *UserDao) QueryByUsername(username string) (model.Userinfo, error) {
	userinfo := model.Userinfo{}

	stmt, err := dao.DB.Prepare(`SELECT uid, username, phone, salt, password, reg_date, email, statement, coins, exp, last_check_in_date FROM userinfo WHERE username = ?`)
	defer stmt.Close()

	if err != nil {
		return userinfo, err
	}

	row := stmt.QueryRow(username)

	err = row.Scan(&userinfo.Uid, &userinfo.Username, &userinfo.Phone, &userinfo.Salt, &userinfo.Password, &userinfo.RegDate, &userinfo.Email, &userinfo.Statement, &userinfo.Coins, &userinfo.Exp, &userinfo.LastCheckInDate)
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
