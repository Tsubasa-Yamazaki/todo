package service

import (
	"github.com/Tsubasa-Yamazaki/todo/api/domain/model"
	"github.com/Tsubasa-Yamazaki/todo/api/domain/repo"
	"github.com/labstack/echo"
)

type TodoFile interface {
	Todos() ([]model.TodoFile, error)
	CreateTodo(c echo.Context) error
	DeleteTodo(c echo.Context) error
}

type todoFile struct {
	repo repo.TodoFile
}

func NewTodoFile(r repo.TodoFile) TodoFile {
	return &todoFile{repo: r}
}

func (t *todoFile) Todos() ([]model.TodoFile, error) {
	return t.repo.Todos()
}

func (t *todoFile) CreateTodo(c echo.Context) error {
	return t.repo.CreateTodo(c)
}

func (t *todoFile) DeleteTodo(c echo.Context) error {
	return t.repo.DeleteTodo(c)
}
