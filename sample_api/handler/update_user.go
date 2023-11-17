package handler

import (
	"net/http"
	"sampleApi/db"
	"sampleApi/entity"

	"github.com/labstack/echo/v4"
)

type UpdateUserRequest struct {
	Name  string `json:"Name"`
	Email string `json:"Email"`
	Pass  string `json:"Pass"`
}

func UpdateUser(c echo.Context) error {
	req := &UpdateUserRequest{}

	// 構造体にリクエスト情報を結び付ける（バインド）
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	dbInstance := db.GetDB()

	// 変数userにリクエストデータを格納
	user := entity.User{
		ID:    0, // Userはautoincrementを指定しているため、自動的にIDが振られます
		Name:  req.Name,
		Email: req.Email,
		Pass:  req.Pass,
	}
	if err := dbInstance.Model(&user).Where("id=?", 1).Updates(entity.User{Name: user.Name, Email: user.Email, Pass: user.Pass}).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	// レスポンス返却
	if err := c.JSON(http.StatusCreated, "ok!"); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return nil
}
