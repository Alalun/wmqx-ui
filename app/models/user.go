package models

import (
	"wmqx-ui/app/utils"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"strings"
	"github.com/snail007/go-activerecord/mysql"
	"time"
)

const (
	USER_ROLE_ROOT  = 2
	USER_ROLE_ADMIN = 1
	USER_ROLE_USER  = 0

	USER_DELETE = 1
	USER_NORMAL = 0
)

const Table_User_Name = "user"

type User struct {
}

var UserModel = User{}

// get user by user_id
func (u *User) GetUserByUserId(userId string) (user map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_User_Name).Where(map[string]interface{}{
		"user_id":   userId,
		"is_delete": USER_NORMAL,
	}))
	if err != nil {
		return
	}
	user = rs.Row()
	return
}

// user_id and username is exists
func (u *User) HasSameUsername(userId, username string) (has bool, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_User_Name).Where(map[string]interface{}{
		"user_id <>": userId,
		"username":   username,
		"is_delete":  USER_NORMAL,
	}).Limit(0, 1))
	if err != nil {
		return
	}
	if rs.Len() > 0 {
		has = true
	}
	return
}

// username is exists
func (u *User) HasUsername(username string) (has bool, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_User_Name).Where(map[string]interface{}{
		"username":  username,
		"is_delete": USER_NORMAL,
	}).Limit(0, 1))
	if err != nil {
		return
	}
	if rs.Len() > 0 {
		has = true
	}
	return
}

// get user by username
func (u *User) GetUserByName(username string) (user map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_User_Name).Where(map[string]interface{}{
		"username":  username,
		"is_delete": USER_NORMAL,
	}).Limit(0, 1))
	if err != nil {
		return
	}
	user = rs.Row()
	return
}

// delete user by user_id
func (u *User) Delete(userId string) (err error) {
	db := G.DB()
	_, err = db.Exec(db.AR().Update(Table_User_Name, map[string]interface{}{
		"is_delete": USER_DELETE,
		"update_time": time.Now().Unix(),
	}, map[string]interface{}{
		"user_id": userId,
	}))
	if err != nil {
		return
	}
	return
}

// insert user
func (u *User) Insert(userValue map[string]interface{}) (id int64, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Exec(db.AR().Insert(Table_User_Name, userValue))
	if err != nil {
		return
	}
	id = rs.LastInsertId
	return
}

// update user by user_id
func (u *User) Update(userId string, userValue map[string]interface{}) (id int64, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	userValue["update_time"] =  time.Now().Unix()
	rs, err = db.Exec(db.AR().Update(Table_User_Name, userValue, map[string]interface{}{
		"user_id":   userId,
		"is_delete": USER_NORMAL,
	}))
	if err != nil {
		return
	}
	id = rs.LastInsertId
	return
}

// update user password
func (u *User) ChangePassword(userId, newPassword, oldPassword string) (err error) {
	db := G.DB()
	userValue, err := u.GetUserByUserId(userId)
	if userValue["password"] != u.EncodePassword(oldPassword) {
		return errors.New("旧密码错误")
	}
	if err != nil {
		return
	}
	_, err = db.Exec(db.AR().Update(Table_User_Name, map[string]interface{}{
		"password": u.EncodePassword(newPassword),
		"update_time": time.Now().Unix(),
	}, map[string]interface{}{
		"user_id":   userId,
		"is_delete": USER_NORMAL,
	}))
	if err != nil {
		return
	}
	return
}

// encode password
func (u *User) EncodePassword(password string) (passwordHash string) {
	hasher := md5.New()
	hasher.Write([]byte(password))
	passwordHash = strings.ToLower(hex.EncodeToString(hasher.Sum(nil)))
	return
}

// get limit users by search keyword
func (u *User) GetUsersByKeywordAndLimit(keyword string, limit int, number int) (users []map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_User_Name).Where(map[string]interface{}{
		"username LIKE": "%" + keyword + "%",
		"is_delete":     USER_NORMAL,
	}).Limit(limit, number).OrderBy("user_id", "DESC"))
	if err != nil {
		return
	}
	users = rs.Rows()

	return
}

// get limit users
func (u *User) GetUsersByLimit(limit int, number int) (users []map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().
			From(Table_User_Name).
			Where(map[string]interface{}{
				"is_delete": USER_NORMAL,
			}).
			Limit(limit, number).
			OrderBy("user_id", "DESC"))
	if err != nil {
		return
	}
	users = rs.Rows()

	return
}

// get user count
func (u *User) CountUsers() (count int64, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().
			Select("count(*) as total").
			From(Table_User_Name).
			Where(map[string]interface{}{
				"is_delete": USER_NORMAL,
			}))
	if err != nil {
		return
	}
	count = utils.NewConvert().StringToInt64(rs.Value("total"))
	return
}

// get user count by keyword
func (u *User) CountUsersByKeyword(keyword string) (count int64, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().
		Select("count(*) as total").
		From(Table_User_Name).
		Where(map[string]interface{}{
			"username LIKE": "%" + keyword + "%",
			"is_delete":     USER_NORMAL,
		}))
	if err != nil {
		return
	}
	count = utils.NewConvert().StringToInt64(rs.Value("total"))
	return
}

// get user by username
func (u *User) GetUserByLikeName(username string) (users []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_User_Name).Where(map[string]interface{}{
		"username Like": "%" + username + "%",
		"is_delete":     USER_NORMAL,
	}).Limit(0, 1))
	if err != nil {
		return
	}
	users = rs.Rows()
	return
}

// get user by many user_id
func (u *User) GetUserByUserIds(userIds []string) (users []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_User_Name).Where(map[string]interface{}{
		"user_id":   userIds,
		"is_delete": USER_NORMAL,
	}))
	if err != nil {
		return
	}
	users = rs.Rows()
	return
}

// update user by username
func (u *User) UpdateUserByUsername(user map[string]interface{}) (affect int64, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	user["update_time"] = time.Now().Unix()
	rs, err = db.Exec(db.AR().Update(Table_User_Name, user, map[string]interface{}{
		"username": user["username"],
	}))
	if err != nil {
		return
	}
	affect = rs.RowsAffected
	return
}
