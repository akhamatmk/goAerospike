package route

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	us "github.com/akhamatvarokah/goAerospike/controllers/userController"
	ac "github.com/akhamatvarokah/goAerospike/controllers"
)

type jwtCustomClaims struct {
	Email string
	Role  string
	jwt.StandardClaims
}

// Route ...
func Route() *echo.Echo {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", ac.Getdata)
	e.GET("/:key", ac.Getdata)

	e.POST("/", ac.Insert)
	e.PUT("/", ac.Edit)
	e.DELETE("/:namespace/:setname/:bin/:key", ac.DeleteData)
	e.GET("/nameSpace", ac.GetNameSpace)
	e.GET("/setname/:namespace", ac.GetSetName)
	e.POST("/login", us.Login)
	e.POST("/user/Adduser", us.AddUser)

	//r := e.Group("/restricted")
	//config := middleware.JWTConfig{
	//	Claims:     &jwtCustomClaims{},
	//	SigningKey: []byte("secret"),
	//}
	//
	//r.Use(middleware.JWTWithConfig(config))
	//r.GET("", restricted, isAdmin)

	return e
}

var isLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("secret"),
})

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	name := claims.Email
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

func isAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*jwtCustomClaims)
		role := claims.Role

		if role != "Admin" {
			return echo.ErrUnauthorized
		}

		return next(c)
	}
}
