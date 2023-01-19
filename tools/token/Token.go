package tools

import (
	"time"

	"github.com/golang-jwt/jwt"
)

// token负载，即token中蕴含的信息
type MyCliams struct {
	Id int // user_id
	jwt.StandardClaims
}

// 签发私钥
var appKey = []byte("404NotFound")

/*
 * 函数功能：创建token
 * 输入参数 id:用户ID，即user_id
 * 返回值 TokenString:token的字符串形式
 * tips:这里没有存入时间，暂时默认创建的token为永久token
 * 如果需要限制token有效时间可以在上面的MyCliams结构体中
 * 添加token签发时间与token有效期
 */
func CreateToken(id int) string {
	c := MyCliams{
		Id: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(), //签发时间
			Issuer:    "yyy",                                //签发人
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	TokenString, err := token.SignedString(appKey)

	if err != nil {
		panic(err)
	}

	return TokenString
}

/*
 * 函数功能：token认证
 * 输入参数 TokenString:token的字符串形式
 * 返回值 flag:鉴权是否成功
 *        id:提取出的user_id
 */
func CheckToke(TokenString string) (flag bool, id int) {
	/*
		这部分后面导出日志时可以使用
		if len(appKey) == 0 {
			log.Fatal("Server unable to start, expected an APP_KEY for JWT auth")
		}
	*/

	token, err := jwt.ParseWithClaims(TokenString, &MyCliams{},
		func(token *jwt.Token) (i interface{}, err error) {
			return appKey, nil
		})
	if err != nil {
		panic(err)
	}
	if claims, ok := token.Claims.(*MyCliams); ok && token.Valid {
		// 校验token
		return true, claims.Id
	}
	return false, -1
}
