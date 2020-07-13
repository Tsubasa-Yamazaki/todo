package main

// 具体的な業務ロジックを実装するパッケージ

import (
	"log"

	"github.com/Tsubasa-Yamazaki/todo/api/app/handler"
	"github.com/Tsubasa-Yamazaki/todo/api/domain/service"
	"github.com/Tsubasa-Yamazaki/todo/api/infra"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"xorm.io/core"
)

var engine *xorm.Engine

func main() {
	var err error

	engine, err = xorm.NewEngine("mysql", "root:0111@/todo283")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer engine.Close()
	engine.SetMapper(core.GonicMapper{})

	log.Println("Open the 283todo.")

	todoFileRepo := infra.NewTodo(engine)
	todoFileService := service.NewTodoFile(todoFileRepo)
	h := handler.NewTodoFileHandler(todoFileService)

	e := echo.New()

	e.Use(middleware.CORS())

	e.POST("/todos", h.CreateTodo)
	e.GET("/todos", h.Todos)
	e.DELETE("/todos", h.DeleteTodo)

	e.Logger.Fatal(e.Start(":8080"))
}
