package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"xorm.io/core"
)

// var db *sql.DB

var engine *xorm.Engine

type todo struct {
	ID         int    `json:"id"`
	Importance string `json:"importance"`
	Task       string `json:"task"`
	Deadline   string `json:"deadline"`
}

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

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.POST("/todos", createTodo)
	e.GET("/todos", todos)
	e.DELETE("/todos", deleteTodo)

	e.Logger.Fatal(e.Start(":8080"))

}

// 表示
func todos(c echo.Context) error {
	todos := []todo{}
	var err error
	err = engine.Table("todolist").Find(&todos)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, todos)
}

func createTodo(c echo.Context) error {
	var body todo
	err := c.Bind(&body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	_, err = engine.Table("todolist").Insert(body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, "OK")
}

func deleteTodo(c echo.Context) error {
	var body todo
	var err error
	deleteId := c.QueryParam("id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	_, err = engine.Table("todolist").Where("id = ?", deleteId).Delete(body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, "OK")
}
