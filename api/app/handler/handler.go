package handler

import (
	"net/http"

	"github.com/Tsubasa-Yamazaki/todo/api/domain/service"
	"github.com/labstack/echo"
)

type TodoFileHandler interface {
	Todos(c echo.Context) error
	CreateTodo(c echo.Context) error
	DeleteTodo(c echo.Context) error
}

type todoFileHandler struct {
	todoFile service.TodoFile
}

func NewTodoFileHandler(todoFile service.TodoFile) TodoFileHandler {
	return &todoFileHandler{todoFile: todoFile}
}

// 一覧
func (t *todoFileHandler) Todos(c echo.Context) error {
	todos, err := t.todoFile.Todos()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, todos)
}

// 登録
func (t *todoFileHandler) CreateTodo(c echo.Context) error {
	err := t.todoFile.CreateTodo(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, "OK")
}

// 削除
func (t *todoFileHandler) DeleteTodo(c echo.Context) error {
	err := t.todoFile.DeleteTodo(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, "OK")
}
