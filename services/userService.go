package services

import (
	"LoveDiary/model"
	"LoveDiary/utils/errmsg"
	"encoding/base64"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/scrypt"
	"log"
)

// CheckUpdateUser 检验用户名是否已经存在
func CheckUpdateUser(name string, id int) int {
	var user model.User
	db.Select("id").Where("user_name = ?", name).First(&user)

	if user.ID > 0 && int(user.ID) != id {
		return errmsg.ERROR_USERNAME_USED
	}
	// 原用户数据不存在时，user_id返回为0
	return errmsg.SUCCESS
}

// GetUserByUsername 根据用户名查找用户的信息
func GetUserByUsername(name string, password string) (model.User, int) {
	var user model.User
	err := db.Limit(1).Where("user_name = ?", name).Find(&user).Error
	if err != nil {
		return user, errmsg.ERROR
	}

	if user.ID == 0 {
		return user, errmsg.ERROR_USER_NOT_EXIST
	} else if ScryptPw(password) != user.Password {
		return user, errmsg.ERROR_PASSWORD_WRONG
	}

	return user, errmsg.SUCCESS
}

// CheckUsername 编辑用户信息时对用户名进行检验
func CheckUsername(id int, name string) int {
	var user model.User
	db.Select("id, user_name").Where("user_name = ?", name).First(&user)

	if user.ID == uint(id) {
		return errmsg.SUCCESS
	} else if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

// GetUserInfo 获取用户信息
func GetUserInfo(id int) (model.User, int) {
	var user model.User
	err := db.Limit(1).Where("id = ?", id).Find(&user).Error
	if err != nil {
		return user, errmsg.ERROR
	}
	return user, errmsg.SUCCESS
}

// CreateUser 新增用户
func CreateUser(data *model.User) int {
	data.Password = ScryptPw(data.Password)
	err := db.Create(&data).Error
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				return errmsg.ERROR_USERNAME_USED
			}
		}
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// ScryptPw 密码加密
func ScryptPw(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{12, 34, 23, 99, 23, 44, 2, 44}

	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	fpw := base64.StdEncoding.EncodeToString(HashPw)
	return fpw
}

// EditUser 编辑用户
func EditUser(id int, data *model.User) int {
	var user model.User
	var maps = make(map[string]interface{})
	maps["user_name"] = data.UserName
	maps["password"] = ScryptPw(data.Password)
	maps["email"] = data.Email
	maps["phone_number"] = data.PhoneNumber
	maps["role"] = data.Role

	err := db.Model(&user).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				return errmsg.ERROR_USERNAME_USED
			}
		}
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// DeleteUser 删除用户
func DeleteUser(id int) int {
	var user model.User
	err := db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetUsers 用户列表
func GetUsers(pageSize int, pageNum int) ([]model.User, int64) {
	var users []model.User
	var total int64
	//err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Count(&total).Error
	err := db.Select("id,user_name,role,created_at,phone_number,email").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	db.Model(&users).Count(&total)

	if err != nil {
		return users, 0
	}
	return users, total
}

// SearchUsers 搜索用户
func SearchUsers(pageSize int, pageNum int, name string) ([]model.User, int64, int) {
	var users []model.User
	var total int64
	err := db.Select("id,user_name,role,created_at,phone_number,email").Where("user_name LIKE ?", "%"+name+"%").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	db.Model(&users).Where("user_name LIKE ?", "%"+name+"%").Count(&total)
	if err != nil {
		total = 0
	}
	if total == 0 {
		return users, total, errmsg.ERROR_USER_NOT_EXIST
	}
	return users, total, errmsg.SUCCESS
}

// CheckLogin 登录验证
func CheckLogin(username string, password string) int {
	var user model.User

	db.Where("user_name = ?", username).First(&user)
	if user.ID == 0 {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	if ScryptPw(password) != user.Password {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	// 可设置管理员role为1,普通用户为0,用来区分二者
	//if user.Role != 1 {
	//	return errmsg.ERROR_USER_NO_RIGHT
	//}
	return errmsg.SUCCESS
}

// ChangePassword 管理员修改用户密码
func ChangePassword(id int, newPassword string) int {
	// 只更新password字段
	var user model.User
	// 两种单行更新的策略
	err := db.Model(&user).Where("id = ?", id).Update("password", newPassword).Error

	// err := db.Select("password").Where("id = ?", id).Updates(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func GetUser(id int) (model.User, int) {
	var user model.User
	db.First(&user, id)
	return user, errmsg.SUCCESS
}

func UpdateUserInfo(id int, data *model.User) int {
	var user model.User
	var maps = make(map[string]interface{})
	maps["user_name"] = data.UserName
	maps["email"] = data.Email
	maps["phone_number"] = data.PhoneNumber
	maps["password"] = ScryptPw(data.Password)
	// print(id)
	// print("\n")
	// print(data.UserName + " " + data.Email + " " + data.PhoneNumber + " " + data.Password + "\n")

	err := db.Model(&user).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
