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

//更新关注中数量
func (dao *UserDao) UpdateFollowing(uid, num int64) error {
	stmt, err := dao.DB.Prepare(`UPDATE userinfo SET followings = followings + ? WHERE uid = ?`)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(num, uid)
	if err != nil {
		return err
	}

	return nil
}

//更新关注者数量
func (dao *UserDao) UpdateFollower(uid, num int64) error {
	stmt, err := dao.DB.Prepare(`UPDATE userinfo SET followers = followers + ? WHERE uid = ?`)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(num, uid)
	if err != nil {
		return err
	}

	return nil
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

func (dao *UserDao) UpdatePassword(uid int64, saltedPassword string, salt string) error {
	stmt, err := dao.DB.Prepare(`UPDATE userinfo SET password = ?, salt = ? WHERE uid = ?`)
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(saltedPassword, salt, uid)
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

func (dao *UserDao) UpdateLastViewDate(uid int64) error {
	timeNow := time.Now()
	stmt, err := dao.DB.Prepare(`UPDATE userinfo SET last_view_date = ? WHERE uid = ?`)
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

func (dao *UserDao) UpdateLastCoinDate(uid int64) error {
	timeNow := time.Now()
	stmt, err := dao.DB.Prepare(`UPDATE userinfo SET last_coin_date = ? WHERE uid = ?`)
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

func (dao *UserDao) UpdateDailyCoin(uid int64, stm int64) error {
	var query string
	if stm == 1 {
		query = "UPDATE userinfo SET daily_coin = 1 WHERE uid = ?"
	} else {
		query = "UPDATE userinfo SET daily_coin = daily_coin + 1 WHERE uid = ?"
	}
	stmt, err := dao.DB.Prepare(query)
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(uid)
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

//根据uid查询可以在个人空间显示的用户信息
func (dao *UserDao) QuerySpaceUserinfoByUid(uid int64) (model.SpaceUserinfo, error) {
	spaceUserinfo := model.SpaceUserinfo{}

	stmt, err := dao.DB.Prepare(`SELECT avatar, uid, username, reg_date, statement, exp, coins, b_coins, birthday, gender, followers, followings, total_views, total_likes FROM userinfo WHERE uid = ?`)
	defer stmt.Close()

	if err != nil {
		return spaceUserinfo, err
	}

	row := stmt.QueryRow(uid)
	var regDate, birthDate time.Time

	err = row.Scan(&spaceUserinfo.Avatar, &spaceUserinfo.Uid, &spaceUserinfo.Username, &regDate, &spaceUserinfo.Statement, &spaceUserinfo.Exp, &spaceUserinfo.Coins, &spaceUserinfo.BCoins, &birthDate, &spaceUserinfo.Gender, &spaceUserinfo.Followers, &spaceUserinfo.Followings, &spaceUserinfo.TotalViews, &spaceUserinfo.TotalLikes)
	if err != nil {
		return spaceUserinfo, err
	}

	spaceUserinfo.Birthday = birthDate.Format("2006-01-02")
	spaceUserinfo.RegDate = regDate.Format("2006-01-02")

	return spaceUserinfo, nil
}

//根据uid查询
func (dao *UserDao) QueryByUid(uid int64) (model.Userinfo, error) {
	userinfo := model.Userinfo{}

	stmt, err := dao.DB.Prepare(`SELECT uid, username, phone, salt, password, reg_date, email, statement, coins, exp, last_check_in_date, b_coins, avatar, birthday, gender, last_coin_date, daily_coin, last_view_date, followers, followings, total_views, total_likes FROM userinfo WHERE uid = ?`)
	if err != nil {
		return userinfo, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(uid)

	err = row.Scan(&userinfo.Uid, &userinfo.Username, &userinfo.Phone, &userinfo.Salt, &userinfo.Password, &userinfo.RegDate, &userinfo.Email, &userinfo.Statement, &userinfo.Coins, &userinfo.Exp, &userinfo.LastCheckInDate, &userinfo.BCoins, &userinfo.Avatar, &userinfo.Birthday, &userinfo.Gender, &userinfo.LastCoinDate, &userinfo.DailyCoin, &userinfo.LastViewDate, &userinfo.Followers, &userinfo.Followings, &userinfo.TotalViews, &userinfo.TotalLikes)
	if err != nil {
		return userinfo, err
	}

	return userinfo, nil
}

//根据邮箱查询
func (dao *UserDao) QueryByEmail(email string) (model.Userinfo, error) {
	userinfo := model.Userinfo{}

	stmt, err := dao.DB.Prepare(`SELECT uid, username, phone, salt, password, reg_date, email, statement, coins, exp, last_check_in_date, b_coins, avatar, birthday, gender, last_coin_date, daily_coin, last_view_date, followers, followings, total_views, total_likes FROM userinfo WHERE email = ?`)
	defer stmt.Close()

	if err != nil {
		return userinfo, err
	}

	row := stmt.QueryRow(email)

	err = row.Scan(&userinfo.Uid, &userinfo.Username, &userinfo.Phone, &userinfo.Salt, &userinfo.Password, &userinfo.RegDate, &userinfo.Email, &userinfo.Statement, &userinfo.Coins, &userinfo.Exp, &userinfo.LastCheckInDate, &userinfo.BCoins, &userinfo.Avatar, &userinfo.Birthday, &userinfo.Gender, &userinfo.LastCoinDate, &userinfo.DailyCoin, &userinfo.LastViewDate, &userinfo.Followers, &userinfo.Followings, &userinfo.TotalViews, &userinfo.TotalLikes)
	if err != nil {
		return userinfo, err
	}

	return userinfo, nil
}

//根据电话查询
func (dao *UserDao) QueryByPhone(phone string) (model.Userinfo, error) {
	userinfo := model.Userinfo{}

	stmt, err := dao.DB.Prepare(`SELECT uid, username, phone, salt, password, reg_date, email, statement, coins, exp, last_check_in_date, b_coins, avatar, birthday, gender, last_coin_date, daily_coin, last_view_date, followers, followings, total_views, total_likes FROM userinfo WHERE phone = ?`)
	defer stmt.Close()

	if err != nil {
		return userinfo, err
	}

	row := stmt.QueryRow(phone)

	err = row.Scan(&userinfo.Uid, &userinfo.Username, &userinfo.Phone, &userinfo.Salt, &userinfo.Password, &userinfo.RegDate, &userinfo.Email, &userinfo.Statement, &userinfo.Coins, &userinfo.Exp, &userinfo.LastCheckInDate, &userinfo.BCoins, &userinfo.Avatar, &userinfo.Birthday, &userinfo.Gender, &userinfo.LastCoinDate, &userinfo.DailyCoin, &userinfo.LastViewDate, &userinfo.Followers, &userinfo.Followings, &userinfo.TotalViews, &userinfo.TotalLikes)
	if err != nil {
		return userinfo, err
	}

	return userinfo, nil
}

//根据用户名查询
func (dao *UserDao) QueryByUsername(username string) (model.Userinfo, error) {
	userinfo := model.Userinfo{}

	stmt, err := dao.DB.Prepare(`SELECT uid, username, phone, salt, password, reg_date, email, statement, coins, exp, last_check_in_date, b_coins, avatar, birthday, gender, last_coin_date, daily_coin, last_view_date, followers, followings, total_views, total_likes FROM userinfo WHERE username = ?`)
	defer stmt.Close()

	if err != nil {
		return userinfo, err
	}

	row := stmt.QueryRow(username)

	err = row.Scan(&userinfo.Uid, &userinfo.Username, &userinfo.Phone, &userinfo.Salt, &userinfo.Password, &userinfo.RegDate, &userinfo.Email, &userinfo.Statement, &userinfo.Coins, &userinfo.Exp, &userinfo.LastCheckInDate, &userinfo.BCoins, &userinfo.Avatar, &userinfo.Birthday, &userinfo.Gender, &userinfo.LastCoinDate, &userinfo.DailyCoin, &userinfo.LastViewDate, &userinfo.Followers, &userinfo.Followings, &userinfo.TotalViews, &userinfo.TotalLikes)
	if err != nil {
		return userinfo, err
	}

	return userinfo, nil
}

func (dao *UserDao) QueryFollowedUid(followingUid int64) ([]int64, error) {
	var uidSlice []int64

	stmt, err := dao.DB.Prepare(`SELECT followed_uid FROM user_follow WHERE following = ?`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(followingUid)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var uid int64
		err = rows.Scan(&uid)
		if err != nil {
			return nil, err
		}

		uidSlice = append(uidSlice, uid)
	}

	return uidSlice, nil
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

func (dao *UserDao) InsertFollow(followerUId, followingUid int64) error {
	stmt, err := dao.DB.Prepare(`INSERT INTO user_follow (follower_uid, following_uid) VALUES (?, ?)`)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(followerUId, followingUid)
	return err
}

func (dao *UserDao) DeleteFollow(followerUid, followingUid int64) error {
	stmt, err := dao.DB.Prepare(`DELETE FROM user_follow WHERE (follower_uid = ? and following_uid = ?)`)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(followerUid, followingUid)
	return err
}
