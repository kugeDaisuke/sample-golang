package handler

import (
	"net/http"
	"sampleApi/db"
	"sampleApi/entity"

	"github.com/labstack/echo/v4"
)

type DeleteUserRequest struct {
	ID int `json:"ID"`
}

func DeleteUser(c echo.Context) error {
	req := &DeleteUserRequest{}

	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	// データベース接続
	dbInstance := db.GetDB()

	// 変数userにリクエストデータを格納
	user := entity.User{
		ID: req.ID,
	}
	if err := dbInstance.Delete(&entity.User{}, user.ID).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	// レスポンス返却
	if err := c.JSON(http.StatusCreated, "ok!"); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return nil
}
