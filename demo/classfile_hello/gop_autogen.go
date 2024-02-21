package main

import "github.com/goplus/yap"

type hello struct {
	yap.App
}
//line demo/classfile_hello/hello_yap.gox:1
func (this *hello) MainEntry() {
//line demo/classfile_hello/hello_yap.gox:1:1
	this.Get("/p/info", func(ctx *yap.Context) {
//line demo/classfile_hello/hello_yap.gox:2:1
		ctx.Json__1(map[string]string{"info": "Test info"})
	})
//line demo/classfile_hello/hello_yap.gox:6:1
	this.Get("/p/:id", func(ctx *yap.Context) {
//line demo/classfile_hello/hello_yap.gox:7:1
		ctx.Json__1(map[string]string{"id": ctx.Param("id")})
	})
//line demo/classfile_hello/hello_yap.gox:11:1
	this.Post("/p/:id", func(ctx *yap.Context) {
//line demo/classfile_hello/hello_yap.gox:12:1
		ctx.Json__1(map[string]string{"id": ctx.Param("id")})
	})
//line demo/classfile_hello/hello_yap.gox:16:1
	this.Put("/p/:id", func(ctx *yap.Context) {
//line demo/classfile_hello/hello_yap.gox:17:1
		ctx.Json__1(map[string]string{"id": ctx.Param("id")})
	})
//line demo/classfile_hello/hello_yap.gox:21:1
	this.Delete("/p/:id", func(ctx *yap.Context) {
//line demo/classfile_hello/hello_yap.gox:22:1
		ctx.Json__1(map[string]string{"id": ctx.Param("id")})
	})
//line demo/classfile_hello/hello_yap.gox:26:1
	this.Handle("/", func(ctx *yap.Context) {
//line demo/classfile_hello/hello_yap.gox:27:1
		ctx.Html__1(`<html><body>Hello, <a href="/p/123">Yap</a>!</body></html>`)
	})
//line demo/classfile_hello/hello_yap.gox:30:1
	this.Run(":8080")
}
func main() {
	yap.Gopt_App_Main(new(hello))
}
