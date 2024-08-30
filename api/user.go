package api

import (
	"LoveDiary/model"
	"LoveDiary/services"
	"LoveDiary/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var code int

func GetTokenInfo(c *gin.Context) {
	username, exists := c.Get("username")
	user_id, _ := c.Get("user_id")
	role, _ := c.Get("role")

	if !exists {
		code := errmsg.ERROR
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": "username not found",
		})
	}

	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":   code,
		"username": username,
		"user_id":  user_id,
		"role":     role,
	})
}

// CreateUser 新增用户
func CreateUser(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)

	code = services.CreateUser(&data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// EditUser 编辑用户
func EditUser(c *gin.Context) {
	idInterface, exists := c.Get("user_id")
	//fmt.Println("idInterface", idInterface)
	if !exists {
		code = errmsg.ERROR
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
	}

	id, _ := strconv.Atoi(idInterface.(string))
	var data model.User
	//id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	// 先检查用户名是否已经存在
	//code = services.CheckUsername(id, data.UserName)
	//if code == errmsg.SUCCESS {
	//	services.EditUser(id, &data)
	//}
	code = services.EditUser(id, &data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"data":    data,
	})
}

func EditUserByAdmin(c *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)

	code = services.EditUser(id, &data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"data":    data,
	})
}

// GetUserInfo 查询用户信息  token认证方式
func GetUserInfo(c *gin.Context) {
	var maps = make(map[string]interface{})
	var data model.User
	//id, _ := strconv.Atoi(c.Param("id"))
	idInterface, exists := c.Get("user_id")
	//fmt.Println("idInterface", idInterface)
	if !exists {
		code = errmsg.ERROR
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
	}

	id, _ := strconv.Atoi(idInterface.(string))

	data, code = services.GetUserInfo(id)
	maps["username"] = data.UserName
	maps["password"] = data.Password
	maps["role"] = data.Role
	maps["phone_number"] = data.PhoneNumber
	maps["email"] = data.Email
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    maps,
		"message": errmsg.GetErrMsg(code),
		"total":   1,
	})
}

// GetUserInfoByAdmin 管理员获取用户信息
func GetUserInfoByAdmin(c *gin.Context) {
	var maps = make(map[string]interface{})
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))

	data, code = services.GetUserInfo(id)
	maps["username"] = data.UserName
	maps["password"] = data.Password
	maps["role"] = data.Role
	maps["phone_number"] = data.PhoneNumber
	maps["email"] = data.Email
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    maps,
		"message": errmsg.GetErrMsg(code),
		"total":   1,
	})
}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = services.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetUsers 用户列表
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	pageNum, _ := strconv.Atoi(c.Query("page_num"))

	if pageSize == 0 {
		// gorm中设置为-1会取消Limit,Offset限制
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}

	data, total := services.GetUsers(pageSize, pageNum)
	code = errmsg.SUCCESS //  ？？？？
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		"total":  total,
		"msg":    errmsg.GetErrMsg(code),
	})
}

// SearchUsers 搜索用户
func SearchUsers(c *gin.Context) {
	name := c.Query("username")
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	pageNum, _ := strconv.Atoi(c.Query("page_num"))

	if pageSize == 0 {
		// gorm中设置为-1会取消Limit,Offset限制
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data, total, code := services.SearchUsers(pageSize, pageNum, name)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// ChangePassword 修改用户密码
func ChangePassword(c *gin.Context) {
	idInterface, exists := c.Get("user_id")
	if !exists {
		code = errmsg.ERROR
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
	}

	id, _ := strconv.Atoi(idInterface.(string))

	var requestData struct {
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	requestData.Password = services.ScryptPw(requestData.Password)
	code = services.ChangePassword(id, requestData.Password)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    requestData,
		"message": errmsg.GetErrMsg(code),
	})
}

// ChangePasswordByAdmin 管理员修改用户密码
func ChangePasswordByAdmin(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var requestData struct {
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	requestData.Password = services.ScryptPw(requestData.Password)
	code = services.ChangePassword(id, requestData.Password)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    requestData,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetUser 获取当前用户的个人信息
func GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	var user model.User
	user, code = services.GetUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    user,
		"message": errmsg.GetErrMsg(code),
	})
}

func UpdateUserInfo(c *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	// print("test\n")
	// print(data.PhoneNumber)
	// 先检查名是否已经存在
	code = services.CheckUpdateUser(data.UserName, id)
	// print(code == errmsg.SUCCESS)
	if code == errmsg.SUCCESS {
		services.UpdateUserInfo(id, &data)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"data":    data,
	})
}

// ScryptPassword 修改密码时，传入密码加密
func ScryptPassword(c *gin.Context) {
	var requestData struct {
		OriginPass string `json:"originPass"`
	}
	_ = c.ShouldBindJSON(&requestData)

	scryPass := services.ScryptPw(requestData.OriginPass)

	c.JSON(http.StatusOK, gin.H{
		"scryPass": scryPass,
	})
}
