package main

import (
	"os"

	"github.com/goplus/yap"
)

func main() {
	y := yap.New(os.DirFS("."))
	y.LoadTemplate("*.*", "**/*.html")
	y.GET("/p/:id", func(ctx *yap.Context) {
		ctx.YAP(200, "article", yap.H{
			"id": ctx.Param("id"),
		})
	})
	y.GET("/api/user", func(ctx *yap.Context) {
		println("user get")
	})
	y.GET("/api/user/unlink", func(ctx *yap.Context) {
		println("unlink")
	})
	y.GET("/api/user/:id/articles", func(ctx *yap.Context) {
		println("article get")
	})
	y.GET("/api/user/:id/articles/:actionId", func(ctx *yap.Context) {
		println("article actionId")
	})
	y.GET("/api/user/:id/articles/:actionId/ok", func(ctx *yap.Context) {
		println("article actionId ok")
	})
	y.GET("/api/user/:id/:actionId/ok", func(ctx *yap.Context) {
		println("article id actionId ok")
	})

	y.Run(":8080")
}
