package auth

import (
	"errors"
	"net/http"
	"time"

	"github.com/msterzhang/onelist/api/database"
	"github.com/msterzhang/onelist/api/models"
	"github.com/msterzhang/onelist/config"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

/*
 * 非系统管理员用户拦截器
 */
func JWTAuthAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		if token == "" {
			token = c.Request.Header.Get("Authorization")
		}
		if token == "" {
			token = c.PostForm("token")
		}
		if token == "" {
			token = c.Request.Header.Get("token")
		}
		if len(token) == 0 {
			c.JSON(http.StatusOK, gin.H{"code": 201, "msg": "not token"})
		}
		claims, err := ParseToken(token)
		if err != nil {
			if err == ErrTokenExpired {
				if token, err = RefreshToken(token); err == nil {
					c.Header("Authorization", token)
					c.JSON(http.StatusOK, gin.H{"code": 201, "msg": "refresh token", "token": token})
					c.AbortWithStatus(http.StatusBadRequest)
					return
				}
			}
			c.JSON(http.StatusUnauthorized, gin.H{"code": 403, "msg": err.Error()})
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if !claims.User.IsAdmin {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 403, "msg": "非系统管理员，禁止访问!"})
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		c.Set("claims", claims)
		c.Set("Id", claims.User.Id)
		c.Set("UserId", claims.User.UserId)
	}
}

/*
 * 非登录用户拦截器
 */
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		if token == "" {
			token = c.Request.Header.Get("Authorization")
		}
		if token == "" {
			token = c.PostForm("token")
		}
		if token == "" {
			token = c.Request.Header.Get("token")
		}
		claims, err := ParseToken(token)
		if err != nil {
			if err == ErrTokenExpired {
				if token, err = RefreshToken(token); err == nil {
					c.Header("Authorization", token)
					c.JSON(http.StatusOK, gin.H{"code": 203, "msg": "refresh token", "token": token})
					c.AbortWithStatus(http.StatusBadRequest)
					return
				}
			}
			c.JSON(http.StatusUnauthorized, gin.H{"code": 403, "msg": err.Error()})
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		db := database.NewDb()
		user := models.User{}
		err = db.Model(&models.User{}).Where("user_email = ?", claims.User.UserEmail).Take(&user).Error
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 403, "msg": "error token!"})
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if user.IsLock {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 403, "msg": "账号被锁定"})
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		c.Set("claims", claims)
		c.Set("Id", claims.User.Id)
		c.Set("UserId", claims.User.UserId)
	}
}

var (
	ErrTokenExpired     = errors.New("token is expired")
	ErrTokenNotValidYet = errors.New("token not active yet")
	ErrTokenMalformed   = errors.New("that's not even a token")
	ErrTokenInvalid     = errors.New("couldn't handle this token")
	SignKey             = "test"
)

/*
 * 新建JWT数据
 */
func GenerateJWT(user models.User) (string, error) {
	claim := models.Claim{
		User: user,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "Alfredo Mendoza",
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(config.SECRETKEY)
}

/*
 * 解析token数据
 */
func ParseToken(tokenString string) (*models.Claim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.Claim{}, func(token *jwt.Token) (interface{}, error) {
		return config.SECRETKEY, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, ErrTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, ErrTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, ErrTokenNotValidYet
			} else {
				return nil, ErrTokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*models.Claim); ok && token.Valid {
		return claims, nil
	}
	return nil, ErrTokenInvalid
}

/*
 * 刷新token数据
 */
func RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &models.Claim{}, func(token *jwt.Token) (interface{}, error) {
		return config.SECRETKEY, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*models.Claim); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return GenerateJWT(claims.User)
	}
	return "", ErrTokenInvalid
}
