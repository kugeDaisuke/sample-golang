package handler

import (
	"net/http"
	"sampleApi/db"
	"sampleApi/entity"

	"github.com/labstack/echo/v4"
)

type CreateUserRequest struct {
	Name  string `json:"Name"`
	Email string `json:"Email"`
	Pass  string `json:"Pass"`
}

func CreateUser(c echo.Context) error {
	req := &CreateUserRequest{}
	// 構造体にリクエスト情報を結び付ける（バインド）
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	// データベース接続
	dbInstance := db.GetDB()

	// 変数userにリクエストデータを格納
	user := entity.User{
		ID:    0, // Userはautoincrementを指定しているため、自動的にIDが振られます
		Name:  req.Name,
		Email: req.Email,
		Pass:  req.Pass,
	}
	if err := dbInstance.Create(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	// レスポンス返却
	if err := c.JSON(http.StatusCreated, "ok!"); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return nil

}
