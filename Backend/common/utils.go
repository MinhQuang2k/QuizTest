package common

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"gopkg.in/go-playground/validator.v8"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

const NBSecretPassword = "A String Very Very Very Strong!!@##$!@#$"
const NBRandomPassword = "A String Very Very Very Niubilty!!@##$!@#4"
const bearerPrefix = "Bearer "

type SignedDetails struct {
	Id uint
	jwt.StandardClaims
}

func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func GenToken(id uint) string {
	claims := &SignedDetails{
		Id: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}
	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(NBSecretPassword))

	return token
}

func BearerAuth(c *gin.Context) string {

	authHeaderValue := c.Request.Header.Get("Authorization")

	if !strings.HasPrefix(authHeaderValue, bearerPrefix) {
		return ""
	}

	token := strings.TrimPrefix(authHeaderValue, bearerPrefix)
	if len(token) == 0 {
		return ""
	}

	return token
}

func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {

	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(NBSecretPassword), nil
		},
	)

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = fmt.Sprintf("the token is invalid")
		msg = err.Error()
		return
	}

	return claims, msg

}

type CommonError struct {
	Errors map[string]interface{} `json:"errors"`
}

func NewValidatorError(err error) CommonError {
	res := CommonError{}
	res.Errors = make(map[string]interface{})
	errs := err.(validator.ValidationErrors)
	for _, v := range errs {
		if v.Param != "" {
			res.Errors[v.Field] = fmt.Sprintf("{%v: %v}", v.Tag, v.Param)
		} else {
			res.Errors[v.Field] = fmt.Sprintf("{key: %v}", v.Tag)
		}

	}
	return res
}

func NewError(key string, err error) CommonError {
	res := CommonError{}
	res.Errors = make(map[string]interface{})
	res.Errors[key] = err.Error()
	return res
}

func Bind(c *gin.Context, obj interface{}) error {
	b := binding.Default(c.Request.Method, c.ContentType())
	return c.ShouldBindWith(obj, b)
}
