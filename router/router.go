package router

// ルーターはCORS,Cookie、エンドポイントの設定のみで不要な変更を行わないようにする
// コントローラーでエンドポイントごとの処理を行う（依存性逆転）

import (
	"go-rest-api/controller"
	"net/http"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(uc controller.IUserController, tc controller.ITaskController) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000", os.Getenv("FE_URL")},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	}))
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath:     "/",
		CookieDomain:   os.Getenv("API_DOMAIN"),
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteNoneMode, // 許可されていないところにCookieを送らない
		// CookieSameSite: http.SameSiteDefaultMode, // Postmanで試すときはこちら
		CookieMaxAge: 7 * 86400, // 1week
	}))
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)
	e.GET("/csrf", uc.CsrfToken)
	t := e.Group("/tasks")                   // タスク関係のルーターをグループ化
	t.Use(echojwt.WithConfig(echojwt.Config{ // tグループに入るためのトークン鍵設定
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	t.GET("", tc.GetAllTasks)
	t.GET("/:taskId", tc.GetTaskById)
	t.POST("", tc.CreateTask)
	t.PUT("/:taskId", tc.UpdateTask)
	t.DELETE("/:taskId", tc.DeleteTask)
	return e
}
