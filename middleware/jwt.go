package middleware

import (
	"LoveDiary/utils"
	"LoveDiary/utils/errmsg"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var JwtKey = []byte(utils.JwtKey)
var code int

type MyClaims struct {
	UserName             string `json:"user_name"`
	jwt.RegisteredClaims        // jwt v3版本为 jwt.StandardClaims
	//jwt.Claims
}

// SetToken 生成token
func SetToken(username string, id uint, role int) (string, int) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["user_id"] = id
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 10).Unix()

	tokenString, err := token.SignedString(JwtKey)

	if err != nil {
		return "", errmsg.ERROR
	}
	return tokenString, errmsg.SUCCESS
	//expireTime := time.Now().Add(10 * time.Hour)
	//setClaims := MyClaims{
	//	UserName: username,
	//	RegisteredClaims: jwt.RegisteredClaims{
	//		// ExpiresAt: expireTime.Unix(),  // 会报错显示Unix不可用,使用jwt.NumericDate()函数进行转换
	//		ExpiresAt: jwt.NewNumericDate(expireTime), // 过期时间10小时
	//		Issuer:    "gin_blog",
	//	},
	//}
	//
	//reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, setClaims)
	//token, err := reqClaim.SignedString(JwtKey)
	//
	//if err != nil {
	//	return "", errmsg.ERROR
	//}
	//return token, errmsg.SUCCESS
}

func Secret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	}
}

// CheckToken 验证token
//func CheckToken(token string) (*MyClaims, int) {
//	setToken, _ := jwt.ParseWithClaims(token, &MyClaims{}, Secret())
//	if key, _ := setToken.Claims.(*MyClaims); setToken.Valid {
//		return key, errmsg.SUCCESS
//	} else {
//		return nil, errmsg.ERROR
//	}
//}

// ParseToken token解析
func ParseToken(tokenString string) (tokenInfo []string, code int) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if err != nil {
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userName := claims["username"].(string)
		userID := strconv.FormatFloat(claims["user_id"].(float64), 'f', -1, 64)
		role := strconv.FormatFloat(claims["role"].(float64), 'f', -1, 64)
		exp := strconv.FormatFloat(claims["exp"].(float64), 'f', -1, 64)
		tokenInfo = []string{userName, userID, role, exp}
		return tokenInfo, errmsg.SUCCESS
	}
	return []string{}, errmsg.ERROR
}

// JwtToken jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.Request.Header.Get("Authorization")
		if tokenHeader == "" {
			code = errmsg.ERROR_TOKEN_NOT_EXIST
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		checkToken := strings.SplitN(tokenHeader, " ", 2)
		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			code = errmsg.ERROR_TOKEN_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		//key, Tcode := CheckToken(checkToken[1])
		key, Tcode := ParseToken(checkToken[1])

		if Tcode == errmsg.ERROR {
			code = errmsg.ERROR_TOKEN_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		if len(key) == 0 {
			code = errmsg.ERROR_TOKEN_RUNTIME
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		c.Set("username", key[0])
		c.Set("user_id", key[1])
		c.Set("role", key[2])
		c.Next()
	}
}
