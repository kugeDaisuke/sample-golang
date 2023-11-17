package handler

import (
	"net/http"
	"sampleApi/db"
	"sampleApi/entity"

	"github.com/labstack/echo/v4"
)

func GetItem(c echo.Context) error {

	// 取得するデータの型
	var item entity.Item
	// データベースに接続
	dbInstance := db.GetDB()

	// Where関数で取得するデータの絞り込みをする
	// データの取得と例外処理
	if err := dbInstance.Where("id = ?", 1).First(&item).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.String(http.StatusOK, item.Name)
}
