package dao

import (
	"bilibili/model"
	"database/sql"
	"time"
)

//更改主键查询
type UserDao struct {
	*sql.DB
}

func (dao *UserDao) UpdateLastCheckInDate(uid int64) error {
	timeNow := time.Now()

	stmt, err := dao.DB.Prepare(`UPDATE userinfo SET last_check_in_date = ? WHERE uid = ?`)
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(timeNow, uid)
	if err != nil {
		return err
	}

	return nil
}

func (dao *UserDao) UpdateCoins(uid int64, num int64) error {
	stmt, err := dao.DB.Prepare(`UPDATE userinfo SET coins = coins + ? WHERE uid = ?`)
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(num, uid)
	if err != nil {
		return err
	}

	return nil
}

func (dao *UserDao) UpdatePassword(uid int64, newPassword string) error {
	stmt, err := dao.DB.Prepare(`UPDATE userinfo SET password = ? WHERE uid = ?`)
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(newPassword, uid)
	if err != nil {
		return err
	}

	return nil
}

func (dao *UserDao) UpdateUsername(uid int64, newUsername string) error {
	stmt, err := dao.DB.Prepare(`UPDATE userinfo SET username = ? WHERE uid = ?`)
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(newUsername, uid)
	if err != nil {
		return err
	}

	return nil
}

func (dao *UserDao) UpdateExp(uid int64, num int64) error {
	stmt, err := dao.DB.Prepare(`UPDATE userinfo SET exp = exp + ? WHERE uid = ?`)
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(num, uid)
	if err != nil {
		return err
	}

	return nil
}

func (dao *UserDao) UpdateStatement(uid int64, newStatement string) error {
	stmt, err := dao.DB.Prepare(`UPDATE userinfo SET statement = ? WHERE uid = ?`)
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(newStatement, uid)
	if err != nil {
		return err
	}

	return nil
}

func (dao *UserDao) UpdateBirthday(uid int64, newBirthday time.Time) error {
	stmt, err := dao.DB.Prepare(`UPDATE userinfo SET birthday = ? WHERE uid = ?`)
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(newBirthday, uid)
	if err != nil {
		return err
	}

	return nil
}

func (dao *UserDao) UpdateGender(uid int64, newGender string) error {
	stmt, err := dao.DB.Prepare(`UPDATE userinfo SET gender = ? WHERE uid = ?`)
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(newGender, uid)
	if err != nil {
		return err
	}

	return nil
}

func (dao *UserDao) UpdateAvatar(uid int64, url string) error {
	stmt, err := dao.DB.Prepare(`UPDATE userinfo SET avatar = ? WHERE uid = ?`)
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(url, uid)
	if err != nil {
		return err
	}

	return nil
}

func (dao *UserDao) UpdatePhone(uid int64, newPhone string) error {
	stmt, err := dao.DB.Prepare(`UPDATE userinfo SET phone = ? WHERE uid = ?`)
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(newPhone, uid)
	if err != nil {
		return err
	}

	return nil
}

func (dao *UserDao) UpdateEmail(uid int64, newEmail string) error {
	stmt, err := dao.DB.Prepare(`UPDATE userinfo SET email = ? WHERE uid = ?`)
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(newEmail, uid)
	if err != nil {
		return err
	}

	return nil
}

//根据uid查询
func (dao *UserDao) QueryByUid(uid int64) (model.Userinfo, error) {
	userinfo := model.Userinfo{}

	stmt, err := dao.DB.Prepare(`SELECT uid, username, phone, salt, password, reg_date, email, statement, coins, exp, last_check_in_date, b_coins, avatar, birthday, gender FROM userinfo WHERE uid = ?`)
	defer stmt.Close()

	if err != nil {
		return userinfo, err
	}

	row := stmt.QueryRow(uid)

	err = row.Scan(&userinfo.Uid, &userinfo.Username, &userinfo.Phone, &userinfo.Salt, &userinfo.Password, &userinfo.RegDate, &userinfo.Email, &userinfo.Statement, &userinfo.Coins, &userinfo.Exp, &userinfo.LastCheckInDate, &userinfo.BCoins, &userinfo.Avatar, &userinfo.Birthday, &userinfo.Gender)
	if err != nil {
		return userinfo, err
	}

	return userinfo, nil
}

//根据邮箱查询
func (dao *UserDao) QueryByEmail(email string) (model.Userinfo, error) {
	userinfo := model.Userinfo{}

	stmt, err := dao.DB.Prepare(`SELECT uid, username, phone, salt, password, reg_date, email, statement, coins, exp, last_check_in_date, b_coins, avatar, birthday, gender FROM userinfo WHERE email = ?`)
	defer stmt.Close()

	if err != nil {
		return userinfo, err
	}

	row := stmt.QueryRow(email)

	err = row.Scan(&userinfo.Uid, &userinfo.Username, &userinfo.Phone, &userinfo.Salt, &userinfo.Password, &userinfo.RegDate, &userinfo.Email, &userinfo.Statement, &userinfo.Coins, &userinfo.Exp, &userinfo.LastCheckInDate, &userinfo.BCoins, &userinfo.Avatar, &userinfo.Birthday, &userinfo.Gender)
	if err != nil {
		return userinfo, err
	}

	return userinfo, nil
}

//根据电话查询
func (dao *UserDao) QueryByPhone(phone string) (model.Userinfo, error) {
	userinfo := model.Userinfo{}

	stmt, err := dao.DB.Prepare(`SELECT uid, username, phone, salt, password, reg_date, email, statement, coins, exp, last_check_in_date, b_coins, avatar, birthday, gender FROM userinfo WHERE phone = ?`)
	defer stmt.Close()

	if err != nil {
		return userinfo, err
	}

	row := stmt.QueryRow(phone)

	err = row.Scan(&userinfo.Uid, &userinfo.Username, &userinfo.Phone, &userinfo.Salt, &userinfo.Password, &userinfo.RegDate, &userinfo.Email, &userinfo.Statement, &userinfo.Coins, &userinfo.Exp, &userinfo.LastCheckInDate, &userinfo.BCoins, &userinfo.Avatar, &userinfo.Birthday, &userinfo.Gender)
	if err != nil {
		return userinfo, err
	}

	return userinfo, nil
}

//根据用户名查询
func (dao *UserDao) QueryByUsername(username string) (model.Userinfo, error) {
	userinfo := model.Userinfo{}

	stmt, err := dao.DB.Prepare(`SELECT uid, username, phone, salt, password, reg_date, email, statement, coins, exp, last_check_in_date, b_coins, avatar, birthday, gender FROM userinfo WHERE username = ?`)
	defer stmt.Close()

	if err != nil {
		return userinfo, err
	}

	row := stmt.QueryRow(username)

	err = row.Scan(&userinfo.Uid, &userinfo.Username, &userinfo.Phone, &userinfo.Salt, &userinfo.Password, &userinfo.RegDate, &userinfo.Email, &userinfo.Statement, &userinfo.Coins, &userinfo.Exp, &userinfo.LastCheckInDate, &userinfo.BCoins, &userinfo.Avatar, &userinfo.Birthday, &userinfo.Gender)
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
