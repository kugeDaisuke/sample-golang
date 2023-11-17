package handler

import (
	"sampleApi/db"
	"sampleApi/entity"

	"github.com/labstack/echo/v4"
)

func UpdateUser(c echo.Context) error {
	var user entity.User

	dbInstance := db.GetDB()
	dbInstance.Model(&user).Updates(entity.User{Name: "三郎", Email: "s3333@enecom-th.jp", Pass: "pass3"})
	return nil
}
