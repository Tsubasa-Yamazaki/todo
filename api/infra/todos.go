package infra

// repoのインターフェイスの実装

import (
	"github.com/Tsubasa-Yamazaki/todo/api/domain/model"
	"github.com/Tsubasa-Yamazaki/todo/api/domain/repo"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/labstack/echo"
)

type todoFile struct {
	engine *xorm.Engine
}

func NewTodo(engine *xorm.Engine) repo.TodoFile {
	return &todoFile{engine: engine}
}

func (t *todoFile) Todos() ([]model.TodoFile, error) {
	todos := []model.TodoFile{}
	var err error
	err = t.engine.Table("todolist").Find(&todos)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (t *todoFile) CreateTodo(c echo.Context) error {
	var body model.TodoFile
	err := c.Bind(&body)
	if err != nil {
		return err
	}
	_, err = t.engine.Table("todolist").Insert(body)
	return err
}

func (t *todoFile) DeleteTodo(c echo.Context) error {
	var body model.TodoFile
	var err error
	deleteId := c.QueryParam("id")
	if err != nil {
		return err
	}
	_, err = t.engine.Table("todolist").Where("id = ?", deleteId).Delete(body)
	return err
}
