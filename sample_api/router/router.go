package router

import (
	"net/http"
	"sampleApi/handler"

	"github.com/labstack/echo/v4"
)

// 新しいリクエスト先が必要な場合はここに追加する
func Route(e *echo.Echo) {
	/*
		httpメソッドは主にGET,POST,PUT,DELETEが使用されます。
		GET: DBからデータを取得する
		POST: DBのデータを更新する
		PUT: DBにデータを新規作成する
		DELTE: DBからデータを削除する

		golangのechoではe.〇〇の〇〇にGET～DELETEを入れることで
		クライアント（ブラウザからの操作）からDBに対して
		どのような処理を実行させたいか決定します。
	*/
	e.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "Accessible")
	})
	// e.Logger.Fatal(e.Start(":1324"))

	// クライアントからhttp://{バックエンドのURL}:{バックエンドのポート}/helloが
	// リクエストされたときに、handler/hello.goのHello関数を実行する
	e.GET("/hello", handler.Hello)
	// e.Logger.Fatal(e.Start(":1324"))

	// サーバーを起動して、以下のコマンドをターミナルから実行しましょう
	// curl -X PUT -H "Content-Type: application/json" -d "{\"Name\": \"smart phone\"}" localhost:1324/items
	e.PUT("/items", handler.CreateItem)
	// e.Logger.Fatal(e.Start(":1324"))

	// 課題１ handlerディレクトリにget_item.goを作成し、ブラウザに文字列"smart phone"を表示させる
	e.GET("/item", handler.GetItem)
	// e.Logger.Fatal(e.Start(":1324"))

	// 課題２ entityに新しく構造体を作成して、任意のテーブルを作成し、データを登録する
	// curl -X PUT -H "Content-Type: application/json" -d "{\"Name\": \"一郎\", \"Email\": \"s1111@enecom-th.jp\", \"Pass\": \"pass1\"}" localhost:1324/users
	e.PUT("/users", handler.CreateUser)
	// e.Logger.Fatal(e.Start(":1324"))

	// 課題３ 課題２で作成したテーブルにデータを入力し、そのデータを更新する処理を作成する
	// curl -X POST -H "Content-Type: application/json" -d "{\"Name\": \"三郎\", \"Email\": \"s3333@enecom-th.jp\", \"Pass\": \"pass3\"}" localhost:1324/userup
	e.POST("/userup", handler.UpdateUser)
	// e.Logger.Fatal(e.Start(":1324"))
}
