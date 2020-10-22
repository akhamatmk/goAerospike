package UserController

import (
	"fmt"
	"github.com/akhamatvarokah/goAerospike/Utils"
	ar "github.com/akhamatvarokah/goAerospike/service/aerospike"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type postUser struct {
	Email    string `json:"email" form:"email" query:"email"`
	Password string `json:"password" form:"password" query:"password"`
	Role     string `json:"role" form:"role" query:"role"`
}

type jwtCustomClaims struct {
	Email string
	Role  string
	jwt.StandardClaims
}

type User struct {
	Role  string `json:"role" form:"role" query:"role"`
	Token string `json:"token" form:"token" query:"token"`
}

func Login(c echo.Context) error {
	u := new(postUser)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	key := u.Email
	namespace := "test"
	setname := "user"
	keybin := "key1"

	result := ar.GetValueByKey(namespace, setname, keybin, key)
	if result == nil {
		return c.JSON(http.StatusBadRequest, Utils.ResponseOk("Email Or password wrong"))
	}

	s := result.([]interface{})
	password := ""
	role := ""

	for key, val := range s {
		if key == 0 {
			password = fmt.Sprint(val)
		}

		if key == 1 {
			role = fmt.Sprint(val)
		}
	}

	cp := CheckPasswordHash(u.Password, password)
	if !cp {
		return c.JSON(http.StatusBadRequest, Utils.ResponseOk("Email Or password wrong"))
	}

	token, _ := Utils.GetToken(u.Email, role)
	return c.JSON(http.StatusOK, Utils.ResponseOk(User{
		Role:  role,
		Token: token,
	}))
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func AddUser(c echo.Context) error {
	u := new(postUser)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if u.Role == "" || u.Email == "" || u.Password == "" {
		if err := c.Bind(u); err != nil {
			return c.JSON(http.StatusBadRequest, Utils.ResponseOk("All data must be filled"))
		}
	}

	password, _ := HashPassword(u.Password)

	result := ar.InsertData(ar.PaylodAerospike{
		NameSpace: "test",
		SetName:   "user",
		Key:       u.Email,
		Value:     []interface{}{password, u.Role},
		KeyBin:    "key1",
	})

	return c.JSON(http.StatusOK, Utils.ResponseOk(result))
}
